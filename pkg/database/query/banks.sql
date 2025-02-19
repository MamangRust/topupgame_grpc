-- Get Banks with Pagination and Total Count
-- name: GetBanks :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM banks
WHERE deleted_at IS NULL
AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Get Active Banks with Pagination and Total Count
-- name: GetBanksActive :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM banks
WHERE deleted_at IS NULL
AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Get Trashed Banks with Pagination and Total Count
-- name: GetBanksTrashed :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM banks
WHERE deleted_at IS NOT NULL
AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Create Bank
-- name: CreateBank :one
INSERT INTO banks (name)
VALUES ($1)
  RETURNING *;


-- Get Bank by ID
-- name: GetBankByID :one
SELECT bank_id, name, created_at, updated_at, deleted_at
FROM banks
WHERE bank_id = $1
  AND deleted_at IS NULL;


-- Update Bank
-- name: UpdateBank :one
UPDATE banks
SET name = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE bank_id = $1
  AND deleted_at IS NULL
  RETURNING *;


-- Trash Bank (Soft Delete)
-- name: TrashBank :one
UPDATE banks
SET deleted_at = CURRENT_TIMESTAMP
WHERE bank_id = $1
  AND deleted_at IS NULL
  RETURNING *;


-- Restore Trashed Bank
-- name: RestoreBank :one
UPDATE banks
SET deleted_at = NULL
WHERE bank_id = $1
  AND deleted_at IS NOT NULL
  RETURNING *;


-- Delete Bank Permanently
-- name: DeleteBankPermanently :exec
DELETE FROM banks WHERE bank_id = $1 AND deleted_at IS NOT NULL;


-- Restore All Trashed Banks
-- name: RestoreAllBanks :exec
UPDATE banks
SET deleted_at = NULL
WHERE deleted_at IS NOT NULL;


-- Delete All Trashed Banks Permanently
-- name: DeleteAllPermanentBanks :exec
DELETE FROM banks WHERE deleted_at IS NOT NULL;
