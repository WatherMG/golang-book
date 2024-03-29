# 2.4 Присваивания

Значение переменной обновляется с помощью оператора присваивания, который в своей простейшей форме содержит переменную,
слева от знака `=`, а справа - выражение.

``` go
x = 1 // Именованная переменная
*p = true // Косвенная переменная
person.name = "Bob" // Поле структуры
count[x] = count[x]*scale // Элемент массива, среза или карты
```

Каждый из арифметических или побитовых бинарных операторов имеет соответствующий **_присваивающий оператор_**,
позволяющий, например, переписать последнюю инструкцию как

``` go
count[x] *= scale
```

Это позволяет не писать лишний раз выражение для переменной. Числовые переменные могут также быть увеличены или
уменьшены с помощью инструкций `++` и `--`.\

``` go
v:=1
v++ // аналог v = v + 1
v-- // аналог v = v - 1
```

## Выводы:

* Значение переменной обновляется с помощью оператора присваивания `=`;
* Оператор присваивания может быть использован для именованных переменных, косвенных переменных, полей структур и
  элементов массива, среза или карты;
* Арифметические и побитовые бинарные операторы имеют соответствующие присваивающие операторы, которые позволяют
  упростить запись и избежать повторения выражений;
* Числовые переменные могут быть увеличены или уменьшены с помощью инструкций `++` и `--`, которые представляют собой
  краткую форму операции сложения или вычитания единицы.