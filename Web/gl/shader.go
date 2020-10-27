package gl

import (
	"errors"
	"github.com/gopherjs/gopherjs/js"
)

type Shader struct {
	shader *js.Object
}

const (
	COMPILE_STATUS  = 0x8B81
	VERTEX_SHADER   = 0x8B31
	FRAGMENT_SHADER = 0x8B30
)

var (
	currentBoundShader *js.Object
)

func CompileShader(vertexCode, fragCode string) (*Shader, error) {
	vertex := webgl.Call("createShader", VERTEX_SHADER)
	webgl.Call("shaderSource", vertex, vertexCode)
	webgl.Call("compileShader", vertex)
	if err := checkError(vertex); err != nil {
		return nil, err
	}

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

	return &Shader{shader}, nil
}

func checkError(shader *js.Object) error {
	if !webgl.Call("getShaderParameter", shader, COMPILE_STATUS).Bool() {
		return errors.New("Shader failed to compile")
	}
	return nil
}

func (s *Shader) Use() *Shader {
	webgl.Call("useProgram", s.shader)
	currentBoundShader = s.shader
	return s
}
