package apigateway

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct{}

func Serve(ctx context.Context) {
	s, _ := InitService(ctx)

	if err := s.run(ctx); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) run(ctx context.Context) error {
	gin.SetMode(gin.DebugMode)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// r.Any("/swagger-ui/*filepath", gin.WrapH(s.swaggerHandle(ctx)))

	r.Any("/swagger-ui/*filepath", gin.WrapH(s.swaggerHandle(ctx)))

	return r.Run(":10000")
}
