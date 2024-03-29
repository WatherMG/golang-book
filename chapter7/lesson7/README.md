# 7.7 Интерфейс `http.Handler`

В главе 1, мы получили представление о том, как использовать пакет `net/http` для реализации веб-клиента (раздел 1.5) и
сервера (раздел 1.7). В этом разделе мы детальнее рассмотрим API сервера, в основе которого лежит
интерфейс `http.Handler`:

``` go
package http

import "net/http"

type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}
func ListenAndServe(address string, h Handler) error
```

Функция `ListenAndServe` требует адрес сервера, такой как `localhost:8000` и экземпляр интерфейса `Handler`, которому
диспетчеризуются все запросы. Он работает бесконечно, если только не происходит ошибка (или при запуске сервера
происходит сбой, так что он не запускается), и в этом случае функция всегда возвращает ненулевую ошибку.

Представим себе сайт электронного магазина с базой данных, отображающий товары на их цены в долларах. Показанная ниже
программа представляет собой простейшую его реализацию. Она моделирует склад как тип карты, `database`, к которому
присоединен метод `ServeHTTP`, так что он соответствует интерфейсу `http.Handler`. Обработчик обходит карту и выводит
его элементы (см. http1.go):

``` go
func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe(":8000", db))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
```

Если мы запустим сервер:

``` shell
$ go build \http1.go
$ .\httpl
```

И подключимся к нему с помощью программы `fetch` из раздела 1.5 (или, если вам так больше нравится, с помощью
веб-браузера), то получим следующий вывод:

``` shell
shoes: $50.00
socks: $5.00
```

Пока что сервер только перечисляет все товары и отвечает этим на любой запрос, независимо от `URL`. Более реалистичный
сервер определяет несколько различных `URL`, каждый из которых приводит к своему поведению. Давайте будем вызвать
имеющееся поведение при запросе `/list` и добавим еще один запрос `/price`, который сообщает о цене конкретного товара,
указанного в параметре запроса наподобие `/price?item=socks` (см. http2.go):

``` go
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "нет товара: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "нет страницы: %s\n", req.URL)
	}
}
```

Теперь обработчик на основе компонента пути URL `req.URL.Path` решает, какая логика должна быть выполнена. Если
обработчик не распознает путь, он сообщает клиенту об ошибке HTTP клиента путем
вызова `w.WriteHeader(http.StatusNotFound)`. Это должно быть сделано до записи любого текста в `w`.
(Кстати, `http.ResponseWriter` предоставляет собой еще один интерфейс. Он дополняет `io.Writer` методами для
отправки `HTTP-заголовков` ответа.) Мы могли бы с тем же результатом использовать и вспомогательную
функцию `http.Error`:

``` go
msg := fmt.Sprintf("нет страницы: %s\n", req.URL)
http.Error(w, msg, http.StatusNotFound) // 404
```

В случае запроса `/price` вызывается метод `URL.Query`, который выполняет запрос преобразования
параметров `HTTP-запроса` в карту, а точнее - в мультикарту типа `url.Values` (раздел 6.2.1) из пакета `net/url`. Затем
находится первый параметр `item` и выясняется его цена. Если товар не найден, выводится сообщение об ошибке.

Пример сеанса работы с новым сервером:

``` shell
$ go build /ch7/http2
$ go build /ch1/fetch
$ ./http2 
$ ./fetch http://localhost:8000/list
shoes: $50.00
socks: $5.00
$ ./fetch http://localhost:8000/price?item=socks
$5.00
$ ./fetch http://localhost:8000/price?item=shoes
$50.00
$ ./fetch http://localhost:8000/price?item=hat
нет товара: "hat"
$ ./fetch http://localhost:8000/help
нет страницы: /help
```

Очевидно, что мы могли бы продолжать добавлять разные варианты действий, в `ServeHTTP`, но в реальных приложениях
удобнее определить логику для каждого случая в виде отдельной функции или метода. Кроме того, связанным `URL` может
потребоваться схожая логика, например несколько изображений могут иметь `URL` вида `/images/*.png`. По этим причинам
`net/http` предоставляет `ServeMux`, `мультиплексор запросов`, упрощающий связь между `URL` и обработчиками. `ServeMux`
собирает целое множество обработчиков `http.Handler` в единый `http.Handler`. И вновь мы видим, что различные типы,
соответствующие одному и тоже не интерфейсу, являются `взаимозаменяемыми`: веб-сервер может диспетчеризовать запросы
любому `http.Handler`, независимо от того, какой конкретный тип скрывается за ним.

В более сложных приложениях для обработки более сложных требований к диспетчеризации несколько `ServeMux` могут
объединяться. Go не имеет канонического веб-каркаса, аналогичного `Ruby on Rails` или `Django в Python`. Это не значит,
что такого каркаса не может существовать, но строительные блоки в стандартной библиотеке `Go` являются столь гибкими,
что конкретный каркас просто не нужен. Кроме того, хотя наличие каркаса на ранних этапах проекта удобно, связанные с ним
дополнительные сложности могут усложнить долгосрочную поддержку проекта.

В приведенной ниже программе мы создаем `ServeMux` и используем его для сопоставления URL с соответствующими
обработчиками для операций `/list`, `/price`, которые были разделены на отдельные методы. Затем мы используем `ServeMux`
как основной обработчик в вызове `ListenAndServe` (см. http3.go):

``` go
type dollars float32
type database map[string]dollars

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)

}

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe(":8000", mux))
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "нет товара: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
```

Давайте сосредоточимся на двух вызовах `mux.Handle`, которые регистрируют обработчики. В первом `db.list` представляет
собой значение-метод (раздел 6.4), т.е. значение типа `func(w http.ResponseWriter, req *http.Request)`, которе при
вызове вызывает метод `database.list` со значением получателя `db`. Так что `db.list` является функцией, которая
реализует поведение обработчика, но так как у этой функции нет методов, она не может соответствовать
интерфейсу `http.Handler` и не может быть передана непосредственно `mux.Handle`.

Выражение `http.HanlerFunc(db.list)` представляет собой преобразование типа, а не вызов функции,
поскольку `http.HanlerFunc` является типом. Он имеет следующее определение:

``` go
package http

type HandlerFunc func(w ResponseWriter, r *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

`HandlerFunc` демонстрирует некоторые необычные возможности механизма интерфейсов Go. Это тип функции, имеющей методы и
соответствующей интерфейсу `http.Handler`. Поведением его метода `ServeHTTP` является вызов базовой функции. Таким
образом, `HandleFunc` является адаптером, который позволяет значению-функции соответствовать интерфейсу, когда функция и
единственный метод интерфейса имеют одинаковую сигнатуру. По сути, этот трюк позволяет одному типу, такому
как `database`, соответствовать интерфейсу `http.Handler` несколькими различными способами: посредством его
метода `list`, метода `price` и т.д.

Поскольку регистрация обработчика таким образом является весьма распространенной, `ServeMux` имеет удобный
метод `HandleFunc`, который делает это для нас, так что мы можем упростить код регистрации обработчика следующим
образом (см. http3a.go):

``` go
mux.HandleFunc("/list", db.list)
mux.HandleFunc("/price", db.price)
```

Из приведенного выше кода легко видеть, как можно создать программу, в которой есть два разных веб-сервера,
прослушивающих различные порты, определяющие различные URL и выполняющие диспетчеризацию различными обработчиками.
Необходимо просто создать еще один `ServeMux` и выполнить еще один вызов `ListenAndServe`, возможно, параллельно. Но для
большинства программ достаточно и одного веб-сервера. Кроме того, обычно обработчики HTTP определяются во многих файлах
приложения, и было бы неудобно, если бы все они должны были быть явно зарегистрированы экземпляром `ServeMux` уровня
приложения.

Поэтому для удобства пакет `net/http` предоставляет глобальный экземпляр `ServeMux` с именем `DefaultServeMux` и
функциям уровня пакета `http.Handle` и `http.HanldeFunc`. Для применения `DefaultServeMux` в качестве основного
обработчика сервера не нужно передавать его `ListenAndServe`; нужно передать `nil`.

Основная функция сервера при этом упрощается до (см. http4.go):

``` go
db := database{"socks": 5, "shoes": 50}
http.HandleFunc("/list", db.list)
http.HandleFunc("/price", db.price)
log.Fatal(http.ListenAndServe(":8000", nil))
```

Наконец еще одно важное напоминание: как мы упоминали в разделе 1.7, веб-сервер вызывает каждый обработчик в
новой `горутине`, так что обработчики должны принимать меры предосторожности, такие как `блокировки` при доступе к
переменным, к которым могут обращаться другие горутины (включая другие запросы того же обработчика). Мы будем говорить о
параллелизме в следующих главах.

## Выводы:

* Функция `ListenAndServe` требует адрес сервера и экземпляр интерфейса `Handler`, которому направляются все запросы. Он
  работает бесконечно, если только не происходит ошибка (или при запуске сервера происходит сбой), в таком случае он
  возвращает ненулевую ошибку;
* Интерфейс `http.Handler` имеет единственный метод `ServeHTTP(w ResponseWriter, *Requests`, который позволяет отвечать
  на `HTTP-запросы`;
* `ResponseWriter` - еще один интерфейс, он дополняет интерфейс `io.Writer` методами для отправки HTTP-заголовков
  ответа, а `http.Requests` - структура, которая содержит данные, соответствующие HTTP-запросу, такие, как URL,
  заголовки, тело ответа и т.д.;
* Если мы создадим тип и реализуем для него метод `ServeHTTP`, он будет соответствовать интерфейсу `http.Handler`;
* Чтобы сообщить клиенту об ошибке HTTP, если она произошла, нужно вызвать `w.WriteHeader(http.Status*...)`. Это должно
  быть сделано до записи любого текста в `w`. Так же можно использовать вспомогательную
  функцию `http.Error(w, msg, status)`;
* `r.URL.Query()` - выполняет запрос преобразования параметров HTTP-запроса в мультикарту типа `url.Values` Например:
  ``` go
    vals := r.URL.Query()
    val1 := vals.Get("key1")
    val2 := vals.Get("key2");
    ```
* Если мы реализуем для нашего типа метод `ServeHTTP`, нам нужно использовать `r.URL.Path` и инструкцию `switch` для
  определения адреса из запроса. Это несколько неудобно. Удобнее будет определить логику для каждого случая в виде
  отдельной функции или метода. Плюс к этому, связанным URL может потребоваться схожая логика, например, несколько
  изображений могут иметь `URL` вида `images/*.png`;
* Чтобы удобно добавлять разные варианты действий, и избежать вышеописанной ситуации, пакет `net/http`
  предоставляет `ServeMux` - `мультиплексор запросов`, упрощающий связь между URL и обработчиками (`Handler`);
* `ServeMux` собирает целое множество обработчиков `http.Handler` в единый `http.Handler`. Различные типы,
  соответствующие одному и тому же интерфейсу, являются `взаимозаменяемыми`. Веб-сервер может диспетчеризовать запросы к
  любому `http.Handler`, независимо от того, какой конкретный тип скрывается за ним;
* В более сложных приложениях могут использоваться несколько `ServeMux` и объединятся;
* В Go нет канонического веб-фреймворка, аналогичного Ruby on Rails или Django. Но это не значит, что такого фреймворка
  не может быть. Просто стандартная библиотека Go является настолько гибкой, что конкретный фреймворк просто не нужен.
  Тем более, что наличие фреймворка удобно на ранних этапах проекта, но связанные с ним дополнительные сложности могут
  усложнить дальнейшую поддержку проекта;
* `mux := http.NewServeMux(); mux.Handle("/list", http.HandlerFunc(db.list))` - создаем новый `ServeMux` и используем
  его для сопоставления URL с соответствующим обработчиком. После этого используем `ServeMux` как основной обработчик в
  вызове `log.Fatal(http.ListenAndServe(":8000", mux))`;
* `db.list` в `mux.Hanlde` представляет собой значение-метод, т.е. значение
  типа `func(w http.ResponseWriter, r *http.Requests)`, которое при вызове вызывает метод `database.list`, со значением
  получателя `db`. Проще говоря, `db.list` является функцией, которая реализует поведение обработчика, но, так как у
  этой функции нет методов, она не может соответствовать интерфейсу `http.Hanlder` и не может быть передана
  непосредственно `mux.Handle`;
* `http.HandlerFunc(db.list)` - это преобразование типа, а не вызов функции, поскольку `http.HandlerFunc` является
  типом;
* `HandlerFunc` демонстрирует некоторые необычные возможности механизма интерфейсов Go. Это тип функции, который имеет
  методы и соответствует интерфейсу `http.Handler`. Поведением его метода `ServeHTTP` является вызов базовой функции. Таким
  образом, `HandlerFunc` является адаптером, который позволяет значению-функции соответствовать интерфейсу, когда
  функция и единственный метод интерфейса имеют одинаковую сигнатуру;
* Этот трюк, позволяет типу `type database map[string]dollar` соответствовать интерфейсу `http.Handler`
  различными способами - его методами, которые имеют ту же сигнатуру как и интерфейс;
* `ServeMux` имеет удобный метод `HandleFunc`, который приводит тип `database` к соответствию интерфейсу `http.Handler`.
  Поэтому можно упростить код регистрации обработчика до `mux.HandleFunc("/list", db.list)`;
* Чтобы создать два разных веб сервера, которые будут прослушивать разные порты, определять разные URL и выполнять
  диспетчеризацию разными обработчиками - нужно создать еще один `ServeMux` и выполнить еще один вызов `ListenAndServe`;
* Пакет `net/http` предоставляет глобальный экземпляр `ServeMux` с именем `DefaultServeMux` и функциями уровня
  пакета `http.Hanlde` и `http.HanldeFunc`;
* Для использования `DefaultServeMux` в качестве основного обработчика сервера в `ListenAndServe` нужно передать `nil`;
* Веб-сервер вызывает каждый обработчик в новой горутине, так что обработчики должны принимать меры предосторожности,
  такие как блокировки при доступе к переменным, к которым могут обращаться другие горутины (включая другие запросы того
  же обработчика);