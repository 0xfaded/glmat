package glmat42

import (
	gl "github.com/chsc/gogl/gl42"
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

func (mf *Mat4f) Sizei() (sizeof gl.Sizei) {
	sizeof = gl.Sizei(unsafe.Sizeof(*mf))
	return sizeof
}

func (vf *Vec4f) Sizei() (sizeof gl.Sizei) {
	sizeof = gl.Sizei(unsafe.Sizeof(*vf))
	return sizeof
}

func (vf *Vec3f) Sizei() (sizeof gl.Sizei) {
	sizeof = gl.Sizei(unsafe.Sizeof(*vf))
	return sizeof
}

func (vf *Vec2f) Sizei() (sizeof gl.Sizei) {
	sizeof = gl.Sizei(unsafe.Sizeof(*vf))
	return sizeof
}

