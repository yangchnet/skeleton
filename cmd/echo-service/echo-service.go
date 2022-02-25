package main

import (
	"context"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	flag "github.com/spf13/pflag"
	"github.com/yangchnet/skeleton/internal/echo/conf"
	"github.com/yangchnet/skeleton/internal/echo/server"
	"github.com/yangchnet/skeleton/pkg/logger"
)

var confLoc *string = flag.StringP("conf", "c", "configs/echo/config.yaml", "config file location")

func main() {
	flag.Parse()

	bc := loadConfig(*confLoc)

	server.GrpcServe(context.Background(), bc)
}

func loadConfig(confLoc string) *conf.Bootstrap {
	c := config.New(
		config.WithSource(
			file.NewSource(confLoc),
		),
		config.WithLogger(logger.StdLogger()),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		logger.Panicf("Error loading config file: %v", err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		logger.Panicf("Error scanning config file: %v", err)
	}

	return &bc
}
