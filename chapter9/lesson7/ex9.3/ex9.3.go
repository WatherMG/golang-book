/*
Exercise 9.3
Расширьте тип Func и метод (*Memo).Get так, чтобы вызывающая функция могла
предоставить необязательный канал done, с помощью которого можно было бы
отменить операцию (раздел 8.9). Результаты отмененного вызова Func кешироваться
не должны.
*/

package memo

import "fmt"

// Func является типом функции с запоминанием.
type Func func(key string, done <-chan struct{}) (interface{}, error)

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
	done     <-chan struct{}
}

type Memo struct {
	requests, cancels chan request
}

// New возвращает f с запоминанием.
// Впоследствии клиенты должны вызывать Close.
func New(f Func) *Memo {
	memo := &Memo{make(chan request), make(chan request)}
	go memo.server(f)
	return memo
}

// Метод Get безопасен с точки зрения параллельности.
func (memo *Memo) Get(key string, done <-chan struct{}) (value interface{}, err error) {
	response := make(chan result)
	req := request{key, response, done}
	memo.requests <- req
	fmt.Println("get: waiting for response")
	res := <-response
	fmt.Println("get: checking if cancelled")
	select {
	case <-done:
		fmt.Println("get: queueing cancellation request")
		memo.cancels <- req
	default:
		// Not cancelled. Continue.
	}
	fmt.Println("get: return")
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

// !+monitor

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
Loop:
	for {
	Cancel:
		for {
			select {
			case req := <-memo.cancels:
				fmt.Println("server: deleting cancelled entry (early)")
				delete(cache, req.key)
			default:
				break Cancel
			}
		}
		// Ожидание запросов или отмен, и прерывание для обработки всех
		// отмен, если они есть.
		select {
		case req := <-memo.cancels:
			fmt.Println("server: deleting cancelled entry")
			delete(cache, req.key)
			continue Loop
		case req, ok := <-memo.requests:
			if !ok {
				return
			}
			fmt.Println("server: request")
			e := cache[req.key]
			if e == nil {
				// Это первый запрос данного ключа key
				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				go e.call(f, req.key, req.done) // Вызов f(key)
			}
			go e.deliver(req.response)
		}
	}
}

func (e *entry) call(f Func, key string, done <-chan struct{}) {
	// Вычисление функции.
	e.res.value, e.res.err = f(key, done)
	fmt.Println("call: returned from f")
	// Оповещение о готовности
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Ожидание готовности
	<-e.ready
	// Отправка результата клиенту
	fmt.Println("deliver: add in cache")
	response <- e.res
}

// !-monitor
