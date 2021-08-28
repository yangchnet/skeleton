package grpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func NewClient(host string, port int) (*grpc.ClientConn, error) {
	endpoint := fmt.Sprintf("%s:%d", host, port)
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, endpoint)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
