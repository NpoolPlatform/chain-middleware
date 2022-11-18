package currencyvalue

import (
	"github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/value"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	value.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	value.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
