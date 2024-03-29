# 1.4 Анимированные GIF-изображения

Объявление `const` дает имена константам - это значения которые неизменны во время компиляции. К ним относятся числовые
параметры циклов, задержки т.д. Как и объявления `var`, `const` могут находиться на уровне пакетов (будут видны во всем
пакете), или в функции (будут видны только в пределах этой функции). Константа может быть числом, строкой или булевым
значением.

Выражения `[]color.Color{...}` и `gif.GIF{...}` являются **_составными литералами (composite literals)_** - это
компактная запись для инстанцирования (это создание экземпляра структуры, интерфейса или типа данных) составных типов
Go из последовательности значений элементов.

В приведенном примере первый из них представляет собой срез, а второй - структуру.

> Тип `gif.GIF` - структурный тип. Структура представляет собой группу значений, именуемых **_полями_**, зачастую
> различных типов, которые собраны в один объект, рассматриваемый как единое целое.

Переменная `anim` - структура типа `gif.GIF`.

> `{x: 1, y: 2}` - структурный литерал (struct literal) - это способ создания и инициализации экземпляра структуры в
> одной строке кода.

`{LoopCount: nframes}` - создает значение структуры, поле которого устанавливается равным `nframes`. Все прочие поля
имеют нулевое значение своих типов.
Обращение к отдельным полям структуры выполняется с помощью записи с точкой `anim.Delay = append(anim.Delay, delay)` -
это явным образом обновляет поле `Delay`

## Выводы:

* Объявление `const` дает имена константам - значениям, которые `неизменны` во время компиляции. Константа может быть
  числом, строкой или булевым значением;
* Выражения вида `[]color.Color{...}` и `gif.GIF{...}` являются `составными литералами` - это компактная запись
  для `создания экземпляра составных типов` Go из последовательности значений элементов;
* Тип `gif.GIF` - структурный тип. Структура представляет собой группу значений, именуемых полями, собранных в один
  объект;
* {x: 1, y: 2} - структурный литерал - это способ создания и инициализации экземпляра структуры в одной строке кода;
* `{LoopCount: nframes}` создает значение структуры, поле которого устанавливается равным `nframes`. Все прочие поля
  имеют нулевое значение своих типов. Обращение к отдельным полям структуры выполняется с помощью записи с точкой.

