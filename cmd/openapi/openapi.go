package main

import (
	"context"

	openapi "github.com/yangchnet/skeleton/pkg/openapi"
)

func main() {
	ctx := context.Background()
	openapi.Serve(ctx)
}
