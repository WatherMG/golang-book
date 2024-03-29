# 5.9 Аварийная ситуация `panic`

Система типов Go отлавливает множество ошибок во время компиляции, но другие ошибки, такие как обращение за границами
массива или разыменование нулевого указателя, требуют проверок времени выполнения. Когда среда выполнения Go
обнаруживает эти ошибки, мы сталкиваемся с аварийной ситуацией (`panic`).

Во время типичной `паники` нормальное выполнение программы останавливается, выполняются все отложенные вызовы функций в
go-подпрограмме и программа аварийно завершает работу с записью соответствующего сообщения. Это журнальное сообщение
содержит `**_значение паники_**` - оно обычно представляет собой некоторое `сообщение об ошибке`, и, для каждой
go-подпрограммы, `**_трассировку стека_**`, показывающую состояние стека вызовов функций, которые были активны во время
`паники`. Этого журнального сообщения часто оказывается достаточно, чтобы диагностировать причину проблемы без
повторного запуска программы, поэтому его _всегда следует включать в передаваемый разработчику отчет о найденных
ошибках_.

Не все `паники` возникают из-за ошибок времени выполнения. Встроенная функция `panic` может вызываться непосредственно.
В качестве аргумента она принимает любое значение. Такой вызов - зачастую наилучшее, что вы можете сделать, когда
случается некоторая "невозможная" ситуация, например выполнение достигает `case` в инструкции `switch`, которого
согласно логике программы достичь не может:

``` go
switch s := suit(drawCard()); s {
case "Spades": //...
case "Hearts": //...
case "Diamonds: //...
case "Clubs": //...
default:
panic(fmt.Sprintf("неверная карта %q", s)) // Джокер?
}
```

Хорошей практикой является проверка выполнения предусловий функции, но такие проверки могут легко оказаться избыточными.
Если только вы не можете предоставить более информативное сообщение об ошибке или обнаружить ошибку заранее, нет смысла
в проверке, которую среда выполнения осуществит сама:

``` go
func Reser(x *Buffer) {
    if x == nil {
        panic("x is nil") // Нет смысла!
    }
    x.elements = nil
}
```

Хотя механизм паник в Go напоминает исключения в других языках программирования, ситуации, в которых он используется,
существенно различаются. Так как при этом происходит аварийное завершение программы, этот механизм обычно используется
для грубых ошибок, таких как логическая несогласованность в программе. Прилежные программисты рассматривают любую панику
как доказательство наличия ошибки в программе. В надежной программе "ожидаемые" ошибки, т.е. те, которые возникают в
результате неправильного ввода, неверной конфигурации или сбоя ввода-вывода, должны быть корректно обработаны. Лучше
всего работать с ними с использованием значений `error`.

Рассмотрим функцию `regexp.Compile`, которая компилирует регулярное выражение в эффективную форму для дальнейшего
сопоставления. Она возвращает `error`, если функция вызывается с неправильно сформированным шаблоном, но проверка этой
ошибки является излишней и обременительной, если вызывающая функция знает, что определенный вызов не может быть
неудачным. В таких случаях разумно, чтобы вызывающая функция обработала ошибку с помощью `генерации паники`, так как
ошибка считается невозможной.
Поскольку большинство регулярных выражений представляют собой литералы в исходном тексте программы, пакет `regexp`
предоставляет функцию-оболочку `regexp.MustCompile`, выполняющую такую проверку:

``` go
package regexp

func Compile(expr string) (*Regexp, error) { /* ... */ }

func MustCompile(expr string) *Regexp {
    re, err := Compile(expr)
    if err != nil {
        panic(err)
    }
    return re
}
```

Эта функция-оболочка позволяет клиентам удобно инициализировать переменные уровня пакета скомпилированным регулярным
выражением, как показано ниже:

``` go 
var httpSchemeRE=regexp.MustCompile(`^https?:`) // "http:" или "https:"
```

Конечно, функция `MustCompile` не должна вызываться с недоверенными входными значениями. Префикс `Must` является
распространенным соглашением именования для функций такого рода наподобие `template.Must` в разделе 4.6.

Когда программа сталкивается с `паникой`, все отложенные функции выполняются в порядке, обратном их появлению в коде,
начиная с функции на вершине стека и опускаясь до функции `main`, что демонстрирует приведенная программа
(см. `defer1.go`):

``` go
func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panic при x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
```

При вызове `f(0)` возникает `паника`, приводящая к выполнению трех отложенных вызовов `fmt.Printf`. Затем среда
выполнения прекращает выполнение программы, выводя соответствующее аварийное сообщение и дамп стека в стандартный поток
ошибок (здесь для ясности вывод упрощен):

```
panic: runtime error: integer divide by zero
main.f(0xd8c128?)
        /lesson9/defer1.go:14 +0x113
main.f(0x1)
        /lesson9/defer1.go:16 +0xf5
main.f(0x2)
        /lesson9/defer1.go:16 +0xf5
main.f(0x3)
        /lesson9/defer1.go:16 +0xf5
main.main()
        /lesson9/defer1.go:10 +0x1e
```

Как мы вскоре увидим, функция может восстановиться после аварийной ситуации так, что программа при этом не будет
аварийно завершена.

Для диагностических целей пакет `runtime` позволяет программисту вывести дамп стека с помощью того же механизма.
Откладывая вызов `printStack`в функции `main` (см. `defer2.go`) мы обеспечиваем вывод дополнительного текста в
стандартный поток вывода (вновь упрощенный здесь для ясности):

```
goroutine 1 [running]:
main.printStack()
        /defer2/defer2.go:20 +0x39
main.f(0x12c1c8?)
        /defer2/defer2.go:25 +0x113
main.f(0x1)
        /defer2/defer2.go:27 +0xf5
main.f(0x2)
        /defer2/defer2.go:27 +0xf5
main.f(0x3)
        /defer2/defer2.go:27 +0xf5
main.main()
        /defer2/defer2.go:15 +0x45
```

Читатели, знакомые с исключениями в других языках программирования, могут быть удивлены тем, что функция `runtime.Stack`
позволяет вывести информацию о функциях, которые кажутся уже "развернутыми". Но механизм паники Go запускает отсроченные
функции **_до_** разворачивания стека.

## Выводы:

* Такие ошибки как, обращение к элементу за границами массива или разыменовывание нулевого указателя требуют проверок в
  `runtime`. Когда среда выполнения Go обнаруживает эти ошибки, возникает `паника` (`panic`);
* Во время паники выполнение программы останавливается, выполняются все отложенные вызовы функций в текущей горутине и
  программа аварийно завершает работу с записью соответствующего сообщения;
* Журнальное сообщение `паники` содержит `значение паники` (обычно это сообщение об ошибке) и `трассировку стека` для
  каждой горутины. Трассировка стека показывает состояние стека вызовов функций, которые были активны во время паники;
* Журнальное сообщение паники содержит много информации о возникшей ошибке, и его следует включать в отчет об ошибке,
  чтобы разработчик мог проанализировать и исправить ошибку;
* Не все `паники` возникают в `runtime`. Встроенная функция `panic` может вызываться в коде, и это не обязательно
  связано с ошибками во время выполнения программы. В этом случае паника также приводит к аварийному завершению
  программы.
* В качестве аргумента `паника` принимает любое значение;
* Не всегда лучшее решение вызывать `panic`, когда происходит "невозможная" ситуация. В некоторых случаях более
  целесообразно использовать другие механизмы обработки ошибок, например, возвращать ошибку или использовать специальный
  тип, который позволяет обрабатывать ошибки;
* Если нельзя предоставить более информативное сообщение или обнаружить ошибку заранее, нет смысла в проверке, например,
  значения на `nil` и вызове `panic`, так как среда выполнения все осуществит сама;
* Механизм `паник` существенно различается с исключениями из других ЯП. Он используется для грубых ошибок, таких как
  логическая несогласованность в программе, при этом происходит аварийное завершение программы;
* Если ошибка, возникает по причине некорректного ввода, неверной конфигурации или сбоя ввода-вывода, лучше обработать
  их с использованием типа `error`;
* Префикс `Must` является распространенным соглашением именования для функции такого рода
  как `template.Must`, `regexp.MustCompile` - такие функции не возвращают ошибки и они должны выполниться обязательно;
* Когда программа сталкивается с `паникой`, все отложенные функции выполняются в порядке, обратном их появлению в коде,
  начиная с функции на вершине стека и опускаясь до функции `main`;
* Функция может восстановиться после паники так, что программа при этом не будет аварийно завершена;
* Если в Go функция вызывает `panic`, а затем восстанавливается с помощью функции `recover`, то эта функция возвращает
  значение, переданное в `panic`;
* Для диагностических целей пакет `runtime` позволяет вывести дамп стека;
* Функция `runtime.Stack` позволяет вывести информацию о функциях. Однако механизм `паники` в Go запускает отложенные
  функции **до** разворачивания стека (вывода его в стандартный поток).