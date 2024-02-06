package coinusedfor

import (
	coinusedfor1 "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/usedfor"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	coinusedfor1.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	coinusedfor1.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
