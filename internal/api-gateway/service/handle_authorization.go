package service

import (
	"context"
	"net/http"
	"strings"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/marmotedu/component-base/pkg/util/stringutil"
	"github.com/yangchnet/skeleton/pkg/logger"
	"github.com/yangchnet/skeleton/pkg/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// handleAuthorization 处理token并做鉴权.
func (s *Server) handleAuthorization(mux *gwruntime.ServeMux, w http.ResponseWriter, req *http.Request) error {
	var err error
	ctx := req.Context()

	auth := strings.SplitN(req.Header.Get("Authorization"), " ", 2)
	if stringutil.FindString([]string{"Bearer"}, auth[0]) < 0 {
		return status.Errorf(codes.InvalidArgument, "http header `Authorization` format error")
	}

	payload, err := s.handleBearer(ctx, auth[1])
	if err != nil {
		return status.Error(codes.Internal, "")
	}

	req.Header.Set("username", payload.Username)
	req.Header.Del(Authorization)

	return nil
}

func (s *Server) handleBearer(ctx context.Context, token string) (*token.Payload, error) {
	payload, err := s.tokenMaker.VerifyToken(token)
	if err != nil {
		logger.Error("verify token failed;", err)
		return nil, err
	}

	return payload, nil
}
