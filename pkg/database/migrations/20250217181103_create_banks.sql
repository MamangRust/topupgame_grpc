-- +goose Up
-- +goose StatementBegin
CREATE TABLE "banks" (
  "bank_id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP DEFAULT NULL
);
CREATE INDEX idx_banks_name ON "banks" ("name");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_banks_name ON "banks" ("name");
DROP TABLE IF EXISTS "banks";
-- +goose StatementEnd