# 8.4.3 Однонаправленные каналы

По мере роста программ крупные функции естественным образом распадаются на более мелкие фрагменты. В нашем предыдущем
примере использованы три горутины, соединяющиеся с помощью двух каналов, которые представляют собой локальные
переменные функции `main`. Программа естественным образом делится на три функции:

``` go
func counter(out chan int) 
func squarer(out, in chan int) 
func printer(in chan int)
```

Функция `squarer`, находящаяся в средине конвейера, получает два параметра — входной и выходной каналы. Оба они имеют
один и тот же тип, но используются в разных целях: `in` работает только на получение, a `out` — только на отправление.
Имена `in` и `out` выражают это предназначение, но, тем не менее, ничто не мешает функции `squarer` отправлять данные в
поток `in` или получать их из потока `out`.

Эта ситуация достаточно типична. Когда канал передается в качестве параметра функции, это почти всегда делается с тем
намерением, что он должен использоваться исключительно для отправления или исключительно для получения данных.

Для документирования этого намерения и предотвращения неверного применения система типов Go предоставляет тип
однонаправленного канала, который обеспечивает только одну операцию — отправление или получение. Тип `chan<-int`
представляет собой канал `int`, предназначенный только для отправления, который позволяет отправлять данные, но не
получать их. Напротив, тип `<-chan int`, канал `int`, предназначенный только для получения, позволяет получать, но не
отправлять данные. (Положение стрелки `<-` относительно ключевого слова `chan` является мнемоническим.) Нарушения
применения таких каналов обнаруживаются во время компиляции.

Поскольку операция `close` утверждает, что в канал больше не будет отправления данных, вызывать ее может только
отправляющая горутина; по этой причине **попытка закрыть канал только для получения приводит к ошибке времени
компиляции**.

Давайте еще раз рассмотрим конвейер, возводящий числа в квадрат, на этот раз — **с применением однонаправленных
каналов** (см. pipeline3.go):

``` go
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
```

Вызов `counter(naturals)` неявно преобразует `naturals`, значение типа `chan int`, в тип параметра, `chan<-int`.
Вызов `printer(squares)` выполняет аналогичное неявное преобразование в `<-chan int`. Преобразования двунаправленных
каналов в однонаправленные разрешается в любом присваивании. Однако обратное неверно: если у вас есть значение
однонаправленного типа, такого как `chan<-int`, то нет способа получить из него значение типа `chan int`, которое
ссылается на структуру данных того же канала.

## Выводы:

* Каналы в Go могут быть `двунаправленными` (для отправления и получения данных) или `однонаправленными` (только для
  отправления или только для получения данных);
  ``` go
  chan int   // двунаправленный канал
  chan<- int // канал только для отправления данных
  <-chan int // канал только для получения данных 
  ```
* Однонаправленные каналы используются для документирования намерений разработчика и предотвращения неправильного
  использования каналов в функциях, ограничения доступа к каналам только для чтения или только для записи, что может
  повысить безопасность, уменьшить сложность кода и улучшить производительность в некоторых случаях.
* Положение стрелки `<-` относительно ключевого слова `chan` является мнемоническим. Нарушения использования
  однонаправленных каналов обнаруживаются на этапе компиляции, что позволяет предотвратить ошибки в реализации функций;
* Операция `Close` для каналов должна вызываться только отправляющей горутиной, потому что она утверждает, что больше не
  будет отправления данных в канал. Попытка закрыть канал только для получения приводит к ошибке времени компиляции;
* Двунаправленные каналы могут быть неявно преобразованы в однонаправленные, но обратное преобразование невозможно;

``` go
naturals := make(chan int)
go counter(naturals) // naturals преобразуется в chan<- int
go printer(squares) // squares преобразуется в <-chan int
```

* Использование однонаправленных каналов и правильная организация горутин позволяют создавать четкую и понятную
  структуру программы, что улучшает ее читаемость и облегчает поддержку.