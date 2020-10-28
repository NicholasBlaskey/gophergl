package main

import (
	"runtime"

	"math"

	"github.com/nicholasblaskey/gophergl/Web/gl"
	//"github.com/nicholasblaskey/gophergl/Open/gl"
)

func init() {
	runtime.LockOSThread()
}

const (
	vertexShader = `#version 410 core
	layout (location = 0) in vec3 position;
	layout (location = 1) in vec2 color;

	out vec2 rgIn;
	void main()
	{
		rgIn = color;
		gl_Position = vec4(position, 1.0);
	}`

	fragShader = `#version 410 core
	out vec4 color;
	in vec2 rgIn;

	uniform float redAmount;

	void main()
	{
		color = vec4(rgIn, redAmount, 1.0);
	}`
)

func main() {
	width, height := int32(800), int32(600)
	window, err := gl.NewWindow(width, height, "square uniform")
	if err != nil {
		panic(err)
	}
	defer window.Terminate()

	shader, err := gl.CompileShader(vertexShader, fragShader)
	if err != nil {
		panic(err)
	}
	shader.Use()

	vao := gl.NewVAO(gl.TRIANGLE_FAN, []int32{3, 2}, []float32{
		// Pos, then color
		+0.5, +0.5, 0.0, 0.1, 0.5,
		+0.5, -0.5, 0.0, 0.1, 0.9,
		-0.5, -0.5, 0.0, 0.5, 0.1,
		-0.5, +0.5, 0.0, 0.9, 0.1,
	})

	window.Run(func() {
		gl.ClearColor(0.1, 0.1, 0.1, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		t := float32(math.Sin(float64(window.GetTime() * 5.0)))
		shader.SetFloat("redAmount", t+1.0)
		vao.Draw()

		window.PollEvents()
		window.SwapBuffers()
	})
}
