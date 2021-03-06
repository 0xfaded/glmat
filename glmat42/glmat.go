package glmat42

import "fmt"

type Mat4 [16]float64

type Vec4 [4]float64
type Vec3 [3]float64
type Vec2 [2]float64

type Quat [4]float64

// Float32 types are not fully supported. They can be cast from float64 types.
type Mat4f [16]float32

type Vec4f [4]float32
type Vec3f [3]float32
type Vec2f [2]float32

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

func (q Quat) String() (s string) {
	s = fmt.Sprintf("[%4g (%4gi %4gj %4gk)]", q[0], q[1], q[2], q[3])
	return s
}

