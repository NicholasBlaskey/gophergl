package gl

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
		-0.5, -0.5, -0.5,
		+0.5, -0.5, -0.5,
		+0.5, +0.5, -0.5,
		+0.5, +0.5, -0.5,
		-0.5, +0.5, -0.5,
		-0.5, -0.5, -0.5,

		-0.5, -0.5, +0.5,
		+0.5, -0.5, +0.5,
		+0.5, +0.5, +0.5,
		+0.5, +0.5, +0.5,
		-0.5, +0.5, +0.5,
		-0.5, -0.5, +0.5,

		-0.5, +0.5, +0.5,
		-0.5, +0.5, -0.5,
		-0.5, -0.5, -0.5,
		-0.5, -0.5, -0.5,
		-0.5, -0.5, +0.5,
		-0.5, +0.5, +0.5,

		+0.5, +0.5, +0.5,
		+0.5, +0.5, -0.5,
		+0.5, -0.5, -0.5,
		+0.5, -0.5, -0.5,
		+0.5, -0.5, +0.5,
		+0.5, +0.5, +0.5,

		-0.5, -0.5, -0.5,
		+0.5, -0.5, -0.5,
		+0.5, -0.5, +0.5,
		+0.5, -0.5, +0.5,
		-0.5, -0.5, +0.5,
		-0.5, -0.5, -0.5,

		-0.5, +0.5, -0.5,
		+0.5, +0.5, -0.5,
		+0.5, +0.5, +0.5,
		+0.5, +0.5, +0.5,
		-0.5, +0.5, +0.5,
		-0.5, +0.5, -0.5,
	}

	numVerts := len(position) / 3

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

		if lenTex != 0 {
			outVerts = append(outVerts, texCoords[i*lenTex:i*lenTex+lenTex]...)
		}
	}

	return TRIANGLES, offsets, outVerts
}
