package recordmapper

import (
	"topup_game/internal/domain/record"
	db "topup_game/pkg/database/schema"
)

type UserRecordMapping interface {
	ToUserRecord(user *db.User) *record.UserRecord
	ToUserRecordPagination(user *db.GetUsersRow) *record.UserRecord
	ToUsersRecordPagination(users []*db.GetUsersRow) []*record.UserRecord

	ToUserRecordActivePagination(user *db.GetUsersActiveRow) *record.UserRecord
	ToUsersRecordActivePagination(users []*db.GetUsersActiveRow) []*record.UserRecord
	ToUserRecordTrashedPagination(user *db.GetUserTrashedRow) *record.UserRecord
	ToUsersRecordTrashedPagination(users []*db.GetUserTrashedRow) []*record.UserRecord
}

type RoleRecordMapping interface {
	ToRoleRecord(role *db.Role) *record.RoleRecord
	ToRolesRecord(roles []*db.Role) []*record.RoleRecord

	ToRoleRecordAll(role *db.GetRolesRow) *record.RoleRecord
	ToRolesRecordAll(roles []*db.GetRolesRow) []*record.RoleRecord

	ToRoleRecordActive(role *db.GetActiveRolesRow) *record.RoleRecord
	ToRolesRecordActive(roles []*db.GetActiveRolesRow) []*record.RoleRecord
	ToRoleRecordTrashed(role *db.GetTrashedRolesRow) *record.RoleRecord
	ToRolesRecordTrashed(roles []*db.GetTrashedRolesRow) []*record.RoleRecord
}

type UserRoleRecordMapping interface {
	ToUserRoleRecord(userRole *db.UserRole) *record.UserRoleRecord
}

type RefreshTokenRecordMapping interface {
	ToRefreshTokenRecord(refreshToken *db.RefreshToken) *record.RefreshTokenRecord
	ToRefreshTokensRecord(refreshTokens []*db.RefreshToken) []*record.RefreshTokenRecord
}

type BankRecordMapping interface {
	ToBankRecord(Bank *db.Bank) *record.BankRecord
	ToBanksRecord(Banks []*db.Bank) []*record.BankRecord
	ToBankRecordAll(Bank *db.GetBanksRow) *record.BankRecord
	ToBanksRecordAll(Banks []*db.GetBanksRow) []*record.BankRecord
	ToBankRecordActive(Bank *db.GetBanksActiveRow) *record.BankRecord
	ToBanksRecordActive(Banks []*db.GetBanksActiveRow) []*record.BankRecord
	ToBankRecordTrashed(Bank *db.GetBanksTrashedRow) *record.BankRecord
	ToBanksRecordTrashed(Banks []*db.GetBanksTrashedRow) []*record.BankRecord
}

type CategoryRecordMapping interface {
	ToCategoryRecord(Category *db.Category) *record.CategoryRecord
	ToCategorysRecord(Categorys []*db.Category) []*record.CategoryRecord
	ToCategoryRecordAll(Category *db.GetCategoriesRow) *record.CategoryRecord
	ToCategorysRecordAll(Categorys []*db.GetCategoriesRow) []*record.CategoryRecord
	ToCategoryRecordActive(Category *db.GetCategoriesActiveRow) *record.CategoryRecord
	ToCategorysRecordActive(Categorys []*db.GetCategoriesActiveRow) []*record.CategoryRecord
	ToCategoryRecordTrashed(Category *db.GetCategoriesTrashedRow) *record.CategoryRecord
	ToCategorysRecordTrashed(Categorys []*db.GetCategoriesTrashedRow) []*record.CategoryRecord
}

type MerchantRecordMapping interface {
	ToMerchantRecord(Merchant *db.Merchant) *record.MerchantRecord
	ToMerchantRecordPagination(Merchant *db.GetMerchantsRow) *record.MerchantRecord
	ToMerchantsRecordPagination(Merchants []*db.GetMerchantsRow) []*record.MerchantRecord
	ToMerchantRecordActivePagination(Merchant *db.GetMerchantsActiveRow) *record.MerchantRecord
	ToMerchantsRecordActivePagination(Merchants []*db.GetMerchantsActiveRow) []*record.MerchantRecord
	ToMerchantRecordTrashedPagination(Merchant *db.GetMerchantsTrashedRow) *record.MerchantRecord
	ToMerchantsRecordTrashedPagination(Merchants []*db.GetMerchantsTrashedRow) []*record.MerchantRecord
}

type NominalRecordMapping interface {
	ToNominalRecord(nominal *db.Nominal) *record.NominalRecord
	ToNominalRecords(nominals []*db.Nominal) []*record.NominalRecord
	ToNominalRecordAll(nominal *db.GetNominalsRow) *record.NominalRecord
	ToNominalRecordsAll(nominals []*db.GetNominalsRow) []*record.NominalRecord
	ToNominalRecordActive(nominal *db.GetNominalsActiveRow) *record.NominalRecord
	ToNominalRecordsActive(nominals []*db.GetNominalsActiveRow) []*record.NominalRecord
	ToNominalRecordTrashed(nominal *db.GetNominalsTrashedRow) *record.NominalRecord
	ToNominalRecordsTrashed(nominals []*db.GetNominalsTrashedRow) []*record.NominalRecord
}

type TransactionRecordMapping interface {
	ToTransactionRecord(transaction *db.Transaction) *record.TransactionRecord
	ToTransactionRecords(transactions []*db.Transaction) []*record.TransactionRecord
	ToTransactionRecordAll(transaction *db.GetTransactionsRow) *record.TransactionRecord
	ToTransactionRecordsAll(transactions []*db.GetTransactionsRow) []*record.TransactionRecord
	ToTransactionRecordActive(transaction *db.GetTransactionsActiveRow) *record.TransactionRecord
	ToTransactionRecordsActive(transactions []*db.GetTransactionsActiveRow) []*record.TransactionRecord
	ToTransactionRecordTrashed(transaction *db.GetTransactionsTrashedRow) *record.TransactionRecord
	ToTransactionRecordsTrashed(transactions []*db.GetTransactionsTrashedRow) []*record.TransactionRecord
}

type VoucherRecordMapping interface {
	ToVoucherRecord(voucher *db.Voucher) *record.VoucherRecord
	ToVouchersRecord(vouchers []*db.Voucher) []*record.VoucherRecord
	ToVoucherRecordAll(voucher *db.GetVouchersRow) *record.VoucherRecord
	ToVouchersRecordAll(vouchers []*db.GetVouchersRow) []*record.VoucherRecord
	ToVoucherRecordActive(voucher *db.GetVouchersActiveRow) *record.VoucherRecord
	ToVouchersRecordActive(vouchers []*db.GetVouchersActiveRow) []*record.VoucherRecord
	ToVoucherRecordTrashed(voucher *db.GetVouchersTrashedRow) *record.VoucherRecord
	ToVouchersRecordTrashed(vouchers []*db.GetVouchersTrashedRow) []*record.VoucherRecord
}
