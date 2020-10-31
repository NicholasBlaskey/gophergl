// Used this tutorial
// https://andreasrohner.at/posts/Web%20Development/JavaScript/Simple-orbital-camera-controls-for-THREE-js/
package gl

import (
	mgl "github.com/go-gl/mathgl/mgl32"
	"math"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type OrbitalCamera struct {
	center     mgl.Vec3
	position   mgl.Vec3
	startDragX float64
	startDragY float64
	mouseX     float64
	mouseY     float64
	dragging   bool
}

var camera *OrbitalCamera = nil

func NewOrbitalCamera(w *Window, radius float32, center mgl.Vec3) *OrbitalCamera {
	camera = &OrbitalCamera{center, mgl.Vec3{radius, 0.0, 0.0},
		-1, -1, -1, -1, false}
	w.window.SetCursorPosCallback(glfw.CursorPosCallback(mouseCallback))
	w.window.SetMouseButtonCallback(mouseButtonCallback)

	return camera
}

func mouseCallback(w *glfw.Window, xPos float64, yPos float64) {
	camera.mouseX, camera.mouseY = xPos, yPos

	if camera.dragging {
		camera.drag(camera.startDragX-xPos, camera.startDragY-yPos)
	} else {

	}
	camera.startDragX = camera.mouseX
	camera.startDragY = camera.mouseY
}

func mouseButtonCallback(window *glfw.Window, button glfw.MouseButton,
	action glfw.Action, mods glfw.ModifierKey) {

	if button == glfw.MouseButtonLeft && action == glfw.Press {
		camera.startDragX = camera.mouseX
		camera.startDragY = camera.mouseY
		camera.dragging = true
	} else if button == glfw.MouseButtonLeft && action == glfw.Release {
		camera.dragging = false
	}
}

func (c *OrbitalCamera) drag(deltaX, deltaY float64) {
	radPerPixel := math.Pi / 450.0 //450.0

	// Convert to sphericial coords
	radius := float64(c.position.Len())
	theta := math.Acos(float64(c.position[2]) / radius)
	phi := math.Atan2(float64(c.position[1]), float64(c.position[0]))

	phi -= radPerPixel * deltaX
	theta = math.Min(math.Max(theta+radPerPixel*-deltaY, 0.001), math.Pi-0.001)

	// Convert back to cartesion
	c.position[0] = float32(radius * math.Sin(theta) * math.Cos(phi))
	c.position[1] = float32(radius * math.Sin(theta) * math.Sin(phi))
	c.position[2] = float32(radius * math.Cos(theta))

	c.position = c.position
}

func (c *OrbitalCamera) LookAt() mgl.Mat4 {
	return mgl.LookAtV(c.position.Add(c.center), c.center, mgl.Vec3{0.0, 0.0, 1.0})
}
