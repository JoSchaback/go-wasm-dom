package dom

import "syscall/js"

// HTMLElement wraps js.Value and adds helpful methods.
type HTMLElement struct {
	jsValue js.Value
	name    string
}

// SetStyle sets the "style" attribute on the element.
// It is a shorthand for jsValue.Set("style", style)
func (c *HTMLElement) SetStyle(style string) {
	c.jsValue.Set("style", style)
}

// SetWidth sets the "width" attribute of the element.
// It is a shorthand for jsValue.Set("width", width)
func (c *HTMLElement) SetWidth(width string) {
	c.jsValue.Set("width", width)
}

// SetWidth sets the "width" attribute of the element as an integer.
func (c *HTMLElement) SetWidthI(width int) {
	c.jsValue.Set("width", width)
}

// SetHeight sets the "height" attribute of the element.
// It is a shorthand for jsValue.Set("width", width)
func (c *HTMLElement) SetHeight(height string) {
	c.jsValue.Set("height", height)
}

// SetHeight sets the "height" attribute of the element as an integer.
func (c *HTMLElement) SetHeightI(height int) {
	c.jsValue.Set("height", height)
}

// JsValue exposes the privat field jsValue that holds the actual HTML element.
func (c *HTMLElement) JsValue() js.Value {
	return c.jsValue
}
