package glmat42

// Create a Mat4f from a a Mat4.
func (m *Mat4) F(mf *Mat4f) {
	mf = new(Mat4f)
	for i := range *m {
		(*mf)[i] = float32((*m)[i])
	}
}

// Create a Vec4f from a a Vec4.
func (v Vec4) F(vf Vec4f) {
	vf[0] = float32(v[0])
	vf[1] = float32(v[1])
	vf[2] = float32(v[2])
	vf[3] = float32(v[3])
}

// Create a Vec3f from a a Vec3.
func (v Vec3) F(vf Vec3f) {
	vf[0] = float32(v[0])
	vf[1] = float32(v[1])
	vf[2] = float32(v[2])
}

// Create a Vec2f from a a Vec2.
func (v Vec2) F(vf Vec2f) {
	vf[0] = float32(v[0])
	vf[1] = float32(v[1])
}

// Create a Mat4 from a a Mat4f.
func (mf *Mat4f) D(m *Mat4) {
	m = new(Mat4)
	for i := range *mf {
		(*m)[i] = float64((*mf)[i])
	}
}

// Create a Vec4 from a a Vec4f.
func (vf Vec4f) D(v Vec4) {
	v[0] = float64(vf[0])
	v[1] = float64(vf[1])
	v[2] = float64(vf[2])
	v[3] = float64(vf[3])
}

// Create a Vec3 from a a Vec3f.
func (vf Vec3f) D(v Vec3) {
	v[0] = float64(vf[0])
	v[1] = float64(vf[1])
	v[2] = float64(vf[2])
}

// Create a Vec2 from a a Vec2f.
func (vf Vec2f) D(v Vec2) {
	v[0] = float64(vf[0])
	v[1] = float64(vf[1])
}
