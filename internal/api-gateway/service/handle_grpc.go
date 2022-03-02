package service

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	echov1 "github.com/yangchnet/skeleton/api/echo/v1"
	iamv1 "github.com/yangchnet/skeleton/api/iam/v1"
	localGrpc "github.com/yangchnet/skeleton/pkg/grpc"
	"github.com/yangchnet/skeleton/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type registerWrapper struct {
	// Register<ServiceName>HandlerFromEndpoint
	f func(ctx context.Context, mux *gwruntime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	// endpoint like EchoService:10001
	endpoint string
}

// here to register service.
var registers []registerWrapper = []registerWrapper{
	{
		f:        echov1.RegisterEchoServiceHandlerFromEndpoint,
		endpoint: "127.0.0.1:10001", // FIXME
	},
	{
		f:        iamv1.RegisterIamServiceHandlerFromEndpoint,
		endpoint: "127.0.0.1:10003",
	},
}

func (s *Server) mainHandler(ctx context.Context) http.Handler {
	gwmux := gwruntime.NewServeMux(
		gwruntime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
			return metadata.Pairs( // FIXME
			// sender.SenderKey, req.Header.Get(sender.SenderKey),
			// RequestIdKey, req.Header.Get(RequestIdKey),
			)
		}))

	opts := localGrpc.ClientOptions

	for _, r := range registers {
		if err := r.f(ctx, gwmux, r.endpoint, opts); err != nil {
			logger.Error("error registers service to apigateway: ", err)
		}
	}

	mux := http.NewServeMux()
	mux.Handle("/", s.serveMuxHandle(ctx, gwmux))

	return s.formWrapper(mux)
}

// serveMuxHandle set skip special route and do auth.
func (s *Server) serveMuxHandle(ctx context.Context, mux *gwruntime.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if s.skipSpecialRoutes(ctx, mux, w, req) {
			// special route skip sender auth check
			mux.ServeHTTP(w, req)
		} else {
			if err := s.handleAuthorization(mux, w, req); err != nil {
				_, outboundMarshaler := gwruntime.MarshalerForRequest(mux, req)
				gwruntime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)

				return
			}
			mux.ServeHTTP(w, req)
		}
	})
}

func (s *Server) formWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)

				return
			}
			jsonMap := make(map[string]interface{}, len(r.Form))
			for k, v := range r.Form {
				if len(v) > 0 {
					jsonMap[k] = v[0]
				}
			}
			jsonBody, err := json.Marshal(jsonMap)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			r.Body = ioutil.NopCloser(bytes.NewReader(jsonBody))
			r.ContentLength = int64(len(jsonBody))
			r.Header.Set("Content-Type", "application/json")
		}

		h.ServeHTTP(w, r)
	})
}
