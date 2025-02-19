package repository

import (
	"context"
	recordmapper "topup_game/internal/mapper/record"
	db "topup_game/pkg/database/schema"
)

type Repositories struct {
	User         UserRepository
	Role         RoleRepository
	UserRole     UserRoleRepository
	RefreshToken RefreshTokenRepository
	Bank         BankRepository
	Category     CategoryRepository
	Merchant     MerchantRepository
	Nominal      NominalRepository
	Transaction  TransactionRepository
	Voucher      VoucherRepository
}

type Deps struct {
	DB           *db.Queries
	Ctx          context.Context
	MapperRecord *recordmapper.RecordMapper
}

func NewRepositories(deps Deps) *Repositories {
	return &Repositories{
		User:         NewUserRepository(deps.DB, deps.Ctx, deps.MapperRecord.UserRecordMapper),
		Role:         NewRoleRepository(deps.DB, deps.Ctx, deps.MapperRecord.RoleRecordMapper),
		UserRole:     NewUserRoleRepository(deps.DB, deps.Ctx, deps.MapperRecord.UserRoleRecordMapper),
		RefreshToken: NewRefreshTokenRepository(deps.DB, deps.Ctx, deps.MapperRecord.RefreshTokenRecordMapper),
		Bank:         NewBankRepository(deps.DB, deps.Ctx, deps.MapperRecord.BankRecordMapper),
		Category:     NewCategoryRepository(deps.DB, deps.Ctx, deps.MapperRecord.CategoryRecordMapper),
		Merchant:     NewMerchantRepository(deps.DB, deps.Ctx, deps.MapperRecord.MerchantRecordMapper),
		Nominal:      NewNominalRepository(deps.DB, deps.Ctx, deps.MapperRecord.NominalRecordMapper),
		Transaction:  NewTransactionRepository(deps.DB, deps.Ctx, deps.MapperRecord.TransactionRecordMapper),
		Voucher:      NewVoucherRepository(deps.DB, deps.Ctx, deps.MapperRecord.VoucherRecordMapper),
	}
}
