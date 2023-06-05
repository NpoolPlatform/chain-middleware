package appcoin

import (
	"github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	coin.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	coin.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
