# 1.1 Краткий обзор основных компонентов GO

Go - компилируемый язык. Его инструментарий преобразует исходный текст программы, а также библиотеки, от которых он
зависит, в команды машинного кода.

Доступ к инструментарию GO в cli - `go` - имеет множество подкоманд.
`go run main.go` - компилирует исходник (может один или несколько), связывает с библиотеками и выполняет исполняемый
файл.

`go build main.go` - создает бинарный исполняемый файл `main.go`, который можно выполнить в любой момент времени,
без предварительной обработки (компиляции, флагов и т.д)
`./main` - выполнит исполняемый бинарный файл.

Go обрабатывает текст в `unicode` - может обработать текст с любым набором символов.

Код в Go организован в виде пакетов, они похожи на модули и библиотеки в других ЯП. Пакет состоит из одного или
нескольких исходников `*.go` в одном каталоге, которые определяют, какие действия выполняет данный пакет.
Каждый исходник начинается с объявления `package` - оно указывает к какому пакету принадлежит данный файл,
после него следует список других пакетов `import`, которые этот файл импортирует, а после него - объявления программы.

Пакет `fmt` - содержит функции для форматированного вывода и сканирования ввода.
Функция `Println()` - одна из основных функций в этом пакете - она выводит одно или несколько значений,
разделенных пробелами и с добавлением символа новой строки в конце, так, что выводимые значения располагаются
в одной строке.

Пакет `main` - определяет отдельную программу, т.е. выполнимый файл, а не библиотеку. Функция `main` - является
точкой входа в приложение, с нее начинается программа.

**Необходимо импортировать только те пакеты, которые нужны.** Программа не будет компилироваться при отсутствии
пакета, либо при наличии лишнего пакета. Это строгое требование предотвращает накопление ссылок на
неиспользуемые пакеты по мере развития программы.

Объявление функции начинается с ключевого слова `func`, имени функции, списка параметров (может быть пустым),
списка возвращаемых значений (может быть пустым) и тела функции в фигурных скобках.

**Go не требует точек с запятой в конце инструкции или объявления, за исключением случаев, когда две или более
инструкций находятся на одной и той же строке**

Местоположение символов новой строки имеет значение для корректного синтаксического анализа кода Go. Например,
открывающая фигурная скобка `{` должна находится в той же строке, что и конец объявления `func`, но не в
отдельной строке, а в выражении `a + b` символ новой строки разрешен после, но не до оператора `+`

Go занимает жесткую позицию относительно форматирования кода. Инструмент `gofmt` приводит код к стандартному формату,
а подкоманда `fmt` применяет `gofmt` ко всем файлам в указанном пакете или к файлам в текущем каталоге по умолчанию.

Еще один инструмент `goimports`, дополнительно управляет вставкой и удалением объявлений импорта, при необходимости.

`gofmt` сортирует имена пакетов в алфавитном порядке, порядок значений импортируемых пакетов значения не имеет.

## Выводы:

* Go - `компилируемый язык`;
* `go run main.go` - компилирует и запускает исходный код;
* `go build main.go` - создает исполняемый файл, который можно запустить позже;
* Go поддерживает `Unicode` и может обрабатывать текст с любым набором символов;
* Код в Go организован в виде пакетов, которые похожи на модули и библиотеки в других языках программирования;
* Исходный код начинается с объявления пакета и списка импортируемых пакетов;
* Пакет `fmt` содержит функции для форматированного вывода и ввода. Функция `Println()` выводит значения в одну строку с
  пробелами между ними и символом новой строки в конце;
* Пакет `main` определяет отдельную программу, а не библиотеку. Функция `main` - это точка входа в приложение;
* Необходимо импортировать только нужные пакеты, иначе будет ошибка компиляции;
* Функции объявляются с ключевым словом `func`, именем функции, параметрами и возвращаемыми значениями (если есть) и
  телом функции в фигурных скобках;
* Точки с запятой не требуются в конце инструкций или объявлений, кроме случаев, когда несколько инструкций находятся на
  одной строке;
* Местоположение символов новой строки важно для корректного синтаксического анализа кода Go;
* Инструмент `gofmt` приводит код к стандартному формату, а `goimports` управляет импортом пакетов;
* Имена пакетов сортируются в алфавитном порядке с помощью `gofmt`.