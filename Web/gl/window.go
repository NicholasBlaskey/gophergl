package gl

import (
	"errors"

	"github.com/gopherjs/gopherjs/js"
)

type Window struct {
	window *js.Object
}

var webgl *js.Object = nil

func (w *Window) Run(frameFunc func()) {
	var render func()
	render = func() {
		frameFunc()
		js.Global.Call("requestAnimationFrame", render)
	}
	js.Global.Call("requestAnimationFrame", render)
}

func NewWindow(width, height int32, title string) (*Window, error) {
	document := js.Global.Get("document")
	canvas := document.Call("createElement", "canvas")
	document.Get("body").Call("appendChild", canvas)
	canvas.Set("width", width)
	canvas.Set("height", height)

	webgl = canvas.Call("getContext", "webgl")
	if webgl == nil {
		webgl = canvas.Call("getContext", "experimental-webgl")
		if webgl == nil {
			return nil, errors.New("Browser does not support webgl context")
		}
	}
	return &Window{canvas}, nil
}

func (w *Window) Terminate() {}

func (w *Window) GetTime() float32 { return 0.0 }

func (w *Window) PollEvents() {}

func (w *Window) SwapBuffers() {}

func ClearColor(r, g, b, a float32) {
	webgl.Call("clearColor", r, g, b, a)
}

const (
	DEPTH_BUFFER_BIT   = 0x00000100
	STENCIL_BUFFER_BIT = 0x00000400
	COLOR_BUFFER_BIT   = 0x00004000
)

func Clear(mask uint32) {
	webgl.Call("clear", mask)
}
