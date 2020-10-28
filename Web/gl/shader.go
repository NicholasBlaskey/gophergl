package gl

import (
	"errors"
	"fmt"
	"strings"

	mgl "github.com/go-gl/mathgl/mgl32"
	"github.com/gopherjs/gopherjs/js"
)

type Shader struct {
	shader      *js.Object
	attribNames []string
}

const (
	COMPILE_STATUS  = 0x8B81
	VERTEX_SHADER   = 0x8B31
	FRAGMENT_SHADER = 0x8B30
)

var (
	currentBoundShader *Shader
)

func CompileShader(vertexCode, fragCode string) (*Shader, error) {
	vertexCode, attribNames := convertToWebShader(vertexCode, true)
	vertex := webgl.Call("createShader", VERTEX_SHADER)
	webgl.Call("shaderSource", vertex, vertexCode)
	webgl.Call("compileShader", vertex)
	if err := checkError(vertex); err != nil {
		return nil, err
	}

	fragCode, _ = convertToWebShader(fragCode, false)
	frag := webgl.Call("createShader", FRAGMENT_SHADER)
	webgl.Call("shaderSource", frag, fragCode)
	webgl.Call("compileShader", frag)
	if err := checkError(frag); err != nil {
		return nil, err
	}

	shader := webgl.Call("createProgram")
	webgl.Call("attachShader", shader, vertex)
	webgl.Call("attachShader", shader, frag)
	webgl.Call("linkProgram", shader)

	return &Shader{shader, attribNames}, nil
}

// TODO actually make this work properly instead of just hacks
func convertToWebShader(shader string, isVertex bool) (string, []string) {
	attribNames := []string{}

	// Remove version tag
	shader = strings.ReplaceAll(shader, "#version 410 core", "")

	// Set a default precision if fragment shader
	if !isVertex {
		shader = "precision mediump float;\n" + shader

		// Remove out vec4 color;
		shader = strings.ReplaceAll(shader, "out vec4 FragColor;", "")
		// Replace color with gl_FragColor
		shader = strings.ReplaceAll(shader, "FragColor", "gl_FragColor")
		// Replace in with varying
		shader = strings.ReplaceAll(shader, "in ", "varying ")
	} else {
		// TODO make this run instead of just relying on location
		// Replace layout with attribute
		for i := 0; i < 10; i++ {
			shader = strings.ReplaceAll(shader,
				fmt.Sprintf("layout (location = %d) in", i), "attribute")
		}

		// Replace out with varying
		shader = strings.ReplaceAll(shader, "out", "varying")

		// Get all attribute names
		for _, v := range strings.Split(shader, "\n") {
			if strings.Contains(v, "void main()") {
				break
			}

			if strings.Contains(v, "attribute") {
				bySpace := strings.Split(v, " ")
				attribName := bySpace[len(bySpace)-1]
				attribNames = append(attribNames, attribName[:len(attribName)-1])
			}
		}
	}

	fmt.Println(shader)
	//fmt.Println("ATTRIBS")
	//fmt.Println(len(attribNames), attribNames)

	return shader, attribNames
}

func checkError(shader *js.Object) error {
	if !webgl.Call("getShaderParameter", shader, COMPILE_STATUS).Bool() {
		return errors.New(webgl.Call("getShaderInfoLog", shader).String())
	}
	return nil
}

func (s *Shader) Use() *Shader {
	webgl.Call("useProgram", s.shader)
	currentBoundShader = s
	return s
}

func (s *Shader) SetBool(name string, value bool) *Shader {
	var intValue int32 = 0
	if value {
		intValue = 1
	}

	webgl.Call("uniform1i", webgl.Call("getUniformLocation", s.shader, name),
		intValue)
	return s
}

func (s *Shader) SetInt(name string, value int32) *Shader {
	webgl.Call("uniform1i", webgl.Call("getUniformLocation", s.shader, name),
		value)
	return s
}

func (s *Shader) SetFloat(name string, value float32) *Shader {
	webgl.Call("uniform1f", webgl.Call("getUniformLocation", s.shader, name),
		value)

	return s
}

func (s *Shader) SetVec2(name string, value mgl.Vec2) *Shader {
	webgl.Call("uniform2fv", webgl.Call("getUniformLocation", s.shader, name),
		value) // TODO verify this
	return s
}

func (s *Shader) SetVec3(name string, value mgl.Vec3) *Shader {
	webgl.Call("uniform3fv", webgl.Call("getUniformLocation", s.shader, name),
		value) // TODO verify this
	return s
}

func (s *Shader) SetMat4(name string, value mgl.Mat4) *Shader {
	webgl.Call("uniformMatrix4fv", webgl.Call("getUniformLocation", s.shader, name),
		value) // TODO verify this
	return s
}
