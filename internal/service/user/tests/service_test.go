package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	repositoryTx "github.com/Mobo140/microservices/auth/internal/client/db"
	dbTxMocks "github.com/Mobo140/microservices/auth/internal/client/db/mocks"
	"github.com/Mobo140/microservices/auth/internal/model"
	repositoryMocks "github.com/Mobo140/microservices/auth/internal/repository/mocks"
	userService "github.com/Mobo140/microservices/auth/internal/service/user"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	type setupMocks func(
		userRepo *repositoryMocks.UserRepositoryMock,
		logRepo *repositoryMocks.LogRepositoryMock,
		txManager *dbTxMocks.TxManagerMock,
	)

	type args struct {
		req *model.User
	}

	var (
		ctxValue = context.Background()
		mc       = minimock.NewController(t)

		id              = gofakeit.Int64()
		name            = gofakeit.Name()
		email           = gofakeit.Email()
		password        = gofakeit.Password(true, true, true, true, true, 10)
		passwordConfirm = gofakeit.Password(true, true, true, true, true, 10)
		role            = (int64)(0)

		repositoryErr  = fmt.Errorf("create userRepo error")
		logErr         = fmt.Errorf("create log error")
		transactionErr = fmt.Errorf("transaction error")

		user = &model.User{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: passwordConfirm,
			Role:            role,
		}

		req = &model.User{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: passwordConfirm,
			Role:            role,
		}

		logEntry = &model.LogEntry{
			UserID:   id,
			Activity: fmt.Sprintf("Create user: Name=%s, Email=%s, Role=%d", name, email, role),
		}

		unknownUser = (int64)(-1)
	)

	tests := []struct {
		name       string
		args       args
		setupMocks setupMocks
		want       int64
		err        error
	}{
		{
			name: "success case",
			args: args{
				req: req,
			},
			want: id,
			err:  nil,
			setupMocks: func(userRepo *repositoryMocks.UserRepositoryMock,
				logRepo *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userRepo.CreateMock.Expect(ctxValue, user).Return(id, nil)
				logRepo.CreateMock.Expect(ctxValue, logEntry).Return(nil)
				txManager.ReadCommitedMock.Set(func(ctx context.Context, f repositoryTx.Handler) error {
					return f(ctx)
				})
			},
		},
		{
			name: "transaction error",
			args: args{
				req: req,
			},
			want: unknownUser,
			err:  transactionErr,
			setupMocks: func(_ *repositoryMocks.UserRepositoryMock,
				_ *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				txManager.ReadCommitedMock.Set(func(_ context.Context, _ repositoryTx.Handler) error {
					return transactionErr
				})
			},
		},
		{
			name: "userRepo error",
			args: args{
				req: req,
			},
			want: unknownUser,
			err:  repositoryErr,
			setupMocks: func(userRepo *repositoryMocks.UserRepositoryMock,
				_ *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userRepo.CreateMock.Expect(ctxValue, user).Return(unknownUser, repositoryErr)
				txManager.ReadCommitedMock.Set(func(ctxValue context.Context, f repositoryTx.Handler) error {
					return f(ctxValue)
				})
			},
		},
		{
			name: "creating log in db error",
			args: args{
				req: req,
			},
			want: unknownUser,
			err:  logErr,
			setupMocks: func(userRepo *repositoryMocks.UserRepositoryMock,
				logRepo *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userRepo.CreateMock.Expect(ctxValue, user).Return(id, nil)
				logRepo.CreateMock.Expect(ctxValue, logEntry).Return(logErr)
				txManager.ReadCommitedMock.Set(func(ctxValue context.Context, f repositoryTx.Handler) error {
					return f(ctxValue)
				})
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepo := repositoryMocks.NewUserRepositoryMock(mc)

			logRepo := repositoryMocks.NewLogRepositoryMock(mc)
			txManager := dbTxMocks.NewTxManagerMock(mc)

			// Настройка моков в соответствии с тестами
			tt.setupMocks(userRepo, logRepo, txManager)

			service := userService.NewService(userRepo, logRepo, txManager)

			gotID, err := service.Create(ctxValue, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, gotID)
		})
	}
}

func TestGet(t *testing.T) {
	t.Parallel()

	type setupMocks func(
		userRepo *repositoryMocks.UserRepositoryMock,
		logRepo *repositoryMocks.LogRepositoryMock,
		txManager *dbTxMocks.TxManagerMock,
	)

	var (
		ctxValue = context.Background()
		mc       = minimock.NewController(t)

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()
		role      = (int64)(0)

		repositoryErr  = fmt.Errorf("update userRepo error")
		logErr         = fmt.Errorf("update log error")
		transactionErr = fmt.Errorf("transaction error")

		info = &model.UserInfo{
			ID:        id,
			Name:      name,
			Email:     email,
			Role:      role,
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{
				Time:  updatedAt,
				Valid: true,
			},
		}

		logEntry = &model.LogEntry{
			UserID:   id,
			Activity: fmt.Sprintf("Get user: ID: %d, Name: %s, Email: %s, Role: %d", id, name, email, role),
		}
	)

	tests := []struct {
		name       string
		setupMocks setupMocks
		want       *model.UserInfo
		err        error
	}{
		{
			name: "success case",
			want: info,
			err:  nil,
			setupMocks: func(userRepo *repositoryMocks.UserRepositoryMock,
				logRepo *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userRepo.GetMock.Expect(ctxValue, id).Return(info, nil)
				logRepo.CreateMock.Expect(ctxValue, logEntry).Return(nil)
				txManager.ReadCommitedMock.Set(func(ctx context.Context, f repositoryTx.Handler) error {
					return f(ctx)
				})
			},
		},
		{
			name: "transaction error",
			want: nil,
			err:  transactionErr,
			setupMocks: func(_ *repositoryMocks.UserRepositoryMock,
				_ *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				txManager.ReadCommitedMock.Set(func(_ context.Context, _ repositoryTx.Handler) error {
					return transactionErr
				})
			},
		},
		{
			name: "userRepo error",
			want: nil,
			err:  repositoryErr,
			setupMocks: func(userRepo *repositoryMocks.UserRepositoryMock,
				_ *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userRepo.GetMock.Expect(ctxValue, id).Return(nil, repositoryErr)
				txManager.ReadCommitedMock.Set(func(ctxValue context.Context, f repositoryTx.Handler) error {
					return f(ctxValue)
				})
			},
		},
		{
			name: "creating log in db error",
			want: nil,
			err:  logErr,
			setupMocks: func(userRepo *repositoryMocks.UserRepositoryMock,
				logRepo *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userRepo.GetMock.Expect(ctxValue, id).Return(info, nil)
				logRepo.CreateMock.Expect(ctxValue, logEntry).Return(logErr)
				txManager.ReadCommitedMock.Set(func(ctxValue context.Context, f repositoryTx.Handler) error {
					return f(ctxValue)
				})
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepo := repositoryMocks.NewUserRepositoryMock(mc)

			logRepo := repositoryMocks.NewLogRepositoryMock(mc)
			txManager := dbTxMocks.NewTxManagerMock(mc)

			// Настройка моков в соответствии с тестами
			tt.setupMocks(userRepo, logRepo, txManager)

			service := userService.NewService(userRepo, logRepo, txManager)

			userInfo, err := service.Get(ctxValue, id)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, userInfo)
		})
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	type setupMocks func(
		userRepo *repositoryMocks.UserRepositoryMock,
		logRepo *repositoryMocks.LogRepositoryMock,
		txManager *dbTxMocks.TxManagerMock,
	)

	var (
		ctxValue = context.Background()
		mc       = minimock.NewController(t)

		id    = gofakeit.Int64()
		name  = gofakeit.Name()
		email = gofakeit.Email()

		repositoryErr  = fmt.Errorf("get userRepo error")
		logErr         = fmt.Errorf("get log error")
		transactionErr = fmt.Errorf("transaction error")

		info = &model.UpdateUserInfo{
			Name:  name,
			Email: email,
		}

		logEntry = &model.LogEntry{
			UserID:   id,
			Activity: fmt.Sprintf("Update user: Name=%s, Email=%s", name, email),
		}
	)

	tests := []struct {
		name       string
		setupMocks setupMocks
		err        error
	}{
		{
			name: "success case",
			err:  nil,
			setupMocks: func(userRepo *repositoryMocks.UserRepositoryMock,
				logRepo *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userRepo.UpdateMock.Expect(ctxValue, id, info).Return(nil)
				logRepo.CreateMock.Expect(ctxValue, logEntry).Return(nil)
				txManager.ReadCommitedMock.Set(func(ctx context.Context, f repositoryTx.Handler) error {
					return f(ctx)
				})
			},
		},
		{
			name: "transaction error",
			err:  transactionErr,
			setupMocks: func(_ *repositoryMocks.UserRepositoryMock,
				_ *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				txManager.ReadCommitedMock.Set(func(_ context.Context, _ repositoryTx.Handler) error {
					return transactionErr
				})
			},
		},
		{
			name: "userRepo error",
			err:  repositoryErr,
			setupMocks: func(userRepo *repositoryMocks.UserRepositoryMock,
				_ *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userRepo.UpdateMock.Expect(ctxValue, id, info).Return(repositoryErr)
				txManager.ReadCommitedMock.Set(func(ctxValue context.Context, f repositoryTx.Handler) error {
					return f(ctxValue)
				})
			},
		},
		{
			name: "creating log in db error",
			err:  logErr,
			setupMocks: func(userRepo *repositoryMocks.UserRepositoryMock,
				logRepo *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userRepo.UpdateMock.Expect(ctxValue, id, info).Return(nil)
				logRepo.CreateMock.Expect(ctxValue, logEntry).Return(logErr)
				txManager.ReadCommitedMock.Set(func(ctxValue context.Context, f repositoryTx.Handler) error {
					return f(ctxValue)
				})
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepo := repositoryMocks.NewUserRepositoryMock(mc)

			logRepo := repositoryMocks.NewLogRepositoryMock(mc)
			txManager := dbTxMocks.NewTxManagerMock(mc)

			// Настройка моков в соответствии с тестами
			tt.setupMocks(userRepo, logRepo, txManager)

			service := userService.NewService(userRepo, logRepo, txManager)

			err := service.Update(ctxValue, id, info)
			require.Equal(t, tt.err, err)
		})
	}
}
func TestDelete(t *testing.T) {
	t.Parallel()

	type setupMocks func(
		userRepo *repositoryMocks.UserRepositoryMock,
		logRepo *repositoryMocks.LogRepositoryMock,
		txManager *dbTxMocks.TxManagerMock,
	)

	var (
		ctxValue = context.Background()
		mc       = minimock.NewController(t)

		id = gofakeit.Int64()

		repositoryErr  = fmt.Errorf("get userRepo error")
		logErr         = fmt.Errorf("get log error")
		transactionErr = fmt.Errorf("transaction error")

		logEntry = &model.LogEntry{
			UserID:   id,
			Activity: fmt.Sprintf("Delete user: ID=%d", id),
		}
	)

	tests := []struct {
		name       string
		setupMocks setupMocks
		err        error
	}{
		{
			name: "success case",
			err:  nil,
			setupMocks: func(userRepo *repositoryMocks.UserRepositoryMock,
				logRepo *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userRepo.DeleteMock.Expect(ctxValue, id).Return(nil)
				logRepo.CreateMock.Expect(ctxValue, logEntry).Return(nil)
				txManager.ReadCommitedMock.Set(func(ctx context.Context, f repositoryTx.Handler) error {
					return f(ctx)
				})
			},
		},
		{
			name: "transaction error",
			err:  transactionErr,
			setupMocks: func(_ *repositoryMocks.UserRepositoryMock,
				_ *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				txManager.ReadCommitedMock.Set(func(_ context.Context, _ repositoryTx.Handler) error {
					return transactionErr
				})
			},
		},
		{
			name: "userRepo error",
			err:  repositoryErr,
			setupMocks: func(userRepo *repositoryMocks.UserRepositoryMock,
				_ *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userRepo.DeleteMock.Expect(ctxValue, id).Return(repositoryErr)
				txManager.ReadCommitedMock.Set(func(ctxValue context.Context, f repositoryTx.Handler) error {
					return f(ctxValue)
				})
			},
		},
		{
			name: "creating log in db error",
			err:  logErr,
			setupMocks: func(userRepo *repositoryMocks.UserRepositoryMock,
				logRepo *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userRepo.DeleteMock.Expect(ctxValue, id).Return(nil)
				logRepo.CreateMock.Expect(ctxValue, logEntry).Return(logErr)
				txManager.ReadCommitedMock.Set(func(ctxValue context.Context, f repositoryTx.Handler) error {
					return f(ctxValue)
				})
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepo := repositoryMocks.NewUserRepositoryMock(mc)

			logRepo := repositoryMocks.NewLogRepositoryMock(mc)
			txManager := dbTxMocks.NewTxManagerMock(mc)

			// Настройка моков в соответствии с тестами
			tt.setupMocks(userRepo, logRepo, txManager)

			service := userService.NewService(userRepo, logRepo, txManager)

			err := service.Delete(ctxValue, id)
			require.Equal(t, tt.err, err)
		})
	}
}