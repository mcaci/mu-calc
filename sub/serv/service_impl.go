package serv

import (
	"context"

	"github.com/mcaci/mu-calc/sub/sub"
)

type subService struct{}

// NewSubService makes a new Service
func NewSubService() Service {
	return subService{}
}

func (subService) Sub(ctx context.Context, a, b int) (int, error) {
	return sub.Sub(a, b), nil
}
