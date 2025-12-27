package cmd

import (
	"context"
	"notification_service/global"
	"sync"
)

func RunCronJobServer(
	ctx context.Context,
	wg *sync.WaitGroup,
	diContainer *global.DIContainer,
) {
	diContainer.StartCronJobs(ctx, wg)
}
