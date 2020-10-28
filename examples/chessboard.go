package main

import (
	"runtime"

	"github.com/nicholasblaskey/gophergl/Web/gl"
	//"github.com/nicholasblaskey/gophergl/Open/gl"

	mgl "github.com/go-gl/mathgl/mgl32"
)

func init() {
	runtime.LockOSThread()
}

const (
	vertexShader = `#version 410 core
	layout (location = 0) in vec2 position;

	void main()
	{
		gl_Position = vec4(position, 0.0, 1.0);
	}`

	fragShader = `#version 410 core
	out vec4 FragColor;

	uniform vec2 dims;
	uniform float squareAmount;
	uniform vec3 col1;
	uniform vec3 col2;

	void main()
	{	
		vec2 uv = floor(vec2(gl_FragCoord) / dims * squareAmount);
		float isEven = mod(uv.x + uv.y, 2.0);	
		FragColor = vec4((1.0 - isEven) * col1 + isEven * col2, 1.0);
	}`
)

func main() {
	width, height := int32(800), int32(600)
	window, err := gl.NewWindow(width, height, "Checkerboard")
	if err != nil {
		panic(err)
	}
	defer window.Terminate()

	shader, err := gl.CompileShader(vertexShader, fragShader)
	if err != nil {
		panic(err)
	}
	shader.Use()

	vao := gl.NewVAO(gl.TRIANGLE_FAN, []int32{2}, []float32{
		+1.0, +1.0,
		+1.0, -1.0,
		-1.0, -1.0,
		-1.0, +1.0,
	})

	window.Run(func() {
		gl.ClearColor(0.1, 0.1, 0.1, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		shader.SetFloat("squareAmount", 10.0)
		shader.SetVec2("dims", mgl.Vec2{float32(width), float32(height)})
		shader.SetVec3("col1", mgl.Vec3{0.3, 0.5, 0.3})
		shader.SetVec3("col2", mgl.Vec3{0.3, 0.3, 0.5})

		vao.Draw()

		window.PollEvents()
		window.SwapBuffers()
	})
}
