package glmat

// lambda * v.
func (v Vec4) Scaled(lambda float32) (scaled Vec4) {
	scaled = v
	scaled.Scale(lambda)
	return scaled
}

// lambda * v.
func (v Vec3) Scaled(lambda float32) (scaled Vec3) {
	scaled = v
	scaled.Scale(lambda)
	return scaled
}

// lambda * v.
func (v Vec2) Scaled(lambda float32) (scaled Vec2) {
	scaled = v
	scaled.Scale(lambda)
	return scaled
}

// v *= lambda.
func (v *Vec4) Scale(lambda float32) *Vec4 {
	v[0] *= lambda
	v[1] *= lambda
	v[2] *= lambda
	v[3] *= lambda

	return v
}

// v *= lambda.
func (v *Vec3) Scale(lambda float32) *Vec3 {
	v[0] *= lambda
	v[1] *= lambda
	v[2] *= lambda

	return v
}

// v *= lambda.
func (v *Vec2) Scale(lambda float32) *Vec2 {
	v[0] *= lambda
	v[1] *= lambda

	return v
}

