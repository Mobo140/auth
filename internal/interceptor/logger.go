package interceptor

import (
	"context"
	"time"

	"github.com/Mobo140/platform_common/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func LogInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	now := time.Now()

	res, err := handler(ctx, req)
	if err != nil {
		logger.Error(err.Error(), zap.String("method", info.FullMethod), zap.Any("req", req))
	}

	logger.Info("request",
		zap.String("method", info.FullMethod),
		zap.Any("req", req),
		zap.Duration("duration",
			time.Since(now)),
	)

	return res, err
}
