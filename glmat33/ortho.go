package glmat33

// Create a 4 by 4 orthogonal projection matrix.
func Ortho4(left, right, bottom, top, nearVal, farVal float64) (m *Mat4) {
	tx := -(right + left)     / (right - left)
	ty := -(top + bottom)     / (top - bottom)
	tz := -(farVal + nearVal) / (farVal - nearVal)

	m = (*Mat4)(&[16]float64{2/(right-left),      0, 0, 0,
                                 0, 2/(top-bottom),      0, 0,
                                 0, 0, -2/(farVal-nearVal), 0,
	                         tx,      ty,      tz,      1})

	return m;
}

// Set m to a 4 by 4 orthogonal projection matrix.
func (m *Mat4) Ortho(left, right, bottom, top, nearVal, farVal float64) *Mat4{
	*m = *Ortho4(left, right, bottom, top, nearVal, farVal)

	return m
}

