package auth

import (
	"context"

	conv "github.com/Mobo140/microservices/auth/internal/converter/auth"
	"github.com/Mobo140/microservices/auth/internal/service"
	desc "github.com/Mobo140/microservices/auth/pkg/auth_v1"
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
		return nil, err
	}

	return &desc.LoginResponse{
		RefreshToken: *refreshTokenData,
	}, nil
}

func (i *Implementation) GetRefreshToken(ctx context.Context, req *desc.GetRefreshTokenRequest) (*desc.GetRefreshTokenResponse, error) {
	refreshTokenStr := req.GetRefreshToken()

	refreshToken, err := i.authService.GetRefreshToken(ctx, refreshTokenStr)
	if err != nil {
		return nil, err
	}

	return &desc.GetRefreshTokenResponse{
		RefreshToken: *refreshToken,
	}, nil
}

func (i *Implementation) GetAccessToken(ctx context.Context, req *desc.GetAccessTokenRequest) (*desc.GetAccessTokenResponse, error) {
	refreshTokenStr := req.GetRefreshToken()

	accessToken, err := i.authService.GetAccessToken(ctx, refreshTokenStr)
	if err != nil {
		return nil, err
	}

	return &desc.GetAccessTokenResponse{
		AccessToken: *accessToken,
	}, nil
}
