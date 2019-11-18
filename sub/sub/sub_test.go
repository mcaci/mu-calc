package sub

import "testing"

func TestSub0To0Is0(t *testing.T) {
	if res := Sub(0, 0); 0 != res {
		t.Fatalf("Expecting res to be 0 but was %d", res)
	}
}
