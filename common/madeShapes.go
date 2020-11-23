package common

import (
	mgl "github.com/go-gl/mathgl/mgl32"
)

const (
	TRIANGLES = 0x0004
)

type VertParams struct {
	Position  bool
	TexCoords bool
	Normals   bool
}

func NewCube(p VertParams) (uint32, []int32, []float32) {
	offsets := []int32{}

	lenPos := 0
	if p.Position {
		lenPos = 3
		offsets = append(offsets, int32(lenPos))
	}
	position := []float32{
		-1.0, -1.0, -1.0,
		+1.0, -1.0, -1.0,
		+1.0, +1.0, -1.0,
		+1.0, +1.0, -1.0,
		-1.0, +1.0, -1.0,
		-1.0, -1.0, -1.0,

		-1.0, -1.0, +1.0,
		+1.0, -1.0, +1.0,
		+1.0, +1.0, +1.0,
		+1.0, +1.0, +1.0,
		-1.0, +1.0, +1.0,
		-1.0, -1.0, +1.0,

		-1.0, +1.0, +1.0,
		-1.0, +1.0, -1.0,
		-1.0, -1.0, -1.0,
		-1.0, -1.0, -1.0,
		-1.0, -1.0, +1.0,
		-1.0, +1.0, +1.0,

		+1.0, +1.0, +1.0,
		+1.0, +1.0, -1.0,
		+1.0, -1.0, -1.0,
		+1.0, -1.0, -1.0,
		+1.0, -1.0, +1.0,
		+1.0, +1.0, +1.0,

		-1.0, -1.0, -1.0,
		+1.0, -1.0, -1.0,
		+1.0, -1.0, +1.0,
		+1.0, -1.0, +1.0,
		-1.0, -1.0, +1.0,
		-1.0, -1.0, -1.0,

		-1.0, +1.0, -1.0,
		+1.0, +1.0, -1.0,
		+1.0, +1.0, +1.0,
		+1.0, +1.0, +1.0,
		-1.0, +1.0, +1.0,
		-1.0, +1.0, -1.0,
	}

	numVerts := len(position) / 3

	lenNormal := 0
	if p.Normals {
		lenNormal = 3
		offsets = append(offsets, int32(lenNormal))
	}
	normals := []float32{
		+0.0, +0.0, -1.0,
		+0.0, +0.0, -1.0,
		+0.0, +0.0, -1.0,
		+0.0, +0.0, -1.0,
		+0.0, +0.0, -1.0,
		+0.0, +0.0, -1.0,

		+0.0, +0.0, +1.0,
		+0.0, +0.0, +1.0,
		+0.0, +0.0, +1.0,
		+0.0, +0.0, +1.0,
		+0.0, +0.0, +1.0,
		+0.0, +0.0, +1.0,

		-1.0, +0.0, +0.0,
		-1.0, +0.0, +0.0,
		-1.0, +0.0, +0.0,
		-1.0, +0.0, +0.0,
		-1.0, +0.0, +0.0,
		-1.0, +0.0, +0.0,

		+1.0, +0.0, +0.0,
		+1.0, +0.0, +0.0,
		+1.0, +0.0, +0.0,
		+1.0, +0.0, +0.0,
		+1.0, +0.0, +0.0,
		+1.0, +0.0, +0.0,

		+0.0, -1.0, +0.0,
		+0.0, -1.0, +0.0,
		+0.0, -1.0, +0.0,
		+0.0, -1.0, +0.0,
		+0.0, -1.0, +0.0,
		+0.0, -1.0, +0.0,

		+0.0, +1.0, +0.0,
		+0.0, +1.0, +0.0,
		+0.0, +1.0, +0.0,
		+0.0, +1.0, +0.0,
		+0.0, +1.0, +0.0,
		+0.0, +1.0, +0.0,
	}

	lenTex := 0
	if p.TexCoords {
		lenTex = 2
		offsets = append(offsets, int32(lenTex))
	}
	texCoords := []float32{
		0.0, 0.0,
		1.0, 0.0,
		1.0, 1.0,
		1.0, 1.0,
		0.0, 1.0,
		0.0, 0.0,

		0.0, 0.0,
		1.0, 0.0,
		1.0, 1.0,
		1.0, 1.0,
		0.0, 1.0,
		0.0, 0.0,

		1.0, 0.0,
		1.0, 1.0,
		0.0, 1.0,
		0.0, 1.0,
		0.0, 0.0,
		1.0, 0.0,

		1.0, 0.0,
		1.0, 1.0,
		0.0, 1.0,
		0.0, 1.0,
		0.0, 0.0,
		1.0, 0.0,

		0.0, 1.0,
		1.0, 1.0,
		1.0, 0.0,
		1.0, 0.0,
		0.0, 0.0,
		0.0, 1.0,

		0.0, 1.0,
		1.0, 1.0,
		1.0, 0.0,
		1.0, 0.0,
		0.0, 0.0,
		0.0, 1.0,
	}

	outVerts := []float32{}
	for i := 0; i < numVerts; i++ {
		if lenPos != 0 {
			outVerts = append(outVerts, position[i*lenPos:i*lenPos+lenPos]...)
		}

		if lenNormal != 0 {
			outVerts = append(outVerts, normals[i*lenNormal:i*lenNormal+lenNormal]...)
		}

		if lenTex != 0 {
			outVerts = append(outVerts, texCoords[i*lenTex:i*lenTex+lenTex]...)
		}
	}

	return TRIANGLES, offsets, outVerts
}

func NewSphere(p VertParams) (uint32, []int32, []int32) {
	xSegments := 64
	ySegments := 64
	pi := float32(math.Pi)
	for y := 0; y <= ySegments; y++ {
		for x := 0; x <= xSegments; x++ {
			xSegment := float32(x) / float32(xSegments)
			ySegment := float32(y) / float32(ySegments)
			xPos := float32(math.Cos(float64(xSegment*2.0*pi)) *
				math.Sin(float64(ySegment*pi)))
			yPos := float32(math.Cos(float64(ySegment * pi)))
			zPos := float32(math.Sin(float64(xSegment*2.0*pi)) *
				math.Sin(float64(ySegment*pi)))

			positions = append(positions, mgl.Vec3{xPos, yPos, zPos})
			uv = append(uv, mgl.Vec2{xSegment, ySegment})
			normals = append(normals, mgl.Vec3{xPos, yPos, zPos})
		}
	}

	oddRow := false
	for y := 0; y < ySegments; y++ {
		if oddRow {
			indices = append(indices, uint32(y*(xSegments+1)+x))
			indices = append(indices, uint32((y+1)*(xSegments+1)+x))
		} else {
			indices = append(indices, uint32((y+1)*(xSegments+1)+x))
			indices = append(indices, uint32(y*(xSegments+1)+x))
		}
		oddRow = !oddRow
	}
	indexCount = uint32(len(indices))

	data := []float32{}
	for i := 0; i < len(positions); i++ {
		data = append(data, positions[i][0], positions[i][1], positions[i][2])
		if len(uv) > 0 {
			data = append(data, uv[i][0], uv[i][1])
		}
		if len(normals) > 0 {
			data = append(data, normals[i][0], normals[i][1], normals[i][2])
		}
	}
}
