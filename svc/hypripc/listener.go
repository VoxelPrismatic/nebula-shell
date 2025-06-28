package hypripc

import (
	"reflect"
	"slices"
	"sync"
)

type Listenable[R any] interface {
	Update(event, value string) bool
	Target() (*R, error)
}

// The callback should return True if it should be dropped.
type EventCallback[T any] func(T) bool

type EventListener[T Listenable[R], R any] struct {
	listenLock sync.Mutex
	store      *T
	listeners  *[]*EventCallback[T]
}

func (lis *EventListener[T, R]) Add(cb EventCallback[T]) *EventCallback[T] {
	lis.listenLock.Lock()
	if lis.listeners == nil {
		lis.listeners = &[]*EventCallback[T]{&cb}
	} else {
		*lis.listeners = append(*lis.listeners, &cb)
	}
	lis.listenLock.Unlock()
	return &cb
}

func (lis *EventListener[T, R]) Drop(cb *EventCallback[T]) {
	lis.listenLock.Lock()
	for i, v := range *lis.listeners {
		if v == cb {
			*lis.listeners = slices.Delete(*lis.listeners, i, i+1)
			return
		}
	}
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
