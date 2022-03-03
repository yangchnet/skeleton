package main

import (
	"context"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	flag "github.com/spf13/pflag"
	"github.com/yangchnet/skeleton/internal/api-gateway/conf"
	apigateway "github.com/yangchnet/skeleton/internal/api-gateway/service"
	"github.com/yangchnet/skeleton/pkg/logger"
)

var confLoc *string = flag.StringP("conf", "c", "configs/api-gateway/configs.yaml", "config file location")

func main() {
	flag.Parse()
	ctx := context.Background()
	bc := loadConfig(*confLoc)

	setLogger(bc)
	apigateway.Serve(ctx, bc)
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

func setLogger(c *conf.Bootstrap) {
	var formatter logger.Formatter
	switch c.Log.Formatter {
	case "text":
		formatter = &logger.TextFormatter{
			c.Log.IgnoreBasicFields,
		}
	case "json":
		formatter = &logger.JsonFormatter{
			c.Log.IgnoreBasicFields,
		}
	}
	logger.SetOptions(
		logger.WithLevel(logger.Level(c.Log.Level)),
		logger.WithFormatter(formatter),
		logger.WithDisableCaller(c.Log.DisableCaller),
		logger.WithDisableColor(c.Log.DisableColor),
		logger.WithEnableAbsPath(c.Log.EnableAbsPath),
		logger.WithEnableFunc(c.Log.EnableFunc),
	)
}
