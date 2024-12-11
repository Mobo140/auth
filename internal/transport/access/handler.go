package access

import (
	"context"
	"errors"
	"strings"

	"github.com/Mobo140/auth/internal/service"
	desc "github.com/Mobo140/auth/pkg/access_v1"
	"github.com/Mobo140/platform_common/pkg/logger"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	authPrefix = "Bearer "
	traceIDKey = "x-trace-id"
)

type Implementation struct {
	desc.UnimplementedAccessV1Server
	accessService service.AccessService
}

func NewImplementation(accessService service.AccessService) *Implementation {
	return &Implementation{
		accessService: accessService,
	}
}

func (i *Implementation) Check(ctx context.Context, req *desc.CheckRequest) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Check/Implementation")
	defer span.Finish()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err := errors.New("metadata is not provided")
		logger.Error("Failed to get metadata from context: ",
			zap.Any("request", req),
			zap.Error(err),
		)

		return nil, err
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		err := errors.New("authorization header is not provided")
		logger.Error("Failed to get auth header from metadata: ",
			zap.Any("request", req),
			zap.Error(err),
		)

		return nil, err
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		err := errors.New("invalid authorization header format")
		logger.Error("Failed to get prefix Bearer: ",
			zap.Any("request", req),
			zap.Error(err),
		)

		return nil, err
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	err := i.accessService.Check(ctx, accessToken, req.GetEndpointAddress())
	if err != nil {
		logger.Error("Failed to check: ",
			zap.Any("endpoint", req.GetEndpointAddress()),
			zap.Any("access token", accessToken),
			zap.Error(err),
		)

		return nil, err
	}

	logger.Info("User successfully authenticated and granted access",
		zap.String("endpoint", req.GetEndpointAddress()),
	)

	return &emptypb.Empty{}, nil
}
