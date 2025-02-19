-- +goose Up
-- +goose StatementBegin
CREATE TABLE "categories" (
  "category_id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP DEFAULT NULL
);
CREATE INDEX "idx_categories_name" ON "categories" ("name");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF "idx_categories_name" ON "categories";

DROP TABLE IF EXISTS "categories";
-- +goose StatementEnd