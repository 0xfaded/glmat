package glmat42

// a . b
func (a Vec4) Dot(b *Vec4) float64 {
	return a[0] * b[0] + a[1] * b[1] + a[2] * b[2] + a[3] * b[3]
}

// a . b
func (a Vec3) Dot(b *Vec3) float64 {
	return a[0] * b[0] + a[1] * b[1] + a[2] * b[2]
}

// a . b
func (a Vec2) Dot(b *Vec2) float64 {
	return a[0] * b[0] + a[1] * b[1]
}
