package user

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"

	"github.com/Mobo140/microservices/auth/internal/client/db"
	"github.com/Mobo140/microservices/auth/internal/model"
	"github.com/Mobo140/microservices/auth/internal/repository"
	"github.com/Mobo140/microservices/auth/internal/repository/user/db/converter"
	modelRepo "github.com/Mobo140/microservices/auth/internal/repository/user/db/model"
)

var _ repository.UserDBRepository = (*userRepo)(nil)

const (
	tableName       = "client"
	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	pshashColumn    = "hash_password"
	roleColumn      = "role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type userRepo struct {
	db db.Client
}

func NewRepository(db db.Client) *userRepo { //nolint:revive // it's ok
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user *model.User) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, pshashColumn, roleColumn).
		Values(user.Name, user.Email, user.HashedPassword, user.Role).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRow: query,
	}

	var userID int64
	if err = r.db.DB().ScanOneContext(ctx, &userID, q, args...); err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *userRepo) Get(ctx context.Context, id int64) (*model.UserInfo, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRow: query,
	}

	var info modelRepo.UserInfo

	if err = r.db.DB().ScanOneContext(ctx, &info, q, args...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrorUserNotFound
		}
		return nil, err
	}

	return converter.ToUserInfoFromRepo(&info), nil
}

func (r *userRepo) GetUsers(ctx context.Context, params *model.GetUsersRequest) ([]*model.UserInfo, error) {
	minID := params.Offset + 1
	maxID := params.Offset + params.Limit
	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.And{
			sq.GtOrEq{idColumn: minID},
			sq.LtOrEq{idColumn: maxID},
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.GetUsers",
		QueryRow: query,
	}

	var usersList []*modelRepo.UserInfo

	if err = r.db.DB().ScanAllContext(ctx, &usersList, q, args...); err != nil {
		return nil, err
	}

	return converter.ToUsersInfoFromRepo(usersList), nil
}

func (r *userRepo) Update(ctx context.Context, id int64, user *model.UpdateUserInfo) error {
	builder := sq.Update(tableName).
		Set(nameColumn, user.Name).
		Set(emailColumn, user.Email).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRow: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Delete",
		QueryRow: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) GetHashAndRoleByUsername(ctx context.Context, username string) (*model.UserAuthData, error) {
	builder := sq.Select(pshashColumn, roleColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{nameColumn: username})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.GetHashAndRoleByUsername",
		QueryRow: query,
	}

	var data *modelRepo.UserAuthData
	if err = r.db.DB().ScanOneContext(ctx, &data, q, args...); err != nil {
		return nil, err
	}

	return converter.ToUserAuthDataFromRepo(data), nil
}

func (r *userRepo) GetRoleByUsername(ctx context.Context, username string) (*string, error) {
	builder := sq.Select(roleColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{nameColumn: username})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.GetRoleByUsername",
		QueryRow: query,
	}

	if err = r.db.DB().ScanOneContext(ctx, &username, q, args...); err != nil {
		return nil, err
	}

	return &username, nil
}
