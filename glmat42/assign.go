package glmat42

// Assign the first 3 elements of v to those of xyz.
func (v *Vec4) Assign3(xyz *Vec3) *Vec4 {
	v[0] = xyz[0]
	v[1] = xyz[1]
	v[2] = xyz[2]

	return v
}

// Assign the first 2 elements of v to those of xy.
func (v *Vec4) Assign2(xy *Vec2) *Vec4 {
	v[0] = xy[0]
	v[1] = xy[1]

	return v
}

// Assign the first 2 elements of v to those of xy.
func (v *Vec3) Assign2(xy *Vec2) *Vec3 {
	v[0] = xy[0]
	v[1] = xy[1]

	return v
}

