// Code generated by sqlc. DO NOT EDIT.

package sqlc

import (
	"database/sql"
	"time"
)

type BannedIp struct {
	IpAddress string
	CreatedAt time.Time
}

type Log struct {
	ID        int64
	UserID    int64
	Slug      string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type LogEntry struct {
	ID           int64
	LogID        int64
	EntryDate    time.Time
	Coffee       string
	Water        sql.NullString
	CoffeeGrams  sql.NullInt32
	WaterGrams   sql.NullInt32
	BrewMethod   sql.NullString
	GrindNotes   sql.NullString
	TastingNotes sql.NullString
	AddlNotes    sql.NullString
	DeletedAt    sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type LoginAttempt struct {
	IpAddress string
	Attempts  int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SchemaMigration struct {
	Version string
}

type User struct {
	ID          int64
	DisplayName string
	Username    string
	Password    string
	TimeZone    sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
