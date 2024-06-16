package document

import (
	"syscall/js"
)

type Element struct {
	js.Value
}

func GetElementByID(id string) *Element {
	return &Element{
		document.Call("getElementById", id),
	}
}

func (e *Element) On(ev string, fn func(this js.Value, args []js.Value) any) {
	e.Call("addEventListener", ev, js.FuncOf(fn))
}

func (e *Element) ToggleClass(class ...string) {
	for _, c := range class {
		e.Get("classList").Call("toggle", c)
	}
}

func (e *Element) RemoveClass(class ...string) {
	for _, c := range class {
		e.Get("classList").Call("remove", c)
	}
}

func (e *Element) AddClass(class ...string) {
	for _, c := range class {
		e.Get("classList").Call("add", c)
	}
}

func (e *Element) ReplaceInner(inner string) {
	e.Set("innerHTML", inner)
}
