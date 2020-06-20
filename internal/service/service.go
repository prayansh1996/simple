package service

import (
	"context"
	"errors"

	"github.com/prayansh1996/simple/proto/simple"
)

type SimpleService struct{}

func (s *SimpleService) SimpleRPC(ctx context.Context, request *simple.SimpleRequest) (*simple.SimpleResponse, error) {
	if request.SimpleString == "" {
		return nil, errors.New("Empty input string")
	}

	return &simple.SimpleResponse{
		SimpleString: "Your requested string is " + request.SimpleString,
	}, nil
}
