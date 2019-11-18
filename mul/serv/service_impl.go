package serv

import (
	"context"

	"github.com/mcaci/mu-calc/mul/mul"
)

type mulService struct{}

// NewMulService makes a new Service
func NewMulService() Service {
	return mulService{}
}

func (mulService) Mul(ctx context.Context, a, b int) (int, error) {
	return mul.Mul(a, b), nil
}
