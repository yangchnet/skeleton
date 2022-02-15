package main

import (
	"context"

	apigateway "github.com/yangchnet/skeleton/pkg/api-gateway"
)

func main() {
	ctx := context.Background()

	// r := gin.New()

	// r.Use(gin.Logger(), gin.Recovery())

	// fs := http.FileServer(http.Dir("../../pkg/api-gateway/swagger-ui"))
	// http.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", fs))
	// if err := http.ListenAndServe(":10000", nil); err != nil {
	// 	log.Fatal("error occur")
	// }

	// r.Any("/swagger-ui/", gin.WrapH(http.StripPrefix("/swagger-ui/", fs)))

	// r.Run(":10000")

	apigateway.Serve(ctx)
}
