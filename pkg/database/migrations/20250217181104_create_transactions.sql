-- +goose Up
-- +goose StatementBegin
CREATE TABLE "transactions" (
  "transaction_id" SERIAL PRIMARY KEY,
  "user_id" INT NOT NULL REFERENCES "users" ("user_id") ON DELETE CASCADE,
  "merchant_id" INT REFERENCES "merchants" ("merchant_id") ON DELETE SET NULL,
  "voucher_id" INT REFERENCES "vouchers" ("voucher_id") ON DELETE SET NULL,
  "nominal_id" INT REFERENCES "nominals" ("nominal_id") ON DELETE SET NULL,
  "bank_id" INT REFERENCES "banks" ("bank_id") ON DELETE SET NULL,
  "payment_method" VARCHAR(255) NOT NULL,
  "amount" INT NOT NULL,
  "status" VARCHAR(20) DEFAULT 'paying',
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP DEFAULT NULL
);

CREATE INDEX idx_transactions_merchant ON "transactions" ("merchant_id");
CREATE INDEX idx_transactions_voucher ON "transactions" ("voucher_id");
CREATE INDEX idx_transactions_nominal ON "transactions" ("nominal_id");
CREATE INDEX idx_transactions_bank ON "transactions" ("bank_id");
CREATE INDEX idx_transactions_status ON "transactions" ("status");
CREATE INDEX idx_transactions_created_at ON "transactions" ("created_at");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_transactions_merchant ON "transactions" ("merchant_id");
DROP INDEX IF EXISTS idx_transactions_voucher ON "transactions" ("voucher_id");
DROP INDEX IF EXISTS idx_transactions_nominal ON "transactions" ("nominal_id");
DROP INDEX IF EXISTS idx_transactions_bank ON "transactions" ("bank_id");
DROP INDEX IF EXISTS idx_transactions_status ON "transactions" ("status");
DROP INDEX IF EXISTS idx_transactions_created_at ON "transactions" ("created_at");
DROP TABLE IF EXISTS "transactions";
-- +goose StatementEnd