package gl

import (
	"github.com/gopherjs/gopherjs/js"
)

const (
	FRAMEBUFFER              = 0x8D40
	RENDERBUFFER             = 0x8D41
	DEPTH24_STENCIL8         = 0x88F0
	DEPTH_STENCIL_ATTACHMENT = 0x821A
)

func BindFramebuffer(fType, index uint32) {
	webgl.Call("bindFramebuffer", fType, index)
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
	// Create framebuffer
	f := &Framebuffer{framebuffer: webgl.Call("createFramebuffer")}
	webgl.Call("bindFramebuffer", FRAMEBUFFER, f.framebuffer)
	//gl.GenFramebuffers(1, &f.framebuffer)
	//gl.BindFramebuffer(gl.FRAMEBUFFER, f.framebuffer)

	// Create color attachment texture
	f.textureColorBuffer = webgl.Call("createTexture")
	webgl.Call("bindTexture", TEXTURE_2D, f.textureColorBuffer)
	webgl.Call("texImage2D", gl.TEXTURE_2D, 0, RGBA, width, height,
		0, RGBA, UNSIGNED_BYTE, nil)
	webgl.Call("texParameteri", TEXTURE_2D, TEXTURE_MIN_FILTER, LINEAR)
	webgl.Call("texParameteri", TEXTURE_2D, TEXTURE_MAG_FILTER, LINEAR)
	webgl.Call("framebufferTexture2D", FRAMEBUFFER, COLOR_ATTACHMENT0,
		TEXTURE_2D, f.textureColorBuffer, 0)
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
	webgl.Call("renderBufferStorage", RENDERBUFFER, DEPTH24_STENCIL8,
		width, height)
	webgl.Call("framebufferRenderbuffer", FRAMEBUFFER, DEPTH_STENCIL_ATTACHMENT,
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

	return f
}

func (f *Framebuffer) Bind() {
	webgl.Call("bindFramebuffer", gl.FRAMEBUFFER, f.framebuffer)
}

func (f *Framebuffer) BindTexture(v uint32) {
	webgl.Call("activeTexture", v)
	webgl.Call("bindTexture", gl.TEXTURE_2D, f.textureColorbuffer)
}

//func
