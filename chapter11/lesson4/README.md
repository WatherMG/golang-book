# 11.4. Функции производительности

Эталонное тестирование представляет собой практику **измерения производительности программы для фиксированной рабочей
нагрузки**. В Go функция производительности выглядит как тестовая функция, но с префиксом `Benchmark` и
параметром `*testing.B`, который предоставляет большую часть тех же методов, что и `*testing.T`, плюс несколько
дополнительных, связанных с измерением производительности. Он также предоставляет целочисленное поле `N`, которое
указывает количество раз измеряемой операции.

Вот как выглядит функция тестирования для `IsPalindrome`, вызывающая эту функцию в цикле `N` раз:

``` go
import "testing"
func BenchmarkIsPamindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}
```

Мы запускаем ее с помощью приведенной ниже команды. В отличие от тестов, функции производительности **по умолчанию не
запускаются**. Аргумент флага `-bench` выбирает, какие функции производительности будут запущены. Это **регулярное
выражение**, соответствующее именам функций `Benchmark`, со значением по умолчанию, которое не соответствует ни одной из
них. Шаблон приводит к соответствию всем функциям в пакете `word`, но так как есть только одна функция, эквивалентная
запись имеет вид `-bench=IsPalindrome`:

``` shell
$ go test -bench='.'
BenchmarkIsPalindrome-16         5335480               226.1 ns/op
ok      GolangBook/chapter11/lesson4/word2      1.637s
```

Числовой суффикс имени функции производительности (здесь — 16) указывает значение `GOMAXPROCS`, которое является важным
для параллельных тестов производительности.

Отчет говорит, что каждый вызов `IsPalindrome` занимает около **0.2261 микросекунды** при усреднении по **5 335 480
запусков**. Поскольку изначально функции производительности не представляют, сколько времени выполняется операция, она
**делает первоначальные измерения для небольшого значения** `N`, а затем **экстраполирует их для получения значения**,
достаточно большого для выполнения устойчивого измерения времени.

Причина, по которой цикл реализуется `функцией производительности`, а не вызывающим кодом в драйвере теста, заключается
в том, что **функция производительности имеет возможность выполнения любого необходимого одноразового кода настройки вне
цикла, без добавления его к измеряемому времени на каждой итерации**. Если этот код настройки **нарушает результаты
тестирования**, параметр `testing.В` предоставляет методы для **остановки**, **возобновления** и **сброса таймера**, но
они требуются очень редко.

Теперь, когда у нас есть результаты эталонных тестов и тестов, легко попробовать идеи для того, чтобы сделать программу
быстрее. Возможно, наиболее очевидная оптимизация — останавливать второй цикл `IsPalindrome` посредине, чтобы избежать
повторных сравнений:

``` go
n := len(letters)/2
for i := 0; i <n; i++ {
	if letters[i] != letters[len(letters)-1-i] {
		return false
	}
}
return true
```

Но как это часто бывает, очевидная оптимизация не всегда дает ожидаемые выгоды. Данная оптимизация дала лишь 4%
улучшения в одном эксперименте:

``` shell
go test -bench='.'
BenchmarkIsPalindrome-16         5473213               216.2 ns/op
PASS
ok      GolangBook/chapter11/lesson4/word2      1.620s
```

Другая идея заключается в том, чтобы предварительно выделить достаточно большой массив для `letters`, вместо того чтобы
расширять его последовательными вызовами `append`. Объявление `letters` как массива подходящего размера, как здесь, дает
улучшение почти на 52%, и теперь показатели производительности усредняются по 10 544 990 итераций:

``` go
letters := make([]rune, 0, len(s))
for _, r := range s {
	if unicode.IsLetter(r) {
		letters = append(letters, unicode.ToLower(r))
	}
}
```

``` shell
go test -bench='.'
BenchmarkIsPalindrome-16        10544990               117.2 ns/op
PASS
ok      GolangBook/chapter11/lesson4/word2      1.550s
```

Как показывает данный пример, наиболее быстрой программой часто оказывается та, которая делает **наименьшее количество
выделений памяти**. Флаг командной строки `-benchmem` включает в отчет статистику распределений памяти. Давайте сравним
количество выделений до оптимизации:

``` shell
go test -bench='.' -benchmem
BenchmarkIsPalindrome-16         5534888               214.0 ns/op           248 B/op          5 allocs/op
PASS
ok      GolangBook/chapter11/lesson4/word2      1.620s
```

и после нее:

``` shell
go test -bench='.' -benchmem
BenchmarkIsPalindrome-16        10615044               119.3 ns/op           128 B/op          1 allocs/op
PASS
ok      GolangBook/chapter11/lesson4/word2      1.587s
```

Сборка всех распределений памяти в один вызов `make` ликвидировала 75% распределений и вдвое уменьшила количество
расходуемой памяти.

Функции производительности говорят нам об **абсолютном времени, затрачиваемом на некоторую операцию**, но во многих
случаях нас интересует **относительное время выполнения двух различных операций**. Например, если функция тратит `1 мс`
для обработки `1000 элементов`, то сколько времени займет обработка `10 ООО или миллиона`? Такие сравнения показывают
`асимптотический рост времени` работы функции. Другой пример: каков наилучший размер буфера ввода-вывода? Измерение
производительности для разных размеров может помочь нам выбрать наименьший буфер, который обеспечивает
удовлетворительные результаты. Третий пример: какой алгоритм наилучшим образом подходит для выполнения данной работы?
Изучение двух различных алгоритмов для одних и тех же входных данных зачастую может показать сильные и слабые стороны
каждого из них при важных или представительных входных данных.

Сравнительные показатели достигаются с помощью обычного кода. Они обычно принимают вид одной параметризованной функции,
вызываемой из нескольких функций Benchmark с различными значениями, например:

``` go
func benchmark(b *testing.B, size int) { /*...*/ }
func Benchmark10(b *testing.B) { benchmark(b, 10) }
func Benchmark100(b *testing.B) { benchmark(b, 100) }
func Benchmark1000(b *testing.B) { benchmark(b, 1000) }
```

Параметр `size`, который определяет размер входных данных, варьируется для разных функций, но является константой в
пределах каждой функции производительности. Постарайтесь противостоять искушению использовать параметр b.N как размер
входных данных. Если только вы не интерпретируете его как количество итераций для входных данных фиксированного размера,
результаты вашего исследования окажутся бессмысленными.

Шаблоны, выявляемые при сравнении результатов функций производительности, особенно полезны во время разработки
программы, но их не следует игнорировать и при рабочей программе. По мере развития программы могут расти ее входные
данные, она может устанавливаться на новых операционных системах или процессорах с различными характеристиками, и мы
можем повторно использовать эти данные для пересмотра проектных решений.

## Выводы:

* Функции производительности в Go используются для измерения производительности программы на разных входных данных;
* Префикс `Benchmark` и параметр `*testing.B` используются для создания функции производительности;
* Поле `N` в `*testing.B` указывает количество итераций, которое будет выполнена измеряемая операция;
* Измерить производительность можно с помощью команды `go test -bench='.'` , где регулярное выражение после
  флага `-bench` определяет, какие функции `Benchmark` будут запущены;
* Настройка `GOMAXPROCS` влияет на параллельные тесты производительности;
* При разработке быстрой программы, оптимизация часто сводится к минимизации выделения памяти;
* Флаг `-benchmem` используется для отображения статистики распределений памяти;