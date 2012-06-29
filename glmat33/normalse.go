package glmat

import "math"

// Normalise the Vec4 v.
func (v *Vec4) Normalise() *Vec4 {
	h := float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2] + v[3]*v[3])))
	if h == 0 {
		v[0] = float32(math.NaN())
		v[1] = float32(math.NaN())
		v[2] = float32(math.NaN())
		v[3] = float32(math.NaN())
		return v
	}

	h1 := 1/h

	v[0] *= h1
	v[1] *= h1
	v[2] *= h1
	v[3] *= h1

	return v
}

// Return a vector parallel to Vec4 v with unit length.
func (v Vec4) Normalised() (normalised Vec4) {
	normalised = v
	normalised.Normalise()
	return normalised
}

// Normalise the Vec3 v.
func (v *Vec3) Normalise() *Vec3 {
	h := float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])))
	if h == 0 {
		v[0] = float32(math.NaN())
		v[1] = float32(math.NaN())
		v[2] = float32(math.NaN())
		return v
	}

	h1 := 1/h

	v[0] *= h1
	v[1] *= h1
	v[2] *= h1

	return v
}

// Return a vector parallel to Vec3 v with unit length.
func (v Vec3) Normalised() (normalised Vec3) {
	normalised = v
	normalised.Normalise()
	return normalised
}

// Normalise the Vec2 v.
func (v *Vec2) Normalise() *Vec2 {
	h := float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1])))
	if h == 0 {
		v[0] = float32(math.NaN())
		v[1] = float32(math.NaN())
		return v
	}

	h1 := 1/h

	v[0] *= h1
	v[1] *= h1

	return v
}

// Return a vector parallel to Vec2 v with unit length.
func (v Vec2) Normalised() (normalised Vec2) {
	normalised = v
	normalised.Normalise()
	return normalised
}


