package gl

import (
	"github.com/go-gl/gl/v4.1-core/gl"

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

func loadImage(file string) ([]byte, int32, int32, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, 0, 0, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, 0, 0, err
	}

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	return rgba.Pix, int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y), nil
}

func TextureFromFile(file string) (*Texture, error) {
	data, width, height, err := loadImage(file)
	if err != nil {
		return nil, err
	}

	t := NewTexture()
	t.Generate(width, height, data)
	return t, nil
}

func (t *Texture) Bind(num uint32) {
	gl.ActiveTexture(num)
	gl.BindTexture(gl.TEXTURE_2D, t.id)
}

type Cubemap struct {
	Right     string
	Left      string
	Top       string
	Bottom    string
	Front     string
	Back      string
	textureID uint32
}

func (cm *Cubemap) Load() error {
	var textureID uint32
	gl.GenTextures(1, &textureID)
	gl.BindTexture(gl.TEXTURE_CUBE_MAP, textureID)

	for i, path := range []string{cm.Right, cm.Left,
		cm.Top, cm.Bottom, cm.Front, cm.Back} {

		data, width, height, err := loadImage(path)
		if err != nil {
			return err
		}

		gl.TexImage2D(gl.TEXTURE_CUBE_MAP_POSITIVE_X+uint32(i), 0, gl.RGBA,
			width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(data))
	}
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_S,
		gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_T,
		gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_R,
		gl.CLAMP_TO_EDGE)

	cm.textureID = textureID

	return nil
}

func (cm *Cubemap) Bind(num uint32) {
	gl.ActiveTexture(num)
	gl.BindTexture(gl.TEXTURE_CUBE_MAP, cm.textureID)
}
