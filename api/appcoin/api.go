package appcoin

import (
	"github.com/NpoolPlatform/message/npool/chain/mw/v1/appcoin"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	appcoin.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	appcoin.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
