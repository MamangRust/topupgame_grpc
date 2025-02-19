-- Get Transactions with Pagination and Total Count
-- name: GetTransactions :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM transactions
WHERE deleted_at IS NULL
AND ($1::TEXT IS NULL OR payment_method ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Get Active Transactions with Pagination and Total Count
-- name: GetTransactionsActive :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM transactions
WHERE deleted_at IS NULL
AND ($1::TEXT IS NULL OR payment_method ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Get Trashed Transactions with Pagination and Total Count
-- name: GetTransactionsTrashed :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM transactions
WHERE deleted_at IS NOT NULL
AND ($1::TEXT IS NULL OR payment_method ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Create Transaction
-- name: CreateTransaction :one
INSERT INTO transactions (user_id, merchant_id, voucher_id, nominal_id, category_id, bank_id, payment_method, status)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
  RETURNING *;


-- Get Transaction by ID
-- name: GetTransactionByID :one
SELECT *
FROM transactions
WHERE transaction_id = $1
  AND deleted_at IS NULL;


-- Update Transaction
-- name: UpdateTransaction :one
UPDATE transactions
SET user_id = $2,
    merchant_id = $3,
    voucher_id = $4,
    nominal_id = $5,
    category_id = $6,
    bank_id = $7,
    payment_method = $8,
    status = $9,
    updated_at = CURRENT_TIMESTAMP
WHERE transaction_id = $1
  AND deleted_at IS NULL
  RETURNING *;



-- Update Transaction Status
-- name: UpdateTransactionStatus :exec
UPDATE transactions
SET 
    status = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE transaction_id = $1
  AND deleted_at IS NULL;




-- Trash Transaction (Soft Delete)
-- name: TrashTransaction :one
UPDATE transactions
SET deleted_at = CURRENT_TIMESTAMP
WHERE transaction_id = $1
  AND deleted_at IS NULL
  RETURNING *;


-- Restore Trashed Transaction
-- name: RestoreTransaction :one
UPDATE transactions
SET deleted_at = NULL
WHERE transaction_id = $1
  AND deleted_at IS NOT NULL
  RETURNING *;


-- Delete Transaction Permanently
-- name: DeleteTransactionPermanently :exec
DELETE FROM transactions WHERE transaction_id = $1 AND deleted_at IS NOT NULL;


-- Restore All Trashed Transactions
-- name: RestoreAllTransactions :exec
UPDATE transactions
SET deleted_at = NULL
WHERE deleted_at IS NOT NULL;


-- Delete All Trashed Transactions Permanently
-- name: DeleteAllPermanentTransactions :exec
DELETE FROM transactions WHERE deleted_at IS NOT NULL;
