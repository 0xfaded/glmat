package glmat33

import "math"

// Returns a Vec4 that is Less than all other Vec4
func MinV4() (v Vec4) {
	v[0] = math.Inf(-1)
	v[1] = math.Inf(-1)
	v[2] = math.Inf(-1)
	v[3] = math.Inf(-1)
	return
}

// Returns a Vec3 that is Less than all other Vec3
func MinV3() (v Vec3) {
	v[0] = math.Inf(-1)
	v[1] = math.Inf(-1)
	v[2] = math.Inf(-1)
	return
}

// Returns a Vec2 that is Less than all other Vec2
func MinV2() (v Vec2) {
	v[0] = math.Inf(-1)
	v[1] = math.Inf(-1)
	return
}

// Returns a Vec4 that is not Less than all other Vec4
func MaxV4() (v Vec4) {
	v[0] = math.Inf(1)
	v[1] = math.Inf(1)
	v[2] = math.Inf(1)
	v[3] = math.Inf(1)
	return
}

// Returns a Vec3 that is not Less than all other Vec3
func MaxV3() (v Vec3) {
	v[0] = math.Inf(1)
	v[1] = math.Inf(1)
	v[2] = math.Inf(1)
	return
}

// Returns a Vec2 that is not Less than all other Vec2
func MaxV2() (v Vec2) {
	v[0] = math.Inf(1)
	v[1] = math.Inf(1)
	return
}

