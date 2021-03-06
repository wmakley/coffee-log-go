// Code generated by sqlc. DO NOT EDIT.
// source: log_entries.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const createLogEntry = `-- name: CreateLogEntry :one
INSERT INTO log_entries
(log_id, entry_date, coffee, water, brew_method, grind_notes, tasting_notes, addl_notes, coffee_grams, water_grams)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING id, log_id, entry_date, coffee, water, coffee_grams, water_grams, brew_method, grind_notes, tasting_notes, addl_notes, deleted_at, created_at, updated_at
`

type CreateLogEntryParams struct {
	LogID        int64
	EntryDate    time.Time
	Coffee       string
	Water        sql.NullString
	BrewMethod   sql.NullString
	GrindNotes   sql.NullString
	TastingNotes sql.NullString
	AddlNotes    sql.NullString
	CoffeeGrams  sql.NullInt32
	WaterGrams   sql.NullInt32
}

func (q *Queries) CreateLogEntry(ctx context.Context, arg CreateLogEntryParams) (LogEntry, error) {
	row := q.db.QueryRowContext(ctx, createLogEntry,
		arg.LogID,
		arg.EntryDate,
		arg.Coffee,
		arg.Water,
		arg.BrewMethod,
		arg.GrindNotes,
		arg.TastingNotes,
		arg.AddlNotes,
		arg.CoffeeGrams,
		arg.WaterGrams,
	)
	var i LogEntry
	err := row.Scan(
		&i.ID,
		&i.LogID,
		&i.EntryDate,
		&i.Coffee,
		&i.Water,
		&i.CoffeeGrams,
		&i.WaterGrams,
		&i.BrewMethod,
		&i.GrindNotes,
		&i.TastingNotes,
		&i.AddlNotes,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listLogEntriesByLogIDOrderByDateDesc = `-- name: ListLogEntriesByLogIDOrderByDateDesc :many
SELECT id, log_id, entry_date, coffee, water, coffee_grams, water_grams, brew_method, grind_notes, tasting_notes, addl_notes, deleted_at, created_at, updated_at FROM log_entries
WHERE deleted_at IS NOT NULL AND log_id = $1
ORDER BY entry_date DESC
`

func (q *Queries) ListLogEntriesByLogIDOrderByDateDesc(ctx context.Context, logID int64) ([]LogEntry, error) {
	rows, err := q.db.QueryContext(ctx, listLogEntriesByLogIDOrderByDateDesc, logID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []LogEntry{}
	for rows.Next() {
		var i LogEntry
		if err := rows.Scan(
			&i.ID,
			&i.LogID,
			&i.EntryDate,
			&i.Coffee,
			&i.Water,
			&i.CoffeeGrams,
			&i.WaterGrams,
			&i.BrewMethod,
			&i.GrindNotes,
			&i.TastingNotes,
			&i.AddlNotes,
			&i.DeletedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateLogEntry = `-- name: UpdateLogEntry :one
UPDATE log_entries
SET entry_date = $2,
	coffee = $3,
	water = $4,
	brew_method = $5,
	grind_notes = $6,
	tasting_notes = $7,
	addl_notes = $8,
	coffee_grams = $9,
	water_grams = $10,
	updated_at = $11
WHERE id = $1
RETURNING id, log_id, entry_date, coffee, water, coffee_grams, water_grams, brew_method, grind_notes, tasting_notes, addl_notes, deleted_at, created_at, updated_at
`

type UpdateLogEntryParams struct {
	ID           int64
	EntryDate    time.Time
	Coffee       string
	Water        sql.NullString
	BrewMethod   sql.NullString
	GrindNotes   sql.NullString
	TastingNotes sql.NullString
	AddlNotes    sql.NullString
	CoffeeGrams  sql.NullInt32
	WaterGrams   sql.NullInt32
	UpdatedAt    time.Time
}

func (q *Queries) UpdateLogEntry(ctx context.Context, arg UpdateLogEntryParams) (LogEntry, error) {
	row := q.db.QueryRowContext(ctx, updateLogEntry,
		arg.ID,
		arg.EntryDate,
		arg.Coffee,
		arg.Water,
		arg.BrewMethod,
		arg.GrindNotes,
		arg.TastingNotes,
		arg.AddlNotes,
		arg.CoffeeGrams,
		arg.WaterGrams,
		arg.UpdatedAt,
	)
	var i LogEntry
	err := row.Scan(
		&i.ID,
		&i.LogID,
		&i.EntryDate,
		&i.Coffee,
		&i.Water,
		&i.CoffeeGrams,
		&i.WaterGrams,
		&i.BrewMethod,
		&i.GrindNotes,
		&i.TastingNotes,
		&i.AddlNotes,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
