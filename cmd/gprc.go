package cmd

import (
	"context"
	"notification_service/global"
	"sync"
)

func RunGrpcServer(
	ctx context.Context,
	wg *sync.WaitGroup,
	diContainer *global.DIContainer,
) {
	diContainer.StartGrpcServer(ctx, wg)
}
