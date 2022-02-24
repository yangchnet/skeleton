package apigateway

import (
	"context"
	"fmt"
	"net/http/httputil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yangchnet/skeleton/pkg/logger"
)

func (s *Server) log(ctx context.Context) gin.HandlerFunc {
	l := logger.StdLogger()

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
				logger.Error("Panic recovered: %+v\n%s", r, string(httpRequest))
				c.JSON(500, gin.H{
					"title": "error",
					"err":   r,
				})
			}
		}()
		c.Next()
	}
}
