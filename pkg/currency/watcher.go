package currency

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	watcher "github.com/NpoolPlatform/go-service-framework/pkg/watcher"
)

func refresh(ctx context.Context, fiat bool) {
	refreshCoins(ctx)
	refreshCoinFiats(ctx)
	if fiat {
		refreshFiats(ctx)
	}
}

var w *watcher.Watcher

func Watch(ctx context.Context) {
	refreshTicker := time.NewTicker(60 * time.Second) //nolint
	fiatRefreshCount := 0
	w = watcher.NewWatcher()

	refresh(ctx, true)

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
