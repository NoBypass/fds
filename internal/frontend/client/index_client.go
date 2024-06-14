package client

import (
	"fmt"
	"github.com/NoBypass/fds/internal/frontend/wasm_utils/document"
	"syscall/js"
	"time"
)

func HandleIndex() {
	var input mainInput
	input.cursor = document.GetElementByID("cursor")
	input.link = document.GetElementByID("player-stats-link")

	mi := document.GetElementByID("main-input")
	mi.On("input", input.mainInputInput)
	mi.On("focus", input.mainInputFocus)
	mi.On("blur", input.mainInputBlur)
}

type mainInput struct {
	value        string
	link, cursor *document.Element
	lastInput    time.Time
}

func (i *mainInput) mainInputInput(this js.Value, _ []js.Value) any {
	i.value = this.Get("value").String()
	i.link.Set("href", "/player/"+i.value)
	i.cursor.Set("style", fmt.Sprintf("transform: translateY(8px) translateX(%fpx)", measureValWith(i.value)-3))
	i.cursor.ToggleClass("animate-blink")
	i.lastInput = time.Now()
	go func() {
		time.Sleep(time.Second)
		if time.Since(i.lastInput) < time.Second {
			return
		}
		i.cursor.ToggleClass("animate-blink")
	}()
	return nil
}

func (i *mainInput) mainInputFocus(this js.Value, _ []js.Value) any {
	i.cursor.Get("classList").Call("remove", "hidden")
	this.Set("placeholder", "")
	return nil
}

func (i *mainInput) mainInputBlur(this js.Value, _ []js.Value) any {
	i.cursor.Get("classList").Call("add", "hidden")
	this.Set("placeholder", "Name")
	return nil
}

func measureValWith(text string) float64 {
	return js.Global().Get("measureValWith").Invoke(text).Float()
}
