package logs

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/Mobo140/microservices/auth/internal/model"
	"github.com/Mobo140/microservices/auth/internal/repository"
	"github.com/Mobo140/platform_common/pkg/db"
)

var _ repository.LogRepository = (*logRepo)(nil)

const (
	tableNameUser   = "logsUser"
	tableNameAuth   = "logsAuth"
	userIdColumn    = "user_id"
	userNameColumn  = "username"
	activityColumn  = "activity"
	createdAtColumn = "created_at"
)

type logRepo struct {
	db db.Client
}

func NewRepository(db db.Client) *logRepo { //nolint:revive // it's ok
	return &logRepo{db: db}
}

func (l *logRepo) CreateLogUser(ctx context.Context, lg *model.LogEntryUser) error {
	builder := sq.Insert(tableNameUser).
		PlaceholderFormat(sq.Dollar).
		Columns(userIdColumn, activityColumn, createdAtColumn).
		Values(lg.UserID, lg.Action, lg.CreatedAt)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "create_log",
		QueryRow: query,
	}

	_, err = l.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (l *logRepo) CreateLogAuth(ctx context.Context, lg *model.LogEntryAuth) error {
	builder := sq.Insert(tableNameAuth).
		PlaceholderFormat(sq.Dollar).
		Columns(userNameColumn, activityColumn, createdAtColumn).
		Values(lg.Username, lg.Action, lg.CreatedAt)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "create_log",
		QueryRow: query,
	}

	_, err = l.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
