package seeder

import (
	"context"
	"database/sql"
	"fmt"
	db "topup_game/pkg/database/schema"
	"topup_game/pkg/logger"

	"go.uber.org/zap"
)

type transactionSeeder struct {
	db     *db.Queries
	ctx    context.Context
	logger logger.LoggerInterface
}

func NewTransactionSeeder(db *db.Queries, ctx context.Context, logger logger.LoggerInterface) *transactionSeeder {
	return &transactionSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

func (r *transactionSeeder) Seed() error {
	transactions := []struct {
		UserID        int32
		MerchantID    sql.NullInt32
		VoucherID     sql.NullInt32
		NominalID     sql.NullInt32
		CategoryID    sql.NullInt32
		BankID        sql.NullInt32
		PaymentMethod string
		Status        sql.NullString
	}{
		{
			UserID:        1,
			MerchantID:    sql.NullInt32{Int32: 1, Valid: true},
			VoucherID:     sql.NullInt32{Int32: 1, Valid: true},
			NominalID:     sql.NullInt32{Int32: 1, Valid: true},
			CategoryID:    sql.NullInt32{Int32: 1, Valid: true},
			BankID:        sql.NullInt32{Int32: 1, Valid: true},
			PaymentMethod: "Credit Card",
			Status:        sql.NullString{String: "Completed", Valid: true},
		},
		{
			UserID:        2,
			MerchantID:    sql.NullInt32{Int32: 2, Valid: true},
			VoucherID:     sql.NullInt32{Int32: 2, Valid: true},
			NominalID:     sql.NullInt32{Int32: 2, Valid: true},
			CategoryID:    sql.NullInt32{Int32: 2, Valid: true},
			BankID:        sql.NullInt32{Int32: 2, Valid: true},
			PaymentMethod: "PayPal",
			Status:        sql.NullString{String: "Pending", Valid: true},
		},
		{
			UserID:        3,
			MerchantID:    sql.NullInt32{Int32: 3, Valid: true},
			VoucherID:     sql.NullInt32{Int32: 3, Valid: true},
			NominalID:     sql.NullInt32{Int32: 3, Valid: true},
			CategoryID:    sql.NullInt32{Int32: 3, Valid: true},
			BankID:        sql.NullInt32{Int32: 3, Valid: true},
			PaymentMethod: "Bank Transfer",
			Status:        sql.NullString{String: "Failed", Valid: true},
		},
	}

	totalTransactions := len(transactions)

	for i, transaction := range transactions {
		_, err := r.db.CreateTransaction(r.ctx, db.CreateTransactionParams{
			UserID:        transaction.UserID,
			MerchantID:    transaction.MerchantID,
			VoucherID:     transaction.VoucherID,
			NominalID:     transaction.NominalID,
			CategoryID:    transaction.CategoryID,
			BankID:        transaction.BankID,
			PaymentMethod: transaction.PaymentMethod,
			Status:        transaction.Status,
		})
		if err != nil {
			r.logger.Error("failed to seed transaction", zap.Int("index", i+1), zap.Int32("userID", transaction.UserID), zap.Error(err))
			return fmt.Errorf("failed to seed transaction %d (User ID: %d): %w", i+1, transaction.UserID, err)
		}
	}

	r.logger.Debug("transactions seeded successfully", zap.Int("totalTransactions", totalTransactions))
	return nil
}
