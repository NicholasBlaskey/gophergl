package gl

import (
	"errors"

	"github.com/gopherjs/gopherjs/js"
	"time"
	//	"fmt"
)

type Window struct {
	window    *js.Object
	startTime time.Time
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
	return &Window{canvas, time.Now()}, nil
}

func (w *Window) Terminate() {}

func (w *Window) GetTime() float32 {
	return float32(time.Now().Sub(w.startTime).Seconds())
}

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

const (
	DEPTH_TEST = 0x0B71
	LEQUAL     = 0x0203
	LESS       = 0x0201

	BLEND                    = 0x0BE2
	ONE_MINUS_SRC_ALPHA      = 0x0303
	ONE_MINUS_CONSTANT_COLOR = 0x8002
	ONE_MINUS_DST_ALPHA      = 0x0305
	ONE_MINUS_DST_COLOR      = 0x0307
	ONE_MINUS_SRC_COLOR      = 0x0301

	SRC_ALPHA = 0x0302
)

func Enable(v uint32) {
	webgl.Call("enable", v)
}

func DepthFunc(v uint32) {
	webgl.Call("depthFunc", v)
}

func BlendFunc(sFact, dFact uint32) {
	webgl.Call("blendFunc", sFact, dFact)
}
