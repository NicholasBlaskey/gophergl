package main

import (
	"runtime"

	mgl "github.com/go-gl/mathgl/mgl32"

	//"github.com/nicholasblaskey/gophergl/Open/gl"
	"github.com/nicholasblaskey/gophergl/Web/gl"
)

func init() {
	runtime.LockOSThread()
}

const (
	vertexShader = `#version 410 core
	layout (location = 0) in vec3 position;

	uniform mat4 projection;
	uniform mat4 view;
	uniform mat4 model;

	void main()
	{
		gl_Position = projection * view * model * vec4(position, 1.0);
	}`

	fragShader = `#version 410 core
	out vec4 FragColor;
	in vec2 uv;
	
	uniform vec3 col;

	void main()
	{	
		FragColor = vec4(col, 1.0);
	}`
)

func main() {
	width, height := int32(800), int32(600)
	window, err := gl.NewWindow(width, height, "cube bouncing")
	if err != nil {
		panic(err)
	}
	defer window.Terminate()
	camera := gl.NewOrbitalCamera(window, 5.0, mgl.Vec3{0.0, 0.0, 0.0})

	gl.Enable(gl.DEPTH_TEST)

	shader, err := gl.CompileShader(vertexShader, fragShader)
	if err != nil {
		panic(err)
	}
	shader.Use()

	vao := gl.NewVAO(gl.NewCube(gl.VertParams{Position: true}))

	projection := mgl.Perspective(mgl.DegToRad(45.0),
		float32(width)/float32(height), 0.1, 100.0)
	shader.SetMat4("projection", projection)

	yVelocity, yPos := float32(0), float32(5)
	window.Run(func() {
		yVelocity, yPos = updatePosVel(window.GetDT(), yVelocity, yPos)

		gl.ClearColor(0.1, 0.1, 0.1, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		shader.SetMat4("view", camera.LookAt())

		// Draw the cube bouncing
		shader.SetVec3("col", mgl.Vec3{0.7, 0.3, 0.3})
		shader.SetMat4("model", mgl.Translate3D(0, 0, yPos).Mul4(
			mgl.Scale3D(0.25, 0.25, 0.25)))
		vao.Draw()

		// Draw the floor
		shader.SetVec3("col", mgl.Vec3{0.3, 0.7, 0.3})
		shader.SetMat4("model", mgl.Translate3D(0, 0, -0.25).Mul4(
			mgl.Scale3D(10, 10, 0.0001)))
		vao.Draw()

		window.PollEvents()
		window.SwapBuffers()
	})
}

const gravityForce = -10.0

func updatePosVel(dt, yVelocity, yPos float32) (float32, float32) {
	yVelocity += gravityForce * dt
	yPos += yVelocity * dt
	if yPos <= 0 {
		yVelocity *= -1
		yPos = 0.0
	}

	return yVelocity, yPos
}
