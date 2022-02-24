package apigateway

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	Authorization = "Authorization"
	RequestIdKey  = "X-Request-Id"
	xForwardedFor = "X-Forwarded-For"
)

type Server struct{}

func Serve(ctx context.Context) {
	s, _ := InitService(ctx)

	if err := s.run(ctx); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) run(ctx context.Context) error {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(s.log(ctx), s.recover(ctx))
	r.GET("/openapi-ui/*filepath", gin.WrapH(getOpenAPIHandler()))

	r.Any("/api/v1/*filepath", gin.WrapH(s.mainHandler(ctx)))

	return r.Run(":10000")
}
