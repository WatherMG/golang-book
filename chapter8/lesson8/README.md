## 8.8. Пример: параллельный обход каталога

В этом разделе мы создадим программу, которая сообщает об использовании дискового пространства одним или несколькими
каталогами, указанными в командной строке наподобие команды Unix `du`. Большую часть работы выполняет показанная ниже
функция `walkDir`, которая перечисляет все записи каталога `dir` с помощью вспомогательной функции `dirents` (см.
du1.go):

``` go
func main() {
	// Определяем начальные каталоги.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// Обход дерева файлов.
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// Вывод информации.
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskusage(nfiles, nbytes)
}

func printDiskusage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.3f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir рекурсивно обходит дерево файлов с корнем в dir,
// и отправляет размер каждого найденного файла в fileSizes.
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}

}

// dirents возвращает записи каталога dir
func dirents(dir string) []os.FileInfo {
	entries, err := os.ReadDir(dir) // ioutils.ReadDir deprecated
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	infos := make([]os.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Fprintf(os.Stderr, "du1: %v\n", err)
			return nil
		}
		infos = append(infos, info)
	}
	return infos
}
```

Функция `ioutil.ReadDir` возвращает срез значений `os.Filelnfo` — ту же информацию, которую функция `os.Stat` возвращает
для одного файла. Для каждого подкаталога `walkDir` рекурсивно вызывает сама себя, а для каждого файла `walkDir`
отправляет сообщение в канал `fileSizes`. Сообщение представляет собой размер файла в байтах.

Показанная ниже функция main использует две горутины. Фоновая горутина вызывает `walkDir` для каждого
каталога, указанного в командной строке, и в конце закрывает канал `fileSizes`. Главная горутина вычисляет сумму
размеров файлов, которые получает из канала, и в конце работы выводит итоговый результат.

Эта программа надолго замирает перед тем, как вывести результат:

``` shell 
28171 files  250.104 GB
```

Программа будет выглядеть лучше, если будет информировать нас о ходе своей работы. Однако простое перемещение вызова
`printDiskUsage` в цикл приведет к выводу тысяч строк.

Вариант `du`, приведенный ниже, периодически выводит итоговую величину, но только если указан флаг `-v`, поскольку не
все пользователи хотят видеть сообщения о ходе выполнения. Фоновая горутина, которая циклически обходит `roots`,
остается неизменной. Главная же горутина теперь использует **таймер для генерации событий** каждые `500 мс` и
инструкцию `select`, чтобы ожидать либо сообщения с размером файла (в этом случае обновляются итоговые значения), либо
события таймера (в этом случае выводятся текущие итоги). Если не указан флаг `-v`, канал `tick` остается нулевым, и его
вариант в инструкции select становится отключенным (см. du2.go):

``` go
var verbose = flag.Bool("v", false, "вывод промежуточных результатов")

func main() {
	// Определяем начальные каталоги.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// Обход дерева файлов.
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// Периодический вывод результатов
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	// Вывод информации.
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes был закрыт
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskusage(nfiles, nbytes)
		}

	}
	printDiskusage(nfiles, nbytes)
}
```

Поскольку программа больше не использует цикл по диапазону, первый вариант `select` должен явно проверять, не закрыт ли
канал `fileSizes`, используя операцию получения с двумя результатами. Если канал был закрыт, выполнение цикла
прерывается.
Помеченная инструкция `break` обеспечивает выход как из инструкции `select`, так и из цикла; инструкция `break` без
метки
обеспечивала бы выход только из инструкции `select` и начало следующей итерации цикла.
Теперь программа дает нам неторопливый поток обновлений итоговых результатов:

``` shell
22693 files  0.700 GB
44134 files  1.108 GB
76515 files  2.906 GB
97908 files  3.087 GB
128994 files  4.631 GB
148029 files  8.989 GB
```

Однако программа все еще работает слишком долго. Нет причин, по которым все вызовы `walkDir` нельзя выполнять
параллельно, тем самым используя параллелизм для дисковой системы. Показанная ниже третья версия `du` создает новую
горутину
для каждого вызова `walkDir`. Она использует `sync.WaitGroup` (раздел 8.5) для подсчета количества активных
вызовов `walkDir`
и горутину, закрывающую канал `fileSizes`, когда счетчик становится равным нулю (см. du3.go).

``` go
func main() {
	// Определяем начальные каталоги.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// Параллельный обход дерева файлов.
	fileSizes := make(chan int64)
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, fileSizes)
	}
	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	// Периодический вывод результатов
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	// Вывод информации.
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes был закрыт
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskusage(nfiles, nbytes)
		}

	}
	printDiskusage(nfiles, nbytes)
}

func printDiskusage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.3f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir рекурсивно обходит дерево файлов с корнем в dir,
// и отправляет размер каждого найденного файла в fileSizes.
func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, wg, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}
```

Поскольку в пике программа создает многие тысячи горутин, мы должны изменить функцию `dirents` так, чтобы она
использовала подсчитывающий семафор для предотвращения открытия слишком большого количества файлов одновременно, так же,
как мы делали для веб-сканера в разделе 8.6:

``` go
// sema - подсчитывающий семафор для ограничения параллельности.
var sema = make(chan struct{}, runtime.GOMAXPROCS(-1))

// dirents возвращает записи каталога dir
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{} // Захват токена
	defer func() {
		<-sema // Освобождение токена
	}()
	entries, err := os.ReadDir(dir) // ioutils.ReadDir deprecated
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	infos := make([]os.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Fprintf(os.Stderr, "du1: %v\n", err)
			return nil
		}
		infos = append(infos, info)
	}
	return infos
}
```

Эта версия работает в несколько раз быстрее предыдущей, хотя ее ускорение сильно варьируется от системы к системе.

## Выводы:

* Функция `walkDir` обходит все файлы и подкаталоги в каталоге, отправляет размеры файлов в канал и вызывает сама себя для
  подкаталогов;
* Главная горутина считывает данные из канала и выводит результаты, второстепенная горутина вызывает walkDir для каждого
  каталога и закрывает канал после завершения;
* Вариант программы, который выводит промежуточные результаты об использовании диска, если установлен флаг "-v" и
  использует таймер для генерации событий каждые 500 мс;
* Параллелизм с использованием sync.WaitGroup и создание новой горутины для каждого вызова walkDir: Третья версия
  программы использует WaitGroup для подсчета вызовов walkDir, создает новые горутины для каждого вызова walkDir, и
  закрывает канал fileSizes, когда счетчик достигает нуля;
* Можно использовать семафор для ограничения количества одновременно открытых файлов и предотвращения создания слишком
  большого количества горутин, что может привести к исчерпанию системных ресурсов.
