package seeder

import (
	"context"
	"fmt"
	db "topup_game/pkg/database/schema"
	"topup_game/pkg/logger"

	"go.uber.org/zap"
)

type categorySeeder struct {
	db     *db.Queries
	ctx    context.Context
	logger logger.LoggerInterface
}

func NewCategorySeeder(db *db.Queries, ctx context.Context, logger logger.LoggerInterface) *categorySeeder {
	return &categorySeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

func (r *categorySeeder) Seed() error {
	categories := []string{"Electronics", "Fashion", "Groceries", "Health & Beauty", "Home & Living"}

	totalCategories := len(categories)

	for i, categoryName := range categories {
		_, err := r.db.CreateCategory(r.ctx, categoryName)
		if err != nil {
			r.logger.Error("failed to seed category", zap.Int("index", i+1), zap.String("categoryName", categoryName), zap.Error(err))
			return fmt.Errorf("failed to seed category %d (%s): %w", i+1, categoryName, err)
		}
	}

	r.logger.Debug("categories seeded successfully", zap.Int("totalCategories", totalCategories))
	return nil
}
