package main

import (
	"context"

	apigateway "github.com/yangchnet/skeleton/pkg/api-gateway"
)

func main() {
	ctx := context.Background()
	apigateway.Serve(ctx)
}
