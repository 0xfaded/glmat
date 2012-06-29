package glmat

import "math"

func Rotation4(angle, x, y, z float32) (m *Mat4) {
	// Work with 64 bits for precision
	xx := float64(x)
	yy := float64(y)
	zz := float64(z)

	// Normalise x y z
	hh := math.Sqrt(xx*xx + yy*yy + zz*zz)
	xx = xx / hh
	yy = yy / hh
	zz = zz / hh

	// Cache results as we go
	c := math.Cos(float64(angle))
	s := math.Sin(float64(angle))
	C := 1-c

	m = new(Mat4)

	// http://en.wikipedia.org/wiki/Rotation_matrix#Rotation_matrix_from_axis_and_angle
	m[ 0] = float32(xx*xx*C+c)   ; m[ 4] = float32(xx*yy*C-zz*s)
	m[ 8] = float32(xx*zz*C+yy*s); m[12] = float32(           0)

        m[ 1] = float32(yy*xx*C+zz*s); m[ 5] = float32(yy*yy*C+c   )
	m[ 9] = float32(yy*zz*C-xx*s); m[13] = float32(           0)

        m[ 2] = float32(zz*xx*C-yy*s); m[ 6] = float32(zz*yy*C+xx*s)
	m[10] = float32(zz*zz*C+c)   ; m[14] = float32(           0)

	m[ 3] = float32(           0); m[ 7] = float32(           0)
	m[11] = float32(           0); m[15] = float32(           1)

	return m
}

func (m *Mat4) Rotate(angle, x, y, z float32) {
	rot := Rotation4(angle, x, y, z)
	m.MulI(rot)
}

