package apigateway

import (
	"context"
	"net/http"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// skipSpecialRoutes will skip the special routes without auth.
func (s *Server) skipSpecialRoutes(ctx context.Context, mux *gwruntime.ServeMux, w http.ResponseWriter, r *http.Request) bool {
	// switch r.URL.Path {
	// case "/v1/app/auth/token",
	// 	"/v1/app/evidences:verify",
	// 	"/v1/app/evidences:new_verify":
	// 	return true
	// case "/v1/internal/test/grpc_error":
	// 	return true
	// }
	return false
}
