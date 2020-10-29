package gl

import (
	"github.com/gopherjs/gopherjs/js"
)

const (
	TEXTURE0 = 0x84C0
	TEXTURE1 = 0x84C1
	TEXTURE2 = 0x84C2
	TEXTURE3 = 0x84C3
	TEXTURE4 = 0x84C4
	TEXTURE5 = 0x84C5
	TEXTURE6 = 0x84C6
	TEXTURE7 = 0x84C7

	RGB  = 0x1907
	RGBA = 0x1908

	UNSIGNED_BYTE   = 0x1401
	REPEAT          = 0x2901
	CLAMP_TO_EDGE   = 0x812F
	MIRRORED_REPEAT = 0x8370

	NEAREST = 0x2600
	LINEAR  = 0x2601

	TEXTURE_2D = 0x0DE1

	TEXTURE_MAG_FILTER = 0x2800
	TEXTURE_MIN_FILTER = 0x2801
	TEXTURE_WRAP_S     = 0x2802
	TEXTURE_WRAP_T     = 0x2803
)

type Texture struct {
	texture        *js.Object
	Width          int32
	Height         int32
	InternalFormat int32
	ImageFormat    uint32
	WrapS          int32
	WrapT          int32
	FilterMin      int32
	FilterMax      int32
}

func NewTexture() *Texture {
	t := Texture{
		texture:        webgl.Call("createTexture"),
		InternalFormat: RGBA,
		ImageFormat:    RGBA,
		WrapS:          REPEAT,
		WrapT:          REPEAT,
		FilterMin:      LINEAR,
		FilterMax:      LINEAR,
	}

	return &t
}

func (t *Texture) Generate(width, height int32, data []byte) {
	t.Width, t.Height = width, height

	webgl.Call("bindTexture", t.texture)
	//webgl.Call("texImage2D", gl.TEXTURE_2D, 0, gl.RGBA,
	//	1, 1, 0, gl.RGBA, gl.UNSIGNED_BYTE, data)

	webgl.Call("texImage2D", TEXTURE_2D, 0, RGBA, RGBA,
		UNSIGNED_BYTE, data)
	t.setTextParams(width, height)
	/*
		var dataPtr unsafe.Pointer
		if data != nil {
			dataPtr = gl.Ptr(data)
		}
		gl.TexImage2D(gl.TEXTURE_2D, 0, t.InternalFormat, width, height, 0,
			t.ImageFormat, gl.UNSIGNED_BYTE, dataPtr)

		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, t.WrapS)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, t.WrapT)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, t.FilterMin)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, t.FilterMax)
		gl.BindTexture(gl.TEXTURE_2D, 0)
	*/
}

func isPowerOf2(value int32) bool {
	return (value & (value - 1)) == 0
}

func (t *Texture) setTextParams(width, height int32) {
	if isPowerOf2(width) && isPowerOf2(height) {
		webgl.Call("texParameteri", TEXTURE_2D, TEXTURE_WRAP_S, t.WrapS)
		webgl.Call("texParameteri", TEXTURE_2D, TEXTURE_WRAP_T, t.WrapT)
		webgl.Call("texParameteri", TEXTURE_2D, TEXTURE_MIN_FILTER, t.FilterMin)
		webgl.Call("texParameteri", TEXTURE_2D, TEXTURE_MAG_FILTER, t.FilterMax)
	} else {
		webgl.Call("texParameteri", TEXTURE_2D, TEXTURE_WRAP_S, CLAMP_TO_EDGE)
		webgl.Call("texParameteri", TEXTURE_2D, TEXTURE_WRAP_T, CLAMP_TO_EDGE)
		webgl.Call("texParameteri", TEXTURE_2D, TEXTURE_MIN_FILTER, LINEAR)
	}

}

func TextureFromFile(file string) *Texture {
	t := NewTexture()
	webgl.Call("bindTexture", t.texture)
	webgl.Call("texImage2D", TEXTURE_2D, 0, RGBA,
		1, 1, 0, RGBA, UNSIGNED_BYTE, []byte{39, 0, 0, 255})

	img := js.Global.Get("Image").New()
	img.Set("src", file)
	img.Set("crossOrigin", "")
	img.Call("addEventListener", "load", func() {
		webgl.Call("texImage2D", TEXTURE_2D, 0, RGBA, RGBA,
			UNSIGNED_BYTE, img)
		t.setTextParams(int32(img.Get("width").Int()),
			int32(img.Get("height").Int()))
	}, false)

	return t
}

func (t *Texture) Bind(num uint32) {
	webgl.Call("activeTexture", num)
	webgl.Call("bindTexture", TEXTURE_2D, t.texture)
}
