package serv

import "context"

// Service provides and interface for adding integers
type Service interface {
	Add(ctx context.Context, a,b int) (int, error)
}
