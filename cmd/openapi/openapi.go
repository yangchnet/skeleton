package main

import (
	"context"

	apigateway "github.com/yangchnet/skeleton/pkg/api-gateway"
)

func main() {
	ctx := context.Background()
	apigateway.Serve(ctx, "326c9f73b0f45705a73eb020a32d790a") // FIXME: secretKey
}
