package serv

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints are exposed
type Endpoints struct {
	AddEndpoint endpoint.Endpoint
}

// MakeAddEndpoint returns the response from our service "Add"
func MakeAddEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addRequest)
		res, err := srv.Add(ctx, req.A, req.B)
		if err != nil {
			return addResponse{res, err.Error()}, nil
		}
		return addResponse{res, ""}, nil
	}
}

// Add endpoint mapping
func (e Endpoints) Add(ctx context.Context, a, b int) (int, error) {
	req := addRequest{A: a, B: b}
	resp, err := e.AddEndpoint(ctx, req)
	if err != nil {
		return 0, err
	}
	AddResp := resp.(addResponse)
	if AddResp.Err != "" {
		return 0, errors.New(AddResp.Err)
	}
	return AddResp.Res, nil
}
