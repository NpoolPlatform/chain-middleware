package chain

import (
	"github.com/NpoolPlatform/message/npool/chain/mw/v1/chain"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	chain.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	chain.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
