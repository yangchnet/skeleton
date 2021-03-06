// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"context"
	"github.com/yangchnet/skeleton/pkg/client/iam"
	"github.com/yangchnet/skeleton/pkg/token"
)

// Injectors from wire.go:

func InitService(ctx context.Context, secretKey string) (*Server, error) {
	iamServiceClient, err := iam.NewIamServiceClient(ctx)
	if err != nil {
		return nil, err
	}
	maker, err := token.NewJWTMaker(secretKey)
	if err != nil {
		return nil, err
	}
	server := NewServer(iamServiceClient, maker)
	return server, nil
}
