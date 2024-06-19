package test10

import (
	"context"

	test10 "test10/api"
)

func (s *test10Server) Ping(ctx context.Context, request *test10.PingRequest) (*test10.PingRespond, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	return &test10.PingRespond{Message: request.GetMessage()}, nil
}
