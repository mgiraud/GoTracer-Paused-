package raytracer

import "math"
import "github.com/mgiraud/raytracer/matrix"

type Camera struct {
	fov     float64
	ratio   float64
	bottom  float64
	top     float64
	right   float64
	left    float64
	near    float64
	far     float64
	ProjMat matrix.Matrix16
}

func (cam *Camera) SetPerspective(fov, near, far, ratio float64) {
	scale := math.Tan(fov*0.5*math.Pi/180) * near
	cam.fov = fov
	cam.ratio = ratio
	cam.right = ratio * scale
	cam.left = -cam.right
	cam.top = scale
	cam.bottom = -cam.top
	cam.near = near
	cam.far = far
}

func (cam *Camera) SetProjectionMatrix() {
	cam.ProjMat = matrix.NewMat16()
	cam.ProjMat[0] = 2 * cam.near / (cam.right - cam.left)
	cam.ProjMat[1] = 0
	cam.ProjMat[2] = 0
	cam.ProjMat[3] = 0

	cam.ProjMat[4] = 0
	cam.ProjMat[5] = 2 * cam.near / (cam.top - cam.bottom)
	cam.ProjMat[6] = 0
	cam.ProjMat[7] = 0

	cam.ProjMat[8] = (cam.right + cam.left) / (cam.right - cam.left)
	cam.ProjMat[9] = (cam.top + cam.bottom) / (cam.top - cam.bottom)
	cam.ProjMat[10] = -(cam.far + cam.near) / (cam.far - cam.near)
	cam.ProjMat[11] = -1

	cam.ProjMat[12] = 0
	cam.ProjMat[13] = 0
	cam.ProjMat[14] = -2 * cam.far * cam.near / (cam.far - cam.near)
	cam.ProjMat[15] = 0
}
