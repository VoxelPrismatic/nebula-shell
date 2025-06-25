package hypripc

import (
	"reflect"
	"sync"
)

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
	if lis.store == nil || reflect.ValueOf(*lis.store).IsNil() {
		underlying := reflect.TypeOf((*T)(nil)).Elem()
		for underlying.Kind() == reflect.Ptr {
			underlying = underlying.Elem()
		}
		if underlying.Kind() != reflect.Struct {
			panic("*listenable must be a struct, got " + underlying.Kind().String())
		}
		inst, ok := reflect.New(underlying).Interface().(T)
		if !ok {
			panic("failed to make event store")
		}
		lis.store = &inst
	}
	if (*lis.store).Update(event, value) {
		fire(lis.store, lis.listeners)
	}
}
