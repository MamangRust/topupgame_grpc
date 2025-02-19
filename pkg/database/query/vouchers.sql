-- Get Vouchers with Pagination and Total Count
-- name: GetVouchers :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM vouchers
WHERE deleted_at IS NULL
AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Get Active Vouchers with Pagination and Total Count
-- name: GetVouchersActive :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM vouchers
WHERE deleted_at IS NULL
AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Get Trashed Vouchers with Pagination and Total Count
-- name: GetVouchersTrashed :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM vouchers
WHERE deleted_at IS NOT NULL
AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Create Voucher
-- name: CreateVoucher :one
INSERT INTO vouchers (merchant_id, category_id, name, image_name)
VALUES ($1, $2, $3, $4)
  RETURNING *;


-- Get Voucher by ID
-- name: GetVoucherByID :one
SELECT voucher_id, merchant_id, category_id, name, image_name, created_at, updated_at, deleted_at
FROM vouchers
WHERE voucher_id = $1
  AND deleted_at IS NULL;


-- Update Voucher
-- name: UpdateVoucher :one
UPDATE vouchers
SET name = $2,
    image_name = $3,
    category_id = $4,
    updated_at = CURRENT_TIMESTAMP
WHERE voucher_id = $1
  AND deleted_at IS NULL
  RETURNING *;




-- Trash Voucher (Soft Delete)
-- name: TrashVoucher :one
UPDATE vouchers
SET deleted_at = CURRENT_TIMESTAMP
WHERE voucher_id = $1
  AND deleted_at IS NULL
  RETURNING *;


-- Restore Trashed Voucher
-- name: RestoreVoucher :one
UPDATE vouchers
SET deleted_at = NULL
WHERE voucher_id = $1
  AND deleted_at IS NOT NULL
  RETURNING *;


-- Delete Voucher Permanently
-- name: DeleteVoucherPermanently :exec
DELETE FROM vouchers WHERE voucher_id = $1 AND deleted_at IS NOT NULL;


-- Restore All Trashed Vouchers
-- name: RestoreAllVouchers :exec
UPDATE vouchers
SET deleted_at = NULL
WHERE deleted_at IS NOT NULL;


-- Delete All Trashed Vouchers Permanently
-- name: DeleteAllPermanentVouchers :exec
DELETE FROM vouchers WHERE deleted_at IS NOT NULL;
