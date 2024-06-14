package wasm_utils

import (
	"log"
	"syscall/js"
)

func init() {
	log.SetOutput(&WASMWriter{})
}

type WASMWriter struct{}

func (ww *WASMWriter) Write(p []byte) (n int, err error) {
	js.Global().Get("console").Call("log", string(p))
	return len(p), nil
}
