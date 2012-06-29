package glmat

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

// *gl.Float casts

func (m *Mat4) FloatP() (pointer *gl.Float) {
	pointer = (*gl.Float)(unsafe.Pointer(m))
	return pointer
}

func (v *Vec4) FloatP() (pointer *gl.Float) {
	pointer = (*gl.Float)(unsafe.Pointer(v))
	return pointer
}

func (v *Vec3) FloatP() (pointer *gl.Float) {
	pointer = (*gl.Float)(unsafe.Pointer(v))
	return pointer
}

func (v *Vec2) FloatP() (pointer *gl.Float) {
	pointer = (*gl.Float)(unsafe.Pointer(v))
	return pointer
}
