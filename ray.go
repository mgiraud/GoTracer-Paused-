package raytracer

import "github.com/mgiraud/raytracer/matrix"

type Ray struct {
	Type string
	Ori  matrix.Vector4
	Dir  matrix.Vector4
}

func NewRay(ori matrix.Vector4, dir matrix.Vector4, name string) *Ray {
	return &Ray{
		Ori:  ori,
		Dir:  dir,
		Type: name,
	}
}

func Reflect(i, n matrix.Vector4) matrix.Vector4 {
	factor := i.Dot(n)
	return matrix.Vector4{
		i[0] - 2*factor*n[0],
		i[1] - 2*factor*n[1],
		i[2] - 2*factor*n[2],
		i[3],
	}
}
