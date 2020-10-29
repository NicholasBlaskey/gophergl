package main

import (
	"runtime"

	"github.com/nicholasblaskey/gophergl/Web/gl"
	//"github.com/nicholasblaskey/gophergl/Open/gl"
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
	out vec4 FragColor;
	in vec2 uv;
	
	uniform sampler2D texture1;

	void main()
	{	
		vec2 tileAmount = floor(uv);
		float isEven = mod(tileAmount.x + tileAmount.y, 2.0);
		vec3 tOut = vec3(texture(texture1, uv));
		
		if (isEven == 0.0) {
			FragColor = vec4(1.0 - tOut, 1.0);
		} else {
			FragColor = vec4(tOut, 1.0);
		}

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
		+0.9, +0.9, 4.0, 0.0,
		+0.9, -0.9, 4.0, 4.0,
		-0.9, -0.9, 0.0, 4.0,
		-0.9, +0.9, 0.0, 0.0,
	})

	t1, err := gl.TextureFromFile("./images/gopher.jpg")
	if err != nil {
		panic(err)
	}
	t2, err := gl.TextureFromFile("./images/turtle.jpg")
	if err != nil {
		panic(err)
	}

	t1.Bind(gl.TEXTURE0)
	t2.Bind(gl.TEXTURE1)

	i := 0
	window.Run(func() {
		gl.ClearColor(0.1, 0.1, 0.1, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		if i%200 == 0 {
			shader.SetInt("texture1", 1)
		} else if i%100 == 0 {
			shader.SetInt("texture1", 0)
		}
		vao.Draw()

		i += 1
		window.PollEvents()
		window.SwapBuffers()
	})
}
