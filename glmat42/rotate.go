package glmat42

import "math"

func Rotation4(angle, x, y, z float64) (m *Mat4) {
	// Work with 64 bits for precision
	xx := x
	yy := y
	zz := z

	// Normalise x y z
	hh := math.Sqrt(xx*xx + yy*yy + zz*zz)
	xx = xx / hh
	yy = yy / hh
	zz = zz / hh

	// Cache results as we go
	c := math.Cos(angle)
	s := math.Sin(angle)
	C := 1-c

	m = new(Mat4)

	// http://en.wikipedia.org/wiki/Rotation_matrix#Rotation_matrix_from_axis_and_angle
	m[ 0] =    xx*xx*C+c; m[ 4] = xx*yy*C-zz*s
	m[ 8] = xx*zz*C+yy*s; m[12] =            0

        m[ 1] = yy*xx*C+zz*s; m[ 5] =    yy*yy*C+c
	m[ 9] = yy*zz*C-xx*s; m[13] =            0

        m[ 2] = zz*xx*C-yy*s; m[ 6] = zz*yy*C+xx*s
	m[10] = zz*zz*C+c   ; m[14] =            0

	m[ 3] =            0; m[ 7] =            0
	m[11] =            0; m[15] =            1

	return m
}

func (m *Mat4) Rotate(angle, x, y, z float64) {
	rot := Rotation4(angle, x, y, z)
	m.MulI(rot)
}

func (m *Mat4) Rotated(angle, x, y, z float64) *Mat4 {
	rot := Rotation4(angle, x, y, z)
	return m.MulM(rot)
}


