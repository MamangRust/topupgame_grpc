-- Get Nominals with Pagination and Total Count
-- name: GetNominals :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM nominals
WHERE deleted_at IS NULL
AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Get Active Nominals with Pagination and Total Count
-- name: GetNominalsActive :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM nominals
WHERE deleted_at IS NULL
AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Get Trashed Nominals with Pagination and Total Count
-- name: GetNominalsTrashed :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM nominals
WHERE deleted_at IS NOT NULL
AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Create Nominal
-- name: CreateNominal :one
INSERT INTO nominals (name, quantity, price, voucher_id)
VALUES ($1, $2, $3, $4)
  RETURNING *;


-- Get Nominal by ID
-- name: GetNominalByID :one
SELECT *
FROM nominals
WHERE nominal_id = $1
  AND deleted_at IS NULL;


-- Update Nominal
-- name: UpdateNominal :one
UPDATE nominals
SET name = $2,
    quantity = $3,
    price = $4,
    voucher_id = $5,
    updated_at = CURRENT_TIMESTAMP
WHERE nominal_id = $1
  AND deleted_at IS NULL
  RETURNING *;


-- name: DecreaseNominalQuantity :exec
UPDATE nominals 
SET quantity = quantity - $1 
WHERE nominal_id = $2 AND quantity >= $1;



-- Trash Nominal (Soft Delete)
-- name: TrashNominal :one
UPDATE nominals
SET deleted_at = CURRENT_TIMESTAMP
WHERE nominal_id = $1
  AND deleted_at IS NULL
  RETURNING *;


-- Restore Trashed Nominal
-- name: RestoreNominal :one
UPDATE nominals
SET deleted_at = NULL
WHERE nominal_id = $1
  AND deleted_at IS NOT NULL
  RETURNING *;


-- Delete Nominal Permanently
-- name: DeleteNominalPermanently :exec
DELETE FROM nominals WHERE nominal_id = $1 AND deleted_at IS NOT NULL;


-- Restore All Trashed Nominals
-- name: RestoreAllNominals :exec
UPDATE nominals
SET deleted_at = NULL
WHERE deleted_at IS NOT NULL;


-- Delete All Trashed Nominals Permanently
-- name: DeleteAllPermanentNominals :exec
DELETE FROM nominals WHERE deleted_at IS NOT NULL;
