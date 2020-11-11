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

	cubemapVertex = `#version 410 core
	layout (location = 0) in vec3 aPos;

	out vec3 uv;

	uniform mat4 projection;
	uniform mat4 view;

	void main() 
	{
		uv = aPos;
		vec4 pos = projection * view * vec4(aPos, 1.0);
		gl_Position = pos.xyww;
	}`

	cubemapFrag = `#version 410 core
	out vec4 FragColor;

	in vec3 uv;
	uniform samplerCube cubemap;
	void main() 
	{
		FragColor = texture(cubemap, uv);
	}`
)

func main() {
	width, height := int32(800), int32(600)
	window, err := gl.NewWindow(width, height, "cubemap")
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

	cubemapShader, err := gl.CompileShader(cubemapVertex, cubemapFrag)
	if err != nil {
		panic(err)
	}

	t1, err := gl.TextureFromFile("./images/gopher.jpg")
	if err != nil {
		panic(err)
	}

	dir := "./images/cubemap/"
	cubemap := gl.Cubemap{
		Right:  dir + "xpos.png",
		Left:   dir + "xneg.png",
		Front:  dir + "zpos.png",
		Back:   dir + "zneg.png",
		Top:    dir + "ypos.png",
		Bottom: dir + "yneg.png",
	}
	err = cubemap.Load()
	if err != nil {
		panic(err)
	}

	projection := mgl.Perspective(mgl.DegToRad(45.0),
		float32(width)/float32(height), 0.1, 100.0)

	shader.Use()
	shader.SetInt("texture1", 0)
	shader.SetMat4("projection", projection)

	cubemapShader.Use()
	cubemapShader.SetInt("skybox", 0)
	cubemapShader.SetMat4("projection", projection)
	window.Run(func() {
		gl.ClearColor(0.1, 0.1, 0.1, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		view := camera.LookAt()

		// Draw our cube
		t1.Bind(gl.TEXTURE0)
		shader.Use().SetMat4("view", view)
		shader.SetMat4("model", mgl.Translate3D(0.0, 0.0, 0.0).Mul4(
			mgl.Scale3D(0.25, 0.25, 0.25)))
		vao.Draw()

		// Draw our cubemap
		gl.DepthFunc(gl.LEQUAL)
		view.SetCol(3, mgl.Vec4{0, 0, 0, 0}) // Strip translation from view
		view.SetRow(3, mgl.Vec4{0, 0, 0, 0})
		cubemapShader.Use().SetMat4("view", view)
		cubemapShader.SetMat4("projection", projection)
		cubemap.Bind(gl.TEXTURE0)
		vao.Draw()
		gl.DepthFunc(gl.LESS)

		window.PollEvents()
		window.SwapBuffers()
	})
}
