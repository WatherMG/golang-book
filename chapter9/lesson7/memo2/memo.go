/*
Пакет memo обеспечивает безопасную для параллельных запросов мемоизацию функции
типа Func. Параллельные запросы сериализуются мьютексом.
*/

package memo2

import "sync"

type Memo struct {
	f     Func
	mu    sync.Mutex // Защита cache
	cache map[string]result
}

// Func является типом функции с запоминанием.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// Метод Get безопасен с точки зрения параллельности.
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()
	return res.value, res.err
}
