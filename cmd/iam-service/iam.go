package main

import (
	"context"

	iam "github.com/yangchnet/skeleton/internal/iam/server"
)

func main() { // BUG
	ctx := context.Background()
	iam.Serve(ctx)
}
