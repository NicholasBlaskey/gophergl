package gl

import (
	"unsafe"

	"errors"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Shader struct {
	id uint32
}

func CompileShader(vertexCode string, fragmentCode string) (*Shader, error) {
	// Compile the shaders
	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	shaderSource, freeVertex := gl.Strs(vertexCode + "\x00")
	defer freeVertex()
	gl.ShaderSource(vertexShader, 1, shaderSource, nil)
	gl.CompileShader(vertexShader)
	err := checkCompileErrors(vertexShader, "VERTEX")
	defer gl.DeleteShader(vertexShader)
	if err != nil {
		return nil, err
	}

	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	shaderSource, freeFragment := gl.Strs(fragmentCode + "\x00")
	defer freeFragment()
	gl.ShaderSource(fragmentShader, 1, shaderSource, nil)
	gl.CompileShader(fragmentShader)
	defer gl.DeleteShader(fragmentShader)
	err = checkCompileErrors(fragmentShader, "FRAGMENT")
	if err != nil {
		return nil, err
	}

	// Create a shader program
	id := gl.CreateProgram()
	gl.AttachShader(id, vertexShader)
	gl.AttachShader(id, fragmentShader)
	gl.LinkProgram(id)
	err = checkCompileErrors(id, "PROGRAM")
	if err != nil {
		return nil, err
	}

	return &Shader{id}, nil
}

func (s *Shader) Use() *Shader {
	gl.UseProgram(s.id)
	return s
}

func checkCompileErrors(shader uint32, shaderType string) error {
	var success int32
	var infoLog [1024]byte

	var status uint32 = gl.COMPILE_STATUS
	stageMessage := "Shader_Compilation_error"
	errorFunc := gl.GetShaderInfoLog
	getIV := gl.GetShaderiv
	if shaderType == "PROGRAM" {
		status = gl.LINK_STATUS
		stageMessage = "Program_link_error"
		errorFunc = gl.GetProgramInfoLog
		getIV = gl.GetProgramiv
	}

	getIV(shader, status, &success)
	if success != 1 {
		test := &success
		errorFunc(shader, 1024, test, (*uint8)(unsafe.Pointer(&infoLog)))
		return errors.New(stageMessage + shaderType +
			"|" + string(infoLog[:1024]) + "|")
	}
	return nil
}
