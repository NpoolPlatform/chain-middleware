package api

import (
	"context"

	chainmw "github.com/NpoolPlatform/message/npool/chain/mw/v1"

	"github.com/NpoolPlatform/chain-middleware/api/appcoin"
	"github.com/NpoolPlatform/chain-middleware/api/appcoin/description"
	"github.com/NpoolPlatform/chain-middleware/api/coin"
	currencyfeed "github.com/NpoolPlatform/chain-middleware/api/coin/currency/feed"
	currencyvalue "github.com/NpoolPlatform/chain-middleware/api/coin/currency/value"
	tran "github.com/NpoolPlatform/chain-middleware/api/tx"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	chainmw.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	chainmw.RegisterMiddlewareServer(server, &Server{})
	appcoin.Register(server)
	description.Register(server)
	coin.Register(server)
	tran.Register(server)
	currencyfeed.Register(server)
	currencyvalue.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := chainmw.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
