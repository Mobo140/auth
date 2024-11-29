package model

import "time"

type LogEntryUser struct {
	UserID    int64
	Action    string
	CreatedAt time.Time
}

type LogEntryAuth struct {
	Username  string
	Action    string
	CreatedAt time.Time
}
