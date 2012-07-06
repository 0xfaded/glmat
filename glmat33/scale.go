package glmat33

// lambda * v.
func (m *Mat4) Scaled(lambda float64) (scaled *Mat4) {
	scaled = new(Mat4)
	for i := range m {
		scaled[i] = lambda * m[i]
	}
	return scaled
}

// Scale matrix along x, y and z axes.
func (m *Mat4) ScaledXYZ(x, y, z float64) (s *Mat4) {
	s = new(Mat4)
	s[0] = m[0] * x; s[4] = m[4] * x; s[ 8] = m[ 8] * x; s[12] = m[12] * x
	s[1] = m[1] * y; s[5] = m[5] * y; s[ 9] = m[ 9] * y; s[13] = m[13] * y
	s[2] = m[2] * z; s[6] = m[6] * z; s[10] = m[10] * z; s[14] = m[14] * z
	s[3] = m[3]    ; s[7] = m[7]    ; s[11] = m[11]    ; s[15] = m[15]
	return s
}

// lambda * v.
func (v Vec4) Scaled(lambda float64) (scaled Vec4) {
	scaled = v
	scaled.Scale(lambda)
	return scaled
}

// lambda * v.
func (v Vec3) Scaled(lambda float64) (scaled Vec3) {
	scaled = v
	scaled.Scale(lambda)
	return scaled
}

// lambda * v.
func (v Vec2) Scaled(lambda float64) (scaled Vec2) {
	scaled = v
	scaled.Scale(lambda)
	return scaled
}

// m *= lambda.
func (m *Mat4) Scale(lambda float64) *Mat4 {
	for i := range m {
		m[i] *= lambda
	}
	return m
}

// Scale matrix along x, y and z axes.
func (m *Mat4) ScaleXYZ(x, y, z float64) *Mat4 {
	m[0] *= x; m[4] *= x; m[ 8] *= x; m[12] *= x
	m[1] *= y; m[5] *= y; m[ 9] *= y; m[13] *= y
	m[2] *= z; m[6] *= z; m[10] *= z; m[14] *= z

	return m
}

// v *= lambda.
func (v *Vec4) Scale(lambda float64) *Vec4 {
	v[0] *= lambda
	v[1] *= lambda
	v[2] *= lambda
	v[3] *= lambda

	return v
}

// v *= lambda.
func (v *Vec3) Scale(lambda float64) *Vec3 {
	v[0] *= lambda
	v[1] *= lambda
	v[2] *= lambda

	return v
}

// v *= lambda.
func (v *Vec2) Scale(lambda float64) *Vec2 {
	v[0] *= lambda
	v[1] *= lambda

	return v
}

