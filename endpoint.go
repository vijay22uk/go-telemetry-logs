package ping

import (
	"context"
	"gokits/ping/model"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	SayHello endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {

	return Endpoints{
		SayHello: makeSayHelloEndpoint(s),
	}

}

func makeSayHelloEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.PingRequest)
		//TODO errr!
		id, msg, _ := s.SayHello(ctx, req.Name)
		return model.PingResponse{
			ID:   id,
			Pong: msg,
		}, nil
	}
}
