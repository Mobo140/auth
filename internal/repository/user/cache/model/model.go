package model

type UserAuthData struct {
	HashedPassword string `redis:"password_hash"`
	Role           int64  `redis:"role"`
}
