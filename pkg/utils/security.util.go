package utils

import (
	"context"

	"github.com/thanvuc/go-core-lib/log"
	"github.com/wagslane/go-rabbitmq"
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

func WithSafePanicEventBus(
	logger log.Logger,
	handler func(d rabbitmq.Delivery) rabbitmq.Action,
) func(d rabbitmq.Delivery) rabbitmq.Action {

	return func(d rabbitmq.Delivery) (action rabbitmq.Action) {
		defer func() {
			if r := recover(); r != nil {
				logger.Error(
					"panic recovered in consumer handler",
					"",
					zap.Any("panic", r),
					zap.Stack("stacktrace"),
				)
				// 🚨 panic = poison message
				action = rabbitmq.NackDiscard
			}
		}()

		return handler(d)
	}
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

func WithSafeWorkerPanic[TReq any, TResp any](
	ctx context.Context,
	logger log.Logger,
	req TReq,
	f func(context.Context, TReq) (TResp, error),
) (TResp, error) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Recovered from panic",
				"Worker panic",
				zap.Any("error", r),
			)
		}
	}()

	return f(ctx, req)
}
