package cmd

import (
	"context"
	"notification_service/global"
	"sync"
)

func RunConsumer(
	ctx context.Context,
	wg *sync.WaitGroup,
	diContainer *global.DIContainer,
) {
	diContainer.StartComsumerWorkers(ctx, wg)
}
