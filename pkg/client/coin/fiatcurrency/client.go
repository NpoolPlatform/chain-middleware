package fiatcurrency

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	fiatcurrencymgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fiatcurrency"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/fiatcurrency"

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

func CreateFiatCurrency(ctx context.Context, in *fiatcurrencymgrpb.FiatCurrencyReq) error {
	_, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateFiatCurrency(ctx, &npool.CreateFiatCurrencyRequest{
			Info: in,
		})
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

func CreateFiatCurrencies(ctx context.Context, in []*fiatcurrencymgrpb.FiatCurrencyReq) error {
	_, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateFiatCurrencies(ctx, &npool.CreateFiatCurrenciesRequest{
			Infos: in,
		})
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

func RefreshFiatCurrencies(ctx context.Context) error {
	_, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.RefreshFiatCurrencies(ctx, &npool.RefreshFiatCurrenciesRequest{})
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

func GetFiatCurrency(ctx context.Context, fiatCurrencyTypeID string) (*npool.FiatCurrency, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetFiatCurrency(ctx, &npool.GetFiatCurrencyRequest{
			FiatCurrencyTypeID: fiatCurrencyTypeID,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.FiatCurrency), nil
}

func GetFiatCurrencies(ctx context.Context, conds *npool.Conds) ([]*npool.FiatCurrency, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetFiatCurrencies(ctx, &npool.GetFiatCurrenciesRequest{
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
	return infos.([]*npool.FiatCurrency), nil
}

func GetCoinFiatCurrencies(ctx context.Context, coinTypeIDs, fiatCurrencyTypeIDs []string) ([]*npool.FiatCurrency, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetCoinFiatCurrencies(ctx, &npool.GetCoinFiatCurrenciesRequest{
			FiatCurrencyTypeIDs: fiatCurrencyTypeIDs,
			CoinTypeIDs:         coinTypeIDs,
		})
		if err != nil {
			return nil, err
		}

		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*npool.FiatCurrency), nil
}

func GetHistories(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.FiatCurrency, uint32, error) {
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
	return infos.([]*npool.FiatCurrency), total, nil
}
