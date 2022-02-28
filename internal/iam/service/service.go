package service

import (
	v1 "github.com/yangchnet/skeleton/api/iam/v1"
	"github.com/yangchnet/skeleton/internal/iam/biz"
)

type service struct {
	v1.UnimplementedIamServiceServer

	uc *biz.IamCase
}

var _ v1.IamServiceServer = (*service)(nil)

func (s *service) mustEmbedUnimplementedIamServiceServer() {
	panic("not implemented") // TODO: Implement
}
