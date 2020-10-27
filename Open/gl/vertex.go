package gl

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type VAO struct {
	VAO           uint32
	VBO           uint32
	VertexAmount  int32
	PrimitiveType uint32
}

const (
	TRIANGLES = gl.TRIANGLES
)

func NewVAO(primitiveType uint32, offsets []int32, verts []float32) *VAO {
	sumOffsets := int32(0)
	for _, offset := range offsets {
		sumOffsets += offset
	}
	v := &VAO{0, 0, int32(len(verts)) / sumOffsets, primitiveType}
	gl.GenVertexArrays(1, &v.VAO)
	gl.GenBuffers(1, &v.VBO)

	gl.BindBuffer(gl.ARRAY_BUFFER, v.VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(verts)*4, gl.Ptr(verts), gl.STATIC_DRAW)

	gl.BindVertexArray(v.VAO)
	offsetAmount := int32(0)
	for i, offset := range offsets {
		gl.EnableVertexAttribArray(uint32(i))
		gl.VertexAttribPointer(uint32(i), offset, gl.FLOAT,
			false, 4*offset, gl.PtrOffset(int(offsetAmount)))
		offsetAmount += offset
	}
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	return v
}

func (v *VAO) Draw() {
	gl.BindVertexArray(v.VAO)
	gl.DrawArrays(v.PrimitiveType, 0, v.VertexAmount)
	gl.BindVertexArray(0)
}
