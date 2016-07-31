package raytracer

import "math"
import "github.com/mgiraud/raytracer/matrix"

type Camera struct {
	Fov            float64        `json:"fov"`
	Position       matrix.Vector4 `json:"position"`
	Rotation       matrix.Vector3 `json:"rotation"`
	Width          float64        `json:"width"`
	Height         float64        `json:"height"`
	Bias           float64        `json:"bias"`
	MaxDepth       int            `json:"maxDepth"`
	BacgroundColor matrix.Vector4 `json:"color"`
	// far, bottom, top, right, left, near float64
	ratio, scale                    float64
	ProjMat, CamToWorld, WorldToCam matrix.Matrix16
	RayOrigin                       matrix.Vector4
}

func (cam *Camera) InitCamera() {
	mat := matrix.NewMat16()
	mat.Mat16Identity()
	invmat := mat.Clone()
	invmat.Inverse(invmat)
	origin := matrix.Vector4{0, 0, 0, 1}
	origin.Mat16MulVec4(mat, origin)
	cam.ratio = (cam.Width / cam.Height)
	cam.scale = math.Tan(cam.Fov * 0.5 * math.Pi / 180)
	cam.CamToWorld = mat
	cam.WorldToCam = invmat
	cam.RayOrigin = origin
	cam.MoveCamera(cam.Position)
	// if cam.Rotation != nil {
	// 	if cam.Rotation[0] != 0 {
	// 		cam.RotateCamera(matrix.Vector3{1, 0, 0}, cam.Rotation[0])
	// 	}
	// 	if cam.Rotation[1] != 0 {
	// 		cam.RotateCamera(matrix.Vector3{0, 1, 0}, cam.Rotation[1])
	// 	}
	// 	if cam.Rotation[2] != 0 {
	// 		cam.RotateCamera(matrix.Vector3{0, 0, 1}, cam.Rotation[2])
	// 	}
	// }
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

/*  Not used for rayTracer -> this is just an example of projection matrix used by OpenGl for example
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
*/
