package serv

import (
	"context"
	"testing"
)

func setup() (srv Service, ctx context.Context) {
	return NewMulService(), context.Background()
}

func TestMul(t *testing.T) {
	srv, ctx := setup()

	res, err := srv.Mul(ctx, 2, 3)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// testing status
	if ok := res == 6; !ok {
		t.Errorf("expected service to be ok")
	}
}
