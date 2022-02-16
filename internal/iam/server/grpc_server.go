package server

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yangchnet/skeleton/internal/iam/service"
	localGrpc "github.com/yangchnet/skeleton/pkg/grpc"
	pb "github.com/yangchnet/skeleton/pkg/pb/iam/v1"
	"google.golang.org/grpc"
)

func Serve(ctx context.Context) {
	conn, err := sql.Open("mysql", "mysql://root:secret@tcp(localhost:13306)/iam")
	if err != nil {
		log.Fatal(err)
	}
	s := service.InitService(context.Background(), conn)
	localGrpc.NewGrpcServer("iam", 10001).Serve(ctx, func(server *grpc.Server) {
		pb.RegisterIamServiceServer(server, &s)
	})
}
