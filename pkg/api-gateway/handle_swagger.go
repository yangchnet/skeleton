package apigateway

import (
	"context"
	"net/http"

	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/httpfs"
	"golang.org/x/tools/godoc/vfs/mapfs"

	spec "github.com/yangchnet/skeleton/pkg/api-gateway/spec"
	swaggerUI "github.com/yangchnet/skeleton/pkg/api-gateway/swagger-ui"
)

func (s *Server) swaggerHandle(ctx context.Context) http.Handler {
	ns := vfs.NameSpace{}
	ns.Bind("/", mapfs.New(swaggerUI.Files), "/", vfs.BindReplace)
	ns.Bind("/", mapfs.New(spec.Files), "/", vfs.BindBefore)
	return http.StripPrefix("/swagger-ui/", http.FileServer(httpfs.New(ns)))
}
