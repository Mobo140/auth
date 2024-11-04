package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/Mobo140/microservices/auth/internal/model"
	"github.com/Mobo140/microservices/auth/internal/service"
	serviceMocks "github.com/Mobo140/microservices/auth/internal/service/mocks"
	userHandler "github.com/Mobo140/microservices/auth/internal/transport/user"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	desc "github.com/Mobo140/microservices/auth/pkg/user_v1"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		req *desc.CreateRequest
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

		serviceErr      = fmt.Errorf("service create error")
		conversationErr = fmt.Errorf("invalid role value: 5")

		req = &desc.CreateRequest{
			User: &desc.User{
				Name:            name,
				Email:           email,
				Password:        password,
				PasswordConfirm: passwordConfirm,
				Role:            desc.Role_USER,
			},
		}

		user = &model.User{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: passwordConfirm,
			Role:            role,
		}

		res = &desc.CreateResponse{
			Id: id,
		}

		unknownUser = (int64)(1)
	)

	tests := []struct {
		name            string
		args            args
		userServiceMock userServiceMockFunc
		want            *desc.CreateResponse
		err             error
	}{
		{
			name: "success",
			args: args{
				req: req,
			},
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.CreateMock.Expect(ctxValue, user).Return(id, nil)
				return mock
			},
			want: res,
			err:  nil,
		},
		{
			name: "service error case",
			args: args{
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.CreateMock.Expect(ctxValue, user).Return(unknownUser, serviceErr)
				return mock
			},
		},
		{
			name: "conversation error case",
			args: args{
				req: &desc.CreateRequest{
					User: &desc.User{
						Name:            name,
						Email:           email,
						Password:        password,
						PasswordConfirm: passwordConfirm,
						Role:            5,
					},
				},
			},
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				return mock
			},
			want: nil,
			err:  conversationErr,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			transport := userHandler.NewImplementation(userServiceMock)

			newID, err := transport.Create(ctxValue, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}

func TestGet(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		req *desc.GetRequest
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

		serviceErr      = fmt.Errorf("service get error")
		conversationErr = fmt.Errorf("invalid role value: 5")

		req = &desc.GetRequest{
			Id: id,
		}

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

		res = &desc.GetResponse{
			Info: &desc.UserInfo{
				Id:        id,
				Name:      name,
				Email:     email,
				Role:      desc.Role_USER,
				CreatedAt: timestamppb.New(createdAt),
				UpdatedAt: timestamppb.New(updatedAt),
			},
		}
	)

	tests := []struct {
		name            string
		args            args
		userServiceMock userServiceMockFunc
		want            *desc.GetResponse
		err             error
	}{
		{
			name: "success",
			args: args{
				req: req,
			},
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.GetMock.Expect(ctxValue, id).Return(info, nil)
				return mock
			},
			want: res,
			err:  nil,
		},
		{
			name: "service error case",
			args: args{
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.GetMock.Expect(ctxValue, id).Return(nil, serviceErr)
				return mock
			},
		},
		{
			name: "conversation error case",
			args: args{
				req: req,
			},
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.GetMock.Expect(ctxValue, id).Return(&model.UserInfo{
					ID:        id,
					Name:      name,
					Email:     email,
					Role:      5,
					CreatedAt: createdAt,
					UpdatedAt: sql.NullTime{
						Time:  updatedAt,
						Valid: true,
					},
				}, nil)

				return mock
			},
			want: nil,
			err:  conversationErr,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			transport := userHandler.NewImplementation(userServiceMock)

			response, err := transport.Get(ctxValue, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, response)
		})
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		req *desc.UpdateRequest
	}

	var (
		ctxValue = context.Background()
		mc       = minimock.NewController(t)

		id    = gofakeit.Int64()
		name  = gofakeit.Name()
		email = gofakeit.Email()

		serviceErr  = fmt.Errorf("service update error")
		converseErr = fmt.Errorf("update user info is empty")

		req = &desc.UpdateRequest{
			Id: id,
			Info: &desc.UpdateUserInfo{
				Name:  wrapperspb.String(name),
				Email: wrapperspb.String(email),
			},
		}

		info = &model.UpdateUserInfo{
			Name:  name,
			Email: email,
		}

		res = &emptypb.Empty{}
	)

	tests := []struct {
		name            string
		args            args
		userServiceMock userServiceMockFunc
		want            *emptypb.Empty
		err             error
	}{
		{
			name: "success",
			args: args{
				req: req,
			},
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctxValue, id, info).Return(nil)
				return mock
			},
			want: res,
			err:  nil,
		},
		{
			name: "service error case",
			args: args{
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctxValue, id, info).Return(serviceErr)
				return mock
			},
		},
		{
			name: "conversation error case",
			args: args{
				req: &desc.UpdateRequest{
					Id:   id,
					Info: nil,
				},
			},
			want: nil,
			err:  converseErr,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			transport := userHandler.NewImplementation(userServiceMock)

			response, err := transport.Update(ctxValue, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, response)
		})
	}
}

func TestDelete(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		req *desc.DeleteRequest
	}

	var (
		ctxValue = context.Background()
		mc       = minimock.NewController(t)

		id = gofakeit.Int64()

		serviceErr = fmt.Errorf("service delete error")

		req = &desc.DeleteRequest{
			Id: id,
		}

		res = &emptypb.Empty{}
	)

	tests := []struct {
		name            string
		args            args
		userServiceMock userServiceMockFunc
		want            *emptypb.Empty
		err             error
	}{
		{
			name: "success",
			args: args{
				req: req,
			},
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.DeleteMock.Expect(ctxValue, id).Return(nil)
				return mock
			},
			want: res,
			err:  nil,
		},
		{
			name: "service error case",
			args: args{
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.DeleteMock.Expect(ctxValue, id).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			transport := userHandler.NewImplementation(userServiceMock)

			response, err := transport.Delete(ctxValue, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, response)
		})
	}
}
