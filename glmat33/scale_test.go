package glmat33

import (
	"testing"
)

func TestScale(t *testing.T) {
	mat  := Mat4{2, 4, 8, 1,
	             2, 4, 8, 1,
	             2, 4, 8, 1,
	             2, 4, 8, 1}

	ones := Mat4{1, 1, 1, 1,
	             1, 1, 1, 1,
	             1, 1, 1, 1,
	             1, 1, 1, 1}

	if (*mat.ScaledXYZ(0.5, 0.25, 0.125) != ones) {
		t.Logf("Scale failed.\n")
		t.Fail()
	}
}

