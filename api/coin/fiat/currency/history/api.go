package currencyhistory

import (
	"github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/fiat/currency/history"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	history.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	history.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
