package glmat

func Translation4(x, y, z float32) (m *Mat4) {
	m = new(Mat4)

	m[ 0] = 1; m[ 1] = 0; m[ 2] = 0; m[ 3] = 0;
        m[ 4] = 0; m[ 5] = 1; m[ 6] = 0; m[ 7] = 0;
        m[ 8] = 0; m[ 9] = 0; m[10] = 1; m[11] = 0;
	m[12] = x; m[13] = y; m[14] = z; m[15] = 1;

	return m
}

func (m *Mat4) Translate(x, y, z float32) *Mat4 {
	m[12] += x * m[15]
	m[13] += y * m[15]
	m[14] += z * m[15]

	return m
}

