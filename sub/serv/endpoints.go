package serv

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints are exposed
type Endpoints struct {
	SubEndpoint endpoint.Endpoint
}

// MakeSubEndpoint returns the response from our service "Sub"
func MakeSubEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(subRequest)
		res, err := srv.Sub(ctx, req.A, req.B)
		if err != nil {
			return subResponse{res, err.Error()}, nil
		}
		return subResponse{res, ""}, nil
	}
}

// Sub endpoint mapping
func (e Endpoints) Sub(ctx context.Context, a, b int) (int, error) {
	req := subRequest{A: a, B: b}
	resp, err := e.SubEndpoint(ctx, req)
	if err != nil {
		return 0, err
	}
	SubResp := resp.(subResponse)
	if SubResp.Err != "" {
		return 0, errors.New(SubResp.Err)
	}
	return SubResp.Res, nil
}
