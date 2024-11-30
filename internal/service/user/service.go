package user

import (
	"context"
	"fmt"

	"github.com/Mobo140/microservices/auth/internal/model"
	"github.com/Mobo140/microservices/auth/internal/repository"
	"github.com/Mobo140/microservices/auth/internal/service"
	"github.com/Mobo140/platform_common/pkg/db"
)

var _ service.UserService = (*serv)(nil)

const (
	unknownUser = -1
)

type serv struct {
	dbRepository  repository.UserDBRepository
	logRepository repository.LogRepository
	txManager     db.TxManager
}

func NewService(
	dbRepository repository.UserDBRepository,
	logRepository repository.LogRepository,
	txManager db.TxManager,
) *serv { //nolint:revive //it's ok
	return &serv{
		dbRepository:  dbRepository,
		logRepository: logRepository,
		txManager:     txManager,
	}
}

func (s *serv) Create(ctx context.Context, user *model.User) (int64, error) {
	var id int64
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var errTx error
		if id, errTx = s.dbRepository.Create(ctx, user); errTx != nil {
			return errTx
		}

		logEntry := model.LogEntryUser{
			UserID: id,
			Action: fmt.Sprintf("Create user: Name=%s, Email=%s, Role=%d", user.Name, user.Email, user.Role),
		}

		errTx = s.logRepository.CreateLogUser(ctx, &logEntry)
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
		info, errTx = s.dbRepository.Get(ctx, id)

		if errTx != nil {
			return errTx
		}

		logEntry := model.LogEntryUser{
			UserID: id,
			Action: fmt.Sprintf("Get user: ID: %d, Name: %s, Email: %s, Role: %d", info.ID, info.Name, info.Email, info.Role),
		}

		errTx = s.logRepository.CreateLogUser(ctx, &logEntry)
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

func (s *serv) GetUsers(ctx context.Context, params *model.GetUsersRequest) ([]*model.UserInfo, error) {
	var usersList []*model.UserInfo
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var errTx error
		if usersList, errTx = s.dbRepository.GetUsers(ctx, params); errTx != nil {
			return errTx
		}

		logEntry := model.LogEntryUser{
			Action: fmt.Sprintf("Get users: from %d to %d", params.Offset+1, params.Offset+(int64)(len(usersList))),
		}

		errTx = s.logRepository.CreateLogUser(ctx, &logEntry)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return usersList, nil
}

func (s *serv) Update(ctx context.Context, id int64, user *model.UpdateUserInfo) error {
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var errTx error
		if errTx = s.dbRepository.Update(ctx, id, user); errTx != nil {
			return errTx
		}

		logEntry := model.LogEntryUser{
			UserID: id,
			Action: fmt.Sprintf("Update user: Name=%s, Email=%s", user.Name, user.Email),
		}

		errTx = s.logRepository.CreateLogUser(ctx, &logEntry)
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
		if errTx = s.dbRepository.Delete(ctx, id); errTx != nil {
			return errTx
		}

		logEntry := model.LogEntryUser{
			UserID: id,
			Action: fmt.Sprintf("Delete user: ID=%d", id),
		}

		errTx = s.logRepository.CreateLogUser(ctx, &logEntry)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	return err
}
