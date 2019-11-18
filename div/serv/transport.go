package serv

import (
	"context"
	"encoding/json"
	"net/http"
)

// In the first part of the file we are mapping requests and responses to their JSON payload.
type divRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type divResponse struct {
	Res int    `json:"res"`
	Err string `json:"err,omitempty"`
}

// In the second part we will write "decoders" for our incoming requests
func decodeDivRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req divRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Last but not least, we have the encoder for the response output
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
