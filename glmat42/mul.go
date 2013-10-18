package glmat42

// a = b * a.
func (a *Mat4) MulI(b *Mat4) *Mat4 {
	// Temporary clone of a.
	t := *a

        a[0]  = b[ 0] * t[ 0] + b[ 1] * t[ 4] + b[ 2] * t[ 8] + b[ 3] * t[12]
        a[1]  = b[ 0] * t[ 1] + b[ 1] * t[ 5] + b[ 2] * t[ 9] + b[ 3] * t[13]
        a[2]  = b[ 0] * t[ 2] + b[ 1] * t[ 6] + b[ 2] * t[10] + b[ 3] * t[14]
        a[3]  = b[ 0] * t[ 3] + b[ 1] * t[ 7] + b[ 2] * t[11] + b[ 3] * t[15]
        a[4]  = b[ 4] * t[ 0] + b[ 5] * t[ 4] + b[ 6] * t[ 8] + b[ 7] * t[12]
        a[5]  = b[ 4] * t[ 1] + b[ 5] * t[ 5] + b[ 6] * t[ 9] + b[ 7] * t[13]
        a[6]  = b[ 4] * t[ 2] + b[ 5] * t[ 6] + b[ 6] * t[10] + b[ 7] * t[14]
        a[7]  = b[ 4] * t[ 3] + b[ 5] * t[ 7] + b[ 6] * t[11] + b[ 7] * t[15]
        a[8]  = b[ 8] * t[ 0] + b[ 9] * t[ 4] + b[10] * t[ 8] + b[11] * t[12]
        a[9]  = b[ 8] * t[ 1] + b[ 9] * t[ 5] + b[10] * t[ 9] + b[11] * t[13]
        a[10] = b[ 8] * t[ 2] + b[ 9] * t[ 6] + b[10] * t[10] + b[11] * t[14]
        a[11] = b[ 8] * t[ 3] + b[ 9] * t[ 7] + b[10] * t[11] + b[11] * t[15]
        a[12] = b[12] * t[ 0] + b[13] * t[ 4] + b[14] * t[ 8] + b[15] * t[12]
        a[13] = b[12] * t[ 1] + b[13] * t[ 5] + b[14] * t[ 9] + b[15] * t[13]
        a[14] = b[12] * t[ 2] + b[13] * t[ 6] + b[14] * t[10] + b[15] * t[14]
        a[15] = b[12] * t[ 3] + b[13] * t[ 7] + b[14] * t[11] + b[15] * t[15]

	return a
}

// c = a * b.
func (a *Mat4) MulM(b *Mat4) (c *Mat4) {
	c = new(Mat4)

        c[0]  = b[ 0] * a[ 0] + b[ 1] * a[ 4] + b[ 2] * a[ 8] + b[ 3] * a[12]
        c[1]  = b[ 0] * a[ 1] + b[ 1] * a[ 5] + b[ 2] * a[ 9] + b[ 3] * a[13]
        c[2]  = b[ 0] * a[ 2] + b[ 1] * a[ 6] + b[ 2] * a[10] + b[ 3] * a[14]
        c[3]  = b[ 0] * a[ 3] + b[ 1] * a[ 7] + b[ 2] * a[11] + b[ 3] * a[15]
        c[4]  = b[ 4] * a[ 0] + b[ 5] * a[ 4] + b[ 6] * a[ 8] + b[ 7] * a[12]
        c[5]  = b[ 4] * a[ 1] + b[ 5] * a[ 5] + b[ 6] * a[ 9] + b[ 7] * a[13]
        c[6]  = b[ 4] * a[ 2] + b[ 5] * a[ 6] + b[ 6] * a[10] + b[ 7] * a[14]
        c[7]  = b[ 4] * a[ 3] + b[ 5] * a[ 7] + b[ 6] * a[11] + b[ 7] * a[15]
        c[8]  = b[ 8] * a[ 0] + b[ 9] * a[ 4] + b[10] * a[ 8] + b[11] * a[12]
        c[9]  = b[ 8] * a[ 1] + b[ 9] * a[ 5] + b[10] * a[ 9] + b[11] * a[13]
        c[10] = b[ 8] * a[ 2] + b[ 9] * a[ 6] + b[10] * a[10] + b[11] * a[14]
        c[11] = b[ 8] * a[ 3] + b[ 9] * a[ 7] + b[10] * a[11] + b[11] * a[15]
        c[12] = b[12] * a[ 0] + b[13] * a[ 4] + b[14] * a[ 8] + b[15] * a[12]
        c[13] = b[12] * a[ 1] + b[13] * a[ 5] + b[14] * a[ 9] + b[15] * a[13]
        c[14] = b[12] * a[ 2] + b[13] * a[ 6] + b[14] * a[10] + b[15] * a[14]
        c[15] = b[12] * a[ 3] + b[13] * a[ 7] + b[14] * a[11] + b[15] * a[15]

	return c
}

// c = a * b.
func (a *Mat4) MulV(b *Vec4) (c Vec4) {
        c[0] = a[ 0] * b[0] + a[ 4] * b[1] + a[ 8] * b[2] + a[12] * b[3]
        c[1] = a[ 1] * b[0] + a[ 5] * b[1] + a[ 9] * b[2] + a[13] * b[3]
        c[2] = a[ 2] * b[0] + a[ 6] * b[1] + a[10] * b[2] + a[14] * b[3]
        c[3] = a[ 3] * b[0] + a[ 7] * b[1] + a[11] * b[2] + a[15] * b[3]

	return c
}

// c = (a * b)[0:3].
func (a *Mat4) MulV3(b *Vec3) (c Vec3) {
        c[0] = a[ 0] * b[0] + a[ 4] * b[1] + a[ 8] * b[2] + a[12] * 1
        c[1] = a[ 1] * b[0] + a[ 5] * b[1] + a[ 9] * b[2] + a[13] * 1
        c[2] = a[ 2] * b[0] + a[ 6] * b[1] + a[10] * b[2] + a[14] * 1

	return c
}

// Hamilton Product c = a x b
func (a Quat) Mul(b Quat) (c Quat) {
	c[0] = a[0]*b[0] - a[1]*b[1] - a[2]*b[2] - a[3]*b[3]
	c[1] = a[0]*b[1] + a[1]*b[0] + a[2]*b[3] - a[3]*b[2]
	c[2] = a[0]*b[2] - a[1]*b[3] + a[2]*b[0] + a[3]*b[1]
	c[3] = a[0]*b[3] + a[1]*b[2] - a[2]*b[1] + a[3]*b[0]
	return c
}

// Equivalent to rotating b, c = a * b
func (a Quat) MulV3(b Vec3) (c Vec3) {
	return a.Mul(AxisQ(b)).Mul(a.Inversed()).Axis()
}
