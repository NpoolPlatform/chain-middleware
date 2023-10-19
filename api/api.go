package api

import (
	"context"

	chainmw "github.com/NpoolPlatform/message/npool/chain/mw/v1"

	appcoin "github.com/NpoolPlatform/chain-middleware/api/app/coin"
	"github.com/NpoolPlatform/chain-middleware/api/app/coin/description"
	chain "github.com/NpoolPlatform/chain-middleware/api/chain"
	"github.com/NpoolPlatform/chain-middleware/api/coin"
	coincurrency "github.com/NpoolPlatform/chain-middleware/api/coin/currency"
	coincurrencyfeed "github.com/NpoolPlatform/chain-middleware/api/coin/currency/feed"
	coincurrencyhis "github.com/NpoolPlatform/chain-middleware/api/coin/currency/history"
	coinfiat "github.com/NpoolPlatform/chain-middleware/api/coin/fiat"
	coinfiatcurrencyhis "github.com/NpoolPlatform/chain-middleware/api/coin/fiat/currency/history"
	"github.com/NpoolPlatform/chain-middleware/api/fiat"
	fiatcurrency "github.com/NpoolPlatform/chain-middleware/api/fiat/currency"
	fiatcurrencyfeed "github.com/NpoolPlatform/chain-middleware/api/fiat/currency/feed"
	fiatcurrencyhis "github.com/NpoolPlatform/chain-middleware/api/fiat/currency/history"
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
	coincurrency.Register(server)
	coincurrencyfeed.Register(server)
	coincurrencyhis.Register(server)
	coinfiat.Register(server)
	coinfiatcurrencyhis.Register(server)
	fiat.Register(server)
	fiatcurrency.Register(server)
	fiatcurrencyfeed.Register(server)
	fiatcurrencyhis.Register(server)
	tran.Register(server)
	chain.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := chainmw.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
