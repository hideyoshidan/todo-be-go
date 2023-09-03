package rpc

import (
	"context"

	"todo.com/proto/appmixer"
	"todo.com/proto/statusmaster"
)

type Appmixer struct {
	sClient statusmaster.StatusMasterClient
	appmixer.AppmixerServer
}

func NewRPC(sClient statusmaster.StatusMasterClient) *Appmixer {
	return &Appmixer{
		sClient: sClient,
	}
}

func (s *Appmixer) SayHello(ctx context.Context, req *appmixer.AppRequest) (*appmixer.AppResponse, error) {
	res, err := s.sClient.SayHello(ctx, &statusmaster.StausRequest{
		Name: req.Name,
	})

	if err != nil {
		return nil, err
	}

	return &appmixer.AppResponse{
		Message: res.Message,
	}, nil
}
