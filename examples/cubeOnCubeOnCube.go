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
	void main()
	{	
		if (uv.x >= 0.98 || uv.y >= 0.98 || uv.x <= 0.02 || uv.y <= 0.02) {
			FragColor = vec4(0.0, 0.0, 0.0, 1.0);
		} else {
			FragColor = texture(texture1, uv);	
		}
	}`
)

func main() {
	width, height := int32(800), int32(600)
	window, err := gl.NewWindow(width, height, "cube on cube on cube on cube")
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

	vao := gl.NewVAO(gl.NewCube(gl.VertParams{Position: true, TexCoords: true}))

	buffer1 := gl.NewFramebuffer(width, height)
	buffer2 := gl.NewFramebuffer(width, height)

	projection := mgl.Perspective(mgl.DegToRad(45.0),
		float32(width)/float32(height), 0.1, 100.0)

	shader.Use().SetMat4("projection", projection)
	shader.SetInt("texture1", 0)

	firstRun := true
	window.Run(func() {
		shader.Use().SetMat4("view", camera.LookAt())
		shader.SetMat4("model", mgl.HomogRotate3D(mgl.DegToRad(45.0),
			mgl.Vec3{0.0, 1.0, 1.0}.Normalize()).Mul4(
			mgl.Scale3D(1.25, 1.25, 1.25)))

		// Render scene to framebuffer (either once or keep going remove the || true if only once)
		var endingBuffer *gl.Framebuffer
		if firstRun || true {
			colors := []mgl.Vec3{
				mgl.Vec3{0.7, 0.3, 0.3},
				mgl.Vec3{0.3, 0.7, 0.3},
				mgl.Vec3{0.3, 0.3, 0.7},
				mgl.Vec3{0.7, 0.7, 0.3},
				mgl.Vec3{0.3, 0.7, 0.7},
				mgl.Vec3{0.7, 0.3, 0.7},
				mgl.Vec3{0.65, 0.75, 0.85},
				mgl.Vec3{0.75, 0.65, 0.85},
				mgl.Vec3{1.0, 1.0, 1.0},
			}
			for i, col := range colors {
				if i%2 != 0 {
					buffer2.Bind()
					buffer1.BindTexture(gl.TEXTURE0)
					endingBuffer = buffer2
				} else {
					buffer1.Bind()
					buffer2.BindTexture(gl.TEXTURE0)
					endingBuffer = buffer1
				}
				gl.ClearColor(col[0], col[1], col[2], 1.0)
				gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
				vao.Draw()
			}
			firstRun = false
		}

		{ // Render to screen
			gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
			gl.ClearColor(0.1, 0.1, 0.1, 1.0)
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

			endingBuffer.BindTexture(gl.TEXTURE0)
			vao.Draw()
		}

		window.PollEvents()
		window.SwapBuffers()
	})
}
