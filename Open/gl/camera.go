// Used this tutorial
// https://andreasrohner.at/posts/Web%20Development/JavaScript/Simple-orbital-camera-controls-for-THREE-js/
package gl

import (
	"fmt"

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

func NewOrbitalCamera(w *Window) *OrbitalCamera {
	camera = &OrbitalCamera{mgl.Vec3{0.0, 0.0, 0.0}, mgl.Vec3{0.0, 0.0, 5.0},
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
		return
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
	//fmt.Println("DRAGGING\n\n\n", deltaX, deltaY)

	radPerPixel := math.Pi / 10.0 //450.0
	//deltaPhi := radPerPixel * deltaX
	//deltaTheta := radPerPixel * deltaY
	//deltaPhi := radPerPixel * -deltaY
	//deltaTheta := radPerPixel * deltaX
	//pos := c.position.Sub(c.center)
	//radius := float64(pos.Len())
	//theta := math.Acos(float64(pos[2]) / radius)
	//phi := math.Atan2(float64(pos[1]), float64(pos[0]))
	//theta = theta - deltaTheta //math.Min(math.Max(theta-deltaTheta, 0), math.Pi)
	//phi -= deltaPhi
	//fmt.Println("RADIUS THETA PHI", radius, theta, phi)

	fmt.Println("deltaX START POS", deltaX, c.position, c.center)

	// Convert to sphericial coords
	radius := float64(c.position.Len())
	theta := math.Acos(float64(c.position[2]) / radius)
	phi := math.Atan2(float64(c.position[1]), float64(c.position[0]))
	fmt.Println("radius theta phi", radius, theta, phi)

	//theta += radPerPixel * deltaX
	fmt.Println("Start PHI", phi)
	phi -= radPerPixel * deltaX
	fmt.Println("POST PHI", phi)

	// Convert back to cartesion
	c.position[0] = float32(radius * math.Sin(theta) * math.Cos(phi))
	c.position[1] = float32(radius * math.Sin(theta) * math.Sin(phi))
	c.position[2] = float32(radius * math.Cos(theta))
	c.position = c.position.Add(c.center)
	fmt.Println("End pos", c.position, c.center)

	//panic("over")
}

func (c *OrbitalCamera) LookAt() mgl.Mat4 {
	//fmt.Println(mgl.LookAtV(c.position, c.center, mgl.Vec3{0.0, 0.0, 1.0}))
	//return mgl.LookAtV(c.position, c.center, mgl.Vec3{0.0, 1.0, 0.0})
	return mgl.LookAtV(c.position, c.center, mgl.Vec3{0.0, 1.0, 0.0})
}
