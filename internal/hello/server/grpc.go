package server

import (
	"context"
	"database/sql"
	"log"

	"github.com/yangchnet/skeleton/internal/hello/service"
	localGrpc "github.com/yangchnet/skeleton/pkg/grpc"
	pb "github.com/yangchnet/skeleton/pkg/pb/basic"
	"google.golang.org/grpc"
)

func Serve(ctx context.Context) {
	conn, err := sql.Open("mysql", "mysql://root:secret@tcp(localhost:13306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	s := service.InitService(ctx, conn)
	localGrpc.NewGrpcServer("hello", 10002).Serve(ctx, func(server *grpc.Server) {
		pb.RegisterHelloServiceServer(server, &s)
	})
}
