package lib

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorResponse(code codes.Code, err error) error {
	return status.New(code, err.Error()).Err()
}
