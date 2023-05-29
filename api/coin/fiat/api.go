package coinfiat

import (
	coinfiat1 "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/fiat"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	coinfiat1.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	coinfiat1.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
