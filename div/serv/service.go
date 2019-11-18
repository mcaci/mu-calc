package serv

import "context"

// Service provides and interface for dividing integers
type Service interface {
	Div(ctx context.Context, a, b int) (int, error)
}
