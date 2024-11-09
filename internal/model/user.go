package model

import (
	"database/sql"
	"time"
)

type UserInfo struct {
	ID        int64
	Name      string
	Email     string
	Role      int64
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type User struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            int64
}

type UpdateUserInfo struct {
	Name  string
	Email string
}

type GetUsersRequest struct {
	Limit  int64
	Offset int64
}
