package appcoin

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/appcoin"
)

func trace(span trace1.Span, in *npool.CoinReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("CoinTypeID.%v", index), in.GetCoinTypeID()),
		attribute.Bool(fmt.Sprintf("ForPay.%v", index), in.GetForPay()),
		attribute.String(fmt.Sprintf("WithdrawAutoReviewAmount.%v", index), in.GetWithdrawAutoReviewAmount()),
		attribute.String(fmt.Sprintf("MarketValue.%v", index), in.GetMarketValue()),
		attribute.Int(fmt.Sprintf("SettlePercent.%v", index), int(in.GetSettlePercent())),
		attribute.String(fmt.Sprintf("Setter.%v", index), in.GetSetter()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.CoinReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Value", in.GetAppID().GetValue()),
		attribute.String("CoinTypeID.Op", in.GetCoinTypeID().GetOp()),
		attribute.String("CoinTypeID.Value", in.GetCoinTypeID().GetValue()),
		attribute.String("ForPay.Op", in.GetForPay().GetOp()),
		attribute.Bool("ForPay.Value", in.GetForPay().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.CoinReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
