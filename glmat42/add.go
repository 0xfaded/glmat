package glmat42

// a + b.
func (a Vec4) Add(b *Vec4) (sum Vec4) {
	sum[0] = a[0] + b[0]
	sum[1] = a[1] + b[1]
	sum[2] = a[2] + b[2]
	sum[3] = a[3] + b[3]

	return sum
}

// a + b.
func (a Vec3) Add(b *Vec3) (sum Vec3) {
	sum[0] = a[0] + b[0]
	sum[1] = a[1] + b[1]
	sum[2] = a[2] + b[2]

	return sum
}

// a + b.
func (a Vec2) Add(b *Vec2) (sum Vec2) {
	sum[0] = a[0] + b[0]
	sum[1] = a[1] + b[1]

	return sum
}

// a += b.
func (a *Vec4) AddI(b *Vec4) *Vec4 {
	a[0] += b[0]
	a[1] += b[1]
	a[2] += b[2]
	a[3] += b[3]

	return a
}

// a += b.
func (a *Vec3) AddI(b *Vec3) *Vec3 {
	a[0] += b[0]
	a[1] += b[1]
	a[2] += b[2]

	return a
}

// a += b.
func (a *Vec2) AddI(b *Vec2) *Vec2 {
	a[0] += b[0]
	a[1] += b[1]

	return a
}

