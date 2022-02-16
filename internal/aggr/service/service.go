package service

import (
	pbHello "github.com/yangchnet/skeleton/pkg/pb/hello/v1"
	pbIam "github.com/yangchnet/skeleton/pkg/pb/iam/v1"
)

type service struct {
	iamServiceClient   pbIam.IamServiceClient
	helloServiceClient pbHello.HelloServiceClient
}
