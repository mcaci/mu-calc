package serv

import (
	"context"

	"github.com/mcaci/calc/div/div"
)

type divService struct{}

// NewDivService makes a new Service
func NewDivService() Service {
	return divService{}
}

func (divService) Div(ctx context.Context, a, b int) (int, error) {
	return div.Div(a, b), nil
}
