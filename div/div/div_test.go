package div

import "testing"

func TestDiv0And1Is0(t *testing.T) {
	if res := Div(0, 1); 0 != res {
		t.Fatalf("Expecting res to be 0 but was %d", res)
	}
}
