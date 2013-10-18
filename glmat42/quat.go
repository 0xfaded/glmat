package glmat42

import (
	"math"
)

func (q Quat) Angle() float64 {
	return math.Acos(q[0])
}

func (q Quat) Axis() (v Vec3) {
	s := 1/math.Sin(q.Angle())
	v[0] = q[1]*s
	v[1] = q[2]*s
	v[2] = q[3]*s
	return v
}

func (q Quat) ScaleAngle(lambda float64) Quat {
	angle := q.Angle() * lambda
	return AngleAxisQ(angle, q.Axis())
}

func (q Quat) Mat4() (m *Mat4) {
	m = new(Mat4)
	m[ 0] = q[0]*q[0] + q[1]*q[1] - q[2]*q[2] - q[3]*q[3]
	m[ 1] = 2*(q[1]*q[2] + q[0]*q[3])
	m[ 2] = 2*(q[1]*q[3] - q[0]*q[2])
	m[ 3] = 0

	m[ 4] = 2*(q[1]*q[2] - q[0]*q[3])
	m[ 5] = q[0]*q[0] - q[1]*q[1] + q[2]*q[2] - q[3]*q[3]
	m[ 6] = 2*(q[2]*q[3] + q[0]*q[1])
	m[ 7] = 0

	m[ 8] = 2*(q[1]*q[3] + q[0]*q[2])
	m[ 9] = 2*(q[2]*q[3] - q[0]*q[1])
	m[10] = q[0]*q[0] - q[1]*q[1] - q[2]*q[2] + q[3]*q[3]
	m[11] = 0

	// m[12] = 0 Already zeroed by runtime
	// m[13] = 0
	// m[14] = 0
	m[15] = 1
	return m
}
