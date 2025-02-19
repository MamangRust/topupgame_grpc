package seeder

import (
	"context"
	"fmt"
	db "topup_game/pkg/database/schema"
	"topup_game/pkg/logger"

	"go.uber.org/zap"
)

type voucherSeeder struct {
	db     *db.Queries
	ctx    context.Context
	logger logger.LoggerInterface
}

func NewVoucherSeeder(db *db.Queries, ctx context.Context, logger logger.LoggerInterface) *voucherSeeder {
	return &voucherSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

func (r *voucherSeeder) Seed() error {
	vouchers := []struct {
		MerchantID int32
		CategoryID int32
		Name       string
		ImageName  string
	}{
		{MerchantID: 2, CategoryID: 2, Name: "Diamond Bonus 50%", ImageName: "diamond_bonus.png"},
		{MerchantID: 2, CategoryID: 2, Name: "Exclusive Skin", ImageName: "exclusive_skin.png"},
		{MerchantID: 2, CategoryID: 2, Name: "Extra Battle Points", ImageName: "battle_points.png"},
	}

	totalVouchers := len(vouchers)

	for i, voucher := range vouchers {
		_, err := r.db.CreateVoucher(r.ctx, db.CreateVoucherParams{
			MerchantID: voucher.MerchantID,
			CategoryID: voucher.CategoryID,
			Name:       voucher.Name,
			ImageName:  voucher.ImageName,
		})
		if err != nil {
			r.logger.Error("failed to seed voucher", zap.Int("index", i+1), zap.String("name", voucher.Name), zap.Error(err))
			return fmt.Errorf("failed to seed voucher %d (%s): %w", i+1, voucher.Name, err)
		}
	}

	r.logger.Debug("vouchers seeded successfully", zap.Int("totalVouchers", totalVouchers))
	return nil
}
