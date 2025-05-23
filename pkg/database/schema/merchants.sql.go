// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: merchants.sql

package db

import (
	"context"
	"database/sql"
)

const createMerchant = `-- name: CreateMerchant :one
INSERT INTO merchants (user_id, name, description, address, contact_email, contact_phone, status)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING merchant_id, user_id, name, description, address, contact_email, contact_phone, status, created_at, updated_at, deleted_at
`

type CreateMerchantParams struct {
	UserID       int32          `json:"user_id"`
	Name         string         `json:"name"`
	Description  sql.NullString `json:"description"`
	Address      sql.NullString `json:"address"`
	ContactEmail sql.NullString `json:"contact_email"`
	ContactPhone sql.NullString `json:"contact_phone"`
	Status       string         `json:"status"`
}

// Create Merchant
func (q *Queries) CreateMerchant(ctx context.Context, arg CreateMerchantParams) (*Merchant, error) {
	row := q.db.QueryRowContext(ctx, createMerchant,
		arg.UserID,
		arg.Name,
		arg.Description,
		arg.Address,
		arg.ContactEmail,
		arg.ContactPhone,
		arg.Status,
	)
	var i Merchant
	err := row.Scan(
		&i.MerchantID,
		&i.UserID,
		&i.Name,
		&i.Description,
		&i.Address,
		&i.ContactEmail,
		&i.ContactPhone,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const deleteAllPermanentMerchants = `-- name: DeleteAllPermanentMerchants :exec
DELETE FROM merchants
WHERE
    deleted_at IS NOT NULL
`

// Delete All Trashed Merchant Permanently
func (q *Queries) DeleteAllPermanentMerchants(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllPermanentMerchants)
	return err
}

const deleteMerchantPermanently = `-- name: DeleteMerchantPermanently :exec
DELETE FROM merchants WHERE merchant_id = $1 AND deleted_at IS NOT NULL
`

// Delete Merchant Permanently
func (q *Queries) DeleteMerchantPermanently(ctx context.Context, merchantID int32) error {
	_, err := q.db.ExecContext(ctx, deleteMerchantPermanently, merchantID)
	return err
}

const getMerchantByID = `-- name: GetMerchantByID :one
SELECT merchant_id, user_id, name, description, address, contact_email, contact_phone, status, created_at, updated_at, deleted_at
FROM merchants
WHERE merchant_id = $1
  AND deleted_at IS NULL
`

func (q *Queries) GetMerchantByID(ctx context.Context, merchantID int32) (*Merchant, error) {
	row := q.db.QueryRowContext(ctx, getMerchantByID, merchantID)
	var i Merchant
	err := row.Scan(
		&i.MerchantID,
		&i.UserID,
		&i.Name,
		&i.Description,
		&i.Address,
		&i.ContactEmail,
		&i.ContactPhone,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getMerchants = `-- name: GetMerchants :many
SELECT
    merchant_id, user_id, name, description, address, contact_email, contact_phone, status, created_at, updated_at, deleted_at,
    COUNT(*) OVER() AS total_count
FROM merchants
WHERE deleted_at IS NULL
AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%' OR contact_email ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetMerchantsParams struct {
	Column1 string `json:"column_1"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

type GetMerchantsRow struct {
	MerchantID   int32          `json:"merchant_id"`
	UserID       int32          `json:"user_id"`
	Name         string         `json:"name"`
	Description  sql.NullString `json:"description"`
	Address      sql.NullString `json:"address"`
	ContactEmail sql.NullString `json:"contact_email"`
	ContactPhone sql.NullString `json:"contact_phone"`
	Status       string         `json:"status"`
	CreatedAt    sql.NullTime   `json:"created_at"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
	DeletedAt    sql.NullTime   `json:"deleted_at"`
	TotalCount   int64          `json:"total_count"`
}

func (q *Queries) GetMerchants(ctx context.Context, arg GetMerchantsParams) ([]*GetMerchantsRow, error) {
	rows, err := q.db.QueryContext(ctx, getMerchants, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetMerchantsRow
	for rows.Next() {
		var i GetMerchantsRow
		if err := rows.Scan(
			&i.MerchantID,
			&i.UserID,
			&i.Name,
			&i.Description,
			&i.Address,
			&i.ContactEmail,
			&i.ContactPhone,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.TotalCount,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMerchantsActive = `-- name: GetMerchantsActive :many
SELECT
    merchant_id, user_id, name, description, address, contact_email, contact_phone, status, created_at, updated_at, deleted_at,
    COUNT(*) OVER() AS total_count
FROM merchants
WHERE deleted_at IS NULL
AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%' OR contact_email ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetMerchantsActiveParams struct {
	Column1 string `json:"column_1"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

type GetMerchantsActiveRow struct {
	MerchantID   int32          `json:"merchant_id"`
	UserID       int32          `json:"user_id"`
	Name         string         `json:"name"`
	Description  sql.NullString `json:"description"`
	Address      sql.NullString `json:"address"`
	ContactEmail sql.NullString `json:"contact_email"`
	ContactPhone sql.NullString `json:"contact_phone"`
	Status       string         `json:"status"`
	CreatedAt    sql.NullTime   `json:"created_at"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
	DeletedAt    sql.NullTime   `json:"deleted_at"`
	TotalCount   int64          `json:"total_count"`
}

// Get Active Merchants with Pagination and Total Count
func (q *Queries) GetMerchantsActive(ctx context.Context, arg GetMerchantsActiveParams) ([]*GetMerchantsActiveRow, error) {
	rows, err := q.db.QueryContext(ctx, getMerchantsActive, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetMerchantsActiveRow
	for rows.Next() {
		var i GetMerchantsActiveRow
		if err := rows.Scan(
			&i.MerchantID,
			&i.UserID,
			&i.Name,
			&i.Description,
			&i.Address,
			&i.ContactEmail,
			&i.ContactPhone,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.TotalCount,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMerchantsTrashed = `-- name: GetMerchantsTrashed :many
SELECT
    merchant_id, user_id, name, description, address, contact_email, contact_phone, status, created_at, updated_at, deleted_at,
    COUNT(*) OVER() AS total_count
FROM merchants
WHERE deleted_at IS NOT NULL
AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%' OR contact_email ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetMerchantsTrashedParams struct {
	Column1 string `json:"column_1"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

type GetMerchantsTrashedRow struct {
	MerchantID   int32          `json:"merchant_id"`
	UserID       int32          `json:"user_id"`
	Name         string         `json:"name"`
	Description  sql.NullString `json:"description"`
	Address      sql.NullString `json:"address"`
	ContactEmail sql.NullString `json:"contact_email"`
	ContactPhone sql.NullString `json:"contact_phone"`
	Status       string         `json:"status"`
	CreatedAt    sql.NullTime   `json:"created_at"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
	DeletedAt    sql.NullTime   `json:"deleted_at"`
	TotalCount   int64          `json:"total_count"`
}

// Get Trashed Merchants with Pagination and Total Count
func (q *Queries) GetMerchantsTrashed(ctx context.Context, arg GetMerchantsTrashedParams) ([]*GetMerchantsTrashedRow, error) {
	rows, err := q.db.QueryContext(ctx, getMerchantsTrashed, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetMerchantsTrashedRow
	for rows.Next() {
		var i GetMerchantsTrashedRow
		if err := rows.Scan(
			&i.MerchantID,
			&i.UserID,
			&i.Name,
			&i.Description,
			&i.Address,
			&i.ContactEmail,
			&i.ContactPhone,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.TotalCount,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const restoreAllMerchants = `-- name: RestoreAllMerchants :exec
UPDATE merchants
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL
`

// Restore All Trashed Merchant
func (q *Queries) RestoreAllMerchants(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, restoreAllMerchants)
	return err
}

const restoreMerchant = `-- name: RestoreMerchant :one
UPDATE merchants
SET
    deleted_at = NULL
WHERE
    merchant_id = $1
    AND deleted_at IS NOT NULL
  RETURNING merchant_id, user_id, name, description, address, contact_email, contact_phone, status, created_at, updated_at, deleted_at
`

// Restore Trashed Merchant
func (q *Queries) RestoreMerchant(ctx context.Context, merchantID int32) (*Merchant, error) {
	row := q.db.QueryRowContext(ctx, restoreMerchant, merchantID)
	var i Merchant
	err := row.Scan(
		&i.MerchantID,
		&i.UserID,
		&i.Name,
		&i.Description,
		&i.Address,
		&i.ContactEmail,
		&i.ContactPhone,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const trashMerchant = `-- name: TrashMerchant :one
UPDATE merchants
SET
    deleted_at = current_timestamp
WHERE
    merchant_id = $1
    AND deleted_at IS NULL
    RETURNING merchant_id, user_id, name, description, address, contact_email, contact_phone, status, created_at, updated_at, deleted_at
`

// Trash Merchant
func (q *Queries) TrashMerchant(ctx context.Context, merchantID int32) (*Merchant, error) {
	row := q.db.QueryRowContext(ctx, trashMerchant, merchantID)
	var i Merchant
	err := row.Scan(
		&i.MerchantID,
		&i.UserID,
		&i.Name,
		&i.Description,
		&i.Address,
		&i.ContactEmail,
		&i.ContactPhone,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const updateMerchant = `-- name: UpdateMerchant :one
UPDATE merchants
SET name = $2,
    description = $3,
    address = $4,
    contact_email = $5,
    contact_phone = $6,
    status = $7,
    updated_at = CURRENT_TIMESTAMP
WHERE merchant_id = $1
  AND deleted_at IS NULL
  RETURNING merchant_id, user_id, name, description, address, contact_email, contact_phone, status, created_at, updated_at, deleted_at
`

type UpdateMerchantParams struct {
	MerchantID   int32          `json:"merchant_id"`
	Name         string         `json:"name"`
	Description  sql.NullString `json:"description"`
	Address      sql.NullString `json:"address"`
	ContactEmail sql.NullString `json:"contact_email"`
	ContactPhone sql.NullString `json:"contact_phone"`
	Status       string         `json:"status"`
}

// Update Merchant
func (q *Queries) UpdateMerchant(ctx context.Context, arg UpdateMerchantParams) (*Merchant, error) {
	row := q.db.QueryRowContext(ctx, updateMerchant,
		arg.MerchantID,
		arg.Name,
		arg.Description,
		arg.Address,
		arg.ContactEmail,
		arg.ContactPhone,
		arg.Status,
	)
	var i Merchant
	err := row.Scan(
		&i.MerchantID,
		&i.UserID,
		&i.Name,
		&i.Description,
		&i.Address,
		&i.ContactEmail,
		&i.ContactPhone,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}
