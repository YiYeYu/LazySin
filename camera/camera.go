package camera

import (
	"github.com/go-gl/gl/v2.1/gl"
)

//Camera is camera
type Camera struct {
}

const (
	projectMatrixNear2D = -100
	projectMatrixFar2D  = 100
)

//PrepareViewPort prepare viewport
func (c *Camera) PrepareViewPort(x, y int32, width, height int32) {
	gl.Viewport(x, y, width, height)
}

//SetCamera2D set a 2D camera
func (c *Camera) SetCamera2D(left float64, right float64, bottom float64, top float64) {
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(left, right, bottom, top, projectMatrixNear2D, projectMatrixFar2D)
}

//MoveCamera move camera position
func (c *Camera) MoveCamera(x float64, y float64, z float64) {
	gl.MatrixMode(gl.PROJECTION)
	gl.Translated(x, y, z)
}
