package document

import "syscall/js"

type Document struct {
	js.Value
}

var document Document

func init() {
	document = Document{
		js.Global().Get("document"),
	}
}
