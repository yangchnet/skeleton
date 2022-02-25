package service

import (
	"context"
	"errors"

	v1 "github.com/yangchnet/skeleton/api/echo/v1"
	"github.com/yangchnet/skeleton/internal/echo/biz"
	"github.com/yangchnet/skeleton/tools/pbtools"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ v1.EchoServiceServer = (*EchoService)(nil)

func (s *EchoService) CreateEcho(ctx context.Context, req *v1.CreateEchoRequest) (*v1.CreateEchoResponse, error) {
	message := req.GetMessage().GetValue()

	echo, err := s.uc.CreateEcho(ctx, message)
	if err != nil {
		return nil, err
	}

	return &v1.CreateEchoResponse{
		Echo: pbtools.ToProtoString(echo.EchoMessage),
	}, nil
}

func (s *EchoService) ListEcho(ctx context.Context, req *v1.ListEchoRequest) (*v1.ListEchoResponse, error) {
	offset := req.GetOffset()
	limit := req.GetLimit()

	echoList, err := s.uc.ListEcho(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	list := make([]*v1.EchoRecord, len(echoList))
	for i, echo := range echoList {
		list[i] = &v1.EchoRecord{
			Id:       pbtools.ToProtoInt64(int64(echo.ID)),
			Message:  pbtools.ToProtoString(echo.Message),
			Echo:     pbtools.ToProtoString(echo.EchoMessage),
			EchoTime: pbtools.ToProtoTimestamp(*echo.CreateTime),
		}
	}

	return &v1.ListEchoResponse{
		Records:    list,
		ListLength: int32(len(echoList)),
	}, nil
}

func (s *EchoService) UpdateEcho(ctx context.Context, req *v1.UpdateEchoRequest) (*v1.UpdateEchoResponse, error) {
	id := req.GetId().GetValue()
	message := req.GetNewEcho().GetValue()

	_, err := s.uc.UpdateEcho(ctx, id, message)
	if err != nil {
		if errors.Is(err, biz.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "id not found")
		}
	}

	return &v1.UpdateEchoResponse{
		Result: v1.OperationResult_SUCCESS,
	}, nil
}

func (s *EchoService) DeleteEcho(ctx context.Context, req *v1.DeleteEchoRequest) (*v1.DeleteEchoResponse, error) {
	if err := s.uc.DeleteEcho(ctx, req.GetId().GetValue()); err != nil {
		return &v1.DeleteEchoResponse{
			Result: v1.OperationResult_FAIL,
		}, err
	}

	return &v1.DeleteEchoResponse{
		Result: v1.OperationResult_SUCCESS,
	}, nil
}

func (s *EchoService) mustEmbedUnimplementedEchoServiceServer() {
	panic("not implemented") // TODO: Implement
}
