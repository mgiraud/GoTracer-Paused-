package raytracer

import "github.com/mgiraud/raytracer/matrix"

type Ray struct {
	I, J int
	Ori  matrix.Vector4
	Dir  matrix.Vector4
}

func NewRay(i int, j int, ori matrix.Vector4, dir matrix.Vector4) Ray {
	dir.Normalize()
	return Ray{
		I:   i,
		J:   j,
		Ori: ori,
		Dir: dir,
	}
}
