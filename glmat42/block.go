package glmat42

// Return a Vec3 constructed using the first 3 elements of v.
func (v Vec4) V3() Vec3 {
	return Vec3{v[0], v[1], v[2]}
}

// Return a Vec2 constructed using the first 2 elements of v.
func (v Vec4) V2() Vec2 {
	return Vec2{v[0], v[1]}
}

// Return a Vec4 constructed using elements of v appended by w.
func (v Vec3) V4(w float64) Vec4 {
	return Vec4{v[0], v[1], v[2], w}
}

// Return a Vec2 constructed using the first 2 elements of v.
func (v Vec3) V2() Vec2 {
	return Vec2{v[0], v[1]}
}

// Return a Vec4 constructed using elements of v appended by z and w.
func (v Vec2) V4(z, w float64) Vec4 {
	return Vec4{v[0], v[1], z, w}
}

// Return a Vec3 constructed using elements of v appended by z.
func (v Vec2) V3(z float64) Vec3 {
	return Vec3{v[0], v[1], z}
}

