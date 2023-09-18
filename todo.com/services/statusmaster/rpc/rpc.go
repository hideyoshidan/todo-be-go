package rpc

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"todo.com/proto/statusmaster"
)

type StatusMasterServer struct {
	statusmaster.StatusMasterServer
	db *gorm.DB
}

func NewRPC(db *gorm.DB) *StatusMasterServer {
	return &StatusMasterServer{
		db: db,
	}
}

func (s *StatusMasterServer) StatusList(ctx context.Context, req *statusmaster.StatusRequest) (*statusmaster.StatusResponse, error) {
	var statuses []*statusmaster.Status
	for i := 0; i < 3; i++ {
		status := &statusmaster.Status{
			Id:        int32(i),
			Name:      fmt.Sprintf("%d 番目", i),
			ColorCode: fmt.Sprintf("%dFFFFF", i),
		}

		statuses = append(statuses, status)
	}
	return &statusmaster.StatusResponse{
		Statuses: statuses,
	}, nil
}
