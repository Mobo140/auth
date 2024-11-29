package model

import (
	"database/sql"
	"time"
)

type AccessEndpoint struct {
	Endpoint  string       `db:"endpoint"`
	Role      int64        `db:"role"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
