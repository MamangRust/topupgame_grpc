package seeder

import (
	"context"
	"fmt"
	db "topup_game/pkg/database/schema"
	"topup_game/pkg/logger"

	"go.uber.org/zap"
)

type nominalSeeder struct {
	db     *db.Queries
	ctx    context.Context
	logger logger.LoggerInterface
}

func NewNominalSeeder(db *db.Queries, ctx context.Context, logger logger.LoggerInterface) *nominalSeeder {
	return &nominalSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

func (r *nominalSeeder) Seed() error {
	nominals := []struct {
		VoucherID int32
		Name      string
		Quantity  int32
		Price     float64
	}{
		{VoucherID: 1, Name: "Small Package", Quantity: 10, Price: 5000},
		{VoucherID: 1, Name: "Medium Package", Quantity: 20, Price: 10000},
		{VoucherID: 2, Name: "Large Package", Quantity: 50, Price: 25000},
		{VoucherID: 3, Name: "VIP Package", Quantity: 100, Price: 50000},
	}

	totalNominals := len(nominals)

	for i, nominal := range nominals {
		_, err := r.db.CreateNominal(r.ctx, db.CreateNominalParams{
			VoucherID: nominal.VoucherID,
			Name:      nominal.Name,
			Quantity:  nominal.Quantity,
			Price:     nominal.Price,
		})
		if err != nil {
			r.logger.Error("failed to seed nominal", zap.Int("index", i+1), zap.String("name", nominal.Name), zap.Error(err))
			return fmt.Errorf("failed to seed nominal %d (%s): %w", i+1, nominal.Name, err)
		}
	}

	r.logger.Debug("nominals seeded successfully", zap.Int("totalNominals", totalNominals))
	return nil
}
