package currency

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	currencymgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreateCurrency(ctx context.Context, in *currencymgrpb.CurrencyReq) (*npool.Currency, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateCurrency(ctx, &npool.CreateCurrencyRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Currency), nil
}

func CreateCurrencies(ctx context.Context, in []*currencymgrpb.CurrencyReq) ([]*npool.Currency, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateCurrencies(ctx, &npool.CreateCurrenciesRequest{
			Infos: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*npool.Currency), nil
}

func RefreshCurrencies(ctx context.Context) error {
	_, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.RefreshCurrencies(ctx, &npool.RefreshCurrenciesRequest{})
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return err
	}
	return nil
}

func GetCurrency(ctx context.Context, id string) (*npool.Currency, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetCurrency(ctx, &npool.GetCurrencyRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Currency), nil
}

func GetCoinCurrency(ctx context.Context, coinTypeID string) (*npool.Currency, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetCoinCurrency(ctx, &npool.GetCoinCurrencyRequest{
			CoinTypeID: coinTypeID,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Currency), nil
}

func GetCurrencies(ctx context.Context, conds *npool.Conds) ([]*npool.Currency, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetCurrencies(ctx, &npool.GetCurrenciesRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}

		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*npool.Currency), nil
}

func GetHistories(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Currency, uint32, error) {
	var total uint32

	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetHistories(ctx, &npool.GetHistoriesRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}

		total = resp.Total

		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*npool.Currency), total, nil
}
