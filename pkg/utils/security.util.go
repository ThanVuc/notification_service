package utils

import (
	"context"

	"github.com/thanvuc/go-core-lib/log"
	"go.uber.org/zap"
)

func WithSafePanic[TReq any, TResp any](
	ctx context.Context,
	logger log.Logger,
	req TReq,
	f func(context.Context, TReq) (TResp, error),
) (TResp, error) {
	requestId := GetRequestIDFromOutgoingContext(ctx)
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Recovered from panic",
				requestId,
				zap.Any("error", r),
			)
		}
	}()

	return f(ctx, req)
}

func WithSafePanicConsumer(ctx context.Context, logger log.Logger, f func(context.Context)) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Recovered from panic",
				"",
				zap.Any("error", r),
			)
		}
	}()

	f(ctx)
}

func WithSafePanicSimple(
	ctx context.Context,
	logger log.Logger,
	f func(context.Context) error,
) error {
	requestId := GetRequestIDFromOutgoingContext(ctx)
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Recovered from panic",
				requestId,
				zap.Any("error", r),
			)
		}
	}()

	return f(ctx)
}
