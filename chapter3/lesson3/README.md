# 3.3 Комплексные числа

Go представляет комплексные числа двух размеров - `complex64` и `complex128`, компонентами которых являются `float32` и
`float64` соответственно.
Встроенная функция `complex` создает комплексное число из действительной и мнимой компонент, а встроенные функции `real`
и `imag` извлекают эти компоненты:

``` go
var x complex128 = complex(1, 2) // 1+2i
var y complex128 = complex(3, 4) // 3+4i
fmt.Println(x*y) // "(5+10i)
fmt.Println(real(x*y)) // "5"
fmt.Println(imag(x*y)) // "10"
```

Если непосредственно за литералом с плавающей точкой или за десятичным целочисленным литералом следует `i`, например,
`3.141592i` или `2i`, такой литерал становится **_мнимым литералом_**, обозначающим комплексное число с нулевым
действительным компонентом:

``` go
fmt.Println(1i * 1i) // "(-1+0i)", i^2 = -1
```

Согласно правилам константной арифметики комплексные константы могут быть прибавлены к другим константам (целочисленным
или с плавающей точкой, действительным или мнимым), позволяя записывать комплексные числа естественным образом, как,
например, `1+2i` или, что, то же самое, `2i + 1`. Показанные выше объявления `x` и `у` могут быть упрощены:

``` go
x := 1 + 2i
y := 3 + 4i
```

Комплексные числа можно проверять на равенство с помощью операторов `==` и `!=`. Два комплексных числа равны тогда, и
только тогда, когда равны их действительные и мнимые части.

Пакет `math/cmplx` предоставляет библиотечные функции для работы с комплексными числами, такие как комплексный
квадратный корень или возведение в степень:

``` go
fmt.Println(cmplx.Sqrt(-1)) // "0+1i"
```

Приведенная далее программа использует арифметику `complex128` для генерации множества Мандельброта.

См. mandelbrot.go

Два вложенных цикла проходят по всем точкам растрового изображения размером 1024ъ1024 в оттенках серого цвета,
представляющего часть комплексной плоскости от -2 до +2.
Программа проверяет, позволяет ли многократное возведение в квадрат и добавление числа, представляющего точку, "сбежать"
из круга радиусом 2. Если позволяет, то данная точка закрашивается оттенком, соответствующим количеству итераций,
потребовавшихся для "побега". Если не позволяет, данное значение принадлежит множеству Мандельброта, и точка остается
черной.
Программа записывает в файл изображение в PNG-кодировке.

## Выводы:

* Go имеет поддержку комплексных чисел с двумя размерами - `complex64` и `complex128`, основанными на `float32`
  и `float64` соответственно;
* Встроенные функции `complex`, `real` и `imag` позволяют создавать комплексные числа и извлекать их действительную и
  мнимую части;
* Мнимые литералы, например `3.141592i` или `2i`, обозначают комплексное число с нулевым действительным компонентом;
* Комплексные числа можно сравнивать на равенство с помощью операторов `==` и  `!=`;
* Пакет `math/cmplx` предоставляет функции для работы с комплексными числами, такие как квадратный корень или возведение
  в степень;
* Арифметика `complex128` может быть использована для генерации множества Мандельброта, что позволяет создавать
  визуализации математических объектов.