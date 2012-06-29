package glmat

import (
	"testing"
)

func TestMat4Mul(t *testing.T) {
	matA  := Mat4{  0,   1,   2,   3,
	                4,   5,   6,   7,
	                8,   9,  10,  11,
	               12,  13,  14,  15}

	matB  := Mat4{  0,  -1,  -2,  -3,
	                4,   3,   2,   1,
	                8,   7,   6,   5,
	               12,  11,  10,   9}

	matAB := Mat4{-56, -62, -68, -74,
                       40,  50,  60,  70,
                      136, 162, 188, 214,
                      232, 274, 316, 358}

	matBA := Mat4{ 56,  50,  44,  38,
                      152, 130, 108,  86,
                      248, 210, 172, 134,
                      344, 290, 236, 182}

	if (*matA.MulM(&matB) != matAB) {
		t.Logf("MulM A*B=AB failed.\n")
		t.Fail()
	}

	if (*matB.MulM(&matA) != matBA) {
		t.Logf("MulM B*A=BA failed.\n")
		t.Fail()
	}
}
