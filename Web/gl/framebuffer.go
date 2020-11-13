package gl

import (
	"github.com/gopherjs/gopherjs/js"

	"fmt"
)

const (
	FRAMEBUFFER       = 0x8D40
	RENDERBUFFER      = 0x8D41
	DEPTH_COMPONENT16 = 0x81A5
	//DEPTH24_STENCIL8         = 0x88F0 WEBGL2.0only
	DEPTH_ATTACHMENT = 0x8D00
	//DEPTH_STENCIL_ATTACHMENT = 0x821A
	COLOR_ATTACHMENT0 = 0x8CE0
)

// This function is broken... since webgl requires the framebuffer to be an
// object rather than an index. We can cheat it for now.
func BindFramebuffer(fType, index uint32) {
	if index != 0 {
		webgl.Call("bindFramebuffer", fType, index)
		return
	}
	webgl.Call("bindFramebuffer", fType, nil)
}

// Perhaps it makes sense to have texture textureColorbuffer as a
// texture.
// I don't understand framebuffers and usage as much as I would like to
// So ideally I want to revisit this for when usage needs more customization.
type Framebuffer struct {
	framebuffer        *js.Object
	textureColorbuffer *js.Object
	rbo                *js.Object
}

func NewFramebuffer(width, height int32) *Framebuffer {
	fmt.Println("HERE?")
	// Create framebuffer
	f := &Framebuffer{framebuffer: webgl.Call("createFramebuffer")}
	webgl.Call("bindFramebuffer", FRAMEBUFFER, f.framebuffer)
	//gl.GenFramebuffers(1, &f.framebuffer)
	//gl.BindFramebuffer(gl.FRAMEBUFFER, f.framebuffer)

	// Create color attachment texture
	f.textureColorbuffer = webgl.Call("createTexture")
	webgl.Call("bindTexture", TEXTURE_2D, f.textureColorbuffer)
	webgl.Call("texImage2D", TEXTURE_2D, 0, RGBA, width, height,
		0, RGBA, UNSIGNED_BYTE, nil)
	webgl.Call("texParameteri", TEXTURE_2D, TEXTURE_WRAP_S, CLAMP_TO_EDGE)
	webgl.Call("texParameteri", TEXTURE_2D, TEXTURE_WRAP_T, CLAMP_TO_EDGE)
	webgl.Call("texParameteri", TEXTURE_2D, TEXTURE_MIN_FILTER, LINEAR)
	webgl.Call("texParameteri", TEXTURE_2D, TEXTURE_MAG_FILTER, LINEAR)
	webgl.Call("framebufferTexture2D", FRAMEBUFFER, COLOR_ATTACHMENT0,
		TEXTURE_2D, f.textureColorbuffer, 0)
	//gl.GenTextures(1, &f.textureColorbuffer)
	//gl.BindTexture(gl.TEXTURE_2D, f.textureColorbuffer)
	//gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height,
	//0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(nil))
	//gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	//gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	//gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0,
	//gl.TEXTURE_2D, f.textureColorbuffer, 0)

	// Create a renderbuffer object for depth and stencil attachment
	f.rbo = webgl.Call("createRenderbuffer")
	webgl.Call("bindRenderbuffer", RENDERBUFFER, f.rbo)
	webgl.Call("renderbufferStorage", RENDERBUFFER, DEPTH_COMPONENT16,
		width, height)
	webgl.Call("framebufferRenderbuffer", FRAMEBUFFER, DEPTH_ATTACHMENT,
		RENDERBUFFER, f.rbo)

	//gl.GenRenderbuffers(1, &f.rbo)
	//gl.BindRenderbuffer(gl.RENDERBUFFER, f.rbo)
	//gl.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH24_STENCIL8,
	//width, height)
	//gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_STENCIL_ATTACHMENT,
	//	gl.RENDERBUFFER, f.rbo)

	// Ensure framebuffer is complete
	//if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
	//gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	//}

	fmt.Println("EXIT")

	return f
}

func (f *Framebuffer) Bind() {
	webgl.Call("bindFramebuffer", FRAMEBUFFER, f.framebuffer)
}

func (f *Framebuffer) BindTexture(v uint32) {
	webgl.Call("activeTexture", v)
	webgl.Call("bindTexture", TEXTURE_2D, f.textureColorbuffer)
}

//func
