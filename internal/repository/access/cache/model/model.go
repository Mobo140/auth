package model

type AccessEndpoint struct {
	Endpoint string `redis:"endpoint"`
	Role     int64  `redis:"role"`
}
