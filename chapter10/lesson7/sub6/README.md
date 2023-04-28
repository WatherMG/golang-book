# 10.7.6. Запрашиваемые пакеты

Инструмент `go list` выводит информацию о доступных пакетах. В своей простейшей форме `go list` проверяет, *
*присутствует ли пакет в рабочем пространстве**, и **выводит его путь импорта**, если это так:

``` shell
$ go list github.com/go-sql-driver/mysql 
github.com/go-sql-driver/mysql
```

Аргумент `go list` может содержать символы `...` которые соответствуют любой подстроке пути импорта пакета. Мы можем
использовать их для перечисления всех пакетов в рабочей области Go:

``` shell
$ go list
archive/tar
archive/zip
bufio
bytes
cmd/addr21ine
cmd/api
... и т.д. ...
```

Или внутри конкретного поддерева:

``` shell
$ go list GolangBook\chapter3\...
GolangBook/chapter3/lesson1
GolangBook/chapter3/lesson2
GolangBook/chapter3/lesson2/ex3.1
GolangBook/chapter3/lesson2/ex3.2
GolangBook/chapter3/lesson2/ex3.3
GolangBook/chapter3/lesson2/ex3.4
GolangBook/chapter3/lesson3
GolangBook/chapter3/lesson3/ex3.5
GolangBook/chapter3/lesson3/ex3.6
GolangBook/chapter3/lesson3/ex3.7
GolangBook/chapter3/lesson3/ex3.8
GolangBook/chapter3/lesson3/ex3.9
GolangBook/chapter3/lesson5/sub4
GolangBook/chapter3/lesson5/sub4/ex3.10
GolangBook/chapter3/lesson5/sub4/ex3.11
GolangBook/chapter3/lesson5/sub4/ex3.12
GolangBook/chapter3/lesson6/sub1
GolangBook/chapter3/lesson6/sub1/ex3.13
```

Или связанных с конкретной темой:

``` shell
$ go list ...xml...
GolangBook/chapter7/lesson14/xmlselect
```

Команда `go list` **получает полные метаданные для каждого пакета**, а не только пути импорта, и делает эту информацию
**доступной для пользователей или других инструментов в различных форматах**. Флаг `-json` заставляет `go list` вывести
всю запись для каждого пакета в формате JSON:

``` shell
$ go list -json .\chapter1\lesson1\...
{
        "Dir": "D:\\Projects\\Golang\\src\\GolangBook\\chapter1\\lesson1",
        "ImportPath": "GolangBook/chapter1/lesson1",
        "Name": "main",
        "Target": "D:\\Projects\\Golang\\bin\\lesson1.exe",
        "Root": "D:\\Projects\\Golang\\src\\GolangBook",
        "Module": {
                "Path": "GolangBook",
                "Main": true,
                "Dir": "D:\\Projects\\Golang\\src\\GolangBook",
                "GoMod": "D:\\Projects\\Golang\\src\\GolangBook\\go.mod",
                "GoVersion": "1.20"
        },
        "Match": [
                "."
        ],
        "Stale": true,
        "StaleReason": "build ID mismatch",
        "GoFiles": [
                "helloworld.go"
        ],
        "Imports": [
                "fmt"
        ],
        "Deps": [
                "errors",
                "fmt",
                "internal/abi",
                "internal/bytealg",
                "internal/coverage/rtcov",
                "internal/cpu",
                "internal/fmtsort",
                "internal/goarch",
                "internal/goexperiment",
                "internal/goos",
                "internal/itoa",
                "internal/oserror",
                "internal/poll",
                "internal/race",
                "internal/reflectlite",
                "internal/safefilepath",
                "internal/syscall/execenv",
                "internal/syscall/windows",
                "internal/syscall/windows/registry",
                "internal/syscall/windows/sysdll",
                "internal/testlog",
                "internal/unsafeheader",
                "io",
                "io/fs",
                "math",
                "math/bits",
                "os",
                "path",
                "reflect",
                "runtime",
                "runtime/internal/atomic",
                "runtime/internal/math",
                "runtime/internal/sys",
                "sort",
                "strconv",
                "sync",
                "sync/atomic",
                "syscall",
                "time",
                "unicode",
                "unicode/utf16",
                "unicode/utf8",
                "unsafe"
        ]
}

```

Флаг `-f` позволяет пользователям **настраивать формат вывода**, используя язык шаблонов пакета `text/template` (раздел
4.6). Следующая команда выводит **транзитивные зависимости пакета** `strconv`, разделенные пробелами:

``` shell
$ go list -f '{{join .Deps " "}}' strconv
errors internal/abi internal/bytealg internal/coverage/rtcov internal/cpu internal/goarch internal/goexperiment internal/goos internal/reflectlite internal/unsafeheader math math/bits runtime runtime/internal/atomic runtime/internal/math runtime/internal/sys runtime/internal/syscall unicode/utf8 unsafe
```

А приведенная ниже команда выводит непосредственные импортирования каждого пакета в поддереве `compress` стандартной
библиотеки:

``` shell
go list -f '{{.ImportPath}} -> {{join .Imports " "}}' compress/...
compress/bzip2 -> bufio io sort
compress/flate -> bufio errors fmt io math math/bits sort strconv sync
compress/gzip -> bufio compress/flate encoding/binary errors fmt hash/crc32 io time
compress/lzw -> bufio errors fmt io
compress/zlib -> bufio compress/flate encoding/binary errors fmt hash hash/adler32 io
```

Команда `go list` полезна как для интерактивных запросов, так и для построения и тестирования сценариев автоматизации.
Мы будем использовать ее снова в разделе 11.2.4. Для получения дополнительной информации, включая набор доступных полей
и их смысл, обратитесь к результату выполнения команды `go help list`.
В этой главе рассмотрены все важные подкоманды инструмента `go`, за исключением одной. В следующей главе мы увидим, как
команда `go test` используется для тестирования программ Go.

## Выводы:

* Инструмент `go list` предоставляет информацию о доступных пакетах в рабочем пространстве;
* `go list` может проверять наличие пакета и выводить его путь импорта, что упрощает поиск и использование пакетов для
  разработчиков;
* Использование символов `...` в аргументе `go list` позволяет перечислять все пакеты в рабочем пространстве, в
  определенном поддереве или связанные с конкретной темой (`...xml...`), упрощая навигацию и организацию кода;
* Команда `go list` предоставляет полные метаданные для каждого пакета, что дает возможность получить более детальную
  информацию для пользователей или других инструментов;
* Флаг `-json` выводит всю информацию о пакете в формате `JSON`, что облегчает автоматическую обработку данных;
* Флаг `-f` позволяет настраивать формат вывода с использованием языка шаблонов `text/template`, что даёт возможность
  выделить определенную информацию о пакете, учитывая индивидуальные требования разработчика;
* Команда `go list` полезна для интерактивных запросов и автоматизации процесса построения и тестирования кода;
* Использование `go list` улучшает организацию и структурирование кода, а также ускоряет поиск нужных пакетов и анализ
  их содержимого.