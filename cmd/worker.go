package cmd

import (
	"context"
	"notification_service/global"
	"sync"
)

func RunWorkerServer(
	ctx context.Context,
	wg *sync.WaitGroup,
	diContainer *global.DIContainer,
) {
	diContainer.StartWorker(ctx, wg)
}
