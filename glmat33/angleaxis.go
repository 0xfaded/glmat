package glmat33

import (
	"math"
)

func AngleAxis4(angle float64, axis Vec3) (m *Mat4) {
	m = new(Mat4)
	m.AngleAxis(angle, axis)
	return m
}

func AngleAxisQ(angle float64, axis Vec3) (q Quat) {
	s := math.Sin(angle)
	q[0] = math.Cos(angle)
	q[1] = s * axis[0]
	q[2] = s * axis[1]
	q[3] = s * axis[2]
	return q
}

// Angle is zero
func AxisQ(axis Vec3) (q Quat) {
	q[0] = 0
	q[1] = axis[0]
	q[2] = axis[1]
	q[3] = axis[2]
	return q
}

func (m *Mat4) AngleAxis(angle float64, u Vec3) *Mat4 {
	c := math.Cos(angle)
	s := math.Sin(angle)
	c1 := 1-c

	m[ 0] = c + u[0] * u[0] * c1
	m[ 1] = u[1] * u[0] * c1 + u[2] * s
	m[ 2] = u[2] * u[0] * c1 - u[1] * s
	m[ 3] = 0

	m[ 4] = u[0] * u[1] * c1 - u[2] * s
	m[ 5] = c + u[1] * u[1] * c1
	m[ 6] = u[2] * u[1] * c1 + u[0] * s
	m[ 7] = 0

	m[ 8] = u[0] * u[2] * c1 + u[1] * s
	m[ 9] = u[1] * u[2] * c1 - u[0] * s
	m[10] = c + u[2] * u[2] * c1
	m[11] = 0

	m[13] = 0
	m[14] = 0
	m[12] = 0
	m[15] = 1

	return m
}

func (q *Quat) AngleAxis(angle float64, axis Vec3) *Quat {
	*q = AngleAxisQ(angle, axis)
	return q
}

