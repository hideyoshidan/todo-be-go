package rpc

import (
	"context"
	"errors"
	"log"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"todo.com/config"
	"todo.com/domain/repositories"
	"todo.com/lib"
	"todo.com/proto/statusmaster"
)

type StatusMasterServer struct {
	statusmaster.StatusMasterServer
	repository *repositories.MStatusRespository
}

func NewRPC(db *gorm.DB) *StatusMasterServer {
	return &StatusMasterServer{
		repository: repositories.NewMStatusRespository(db),
	}
}

func (s *StatusMasterServer) errorResponse(code codes.Code, err error) error {
	return status.New(code, err.Error()).Err()
}

// StatusList gets all status from Database
func (s *StatusMasterServer) StatusList(ctx context.Context, req *statusmaster.StatusRequest) (*statusmaster.StatusResponse, error) {
	log.Printf("StatusList is ENTRY")
	mStatuses, err := s.repository.FetchAll()
	if err != nil {
		slog.Error("Error occurred while fetching data. ERROR : %v", err)
		return nil, lib.ErrorResponse(codes.Internal, err)
	}

	// make response
	var response []*statusmaster.Status
	for _, r := range mStatuses {
		status := &statusmaster.Status{
			Id:        int32(r.ID),
			Name:      r.Name,
			ColorCode: r.ColorCode,
		}

		response = append(response, status)
	}

	// if no data, return error
	if len(response) < 1 {
		return nil, lib.ErrorResponse(codes.Internal, errors.New(config.ERROR_NO_DATA))
	}

	return &statusmaster.StatusResponse{
		Statuses: response,
	}, nil
}
