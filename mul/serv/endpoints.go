package serv

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints are exposed
type Endpoints struct {
	MulEndpoint endpoint.Endpoint
}

// MakeMulEndpoint returns the response from our service "Mul"
func MakeMulEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(mulRequest)
		res, err := srv.Mul(ctx, req.A, req.B)
		if err != nil {
			return mulResponse{res, err.Error()}, nil
		}
		return mulResponse{res, ""}, nil
	}
}

// Mul endpoint mapping
func (e Endpoints) Mul(ctx context.Context, a, b int) (int, error) {
	req := mulRequest{A: a, B: b}
	resp, err := e.MulEndpoint(ctx, req)
	if err != nil {
		return 0, err
	}
	MulResp := resp.(mulResponse)
	if MulResp.Err != "" {
		return 0, errors.New(MulResp.Err)
	}
	return MulResp.Res, nil
}
