package service

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	v1 "github.com/yangchnet/skeleton/api/iam/v1"
	"github.com/yangchnet/skeleton/internal/api-gateway/conf"
	"github.com/yangchnet/skeleton/internal/api-gateway/middleware"
	"github.com/yangchnet/skeleton/pkg/token"
)

const (
	Authorization = "Authorization"
	RequestIdKey  = "X-Request-Id"
	xForwardedFor = "X-Forwarded-For"
)

var ProviderSet = wire.NewSet(NewServer)

// Server is a openapi server.
type Server struct {
	iamClient  v1.IamServiceClient
	tokenMaker token.Maker
}

// NewServer creates a new Server for openapi.
func NewServer(iamClient v1.IamServiceClient, tokenMaker token.Maker) *Server {
	return &Server{iamClient: iamClient, tokenMaker: tokenMaker}
}

// Serve start openapi server.
func Serve(ctx context.Context, bc *conf.Bootstrap) {
	s, _ := InitService(ctx, bc.Token.Secret) // FIXME: secret

	if err := s.run(ctx); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) run(ctx context.Context) error {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(middleware.Log(ctx), middleware.Recover(ctx))

	r.GET("/openapi-ui/*filepath", gin.WrapH(getOpenAPIHandler()))

	r.Any("/api/v1/*filepath", gin.WrapH(s.mainHandler(ctx)))

	return r.Run(":10000")
}
