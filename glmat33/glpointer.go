package glmat33

import (
	gl "github.com/chsc/gogl/gl33"
	"unsafe"
)

// gl.Pointer casts

func (m *Mat4) Pointer() (pointer gl.Pointer) {
	pointer = gl.Pointer(unsafe.Pointer(m))
	return pointer
}

func (v *Vec4) Pointer() (pointer gl.Pointer) {
	pointer = gl.Pointer(unsafe.Pointer(v))
	return pointer
}

func (v *Vec3) Pointer() (pointer gl.Pointer) {
	pointer = gl.Pointer(unsafe.Pointer(v))
	return pointer
}

func (v *Vec2) Pointer() (pointer gl.Pointer) {
	pointer = gl.Pointer(unsafe.Pointer(v))
	return pointer
}

// *gl.Double casts

func (m *Mat4) DoubleP() (pointer *gl.Double) {
	pointer = (*gl.Double)(unsafe.Pointer(m))
	return pointer
}

func (v *Vec4) DoubleP() (pointer *gl.Double) {
	pointer = (*gl.Double)(unsafe.Pointer(v))
	return pointer
}

func (v *Vec3) DoubleP() (pointer *gl.Double) {
	pointer = (*gl.Double)(unsafe.Pointer(v))
	return pointer
}

func (v *Vec2) DoubleP() (pointer *gl.Double) {
	pointer = (*gl.Double)(unsafe.Pointer(v))
	return pointer
}
