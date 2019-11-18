package serv

import "context"

// Service provides and interface for multiplying integers
type Service interface {
	Mul(ctx context.Context, a, b int) (int, error)
}
