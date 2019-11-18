package add

import "testing"

func TestAdd0And0Is0(t *testing.T) {
	if res := Add(0, 0); 0 != res {
		t.Fatalf("Expecting res to be 0 but was %d", res)
	}
}
