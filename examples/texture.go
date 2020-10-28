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
	layout (location = 0) in vec2 position;	
	layout (location = 1) in vec2 auv;
	
	out vec2 uv;

	void main()
	{
		uv = auv;
		gl_Position = vec4(position, 0.0, 1.0);
	}`

	fragShader = `#version 410 core
	out vec4 color;
	in vec2 uv;
	
	uniform sampler2D texture1;

	void main()
	{	
		color = vec4(vec3(1.0 - texture(texture1, uv)), 1.0);
	}`
)

func main() {
	width, height := int32(800), int32(600)
	window, err := gl.NewWindow(width, height, "texture")
	if err != nil {
		panic(err)
	}
	defer window.Terminate()

	shader, err := gl.CompileShader(vertexShader, fragShader)
	if err != nil {
		panic(err)
	}
	shader.Use()

	vao := gl.NewVAO(gl.TRIANGLE_FAN, []int32{2, 2}, []float32{
		// Pos, then tex coords
		+0.9, +0.9, 1.0, 0.0,
		+0.9, -0.9, 1.0, 1.0,
		-0.9, -0.9, 0.0, 1.0,
		-0.9, +0.9, 0.0, 0.0,
	})

	texture := gl.TextureFromFile("./images/gopher.jpg")
	texture.Bind(gl.TEXTURE0)

	window.Run(func() {
		gl.ClearColor(0.1, 0.1, 0.1, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		shader.SetInt("texture1", 0)

		vao.Draw()

		window.PollEvents()
		window.SwapBuffers()
	})
}
