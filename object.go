package raytracer

import "github.com/mgiraud/raytracer/matrix"

const epsilon = 0.001

type Object interface {
	Intersect(ray *Ray) (bool, float64, matrix.Vector4)
	Normale(inter matrix.Vector4, ray *Ray) matrix.Vector4
	ComputeMatrix()
	GetColor() matrix.Vector4
	GetAlbedo() float64
	GetKs() float64
	GetKd() float64
	GetIor() float64
	GetPhong() float64
}
