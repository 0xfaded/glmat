package glmat42

import "math"

func Perspective4(fovy, aspect, zNear, zFar float64) (m *Mat4) {
	cotangent := 1 / math.Tan(fovy)
	deltaZ    := zFar - zNear

	m = (*Mat4)(&[16]float64{cotangent / aspect         , 0, 0,  0,
	                         0, cotangent                  , 0,  0,
	                         0, 0, -(zFar + zNear) / deltaZ   , -1,
	                         0, 0, -2 * zNear * zFar / deltaZ ,  0})

	return m;
}

func (m *Mat4) Perspective(fovy, aspect, zNear, zFar float64) *Mat4 {
	*m = *Perspective4(fovy, aspect, zNear, zFar)

	return m
}

