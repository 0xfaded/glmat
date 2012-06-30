package glmat33

import "math"

func (v *Vec2) Arg() float32 {
	return float32(math.Atan2(float64(v[1]), float64(v[0])))
}

