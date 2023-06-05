package fiat

import (
	"github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	fiat.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	fiat.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
