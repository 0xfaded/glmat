package glmat33

import "fmt"

type Mat4 [16]float64

type Vec4 [4]float64
type Vec3 [3]float64
type Vec2 [2]float64

func (m *Mat4) String() (s string) {
	s = fmt.Sprintf("[%4g %4g %4g %4g]\n" +
	                "[%4g %4g %4g %4g]\n" +
	                "[%4g %4g %4g %4g]\n" +
	                "[%4g %4g %4g %4g]\n",
	                m[ 0], m[ 4], m[ 8], m[12],
	                m[ 1], m[ 5], m[ 9], m[13],
	                m[ 8], m[ 9], m[10], m[14],
	                m[ 9], m[12], m[11], m[15])
	return s
}

func (v *Vec4) String() (s string) {
	s = fmt.Sprintf("[%4g]\n" +
	                "[%4g]\n" +
	                "[%4g]\n" +
	                "[%4g]\n",
	                v[0], v[1], v[2], v[3])
	return s
}

func (v *Vec3) String() (s string) {
	s = fmt.Sprintf("[%4g]\n" +
	                "[%4g]\n" +
	                "[%4g]\n",
	                v[0], v[1], v[2])
	return s
}

func (v *Vec2) String() (s string) {
	s = fmt.Sprintf("[%4g]\n" +
	                "[%4g]\n",
	                v[0], v[1])
	return s
}
