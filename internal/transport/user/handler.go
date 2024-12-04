package user

import (
	"context"
	"errors"

	conv "github.com/Mobo140/auth/internal/converter/user"
	"github.com/Mobo140/auth/internal/service"
	desc "github.com/Mobo140/auth/pkg/user_v1"
	"github.com/Mobo140/platform_common/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	logger.Info("Creating user...", zap.Any("user", req.GetUser()))

	if !i.validatePassword(req.User.Password, req.User.PasswordConfirm) {
		err := errors.New("passwords don't match")
		logger.Error("validate Password: ", zap.Error(err))

		return nil, err
	}

	user, err := conv.ToUserFromDesc(req.User)
	if err != nil {
		logger.Error("Failed to convert to user from desc: ", zap.Error(err))

		return nil, err
	}

	id, err := i.userService.Create(ctx, user)
	if err != nil {
		logger.Error("Failed to create user: ", zap.Error(err))

		return nil, err
	}

	logger.Info("Create user: ", zap.Any("id", id))

	return &desc.CreateResponse{Id: id}, nil
}

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	logger.Info("Getting user...", zap.Int64("id", req.GetId()))

	info, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		logger.Error("Failed to get user info: ", zap.Error(err))

		return nil, err
	}

	logger.Info("Get user: ", zap.Any("info", info))

	userInfo, err := conv.ToUserInfoFromService(info)
	if err != nil {
		logger.Error("Failed to convert to user info from service: ", zap.Error(err))

		return nil, err
	}

	return &desc.GetResponse{
		Info: userInfo,
	}, nil
}

func (i *Implementation) GetUsers(ctx context.Context, req *desc.GetUsersRequest) (*desc.GetUsersResponse, error) {
	logger.Info("Getting users...",
		zap.Int64("with limit", req.GetLimit()),
		zap.Int64("with offset", req.GetOffset()),
	)

	params := conv.ToGetUsersParamsFromDesc(req.GetLimit(), req.GetOffset())

	users, err := i.userService.GetUsers(ctx, params)
	if err != nil {
		logger.Error("Failed to get users: ", zap.Error(err))

		return nil, err
	}

	logger.Info("Get users: ", zap.Any("users", users))

	usersInfo, err := conv.ToUsersListFromService(users)
	if err != nil {
		logger.Error("Failed to convert to users info from service: ", zap.Error(err))

		return nil, err
	}

	return &desc.GetUsersResponse{
		Users: usersInfo,
	}, nil
}

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	logger.Info("Update user info...",
		zap.Int64("with id", req.GetId()),
		zap.Any("and info", req.GetInfo()),
	)

	user, err := conv.ToUpdateUserInfoFromDesc(req.Info)
	if err != nil {
		logger.Error("Failed to convert to info from desc: ", zap.Error(err))

		return nil, err
	}

	err = i.userService.Update(ctx, req.Id, user)
	if err != nil {
		logger.Error("Failed to update user info: ", zap.Error(err))

		return nil, err
	}

	logger.Info("Update user: ", zap.Any("id", req.GetId()))

	return &emptypb.Empty{}, nil
}

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	logger.Info("Delete user", zap.Int64("with id", req.GetId()))

	err := i.userService.Delete(ctx, req.Id)
	if err != nil {
		logger.Error("Failed to delete user: ", zap.Error(err))

		return nil, err
	}

	logger.Info("Delete user: ", zap.Int64("with id", req.GetId()))

	return &emptypb.Empty{}, nil
}

func (i *Implementation) validatePassword(password string, passwordConfirm string) bool {
	return password == passwordConfirm
}
