package main

import (
	"runtime"

	mgl "github.com/go-gl/mathgl/mgl32"

	"github.com/nicholasblaskey/gophergl/Open/gl"
	//"github.com/nicholasblaskey/gophergl/Web/gl"
)

func init() {
	runtime.LockOSThread()
}

const (
	vertexShader = `#version 410 core
	layout (location = 0) in vec3 position;
	layout (location = 1) in vec2 auv;
	
	out vec2 uv;

	uniform mat4 projection;
	uniform mat4 view;
	uniform mat4 model;


	void main()
	{
		uv = auv;
		gl_Position = projection * view * model * vec4(position, 1.0);
	}`

	fragShader = `#version 410 core
	out vec4 FragColor;
	in vec2 uv;
	
	uniform sampler2D texture1;
	uniform vec3 color;

	void main()
	{	
		vec4 sampleCol = texture(texture1, uv);
		if (sampleCol.a == 1.0) {
			FragColor = sampleCol;
		} else {
			FragColor = vec4(color, 0.2);	
		}
	}`
)

func main() {
	width, height := int32(800), int32(600)
	window, err := gl.NewWindow(width, height, "blending")
	if err != nil {
		panic(err)
	}
	defer window.Terminate()
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	camera := gl.NewOrbitalCamera(window, 5.0, mgl.Vec3{0.0, 0.0, 0.0})

	shader, err := gl.CompileShader(vertexShader, fragShader)
	if err != nil {
		panic(err)
	}
	shader.Use()

	vao := gl.NewVAO(gl.NewCube(gl.VertParams{Position: true, TexCoords: true}))
	t1, err := gl.TextureFromFile("./images/window.png")
	if err != nil {
		panic(err)
	}
	t1.Bind(gl.TEXTURE0)
	shader.SetInt("texture1", 0)

	projection := mgl.Perspective(mgl.DegToRad(45.0),
		float32(width)/float32(height), 0.1, 100.0)
	shader.SetMat4("projection", projection)

	window.Run(func() {
		gl.ClearColor(0.1, 0.1, 0.1, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		shader.SetMat4("view", camera.LookAt())

		distApart := float32(0.75)

		shader.SetVec3("color", mgl.Vec3{0.9, 0.3, 0.3})
		drawCube(shader, vao, 0.0, 0.0, 0.0)

		shader.SetVec3("color", mgl.Vec3{0.3, 0.9, 0.3})
		drawCube(shader, vao, 0.0, +distApart, 0.0)

		shader.SetVec3("color", mgl.Vec3{0.3, 0.3, 0.9})
		drawCube(shader, vao, 0.0, -distApart, 0.0)

		shader.SetVec3("color", mgl.Vec3{0.9, 0.9, 0.3})
		drawCube(shader, vao, +distApart, 0.0, 0.0)

		shader.SetVec3("color", mgl.Vec3{0.3, 0.9, 0.9})
		drawCube(shader, vao, -distApart, 0.0, 0.0)

		shader.SetVec3("color", mgl.Vec3{0.9, 0.3, 0.9})
		drawCube(shader, vao, 0.0, 0.0, +distApart)

		shader.SetVec3("color", mgl.Vec3{0.9, 0.9, 0.9})
		drawCube(shader, vao, 0.0, 0.0, -distApart)

		//drawCube(shader, vao, 0.0, 0.0, 0.0)
		//xdrawCube(shader, vao, 0.0, 0.0, 0.0)

		window.PollEvents()
		window.SwapBuffers()
	})
}

func drawCube(shader *gl.Shader, vao *gl.VAO, x, y, z float32) {
	shader.SetMat4("model", mgl.Translate3D(x, y, z).Mul4(
		mgl.Scale3D(0.25, 0.25, 0.25)))
	vao.Draw()
}
