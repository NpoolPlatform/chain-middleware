package currencyfeed

import (
	feed "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/feed"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	feed.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	feed.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
