package seeder

import (
	"context"
	"fmt"
	db "topup_game/pkg/database/schema"
	"topup_game/pkg/logger"

	"go.uber.org/zap"
)

type bankSeeder struct {
	db     *db.Queries
	ctx    context.Context
	logger logger.LoggerInterface
}

func NewBankSeeder(db *db.Queries, ctx context.Context, logger logger.LoggerInterface) *bankSeeder {
	return &bankSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

func (r *bankSeeder) Seed() error {
	banks := []string{
		"Bank Nusantara",
		"Bank Aman Sentosa",
		"Bank Prima Sejahtera",
		"Bank Harmoni",
		"Bank Cemerlang",
		"Bank Sentral Nusantara",
		"Bank Fortuna",
		"Bank Mega Jaya",
		"Bank Sahabat Rakyat",
		"Bank Mandala",
		"Bank Dinamika",
		"Bank Global Trust",
		"Bank Prima Internasional",
		"Bank Andalan",
		"Bank Serba Guna",
	}

	totalBanks := len(banks)

	for i, bankName := range banks {
		_, err := r.db.CreateBank(r.ctx, bankName)
		if err != nil {
			r.logger.Error("failed to seed bank", zap.Int("index", i+1), zap.String("bankName", bankName), zap.Error(err))
			return fmt.Errorf("failed to seed bank %d (%s): %w", i+1, bankName, err)
		}
	}

	r.logger.Debug("banks seeded successfully", zap.Int("totalBanks", totalBanks))
	return nil
}
