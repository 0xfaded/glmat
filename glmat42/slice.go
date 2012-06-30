package glmat42

import (
	gl "github.com/chsc/gogl/gl42"
)

type (
	Vec4S []Vec4
	Vec3S []Vec3
	Vec2S []Vec2
)

// Returns number of vectors in the slice.
func (s Vec4S) Len() int {
	return len(s)
}

// Returns true if vector at position a is less than b. @see less.go.
func (s Vec4S) Less(a, b int) bool {
	return s[a].Less(&s[b])
}

// Swap vectors at position a and b.
func (s Vec4S) Swap(a, b int) {
	tmp := s[a]
	s[a] = s[b]
	s[b] = tmp
}

// Return size of array pointed to by s.Pointer().
func (s Vec4S) Sizeiptr() gl.Sizeiptr {
	if len(s) > 0 {
		return gl.Sizeiptr(int(s[0].Sizei()) * len(s))
	}
	return gl.Sizeiptr(0)
}

// Returns number of vectors in the slice.
func (s Vec3S) Len() int {
	return len(s)
}

// Returns true if vector at position a is less than b. @see less.go.
func (s Vec3S) Less(a, b int) bool {
	return s[a].Less(&s[b])
}

// Swap vectors at position a and b.
func (s Vec3S) Swap(a, b int) {
	tmp := s[a]
	s[a] = s[b]
	s[b] = tmp
}

// Return size of array pointed to by s.Pointer().
func (s Vec3S) Sizeiptr() gl.Sizeiptr {
	if len(s) > 0 {
		return gl.Sizeiptr(int(s[0].Sizei()) * len(s))
	}
	return gl.Sizeiptr(0)
}

// Returns number of vectors in the slice.
func (s Vec2S) Len() int {
	return len(s)
}

// Returns true if vector at position a is less than b. @see less.go.
func (s Vec2S) Less(a, b int) bool {
	return s[a].Less(&s[b])
}

// Swap vectors at position a and b.
func (s Vec2S) Swap(a, b int) {
	tmp := s[a]
	s[a] = s[b]
	s[b] = tmp
}

// Return size of array pointed to by s.Pointer().
func (s Vec2S) Sizeiptr() gl.Sizeiptr {
	if len(s) > 0 {
		return gl.Sizeiptr(int(s[0].Sizei()) * len(s))
	}
	return gl.Sizeiptr(0)
}

