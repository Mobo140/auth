package auth

import (
	"context"

	conv "github.com/Mobo140/auth/internal/converter/auth"
	"github.com/Mobo140/auth/internal/service"
	desc "github.com/Mobo140/auth/pkg/auth_v1"
	"github.com/Mobo140/platform_common/pkg/logger"
	"go.uber.org/zap"
)

type Implementation struct {
	desc.UnimplementedAuthV1Server
	authService service.AuthService
}

func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}

func (i *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	data := conv.ToLoginDataFromDesc(req)

	refreshTokenData, err := i.authService.Login(ctx, data)
	if err != nil {
		logger.Error("Failed to login: ", zap.Error(err))

		return nil, err
	}

	logger.Info("User authorized", zap.Any("login data", data))

	return &desc.LoginResponse{
		RefreshToken: *refreshTokenData,
	}, nil
}

func (i *Implementation) GetRefreshToken(ctx context.Context,
	req *desc.GetRefreshTokenRequest,
) (*desc.GetRefreshTokenResponse, error) {
	refreshTokenStr := req.GetRefreshToken()

	refreshToken, err := i.authService.GetRefreshToken(ctx, refreshTokenStr)
	if err != nil {
		logger.Error("Failed to get new refresh token by old: ",
			zap.Any("old refresh token", refreshTokenStr),
			zap.Error(err),
		)

		return nil, err
	}

	logger.Info("Get new refresh token: ", zap.Any("refresh token", refreshToken))

	return &desc.GetRefreshTokenResponse{
		RefreshToken: *refreshToken,
	}, nil
}

func (i *Implementation) GetAccessToken(
	ctx context.Context,
	req *desc.GetAccessTokenRequest,
) (*desc.GetAccessTokenResponse, error) {
	refreshTokenStr := req.GetRefreshToken()

	accessToken, err := i.authService.GetAccessToken(ctx, refreshTokenStr)
	if err != nil {
		logger.Error("Failed to get access token: ",
			zap.Any("refresh token", refreshTokenStr),
			zap.Error(err),
		)

		return nil, err
	}

	logger.Info("Get new access token: ", zap.Any("access token", accessToken))

	return &desc.GetAccessTokenResponse{
		AccessToken: *accessToken,
	}, nil
}
