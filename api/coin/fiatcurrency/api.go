package fiatcurrency

import (
	"github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/fiatcurrency"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	fiatcurrency.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	fiatcurrency.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
