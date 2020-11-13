package main

import (
	"runtime"

	"fmt"

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
		FragColor = texture(texture1, uv);	
	}`
)

func main() {
	width, height := int32(800), int32(600)
	window, err := gl.NewWindow(width, height, "framebuffer")
	if err != nil {
		panic(err)
	}
	defer window.Terminate()
	camera := gl.NewOrbitalCamera(window, 5.0, mgl.Vec3{0.0, 0.0, 0.0})

	gl.Enable(gl.DEPTH_TEST)

	fmt.Println("b4 shader")
	shader, err := gl.CompileShader(vertexShader, fragShader)
	if err != nil {
		panic(err)
	}
	shader.Use()

	fmt.Println("POST SHADER")
	vao := gl.NewVAO(gl.NewCube(gl.VertParams{Position: true, TexCoords: true}))

	fmt.Println("B4 texturE")
	t1, err := gl.TextureFromFile("./images/gopher.jpg")
	if err != nil {
		panic(err)
	}
	fmt.Println("B4 buffer")
	framebuffer := gl.NewFramebuffer(width, height)

	fmt.Println("POST FBO")
	projection := mgl.Perspective(mgl.DegToRad(45.0),
		float32(width)/float32(height), 0.1, 100.0)

	shader.Use().SetMat4("projection", projection)
	shader.SetInt("texture1", 0)

	//gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	firstRun := true
	window.Run(func() {
		shader.Use().SetMat4("view", camera.LookAt())
		shader.SetMat4("model", mgl.HomogRotate3D(mgl.DegToRad(45.0),
			mgl.Vec3{0.0, 1.0, 1.0}.Normalize()).Mul4(
			mgl.Scale3D(1.25, 1.25, 1.25)))

		// Render scene to framebuffer (either once or keep going remove the || true if only once)
		if firstRun || true {
			framebuffer.Bind()
			t1.Bind(gl.TEXTURE0)
			gl.Enable(gl.DEPTH_TEST)
			gl.ClearColor(1.0, 1.0, 1.0, 1.0)
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

			vao.Draw()

			firstRun = false
		}

		{ // Render
			gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
			//gl.Disable(gl.DEPTH_TEST)
			gl.ClearColor(0.1, 0.1, 0.1, 0.1)
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

			framebuffer.BindTexture(gl.TEXTURE0)
			vao.Draw()
		}

		window.PollEvents()
		window.SwapBuffers()
	})
}
