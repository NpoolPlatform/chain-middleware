package currency

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	watcher "github.com/NpoolPlatform/go-service-framework/pkg/watcher"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
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
	lockKey := fmt.Sprintf("%v", basetypes.Prefix_PrefixUpdateCoinCurrency)
	for {
		if err := redis2.TryLock(lockKey, 0); err == nil {
			break
		}
		<-time.After(time.Minute)
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

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

func Shutdown(ctx context.Context) {
	if w != nil {
		w.Shutdown(ctx)
	}
}
