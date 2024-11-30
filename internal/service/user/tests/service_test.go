package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/Mobo140/microservices/auth/internal/model"
	repositoryMocks "github.com/Mobo140/microservices/auth/internal/repository/mocks"
	userService "github.com/Mobo140/microservices/auth/internal/service/user"
	repositoryTx "github.com/Mobo140/platform_common/pkg/db"
	dbTxMocks "github.com/Mobo140/platform_common/pkg/db/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	type setupMocks func(
		userDBRepo *repositoryMocks.UserDBRepositoryMock,
		logRepo *repositoryMocks.LogRepositoryMock,
		txManager *dbTxMocks.TxManagerMock,
	)

	type args struct {
		req *model.User
	}

	var (
		ctxValue = context.Background()
		mc       = minimock.NewController(t)

		id       = gofakeit.Int64()
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, true, 10)
		role     = (int64)(0)

		hashPassword, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //nolint:errcheck //it's ok
		repositoryErr   = fmt.Errorf("create userRepo error")
		logErr          = fmt.Errorf("create log error")
		transactionErr  = fmt.Errorf("transaction error")

		user = &model.User{
			Name:           name,
			Email:          email,
			HashedPassword: (string)(hashPassword),
			Role:           role,
		}

		req = &model.User{
			Name:           name,
			Email:          email,
			HashedPassword: (string)(hashPassword),
			Role:           role,
		}

		logEntry = &model.LogEntryUser{
			UserID: id,
			Action: fmt.Sprintf("Create user: Name=%s, Email=%s, Role=%d", name, email, role),
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
			setupMocks: func(userDBRepo *repositoryMocks.UserDBRepositoryMock,
				logRepo *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userDBRepo.CreateMock.Expect(ctxValue, user).Return(id, nil)
				logRepo.CreateLogUserMock.Expect(ctxValue, logEntry).Return(nil)
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
			setupMocks: func(_ *repositoryMocks.UserDBRepositoryMock,
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
			setupMocks: func(userRepo *repositoryMocks.UserDBRepositoryMock,
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
			setupMocks: func(userRepo *repositoryMocks.UserDBRepositoryMock,
				logRepo *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userRepo.CreateMock.Expect(ctxValue, user).Return(id, nil)
				logRepo.CreateLogUserMock.Expect(ctxValue, logEntry).Return(logErr)
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

			dbRepo := repositoryMocks.NewUserDBRepositoryMock(mc)
			logRepo := repositoryMocks.NewLogRepositoryMock(mc)
			txManager := dbTxMocks.NewTxManagerMock(mc)

			// Настройка моков в соответствии с тестами
			tt.setupMocks(dbRepo, logRepo, txManager)

			service := userService.NewService(dbRepo, logRepo, txManager)

			gotID, err := service.Create(ctxValue, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, gotID)
		})
	}
}

func TestGet(t *testing.T) {
	t.Parallel()

	type setupMocks func(
		userDBRepo *repositoryMocks.UserDBRepositoryMock,
		logRepo *repositoryMocks.LogRepositoryMock,
		txManager *dbTxMocks.TxManagerMock,
	)

	type args struct {
		req int64
	}

	var (
		ctxValue = context.Background()
		mc       = minimock.NewController(t)

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()
		role      = (int64)(0)

		repositoryErr  = fmt.Errorf("update chatRepo error")
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

		logEntry = &model.LogEntryUser{
			UserID: id,
			Action: fmt.Sprintf("Get user: ID: %d, Name: %s, Email: %s, Role: %d", id, name, email, role),
		}
	)

	tests := []struct {
		name       string
		setupMocks setupMocks
		args       args
		want       *model.UserInfo
		err        error
	}{
		{
			name: "success case",
			want: info,
			args: args{
				req: id,
			},
			err: nil,
			setupMocks: func(userRepo *repositoryMocks.UserDBRepositoryMock,
				logRepo *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userRepo.GetMock.Expect(ctxValue, id).Return(info, nil)
				logRepo.CreateLogUserMock.Expect(ctxValue, logEntry).Return(nil)
				txManager.ReadCommitedMock.Set(func(ctx context.Context, f repositoryTx.Handler) error {
					return f(ctx)
				})
			},
		},
		{
			name: "transaction error",
			want: nil,
			args: args{
				req: id,
			},
			err: transactionErr,
			setupMocks: func(_ *repositoryMocks.UserDBRepositoryMock,
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
			args: args{
				req: id,
			},
			err: repositoryErr,
			setupMocks: func(userDBRepo *repositoryMocks.UserDBRepositoryMock,
				_ *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userDBRepo.GetMock.Expect(ctxValue, id).Return(nil, repositoryErr)
				txManager.ReadCommitedMock.Set(func(ctxValue context.Context, f repositoryTx.Handler) error {
					return f(ctxValue)
				})
			},
		},
		{
			name: "creating log in db error",
			want: nil,
			err:  logErr,
			args: args{
				req: id,
			},
			setupMocks: func(userDBRepo *repositoryMocks.UserDBRepositoryMock,
				logRepo *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userDBRepo.GetMock.Expect(ctxValue, id).Return(info, nil)
				logRepo.CreateLogUserMock.Expect(ctxValue, logEntry).Return(logErr)
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

			userRepo := repositoryMocks.NewUserDBRepositoryMock(mc)

			logRepo := repositoryMocks.NewLogRepositoryMock(mc)
			txManager := dbTxMocks.NewTxManagerMock(mc)

			// Настройка моков в соответствии с тестами
			tt.setupMocks(userRepo, logRepo, txManager)

			service := userService.NewService(userRepo, logRepo, txManager)

			userInfo, err := service.Get(ctxValue, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, userInfo)
		})
	}
}

func TestGetUsers(t *testing.T) {
	t.Parallel()

	type setupMocks func(
		userRepo *repositoryMocks.UserDBRepositoryMock,
		logRepo *repositoryMocks.LogRepositoryMock,
		txManager *dbTxMocks.TxManagerMock,
	)

	type args struct {
		req *model.GetUsersRequest
	}

	var (
		ctxValue = context.Background()
		mc       = minimock.NewController(t)

		idFirst        = gofakeit.Int64()
		nameFirst      = gofakeit.Name()
		emailFirst     = gofakeit.Email()
		createdAtFirst = gofakeit.Date()
		updatedAtFirst = gofakeit.Date()
		roleFirst      = (int64)(0)

		idSecond        = gofakeit.Int64()
		nameSecond      = gofakeit.Name()
		emailSecond     = gofakeit.Email()
		createdAtSecond = gofakeit.Date()
		updatedAtSecond = gofakeit.Date()
		roleSecond      = (int64)(0)

		repositoryErr  = fmt.Errorf("update chatRepo error")
		logErr         = fmt.Errorf("update log error")
		transactionErr = fmt.Errorf("transaction error")

		req = &model.GetUsersRequest{
			Limit:  2,
			Offset: 0,
		}

		params = &model.GetUsersRequest{
			Limit:  2,
			Offset: 0,
		}

		usersList = []*model.UserInfo{
			{
				ID:        idFirst,
				Name:      nameFirst,
				Email:     emailFirst,
				Role:      roleFirst,
				CreatedAt: createdAtFirst,
				UpdatedAt: sql.NullTime{
					Time:  updatedAtFirst,
					Valid: true,
				},
			},
			{
				ID:        idSecond,
				Name:      nameSecond,
				Email:     emailSecond,
				Role:      roleSecond,
				CreatedAt: createdAtSecond,
				UpdatedAt: sql.NullTime{
					Time:  updatedAtSecond,
					Valid: true,
				},
			},
		}

		res = []*model.UserInfo{
			{
				ID:        idFirst,
				Name:      nameFirst,
				Email:     emailFirst,
				Role:      roleFirst,
				CreatedAt: createdAtFirst,
				UpdatedAt: sql.NullTime{
					Time:  updatedAtFirst,
					Valid: true,
				},
			},
			{
				ID:        idSecond,
				Name:      nameSecond,
				Email:     emailSecond,
				Role:      roleSecond,
				CreatedAt: createdAtSecond,
				UpdatedAt: sql.NullTime{
					Time:  updatedAtSecond,
					Valid: true,
				},
			},
		}

		logEntry = &model.LogEntryUser{
			Action: fmt.Sprintf("Get users: from %d to %d", params.Offset+1, params.Offset+(int64)(len(usersList))),
		}
	)

	tests := []struct {
		name       string
		setupMocks setupMocks
		args       args
		want       []*model.UserInfo
		err        error
	}{
		{
			name: "success case",
			want: res,
			args: args{
				req: req,
			},
			err: nil,
			setupMocks: func(userRepo *repositoryMocks.UserDBRepositoryMock,
				logRepo *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userRepo.GetUsersMock.Expect(ctxValue, params).Return(usersList, nil)
				logRepo.CreateLogUserMock.Expect(ctxValue, logEntry).Return(nil)
				txManager.ReadCommitedMock.Set(func(ctx context.Context, f repositoryTx.Handler) error {
					return f(ctx)
				})
			},
		},
		{
			name: "transaction error",
			want: nil,
			args: args{
				req: req,
			},
			err: transactionErr,
			setupMocks: func(_ *repositoryMocks.UserDBRepositoryMock,
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
			args: args{
				req: req,
			},
			err: repositoryErr,
			setupMocks: func(userRepo *repositoryMocks.UserDBRepositoryMock,
				_ *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userRepo.GetUsersMock.Expect(ctxValue, params).Return(nil, repositoryErr)
				txManager.ReadCommitedMock.Set(func(ctxValue context.Context, f repositoryTx.Handler) error {
					return f(ctxValue)
				})
			},
		},
		{
			name: "creating log in db error",
			want: nil,
			err:  logErr,
			args: args{
				req: req,
			},
			setupMocks: func(userRepo *repositoryMocks.UserDBRepositoryMock,
				logRepo *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userRepo.GetUsersMock.Expect(ctxValue, params).Return(usersList, nil)
				logRepo.CreateLogUserMock.Expect(ctxValue, logEntry).Return(logErr)
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

			userRepo := repositoryMocks.NewUserDBRepositoryMock(mc)

			logRepo := repositoryMocks.NewLogRepositoryMock(mc)
			txManager := dbTxMocks.NewTxManagerMock(mc)

			// Настройка моков в соответствии с тестами
			tt.setupMocks(userRepo, logRepo, txManager)

			service := userService.NewService(userRepo, logRepo, txManager)

			userInfo, err := service.GetUsers(ctxValue, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, userInfo)
		})
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	type setupMocks func(
		userRepo *repositoryMocks.UserDBRepositoryMock,
		logRepo *repositoryMocks.LogRepositoryMock,
		txManager *dbTxMocks.TxManagerMock,
	)

	type args struct {
		ID  int64
		req *model.UpdateUserInfo
	}

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

		logEntry = &model.LogEntryUser{
			UserID: id,
			Action: fmt.Sprintf("Update user: Name=%s, Email=%s", name, email),
		}
	)

	tests := []struct {
		name       string
		setupMocks setupMocks
		args       args
		err        error
	}{
		{
			name: "success case",
			err:  nil,
			args: args{
				ID: id,
				req: &model.UpdateUserInfo{
					Name:  name,
					Email: email,
				},
			},

			setupMocks: func(userRepo *repositoryMocks.UserDBRepositoryMock,
				logRepo *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userRepo.UpdateMock.Expect(ctxValue, id, info).Return(nil)
				logRepo.CreateLogUserMock.Expect(ctxValue, logEntry).Return(nil)
				txManager.ReadCommitedMock.Set(func(ctx context.Context, f repositoryTx.Handler) error {
					return f(ctx)
				})
			},
		},
		{
			name: "transaction error",
			args: args{
				ID: id,
				req: &model.UpdateUserInfo{
					Name:  name,
					Email: email,
				},
			},

			err: transactionErr,
			setupMocks: func(_ *repositoryMocks.UserDBRepositoryMock,
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
				ID: id,
				req: &model.UpdateUserInfo{
					Name:  name,
					Email: email,
				},
			},
			err: repositoryErr,
			setupMocks: func(userRepo *repositoryMocks.UserDBRepositoryMock,
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
			args: args{
				ID: id,
				req: &model.UpdateUserInfo{
					Name:  name,
					Email: email,
				},
			},
			err: logErr,
			setupMocks: func(userRepo *repositoryMocks.UserDBRepositoryMock,
				logRepo *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userRepo.UpdateMock.Expect(ctxValue, id, info).Return(nil)
				logRepo.CreateLogUserMock.Expect(ctxValue, logEntry).Return(logErr)
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

			userRepo := repositoryMocks.NewUserDBRepositoryMock(mc)

			logRepo := repositoryMocks.NewLogRepositoryMock(mc)
			txManager := dbTxMocks.NewTxManagerMock(mc)

			// Настройка моков в соответствии с тестами
			tt.setupMocks(userRepo, logRepo, txManager)

			service := userService.NewService(userRepo, logRepo, txManager)

			err := service.Update(ctxValue, tt.args.ID, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
func TestDelete(t *testing.T) {
	t.Parallel()

	type setupMocks func(
		userRepo *repositoryMocks.UserDBRepositoryMock,
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

		logEntry = &model.LogEntryUser{
			UserID: id,
			Action: fmt.Sprintf("Delete user: ID=%d", id),
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
			setupMocks: func(userRepo *repositoryMocks.UserDBRepositoryMock,
				logRepo *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userRepo.DeleteMock.Expect(ctxValue, id).Return(nil)
				logRepo.CreateLogUserMock.Expect(ctxValue, logEntry).Return(nil)
				txManager.ReadCommitedMock.Set(func(ctx context.Context, f repositoryTx.Handler) error {
					return f(ctx)
				})
			},
		},
		{
			name: "transaction error",
			err:  transactionErr,
			setupMocks: func(_ *repositoryMocks.UserDBRepositoryMock,
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
			setupMocks: func(userRepo *repositoryMocks.UserDBRepositoryMock,
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
			setupMocks: func(userRepo *repositoryMocks.UserDBRepositoryMock,
				logRepo *repositoryMocks.LogRepositoryMock,
				txManager *dbTxMocks.TxManagerMock,
			) {
				userRepo.DeleteMock.Expect(ctxValue, id).Return(nil)
				logRepo.CreateLogUserMock.Expect(ctxValue, logEntry).Return(logErr)
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

			userRepo := repositoryMocks.NewUserDBRepositoryMock(mc)

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
