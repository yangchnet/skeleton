package service

import (
	"context"

	v1 "github.com/yangchnet/skeleton/api/echo/v1"
)

var _ v1.EchoServiceServer = (*EchoService)(nil)

func (s *EchoService) CreateEcho(_ context.Context, _ *v1.CreateEchoRequest) (*v1.CreateEchoResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *EchoService) ListEcho(_ context.Context, _ *v1.ListEchoRequest) (*v1.ListEchoResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *EchoService) UpdateEcho(_ context.Context, _ *v1.UpdateEchoRequest) (*v1.UpdateEchoResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *EchoService) DeleteEcho(_ context.Context, _ *v1.DeleteEchoRequest) (*v1.DeleteEchoResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *EchoService) mustEmbedUnimplementedEchoServiceServer() {
	panic("not implemented") // TODO: Implement
}
