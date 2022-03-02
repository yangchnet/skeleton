package service

import (
	"net/http"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// handleAuthorization 处理token并做鉴权.
func (s *Server) handleAuthorization(mux *gwruntime.ServeMux, w http.ResponseWriter, req *http.Request) error {
	// panic("authorization not implemented") // TODO: implement
	return nil
}
