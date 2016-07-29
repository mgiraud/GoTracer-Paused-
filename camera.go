package raytracer

import "math"
import "github.com/mgiraud/raytracer/matrix"

type Camera struct {
	Fov        float64        `json:"fov"`
	Position   matrix.Vector4 `json:"position"`
	Rotation   matrix.Vector3 `json:"rotation"`
	ratio      float64
	bottom     float64
	top        float64
	right      float64
	left       float64
	near       float64
	far        float64
	scale      float64
	ProjMat    matrix.Matrix16
	CamToWorld matrix.Matrix16
	WorldToCam matrix.Matrix16
	RayOrigin  matrix.Vector4
}

func (cam *Camera) InitCamera(width, height int, fov float64) {
	mat := matrix.NewMat16()
	mat.Mat16Identity()
	invmat := mat.Clone()
	invmat.Inverse(invmat)
	origin := matrix.Vector4{0, 0, 0, 1}
	origin.Mat16MulVec4(mat, origin)
	cam.ratio = (float64(width) / float64(height))
	cam.scale = math.Tan(fov * 0.5 * math.Pi / 180)
	cam.CamToWorld = mat
	cam.WorldToCam = invmat
	cam.RayOrigin = origin
}

func (cam *Camera) MoveCamera(point matrix.Vector4) {
	cam.CamToWorld.Translate(cam.CamToWorld, point)
	cam.RayOrigin = matrix.Vector4{0, 0, 0, 1}
	cam.RayOrigin.Mat16MulVec4(cam.CamToWorld, cam.RayOrigin)
	cam.WorldToCam = matrix.NewMat16()
	cam.WorldToCam.Inverse(cam.CamToWorld)
}

func (cam *Camera) RotateCamera(axis matrix.Vector3, angle float64) {
	cam.CamToWorld.Rotate(cam.CamToWorld, axis, angle)
	cam.RayOrigin = matrix.Vector4{0, 0, 0, 1}
	cam.RayOrigin.Mat16MulVec4(cam.CamToWorld, cam.RayOrigin)
	cam.WorldToCam = matrix.NewMat16()
	cam.WorldToCam.Inverse(cam.CamToWorld)
}

func (cam *Camera) LookAt(eye matrix.Vector3, center matrix.Vector3, up matrix.Vector3) {
	cam.CamToWorld.LookAt(eye, center, up)
	cam.WorldToCam = matrix.NewMat16()
	cam.WorldToCam.Inverse(cam.CamToWorld)
}

func (cam *Camera) SetPerspective(fov, near, far, ratio float64) {
	scale := math.Tan(fov*0.5*math.Pi/180) * near
	cam.Fov = fov
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
