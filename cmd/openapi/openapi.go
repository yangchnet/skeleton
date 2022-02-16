package main

import (
	"context"

	openapi "github.com/yangchnet/skeleton/pkg/openapi"
)

func main() {
	openapi.Serve(context.Background())
}
