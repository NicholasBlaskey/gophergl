package gl

import (
	"fmt"
)

type VAO struct {
	VAO           uint32
	VBO           uint32
	VertexAmount  int32
	PrimitiveType uint32
}

const (
	ARRAY_BUFFER = 0x8892
	STATIC_DRAW  = 0x88E4

	POINTS         = 0x0000
	LINES          = 0x0001
	LINE_LOOP      = 0x0002
	LINE_STRIP     = 0x0003
	TRIANGLES      = 0x0004
	TRIANGLE_STRIP = 0x0005
	TRIANGLE_FAN   = 0x0006

	FLOAT = 0x1406
)

func NewVAO(primitiveType uint32, offsets []int32, verts []float32) *VAO {
	sumOffsets := int32(0)
	for _, offset := range offsets {
		sumOffsets += offset
	}

	VBO := webgl.Call("createBuffer")
	webgl.Call("bindBuffer", ARRAY_BUFFER, VBO)
	webgl.Call("bufferData", ARRAY_BUFFER, verts, STATIC_DRAW)

	offsetAmount := int32(0)
	for i, offset := range offsets {
		fmt.Println("IN CREATE VAO", currentBoundShader.attribNames[i])
		attribLoc := webgl.Call("getAttribLocation", currentBoundShader.shader,
			currentBoundShader.attribNames[i])
		fmt.Println(currentBoundShader.attribNames[i], attribLoc)

		webgl.Call("enableVertexAttribArray", attribLoc)
		webgl.Call("vertexAttribPointer", attribLoc, offset, FLOAT, false,
			4*sumOffsets, offsetAmount*4)
		offsetAmount += offset
	}
	//gl.enableVertexAttribArray(positionAttributeLocation);
	//var size = 2;          // 2 components per iteration
	//var type = gl.FLOAT;   // the data is 32bit floats
	//var normalize = false; // use the data as is
	//var stride = 0;        // 0 = move size * sizeof(type) each iteration
	//var offset = 0;        // start at the beginning of the buffer
	//gl.vertexAttribPointer(
	//	positionAttributeLocation, size, type, normalize, stride, offset)

	return &VAO{PrimitiveType: primitiveType,
		VertexAmount: int32(len(verts)) / sumOffsets}
}

func (v *VAO) Draw() {
	webgl.Call("drawArrays", v.PrimitiveType, 0, v.VertexAmount)
}
