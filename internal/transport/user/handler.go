package user

import (
	"context"
	"errors"
	"log"

	conv "github.com/Mobo140/auth/internal/converter/user"
	"github.com/Mobo140/auth/internal/service"
	desc "github.com/Mobo140/auth/pkg/user_v1"
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
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	if !i.validatePassword(req.User.Password, req.User.PasswordConfirm) {
		return nil, errors.New("passwords don't match")
	}

	user, err := conv.ToUserFromDesc(req.User)
	if err != nil {
		return nil, err
	}

	id, err := i.userService.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d", id)

	return &desc.CreateResponse{Id: id}, nil
}

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	info, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, name: %s, email: %s, role: %d, created_at: %v, udpdated_at: %v\n",
		info.ID, info.Name, info.Email, info.Role, info.CreatedAt, info.UpdatedAt,
	)

	userInfo, err := conv.ToUserInfoFromService(info)
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Info: userInfo,
	}, nil
}

func (i *Implementation) GetUsers(ctx context.Context, req *desc.GetUsersRequest) (*desc.GetUsersResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	params := conv.ToGetUsersParamsFromDesc(req.GetLimit(), req.GetOffset())

	users, err := i.userService.GetUsers(ctx, params)
	if err != nil {
		return nil, err
	}

	usersInfo, err := conv.ToUsersListFromService(users)
	if err != nil {
		return nil, err
	}

	return &desc.GetUsersResponse{
		Users: usersInfo,
	}, nil
}

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	user, err := conv.ToUpdateUserInfoFromDesc(req.Info)
	if err != nil {
		return nil, err
	}

	err = i.userService.Update(ctx, req.Id, user)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	err = i.userService.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (i *Implementation) validatePassword(password string, passwordConfirm string) bool {
	return password == passwordConfirm
}
