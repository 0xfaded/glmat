package glmat33

import (
	"testing"
)

func TestScaleAngleIdentity(t *testing.T) {

	q0 := IdentityQ()
	q1 := q0.ScaleAngle(5)

	if (q0 != q1) {
		t.Logf("Idenity Scale Broke.\n")
		t.Fail()
	}
}

