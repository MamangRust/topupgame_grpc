package seeder

import (
	"context"
	"fmt"
	"time"
	db "topup_game/pkg/database/schema"
	"topup_game/pkg/hash"
	"topup_game/pkg/logger"
)

type Deps struct {
	Db     *db.Queries
	Ctx    context.Context
	Logger logger.LoggerInterface
	Hash   hash.HashPassword
}

type Seeder struct {
	User        *userSeeder
	Role        *roleSeeder
	UserRole    *userRoleSeeder
	Bank        *bankSeeder
	Merchant    *merchantSeeder
	Category    *categorySeeder
	Nominal     *nominalSeeder
	Transaction *transactionSeeder
	Voucher     *voucherSeeder
}

func NewSeeder(deps Deps) *Seeder {
	return &Seeder{
		User:        NewUserSeeder(deps.Db, deps.Hash, deps.Ctx, deps.Logger),
		Role:        NewRoleSeeder(deps.Db, deps.Ctx, deps.Logger),
		UserRole:    NewUserRoleSeeder(deps.Db, deps.Ctx, deps.Logger),
		Merchant:    NewMerchantSeeder(deps.Db, deps.Ctx, deps.Logger),
		Bank:        NewBankSeeder(deps.Db, deps.Ctx, deps.Logger),
		Category:    NewCategorySeeder(deps.Db, deps.Ctx, deps.Logger),
		Nominal:     NewNominalSeeder(deps.Db, deps.Ctx, deps.Logger),
		Transaction: NewTransactionSeeder(deps.Db, deps.Ctx, deps.Logger),
		Voucher:     NewVoucherSeeder(deps.Db, deps.Ctx, deps.Logger),
	}
}

func (s *Seeder) Run() error {
	if err := s.seedWithDelay("users", s.User.Seed); err != nil {
		return err
	}

	if err := s.seedWithDelay("roles", s.Role.Seed); err != nil {
		return err
	}

	if err := s.seedWithDelay("user_roles", s.UserRole.Seed); err != nil {
		return err
	}

	if err := s.seedWithDelay("merchant", s.Merchant.Seed); err != nil {
		return err
	}

	if err := s.seedWithDelay("bank", s.Bank.Seed); err != nil {
		return err
	}

	if err := s.seedWithDelay("category", s.Category.Seed); err != nil {
		return err
	}

	if err := s.seedWithDelay("voucher", s.Voucher.Seed); err != nil {
		return err
	}

	if err := s.seedWithDelay("nominal", s.Nominal.Seed); err != nil {
		return err
	}

	if err := s.seedWithDelay("transaction", s.Transaction.Seed); err != nil {
		return err
	}

	return nil
}

func (s *Seeder) seedWithDelay(entityName string, seedFunc func() error) error {
	if err := seedFunc(); err != nil {
		return fmt.Errorf("failed to seed %s: %w", entityName, err)
	}

	time.Sleep(30 * time.Second)
	return nil
}
