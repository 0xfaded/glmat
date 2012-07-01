package glmat33

import "math"

func (v *Vec2) Arg() float64 {
	return math.Atan2(v[1], v[0])
}

