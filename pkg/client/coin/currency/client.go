package currency

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency"

	servicename "github.com/NpoolPlatform/chain-middleware/pkg/servicename"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
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

func GetCurrencies(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Currency, uint32, error) {
	total := uint32(0)

	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetCurrencies(ctx, &npool.GetCurrenciesRequest{
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

func GetCurrencyOnly(ctx context.Context, conds *npool.Conds) (*npool.Currency, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		const singleRowLimit = 2
		resp, err := cli.GetCurrencies(ctx, &npool.GetCurrenciesRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  singleRowLimit,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.Currency)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Currency)) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	return infos.([]*npool.Currency)[0], nil
}
