package service

import (
	"github.com/google/wire"
	v1 "github.com/yangchnet/skeleton/api/iam/v1"
	"github.com/yangchnet/skeleton/internal/iam/biz"
)

var ProviderSet = wire.NewSet(NewIamService)

type service struct {
	v1.UnimplementedIamServiceServer

	uc *biz.IamCase
}

var _ v1.IamServiceServer = (*service)(nil)

func (s *service) mustEmbedUnimplementedIamServiceServer() {
	panic("not implemented") // TODO: Implement
}

func NewIamService(uc *biz.IamCase) *service {
	return &service{uc: uc}
}
