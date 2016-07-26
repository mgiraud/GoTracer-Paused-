package matrix

import "fmt"
import "math"

type Matrix interface {
	Mul() Matrix16
}

type Matrix16 []float64

func NewMat16() Matrix16 {
	return Matrix16{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func (mat Matrix16) String() string {
	return fmt.Sprintf("%f %f %f %f \n%f %f %f %f \n%f %f %f %f \n%f %f %f %f",
		mat[0], mat[1], mat[2], mat[3],
		mat[4], mat[5], mat[6], mat[7],
		mat[8], mat[9], mat[10], mat[11],
		mat[12], mat[13], mat[14], mat[15])
}

func (dest Matrix16) Transpose(mat Matrix16) {
	a01 := mat[1]
	a02 := mat[2]
	a03 := mat[3]
	a12 := mat[6]
	a13 := mat[7]
	a23 := mat[11]

	if &mat == &dest {
		mat[1] = mat[4]
		mat[2] = mat[8]
		mat[3] = mat[12]
		mat[4] = a01
		mat[6] = mat[9]
		mat[7] = mat[13]
		mat[8] = a02
		mat[9] = a12
		mat[11] = mat[14]
		mat[12] = a03
		mat[13] = a13
		mat[14] = a23
		return
	}
	dest[0] = mat[0]
	dest[1] = mat[4]
	dest[2] = mat[8]
	dest[3] = mat[12]
	dest[4] = mat[1]
	dest[5] = mat[5]
	dest[6] = mat[9]
	dest[7] = mat[13]
	dest[8] = mat[2]
	dest[9] = mat[6]
	dest[10] = mat[10]
	dest[11] = mat[14]
	dest[12] = mat[3]
	dest[13] = mat[7]
	dest[14] = mat[11]
	dest[15] = mat[15]
}

func (mat Matrix16) Determinant() float64 {
	a00 := mat[0]
	a01 := mat[1]
	a02 := mat[2]
	a03 := mat[3]

	a10 := mat[4]
	a11 := mat[5]
	a12 := mat[6]
	a13 := mat[7]

	a20 := mat[8]
	a21 := mat[9]
	a22 := mat[10]
	a23 := mat[11]

	a30 := mat[12]
	a31 := mat[13]
	a32 := mat[14]
	a33 := mat[15]

	return (a30*a21*a12*a03 - a20*a31*a12*a03 - a30*a11*a22*a03 + a10*a31*a22*a03 +
		a20*a11*a32*a03 - a10*a21*a32*a03 - a30*a21*a02*a13 + a20*a31*a02*a13 +
		a30*a01*a22*a13 - a00*a31*a22*a13 - a20*a01*a32*a13 + a00*a21*a32*a13 +
		a30*a11*a02*a23 - a10*a31*a02*a23 - a30*a01*a12*a23 + a00*a31*a12*a23 +
		a10*a01*a32*a23 - a00*a11*a32*a23 - a20*a11*a02*a33 + a10*a21*a02*a33 +
		a20*a01*a12*a33 - a00*a21*a12*a33 - a10*a01*a22*a33 + a00*a11*a22*a33)
}

func (dest Matrix16) Inverse(mat Matrix16) {
	a00 := mat[0]
	a01 := mat[1]
	a02 := mat[2]
	a03 := mat[3]

	a10 := mat[4]
	a11 := mat[5]
	a12 := mat[6]
	a13 := mat[7]

	a20 := mat[8]
	a21 := mat[9]
	a22 := mat[10]
	a23 := mat[11]

	a30 := mat[12]
	a31 := mat[13]
	a32 := mat[14]
	a33 := mat[15]

	b00 := a00*a11 - a01*a10
	b01 := a00*a12 - a02*a10
	b02 := a00*a13 - a03*a10
	b03 := a01*a12 - a02*a11
	b04 := a01*a13 - a03*a11
	b05 := a02*a13 - a03*a12
	b06 := a20*a31 - a21*a30
	b07 := a20*a32 - a22*a30
	b08 := a20*a33 - a23*a30
	b09 := a21*a32 - a22*a31
	b10 := a21*a33 - a23*a31
	b11 := a22*a33 - a23*a32

	d := (b00*b11 - b01*b10 + b02*b09 + b03*b08 - b04*b07 + b05*b06)

	// Calculate the determinant
	if d == 0 {
		return
	}
	invDet := 1 / d

	dest[0] = (a11*b11 - a12*b10 + a13*b09) * invDet
	dest[1] = (-a01*b11 + a02*b10 - a03*b09) * invDet
	dest[2] = (a31*b05 - a32*b04 + a33*b03) * invDet
	dest[3] = (-a21*b05 + a22*b04 - a23*b03) * invDet
	dest[4] = (-a10*b11 + a12*b08 - a13*b07) * invDet
	dest[5] = (a00*b11 - a02*b08 + a03*b07) * invDet
	dest[6] = (-a30*b05 + a32*b02 - a33*b01) * invDet
	dest[7] = (a20*b05 - a22*b02 + a23*b01) * invDet
	dest[8] = (a10*b10 - a11*b08 + a13*b06) * invDet
	dest[9] = (-a00*b10 + a01*b08 - a03*b06) * invDet
	dest[10] = (a30*b04 - a31*b02 + a33*b00) * invDet
	dest[11] = (-a20*b04 + a21*b02 - a23*b00) * invDet
	dest[12] = (-a10*b09 + a11*b07 - a12*b06) * invDet
	dest[13] = (a00*b09 - a01*b07 + a02*b06) * invDet
	dest[14] = (-a30*b03 + a31*b01 - a32*b00) * invDet
	dest[15] = (a20*b03 - a21*b01 + a22*b00) * invDet
}

func (dest Matrix16) Mat16Identity() {
	dest[0] = 1
	dest[1] = 0
	dest[2] = 0
	dest[3] = 0
	dest[4] = 0
	dest[5] = 1
	dest[6] = 0
	dest[7] = 0
	dest[8] = 0
	dest[9] = 0
	dest[10] = 1
	dest[11] = 0
	dest[12] = 0
	dest[13] = 0
	dest[14] = 0
	dest[15] = 1
}

func (ret Matrix16) Mat16MulMat16(mat, mat2 Matrix16) {

	// Cache the matrix values (makes for huge speed increses!)
	a00, a01, a02, a03, a10, a11, a12, a13, a20, a21, a22, a23, a30, a31, a32, a33 :=
		mat[0], mat[1], mat[2], mat[3],
		mat[4], mat[5], mat[6], mat[7],
		mat[8], mat[9], mat[10], mat[11],
		mat[12], mat[13], mat[14], mat[15]

	b00, b01, b02, b03, b10, b11, b12, b13, b20, b21, b22, b23, b30, b31, b32, b33 :=
		mat2[0], mat2[1], mat2[2], mat2[3],
		mat2[4], mat2[5], mat2[6], mat2[7],
		mat2[8], mat2[9], mat2[10], mat2[11],
		mat2[12], mat2[13], mat2[14], mat2[15]

	ret[0] = b00*a00 + b01*a10 + b02*a20 + b03*a30
	ret[1] = b00*a01 + b01*a11 + b02*a21 + b03*a31
	ret[2] = b00*a02 + b01*a12 + b02*a22 + b03*a32
	ret[3] = b00*a03 + b01*a13 + b02*a23 + b03*a33
	ret[4] = b10*a00 + b11*a10 + b12*a20 + b13*a30
	ret[5] = b10*a01 + b11*a11 + b12*a21 + b13*a31
	ret[6] = b10*a02 + b11*a12 + b12*a22 + b13*a32
	ret[7] = b10*a03 + b11*a13 + b12*a23 + b13*a33
	ret[8] = b20*a00 + b21*a10 + b22*a20 + b23*a30
	ret[9] = b20*a01 + b21*a11 + b22*a21 + b23*a31
	ret[10] = b20*a02 + b21*a12 + b22*a22 + b23*a32
	ret[11] = b20*a03 + b21*a13 + b22*a23 + b23*a33
	ret[12] = b30*a00 + b31*a10 + b32*a20 + b33*a30
	ret[13] = b30*a01 + b31*a11 + b32*a21 + b33*a31
	ret[14] = b30*a02 + b31*a12 + b32*a22 + b33*a32
	ret[15] = b30*a03 + b31*a13 + b32*a23 + b33*a33
}

func (dest Matrix16) Translate(mat Matrix16, vec Vector4) {
	x := vec[0]
	y := vec[1]
	z := vec[2]

	a00 := mat[0]
	a01 := mat[1]
	a02 := mat[2]
	a03 := mat[3]
	a10 := mat[4]
	a11 := mat[5]
	a12 := mat[6]
	a13 := mat[7]
	a20 := mat[8]
	a21 := mat[9]
	a22 := mat[10]
	a23 := mat[11]

	dest[0] = a00
	dest[1] = a01
	dest[2] = a02
	dest[3] = a03
	dest[4] = a10
	dest[5] = a11
	dest[6] = a12
	dest[7] = a13
	dest[8] = a20
	dest[9] = a21
	dest[10] = a22
	dest[11] = a23

	dest[12] = a00*x + a10*y + a20*z + mat[12]
	dest[13] = a01*x + a11*y + a21*z + mat[13]
	dest[14] = a02*x + a12*y + a22*z + mat[14]
	dest[15] = a03*x + a13*y + a23*z + mat[15]
}

func (dest Matrix16) Scale(mat Matrix16, vec Vector4) {
	x := vec[0]
	y := vec[1]
	z := vec[2]

	dest[0] = mat[0] * x
	dest[1] = mat[1] * x
	dest[2] = mat[2] * x
	dest[3] = mat[3] * x
	dest[4] = mat[4] * y
	dest[5] = mat[5] * y
	dest[6] = mat[6] * y
	dest[7] = mat[7] * y
	dest[8] = mat[8] * z
	dest[9] = mat[9] * z
	dest[10] = mat[10] * z
	dest[11] = mat[11] * z
	dest[12] = mat[12]
	dest[13] = mat[13]
	dest[14] = mat[14]
	dest[15] = mat[15]
}

func (dest Matrix16) Rotate(mat Matrix16, axis Vector3, angle float64) {
	x := axis[0]
	y := axis[1]
	z := axis[2]
	len := math.Sqrt(x*x + y*y + z*z)

	if len != 1 {
		len := 1 / len
		x *= len
		y *= len
		z *= len
	}

	s := math.Sin(angle)
	c := math.Cos(angle)
	t := 1 - c

	a00 := mat[0]
	a01 := mat[1]
	a02 := mat[2]
	a03 := mat[3]
	a10 := mat[4]
	a11 := mat[5]
	a12 := mat[6]
	a13 := mat[7]
	a20 := mat[8]
	a21 := mat[9]
	a22 := mat[10]
	a23 := mat[11]

	// Construct the elements of the rotation matrix
	b00 := x*x*t + c
	b01 := y*x*t + z*s
	b02 := z*x*t - y*s
	b10 := x*y*t - z*s
	b11 := y*y*t + c
	b12 := z*y*t + x*s
	b20 := x*z*t + y*s
	b21 := y*z*t - x*s
	b22 := z*z*t + c

	dest[12] = mat[12]
	dest[13] = mat[13]
	dest[14] = mat[14]
	dest[15] = mat[15]

	// Perform rotation-specific matrix multiplication
	dest[0] = a00*b00 + a10*b01 + a20*b02
	dest[1] = a01*b00 + a11*b01 + a21*b02
	dest[2] = a02*b00 + a12*b01 + a22*b02
	dest[3] = a03*b00 + a13*b01 + a23*b02

	dest[4] = a00*b10 + a10*b11 + a20*b12
	dest[5] = a01*b10 + a11*b11 + a21*b12
	dest[6] = a02*b10 + a12*b11 + a22*b12
	dest[7] = a03*b10 + a13*b11 + a23*b12

	dest[8] = a00*b20 + a10*b21 + a20*b22
	dest[9] = a01*b20 + a11*b21 + a21*b22
	dest[10] = a02*b20 + a12*b21 + a22*b22
	dest[11] = a03*b20 + a13*b21 + a23*b22
}

func (dest Matrix16) LookAt(eye Vector3, center Vector3, up Vector3) {
	eyex := eye[0]
	eyey := eye[1]
	eyez := eye[2]
	upx := up[0]
	upy := up[1]
	upz := up[2]
	centerx := center[0]
	centery := center[1]
	centerz := center[2]

	if eyex == centerx && eyey == centery && eyez == centerz {
		dest.Mat16Identity()
		return
	}

	//vec3.direction(eye, center, z);
	z0 := eyex - centerx
	z1 := eyey - centery
	z2 := eyez - centerz

	// normalize (no check needed for 0 because of early return)
	len := 1 / math.Sqrt(z0*z0+z1*z1+z2*z2)
	z0 *= len
	z1 *= len
	z2 *= len

	//vec3.normalize(vec3.cross(up, z, x));
	x0 := upy*z2 - upz*z1
	x1 := upz*z0 - upx*z2
	x2 := upx*z1 - upy*z0
	len = math.Sqrt(x0*x0 + x1*x1 + x2*x2)
	if len == 0 {
		x0 = 0
		x1 = 0
		x2 = 0
	} else {
		len = 1 / len
		x0 *= len
		x1 *= len
		x2 *= len
	}

	//vec3.normalize(vec3.cross(z, x, y));
	y0 := z1*x2 - z2*x1
	y1 := z2*x0 - z0*x2
	y2 := z0*x1 - z1*x0

	len = math.Sqrt(y0*y0 + y1*y1 + y2*y2)
	if len == 0 {
		y0 = 0
		y1 = 0
		y2 = 0
	} else {
		len = 1 / len
		y0 *= len
		y1 *= len
		y2 *= len
	}

	dest[0] = x0
	dest[1] = y0
	dest[2] = z0
	dest[3] = 0
	dest[4] = x1
	dest[5] = y1
	dest[6] = z1
	dest[7] = 0
	dest[8] = x2
	dest[9] = y2
	dest[10] = z2
	dest[11] = 0
	dest[12] = -(x0*eyex + x1*eyey + x2*eyez)
	dest[13] = -(y0*eyex + y1*eyey + y2*eyez)
	dest[14] = -(z0*eyex + z1*eyey + z2*eyez)
	dest[15] = 1

}
