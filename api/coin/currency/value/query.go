package currencyvalue

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/value"
)

func (s *Server) GetCurrency(ctx context.Context, in *npool.GetCurrencyRequest) (*npool.GetCurrencyResponse, error) {
	return nil, nil
}

func (s *Server) GetCoinCurrency(ctx context.Context, in *npool.GetCoinCurrencyRequest) (*npool.GetCoinCurrencyResponse, error) {
	return nil, nil
}

func (s *Server) GetCurrencies(ctx context.Context, in *npool.GetCurrenciesRequest) (*npool.GetCurrenciesResponse, error) {
	return nil, nil
}

func (s *Server) GetHistories(ctx context.Context, in *npool.GetHistoriesRequest) (*npool.GetHistoriesResponse, error) {
	return nil, nil
}
