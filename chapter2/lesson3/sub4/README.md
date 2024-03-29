# 2.3.4 Время жизни переменных

**_Время жизни переменной_** - это интервал времени выполнения программы, в течение которого она существует. Время жизни
переменной **уровня пакета** равно времени работы всей программы.
**Локальные** же переменные, напротив, имеют **динамическое время жизни**: новый экземпляр создается всякий раз, когда
выполняется оператор объявления, и переменная живет до тех пор, пока она не станет **_недоступной_**, после чего
выделенная для нее память может быть использована повторно (освобождается). Параметры и результаты функций являются
локальными переменными - они создаются всякий раз, когда вызывается их функция.

Например, во фрагменте программы Lissajous переменная `t` создается в начале цикла `for`, а новые переменные
`x` и `y` создаются на каждой итерации цикла:
https://github.com/WatherMG/golang-book/blob/eed1d3f10c639d5dae4f96bfb3624111115703f5/chapter1%20-%20Tutorial/lesson4%20-%20%D0%90%D0%BD%D0%B8%D0%BC%D0%B8%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%BD%D1%8B%D0%B5%20GIF-%D0%B8%D0%B7%D0%BE%D0%B1%D1%80%D0%B0%D0%B6%D0%B5%D0%BD%D0%B8%D1%8F/lissajous.go#L52-L55

### Как сборщик мусора узнает, что память, выделенная для хранения переменной, может быть освобождена?

Go использует алгоритм сборки мусора, который основан на отслеживании ссылок на переменные. Если нет доступных ссылок на
переменную, то сборщик мусора предполагает, что переменная больше не нужна и освобождает память, выделенную для нее. Это
помогает предотвратить утечку памяти в программе и переиспользовать память.

Поскольку время жизни переменной определяется только ее доступностью, локальная переменная может пережить итерацию
охватывающего цикла. Она может продолжать существовать даже после возврата значения из охватывающей функции.

Для **локальных переменных** компилятор может выбрать куда их поместить в **стек или кучу**, но этот выбор не зависит от
того, была ли переменная объявлена с помощью `var` или `new`.

``` go
package main

var global *int

func f() {
	var x int
	x = 1
	global = &x
}

func g() {
	y := new(int)
	*y = 1
}
```

В функции `f()` память для переменной `x` должна быть выделена в **куче**, потому что она остается доступной с помощью
переменной `global` после возвращения значения из `f()`, несмотря на ее объявление как локальной переменной. `x`
_сбегает от_ `f()`.
Когда выполняется возврат из функции `g()` переменная `*y` становится недоступной, и выделенная для нее память может
быть использована повторно. Поскольку `*y` _не сбегает от_ `g()`, компилятор может безопасно выделить память для `*y` в
стеке, несмотря на то, что память выделяется с помощью функции `new`. В любом случае понятие `побега` не является тем,
о чем нужно беспокоиться, чтобы писать правильный код. Тем не менее его хорошо иметь в виду во время оптимизации
производительности, поскольку каждая _сбегающая_ переменная требует дополнительной памяти.

Сборка мусора представляет огромную помощь при написании корректных программ, но не освобождает от бремени размышлений о
памяти. Не нужно явно выделять и освобождать память, но, чтобы писать эффективные программы, все же необходимо знать о
времени жизни переменных. Например, сохранение ненужных указателей на короткоживущие объекты внутри долгоживущих
объектов, в особенности в глобальных переменных, мешает сборщику мусора освобождать память, выделенную для
короткоживущих объектов.

## Выводы:

* Время жизни переменной - это период, в течение которого она существует в программе;
* Переменные уровня пакета имеют время жизни, равное времени работы всей программы;
* Локальные переменные имеют `динамическое время жизни` и создаются каждый раз при выполнении оператора объявления;
* Параметры и результаты функций являются `локальными переменными` и создаются при каждом вызове функции;
* Go использует алгоритм сборки мусора, основанный на отслеживании `ссылок на переменные`, что позволяет освободить
  память, выделенную для ненужных переменных;
* Локальная переменная может продолжать существовать даже после возврата значения из охватывающей функции:
  * ``` go
    func f() *int {
        i := 10
        return &i // возвращаем ссылку на локальную переменную i
    }
    ```
* Компилятор может выбрать, где разместить локальные переменные `(в стеке или куче)`, независимо от того, была ли
  переменная объявлена с помощью `var` или `new`;
* Необходимо учитывать время жизни переменных для написания эффективных программ, избегая утечек памяти и ненужного
  использования ресурсов;
* Сборка мусора облегчает написание корректных программ, но не освобождает от бремени размышлений о памяти и времени
  жизни переменных.