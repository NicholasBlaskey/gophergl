package main

import (
	"runtime"

	mgl "github.com/go-gl/mathgl/mgl32"
	"math"

	//"github.com/nicholasblaskey/gophergl/Open/gl"
	"github.com/nicholasblaskey/gophergl/Web/gl"
)

func init() {
	runtime.LockOSThread()
}

const (
	vertexShader = `#version 410 core
	layout (location = 0) in vec3 aPosition;
	layout (location = 1) in vec3 aNormal;
	layout (location = 2) in vec2 aUV;
	
	out vec3 position;
	out vec3 normal;
	out vec2 uv;

	uniform mat4 projection;
	uniform mat4 view;
	uniform mat4 model;
	uniform mat4 normalMatrix;	

	void main()
	{
		position = vec3(model * vec4(aPosition, 1.0));
		normal = normalize(mat3(normalMatrix) * aNormal);
		uv = aUV;

		gl_Position = projection * view * model * vec4(position, 1.0);
	}`

	fragShader = `#version 410 core
	out vec4 FragColor;	

	in vec3 position;
	in vec3 normal;
	in vec2 uv;

	struct Material {
		sampler2D diffuse;
		sampler2D specular;
		sampler2D glow;
		float glowIntensity;
		float shininess;
	};

	struct Light {
		vec3 position;
		vec3 diffuse;
		vec3 specular;
		vec3 ambient;
	};

	uniform vec3 viewPos;
	uniform Material material;
	uniform Light light;

	void main()
	{	
		vec3 ambient = light.ambient * texture(material.diffuse, uv).rgb;
		
		vec3 norm = normalize(normal);
		vec3 lightDir = normalize(light.position - position);
		float diff = max(dot(norm, lightDir), 0.0);
		vec3 diffuse = light.diffuse * diff * texture(material.diffuse, uv).rgb;

		vec3 viewDir = normalize(viewPos - position);
		vec3 reflectDir = reflect(-lightDir, norm);
		float spec = pow(max(dot(viewDir, reflectDir), 0.0), material.shininess);
		vec3 specular = light.specular * spec * texture(material.specular, uv).rgb;

		vec3 emission = texture(material.glow, uv).rgb * material.glowIntensity;		

		FragColor = vec4(ambient + diffuse + specular + emission, 1.0);
		//FragColor = texture(material.glow, uv);
	}`

	lampVertex = `#version 410 core
	layout (location = 0) in vec3 aPosition;

	uniform mat4 model;
	uniform mat4 view;
	uniform mat4 projection;

	void main() 
	{
		gl_Position = projection * view * model * vec4(aPosition, 1.0);
	}`

	lampFrag = `#version 410 core
	out vec4 FragColor;

	void main() 
	{
		FragColor = vec4(1.0);
	}`
)

func main() {
	width, height := int32(800), int32(600)
	window, err := gl.NewWindow(width, height, "texture")
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
	lampShader, err := gl.CompileShader(lampVertex, lampFrag)
	if err != nil {
		panic(err)
	}

	shader.Use()
	vao := gl.NewVAO(gl.TRIANGLES, []int32{3, 3, 2}, []float32{
		// positions          // normals           // texture coords
		-0.5, -0.5, -0.5, 0.0, 0.0, -1.0, 0.0, 0.0,
		0.5, -0.5, -0.5, 0.0, 0.0, -1.0, 1.0, 0.0,
		0.5, 0.5, -0.5, 0.0, 0.0, -1.0, 1.0, 1.0,
		0.5, 0.5, -0.5, 0.0, 0.0, -1.0, 1.0, 1.0,
		-0.5, 0.5, -0.5, 0.0, 0.0, -1.0, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 0.0, -1.0, 0.0, 0.0,

		-0.5, -0.5, 0.5, 0.0, 0.0, 1.0, 0.0, 0.0,
		0.5, -0.5, 0.5, 0.0, 0.0, 1.0, 1.0, 0.0,
		0.5, 0.5, 0.5, 0.0, 0.0, 1.0, 1.0, 1.0,
		0.5, 0.5, 0.5, 0.0, 0.0, 1.0, 1.0, 1.0,
		-0.5, 0.5, 0.5, 0.0, 0.0, 1.0, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0, 1.0, 0.0, 0.0,

		-0.5, 0.5, 0.5, -1.0, 0.0, 0.0, 1.0, 0.0,
		-0.5, 0.5, -0.5, -1.0, 0.0, 0.0, 1.0, 1.0,
		-0.5, -0.5, -0.5, -1.0, 0.0, 0.0, 0.0, 1.0,
		-0.5, -0.5, -0.5, -1.0, 0.0, 0.0, 0.0, 1.0,
		-0.5, -0.5, 0.5, -1.0, 0.0, 0.0, 0.0, 0.0,
		-0.5, 0.5, 0.5, -1.0, 0.0, 0.0, 1.0, 0.0,

		0.5, 0.5, 0.5, 1.0, 0.0, 0.0, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 0.0, 0.0, 1.0, 1.0,
		0.5, -0.5, -0.5, 1.0, 0.0, 0.0, 0.0, 1.0,
		0.5, -0.5, -0.5, 1.0, 0.0, 0.0, 0.0, 1.0,
		0.5, -0.5, 0.5, 1.0, 0.0, 0.0, 0.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0, 0.0, 1.0, 0.0,

		-0.5, -0.5, -0.5, 0.0, -1.0, 0.0, 0.0, 1.0,
		0.5, -0.5, -0.5, 0.0, -1.0, 0.0, 1.0, 1.0,
		0.5, -0.5, 0.5, 0.0, -1.0, 0.0, 1.0, 0.0,
		0.5, -0.5, 0.5, 0.0, -1.0, 0.0, 1.0, 0.0,
		-0.5, -0.5, 0.5, 0.0, -1.0, 0.0, 0.0, 0.0,
		-0.5, -0.5, -0.5, 0.0, -1.0, 0.0, 0.0, 1.0,

		-0.5, 0.5, -0.5, 0.0, 1.0, 0.0, 0.0, 1.0,
		0.5, 0.5, -0.5, 0.0, 1.0, 0.0, 1.0, 1.0,
		0.5, 0.5, 0.5, 0.0, 1.0, 0.0, 1.0, 0.0,
		0.5, 0.5, 0.5, 0.0, 1.0, 0.0, 1.0, 0.0,
		-0.5, 0.5, 0.5, 0.0, 1.0, 0.0, 0.0, 0.0,
		-0.5, 0.5, -0.5, 0.0, 1.0, 0.0, 0.0, 1.0,
	})

	// Load textures
	diffuse, err1 := gl.TextureFromFile("./images/glowCube/diffus.png")
	specular, err2 := gl.TextureFromFile("./images/glowCube/specular.png")
	glow, err3 := gl.TextureFromFile("./images/glowCube/glow.png")
	if err1 != nil || err2 != nil || err3 != nil {
		panic("Could not open a texture")
	}

	// Set lighting details
	shader.Use()
	shader.SetInt("material.diffuse", 0)
	shader.SetInt("material.specular", 1)
	shader.SetInt("material.glow", 2)
	shader.SetFloat("material.shininess", 64.0)

	lightPos := mgl.Vec3{1.2, 1.0, 2.0}
	shader.SetVec3("light.ambient", mgl.Vec3{0.3, 0.3, 0.3})
	shader.SetVec3("light.diffuse", mgl.Vec3{0.55, 0.55, 0.55})
	shader.SetVec3("light.specular", mgl.Vec3{1.0, 1.0, 1.0})

	projection := mgl.Perspective(mgl.DegToRad(45.0),
		float32(width)/float32(height), 0.1, 100.0)
	shader.SetMat4("projection", projection)

	lampShader.Use()
	lampShader.SetMat4("projection", projection)

	window.Run(func() {
		gl.ClearColor(0.1, 0.1, 0.1, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		diffuse.Bind(gl.TEXTURE0)
		specular.Bind(gl.TEXTURE1)
		glow.Bind(gl.TEXTURE2)

		// Draw cube
		view := camera.LookAt()
		model := mgl.HomogRotate3D(mgl.DegToRad(45.0),
			mgl.Vec3{1.0, 1.0, 0.0}.Normalize())

		shader.Use()
		shader.SetVec3("light.position", lightPos)

		shader.SetMat4("view", view)
		shader.SetVec3("viewPos", camera.Position)
		shader.SetFloat("material.glowIntensity",
			float32(math.Sin(float64(window.GetTime()))))
		shader.SetMat4("model", model)
		shader.SetMat4("normalMatrix", model.Inv().Transpose())
		vao.Draw()

		// Draw lamp
		model = mgl.Translate3D(lightPos[0], lightPos[1], lightPos[2]).Mul4(
			mgl.Scale3D(0.25, 0.25, 0.25))
		lampShader.Use()
		lampShader.SetMat4("view", view)
		lampShader.SetMat4("model", model)
		vao.Draw()

		lightPos[0] = float32(math.Sin(float64(window.GetTime())))
		lightPos[1] = float32(math.Cos(float64(window.GetTime())))

		window.PollEvents()
		window.SwapBuffers()
	})
}
