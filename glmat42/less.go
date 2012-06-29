package glmat

// Returns true if a.x < b.x. If a.x == b.x then return true if a.y < b.y.
// Similarily for z and w. This measure is only really useful when sorting
// vectors already known to be colinear.
func (a *Vec4) Less (b *Vec4) bool {
	return a[0] < b[0] || (a[0] == b[0] &&
	         (a[1] < b[1] || (a[1] == b[1] &&
		   (a[2] < b[2] || (a[2] == a[2] && a[3] < b[3])))))
}

// Returns true if a.x < b.x. If a.x == b.x then return true if a.y < b.y.
// Similarily for z. This measure is only really useful when sorting
// vectors already known to be colinear.
func (a *Vec3) Less (b *Vec3) bool {
	return a[0] < b[0] || (a[0] == b[0] &&
	         (a[1] < b[1] || (a[1] == b[1] && a[2] < b[2])))
}

// Returns true if a.x < b.x. If a.x == b.x then return true if a.y < b.y.
// This measure is only really useful when sorting vectors already known
// to be colinear.
func (a *Vec2) Less (b *Vec2) bool {
	return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
}

