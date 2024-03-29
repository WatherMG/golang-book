# 4.2.1 Функция append

Встроенная функция `append` добавляет элементы в срез:

``` go 
var runes []rune
for _, r := range "Hello, World!!!" {
    runes = append(runes, r)
}
fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' 'W' 'o' 'r' 'l' 'd' '!' '!' '!']"
```

Цикл использует функцию `append` для построения среза из десяти рун, закодированных строковым литералом. Эта задача
конкретная задача решается куда проще с помощью встроенного преобразования `[]rune("Hello, World!!!")`.

Функция `append` имеет решающее значение для понимания работы срезов. (см. append.go) - версия функции под названием
`appendInt`, которая специализирована для срезов `[]int`.

Каждый вызов `appendInt` должен проверять, имеет ли срез достаточную емкость для добавления новых элементов в
существующий массив.

Если да, функция расширяет срез, определяя срез большего размера (но все еще в пределах исходного массива),
копирует элемент `y` во вновь добавленную к срезу память и возвращает срез. Входной срез `x` и результирующий срез `z`
при этом используют один и тот же базовый массив.

Если для роста недостаточно памяти (емкости), функции `appendInt` необходимо выделить новый массив, достаточно большой,
чтобы хранить результат, затем скопировать в него значения из `x` и добавить новый элемент `y`. Результат `z` теперь
относится к другому базовому массиву, отличному от базового массива x.

Можно было бы просто скопировать элементы с помощью явного массива, но проще использовать встроенную функцию `copy`,
которая копирует элементы из одного среза в другой того же типа. Ее первый аргумент - целевой срез, второй - исходный
(`dst = src`). Срезы могут относиться к одному и тому же базовому массиву, они могут даже перекрываться.
Функция `copy` возвращает количество фактически скопированных элементов, которое представляет собой меньшую из длин двух
срезов, поэтому опасность выйти за пределы диапазона отсутствует.

Для повышения эффективности новый массив обычно несколько больше, чем минимально необходимо для хранения `x` и `y`.
Увеличение массива путем удвоения его размера при каждом расширении предотвращает чрезмерное количество выделений памяти
и гарантирует, что добавление одного элемента в среднем выполняется за константное время. Код в `append.go`
демонстрирует это.

При каждом добавлении выводится выделенное количество памяти (емкости) для массива:

``` 
0 cap=1  [0]
1 cap=2  [0 1]
2 cap=4  [0 1 2]
3 cap=4  [0 1 2 3]
4 cap=8  [0 1 2 3 4]
5 cap=8  [0 1 2 3 4 5]
6 cap=8  [0 1 2 3 4 5 6]
7 cap=8  [0 1 2 3 4 5 6 7]
8 cap=16 [0 1 2 3 4 5 6 7 8]
9 cap=16 [0 1 2 3 4 5 6 7 8 9]
```

Рассмотрим итерацию `№3`. Срез `x` содержит 3 элемента `[0 1 2]`, но имеет емкость `4`, поэтому имеется еще один
незанятый элемент в конце массива, и функция `appendInt` может добавить элемент `3` без выделения памяти. Результирующий
срез `у` имеет длину и емкость `4` и тот же базовый массив, что и `у` исходного среза `x`, как показано на рисунке:
![img.png](img.png)
На следующей итерации, `i=4`, запаса пустых мест нет, так что функция `appendInt` выделяет новый массив размером `8`,
копирует в него четыре элемента `[0 1 2 3]` из `x` и добавляет значение `i`, равное `4`. Результирующий срез `у` имеет
длину `5` и емкость `8`. Остаток в `3` элемента предназначен для сохранения вставляемых значений на следующих трех
итерациях без необходимости выделять для них память. Срезу `у` и `x` являются представлениями разных массивов (см.
рисунок).
![img_1.png](img_1.png)

Встроенная функция `append` может использовать более сложные стратегии роста, чем функция `appendInt`. Обычно мы не
знаем, приведет ли данный вызов функции `append` к перераспределению памяти, поэтому мы не можем считать ни что исходный
срез относится к тому же массиву, что и результирующий срез, ни что он относится к другому массиву. Аналогично мы не
должны предполагать, что операции над элементами старого срезу будут (или не будут) отражены в новом срезе. Поэтому
обычно выполняется присваивание результата вызова функции `append`  той же переменной среза, которая была передана в
функцию `append`:

``` runes = append(runes, r)```

Обновление переменной среза требуется не только при вызове `append`, но и для любой функции, которая может изменить
длину или емкость среза, или сделать его ссылающимся на другой базовый массив.

Чтобы правильно использовать срезы, важно иметь в виду, что хотя элементы базового массива и доступны косвенно,
указатель среза, его длина и емкость таковыми не являются. Чтобы обновить их, требуется присваивание, такое как
показано выше. В этом смысле срезы не являются "чисто" ссылочными типами, а напоминают составной тип, такой как
приведенная структура:

``` go
type IntSlice struct {
    prt *int
    len, cap int
}
```

Наша функция `appendInt` добавляет в срез единственный элемент, но встроенная функция `append` позволяет добавлять
больше
одного нового элемента или даже целый срез:

``` go
var x []int
x = append(x, 1)
x = append(x, 2, 3)
x = append(x, 4, 5, 6)
x = append(x, x...) // добавление среза x
fmt.Println(x) // "[1 2 3 4 5 6 1 2 3 4 5 6]"
```

С помощью небольшой модификации (код ниже), можно обеспечить поведение нашей функции, совпадающее с поведением
встроенной функции `append`. Троеточие `...` в объявлении функции `appendInt` делает ее **_вариадической_** функцией,
т.е. функцией с переменным числом аргументов. Такое троеточие в вызове `append` во фрагменте исходного текста выше
показывает, как передавать список аргументов из среза.

``` go
func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	// ... Расширяем z до как минимум длины len...
	copy(z[len(x):], y)
	return z
}
```

Логика расширения базового массива z остается неизменной и здесь не показана.

## Выводы:

* Функция `append` в Golang используется для добавления элементов к `срезу (slice)`; это встроенная функция, которая
  имеет ключевое значение для работы со срезами;
* Функция `append` проверяет, имеет ли срез `достаточную емкость (capacity)` для добавления элементов; если да,
  она `расширяет` срез,
  если нет – выделяет новый массив и копирует в него элементы старого среза и новые элементы;
* Для копирования элементов из одного среза в другой можно использовать функцию `copy`; она принимает два аргумента –
  целевой срез и исходный срез, и возвращает количество фактически скопированных элементов;
* Увеличение размера массива при добавлении элементов выполняется путем удвоения его размера, что позволяет снизить
  количество выделений памяти и гарантировать константное время добавления одного элемента в среднем;
* Встроенная функция `append` может использоваться для добавления одного элемента, нескольких элементов или даже другого
  среза; например: `x = append(x, 1, 2, 3)` или `x = append(x, x...)`;
* `x...` - распаковка среза - используется для передачи каждого элемента, например в функцию `append`, в качестве
  отдельного аргумента;
* Для корректной работы со срезами важно помнить, что хотя элементы базового массива доступны косвенно, указатель среза,
  его длина и емкость не являются ссылочными; для обновления их требуется присваивание, например: `runes = append(runes,
  r)`;
* Обновление переменной среза требуется не только при вызове функции `append`, но и для любой функции, которая может
  изменить длину или емкость среза, или сделать его ссылающимся на другой базовый массив;
* `Вариадическая функция` – функция, которая принимает переменное количество аргументов; объявление вариадической
  функции выглядит как функция с аргументом `y ...int`; например, можно использовать вариадическую
  функцию `appendInt(x []int, y ...int)` для имитации встроенной функции append.

