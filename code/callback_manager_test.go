package code

/*
	这个 demo 是一个回调函数的管理器
*/

import (
	"fmt"
	"sync"
	"testing"
)

var (
	callback map[string]func()
	mu       sync.Mutex
)

func TestCallbackManager(t *testing.T) {

	Register(foo1, "foo1")
	Register(foo2, "foo2")

	if err := Deregister("foo1"); err != nil {
		panic(err)
	}

	for _, f := range callback {
		f()
	}
}

func Register(fun func(), key string) {
	mu.Lock()
	defer mu.Unlock()
	if callback == nil {
		callback = make(map[string]func())
	}
	callback[key] = fun
}

func Deregister(key string) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := callback[key]; !exists {
		return fmt.Errorf("callback %s not registered", key)
	}
	delete(callback, key)
	return nil
}

func foo1() { fmt.Println("foo1") }

func foo2() { fmt.Println("foo2") }
