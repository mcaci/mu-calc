package serv

import (
	"context"

	"github.com/mcaci/calc/add/add"
)

type addService struct{}

// NewAddService makes a new Service
func NewAddService() Service {
	return addService{}
}

func (addService) Add(ctx context.Context, a, b int) (int, error) {
	return add.Add(a, b), nil
}
