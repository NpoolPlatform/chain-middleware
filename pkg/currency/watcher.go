//nolint:dupl
package currency

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	watcher "github.com/NpoolPlatform/go-service-framework/pkg/watcher"
	// cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	// basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	// "github.com/shopspring/decimal"
)

func refresh(ctx context.Context, fiat bool) {
	logger.Sugar().Infow(
		"refresh",
		"Fiat", fiat,
	)
}

var w *watcher.Watcher

func Watch(ctx context.Context) {
	refreshTicker := time.NewTicker(60 * time.Second)
	fiatRefreshCount := 0
	w = watcher.NewWatcher()

	for {
		select {
		case <-refreshTicker.C:
			refresh(ctx, fiatRefreshCount%60 == 0)
			fiatRefreshCount++
		case <-ctx.Done():
			logger.Sugar().Infow(
				"Watch",
				"State", "Done",
				"Error", ctx.Err(),
			)
			close(w.ClosedChan())
			return
		case <-w.CloseChan():
			close(w.ClosedChan())
			return
		}
	}
}

func Shutdown() {
	if w != nil {
		w.Shutdown()
	}
}
