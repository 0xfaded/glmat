package glmat33

import "math"

// Normalise the Vec4 v.
func (v *Vec4) Normalise() *Vec4 {
	h := v.Norm()
	if h == 0 {
		v[0] = math.NaN()
		v[1] = math.NaN()
		v[2] = math.NaN()
		v[3] = math.NaN()
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
	h := v.Norm()
	if h == 0 {
		v[0] = math.NaN()
		v[1] = math.NaN()
		v[2] = math.NaN()
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
	h := v.Norm()
	if h == 0 {
		v[0] = math.NaN()
		v[1] = math.NaN()
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

// |v|
func (v Vec4) Norm() float64 {
	return math.Sqrt(v.Dot(&v))
}

// |v|
func (v Vec3) Norm() float64 {
	return math.Sqrt(v.Dot(&v))
}

// |v|
func (v Vec2) Norm() float64 {
	return math.Sqrt(v.Dot(&v))
}

