// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"context"
	"github.com/yangchnet/skeleton/pkg/client/hello"
	"github.com/yangchnet/skeleton/pkg/client/iam"
)

// Injectors from wire.go:

func InitService(ctx context.Context) (service, error) {
	iamServiceClient, err := iam.NewIamServiceClient(ctx)
	if err != nil {
		return service{}, err
	}
	helloServiceClient, err := hello.NewHelloServiceClient(ctx)
	if err != nil {
		return service{}, err
	}
	serviceService := service{
		iamServiceClient:   iamServiceClient,
		helloServiceClient: helloServiceClient,
	}
	return serviceService, nil
}