-- +goose Up
-- +goose StatementBegin
CREATE TABLE "nominals" (
  "nominal_id" SERIAL PRIMARY KEY,
  "voucher_id" INT NOT NULL REFERENCES "vouchers" ("voucher_id"),
  "name" VARCHAR(255) NOT NULL,
  "quantity" INT NOT NULL,
  "price" DECIMAL(10,2) NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_nominals_name ON "nominals" ("name");
DROP INDEX IF EXISTS idx_nominals_price ON "nominals" ("price");
DROP TABLE IF EXISTS "nominals";
-- +goose StatementEnd
