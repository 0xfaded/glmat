package glmat42

import (
	"fmt"
	"errors"
)

// Calucalte the inverse of a 4 by 4 matrix. Panics if matrix is non invertable.
// See Inverse for safe version that returns an error.
func (m *Mat4) Inversed() (inv *Mat4) {
	// Lifted from http://www.geometrictools.com/LibMathematics/Algebra/Wm5Matrix4.inl
	a0 := m[ 0]*m[ 5] - m[ 1]*m[ 4]
	a1 := m[ 0]*m[ 6] - m[ 2]*m[ 4]
	a2 := m[ 0]*m[ 7] - m[ 3]*m[ 4]
	a3 := m[ 1]*m[ 6] - m[ 2]*m[ 5]
	a4 := m[ 1]*m[ 7] - m[ 3]*m[ 5]
	a5 := m[ 2]*m[ 7] - m[ 3]*m[ 6]
	b0 := m[ 8]*m[13] - m[ 9]*m[12]
	b1 := m[ 8]*m[14] - m[10]*m[12]
	b2 := m[ 8]*m[15] - m[11]*m[12]
	b3 := m[ 9]*m[14] - m[10]*m[13]
	b4 := m[ 9]*m[15] - m[11]*m[13]
	b5 := m[10]*m[15] - m[11]*m[14]

	det := a0*b5 - a1*b4 + a2*b3 + a3*b2 - a4*b1 + a5*b0
	if det == 0 {
		errStr := fmt.Sprintf("Matrix (%v) is non invertable", (*[16]float32)(m))
		panic(errors.New(errStr))
	}

	inv = new(Mat4)

	inv[ 0] =   m[ 5]*b5 - m[ 6]*b4 + m[ 7]*b3
	inv[ 4] = - m[ 4]*b5 + m[ 6]*b2 - m[ 7]*b1
	inv[ 8] =   m[ 4]*b4 - m[ 5]*b2 + m[ 7]*b0
	inv[12] = - m[ 4]*b3 + m[ 5]*b1 - m[ 6]*b0
	inv[ 1] = - m[ 1]*b5 + m[ 2]*b4 - m[ 3]*b3
	inv[ 5] =   m[ 0]*b5 - m[ 2]*b2 + m[ 3]*b1
	inv[ 9] = - m[ 0]*b4 + m[ 1]*b2 - m[ 3]*b0
	inv[13] =   m[ 0]*b3 - m[ 1]*b1 + m[ 2]*b0
	inv[ 2] =   m[13]*a5 - m[14]*a4 + m[15]*a3
	inv[ 6] = - m[12]*a5 + m[14]*a2 - m[15]*a1
	inv[10] =   m[12]*a4 - m[13]*a2 + m[15]*a0
	inv[14] = - m[12]*a3 + m[13]*a1 - m[14]*a0
	inv[ 3] = - m[ 9]*a5 + m[10]*a4 - m[11]*a3
	inv[ 7] =   m[ 8]*a5 - m[10]*a2 + m[11]*a1
	inv[11] = - m[ 8]*a4 + m[ 9]*a2 - m[11]*a0
	inv[15] =   m[ 8]*a3 - m[ 9]*a1 + m[10]*a0

	det = 1.0 / det

	for i := range inv {
		inv[i] *= det
	}
	return inv
}

// Inverse m. Returns an error if m is non-invertable and m remains
// untouched.
func (m *Mat4) Inverse() (err error) {
	// Lifted from http://www.geometrictools.com/LibMathematics/Algebra/Wm5Matrix4.inl
	a0 := m[ 0]*m[ 5] - m[ 1]*m[ 4]
	a1 := m[ 0]*m[ 6] - m[ 2]*m[ 4]
	a2 := m[ 0]*m[ 7] - m[ 3]*m[ 4]
	a3 := m[ 1]*m[ 6] - m[ 2]*m[ 5]
	a4 := m[ 1]*m[ 7] - m[ 3]*m[ 5]
	a5 := m[ 2]*m[ 7] - m[ 3]*m[ 6]
	b0 := m[ 8]*m[13] - m[ 9]*m[12]
	b1 := m[ 8]*m[14] - m[10]*m[12]
	b2 := m[ 8]*m[15] - m[11]*m[12]
	b3 := m[ 9]*m[14] - m[10]*m[13]
	b4 := m[ 9]*m[15] - m[11]*m[13]
	b5 := m[10]*m[15] - m[11]*m[14]

	det := a0*b5 - a1*b4 + a2*b3 + a3*b2 - a4*b1 + a5*b0
	if det == 0 {
		errStr := fmt.Sprintf("Matrix (%v) is non invertable", (*[16]float32)(m))
		return errors.New(errStr)
	}

	inv := new(Mat4)

	inv[ 0] =   m[ 5]*b5 - m[ 6]*b4 + m[ 7]*b3
	inv[ 4] = - m[ 4]*b5 + m[ 6]*b2 - m[ 7]*b1
	inv[ 8] =   m[ 4]*b4 - m[ 5]*b2 + m[ 7]*b0
	inv[12] = - m[ 4]*b3 + m[ 5]*b1 - m[ 6]*b0
	inv[ 1] = - m[ 1]*b5 + m[ 2]*b4 - m[ 3]*b3
	inv[ 5] =   m[ 0]*b5 - m[ 2]*b2 + m[ 3]*b1
	inv[ 9] = - m[ 0]*b4 + m[ 1]*b2 - m[ 3]*b0
	inv[13] =   m[ 0]*b3 - m[ 1]*b1 + m[ 2]*b0
	inv[ 2] =   m[13]*a5 - m[14]*a4 + m[15]*a3
	inv[ 6] = - m[12]*a5 + m[14]*a2 - m[15]*a1
	inv[10] =   m[12]*a4 - m[13]*a2 + m[15]*a0
	inv[14] = - m[12]*a3 + m[13]*a1 - m[14]*a0
	inv[ 3] = - m[ 9]*a5 + m[10]*a4 - m[11]*a3
	inv[ 7] =   m[ 8]*a5 - m[10]*a2 + m[11]*a1
	inv[11] = - m[ 8]*a4 + m[ 9]*a2 - m[11]*a0
	inv[15] =   m[ 8]*a3 - m[ 9]*a1 + m[10]*a0

	det = 1.0 / det

	for i := range m {
		m[i] = inv[i] * det
	}

	return nil
}

