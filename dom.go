package dom

import (
	"fmt"
	"syscall/js"
)

var consoleHandle = js.Global().Get("console")

// Document caches the global document DOM element
var Document = js.Global().Get("document")

// Body caches the global <body> element
var Body = Document.Get("body")

// Log prints the provided message string to the browser console.
func Log(msg string) {
	consoleHandle.Call("log", msg)
}

// Logf prints the formated message with parameters to the browser console.
// It uses fmt.Sprintf(msg, args...) internally.
func Logf(msg string, args ...interface{}) {
	consoleHandle.Call("log", fmt.Sprintf(msg, args...))
}

// Error prints a string to the error channel in the browser console.
func Error(msg string) {
	consoleHandle.Call("error", msg)
}

// Warning prints the given string to the warning channel in the browser console.
func Warning(msg string) {
	consoleHandle.Call("warning", msg)
}

// CreateElement creates HTMLElements on the DOM and adds it to the parent
func CreateElement(elType string, classes string, parent *js.Value) HTMLElement {
	element := Document.Call("createElement", elType)
	if len(classes) > 0 {
		element.Call("setAttribute", "class", classes)
	}
	if parent != nil {
		parent.Call("appendChild", element)
	}
	e := HTMLElement{jsValue: element, name: elType}
	return e
}

// Div creates a a <div> element with the provided classes and adds it to the parent.
func Div(classes string, parent *js.Value) HTMLElement {
	return CreateElement("div", classes, parent)
}

// RemoveAllChildren removes all child nodes from given node.
func RemoveAllChildren(node *js.Value) {
	for !node.Get("lastChild").IsNull() {
		node.Call("removeChild", node.Get("lastChild"))
	}
}

// FullPageCanvas creates a canvas element that spans over the entire body.
func FullPageCanvas() HTMLElement {
	Body.Set("style", "margin: 0;padding: 0;height:100%; overflow: hidden;")

	RemoveAllChildren(&Body)

	canvas := CreateElement("canvas", "", &Body)
	canvas.SetStyle("width:100%;height:100%;margin 0;padding 0;position:absolute")

	return canvas
}
