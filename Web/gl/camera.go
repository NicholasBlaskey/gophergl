// Used this resource
// https://andreasrohner.at/posts/Web%20Development/JavaScript/Simple-orbital-camera-controls-for-THREE-js/

package gl

import (
	"math"

	mgl "github.com/go-gl/mathgl/mgl32"
	"github.com/gopherjs/gopherjs/js"
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

	w.window.Call("addEventListener", "mousedown", mouseDown)
	w.window.Call("addEventListener", "mouseup", mouseUp)
	w.window.Call("addEventListener", "mousemove", mouseMove)
	w.window.Call("addEventListener", "mousewheel", mouseWheel)
	w.window.Call("addEventListener", "DOMMouseScroll", mouseWheel)

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

func mouseDown(event *js.Object) {
	camera.dragging = true
	camera.startDragX = event.Get("clientX").Float()
	camera.startDragY = event.Get("clientY").Float()

	event.Call("preventDefault")
}

func mouseUp(event *js.Object) {
	camera.dragging = false

	event.Call("preventDefault")
}

func mouseMove(event *js.Object) {
	xPos := event.Get("clientX").Float()
	yPos := event.Get("clientY").Float()
	if camera.dragging {
		camera.drag(camera.startDragX-xPos, camera.startDragY-yPos)
	}

	camera.startDragX = xPos
	camera.startDragY = yPos

	event.Call("preventDefault")
}

func mouseWheel(event *js.Object) {
	event.Call("preventDefault")

	scrollDelta := 0.0
	detail := event.Get("detail")
	if detail != nil {
		scrollDelta = detail.Float()
	} else {
		scrollDelta = event.Get("wheelDelta").Float()
	}
	if scrollDelta == 0 {
		return
	}
	camera.zoom(scrollDelta < 0.0)
}
