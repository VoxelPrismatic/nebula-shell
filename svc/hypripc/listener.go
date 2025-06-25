package hypripc

import "sync"

type Listenable[R any] interface {
	Update(event, value string) bool
	Target() (*R, error)
}

type EventListener[T Listenable[R], R any] struct {
	listenLock sync.Mutex
	store      *T
	listeners  []func(T)
}

func (lis *EventListener[T, R]) Add(cb func(T)) {
	lis.listenLock.Lock()
	lis.listeners = append(lis.listeners, cb)
	lis.listenLock.Unlock()
}

func (lis *EventListener[T, R]) update(event, value string) {
	if lis.store == nil {
		lis.store = new(T)
	}
	if (*lis.store).Update(event, value) {
		fire(lis.store, lis.listeners)
	}
}
