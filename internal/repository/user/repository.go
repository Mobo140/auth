package user

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Mobo140/microservices/auth/internal/model"
	"github.com/Mobo140/microservices/auth/internal/repository"
	"github.com/Mobo140/microservices/auth/internal/repository/user/converter"
	modelRepo "github.com/Mobo140/microservices/auth/internal/repository/user/model"
)

var _ repository.UserRepository = (*repo)(nil)

const (
	tableName       = "client"
	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	roleColumn      = "role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *repo {
	return &repo{db: db}
}
func (r *repo) Create(ctx context.Context, user *model.User) (int64, error) {

	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, roleColumn).
		Values(user.Name, user.Email, user.Role).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query:%v", err)
	}

	var userID int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		log.Fatalf("failed to insert user: %v", err)
	}

	log.Printf("inserted user with id: %d", userID)

	return userID, nil

}

func (r *repo) Get(ctx context.Context, id int64) (*model.UserInfo, error) {

	builderSelectOne := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var info modelRepo.UserInfo
	err = r.db.QueryRow(ctx, query, args...).Scan(&info.ID, &info.Name, &info.Email, &info.Role, &info.CreatedAt, &info.UpdatedAt)
	if err != nil {
		log.Fatalf("failed to select user: %v", err)
	}

	log.Printf("id: %d, name: %s, email: %s, role: %d, created_at: %v, updated_at: %v", info.ID, info.Name, info.Email, info.Role, info.CreatedAt, info.UpdatedAt)

	return converter.ToUserInfoFromRepo(&info), nil
}

func (r *repo) Update(ctx context.Context, id int64, user *model.UpdateUserInfo) error {

	builderUpdate := sq.Update(tableName).
		Set(nameColumn, user.Name).
		Set(emailColumn, user.Email).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to update user: %v", err)
	}

	log.Printf("updated user with id: %d", id)

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {

	builderDelete := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		log.Fatalf("failed to build user: %v", err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to delete user: %v", err)
	}

	log.Printf("deleted user with id: %d", id)

	return nil

}
