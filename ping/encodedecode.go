package ping

import (
	"context"
	"encoding/json"
	"gokits/ping/model"
	"net/http"
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.PingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
