package raytracer

import "github.com/mgiraud/raytracer/matrix"

type Object interface {
	Intersect(ray *Ray) (bool, float64, matrix.Vector4)
	Normale(inter matrix.Vector4, ray *Ray) matrix.Vector4
	ComputeMatrix()
	GetColor() matrix.Vector4
	GetAlbedo() float64
	GetKs() float64
	GetKd() float64
	GetN() float64
}

type Sphere struct {
	Radius        float64        `json:"radius"`
	Position      matrix.Vector4 `json:"position"`
	Rotation      matrix.Vector3 `json:"rotation"`
	Color         matrix.Vector4 `json:"color"`
	Albedo        float64        `json:"albedo"`
	Kd            float64        `json:"Kd"`
	Ks            float64        `json:"Ks"`
	N             float64        `json:"n"`
	ObjectToWorld matrix.Matrix16
	WorldToObject matrix.Matrix16
}

func (sph *Sphere) GetKs() float64 {
	return sph.Ks
}

func (sph *Sphere) GetKd() float64 {
	return sph.Kd
}

func (sph *Sphere) GetN() float64 {
	return sph.N
}

func (sph *Sphere) GetColor() matrix.Vector4 {
	return sph.Color
}

func (sph *Sphere) GetAlbedo() float64 {
	return sph.Albedo
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

func (sph *Sphere) Intersect(ray *Ray) (bool, float64, matrix.Vector4) {
	var inter matrix.Vector4
	ori := matrix.Vector4{0, 0, 0, 1}
	ori.Mat16MulVec4(sph.WorldToObject, ray.Ori)
	dir := matrix.Vector4{0, 0, 0, 0}
	dir.Mat16MulVec4(sph.WorldToObject, ray.Dir)

	a := dir.Dot(dir)
	b := 2 * dir.Dot(ori)
	c := ori.Dot(ori) - sph.Radius*sph.Radius
	if ray.Type == "SHADOW_DIRECT_LIGHT" {
		c += 4.0
	}
	res, x0 := solveQuadratic(a, b, c)
	if res {
		// fmt.Println(dir)
		inter = matrix.Vector4{
			x0*dir[0] + ori[0]/ori[3],
			x0*dir[1] + ori[1]/ori[3],
			x0*dir[2] + ori[2]/ori[3],
			1.0,
		}
		x0 = ori.Distance(inter)
		inter.Mat16MulVec4(sph.ObjectToWorld, inter)
	} else {
		inter = nil
	}
	if res && ray.Type == "SHADOW_DIRECT_LIGHT" {
		// fmt.Println(a, b, c)
		// fmt.Println(ori)
		// fmt.Println(dir)
		// fmt.Println(inter, x0)
	} else {
		// fmt.Println(dir)
	}

	return res, x0, inter
}

func (sph *Sphere) Normale(inter matrix.Vector4, ray *Ray) matrix.Vector4 {
	normale := matrix.NewVec4()
	normale.Mat16MulVec4(sph.WorldToObject, inter)
	normale = matrix.Vector4{
		normale[0] / normale[3],
		normale[1] / normale[3],
		normale[2] / normale[3],
		0.0,
	}
	normale.Mat16MulVec4(sph.ObjectToWorld, normale)
	normale.Normalize()
	return normale
}
