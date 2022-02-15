package apigateway

import (
	"context"
	"fmt"
	"log"
	"net/http/httputil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yangchnet/skeleton/pkg/logger"
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

	r.Any("/swagger-ui/*filepath", gin.WrapH(s.swaggerHandle(ctx)))

	return r.Run(":10000")
}

func (s *Server) log(ctx context.Context) gin.HandlerFunc {
	l := logger.NewLogger()
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		path := c.Request.URL.Path

		logStr := fmt.Sprintf("%3d | %v | %s | %s %s %s",
			statusCode,
			latency,
			clientIP, method,
			path,
			c.Errors.String(),
		)

		switch {
		case statusCode >= 400 && statusCode <= 499:
			l.Warn(ctx, logStr)
		case statusCode >= 500:
			l.Error(ctx, logStr)
		default:
			l.Info(ctx, logStr)
		}
	}
}

func (s *Server) recover(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				logger.Critical(ctx, "Panic recovered: %+v\n%s", r, string(httpRequest))
				c.JSON(500, gin.H{
					"title": "error",
					"err":   r,
				})
			}
		}()
		c.Next()
	}
}
