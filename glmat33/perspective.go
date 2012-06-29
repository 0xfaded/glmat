package glmat

import "math"

func Perspective4(fovy, aspect, zNear, zFar float32) (m *Mat4) {
	cotangent := 1 / float32(math.Tan(float64(fovy)))
	deltaZ    := zFar - zNear

	m = (*Mat4)(&[16]float32{cotangent / aspect         , 0, 0,  0,
	                         0, cotangent                  , 0,  0,
	                         0, 0, -(zFar + zNear) / deltaZ   , -1,
	                         0, 0, -2 * zNear * zFar / deltaZ ,  0})

	return m;
}

func (m *Mat4) Perspective(fovy, aspect, zNear, zFar float32) *Mat4 {
	*m = *Perspective4(fovy, aspect, zNear, zFar)

	return m
}

