package gl

import (
	"github.com/nicholasblaskey/gophergl/common"
)

type VertParams common.VertParams

func NewCube(p VertParams) (uint32, []int32, []float32) {
	return common.NewCube(common.VertParams(p))
}
