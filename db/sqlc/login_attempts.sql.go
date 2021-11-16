// Code generated by sqlc. DO NOT EDIT.
// source: login_attempts.sql

package sqlc

import (
	"context"
)

const createBannedIP = `-- name: CreateBannedIP :one
INSERT INTO banned_ips (ip_address, created_at)
VALUES ($1, NOW())
RETURNING ip_address, created_at
`

func (q *Queries) CreateBannedIP(ctx context.Context, ipAddress string) (BannedIp, error) {
	row := q.db.QueryRowContext(ctx, createBannedIP, ipAddress)
	var i BannedIp
	err := row.Scan(&i.IpAddress, &i.CreatedAt)
	return i, err
}

const createLoginAttempt = `-- name: CreateLoginAttempt :one
INSERT INTO login_attempts (ip_address, attempts, created_at, updated_at)
VALUES ($1, 1, NOW(), NOW())
RETURNING ip_address, attempts, created_at, updated_at
`

func (q *Queries) CreateLoginAttempt(ctx context.Context, ipAddress string) (LoginAttempt, error) {
	row := q.db.QueryRowContext(ctx, createLoginAttempt, ipAddress)
	var i LoginAttempt
	err := row.Scan(
		&i.IpAddress,
		&i.Attempts,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAllBannedIPs = `-- name: DeleteAllBannedIPs :exec
DELETE FROM banned_ips
`

func (q *Queries) DeleteAllBannedIPs(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllBannedIPs)
	return err
}

const deleteAllLoginAttempts = `-- name: DeleteAllLoginAttempts :exec
DELETE FROM login_attempts
`

func (q *Queries) DeleteAllLoginAttempts(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllLoginAttempts)
	return err
}

const getBannedIP = `-- name: GetBannedIP :one
SELECT ip_address, created_at FROM banned_ips
WHERE ip_address = $1
LIMIT 1
`

func (q *Queries) GetBannedIP(ctx context.Context, ipAddress string) (BannedIp, error) {
	row := q.db.QueryRowContext(ctx, getBannedIP, ipAddress)
	var i BannedIp
	err := row.Scan(&i.IpAddress, &i.CreatedAt)
	return i, err
}

const getLoginAttempt = `-- name: GetLoginAttempt :one
SELECT ip_address, attempts, created_at, updated_at FROM login_attempts
WHERE ip_address = $1 LIMIT 1
`

func (q *Queries) GetLoginAttempt(ctx context.Context, ipAddress string) (LoginAttempt, error) {
	row := q.db.QueryRowContext(ctx, getLoginAttempt, ipAddress)
	var i LoginAttempt
	err := row.Scan(
		&i.IpAddress,
		&i.Attempts,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const incrementLoginAttempt = `-- name: IncrementLoginAttempt :one
UPDATE login_attempts
SET attempts = attempts + 1, updated_at = NOW()
WHERE ip_address = $1
RETURNING ip_address, attempts, created_at, updated_at
`

func (q *Queries) IncrementLoginAttempt(ctx context.Context, ipAddress string) (LoginAttempt, error) {
	row := q.db.QueryRowContext(ctx, incrementLoginAttempt, ipAddress)
	var i LoginAttempt
	err := row.Scan(
		&i.IpAddress,
		&i.Attempts,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}