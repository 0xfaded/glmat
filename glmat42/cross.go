package glmat42

// a x b.
func (a Vec3) Cross(b *Vec3) (cross Vec3) {
	cross[0] = a[1] * b[2] - a[2] * b[1]
	cross[1] = a[0] * b[2] - a[2] * b[0]
	cross[2] = a[0] * b[1] - a[1] * b[0]

	return cross
}

// a x b.
func (a Vec2) Cross(b *Vec2) (cross float64) {
	return a[0] * b[1] - a[1] * b[0]
}

// a x= b.
func (a *Vec3) CrossI(b *Vec3) *Vec3 {
	*a = a.Cross(b)

	return a
}
