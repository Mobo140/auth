package model

import "time"

type LogEntry struct {
	UserID    int64
	Activity  string
	CreatedAt time.Time
}
