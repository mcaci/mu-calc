package serv

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints are exposed
type Endpoints struct {
	DivEndpoint endpoint.Endpoint
}

// MakeDivEndpoint returns the response from our service "Div"
func MakeDivEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(divRequest)
		res, err := srv.Div(ctx, req.A, req.B)
		if err != nil {
			return divResponse{res, err.Error()}, nil
		}
		return divResponse{res, ""}, nil
	}
}

// Div endpoint mapping
func (e Endpoints) Div(ctx context.Context, a, b int) (int, error) {
	req := divRequest{A: a, B: b}
	resp, err := e.DivEndpoint(ctx, req)
	if err != nil {
		return 0, err
	}
	DivResp := resp.(divResponse)
	if DivResp.Err != "" {
		return 0, errors.New(DivResp.Err)
	}
	return DivResp.Res, nil
}
