package glmat33

func Identity4() (m *Mat4) {
	m = (*Mat4)(&[16]float64{1.0, 0.0, 0.0, 0.0,
	                         0.0, 1.0, 0.0, 0.0,
	                         0.0, 0.0, 1.0, 0.0,
	                         0.0, 0.0, 0.0, 1.0})
	return m
}

func (m *Mat4) Identity() *Mat4 {
	*m = [16]float64{1.0, 0.0, 0.0, 0.0,
	                 0.0, 1.0, 0.0, 0.0,
	                 0.0, 0.0, 1.0, 0.0,
	                 0.0, 0.0, 0.0, 1.0}
	return m
}

