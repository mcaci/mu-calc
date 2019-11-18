package serv

import (
	"context"
	"testing"
)

func setup() (srv Service, ctx context.Context) {
	return NewSubService(), context.Background()
}

func TestSub(t *testing.T) {
	srv, ctx := setup()

	res, err := srv.Sub(ctx, 2, 3)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// testing status
	if ok := res == -1; !ok {
		t.Errorf("expected service to be ok")
	}
}
