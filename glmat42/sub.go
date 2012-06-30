package glmat42

// a - b.
func (a Vec4) Sub(b *Vec4) (difference Vec4) {
	difference[0] = a[0] - b[0]
	difference[1] = a[1] - b[1]
	difference[2] = a[2] - b[2]
	difference[3] = a[3] - b[3]

	return difference
}

// a - b.
func (a Vec3) Sub(b *Vec3) (difference Vec3) {
	difference[0] = a[0] - b[0]
	difference[1] = a[1] - b[1]
	difference[2] = a[2] - b[2]

	return difference
}

// a - b.
func (a Vec2) Sub(b *Vec2) (difference Vec2) {
	difference[0] = a[0] - b[0]
	difference[1] = a[1] - b[1]

	return difference
}

// a -= b.
func (a *Vec4) SubI(b *Vec4) *Vec4 {
	a[0] -= b[0]
	a[1] -= b[1]
	a[2] -= b[2]
	a[3] -= b[3]

	return a
}

// a -= b.
func (a *Vec3) SubI(b *Vec3) *Vec3 {
	a[0] -= b[0]
	a[1] -= b[1]
	a[2] -= b[2]

	return a
}

// a -= b.
func (a *Vec2) SubI(b *Vec2) *Vec2 {
	a[0] -= b[0]
	a[1] -= b[1]

	return a
}

