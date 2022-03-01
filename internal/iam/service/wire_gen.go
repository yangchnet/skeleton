// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"context"
	"github.com/yangchnet/skeleton/internal/iam/biz"
	"github.com/yangchnet/skeleton/internal/iam/conf"
	"github.com/yangchnet/skeleton/internal/iam/data"
)

// Injectors from wire.go:

func InitService(ctx context.Context, conf2 *conf.Bootstrap) (*service, error) {
	client, err := data.NewEnt(conf2)
	if err != nil {
		return nil, err
	}
	cacheInterface := data.NewCache(conf2)
	iamRepo := data.NewData(client, cacheInterface)
	iamCase := biz.NewIamCase(iamRepo)
	serviceService := NewIamService(iamCase)
	return serviceService, nil
}