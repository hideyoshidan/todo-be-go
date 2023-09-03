package rpc

import (
	"context"
	"fmt"

	"todo.com/proto/statusmaster"
)

type StatusMasterServer struct {
	statusmaster.StatusMasterServer
}

func NewRPC() *StatusMasterServer {
	return &StatusMasterServer{}
}

func (s *StatusMasterServer) SayHello(ctx context.Context, req *statusmaster.StausRequest) (*statusmaster.StausResponse, error) {
	return &statusmaster.StausResponse{
		Message: fmt.Sprintf("Hello %v From Status Master!!!!", req.Name),
	}, nil
}
