package gl

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

type Window struct {
	window *glfw.Window
}

func (w *Window) Run(frameFunc func()) {
	for !w.window.ShouldClose() {
		frameFunc()
	}
}

func NewWindow(width, height int32, title string) (*Window, error) {
	if err := glfw.Init(); err != nil {
		return nil, err
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(
		int(width), int(height), title, nil, nil)
	if err != nil {
		return nil, err
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		return nil, err
	}
	return &Window{window}, nil
}

func (w *Window) Terminate() {
	glfw.Terminate()
}

func (w *Window) GetTime() float32 {
	return float32(glfw.GetTime())
}

func (w *Window) PollEvents() {
	glfw.PollEvents()
}

func (w *Window) SwapBuffers() {
	w.window.SwapBuffers()
}

func ClearColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
}

const (
	DEPTH_BUFFER_BIT   = gl.DEPTH_BUFFER_BIT
	STENCIL_BUFFER_BIT = gl.STENCIL_BUFFER_BIT
	COLOR_BUFFER_BIT   = gl.COLOR_BUFFER_BIT
)

func Clear(mask uint32) {
	gl.Clear(mask)
}
