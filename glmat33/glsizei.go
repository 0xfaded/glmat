package glmat33

import (
	gl "github.com/chsc/gogl/gl33"
	"unsafe"
)

// gl sizeof

func (m *Mat4) Sizei() (sizeof gl.Sizei) {
	sizeof = gl.Sizei(unsafe.Sizeof(*m))
	return sizeof
}

func (v *Vec4) Sizei() (sizeof gl.Sizei) {
	sizeof = gl.Sizei(unsafe.Sizeof(*v))
	return sizeof
}

func (v *Vec3) Sizei() (sizeof gl.Sizei) {
	sizeof = gl.Sizei(unsafe.Sizeof(*v))
	return sizeof
}

func (v *Vec2) Sizei() (sizeof gl.Sizei) {
	sizeof = gl.Sizei(unsafe.Sizeof(*v))
	return sizeof
}

