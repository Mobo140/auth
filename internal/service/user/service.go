package user

import (
	"context"
	"fmt"

	"github.com/Mobo140/microservices/auth/internal/client/db"
	"github.com/Mobo140/microservices/auth/internal/model"
	"github.com/Mobo140/microservices/auth/internal/repository"
	"github.com/Mobo140/microservices/auth/internal/service"
)

var _ service.UserService = (*serv)(nil)

const (
	unknownUser = -1
)

type serv struct {
	userRepository repository.UserRepository
	logRepository  repository.LogRepository
	txManager      db.TxManager
}

func NewService(
	userRepository repository.UserRepository,
	logRepository repository.LogRepository,
	txManager db.TxManager,
) *serv { // nolint
	return &serv{
		userRepository: userRepository,
		logRepository:  logRepository,
		txManager:      txManager,
	}
}

func (s *serv) Create(ctx context.Context, user *model.User) (int64, error) {
	var id int64
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var errTx error
		if id, errTx = s.userRepository.Create(ctx, user); errTx != nil {
			return errTx
		}

		logEntry := model.LogEntry{
			UserID:   id,
			Activity: fmt.Sprintf("Create user: Name=%s, Email=%s, Role=%d", user.Name, user.Email, user.Role),
		}

		errTx = s.logRepository.Create(ctx, &logEntry)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return unknownUser, err
	}

	return id, nil
}

func (s *serv) Get(ctx context.Context, id int64) (*model.UserInfo, error) {
	var info *model.UserInfo
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var errTx error
		if info, errTx = s.userRepository.Get(ctx, id); errTx != nil {
			return errTx
		}

		logEntry := model.LogEntry{
			UserID:   id,
			Activity: fmt.Sprintf("Get user: ID: %d, Name: %s, Email: %s, Role: %d", info.ID, info.Name, info.Email, info.Role),
		}

		errTx = s.logRepository.Create(ctx, &logEntry)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return info, nil
}

func (s *serv) Update(ctx context.Context, id int64, user *model.UpdateUserInfo) error {
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var errTx error
		if errTx = s.userRepository.Update(ctx, id, user); errTx != nil {
			return errTx
		}

		logEntry := model.LogEntry{
			UserID:   id,
			Activity: fmt.Sprintf("Update user: Name=%s, Email=%s", user.Name, user.Email),
		}

		errTx = s.logRepository.Create(ctx, &logEntry)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	return err
}

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var errTx error
		if errTx = s.userRepository.Delete(ctx, id); errTx != nil {
			return errTx
		}

		logEntry := model.LogEntry{
			UserID:   id,
			Activity: fmt.Sprintf("Delete user: ID=%d", id),
		}

		errTx = s.logRepository.Create(ctx, &logEntry)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	return err
}
