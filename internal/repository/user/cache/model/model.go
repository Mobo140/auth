package model

type UserInfo struct {
	ID          int64  `redis:"id"`
	Name        string `redis:"name"`
	Email       string `redis:"email"`
	Role        int64  `redis:"role"`
	CreatedAtNs int64  `redis:"created_at"`
	UpdatedAtNs *int64 `redis:"updated_at"`
}

type User struct {
	Name           string
	Email          string
	HashedPassword string
	Role           int64
}
