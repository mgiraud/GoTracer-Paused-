package raytracer

import (
	"image/color"

	"github.com/mgiraud/raytracer/matrix"
)

type Object interface {
	Intersect(ray *Ray) (bool, float64)
	ComputeMatrix()
	GetColor() color.RGBA
}

type Sphere struct {
	Radius        float64        `json:"radius"`
	Position      matrix.Vector4 `json:"position"`
	Rotation      matrix.Vector3 `json:"rotation"`
	Color         color.RGBA     `json:"color"`
	ObjectToWorld matrix.Matrix16
	WorldToObject matrix.Matrix16
}

func (sph *Sphere) GetColor() color.RGBA {
	return sph.Color
}

func (sph *Sphere) ComputeMatrix() {
	sph.ObjectToWorld = matrix.NewMat16()
	sph.ObjectToWorld.Mat16Identity()
	sph.ObjectToWorld.Translate(sph.ObjectToWorld, sph.Position)
	if sph.Rotation[0] != 0 {
		sph.ObjectToWorld.Rotate(sph.ObjectToWorld, matrix.Vector3{1, 0, 0}, sph.Rotation[0])
	}
	if sph.Rotation[1] != 0 {
		sph.ObjectToWorld.Rotate(sph.ObjectToWorld, matrix.Vector3{0, 1, 0}, sph.Rotation[1])
	}
	if sph.Rotation[2] != 0 {
		sph.ObjectToWorld.Rotate(sph.ObjectToWorld, matrix.Vector3{0, 0, 1}, sph.Rotation[2])
	}
	sph.WorldToObject = matrix.NewMat16()
	sph.WorldToObject.Inverse(sph.ObjectToWorld)
}

func (sph *Sphere) Intersect(ray *Ray) (bool, float64) {
	ori := matrix.Vector4{0, 0, 0, 1}
	ori.Mat16MulVec4(sph.WorldToObject, ray.Ori)
	dir := matrix.Vector4{0, 0, 0, 0}
	dir.Mat16MulVec4(sph.WorldToObject, ray.Dir)

	a := dir.Dot(dir)
	b := 2 * dir.Dot(ori)
	c := ori.Dot(ori) - sph.Radius*sph.Radius
	res, x0 := solveQuadratic(a, b, c)
	return res, x0
}
