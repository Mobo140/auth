package access

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/Mobo140/auth/internal/model"
	"github.com/Mobo140/auth/internal/repository"
	"github.com/Mobo140/auth/internal/repository/access/db/converter"
	modelRepo "github.com/Mobo140/auth/internal/repository/access/db/model"
	"github.com/Mobo140/platform_common/pkg/db"
)

var _ repository.AccessDBRepository = (*accessRepo)(nil)

const (
	tableName      = "access"
	columnEndpoint = "endpoint"
	columnRole     = "role"
)

type accessRepo struct {
	db db.Client
}

func NewRepository(db db.Client) *accessRepo { //nolint:revive //it's ok
	return &accessRepo{db: db}
}

func (r *accessRepo) GetEndpointsAccess(ctx context.Context) ([]*model.AccessEndpoint, error) {
	builder := sq.Select(columnEndpoint, columnRole).
		From(tableName)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "access_repository.GetEndpointsAccess",
		QueryRow: query,
	}

	var endpoints []*modelRepo.AccessEndpoint
	if err := r.db.DB().ScanAllContext(ctx, &endpoints, q, args...); err != nil {
		return nil, err
	}

	return converter.ToEndpointsAccessFromRepo(endpoints), nil
}
