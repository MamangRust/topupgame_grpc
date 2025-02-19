-- Get Categories with Pagination and Total Count
-- name: GetCategories :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM categories
WHERE deleted_at IS NULL
AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Get Active Categories with Pagination and Total Count
-- name: GetCategoriesActive :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM categories
WHERE deleted_at IS NULL
AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Get Trashed Categories with Pagination and Total Count
-- name: GetCategoriesTrashed :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM categories
WHERE deleted_at IS NOT NULL
AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Create Category
-- name: CreateCategory :one
INSERT INTO categories (name)
VALUES ($1)
  RETURNING *;


-- Get Category by ID
-- name: GetCategoryByID :one
SELECT category_id, name, created_at, updated_at, deleted_at
FROM categories
WHERE category_id = $1
  AND deleted_at IS NULL;


-- Update Category
-- name: UpdateCategory :one
UPDATE categories
SET name = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE category_id = $1
  AND deleted_at IS NULL
  RETURNING *;


-- Trash Category (Soft Delete)
-- name: TrashCategory :one
UPDATE categories
SET deleted_at = CURRENT_TIMESTAMP
WHERE category_id = $1
  AND deleted_at IS NULL
  RETURNING *;


-- Restore Trashed Category
-- name: RestoreCategory :one
UPDATE categories
SET deleted_at = NULL
WHERE category_id = $1
  AND deleted_at IS NOT NULL
  RETURNING *;


-- Delete Category Permanently
-- name: DeleteCategoryPermanently :exec
DELETE FROM categories WHERE category_id = $1 AND deleted_at IS NOT NULL;


-- Restore All Trashed Categories
-- name: RestoreAllCategories :exec
UPDATE categories
SET deleted_at = NULL
WHERE deleted_at IS NOT NULL;


-- Delete All Trashed Categories Permanently
-- name: DeleteAllPermanentCategories :exec
DELETE FROM categories WHERE deleted_at IS NOT NULL;
