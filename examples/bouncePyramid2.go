package main

import (
	"runtime"

	mgl "github.com/go-gl/mathgl/mgl32"

	"math"
	"math/rand"

	"github.com/nicholasblaskey/gophergl/Open/gl"
	//	"github.com/nicholasblaskey/gophergl/Web/gl"
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
	window, err := gl.NewWindow(width, height, "bounce pyramid 2")
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
		float32(width)/float32(height), 0.1, 300.0)
	shader.SetMat4("projection", projection)

	// Create the cube pyramid
	startingHeight := float32(50)
	levelDiff := float32(0.01)
	cubeWidth := float32(0.10)
	dim := 81 // Only works for odd values as a simplification
	center := float32(dim/2) * cubeWidth * 2
	cubeVelocities := []float32{}
	cubePositions := []mgl.Vec3{}
	cubeCols := []mgl.Vec3{}

	colors := []mgl.Vec3{}
	for i := 0; i < dim/2+1; i++ {
		colors = append(colors, mgl.Vec3{
			rand.Float32(), rand.Float32(), rand.Float32()})
	}

	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			x, y := float32(i)*cubeWidth*2, float32(j)*cubeWidth*2

			level := float32(math.Round(math.Max(
				math.Abs(float64(center-x)),
				math.Abs(float64(center-y)),
			) / float64(cubeWidth*2)))

			cubePositions = append(cubePositions, mgl.Vec3{
				x, y, startingHeight})
			cubeCols = append(cubeCols, colors[int(level)])
			cubeVelocities = append(cubeVelocities, 0.0)

			index := len(cubePositions) - 1
			for i := 0; i < int(level); i++ {
				vel, pos := updatePosVel(levelDiff, cubeVelocities[index], cubePositions[index][2])
				cubeVelocities[index] = vel
				cubePositions[index][2] = pos
			}

		}
	}

	window.Run(func() {
		gl.ClearColor(0.1, 0.1, 0.1, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		shader.SetMat4("view", camera.LookAt())

		dt := window.GetDT()
		for i := 0; i < len(cubeCols); i++ {
			vel, pos := updatePosVel(dt, cubeVelocities[i], cubePositions[i][2])
			cubeVelocities[i] = vel
			cubePositions[i][2] = pos

			shader.SetVec3("col", cubeCols[i])
			shader.SetMat4("model", mgl.Translate3D(
				cubePositions[i][0], cubePositions[i][1],
				cubePositions[i][2]).Mul4(
				mgl.Scale3D(cubeWidth, cubeWidth, cubeWidth)))
			vao.Draw()
		}

		// Draw the floor
		shader.SetVec3("col", mgl.Vec3{0.3, 0.7, 0.3})
		shader.SetMat4("model", mgl.Translate3D(0, 0, -0.25).Mul4(
			mgl.Scale3D(100, 100, 0.0001)))
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
