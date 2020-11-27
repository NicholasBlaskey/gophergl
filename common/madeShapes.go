package common

import (
	mgl "github.com/go-gl/mathgl/mgl32"
	"math"
)

const (
	TRIANGLES      = 0x0004
	TRIANGLE_STRIP = 0x0005
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

// TODO rework premade shapes to use an ebo
func NewSphere(p VertParams) (uint32, []int32, []float32) {
	positions := []mgl.Vec3{}
	uv := []mgl.Vec2{}
	normals := []mgl.Vec3{}

	offsets := []int32{}
	if p.Position {
		offsets = append(offsets, 3)
	}
	if p.TexCoords {
		offsets = append(offsets, 2)
	}
	if p.Normals {
		offsets = append(offsets, 3)
	}

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

			if p.Position {
				positions = append(positions, mgl.Vec3{xPos, yPos, zPos})
			}
			if p.TexCoords {
				uv = append(uv, mgl.Vec2{xSegment, ySegment})
			}
			if p.Normals {
				normals = append(normals, mgl.Vec3{xPos, yPos, zPos})
			}
		}
	}

	data := []float32{}
	oddRow := false
	for y := 0; y < ySegments; y++ {
		if oddRow {
			for x := 0; x <= xSegments; x++ {
				indexes := []int{y*(xSegments+1) + x, (y+1)*(xSegments+1) + x}
				for _, i := range indexes {
					appendDataSphere(&data, positions, uv, normals, p, i)
				}
			}
		} else {
			for x := xSegments; x >= 0; x-- {
				indexes := []int{(y+1)*(xSegments+1) + x, y*(xSegments+1) + x}
				for _, i := range indexes {
					appendDataSphere(&data, positions, uv, normals, p, i)
				}
			}
		}
		oddRow = !oddRow
	}

	return TRIANGLE_STRIP, offsets, data
}

func appendDataSphere(data *[]float32, positions []mgl.Vec3,
	uv []mgl.Vec2, normals []mgl.Vec3, p VertParams, i int) {

	if p.Position {
		*data = append(*data, positions[i][0], positions[i][1],
			positions[i][2])
	}
	if p.TexCoords {
		*data = append(*data, uv[i][0], uv[i][1])
	}
	if p.Normals {
		*data = append(*data, normals[i][0], normals[i][1], normals[i][2])
	}
}
