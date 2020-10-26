package main

import (
	"runtime"

	"github.com/nicholasblaskey/gophergl/Web/gl"
	//"github.com/nicholasblaskey/gophergl/Open/gl"
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

	i := 0
	window.Run(func() {
		if i%200 == 0 {
			gl.ClearColor(0.3, 0.5, 0.3, 1.0)
		} else if i%100 == 0 {
			gl.ClearColor(0.5, 0.3, 0.3, 1.0)
		}
		gl.Clear(gl.COLOR_BUFFER_BIT)

		window.PollEvents()
		window.SwapBuffers()

		i += 1
	})
}
