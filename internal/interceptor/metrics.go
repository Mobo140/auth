package interceptor

import (
	"context"
	"time"

	"github.com/Mobo140/auth/internal/metric"
	"google.golang.org/grpc"
)

func MetricsInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	metric.IncRequestCounter()

	timeStart := time.Now()

	res, err := handler(ctx, req)
	diffTime := time.Since(timeStart)

	if err != nil {
		metric.IncResponseCounter("error", info.FullMethod)
		metric.HistogramResponseTimeObserve("error", diffTime.Seconds())
	} else {
		metric.IncResponseCounter("success", info.FullMethod)
		metric.HistogramResponseTimeObserve("success", diffTime.Seconds())
	}

	return res, err
}
