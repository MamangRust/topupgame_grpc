-- +goose Up
-- +goose StatementBegin
CREATE TABLE "vouchers" (
  "voucher_id" SERIAL PRIMARY KEY,
  "merchant_id" INT NOT NULL REFERENCES "merchants" ("merchant_id"),
  "category_id" INT NOT NULL REFERENCES "categories" ("category_id"),
  "name" VARCHAR(255) NOT NULL,
  "image_name" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP DEFAULT NULL
);
CREATE INDEX idx_vouchers_merchant ON "vouchers" ("merchant_id");
CREATE INDEX idx_vouchers_category ON "vouchers" ("category_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_vouchers_merchant ON "vouchers" ("merchant_id");
DROP INDEX IF EXISTS idx_vouchers_category ON "vouchers" ("category_id");
DROP TABLE IF EXISTS "vouchers";
-- +goose StatementEnd