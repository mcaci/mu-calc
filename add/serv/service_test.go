package serv

import (
	"context"
	"testing"
)

func setup() (srv Service, ctx context.Context) {
	return NewAddService(), context.Background()
}

func TestAdd(t *testing.T) {
	srv, ctx := setup()

	res, err := srv.Add(ctx, 2, 3)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// testing status
	if ok := res == 5; !ok {
		t.Errorf("expected service to be ok")
	}
}
