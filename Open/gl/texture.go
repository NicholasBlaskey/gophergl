package gl

import (
	"github.com/go-gl/gl/v4.1-core/gl"

	//"fmt"

	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"unsafe"
)

const (
	TEXTURE0 = gl.TEXTURE0
	TEXTURE1 = gl.TEXTURE1
	TEXTURE2 = gl.TEXTURE2
	TEXTURE3 = gl.TEXTURE3
	TEXTURE4 = gl.TEXTURE4
	TEXTURE5 = gl.TEXTURE5
	TEXTURE6 = gl.TEXTURE6
	TEXTURE7 = gl.TEXTURE7
)

type Texture struct {
	id             uint32
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
		InternalFormat: gl.RGBA,
		ImageFormat:    gl.RGBA,
		WrapS:          gl.REPEAT,
		WrapT:          gl.REPEAT,
		FilterMin:      gl.LINEAR,
		FilterMax:      gl.LINEAR,
	}
	gl.GenTextures(1, &t.id)
	return &t
}

func (t *Texture) Generate(width, height int32, data []byte) {
	t.Width, t.Height = width, height
	gl.BindTexture(gl.TEXTURE_2D, t.id)

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
}

func TextureFromFile(file string) (*Texture, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	t := NewTexture()
	t.Generate(int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y), rgba.Pix)
	return t, nil
}

func (t *Texture) Bind(num uint32) {
	gl.ActiveTexture(num)
	gl.BindTexture(gl.TEXTURE_2D, t.id)
}
