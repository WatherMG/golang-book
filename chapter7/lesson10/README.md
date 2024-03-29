# 7.10 Декларации типов

`Декларация типа (type assertion)` представляет собой операцию, применяемую к значению-интерфейсу. Синтаксически она
выглядит, как `x.(T)`, где `x` - выражение интерфейсного типа, а `T` является типом,
именуемым `декларируемым (asserted)`. Декларация типа проверяет, что динамический тип его операнда `x` соответствует
декларируемому типу.

Существуют две возможности. Во-первых, если декларируемый тип `T` является конкретным типом, то декларация
типов проверяет, `идентичен` ли динамический тип `x` к типу `T`. Если эта проверка завершается успешно, результатом
декларации типа является динамическое значение `x`, типом которого, конечно, является `T`. Другими
словами, `декларация типа` для конкретного типа извлекает конкретное значение из своего операнда. Если проверка
неудачна, то создается паника, например:

``` go
var w io.Writer
w = os.Stdout
f := w.(*os.File)      // Успешно: f == os.Stdout
c := w.(*bytes.Buffer) // panic: интерфейс хранит *os.File, а не *bytes.Buffer
```

Во-вторых, если вместо этого декларированный тип `T` является типом интерфейса, то декларация типов проверяет,
`соответствует` ли динамический тип `x` интерфейсу `T`. Если проверка прошла успешно, динамическое значение не
извлекается, значением результата остается значение интерфейса с тем же типом и значениями компонентов, но сам результат
имеет тип интерфейса `T`. Другими словами, декларация типа для типа интерфейса изменяет тип выражения, делая доступным
иной (обычно больший) набор методов, но сохраняет динамический тип и значения компонентов внутри значения интерфейса.

После первой декларации типа, показанной ниже, и `w`, и `rw` хранят `os.Stdout`, так что каждая их этих переменных имеет
динамический тип `*os.File`, но `w`, которая является `io.Writer`, демонстрирует только метод `Write`, в то время
как `rw` имеет еще и метод `Read`:

``` go
var w io.Writer
w = os.Stdout
rw := w.(io.ReadWriter) // Успех: *os.File имеет методы Read и Write
w = new(ByteCounter)
rw = w.(io.ReadWriter)  // panic: *ByteCounter не имеет метода Read
```

Независимо от того, какой тип был декларирован, если операнд представляет собой интерфейсное значение `nil`, декларация
типа является неудачной. Декларация типа для менее ограничивающего типа интерфейса (с меньшим количеством методов)
требуется редко так как ведет себя так же, как присваивание, за исключением случая `nil`:

``` go
w = rw             // io.ReadWriter присваиваем io.Writer
w = rw.(io.Writer) // Ошибка только при rw == nil
```

Часто мы не уверены в динамическом типе значения интерфейса и хотим проверить, не является ли он некоторым определенным
типом. Если декларация типа появляется в присваивании, в котором ожидаются два результата, как, например, в приведенных
ниже объяснениях, паника при неудаче операции не происходит. Вместо этого возвращается дополнительный второй результат
булева типа, указывающий успех или неудачу операции:

``` go
var w io.Writer = os.Stdout
f, ok := w.(*os.File)      // Успех:   ok, f == os.Stdout
b, ok := w.(*bytes.Buffer) // Неудача: !ok, b == nil
```

Второй результат присваивается переменной с именем `ok`. Если операция не удалась, `ok` получает значение `false`, а
первый результат равен нулевому значению декларированного типа, который в данном примере представляет собой нулевое
значение `*bytes.Buffer`.

Результат `ok` часто немедленно используется для принятия решения о последующих действиях. Расширенная форма инструкции
`if` позволяет сделать это достаточно компактно:

``` go
if f, ok := w.(*os.File); ok {
  // ... использование f...
}
```

Когда операнд декларации типа представляет собой переменную, вместо другого имени для новой локальной переменной иногда
повторно используется исходное имя, затеняющее оригинал:

``` go
if w, ok := w.(*os.File); ok {
  // ... использование w...
}
```

## Выводы:

* `Декларация типа (type assertion)` - это операция, которая проверяет, что значение интерфейса имеет определенный тип.
  Синтаксически она выглядит, как `x.(T)`;
* Если декларация типа успешна, то результатом операции будет значение этого типа. Если нет, то операция вызовет панику;
* Декларация типа может использоваться для извлечения значения из интерфейса;
* Если значение интерфейса равно `nil`, то декларация типа не будет успешной;
* Можно использовать декларацию типа с двумя результатами, чтобы избежать паники. В этом случае второй результат будет
  иметь тип `bool` и указывать на успех или неудачу операции. `f, ok := w.(*os.File)`;
* Если декларация типа используется в присваивании с двумя результатами, то второй результат можно использовать для
  принятия решения о последующих действиях. `if f, ok := w.(*os.File); ok {...}`;
* Иногда имя переменной может быть повторно использовано в декларации типа, чтобы затенить (переприсвоить) оригинальную
  переменную. `w, ok := w.(*os.File)`;
* В общем, декларация типа - это способ проверить и извлечь значение из интерфейса в Go. Она может использоваться для
  обработки ошибок и принятия решений в зависимости от типа значения интерфейса.