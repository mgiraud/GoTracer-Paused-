package raytracer

import (
	"math"

	"github.com/mgiraud/raytracer/matrix"
)

type Ray struct {
	Type  string
	Ori   matrix.Vector4
	Dir   matrix.Vector4
	Depth int
}

func NewRay(ori matrix.Vector4, dir matrix.Vector4, name string, depth int) *Ray {
	return &Ray{
		Ori:   ori,
		Dir:   dir,
		Type:  name,
		Depth: depth,
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

func Fresnel(i, n matrix.Vector4, ior float64) float64 {
	var cosi, etai, etat, sint, cost, kr, Rs, Rp float64
	cosi = math.Min(1, math.Max(-1, i.Dot(n)))
	if cosi > 0 {
		etat, etai = 1, ior
	} else {
		etat, etai = ior, 1
	}
	sint = etai / etat * math.Sqrt(math.Max(0.0, 1-cosi*cosi))
	if sint >= 1 {
		kr = 1
	} else {
		cost = math.Sqrt(math.Max(0.0, 1-sint*sint))
		cosi = math.Abs(cosi)
		Rs = ((etat * cosi) - (etai * cost)) / ((etat * cosi) + (etai * cost))
		Rp = ((etai * cosi) - (etat * cost)) / ((etai * cosi) + (etat * cost))
		kr = (Rs*Rs + Rp*Rp) / 2
	}
	return kr
}

func Refract(i, N matrix.Vector4, ior float64) matrix.Vector4 {
	var cosi, etai, eta, etat, k float64
	var n matrix.Vector4 = N.Clone()
	cosi = math.Min(1, math.Max(-1, i.Dot(n)))
	etai, etat = 1, ior
	if cosi < 0 {
		cosi = -cosi
	} else {
		etai = etat
		etat = 1
		n = n.Neg()
	}
	eta = etai / etat
	k = 1 - eta*eta*(1-cosi*cosi)
	if k < 0 {
		return matrix.Vector4{0, 0, 0, 0}
	} else {
		n = n.MulFloat(eta*cosi - math.Sqrt(k))
		return i.MulFloat(eta).Add(n)
	}
}
