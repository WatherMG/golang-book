/*
Пакет memo обеспечивает параллельно безопасную неблокирующую мемоизацию
функции. Запросы к разным ключам выполняются параллельно.
Последовательные запросы к одному и тому же ключу блокируются до завершения первого.
В данной реализации используется горутина монитора.
*/

package memo

// Func является типом функции с запоминанием.
type Func func(key string) (interface{}, error)

// result - это результат вызова функции Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // Закрывается, когда res готов
}

// requests представляет собой сообщение,
// требующее применения Func к key.
type request struct {
	key      string
	response chan<- result // Клиенту нужен только result
}

type Memo struct {
	requests chan request
}

// New возвращает f с запоминанием.
// Впоследствии клиенты должны вызывать Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

// Метод Get безопасен с точки зрения параллельности.
func (memo *Memo) Get(key string) (value interface{}, err error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

// !+monitor

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// Это первый запрос данного ключа key
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // Вызов f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// Вычисление функции.
	e.res.value, e.res.err = f(key)
	// Оповещение о готовности
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Ожидание готовности
	<-e.ready
	// Отправка результата клиенту
	response <- e.res
}

// !-monitor
