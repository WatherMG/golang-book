# 2.3.1. Краткое объявление переменной

В функции может использоваться **_краткое объявление переменной_** - альтернативная форма объявления.
Она имеет вид: `name := expression`, и тип переменной `name` определяется как тип выражения `expression`.

https://github.com/WatherMG/golang-book/blob/f0149d829ef8ac8f4720f21295927b750916c134/chapter1%20-%20Tutorial/lesson4%20-%20%D0%90%D0%BD%D0%B8%D0%BC%D0%B8%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%BD%D1%8B%D0%B5%20GIF-%D0%B8%D0%B7%D0%BE%D0%B1%D1%80%D0%B0%D0%B6%D0%B5%D0%BD%D0%B8%D1%8F/lissajous.go#L46-L48

Для краткости и гибкости краткие объявления переменных используются для объявления и инициализации большинства локальных
переменных. Объявление `var`, как правило, резервируется для локальных переменных, требующих явного указания типа,
который отличается от выражения инициализатора, или когда значение переменной будет присвоено позже, а его начальное
значение не играет роли.

``` go
i := 100
var boiling float64 = 100
var names []string
var err error
var p Point
```

Как и в случае объявления `var`, в одном кратком объявлении можно объявить и инициализировать несколько переменных:

`i, j := 0, 1`

Однако объявления с несколькими выражениями инициализации следует использовать только тогда, когда они могут повысить
удобочитаемость, как, например, в коротких и естественных группах наподобие инициализирующей части цикла `for`.

Объявление переменной это - `:=`, присваивание переменной это - `=`

Объявление переменных не следует путать с **_присваиванием кортежу_**, в котором каждой переменной в левой части
инструкции присваивается значение из правой части:

`i, j = j, i // Обмен значений i и j`

Краткие объявления можно использовать для вызовов функций наподобие `os.Open`, которые возвращают два или больше
значений:

``` go
package main

import (
	"fmt"
	"os"
)

func open() {
	f, err := os.Open("name")
	if err != nil {
		fmt.Println(err)
	}
	// ... использование f ...
	f.Close()
}
```

**_Краткое объявление переменной необязательно объявляет все переменные в своей левой части_**. Если некоторые из них
уже были объявлены в **_том же_** лексическом блоке, то для этих переменных краткие объявления действуют как
**_присваивания_**.

В приведенном ниже коде, первая инструкция объявляет `in` и `err`. Вторая объявляет только `out`, а уже существующей
переменной `err` она присваивает значение:

```
in, err := os.Open(infile)
...
out, err := os.Create(outfile)
```

> Но, краткое объявление переменной должно объявить по крайней мере одну новую переменную, так что приведенный ниже код
> не компилируется:

```
f, err := os.Open(infile)
...
f, err := os.Create(outfile)
```

Чтобы исправить ошибку, во второй инструкции следует использовать обычное присваивание.

Краткое объявление переменной действует как присваивание только для переменных, которые уже были объявлены в том же
лексическом блоке. Объявления во внешнем блоке игнорируются.

## Выводы:

* Краткое объявление переменной использует оператор `:=` и позволяет объявить и инициализировать локальные переменные
  без явного указания их типа (`s := i`);
* Краткое объявление переменной определяет тип переменной на основе типа выражения справа от оператора `:=`;
* Можно использовать краткое объявление переменной для нескольких переменных одновременно, например `i, j := 0, 1`;
* Присваивание значений переменным осуществляется через оператор `=`, например, `i, j = j, i` для обмена значений
  переменных `i` и `j`;
* Если переменная уже была объявлена в том же лексическом блоке, то краткое объявление переменной действует как
  присваивание;
* Краткое объявление переменной должно объявлять хотя бы одну новую переменную, иначе код не будет компилироваться;
* Краткое объявление переменной упрощает работу с локальными переменными, делает код более читабельным и понятным.