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

	ToBankRecordMonthAmountSuccess(b *db.GetMonthAmountBankSuccessRow) *record.MonthAmountBankSuccessRecord
	ToBanksRecordMonthAmountSuccess(b []*db.GetMonthAmountBankSuccessRow) []*record.MonthAmountBankSuccessRecord
	ToBankRecordYearAmountSuccess(b *db.GetYearAmountBankSuccessRow) *record.YearAmountBankSuccessRecord
	ToBanksRecordYearAmountSuccess(b []*db.GetYearAmountBankSuccessRow) []*record.YearAmountBankSuccessRecord
	ToBankRecordMonthAmountFailed(b *db.GetMonthAmountBankFailedRow) *record.MonthAmountBankFailedRecord
	ToBanksRecordMonthAmountFailed(b []*db.GetMonthAmountBankFailedRow) []*record.MonthAmountBankFailedRecord

	ToBankRecordYearAmountFailed(b *db.GetYearAmountBankFailedRow) *record.YearAmountBankFailedRecord
	ToBanksRecordYearAmountFailed(b []*db.GetYearAmountBankFailedRow) []*record.YearAmountBankFailedRecord
	ToBankRecordMonthMethodSuccess(b *db.GetMonthBankMethodsSuccessRow) *record.MonthMethodBankRecord
	ToBanksRecordMonthMethodSuccess(b []*db.GetMonthBankMethodsSuccessRow) []*record.MonthMethodBankRecord

	ToBankRecordMonthMethodFailed(b *db.GetMonthBankMethodsFailedRow) *record.MonthMethodBankRecord
	ToBanksRecordMonthMethodFailed(b []*db.GetMonthBankMethodsFailedRow) []*record.MonthMethodBankRecord
	ToBankRecordYearMethodSuccess(b *db.GetYearBankMethodsSuccessRow) *record.YearMethodBankRecord
	ToBanksRecordYearMethodSuccess(b []*db.GetYearBankMethodsSuccessRow) []*record.YearMethodBankRecord
	ToBankRecordYearMethodFailed(b *db.GetYearBankMethodsFailedRow) *record.YearMethodBankRecord
	ToBanksRecordYearMethodFailed(b []*db.GetYearBankMethodsFailedRow) []*record.YearMethodBankRecord

	ToBankRecordMonthAmountSuccessById(b *db.GetMonthAmountBankSuccessByIdRow) *record.MonthAmountBankSuccessRecord
	ToBanksRecordMonthAmountSuccessById(b []*db.GetMonthAmountBankSuccessByIdRow) []*record.MonthAmountBankSuccessRecord
	ToBankRecordYearAmountSuccessById(b *db.GetYearAmountBankSuccessByIdRow) *record.YearAmountBankSuccessRecord
	ToBanksRecordYearAmountSuccessById(b []*db.GetYearAmountBankSuccessByIdRow) []*record.YearAmountBankSuccessRecord

	ToBankRecordMonthAmountFailedById(b *db.GetMonthAmountBankFailedByIdRow) *record.MonthAmountBankFailedRecord
	ToBanksRecordMonthAmountFailedById(b []*db.GetMonthAmountBankFailedByIdRow) []*record.MonthAmountBankFailedRecord
	ToBankRecordYearAmountFailedById(b *db.GetYearAmountBankFailedByIdRow) *record.YearAmountBankFailedRecord
	ToBanksRecordYearAmountFaileById(b []*db.GetYearAmountBankFailedByIdRow) []*record.YearAmountBankFailedRecord

	ToBankRecordMonthMethodSuccessById(b *db.GetMonthBankMethodsSuccessByIdRow) *record.MonthMethodBankRecord
	ToBanksRecordMonthMethodSuccessById(b []*db.GetMonthBankMethodsSuccessByIdRow) []*record.MonthMethodBankRecord
	ToBankRecordMonthMethodFailedById(b *db.GetMonthBankMethodsFailedByIdRow) *record.MonthMethodBankRecord
	ToBanksRecordMonthMethodFailedById(b []*db.GetMonthBankMethodsFailedByIdRow) []*record.MonthMethodBankRecord
	ToBankRecordYearMethodSuccessById(b *db.GetYearBankMethodsSuccessByIdRow) *record.YearMethodBankRecord
	ToBanksRecordYearMethodSuccessById(b []*db.GetYearBankMethodsSuccessByIdRow) []*record.YearMethodBankRecord
	ToBankRecordYearMethodFailedById(b *db.GetYearBankMethodsFailedByIdRow) *record.YearMethodBankRecord
	ToBanksRecordYearMethodFailedById(b []*db.GetYearBankMethodsFailedByIdRow) []*record.YearMethodBankRecord

	ToBankRecordMonthAmountSuccessByMerchant(b *db.GetMonthAmountBankSuccessByMerchantRow) *record.MonthAmountBankSuccessRecord
	ToBanksRecordMonthAmountSuccessByMerchant(b []*db.GetMonthAmountBankSuccessByMerchantRow) []*record.MonthAmountBankSuccessRecord
	ToBankRecordYearAmountSuccessByMerchant(b *db.GetYearAmountBankSuccessByMerchantRow) *record.YearAmountBankSuccessRecord
	ToBanksRecordYearAmountSuccessByMerchant(b []*db.GetYearAmountBankSuccessByMerchantRow) []*record.YearAmountBankSuccessRecord

	ToBankRecordMonthAmountFailedByMerchant(b *db.GetMonthAmountBankFailedByMerchantRow) *record.MonthAmountBankFailedRecord
	ToBanksRecordMonthAmountFailedByMerchant(b []*db.GetMonthAmountBankFailedByMerchantRow) []*record.MonthAmountBankFailedRecord
	ToBankRecordYearAmountFailedByMerchant(b *db.GetYearAmountBankFailedByMerchantRow) *record.YearAmountBankFailedRecord
	ToBanksRecordYearAmountFaileByMerchant(b []*db.GetYearAmountBankFailedByMerchantRow) []*record.YearAmountBankFailedRecord

	ToBankRecordMonthMethodSuccessByMerchant(b *db.GetMonthBankMethodsSuccessByMerchantRow) *record.MonthMethodBankRecord
	ToBanksRecordMonthMethodSuccessByMerchant(b []*db.GetMonthBankMethodsSuccessByMerchantRow) []*record.MonthMethodBankRecord
	ToBankRecordMonthMethodFailedByMerchant(b *db.GetMonthBankMethodsFailedByMerchantRow) *record.MonthMethodBankRecord
	ToBanksRecordMonthMethodFailedByMerchant(b []*db.GetMonthBankMethodsFailedByMerchantRow) []*record.MonthMethodBankRecord
	ToBankRecordYearMethodSuccessByMerchant(b *db.GetYearBankMethodsSuccessByMerchantRow) *record.YearMethodBankRecord
	ToBanksRecordYearMethodSuccessByMerchant(b []*db.GetYearBankMethodsSuccessByMerchantRow) []*record.YearMethodBankRecord
	ToBankRecordYearMethodFailedByMerchant(b *db.GetYearBankMethodsFailedByMerchantRow) *record.YearMethodBankRecord
	ToBanksRecordYearMethodFailedByMerchant(b []*db.GetYearBankMethodsFailedByMerchantRow) []*record.YearMethodBankRecord
}

type CategoryRecordMapping interface {
	ToCategoryRecord(Category *db.Category) *record.CategoryRecord
	ToCategoriesRecord(Categorys []*db.Category) []*record.CategoryRecord
	ToCategoryRecordAll(Category *db.GetCategoriesRow) *record.CategoryRecord
	ToCategoriesRecordAll(Categorys []*db.GetCategoriesRow) []*record.CategoryRecord
	ToCategoryRecordActive(Category *db.GetCategoriesActiveRow) *record.CategoryRecord
	ToCategoriesRecordActive(Categorys []*db.GetCategoriesActiveRow) []*record.CategoryRecord
	ToCategoryRecordTrashed(Category *db.GetCategoriesTrashedRow) *record.CategoryRecord
	ToCategoriesRecordTrashed(Categorys []*db.GetCategoriesTrashedRow) []*record.CategoryRecord

	ToCategoryRecordMonthAmountSuccess(b *db.GetMonthAmountCategorySuccessRow) *record.MonthAmountCategorySuccessRecord
	ToCategoriesRecordMonthAmountSuccess(b []*db.GetMonthAmountCategorySuccessRow) []*record.MonthAmountCategorySuccessRecord
	ToCategoryRecordYearAmountSuccess(b *db.GetYearAmountCategorySuccessRow) *record.YearAmountCategorySuccessRecord
	ToCategoriesRecordYearAmountSuccess(b []*db.GetYearAmountCategorySuccessRow) []*record.YearAmountCategorySuccessRecord
	ToCategoryRecordMonthAmountFailed(b *db.GetMonthAmountCategoryFailedRow) *record.MonthAmountCategoryFailedRecord
	ToCategoriesRecordMonthAmountFailed(b []*db.GetMonthAmountCategoryFailedRow) []*record.MonthAmountCategoryFailedRecord

	ToCategoryRecordYearAmountFailed(b *db.GetYearAmountCategoryFailedRow) *record.YearAmountCategoryFailedRecord
	ToCategoriesRecordYearAmountFailed(b []*db.GetYearAmountCategoryFailedRow) []*record.YearAmountCategoryFailedRecord
	ToCategoryRecordMonthMethodSuccess(b *db.GetMonthMethodCategoriesSuccessRow) *record.MonthMethodCategoryRecord
	ToCategoriesRecordMonthMethodSuccess(b []*db.GetMonthMethodCategoriesSuccessRow) []*record.MonthMethodCategoryRecord

	ToCategoryRecordMonthMethodFailed(b *db.GetMonthMethodCategoriesFailedRow) *record.MonthMethodCategoryRecord
	ToCategoriesRecordMonthMethodFailed(b []*db.GetMonthMethodCategoriesFailedRow) []*record.MonthMethodCategoryRecord
	ToCategoryRecordYearMethodSuccess(b *db.GetYearMethodCategoriesSuccessRow) *record.YearMethodCategoryRecord
	ToCategoriesRecordYearMethodSuccess(b []*db.GetYearMethodCategoriesSuccessRow) []*record.YearMethodCategoryRecord
	ToCategoryRecordYearMethodFailed(b *db.GetYearMethodCategoriesFailedRow) *record.YearMethodCategoryRecord
	ToCategoriesRecordYearMethodFailed(b []*db.GetYearMethodCategoriesFailedRow) []*record.YearMethodCategoryRecord

	ToCategoryRecordMonthAmountSuccessById(b *db.GetMonthAmountCategorySuccessByIdRow) *record.MonthAmountCategorySuccessRecord
	ToCategoriesRecordMonthAmountSuccessById(b []*db.GetMonthAmountCategorySuccessByIdRow) []*record.MonthAmountCategorySuccessRecord
	ToCategoryRecordYearAmountSuccessById(b *db.GetYearAmountCategorySuccessByIdRow) *record.YearAmountCategorySuccessRecord
	ToCategoriesRecordYearAmountSuccessById(b []*db.GetYearAmountCategorySuccessByIdRow) []*record.YearAmountCategorySuccessRecord

	ToCategoryRecordMonthAmountFailedById(b *db.GetMonthAmountCategoryFailedByIdRow) *record.MonthAmountCategoryFailedRecord
	ToCategoriesRecordMonthAmountFailedById(b []*db.GetMonthAmountCategoryFailedByIdRow) []*record.MonthAmountCategoryFailedRecord
	ToCategoryRecordYearAmountFailedById(b *db.GetYearAmountCategoryFailedByIdRow) *record.YearAmountCategoryFailedRecord
	ToCategoriesRecordYearAmountFaileById(b []*db.GetYearAmountCategoryFailedByIdRow) []*record.YearAmountCategoryFailedRecord

	ToCategoryRecordMonthMethodSuccessById(b *db.GetMonthMethodCategoriesSuccessByIdRow) *record.MonthMethodCategoryRecord
	ToCategoriesRecordMonthMethodSuccessById(b []*db.GetMonthMethodCategoriesSuccessByIdRow) []*record.MonthMethodCategoryRecord
	ToCategoryRecordMonthMethodFailedById(b *db.GetMonthMethodCategoriesFailedByIdRow) *record.MonthMethodCategoryRecord
	ToCategoriesRecordMonthMethodFailedById(b []*db.GetMonthMethodCategoriesFailedByIdRow) []*record.MonthMethodCategoryRecord
	ToCategoryRecordYearMethodSuccessById(b *db.GetYearMethodCategoriesSuccessByIdRow) *record.YearMethodCategoryRecord
	ToCategoriesRecordYearMethodSuccessById(b []*db.GetYearMethodCategoriesSuccessByIdRow) []*record.YearMethodCategoryRecord
	ToCategoryRecordYearMethodFailedById(b *db.GetYearMethodCategoriesFailedByIdRow) *record.YearMethodCategoryRecord
	ToCategoriesRecordYearMethodFailedById(b []*db.GetYearMethodCategoriesFailedByIdRow) []*record.YearMethodCategoryRecord

	ToCategoryRecordMonthAmountSuccessByMerchant(b *db.GetMonthAmountCategorySuccessByMerchantRow) *record.MonthAmountCategorySuccessRecord
	ToCategoriesRecordMonthAmountSuccessByMerchant(b []*db.GetMonthAmountCategorySuccessByMerchantRow) []*record.MonthAmountCategorySuccessRecord
	ToCategoryRecordYearAmountSuccessByMerchant(b *db.GetYearAmountCategorySuccessByMerchantRow) *record.YearAmountCategorySuccessRecord
	ToCategoriesRecordYearAmountSuccessByMerchant(b []*db.GetYearAmountCategorySuccessByMerchantRow) []*record.YearAmountCategorySuccessRecord

	ToCategoryRecordMonthAmountFailedByMerchant(b *db.GetMonthAmountCategoryFailedByMerchantRow) *record.MonthAmountCategoryFailedRecord
	ToCategoriesRecordMonthAmountFailedByMerchant(b []*db.GetMonthAmountCategoryFailedByMerchantRow) []*record.MonthAmountCategoryFailedRecord
	ToCategoryRecordYearAmountFailedByMerchant(b *db.GetYearAmountCategoryFailedByMerchantRow) *record.YearAmountCategoryFailedRecord
	ToCategoriesRecordYearAmountFaileByMerchant(b []*db.GetYearAmountCategoryFailedByMerchantRow) []*record.YearAmountCategoryFailedRecord

	ToCategoryRecordMonthMethodSuccessByMerchant(b *db.GetMonthMethodCategoriesSuccessByMerchantRow) *record.MonthMethodCategoryRecord
	ToCategoriesRecordMonthMethodSuccessByMerchant(b []*db.GetMonthMethodCategoriesSuccessByMerchantRow) []*record.MonthMethodCategoryRecord
	ToCategoryRecordMonthMethodFailedByMerchant(b *db.GetMonthMethodCategoriesFailedByMerchantRow) *record.MonthMethodCategoryRecord
	ToCategoriesRecordMonthMethodFailedByMerchant(b []*db.GetMonthMethodCategoriesFailedByMerchantRow) []*record.MonthMethodCategoryRecord
	ToCategoryRecordYearMethodSuccessByMerchant(b *db.GetYearMethodCategoriesSuccessByMerchantRow) *record.YearMethodCategoryRecord
	ToCategoriesRecordYearMethodSuccessByMerchant(b []*db.GetYearMethodCategoriesSuccessByMerchantRow) []*record.YearMethodCategoryRecord
	ToCategoryRecordYearMethodFailedByMerchant(b *db.GetYearMethodCategoriesFailedByMerchantRow) *record.YearMethodCategoryRecord
	ToCategoriesRecordYearMethodFailedByMerchant(b []*db.GetYearMethodCategoriesFailedByMerchantRow) []*record.YearMethodCategoryRecord
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

	ToNominalRecordMonthAmountSuccess(b *db.GetMonthAmountNominalsSuccessRow) *record.MonthAmountNominalSuccessRecord
	ToNominalsRecordMonthAmountSuccess(b []*db.GetMonthAmountNominalsSuccessRow) []*record.MonthAmountNominalSuccessRecord
	ToNominalRecordYearAmountSuccess(b *db.GetYearAmountNominalsSuccessRow) *record.YearAmountNominalSuccessRecord
	ToNominalsRecordYearAmountSuccess(b []*db.GetYearAmountNominalsSuccessRow) []*record.YearAmountNominalSuccessRecord
	ToNominalRecordMonthAmountFailed(b *db.GetMonthAmountNominalsFailedRow) *record.MonthAmountNominalFailedRecord
	ToNominalsRecordMonthAmountFailed(b []*db.GetMonthAmountNominalsFailedRow) []*record.MonthAmountNominalFailedRecord
	ToNominalRecordYearAmountFailed(b *db.GetYearAmountNominalsFailedRow) *record.YearAmountNominalFailedRecord
	ToNominalsRecordYearAmountFailed(b []*db.GetYearAmountNominalsFailedRow) []*record.YearAmountNominalFailedRecord

	ToNominalRecordMonthMethodSuccess(b *db.GetMonthMethodNominalsSuccessRow) *record.MonthMethodNominalRecord
	ToNominalsRecordMonthMethodSuccess(b []*db.GetMonthMethodNominalsSuccessRow) []*record.MonthMethodNominalRecord
	ToNominalRecordMonthMethodFailed(b *db.GetMonthMethodNominalsFailedRow) *record.MonthMethodNominalRecord
	ToNominalsRecordMonthMethodFailed(b []*db.GetMonthMethodNominalsFailedRow) []*record.MonthMethodNominalRecord
	ToNominalRecordYearMethodSuccess(b *db.GetYearMethodNominalsSuccessRow) *record.YearMethodNominalRecord
	ToNominalsRecordYearMethodSuccess(b []*db.GetYearMethodNominalsSuccessRow) []*record.YearMethodNominalRecord
	ToNominalRecordYearMethodFailed(b *db.GetYearMethodNominalsFailedRow) *record.YearMethodNominalRecord
	ToNominalsRecordYearMethodFailed(b []*db.GetYearMethodNominalsFailedRow) []*record.YearMethodNominalRecord

	ToNominalRecordMonthAmountSuccessById(b *db.GetMonthAmountNominalsSuccessByIdRow) *record.MonthAmountNominalSuccessRecord
	ToNominalsRecordMonthAmountSuccessById(b []*db.GetMonthAmountNominalsSuccessByIdRow) []*record.MonthAmountNominalSuccessRecord
	ToNominalRecordYearAmountSuccessById(b *db.GetYearAmountNominalsSuccessByIdRow) *record.YearAmountNominalSuccessRecord
	ToNominalsRecordYearAmountSuccessById(b []*db.GetYearAmountNominalsSuccessByIdRow) []*record.YearAmountNominalSuccessRecord

	ToNominalRecordMonthAmountFailedById(b *db.GetMonthAmountNominalsFailedByIdRow) *record.MonthAmountNominalFailedRecord
	ToNominalsRecordMonthAmountFailedById(b []*db.GetMonthAmountNominalsFailedByIdRow) []*record.MonthAmountNominalFailedRecord
	ToNominalRecordYearAmountFailedById(b *db.GetYearAmountNominalsFailedByIdRow) *record.YearAmountNominalFailedRecord
	ToNominalsRecordYearAmountFaileById(b []*db.GetYearAmountNominalsFailedByIdRow) []*record.YearAmountNominalFailedRecord

	ToNominalRecordMonthMethodSuccessById(b *db.GetMonthMethodNominalsSuccessByIdRow) *record.MonthMethodNominalRecord
	ToNominalsRecordMonthMethodSuccessById(b []*db.GetMonthMethodNominalsSuccessByIdRow) []*record.MonthMethodNominalRecord
	ToNominalRecordMonthMethodFailedById(b *db.GetMonthMethodNominalsFailedByIdRow) *record.MonthMethodNominalRecord
	ToNominalsRecordMonthMethodFailedById(b []*db.GetMonthMethodNominalsFailedByIdRow) []*record.MonthMethodNominalRecord
	ToNominalRecordYearMethodSuccessById(b *db.GetYearMethodNominalsSuccessByIdRow) *record.YearMethodNominalRecord
	ToNominalsRecordYearMethodSuccessById(b []*db.GetYearMethodNominalsSuccessByIdRow) []*record.YearMethodNominalRecord
	ToNominalRecordYearMethodFailedById(b *db.GetYearMethodNominalsFailedByIdRow) *record.YearMethodNominalRecord
	ToNominalsRecordYearMethodFailedById(b []*db.GetYearMethodNominalsFailedByIdRow) []*record.YearMethodNominalRecord

	ToNominalRecordMonthAmountSuccessByMerchant(b *db.GetMonthAmountNominalsSuccessByMerchantRow) *record.MonthAmountNominalSuccessRecord
	ToNominalsRecordMonthAmountSuccessByMerchant(b []*db.GetMonthAmountNominalsSuccessByMerchantRow) []*record.MonthAmountNominalSuccessRecord
	ToNominalRecordYearAmountSuccessByMerchant(b *db.GetYearAmountNominalsSuccessByMerchantRow) *record.YearAmountNominalSuccessRecord
	ToNominalsRecordYearAmountSuccessByMerchant(b []*db.GetYearAmountNominalsSuccessByMerchantRow) []*record.YearAmountNominalSuccessRecord

	ToNominalRecordMonthAmountFailedByMerchant(b *db.GetMonthAmountNominalsFailedByMerchantRow) *record.MonthAmountNominalFailedRecord
	ToNominalsRecordMonthAmountFailedByMerchant(b []*db.GetMonthAmountNominalsFailedByMerchantRow) []*record.MonthAmountNominalFailedRecord
	ToNominalRecordYearAmountFailedByMerchant(b *db.GetYearAmountNominalsFailedByMerchantRow) *record.YearAmountNominalFailedRecord
	ToNominalsRecordYearAmountFaileByMerchant(b []*db.GetYearAmountNominalsFailedByMerchantRow) []*record.YearAmountNominalFailedRecord

	ToNominalRecordMonthMethodSuccessByMerchant(b *db.GetMonthMethodNominalsSuccessByMerchantRow) *record.MonthMethodNominalRecord
	ToNominalsRecordMonthMethodSuccessByMerchant(b []*db.GetMonthMethodNominalsSuccessByMerchantRow) []*record.MonthMethodNominalRecord
	ToNominalRecordMonthMethodFailedByMerchant(b *db.GetMonthMethodNominalsFailedByMerchantRow) *record.MonthMethodNominalRecord
	ToNominalsRecordMonthMethodFailedByMerchant(b []*db.GetMonthMethodNominalsFailedByMerchantRow) []*record.MonthMethodNominalRecord
	ToNominalRecordYearMethodSuccessByMerchant(b *db.GetYearMethodNominalsSuccessByMerchantRow) *record.YearMethodNominalRecord
	ToNominalsRecordYearMethodSuccessByMerchant(b []*db.GetYearMethodNominalsSuccessByMerchantRow) []*record.YearMethodNominalRecord
	ToNominalRecordYearMethodFailedByMerchant(b *db.GetYearMethodNominalsFailedByMerchantRow) *record.YearMethodNominalRecord
	ToNominalsRecordYearMethodFailedByMerchant(b []*db.GetYearMethodNominalsFailedByMerchantRow) []*record.YearMethodNominalRecord
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

	ToTransactionRecordMonthAmountSuccess(b *db.GetMonthlyAmountTransactionSuccessRow) *record.MonthAmountTransactionSuccessRecord
	ToTransactionsRecordMonthAmountSuccess(b []*db.GetMonthlyAmountTransactionSuccessRow) []*record.MonthAmountTransactionSuccessRecord
	ToTransactionRecordYearAmountSuccess(b *db.GetYearlyAmountTransactionSuccessRow) *record.YearAmountTransactionSuccessRecord
	ToTransactionsRecordYearAmountSuccess(b []*db.GetYearlyAmountTransactionSuccessRow) []*record.YearAmountTransactionSuccessRecord
	ToTransactionRecordMonthAmountFailed(b *db.GetMonthlyAmountTransactionFailedRow) *record.MonthAmountTransactionFailedRecord
	ToTransactionsRecordMonthAmountFailed(b []*db.GetMonthlyAmountTransactionFailedRow) []*record.MonthAmountTransactionFailedRecord
	ToTransactionRecordYearAmountFailed(b *db.GetYearlyAmountTransactionFailedRow) *record.YearAmountTransactionFailedRecord
	ToTransactionsRecordYearAmountFailed(b []*db.GetYearlyAmountTransactionFailedRow) []*record.YearAmountTransactionFailedRecord

	ToTransactionRecordMonthMethodSuccess(b *db.GetMonthlyTransactionMethodsSuccessRow) *record.MonthMethodTransactionRecord
	ToTransactionsRecordMonthMethodSuccess(b []*db.GetMonthlyTransactionMethodsSuccessRow) []*record.MonthMethodTransactionRecord
	ToTransactionRecordMonthMethodFailed(b *db.GetMonthlyTransactionMethodsFailedRow) *record.MonthMethodTransactionRecord
	ToTransactionsRecordMonthMethodFailed(b []*db.GetMonthlyTransactionMethodsFailedRow) []*record.MonthMethodTransactionRecord
	ToTransactionRecordYearMethodSuccess(b *db.GetYearlyTransactionMethodsSuccessRow) *record.YearMethodTransactionRecord
	ToTransactionsRecordYearMethodSuccess(b []*db.GetYearlyTransactionMethodsSuccessRow) []*record.YearMethodTransactionRecord
	ToTransactionRecordYearMethodFailed(b *db.GetYearlyTransactionMethodsFailedRow) *record.YearMethodTransactionRecord
	ToTransactionsRecordYearMethodFailed(b []*db.GetYearlyTransactionMethodsFailedRow) []*record.YearMethodTransactionRecord

	ToTransactionRecordMonthAmountSuccessByMerchant(b *db.GetMonthlyAmountTransactionSuccessByMerchantRow) *record.MonthAmountTransactionSuccessRecord
	ToTransactionsRecordMonthAmountSuccessByMerchant(b []*db.GetMonthlyAmountTransactionSuccessByMerchantRow) []*record.MonthAmountTransactionSuccessRecord
	ToTransactionRecordYearAmountSuccessByMerchant(b *db.GetYearlyAmountTransactionSuccessByMerchantRow) *record.YearAmountTransactionSuccessRecord
	ToTransactionsRecordYearAmountSuccessByMerchant(b []*db.GetYearlyAmountTransactionSuccessByMerchantRow) []*record.YearAmountTransactionSuccessRecord
	ToTransactionRecordMonthAmountFailedByMerchant(b *db.GetMonthlyAmountTransactionFailedByMerchantRow) *record.MonthAmountTransactionFailedRecord
	ToTransactionsRecordMonthAmountFailedByMerchant(b []*db.GetMonthlyAmountTransactionFailedByMerchantRow) []*record.MonthAmountTransactionFailedRecord
	ToTransactionRecordYearAmountFailedByMerchant(b *db.GetYearlyAmountTransactionFailedByMerchantRow) *record.YearAmountTransactionFailedRecord
	ToTransactionsRecordYearAmountFailedByMerchant(b []*db.GetYearlyAmountTransactionFailedByMerchantRow) []*record.YearAmountTransactionFailedRecord

	ToTransactionRecordMonthMethodSuccessByMerchant(b *db.GetMonthlyTransactionMethodsSuccessByMerchantRow) *record.MonthMethodTransactionRecord
	ToTransactionsRecordMonthMethodSuccessByMerchant(b []*db.GetMonthlyTransactionMethodsSuccessByMerchantRow) []*record.MonthMethodTransactionRecord
	ToTransactionRecordMonthMethodFailedByMerchant(b *db.GetMonthlyTransactionMethodsFailedByMerchantRow) *record.MonthMethodTransactionRecord
	ToTransactionsRecordMonthMethodFailedByMerchant(b []*db.GetMonthlyTransactionMethodsFailedByMerchantRow) []*record.MonthMethodTransactionRecord
	ToTransactionRecordYearMethodSuccessByMerchant(b *db.GetYearlyTransactionMethodsSuccessByMerchantRow) *record.YearMethodTransactionRecord
	ToTransactionsRecordYearMethodSuccessByMerchant(b []*db.GetYearlyTransactionMethodsSuccessByMerchantRow) []*record.YearMethodTransactionRecord
	ToTransactionRecordYearMethodFailedByMerchant(b *db.GetYearlyTransactionMethodsFailedByMerchantRow) *record.YearMethodTransactionRecord
	ToTransactionsRecordYearMethodFailedByMerchant(b []*db.GetYearlyTransactionMethodsFailedByMerchantRow) []*record.YearMethodTransactionRecord
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

	ToVoucherRecordMonthAmountSuccess(b *db.GetMonthAmountVouchersSuccessRow) *record.MonthAmountVoucherSuccessRecord
	ToVouchersRecordMonthAmountSuccess(b []*db.GetMonthAmountVouchersSuccessRow) []*record.MonthAmountVoucherSuccessRecord
	ToVoucherRecordYearAmountSuccess(b *db.GetYearAmountVouchersSuccessRow) *record.YearAmountVoucherSuccessRecord
	ToVouchersRecordYearAmountSuccess(b []*db.GetYearAmountVouchersSuccessRow) []*record.YearAmountVoucherSuccessRecord
	ToVoucherRecordMonthAmountFailed(b *db.GetMonthAmountVouchersFailedRow) *record.MonthAmountVoucherFailedRecord
	ToVouchersRecordMonthAmountFailed(b []*db.GetMonthAmountVouchersFailedRow) []*record.MonthAmountVoucherFailedRecord
	ToVoucherRecordYearAmountFailed(b *db.GetYearAmountVouchersFailedRow) *record.YearAmountVoucherFailedRecord
	ToVouchersRecordYearAmountFailed(b []*db.GetYearAmountVouchersFailedRow) []*record.YearAmountVoucherFailedRecord

	ToVoucherRecordMonthMethodSuccess(b *db.GetMonthMethodVouchersSuccessRow) *record.MonthMethodVoucherRecord
	ToVouchersRecordMonthMethodSuccess(b []*db.GetMonthMethodVouchersSuccessRow) []*record.MonthMethodVoucherRecord

	ToVoucherRecordMonthMethodFailed(b *db.GetMonthMethodVouchersFailedRow) *record.MonthMethodVoucherRecord
	ToVouchersRecordMonthMethodFailed(b []*db.GetMonthMethodVouchersFailedRow) []*record.MonthMethodVoucherRecord
	ToVoucherRecordYearMethodSuccess(b *db.GetYearMethodVouchersSuccessRow) *record.YearMethodVoucherRecord
	ToVouchersRecordYearMethodSuccess(b []*db.GetYearMethodVouchersSuccessRow) []*record.YearMethodVoucherRecord
	ToVoucherRecordYearMethodFailed(b *db.GetYearMethodVouchersFailedRow) *record.YearMethodVoucherRecord
	ToVouchersRecordYearMethodFailed(b []*db.GetYearMethodVouchersFailedRow) []*record.YearMethodVoucherRecord

	ToVoucherRecordMonthAmountSuccessById(b *db.GetMonthAmountVouchersSuccessByIdRow) *record.MonthAmountVoucherSuccessRecord
	ToVouchersRecordMonthAmountSuccessById(b []*db.GetMonthAmountVouchersSuccessByIdRow) []*record.MonthAmountVoucherSuccessRecord
	ToVoucherRecordYearAmountSuccessById(b *db.GetYearAmountVouchersSuccessByIdRow) *record.YearAmountVoucherSuccessRecord
	ToVouchersRecordYearAmountSuccessById(b []*db.GetYearAmountVouchersSuccessByIdRow) []*record.YearAmountVoucherSuccessRecord

	ToVoucherRecordMonthAmountFailedById(b *db.GetMonthAmountVouchersFailedByIdRow) *record.MonthAmountVoucherFailedRecord
	ToVouchersRecordMonthAmountFailedById(b []*db.GetMonthAmountVouchersFailedByIdRow) []*record.MonthAmountVoucherFailedRecord
	ToVoucherRecordYearAmountFailedById(b *db.GetYearAmountVouchersFailedByIdRow) *record.YearAmountVoucherFailedRecord
	ToVouchersRecordYearAmountFailedById(b []*db.GetYearAmountVouchersFailedByIdRow) []*record.YearAmountVoucherFailedRecord

	ToVoucherRecordMonthMethodSuccessById(b *db.GetMonthMethodVouchersSuccessByIdRow) *record.MonthMethodVoucherRecord
	ToVouchersRecordMonthMethodSuccessById(b []*db.GetMonthMethodVouchersSuccessByIdRow) []*record.MonthMethodVoucherRecord
	ToVoucherRecordMonthMethodFailedById(b *db.GetMonthMethodVouchersFailedByIdRow) *record.MonthMethodVoucherRecord
	ToVouchersRecordMonthMethodFailedById(b []*db.GetMonthMethodVouchersFailedByIdRow) []*record.MonthMethodVoucherRecord
	ToVoucherRecordYearMethodSuccessById(b *db.GetYearMethodVouchersSuccessByIdRow) *record.YearMethodVoucherRecord
	ToVouchersRecordYearMethodSuccessById(b []*db.GetYearMethodVouchersSuccessByIdRow) []*record.YearMethodVoucherRecord
	ToVoucherRecordYearMethodFailedById(b *db.GetYearMethodVouchersFailedByIdRow) *record.YearMethodVoucherRecord
	ToVouchersRecordYearMethodFailedById(b []*db.GetYearMethodVouchersFailedByIdRow) []*record.YearMethodVoucherRecord

	ToVoucherRecordMonthAmountSuccessByMerchant(b *db.GetMonthAmountVouchersSuccessByMerchantRow) *record.MonthAmountVoucherSuccessRecord
	ToVouchersRecordMonthAmountSuccessByMerchant(b []*db.GetMonthAmountVouchersSuccessByMerchantRow) []*record.MonthAmountVoucherSuccessRecord
	ToVoucherRecordYearAmountSuccessByMerchant(b *db.GetYearAmountVouchersSuccessByMerchantRow) *record.YearAmountVoucherSuccessRecord
	ToVouchersRecordYearAmountSuccessByMerchant(b []*db.GetYearAmountVouchersSuccessByMerchantRow) []*record.YearAmountVoucherSuccessRecord

	ToVoucherRecordMonthAmountFailedByMerchant(b *db.GetMonthAmountVouchersFailedByMerchantRow) *record.MonthAmountVoucherFailedRecord
	ToVouchersRecordMonthAmountFailedByMerchant(b []*db.GetMonthAmountVouchersFailedByMerchantRow) []*record.MonthAmountVoucherFailedRecord
	ToVoucherRecordYearAmountFailedByMerchant(b *db.GetYearAmountVouchersFailedByMerchantRow) *record.YearAmountVoucherFailedRecord
	ToVouchersRecordYearAmountFailedByMerchant(b []*db.GetYearAmountVouchersFailedByMerchantRow) []*record.YearAmountVoucherFailedRecord

	ToVoucherRecordMonthMethodSuccessByMerchant(b *db.GetMonthMethodVouchersSuccessByMerchantRow) *record.MonthMethodVoucherRecord
	ToVouchersRecordMonthMethodSuccessByMerchant(b []*db.GetMonthMethodVouchersSuccessByMerchantRow) []*record.MonthMethodVoucherRecord
	ToVoucherRecordMonthMethodFailedByMerchant(b *db.GetMonthMethodVouchersFailedByMerchantRow) *record.MonthMethodVoucherRecord
	ToVouchersRecordMonthMethodFailedByMerchant(b []*db.GetMonthMethodVouchersFailedByMerchantRow) []*record.MonthMethodVoucherRecord
	ToVoucherRecordYearMethodSuccessByMerchant(b *db.GetYearMethodVouchersSuccessByMerchantRow) *record.YearMethodVoucherRecord
	ToVouchersRecordYearMethodSuccessByMerchant(b []*db.GetYearMethodVouchersSuccessByMerchantRow) []*record.YearMethodVoucherRecord
	ToVoucherRecordYearMethodFailedByMerchant(b *db.GetYearMethodVouchersFailedByMerchantRow) *record.YearMethodVoucherRecord
	ToVouchersRecordYearMethodFailedByMerchant(b []*db.GetYearMethodVouchersFailedByMerchantRow) []*record.YearMethodVoucherRecord
}
