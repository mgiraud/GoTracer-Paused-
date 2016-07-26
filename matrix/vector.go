package matrix

type Vector4 []float64
type Vector3 []float64

func NewVec4() Vector4 {
	return Vector4{0, 0, 0, 0}
}

func NewVec3() Vector3 {
	return Vector3{0, 0, 0}
}

func (dest Vector4) Mat16MulVec4(mat Matrix16, vec Vector4) {
	x := vec[0]
	y := vec[1]
	z := vec[2]
	w := vec[3]

	dest[0] = mat[0]*x + mat[4]*y + mat[8]*z + mat[12]*w
	dest[1] = mat[1]*x + mat[5]*y + mat[9]*z + mat[13]*w
	dest[2] = mat[2]*x + mat[6]*y + mat[10]*z + mat[14]*w
	dest[3] = mat[3]*x + mat[7]*y + mat[11]*z + mat[15]*w
}
