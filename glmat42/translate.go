package glmat42

func Translation4(x, y, z float64) (m *Mat4) {
	m = new(Mat4)

	m[ 0] = 1; m[ 1] = 0; m[ 2] = 0; m[ 3] = 0;
        m[ 4] = 0; m[ 5] = 1; m[ 6] = 0; m[ 7] = 0;
        m[ 8] = 0; m[ 9] = 0; m[10] = 1; m[11] = 0;
	m[12] = x; m[13] = y; m[14] = z; m[15] = 1;

	return m
}

func (m *Mat4) Translate(x, y, z float64) *Mat4 {
	m[12] += m[0]*x + m[4]*y + m[ 8]*z
	m[13] += m[1]*x + m[5]*y + m[ 9]*z
	m[14] += m[2]*x + m[6]*y + m[10]*z

	return m
}

func (m *Mat4) Translated(x, y, z float64) (t *Mat4) {
	t = new(Mat4)
	*t = *m
	t.Translate(x, y, z)
	return t
}

