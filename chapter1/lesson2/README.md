# 1.2 Аргументы командной строки

Пакет `os` предоставляет функции и различные значения для работы с ОС кроссплатформенным образом. Аргументы cli доступны
в программе в виде переменной с именем `Args`, которая является частью пакета `os`. Таким образом, ее имя в любом месте
за пределами пакета `os` выглядит как `os.Args`.

Переменная `os.Args` представляет собой **_срез_** строк. Срезы являются фундаментальным понятием в Go.
Индексация срезов в Go использует **полуоткрытые интервалы**, которые включают первый индекс, но исключают последний,
потому что это упрощает логику.

`os.Args[0]` - имя самой команды, остальные элементы представляют собой аргументы переданные программе,
когда началось выполнение.

Выражение вида `s[m:n]` - дает срез, который указывает на элементы от `m` до `n-1`. Можно использовать такую
конструкцию:
`os.Args[1:len(os.Args)]`. Если опущено значение `m` или `n`, используются значения по умолчанию - `0` или `len(s)`
соответственно, так что можно сократить запись нужного среза до `os.Args[1:]`.

Однострочное комментирование - `//`. Многострочное - `/**/`.

Переменная может быть инициализирована `var s = 'string'`. Если переменная не инициализированная явно (указан тип), она
неявно инициализируется _нулевым значением_ соответствующего типа (`0` для числовых, пустая строка `""` для строк).

В примере из этого урока, объявление переменных неявно инициализирует строки `s` и `sep`.

Для чисел Go предоставляет обычные арифметические и логические операторы. Однако, при применении к строкам оператор `+`
выполняет [конкатенацию](/ "Соединение двух и более строк в одну") значений (соединение двух и более строк в одну).

`s += sep + os.Args[i]` - представляет собой инструкцию присваивания, которая выполняет конкатенацию старого значения
`s` с `sep` и `os.Args[i]` и присваивает новое значение переменной `s`. Это эквивалентно
выражению `s = s + sep + os.Args[i]`

Оператор `+=` является **_присваивающим оператором_**. Каждый арифметический и логический оператор наподобие `+` или `*`
имеет соответствующий присваивающий оператор.

Оператор `:=` являются частью краткого объявления переменной, благодаря ей, можно объявить одну или несколько переменных
и назначить им соответствующие типы на основе значения инициализатора.

Оператор `i++` инкремента\декремента `i--`, добавляет\отнимает единицу. Эта запись эквивалентна `i += 1`
или `i = i + 1`.
Это инструкции, а не выражения, как в большинстве языков семейства `C`, поэтому запись `j = i++` является некорректной.
Эти операторы могут быть только постфиксными.

Цикл `for` является единственной инструкцией цикла в Go. Он имеет ряд разновидностей, одна из которых показана в
примере:

`for инициализация; условие; последействие {
// ноль или несколько инструкций
}`

Вокруг трех компонентов цикла `for` скобки не используются. Фигурные скобки - обязательны, причем открывающая фигурная
скобка **обязана** находится в той же строке, что и инструкция последействие.

Необязательная инструкция **_инициализации_** выполняется до начала работы цикла. Если она имеется, она обязана быть
_**простой инструкцией**_, т.е. кратким объявлением переменной, инструкцией инкремента или присваивания, или вызовом
функции.

**_Условие_** - логическое выражение, вычисляется в начале каждой итерации цикла. Если `true` - выполняется тело цикла.

**_Последействие_** - выполняется после тела цикла, после чего заново вычисляется условие. Когда условие равно `false` -
цикл завершается.

Любая из компонентов цикла может быть опущена, при этом можно не ставить точки с запятыми.

**_Традиционный цикл while:_**
`for condition {
//
}`

**_Традиционный бесконечный цикл:_**
`for {
//
}`

Если условие опущено полностью в любой из разновидностей цикла, мы получим бесконечный цикл, который должен завершится
другим путем, например с помощью инструкции `break` или `return`.

Еще одна разновидность цикла `for` выполняет итерации для **_диапазона_** значений для типа данных наподобие строки или
среза. (см. echo2.go)

В каждой итерации цикла, **_range_** возвращает пару значений: **_индекс_** и **_значение элемента_** с этим индексом.
В данном примере индекс не нужен, но синтаксис цикла `range` требует, чтобы работа велась и с индексом. Одно из решений
заключается в том, чтобы присваивать значение индекса временной переменной с очевидным именем, например `temp` и
игнорировать его.
Однако Go не допускает наличия неиспользуемых локальных переменных, так что этот способ приведет к ошибке компиляции.

Решение заключается в применении **_пустого идентификатора (blank identifier)_** с именем `_ (символ подчеркивания)`.
Пустой идентификатор может использоваться везде, где синтаксис требует имя переменной, но логике программы он не нужен.

В файле `echo2.go` для объявления и инициализации `s` и `sep` используется **_краткое объявление переменной_**, но,
можно объявить эти переменные и отдельно.
Эти объявления эквивалентны:

``` go
s := "" // может использоватся только внутри функции, но не для переменных уровня пакета.
var s string // инициализация по умолчанию. (для строки "")
var s = "" // используется редко, в основном при объявлении нескольких переменных.
var s string = "" // явное указание типа переменной, лишнее, т.к. тип совпадает с начальным значением
// последний пример может использоваться, когда тип переменной и значение разные.
```

На практике обычно следует использовать первые два варианта.

Оператор `+=` создает новую строку путем конкатенации старой строки, символа пробела и очередного аргумента, а затем
присваивает новую строку переменной `s`. Старое содержимое строки `s` больше не используется, поэтому оно будет
обработано **_сборщиком мусора_**

Если объем данных (передаваемых аргументов) большой, пример из `echo1.go` и `echo2.go` может быть дорогостоящим
решением.
Более простое и эффективное решение - использование функции `Join()` из пакета `strings` (см. echo3.go)

Если нам не нужно беспокоится о формате и нужно увидеть только значения, например, для отладки, можно использовать
функцию `Println()`

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:])
}

```

Вывод этой инструкции похож на вывод, полученный в версии с применением `strings.Join()`, но с квадратными скобками
вокруг. Таким образом может быть выведен любой срез

## Выводы:

* Пакет `os` предоставляет функции для работы с операционной системой кроссплатформенно;
* Аргументы командной строки доступны в программе через переменную `os.Args`, которая является `срезом` строк;
* `os.Args[0]` - имя команды, остальные элементы - аргументы программы;
* Срезы в Go используют полуоткрытые интервалы, включая первый индекс и исключая последний;
* Выражение `s[m:n]` дает срез элементов от `m` до `n-1`. Если опущено значение `m` или `n`, используются значения по
  умолчанию `0` или `len(s)` соответственно;
* Однострочные комментарии начинаются с `//`, многострочные - с `/*...*/`;
* Неинициализированные переменные имеют нулевое значение соответствующего типа;
* Оператор `+` для строк выполняет конкатенацию;
* Оператор `+=` является присваивающим оператором, который выполняет конкатенацию и присваивание нового значения
  переменной;
* Оператор `:=` используется для `краткого объявления переменных` с присваиванием им типа на основе значения
  инициализатора;
* Операторы i++ и i-- (инкремент\декремент) увеличивают или уменьшают значение переменной на единицу;
* Цикл `for` - единственная инструкция цикла в Go с различными вариантами использования;
* Скобки не используются вокруг компонентов цикла for, фигурные скобки обязательны;
* `Инициализация` выполняется до начала работы цикла и может быть опущена;
* `Условие` вычисляется в начале каждой итерации цикла. Если оно равно `true`, выполняется тело цикла;
* `Последействие` выполняется после тела цикла, после чего заново вычисляется условие. Если условие равно `false`, цикл
  завершается;
* Любая из компонентов цикла может быть опущена без использования точек с запятой;
* `Традиционный` цикл `while` выглядит как `for condition { }`
* `Бесконечный` цикл выглядит как `for { }` и должен быть завершен другим способом, например с помощью
  инструкции `break` или `return`;
* Цикл `for` с `range` выполняет итерации для диапазона значений для типа данных, например строки или среза;
* В каждой итерации цикла `range` возвращает `пару значений`: `индекс` и `значение элемента` с этим индексом;
* Если индекс не нужен, можно использовать пустой идентификатор `_`, который может использоваться везде, где требуется
  имя переменной, но логике программы он не нужен;
* Функция `Join` из пакета `strings` объединяет элементы с указанным разделителем;
* Для вывода значений без форматирования можно использовать функцию `Println`.