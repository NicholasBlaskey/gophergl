package gl

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// I don't understand framebuffers and usage as much as I would like to
// So ideally I want to revisit this for when usage needs more customization.
type Framebuffer struct {
	framebuffer        uint32
	textureColorbuffer uint32
	rbo                uint32
}

func New(width, height int32) *Framebuffer {
	// Create framebuffer
	f := &Framebuffer{}
	gl.GenFramebuffers(1, &f.framebuffer)
	gl.BindFramebuffer(gl.FRAMEBUFFER, framebuffer)

	// Create color attachment texture
	gl.GenTexture(1, &f.textureColorbuffer)
	gl.BindTexture(gl.TEXTURE_2D, f.textureColorBuffer)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height,
		0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(nil))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0,
		gl.TEXTURE_2D, textureColorbuffer, 0)

	// Create a renderbuffer object for depth and stencil attachment
	gl.GenRenderbuffers(1, &f.rbo)
	gl.BindRenderbuffer(gl.RENDERBUFFER, f.rbo)
	gl.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH24_STENCIL8,
		width, height)
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_STENCIL_ATTACHMENT,
		gl.RENDERBUFFER, rbo)

	// Ensure framebuffer is complete
	if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
		gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	}

	return f
}

func (f *Framebuffer) Bind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, f.framebuffer)
}

//func
