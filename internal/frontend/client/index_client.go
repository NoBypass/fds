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
	input.wrapper = document.GetElementByID("input-wrapper")
	input.colors = colorPrimary
	input.wrapper.AddClass(input.colors...)

	mi.On("input", input.mainInputInput)
	mi.On("focus", input.mainInputFocus)
	mi.On("blur", input.mainInputBlur)
}

var (
	colorPrimary = []string{"border-neutral-600", "has-[:focus]:border-indigo-500", "outline-indigo-900"}
	colorSuccess = []string{"border-emerald-600", "has-[:focus]:border-emerald-500", "outline-emerald-900"}
	colorError   = []string{"border-rose-600", "has-[:focus]:border-rose-500", "outline-rose-900"}
)

type mainInput struct {
	value                 string
	link, cursor, wrapper *document.Element
	lastInput             time.Time
	colors                []string
}

func (i *mainInput) mainInputInput(this js.Value, _ []js.Value) any {
	i.value = this.Get("value").String()
	i.link.Set("href", "/player/"+i.value)
	i.cursor.Set("style", fmt.Sprintf("transform: translateY(8px) translateX(%fpx)", measureValWith(i.value)-3))
	i.cursor.RemoveClass("animate-blink")
	i.lastInput = time.Now()
	go func() {
		time.Sleep(500 * time.Millisecond)
		if time.Since(i.lastInput) < 500*time.Millisecond {
			return
		}
		i.cursor.AddClass("animate-blink")
	}()
	return nil
}

func (i *mainInput) mainInputFocus(this js.Value, _ []js.Value) any {
	i.cursor.RemoveClass("hidden")
	this.Set("placeholder", "")
	return nil
}

func (i *mainInput) mainInputBlur(this js.Value, _ []js.Value) any {
	i.cursor.AddClass("hidden")
	this.Set("placeholder", "Name")
	return nil
}

func measureValWith(text string) float64 {
	return js.Global().Get("measureValWith").Invoke(text).Float()
}
