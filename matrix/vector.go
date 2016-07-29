package matrix

import (
	"fmt"
	"math"
)

type Vector4 []float64
type Vector3 []float64

func NewVec4() Vector4 {
	return Vector4{0, 0, 0, 0}
}

func NewVec3() Vector3 {
	return Vector3{0, 0, 0}
}

func (vec Vector4) String() string {
	return fmt.Sprintf("%f %f %f %f",
		vec[0], vec[1], vec[2], vec[3])
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

func (dest Vector4) Add(vec Vector4) {
	dest[0] += vec[0]
	dest[1] += vec[1]
	dest[2] += vec[2]
	dest[3] += vec[3]
}

func (dest Vector4) Sub(vec Vector4) {
	dest[0] -= vec[0]
	dest[1] -= vec[1]
	dest[2] -= vec[2]
	dest[3] -= vec[3]
}

func (dest Vector4) Mum(vec Vector4) {
	dest[0] *= vec[0]
	dest[1] *= vec[1]
	dest[2] *= vec[2]
	dest[3] *= vec[3]
}

func (dest Vector4) Div(vec Vector4) {
	dest[0] /= vec[0]
	dest[1] /= vec[1]
	dest[2] /= vec[2]
	dest[3] /= vec[3]
}

func (dest Vector4) Scale(i float64) {
	dest[0] *= i
	dest[1] *= i
	dest[2] *= i
	dest[3] *= i
}

func (dest Vector4) Distance(vec Vector4) float64 {
	x := dest[0] - vec[0]
	y := dest[1] - vec[1]
	z := dest[2] - vec[2]
	w := dest[3] - vec[3]
	return math.Sqrt(x*x + y*y + z*z + w*w)
}

func (dest Vector4) Length() float64 {
	x := dest[0]
	y := dest[1]
	z := dest[2]
	w := dest[3]
	return math.Sqrt(x*x + y*y + z*z + w*w)
}

func (dest Vector4) Normalize() {
	x := dest[0]
	y := dest[1]
	z := dest[2]
	w := dest[3]
	len := x*x + y*y + z*z + w*w
	if len > 0 {
		len = 1 / math.Sqrt(len)
		dest[0] = x * len
		dest[1] = y * len
		dest[2] = z * len
		dest[3] = w * len
	}
}

func (a Vector4) Dot(b Vector4) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2] + a[3]*b[3]
}
