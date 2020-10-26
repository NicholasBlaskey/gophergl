package main

import (
	"runtime"

	"github.com/nicholasblaskey/gophergl/Open/gl"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	width, height := int32(800), int32(600)
	window, err := gl.NewWindow(width, height, "hello window!")
	if err != nil {
		panic(err)
	}
	defer window.Terminate()

	window.Run(func() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.ClearColor(0.3, 0.5, 0.3, 1.0)

		window.PollEvents()
		window.SwapBuffers()
	})
}
