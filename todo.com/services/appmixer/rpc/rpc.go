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
	return &appmixer.AppResponse{
		Message: "Hello!!!",
	}, nil
}

func (s *Appmixer) GetStatus(ctx context.Context, req *statusmaster.StatusRequest) (*statusmaster.StatusResponse, error) {
	res, err := s.sClient.StatusList(ctx, &statusmaster.StatusRequest{})

	if err != nil {
		return nil, err
	}
	return res, nil

	// return res, nil
}
