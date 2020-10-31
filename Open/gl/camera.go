// Used this resource
// https://andreasrohner.at/posts/Web%20Development/JavaScript/Simple-orbital-camera-controls-for-THREE-js/
package gl

import (
	mgl "github.com/go-gl/mathgl/mgl32"
	"math"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type OrbitalCamera struct {
	center      mgl.Vec3
	position    mgl.Vec3
	startDragX  float64
	startDragY  float64
	mouseX      float64
	mouseY      float64
	dragging    bool
	ScaleFactor float32
	Sensativity float32
}

var camera *OrbitalCamera = nil

func NewOrbitalCamera(w *Window, radius float32, center mgl.Vec3) *OrbitalCamera {
	camera = &OrbitalCamera{center: center,
		position:    mgl.Vec3{radius, 0.0, 0.0},
		ScaleFactor: 0.1, Sensativity: math.Pi / 450.0}
	w.window.SetCursorPosCallback(glfw.CursorPosCallback(mouseCallback))
	w.window.SetMouseButtonCallback(mouseButtonCallback)
	w.window.SetScrollCallback(scrollCallback)
	return camera
}

func (c *OrbitalCamera) drag(deltaX, deltaY float64) {
	radius, theta, phi := mgl.CartesianToSpherical(c.position)

	phi -= float32(float64(c.Sensativity) * deltaX)
	theta = float32(math.Min(math.Max(
		float64(theta)+float64(c.Sensativity)*deltaY, 0.001), math.Pi-0.001))

	c.position = mgl.SphericalToCartesian(radius, theta, phi)
}

func (c *OrbitalCamera) LookAt() mgl.Mat4 {
	return mgl.LookAtV(c.position.Add(c.center), c.center, mgl.Vec3{0.0, 0.0, 1.0})
}

func (c *OrbitalCamera) zoom(zoomIn bool) {
	if zoomIn {
		c.position = c.position.Sub(c.center).Mul(1.0 - c.ScaleFactor).Add(c.center)
	} else {
		c.position = c.position.Sub(c.center).Mul(1.0 + c.ScaleFactor).Add(c.center)
	}
}

func mouseCallback(w *glfw.Window, xPos float64, yPos float64) {
	camera.mouseX, camera.mouseY = xPos, yPos

	if camera.dragging {
		camera.drag(camera.startDragX-xPos, camera.startDragY-yPos)
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

func scrollCallback(w *glfw.Window, xOffset float64, yOffset float64) {
	if yOffset == 0 {
		return
	}
	camera.zoom(yOffset > 0)
}
