# Golang Book: Решения упражнений и выводы по книге

Этот репозиторий используется в образовательных целях. Он содержит решения упражнений и выводы по книге "Язык
программирования Go" | Керниган Брайан, Донован Алан А. А..

Структура проекта соответствует структуре книги, каждый раздел содержит текст из книги, примеры кода, решенные
упражнения и выводы. В некоторых частях текста и примеров есть исправления, так как в текущей версии Go некоторые методы
устарели и были заменены или удалены. Для некоторых упражнений в комментариях добавлены примечания о проблемах, с
которыми я столкнулся во время их решения, и как я решил эти проблемы.

Если вам понравился этот репозиторий и вы сочли его полезным, пожалуйста, поставьте звезду ⭐️. Это поможет другим людям
найти этот проект и воспользоваться им.

<details>
<summary>English</summary>

This repository is used for educational purposes. It contains exercise solutions and conclusions from the book Go
Programming Language | Kernighan Brian, Donovan Alan A. A.

The structure of the project follows the structure of the book, each section contains text from the book, code examples,
solved exercises and conclusions. There are corrections in some parts of the text and examples, as the current version
of Go some methods are obsolete and have been replaced or removed. For some of the exercises, notes have been added in
the comments about problems I encountered while solving them and how I solved those problems.

If you enjoyed this repository and found it useful, please put a star at ⭐️. This will help other people
find this project and use it.

</details>

## Структура проекта по главам:

Все ссылки ведут на конкретную главу книги.

All links lead to a specific chapter of the book.
<details>
<summary>
Глава 1. Учебник
</summary>

* [x] [Глава 1. Учебник](./chapter1)
	* [x] [1.1. Hello, World](./chapter1/lesson1)
	* [x] [1.2. Аргументы командной строки](./chapter1/lesson2)
	* [x] [1.3. Поиск повторяющихся строк](./chapter1/lesson3)
	* [x] [1.4. Анимированные GIF-изображения](./chapter1/lesson4)
	* [x] [1.5. Выборка URL](./chapter1/lesson5)
	* [x] [1.6. Параллельная выборка URL](./chapter1/lesson6)
	* [x] [1.7. Веб-сервер](./chapter1/lesson7)
	* [x] [1.8. Некоторые мелочи](./chapter1/lesson8)

</details>

<details>
<summary>
Глава 2. Структура программы
</summary>

* [x] [Глава 2. Структура программы](./chapter2)
	* [x] [2.1. Имена](./chapter2/lesson1)
	* [x] [2.2. Объявления](./chapter2/lesson2)
	* [x] [2.3. Переменные](./chapter2/lesson3)
		* [x] [2.3.1 Краткое объявление переменной](./chapter2/lesson3/sub1)
		* [x] [2.3.2 Указатели](./chapter2/lesson3/sub2)
		* [x] [2.3.3 Функция new](./chapter2/lesson3/sub3)
		* [x] [2.3.4 Время жизни переменных](./chapter2/lesson3/sub4)
	* [x] [2.4. Присваивания](./chapter2/lesson4)
		* [x] [2.4.1 Присваивание кортежу](./chapter2/lesson4/sub1)
		* [x] [2.4.2 Присваиваемость](./chapter2/lesson4/sub2)
	* [x] [2.5. Объявления типов](./chapter2/lesson5)
	* [x] [2.6. Пакеты и файлы](./chapter2/lesson6)
		* [x] [2.6.1 Импорт](./chapter2/lesson6/sub1)
		* [x] [2.6.2 Инициализация пакетов](./chapter2/lesson6/sub2)
	* [x] [2.7. Область видимости](./chapter2/lesson7)

</details>

<details>
<summary>
Глава 3. Фундаментальные типы данных
</summary>

* [x] [Глава 3. Фундаментальные типы данных](./chapter3)
	* [x] [3.1. Целые числа](./chapter3/lesson1)
	* [x] [3.2. Числа с плавающей точкой](./chapter3/lesson2)
	* [x] [3.3. Комплексные числа](./chapter3/lesson3)
	* [x] [3.4. Булевы значения](./chapter3/lesson4)
	* [x] [3.5. Строки](./chapter3/lesson5)
		* [x] [3.5.1 Строковые литералы](./chapter3/lesson5/sub1)
		* [x] [3.5.2 Unicode](./chapter3/lesson5/sub2)
		* [x] [3.5.3 UTF-8](./chapter3/lesson5/sub3)
		* [x] [3.5.4 Строки и байтовые срезы](./chapter3/lesson5/sub4)
		* [x] [3.5.5 Преобразования между строками и числами](./chapter3/lesson5/sub5)
	* [x] [3.6. Константы](./chapter3/lesson6)
		* [x] [3.6.1 Генератор констант iota](./chapter3/lesson6/sub1)
		* [x] [3.6.2 Нетипизированные константы](./chapter3/lesson6/sub2)

</details>
<details>
<summary>
Глава 4. Составные типы
</summary>

* [x] [Глава 4. Составные типы](./chapter4)
	* [x] [4.1. Массивы](./chapter4/lesson1)
	* [x] [4.2. Срезы](./chapter4/lesson2)
		* [x] [4.2.1 Функция append](./chapter4/lesson2/sub1)
		* [x] [4.2.2 Работа со срезами "на лету"](./chapter4/lesson2/sub2)
	* [x] [4.3. Отображения](./chapter4/lesson3)
	* [x] [4.4. Структуры](./chapter4/lesson4)
		* [x] [4.4.1 Структурные литералы](./chapter4/lesson4/sub1)
		* [x] [4.4.2 Сравнение структур](./chapter4/lesson4/sub2)
		* [x] [4.4.3 Встраивание структур и анонимные поля](./chapter4/lesson4/sub3)
	* [x] [4.5. JSON](./chapter4/lesson5)
	* [x] [4.6. Текстовые и HTML-шаблоны](./chapter4/lesson6)

</details>
<details>
<summary>
Глава 5. Функции
</summary>

* [x] [Глава 5. Функции](./chapter5)
	* [x] [5.1. Объявления функций](./chapter5/lesson1)
	* [x] [5.2. Рекурсия](./chapter5/lesson2)
	* [x] [5.3. Множественные возвращаемые значения](./chapter5/lesson3)
	* [x] [5.4. Ошибки](./chapter5/lesson4)
		* [x] [5.4.1 Стратегии обработки ошибок](./chapter5/lesson4/sub1)
		* [x] [5.4.2 Конец файла (EOF)](./chapter5/lesson4/sub2)
	* [x] [5.5. Значения-функции](./chapter5/lesson5)
	* [x] [5.6. Анонимные функции](./chapter5/lesson6)
		* [x] [5.6.1 Предупреждение о захвате переменных итераций](./chapter5/lesson6/sub1)
	* [x] [5.7. Вариативные функции](./chapter5/lesson7)
	* [x] [5.8. Отложенные вызовы функций](./chapter5/lesson8)
	* [x] [5.9. Аварийная ситуация](./chapter5/lesson9)
	* [x] [5.10. Восстановление](./chapter5/lesson10)

</details>
<details>
<summary>
Глава б. Методы
</summary>

* [x] [Глава б. Методы](./chapter6)
	* [x] [6.1. Объявления методов](./chapter6/lesson1)
	* [x] [6.2. Методы с указателем в роли получателя](./chapter6/lesson2)
		* [x] [6.2.1 Значение nil является корректным получателем](./chapter6/lesson2/sub1)
	* [x] [6.3. Создание типов путем встраивания структур](./chapter6/lesson3)
	* [x] [6.4. Значения-методы и выражения-методы](./chapter6/lesson4)
	* [x] [6.5. Пример: тип битового вектора](./chapter6/lesson5)
	* [x] [6.6. Инкапсуляция](./chapter6/lesson6)

</details>
<details>
<summary>
Глава 7. Интерфейсы
</summary>

* [x] [Глава 7. Интерфейсы](./chapter7)
	* [x] [7.1. Интерфейсы как контракты](./chapter7/lesson1)
	* [x] [7.2. Типы интерфейсов](./chapter7/lesson2)
	* [x] [7.3. Соответствие интерфейсу](./chapter7/lesson3)
	* [x] [7.4. Анализ флагов с помощью flag.Value](./chapter7/lesson4)
	* [x] [7.5. Значения интерфейсов](./chapter7/lesson5)
		* [x] [7.5.1 Осторожно: интерфейс, содержащий нулевой указатель не является нулевым](./chapter7/lesson5/sub1)
	* [x] [7.6. Сортировка с помощью sort.Interface](./chapter7/lesson6)
	* [x] [7.7. Интерфейс http.Handler](./chapter7/lesson7)
	* [x] [7.8. Интерфейс error](./chapter7/lesson8)
	* [x] [7.9. Пример: вычислитель выражения](./chapter7/lesson9)
	* [x] [7.10. Декларации типов](./chapter7/lesson10)
	* [x] [7.11. Распознавание ошибок с помощью деклараций типов](./chapter7/lesson11)
	* [x] [7.12. Запрос поведения с помощью деклараций типов](./chapter7/lesson12)
	* [x] [7.13. Выбор типа](./chapter7/lesson13)
	* [x] [7.14. Пример: XML-декодирование на основе лексем](./chapter7/lesson14)
	* [x] [7.15. Несколько советов](./chapter7/lesson15)

</details>
<details>
<summary>
Глава 8. Горутины и каналы
</summary>

* [x] [Глава 8. Горутины и каналы](./chapter8)
	* [x] [8.1. Горутины](./chapter8/lesson1)
	* [x] [8.2. Пример: параллельный сервер часов](./chapter8/lesson2)
	* [x] [8.3. Пример: параллельный эхо-сервер](./chapter8/lesson3)
	* [x] [8.4. Каналы](./chapter8/lesson4)
		* [x] [8.4.1 Небуферизованные каналы](./chapter8/lesson4/sub1)
		* [x] [8.4.2 Конвейеры](./chapter8/lesson4/sub2)
		* [x] [8.4.3 Однонаправленные каналы](./chapter8/lesson4/sub3)
		* [x] [8.4.4 Буферизованные каналы](./chapter8/lesson4/sub4)
	* [x] [8.5. Параллельные циклы](./chapter8/lesson5)
	* [x] [8.6. Пример: параллельный веб-сканер](./chapter8/lesson6)
	* [x] [8.7. Мультиплексирование с помощью select](./chapter8/lesson7)
	* [x] [8.8. Пример: параллельный обход каталога](./chapter8/lesson8)
	* [x] [8.9. Отмена](./chapter8/lesson9)
	* [x] [8.10. Пример: чат-сервер](./chapter8/lesson10)

</details>
<details>
<summary>
Глава 9. Параллельность и совместно используемые переменные
</summary>

* [x] [Глава 9. Параллельность и совместно используемые переменные](./chapter9)
	* [x] [9.1. Состояния гонки](./chapter9/lesson1)
	* [x] [9.2. Взаимные исключения: sync.Mutex](./chapter9/lesson2)
	* [x] [9.3. Мьютексы чтения/записи: sync.RWMutex](./chapter9/lesson3)
	* [x] [9.4. Синхронизация памяти](./chapter9/lesson4)
	* [x] [9.5. Отложенная инициализация: sync.Once](./chapter9/lesson5)
	* [x] [9.6. Детектор гонки](./chapter9/lesson6)
	* [x] [9.7. Пример: параллельный неблокирующий кеш](./chapter9/lesson7)
	* [x] [9.8. Go-подпрограммы и потоки](./chapter9/lesson8)
		* [x] [9.8.1 Растущие стеки](./chapter9/lesson8/sub1)
		* [x] [9.8.2 Планирование go-подпрограмм](./chapter9/lesson8/sub2)
		* [x] [9.8.3 GOMAXPROCS](./chapter9/lesson8/sub3)
		* [x] [9.8.4 Go-подпрограммы не имеют идентификации](./chapter9/lesson8/sub4)

</details>
<details>
<summary>
Глава 10. Пакеты и инструменты Go
</summary>

* [x] [Глава 10. Пакеты и инструменты Go](./chapter10)
	* [x] [10.1. Введение](./chapter10/lesson1)
	* [x] [10.2. Пути импорта](./chapter10/lesson2)
	* [x] [10.3. Объявление пакета](./chapter10/lesson3)
	* [x] [10.4. Объявления импорта](./chapter10/lesson4)
	* [x] [10.5. Пустой импорт](./chapter10/lesson5)
	* [x] [10.6. Пакеты и именование](./chapter10/lesson6)
	* [x] [10.7. Инструментарий Go](./chapter10/lesson7)
		* [x] [10.7.1 Организация рабочего пространства](./chapter10/lesson7/sub1)
		* [x] [10.7.2 Загрузка пакетов](./chapter10/lesson7/sub2)
		* [x] [10.7.3 Построение пакетов](./chapter10/lesson7/sub3)
		* [x] [10.7.4 Документирование пакетов](./chapter10/lesson7/sub4)
		* [x] [10.7.5 Внутренние пакеты](./chapter10/lesson7/sub5)
		* [x] [10.7.6 Запрашиваемые пакеты](./chapter10/lesson7/sub6)

</details>
<details>
<summary>
Глава 11. Тестирование
</summary>

* [x] [Глава 11. Тестирование](./chapter11)
	* [x] [11.1. Инструмент gotest](./chapter11/lesson1)
	* [x] [11.2. Тестовые функции](./chapter11/lesson2)
		* [x] [11.2.1 Рандомизированное тестирование](./chapter11/lesson2/sub1)
		* [x] [11.2.2 Тестирование команд](./chapter11/lesson2/sub2)
		* [x] [11.2.3 Тестирование белого ящика](./chapter11/lesson2/sub3)
		* [x] [11.2.4 Внешние тестовые пакеты](./chapter11/lesson2/sub4)
		* [x] [11.2.5 Написание эффективных тестов](./chapter11/lesson2/sub5)
		* [x] [11.2.6 Избегайте хрупких тестов](./chapter11/lesson2/sub6)
	* [x] [11.3. Охват](./chapter11/lesson3)
	* [x] [11.4. Функции производительности](./chapter11/lesson4)
	* [x] [11.5. Профилирование](./chapter11/lesson5)
	* [x] [11.6. Функции-примеры](./chapter11/lesson6)

</details>
<details>
<summary>
Глава 12. Рефлексия
</summary>

* [ ] [Глава 12. Рефлексия](./chapter12)
	* [ ] [12.1. Почему рефлексия?](./chapter12/lesson1)
	* [ ] [12.2. reflect.Туре и reflect.Value](./chapter12/lesson2)
	* [ ] [12.3. Рекурсивный вывод значения](./chapter12/lesson3)
	* [ ] [12.4. Пример: кодирование S-выражений](./chapter12/lesson4)
	* [ ] [12.5. Установка переменных с помощью reflect.Value](./chapter12/lesson5)
	* [ ] [12.6. Пример: декодирование S-выражений](./chapter12/lesson6)
	* [ ] [12.7. Доступ к дескрипторам полей структур](./chapter12/lesson7)
	* [ ] [12.8. Вывод методов типа](./chapter12/lesson8)
	* [ ] [12.9. Предостережение](./chapter12/lesson9)

</details>
<details>
<summary>
Глава 13. Низкоуровневое программирование
</summary>

* [ ] [Глава 13. Низкоуровневое программирование](./chapter13)
	* [ ] [13.1. unsafe.Sizeof, Alignof и Offsetof](./chapter13/lesson1)
	* [ ] [13.2. unsafe.Pointer](./chapter13/lesson2)
	* [ ] [13.3. Пример: глубокое равенство](./chapter13/lesson3)
	* [ ] [13.4. Вызов кода "С" с помощью сgo](./chapter13/lesson4)
	* [ ] [13.5. Еще одно предостережение](./chapter13/lesson5)

</details>