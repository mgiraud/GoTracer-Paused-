package raytracer

import (
	"math"

	"github.com/mgiraud/raytracer/matrix"
)

type Plane struct {
	Position      matrix.Vector4 `json:"position"`
	Rotation      matrix.Vector3 `json:"rotation"`
	Color         matrix.Vector4 `json:"color"`
	Albedo        float64        `json:"albedo"`
	Kd            float64        `json:"Kd"`
	Ks            float64        `json:"Ks"`
	Ior           float64        `json:"ior"`
	Phong         float64        `json:"phong"`
	ObjectToWorld matrix.Matrix16
	WorldToObject matrix.Matrix16
}

func (plane *Plane) GetKs() float64 {
	return plane.Ks
}

func (plane *Plane) GetKd() float64 {
	return plane.Kd
}

func (plane *Plane) GetIor() float64 {
	return plane.Ior
}

func (plane *Plane) GetPhong() float64 {
	return plane.Phong
}

func (plane *Plane) GetColor() matrix.Vector4 {
	return plane.Color
}

func (plane *Plane) GetAlbedo() float64 {
	return plane.Albedo
}

func (plane *Plane) ComputeMatrix() {
	plane.ObjectToWorld = matrix.NewMat16()
	plane.ObjectToWorld.Mat16Identity()
	plane.ObjectToWorld.Translate(plane.ObjectToWorld, plane.Position)
	if plane.Rotation != nil {
		if plane.Rotation[0] != 0 {
			plane.ObjectToWorld.Rotate(plane.ObjectToWorld, matrix.Vector3{1, 0, 0}, plane.Rotation[0]*math.Pi/180)
		}
		if plane.Rotation[1] != 0 {
			plane.ObjectToWorld.Rotate(plane.ObjectToWorld, matrix.Vector3{0, 1, 0}, plane.Rotation[1]*math.Pi/180)
		}
		if plane.Rotation[2] != 0 {
			plane.ObjectToWorld.Rotate(plane.ObjectToWorld, matrix.Vector3{0, 0, 1}, plane.Rotation[2]*math.Pi/180)
		}
	}
	plane.WorldToObject = matrix.NewMat16()
	plane.WorldToObject.Inverse(plane.ObjectToWorld)
}

func (plane *Plane) Intersect(ray *Ray) (bool, float64, matrix.Vector4) {
	var inter matrix.Vector4 = nil
	var res bool = false
	var x0 float64
	ori := matrix.Vector4{0, 0, 0, 1}
	ori.Mat16MulVec4(plane.WorldToObject, ray.Ori)
	dir := matrix.Vector4{0, 0, 0, 0}
	dir.Mat16MulVec4(plane.WorldToObject, ray.Dir)

	if dir[1] != 0 {
		t := -ori[1] / dir[1]
		if t > epsilon {
			res = true
			inter = matrix.Vector4{
				ori[0] + dir[0]*t,
				0.0,
				ori[2] + dir[2]*t,
				1.0,
			}
			x0 = ori.Distance(inter)
			inter.Mat16MulVec4(plane.ObjectToWorld, inter)
			// fmt.Println(inter)
		}
	} else {
		inter = nil
	}
	return res, x0, inter
}

func (plane *Plane) Normale(inter matrix.Vector4, ray *Ray) matrix.Vector4 {
	// normale := matrix.NewVec4()
	// normale.Mat16MulVec4(plane.WorldToObject, inter)
	normale := matrix.Vector4{0.0, 1.0, 0.0, 0.0}
	plane.CorrigerNormale(normale, inter, ray)
	return normale
}

func (plane *Plane) CorrigerNormale(n, i matrix.Vector4, ray *Ray) {
	dir := matrix.Vector4{0, 0, 0, 0}
	dir.Mat16MulVec4(plane.WorldToObject, ray.Dir)
	ps := n.Dot(dir)
	if ps > 0.0 {
		n[0] = -n[0]
		n[1] = -n[1]
		n[2] = -n[2]
		n[3] = 0
	}
	n.Mat16MulVec4(plane.ObjectToWorld, n)
	n.Normalize()
}
