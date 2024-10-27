package user

import (
	"context"
	"log"

	conv "github.com/Mobo140/microservices/auth/internal/converter"
	"github.com/Mobo140/microservices/auth/internal/service"
	"github.com/Mobo140/microservices/auth/internal/transport"
	desc "github.com/Mobo140/microservices/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ transport.UserHandler = (*implementation)(nil)

type implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *implementation {
	return &implementation{
		userService: userService,
	}
}

func (i *implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	user, err := conv.ToUserFromDesc(req.User)
	if err != nil {
		return nil, err
	}
	id, err := i.userService.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{Id: id}, nil
}

func (i *implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
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

func (i *implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	user, err := conv.ToUpdateUserInfoFromDesc(req.Info)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	err = i.userService.Update(ctx, req.Id, user)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (i *implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {

	err := i.userService.Delete(ctx, req.Id)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
