package serv

import (
	"context"
	"testing"
)

func setup() (srv Service, ctx context.Context) {
	return NewDivService(), context.Background()
}

func TestDiv(t *testing.T) {
	srv, ctx := setup()

	res, err := srv.Div(ctx, 2, 3)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// testing status
	if ok := res == 0; !ok {
		t.Errorf("expected service to be ok")
	}
}
