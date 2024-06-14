package document

import (
	"sync"
	"syscall/js"
)

type Element struct {
	js.Value
	mu sync.Mutex
}

func GetElementByID(id string) *Element {
	return &Element{
		document.Call("getElementById", id),
		sync.Mutex{},
	}
}

func (e *Element) On(ev string, fn func(this js.Value, args []js.Value) any) {
	e.Call("addEventListener", ev, js.FuncOf(fn))
}

func (e *Element) ToggleClass(class string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.Get("classList").Call("toggle", class)
}
