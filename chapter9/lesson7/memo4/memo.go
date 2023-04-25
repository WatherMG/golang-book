/*
Пакет memo предоставляет безопасную с точки зрения параллелизма мемоизацию функции f.
Запросы на разные ключи выполняются параллельно.
Параллельные запросы на один и тот же ключ блокируются до завершения первого.
В данной реализации используется мьютекс.
*/

package memo4

import "sync"

type entry struct {
	res   result
	ready chan struct{} // Закрывается, когда res готов
}

type Memo struct {
	f     Func
	mu    sync.Mutex // Защита cache
	cache map[string]*entry
}

// Func является типом функции с запоминанием.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

// Метод Get безопасен с точки зрения параллельности.
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// Это первый запрос данного ключа.
		// Эта горутина становиться ответственной за
		// вычисление значения и оповещение о готовности.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)
		close(e.ready) // Широковещательное оповещение о готовности
	} else {
		// Это повторный запрос данного ключа.
		memo.mu.Unlock()

		<-e.ready // Ожидание готовности
	}
	return e.res.value, e.res.err
}
