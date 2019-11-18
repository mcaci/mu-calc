package serv

import "context"

// Service provides and interface for subtracting integers
type Service interface {
	Sub(ctx context.Context, a, b int) (int, error)
}
