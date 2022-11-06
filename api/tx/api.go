package tran

import (
	tran1 "github.com/NpoolPlatform/message/npool/chain/mw/v1/tx"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	tran1.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	tran1.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
