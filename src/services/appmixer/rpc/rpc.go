package rpc

import (
	"context"

	"todo.com/proto/appmixer"
)

type AppmixerServer struct {
	appmixer.AppmixerServer
}

func NewRPC() *AppmixerServer {
	return &AppmixerServer{}
}

func (s *AppmixerServer) SayHello(ctx context.Context, req *appmixer.AppRequest) (*appmixer.AppResponse, error) {
	return &appmixer.AppResponse{
		Message: "Hello World",
	}, nil
}
