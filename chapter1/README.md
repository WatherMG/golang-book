# Выводы к главе:

<!-- TOC -->

* [1.1 Краткий обзор основных компонентов GO](#11-краткий-обзор-основных-компонентов-go)
* [1.2 Аргументы командной строки](#12-аргументы-командной-строки)
* [1.3 Поиск повторяющихся строк](#13-поиск-повторяющихся-строк)
* [1.4 Анимированные GIF-изображения](#14-анимированные-gif-изображения)
* [1.5 Выборка URL](#15-выборка-url)
* [1.6. Параллельная выборка URL](#16-параллельная-выборка-url)
* [1.7. Веб-сервер](#17-веб-сервер)
* [1.8. Некоторые мелочи](#18-некоторые-мелочи)

<!-- TOC -->
____

## 1.1 Краткий обзор основных компонентов GO

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

____

## 1.2 Аргументы командной строки

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

____

## 1.3 Поиск повторяющихся строк

* Для условия инструкции `if` скобки не нужны, но для тела инструкции фигурные скобки `обязательны`. Может существовать
  необязательная часть `else`;
* `Карта` (словарь) содержит набор пар `“ключ-значение”` и обеспечивает `константное` время выполнения операций
  хранения, извлечения или проверки наличия элемента;
* `Ключ` может быть любого типа, если значения этого типа `можно сравнить` при помощи оператора `==`. Значение может
  быть любого типа;
* Если в `карте` (словаре) нет нужного ключа, выражение в правой части присваивает нулевое значение типа новому
  элементу;
* `Цикл по диапазону` по `карте` (словарю) возвращает ключ и значение элемента для конкретного ключа. Порядок обхода
  карты (словаря) не определен;
* Элементы карты (словаря) не отсортированы;
* Пакет `bufio` помогает сделать ввод и вывод эффективным и удобным. Тип `Scanner` считывает входные данные и разбивает
  их на строки или слова;
* Функция `Scan` возвращает `true`, если строка считана и доступна, и `false`, если входные данные исчерпаны;
* Функция `fmt.Printf()` выполняет форматированный вывод на основе списка выражений. Первый аргумент - строка формата,
  которая указывает, как должны быть отформатированы последующие аргументы;
* Функция `Printf()` имеет множество преобразований для форматирования вывода;
* Строки могут содержать управляющие последовательности символов, такие как `\n` для новой строки и `\t` для табуляции;
* По умолчанию `Printf()` не содержит символ новой строки. Функции с окончанием `f` используют такие же правила
  форматирования как `fmt.Printf()`. Функции с окончанием `ln` форматируют аргументы как `%v` и добавляют `\n`;
* Функция `os.Open()` возвращает `открытый файл` и значение типа `error`. Если ошибка равна `nil`, файл открыт успешно;
  Если ошибка не равна `nil`, что-то пошло не так;
* Функции и другие объекты уровня пакета могут быть объявлены в любом порядке;
* Карта (словарь) является `ссылкой на структуру данных`, созданную функцией `make`. При передаче карты в функцию
  передается `копия ссылки` на объект карты, но сам объект карты не копируется. Изменения в переданной карте будут
  отражены на исходной карте.

____

## 1.4 Анимированные GIF-изображения

* Объявление `const` дает имена константам - значениям, которые `неизменны` во время компиляции. Константа может быть
  числом, строкой или булевым значением;
* Выражения вида `[]color.Color{...}` и `gif.GIF{...}` являются `составными литералами` - это компактная запись
  для `создания экземпляра составных типов` Go из последовательности значений элементов;
* Тип `gif.GIF` - структурный тип. Структура представляет собой группу значений, именуемых полями, собранных в один
  объект;
* {x: 1, y: 2} - структурный литерал - это способ создания и инициализации экземпляра структуры в одной строке кода;
* `{LoopCount: nframes}` создает значение структуры, поле которого устанавливается равным `nframes`. Все прочие поля
  имеют нулевое значение своих типов. Обращение к отдельным полям структуры выполняется с помощью записи с точкой.

____

## 1.5 Выборка URL

* Функция `http.Get` из пакета `net/http` выполняет HTTP-запрос и при отсутствии ошибок возвращает результат в
  структуру. Поле `Body` этой структуры содержит ответ сервера в виде потока, доступного для чтения;
* Метод `io.ReadAll` из пакета `io` считывает поток, результат сохраняется в переменную;
* Поток `Body` нужно закрыть для предотвращения утечки ресурсов.

____

## 1.6. Параллельная выборка URL

* `Горутина` - параллельное выполнение функции;
* `Канал` является механизмом связи, который позволяет одной горутине `передавать значения определенного типа` другой
  горутине. Функция `main` выполняется в горутине, а инструкция `go` создает дополнительные горутины;
* Функция `io.Copy()` считывает тело ответа и игнорирует его, записывая в выходной поток `io.Discard`. `Copy` возвращает
  количество байтов и информацию о произошедших ошибках;
* Когда одна горутина пытается отправить или получить информацию по каналу, она блокируется, пока другая горутина
  пытается выполнить те же действия. После передачи информации обе горутины продолжают работу;
* Весь вывод осуществляется функцией main, что гарантирует, что вывод каждой горутины будет обработан как единое целое
  без чередования вывода при завершении двух горутин в один и тот же момент времени.

____

## 1.7. Веб-сервер

* Запрос представлен структурой типа `http.Request`, которая содержит ряд связанных полей, включая `URL` входящего
  запроса. Полученный запрос передается функции-обработчику, которая извлекает компонент пути из URL запроса и
  отправляет его обратно в качестве ответа с помощью `fmt.Fprintf()`;
* Чтобы запустить сервер в фоновом режиме в ОС Linux и Mac, нужно ввести символ `&`. Чтобы завершить процесс, нужно
  найти его с помощью команды `ps -ef | grep server1.go` и выполнить команду `kill PID`. В ОС Windows символ амперсанда
  не нужен;
* Сервер запускает обработчик для каждого входящего запроса в `отдельной горутине`. Однако, если два параллельных
  запроса попытаются обновить переменную в один и тот же момент времени, может возникнуть ошибка под
  названием `состояние гонки ( race condition)`;
* `Состояние гонки (race condition)` - это ошибка, которая может возникнуть в многопоточных или параллельных программах.
  Она происходит, когда две или более горутин `пытаются обновить переменную в один и тот же момент времени`. В
  результате значение переменной может быть увеличено не согласованно и стать непредсказуемым. Чтобы избежать этой
  проблемы, нужно гарантировать, что доступ к переменной получает не более одной горутины одновременно. Для этого каждый
  доступ к переменной должен быть окружен вызовами `mu.Lock()` и `mu.Unlock()`;
* Go разрешает простым инструкциям предшествовать условию `if`, что особенно полезно при обработке ошибок. Это делает
  код короче и уменьшает область видимости переменной err, что является хорошей
  практикой. `if err := r.ParseForm(); err != nil {}`.

____

## 1.8. Некоторые мелочи

* Инструкция `switch` представляет собой инструкцию множественного ветвления. Результат вызова функции или значения
  сравнивается со значением в каждой части `case`;
* Значения проверяются `сверху вниз`, и при первом найденном совпадении выполняется соответствующий код. Необязательный
  вариант `default` выполняется, если `нет совпадений`;
* Инструкция `switch` может обойтись и `без операнда` и просто перечислять различные инструкции `case`, каждая из
  которых представляет собой логическое выражение. Такая инструкция называется `переключатель без тегов`;
* Инструкции `break` и `continue` модифицируют поток управления. Инструкция `break` передает управление следующей
  инструкции после наиболее глубоко вложенной инструкции `for`, `switch` или `select`. Инструкция `continue` заставляет
  наиболее глубоко вложенный цикл `for` начать очередную итерацию;
* Инструкции могут иметь `метки`, так что `break` и `continue` могут на них ссылаться. Имеется инструкция `goto`, она
  предназначена для машинно-генерируемого кода;
* Объявление `type` позволяет присваивать имена существующим типам. Структурные типы зачастую длинны, поэтому они почти
  всегда именованны;
* Go предоставляет `указатели` - `значения, содержащие адреса переменных`. Оператор `&` дает адрес переменной, а
  оператор `*` позволяет получить значение переменной, на которую указывает указатель. Однако, арифметики указателей в
  Go нет;* Метод представляет собой функцию, связанную с именованным типом. Методы могут быть связаны почти с любым
  именованным типом;
* `Интерфейс` представляет собой `абстрактные типы`, которые позволяют рассматривать различные конкретные типы
  одинаково, на основе имеющихся у них методов;
* Go поставляется с обширной стандартной библиотекой полезных пакетов. Программирование часто состоит в использовании
  существующих пакетов;
* Хороший стиль требует написания комментария перед объявлением каждой функции, описывающим ее поведение. Это важно для
  инструментов `go doc` и `godoc`, которые используются для поиска и отображения документации. Для многострочных
  комментариев можно использовать запись `/*...*/`.