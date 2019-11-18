package mul

import "testing"

func TestMul0And0Is0(t *testing.T) {
	if res := Mul(0, 0); 0 != res {
		t.Fatalf("Expecting res to be 0 but was %d", res)
	}
}
