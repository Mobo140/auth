package user

import (
	"context"

	"github.com/Mobo140/microservices/auth/internal/model"
	"github.com/Mobo140/microservices/auth/internal/repository"
	"github.com/Mobo140/microservices/auth/internal/service"
)

var _ service.UserService = (*serv)(nil)

type serv struct {
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) *serv {
	return &serv{
		userRepository: userRepository,
	}
}

func (s *serv) Create(ctx context.Context, user *model.User) (int64, error) {

	id, err := s.userRepository.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *serv) Get(ctx context.Context, id int64) (*model.UserInfo, error) {

	info, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (s *serv) Update(ctx context.Context, id int64, user *model.UpdateUserInfo) error {
	return s.userRepository.Update(ctx, id, user)
}

func (s *serv) Delete(ctx context.Context, id int64) error {
	return s.userRepository.Delete(ctx, id)
}
