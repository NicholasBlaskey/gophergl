package gl

import (
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	startDragX = float32(-1)
	startDragY = float32(-1)
	mouseX     = float32(-1)
	mouseY     = float32(-1)
)

func (w *Window) OrbitialCamera() {
	fmt.Println("GOING ON")
	w.window.SetCursorPosCallback(glfw.CursorPosCallback(mouseCallback))
	w.window.SetMouseButtonCallback(mouseButtonCallback)
}

func mouseCallback(w *glfw.Window, xPos float64, yPos float64) {
	mouseX, mouseY = float32(xPos), float32(yPos)
	//if firstMouse {
	//lastX = float32(xPos)
	//lastY = float32(yPos)
	//firstMouse = false
	//}
	//xOffset := float32(xPos) - lastX
	// Reversed due to y coords go from bot up
	//yOffset := lastY - float32(yPos)

	//lastX = float32(xPos)
	//lastY = float32(yPos)

	//ourCamera.ProcessMouseMovement(xOffset, yOffset, true)
}

func mouseButtonCallback(window *glfw.Window, button glfw.MouseButton,
	action glfw.Action, mods glfw.ModifierKey) {

	if button == glfw.MouseButtonLeft && action == glfw.Press {
		startDragX, startDragY = mouseX, mouseY
	} else if button == glfw.MouseButtonLeft && action == glfw.Release {
		fmt.Println(startDragX, startDragY)
	}
}
