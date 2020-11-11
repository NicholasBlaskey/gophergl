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

	shader, err := gl.CompileShader(vertexShader, fragShader)
	if err != nil {
		panic(err)
	}
	shader.Use()

	vao := gl.NewVAO(gl.NewCube(gl.VertParams{Position: true, TexCoords: true}))
	t1, err := gl.TextureFromFile("./images/gopher.jpg")
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

		shader.SetMat4("model", mgl.HomogRotate3D(mgl.DegToRad(45.0),
			mgl.Vec3{0.0, 1.0, 1.0}.Normalize()).Mul4(
			mgl.Scale3D(0.25, 0.25, 0.25)))
		vao.Draw()

		window.PollEvents()
		window.SwapBuffers()
	})
}

/*

	// Frambuffer config
	var framebuffer uint32
	gl.GenFramebuffers(1, &framebuffer)
	gl.BindFramebuffer(gl.FRAMEBUFFER, framebuffer)
	// Create a color attachment texture
	var textureColorbuffer uint32
	gl.GenTextures(1, &textureColorbuffer)
	gl.BindTexture(gl.TEXTURE_2D, textureColorbuffer)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, windowWidth, windowHeight,
		0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(nil))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0,
		gl.TEXTURE_2D, textureColorbuffer, 0)
	// Create a renderbuffer object for depth and stencil atachment
	var rbo uint32
	gl.GenRenderbuffers(1, &rbo)
	gl.BindRenderbuffer(gl.RENDERBUFFER, rbo)
	gl.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH24_STENCIL8,
		windowWidth, windowHeight)
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_STENCIL_ATTACHMENT,
		gl.RENDERBUFFER, rbo)
	// Ensure framebuffer is complete
	if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
		gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	}


//

		// Binf frame buffer
		gl.BindFramebuffer(gl.FRAMEBUFFER, framebuffer)
		gl.Enable(gl.DEPTH_TEST)

		// Bind back original frame buffer
		gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
		gl.Disable(gl.DEPTH_TEST)
		// Clear all relevant buffers
		gl.ClearColor(1.0, 1.0, 1.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

*/
