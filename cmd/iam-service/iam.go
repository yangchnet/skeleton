package main

import (
	"context"

	iam "github.com/yangchnet/skeleton/internal/iam/server"
)

func main() {
	ctx := context.Background()
	iam.Serve(ctx)
}
