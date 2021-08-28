package service

import (
	pb "github.com/yangchnet/skeleton/pkg/pb/basic"
)

type service struct {
	iamServiceClient   pb.IamServiceClient
	helloServiceClient pb.HelloServiceClient
}
