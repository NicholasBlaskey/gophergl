package gl

import (
	//"errors"

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
	TEXTURE_WRAP_R     = 0x8072
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
	webgl.Call("texImage2D", TEXTURE_2D, 0, RGBA, RGBA,
		UNSIGNED_BYTE, data)
	t.setTextParams(width, height)
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
		webgl.Call("texParameteri", TEXTURE_2D, TEXTURE_MAG_FILTER, LINEAR)
	}

}

func TextureFromFile(file string) (*Texture, error) {
	t := NewTexture()
	webgl.Call("bindTexture", TEXTURE_2D, t.texture)
	webgl.Call("texImage2D", TEXTURE_2D, 0, RGBA,
		1, 1, 0, RGBA, UNSIGNED_BYTE, []byte{39, 0, 0, 255})

	img := js.Global.Get("Image").New()
	img.Set("src", file)

	h, w := img.Get("height").Int(), img.Get("width").Int()

	img.Call("addEventListener", "load", func() {
		webgl.Call("bindTexture", TEXTURE_2D, t.texture)
		webgl.Call("texImage2D", TEXTURE_2D, 0, t.InternalFormat,
			t.ImageFormat, UNSIGNED_BYTE, img)
		t.setTextParams(int32(w), int32(h))
	}, false)

	return t, nil
}

func (t *Texture) Bind(num uint32) {
	webgl.Call("activeTexture", num)
	webgl.Call("bindTexture", TEXTURE_2D, t.texture)
}

type Cubemap struct {
	Right   string
	Left    string
	Top     string
	Bottom  string
	Front   string
	Back    string
	texture *js.Object
}

func (cm *Cubemap) Load() error {
	cm.texture = webgl.Call("createTexture")
	webgl.Call("bindTexture", TEXTURE_CUBE_MAP, cm.texture)

	// Bind a single pixel for texture for the time being
	//webgl.Call("texImage2D", TEXTURE_2D, 0, RGBA,
	//1, 1, 0, RGBA, UNSIGNED_BYTE, []byte{39, 0, 0, 255})

	for i, path := range []string{cm.Right, cm.Left,
		cm.Top, cm.Bottom, cm.Front, cm.Back} {

		img := js.Global.Get("Image").New()
		img.Set("src", path)

		img.Call("addEventListener", "load", func() {
			webgl.Call("bindTexture", TEXTURE_CUBE_MAP, cm.texture)
			webgl.Call("texImage2D", TEXTURE_CUBE_MAP_POSITIVE_X+uint32(i),
				0, RGBA, RGBA, UNSIGNED_BYTE, img)
		}, false)
	}

	webgl.Call("texParameteri", TEXTURE_2D, TEXTURE_MIN_FILTER, LINEAR)
	webgl.Call("texParameteri", TEXTURE_2D, TEXTURE_MAG_FILTER, LINEAR)
	webgl.Call("texParameteri", TEXTURE_2D, TEXTURE_WRAP_S, CLAMP_TO_EDGE)
	webgl.Call("texParameteri", TEXTURE_2D, TEXTURE_WRAP_T, CLAMP_TO_EDGE)
	webgl.Call("texParameteri", TEXTURE_2D, TEXTURE_WRAP_R, CLAMP_TO_EDGE)

	return nil
}

const (
	TEXTURE_CUBE_MAP            = 0x8513
	TEXTURE_CUBE_MAP_POSITIVE_X = 0x8515
)

func (cm *Cubemap) Bind(num uint32) {
	webgl.Call("activeTexture", num)
	webgl.Call("bindTexture", TEXTURE_CUBE_MAP, cm.texture)
}
