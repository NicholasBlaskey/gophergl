package main

import (
	"runtime"

	//"github.com/nicholasblaskey/gophergl/Web/gl"
	"github.com/nicholasblaskey/gophergl/Open/gl"
)

func init() {
	runtime.LockOSThread()
}

const (
	vertexShader = `#version 410 core
	layout (location = 0) in vec3 position;
	void main()
	{
		gl_Position = vec4(position.x, position.y, position.z, 1.0);
	}`

	fragShader = `#version 410 core
	out vec4 color;
	void main()
	{
		color = vec4(1.0f, 0.5f, 0.2f, 1.0f);
	}`
)

func main() {
	width, height := int32(800), int32(600)
	window, err := gl.NewWindow(width, height, "hello triangle!")
	if err != nil {
		panic(err)
	}
	defer window.Terminate()

	shader, err := gl.CompileShader(vertexShader, fragShader)
	if err != nil {
		panic(err)
	}

	shader.Use()
	vao := gl.NewVAO(gl.TRIANGLES, []int32{3}, []float32{
		-0.5, -0.5, 0.0,
		0.5, -0.5, 0.0,
		0.0, 0.5, 0.0,
	})

	i := 0
	window.Run(func() {
		gl.ClearColor(0.3, 0.5, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		vao.Draw()

		window.PollEvents()
		window.SwapBuffers()

		i += 1
	})
}
