package repository

import (
	"topup_game/internal/domain/record"
	"topup_game/internal/domain/requests"
)

type UserRepository interface {
	FindAllUsers(req *requests.FindAllUsers) ([]*record.UserRecord, *int, error)
	FindByActive(req *requests.FindAllUsers) ([]*record.UserRecord, *int, error)
	FindByTrashed(req *requests.FindAllUsers) ([]*record.UserRecord, *int, error)
	FindById(user_id int) (*record.UserRecord, error)
	FindByEmail(email string) (*record.UserRecord, error)
	CreateUser(request *requests.CreateUserRequest) (*record.UserRecord, error)
	UpdateUser(request *requests.UpdateUserRequest) (*record.UserRecord, error)
	TrashedUser(user_id int) (*record.UserRecord, error)
	RestoreUser(user_id int) (*record.UserRecord, error)
	DeleteUserPermanent(user_id int) (bool, error)
	RestoreAllUser() (bool, error)
	DeleteAllUserPermanent() (bool, error)
}

type RoleRepository interface {
	FindAllRoles(req *requests.FindAllRoles) ([]*record.RoleRecord, *int, error)
	FindByActiveRole(req *requests.FindAllRoles) ([]*record.RoleRecord, *int, error)
	FindByTrashedRole(req *requests.FindAllRoles) ([]*record.RoleRecord, *int, error)
	FindById(role_id int) (*record.RoleRecord, error)
	FindByName(name string) (*record.RoleRecord, error)
	FindByUserId(user_id int) ([]*record.RoleRecord, error)
	CreateRole(request *requests.CreateRoleRequest) (*record.RoleRecord, error)
	UpdateRole(request *requests.UpdateRoleRequest) (*record.RoleRecord, error)
	TrashedRole(role_id int) (*record.RoleRecord, error)

	RestoreRole(role_id int) (*record.RoleRecord, error)
	DeleteRolePermanent(role_id int) (bool, error)
	RestoreAllRole() (bool, error)
	DeleteAllRolePermanent() (bool, error)
}

type RefreshTokenRepository interface {
	FindByToken(token string) (*record.RefreshTokenRecord, error)
	FindByUserId(user_id int) (*record.RefreshTokenRecord, error)
	CreateRefreshToken(req *requests.CreateRefreshToken) (*record.RefreshTokenRecord, error)
	UpdateRefreshToken(req *requests.UpdateRefreshToken) (*record.RefreshTokenRecord, error)
	DeleteRefreshToken(token string) error
	DeleteRefreshTokenByUserId(user_id int) error
}

type UserRoleRepository interface {
	AssignRoleToUser(req *requests.CreateUserRoleRequest) (*record.UserRoleRecord, error)
	RemoveRoleFromUser(req *requests.RemoveUserRoleRequest) error
}

type BankRepository interface {
	FindAllBanks(req *requests.FindAllBanks) ([]*record.BankRecord, *int, error)
	FindByActiveBanks(req *requests.FindAllBanks) ([]*record.BankRecord, *int, error)
	FindByTrashedBanks(req *requests.FindAllBanks) ([]*record.BankRecord, *int, error)

	FindById(id int) (*record.BankRecord, error)
	CreateBank(req *requests.CreateBankRequest) (*record.BankRecord, error)
	UpdateBank(req *requests.UpdateBankRequest) (*record.BankRecord, error)
	TrashedBank(id int) (*record.BankRecord, error)
	RestoreBank(id int) (*record.BankRecord, error)
	DeleteBankPermanent(bank_id int) (bool, error)
	RestoreAllBanks() (bool, error)
	DeleteAllBanksPermanent() (bool, error)

	FindMonthAmountBankSuccess(req *requests.MonthAmountBankRequest) ([]*record.MonthAmountBankSuccessRecord, error)
	FindYearAmountBankSuccess(year int) ([]*record.YearAmountBankSuccessRecord, error)
	FindMonthAmountBankFailed(req *requests.MonthAmountBankRequest) ([]*record.MonthAmountBankFailedRecord, error)
	FindYearAmountBankFailed(year int) ([]*record.YearAmountBankFailedRecord, error)
	FindMonthMethodBankSuccess(year int) ([]*record.MonthMethodBankRecord, error)
	FindYearMethodBankSuccess(year int) ([]*record.YearMethodBankRecord, error)
	FindMonthMethodBankFailed(year int) ([]*record.MonthMethodBankRecord, error)
	FindYearMethodBankFailed(year int) ([]*record.YearMethodBankRecord, error)

	FindMonthAmountBankSuccessById(req *requests.MonthAmountBankByIdRequest) ([]*record.MonthAmountBankSuccessRecord, error)
	FindYearAmountBankSuccessById(req *requests.YearAmountBankByIdRequest) ([]*record.YearAmountBankSuccessRecord, error)
	FindMonthAmountBankFailedById(req *requests.MonthAmountBankByIdRequest) ([]*record.MonthAmountBankFailedRecord, error)
	FindYearAmountBankFailedById(req *requests.YearAmountBankByIdRequest) ([]*record.YearAmountBankFailedRecord, error)
	FindMonthMethodBankSuccessById(req *requests.MonthMethodBankByIdRequest) ([]*record.MonthMethodBankRecord, error)
	FindYearMethodBankSuccessById(req *requests.YearMethodBankByIdRequest) ([]*record.YearMethodBankRecord, error)
	FindMonthMethodBankFailedById(req *requests.MonthMethodBankByIdRequest) ([]*record.MonthMethodBankRecord, error)
	FindYearMethodBankFailedById(req *requests.YearMethodBankByIdRequest) ([]*record.YearMethodBankRecord, error)

	FindMonthAmountBankSuccessByMerchant(req *requests.MonthAmountBankByMerchantRequest) ([]*record.MonthAmountBankSuccessRecord, error)
	FindYearAmountBankSuccessByMerchant(req *requests.YearAmountBankByMerchantRequest) ([]*record.YearAmountBankSuccessRecord, error)
	FindMonthAmountBankFailedByMerchant(req *requests.MonthAmountBankByMerchantRequest) ([]*record.MonthAmountBankFailedRecord, error)
	FindYearAmountBankFailedByMerchant(req *requests.YearAmountBankByMerchantRequest) ([]*record.YearAmountBankFailedRecord, error)
	FindMonthMethodBankSuccessByMerchant(req *requests.MonthMethodBankByMerchantRequest) ([]*record.MonthMethodBankRecord, error)
	FindYearMethodBankSuccessByMerchant(req *requests.YearMethodBankByMerchantRequest) ([]*record.YearMethodBankRecord, error)
	FindMonthMethodBankFailedByMerchant(req *requests.MonthMethodBankByMerchantRequest) ([]*record.MonthMethodBankRecord, error)
	FindYearMethodBankFailedByMerchant(req *requests.YearMethodBankByMerchantRequest) ([]*record.YearMethodBankRecord, error)
}

type CategoryRepository interface {
	FindAllCategories(req *requests.FindAllCategory) ([]*record.CategoryRecord, *int, error)
	FindById(id int) (*record.CategoryRecord, error)
	FindByActiveCategories(req *requests.FindAllCategory) ([]*record.CategoryRecord, *int, error)
	FindByTrashedCategories(req *requests.FindAllCategory) ([]*record.CategoryRecord, *int, error)
	CreateCategory(req *requests.CreateCategoryRequest) (*record.CategoryRecord, error)
	UpdateCategory(req *requests.UpdateCategoryRequest) (*record.CategoryRecord, error)
	TrashedCategory(id int) (*record.CategoryRecord, error)
	RestoreCategory(id int) (*record.CategoryRecord, error)
	DeleteCategoryPermanent(category_id int) (bool, error)
	RestoreAllCategories() (bool, error)
	DeleteAllCategoriesPermanent() (bool, error)

	FindMonthAmountCategorySuccess(req *requests.MonthAmountCategoryRequest) ([]*record.MonthAmountCategorySuccessRecord, error)
	FindYearAmountCategorySuccess(year int) ([]*record.YearAmountCategorySuccessRecord, error)
	FindMonthAmountCategoryFailed(req *requests.MonthAmountCategoryRequest) ([]*record.MonthAmountCategoryFailedRecord, error)
	FindYearAmountCategoryFailed(year int) ([]*record.YearAmountCategoryFailedRecord, error)
	FindMonthMethodCategorySuccess(year int) ([]*record.MonthMethodCategoryRecord, error)
	FindYearMethodCategorySuccess(year int) ([]*record.YearMethodCategoryRecord, error)
	FindMonthMethodCategoryFailed(year int) ([]*record.MonthMethodCategoryRecord, error)
	FindYearMethodCategoryFailed(year int) ([]*record.YearMethodCategoryRecord, error)

	FindMonthAmountCategorySuccessById(req *requests.MonthAmountCategoryByIdRequest) ([]*record.MonthAmountCategorySuccessRecord, error)
	FindYearAmountCategorySuccessById(req *requests.YearAmountCategoryByIdRequest) ([]*record.YearAmountCategorySuccessRecord, error)
	FindMonthAmountCategoryFailedById(req *requests.MonthAmountCategoryByIdRequest) ([]*record.MonthAmountCategoryFailedRecord, error)
	FindYearAmountCategoryFailedById(req *requests.YearAmountCategoryByIdRequest) ([]*record.YearAmountCategoryFailedRecord, error)
	FindMonthMethodCategorySuccessById(req *requests.MonthMethodCategoryByIdRequest) ([]*record.MonthMethodCategoryRecord, error)
	FindYearMethodCategorySuccessById(req *requests.YearMethodCategoryByIdRequest) ([]*record.YearMethodCategoryRecord, error)
	FindMonthMethodCategoryFailedById(req *requests.MonthMethodCategoryByIdRequest) ([]*record.MonthMethodCategoryRecord, error)
	FindYearMethodCategoryFailedById(req *requests.YearMethodCategoryByIdRequest) ([]*record.YearMethodCategoryRecord, error)

	FindMonthAmountCategorySuccessByMerchant(req *requests.MonthAmountCategoryByMerchantRequest) ([]*record.MonthAmountCategorySuccessRecord, error)
	FindYearAmountCategorySuccessByMerchant(req *requests.YearAmountCategoryByMerchantRequest) ([]*record.YearAmountCategorySuccessRecord, error)
	FindMonthAmountCategoryFailedByMerchant(req *requests.MonthAmountCategoryByMerchantRequest) ([]*record.MonthAmountCategoryFailedRecord, error)
	FindYearAmountCategoryFailedByMerchant(req *requests.YearAmountCategoryByMerchantRequest) ([]*record.YearAmountCategoryFailedRecord, error)
	FindMonthMethodCategorySuccessByMerchant(req *requests.MonthMethodCategoryByMerchantRequest) ([]*record.MonthMethodCategoryRecord, error)
	FindYearMethodCategorySuccessByMerchant(req *requests.YearMethodCategoryByMerchantRequest) ([]*record.YearMethodCategoryRecord, error)
	FindMonthMethodCategoryFailedByMerchant(req *requests.MonthMethodCategoryByMerchantRequest) ([]*record.MonthMethodCategoryRecord, error)
	FindYearMethodCategoryFailedByMerchant(req *requests.YearMethodCategoryByMerchantRequest) ([]*record.YearMethodCategoryRecord, error)
}

type MerchantRepository interface {
	FindAllMerchants(req *requests.FindAllMerchants) ([]*record.MerchantRecord, *int, error)
	FindByActive(req *requests.FindAllMerchants) ([]*record.MerchantRecord, *int, error)
	FindByTrashed(req *requests.FindAllMerchants) ([]*record.MerchantRecord, *int, error)
	FindById(user_id int) (*record.MerchantRecord, error)
	CreateMerchant(request *requests.CreateMerchantRequest) (*record.MerchantRecord, error)
	UpdateMerchant(request *requests.UpdateMerchantRequest) (*record.MerchantRecord, error)
	TrashedMerchant(merchant_id int) (*record.MerchantRecord, error)
	RestoreMerchant(merchant_id int) (*record.MerchantRecord, error)
	DeleteMerchantPermanent(Merchant_id int) (bool, error)
	RestoreAllMerchant() (bool, error)
	DeleteAllMerchantPermanent() (bool, error)
}

type NominalRepository interface {
	FindAllNominals(req *requests.FindAllNominals) ([]*record.NominalRecord, *int, error)
	FindByActiveNominal(req *requests.FindAllNominals) ([]*record.NominalRecord, *int, error)
	FindByTrashedNominal(req *requests.FindAllNominals) ([]*record.NominalRecord, *int, error)
	FindById(id int) (*record.NominalRecord, error)
	CreateNominal(req *requests.CreateNominalRequest) (*record.NominalRecord, error)
	UpdateNominal(req *requests.UpdateNominalRequest) (*record.NominalRecord, error)
	UpdateQuantity(nominal int, quantity int) (bool, error)
	TrashedNominal(id int) (*record.NominalRecord, error)
	RestoreNominal(id int) (*record.NominalRecord, error)
	DeleteNominalPermanent(nominal_id int) (bool, error)
	RestoreAllNominal() (bool, error)
	DeleteAllNominalsPermanent() (bool, error)

	FindMonthAmountNominalSuccess(req *requests.MonthAmountNominalRequest) ([]*record.MonthAmountNominalSuccessRecord, error)
	FindYearAmountNominalSuccess(year int) ([]*record.YearAmountNominalSuccessRecord, error)
	FindMonthAmountNominalFailed(req *requests.MonthAmountNominalRequest) ([]*record.MonthAmountNominalFailedRecord, error)
	FindYearAmountNominalFailed(year int) ([]*record.YearAmountNominalFailedRecord, error)
	FindMonthMethodNominalSuccess(year int) ([]*record.MonthMethodNominalRecord, error)
	FindYearMethodNominalSuccess(year int) ([]*record.YearMethodNominalRecord, error)
	FindMonthMethodNominalFailed(year int) ([]*record.MonthMethodNominalRecord, error)
	FindYearMethodNominalFailed(year int) ([]*record.YearMethodNominalRecord, error)

	FindMonthAmountNominalSuccessById(req *requests.MonthAmountNominalByIdRequest) ([]*record.MonthAmountNominalSuccessRecord, error)
	FindYearAmountNominalSuccessById(req *requests.YearAmountNominalByIdRequest) ([]*record.YearAmountNominalSuccessRecord, error)
	FindMonthAmountNominalFailedById(req *requests.MonthAmountNominalByIdRequest) ([]*record.MonthAmountNominalFailedRecord, error)
	FindYearAmountNominalFailedById(req *requests.YearAmountNominalByIdRequest) ([]*record.YearAmountNominalFailedRecord, error)
	FindMonthMethodNominalSuccessById(req *requests.MonthMethodNominalByIdRequest) ([]*record.MonthMethodNominalRecord, error)
	FindYearMethodNominalSuccessById(req *requests.YearMethodNominalByIdRequest) ([]*record.YearMethodNominalRecord, error)
	FindMonthMethodNominalFailedById(req *requests.MonthMethodNominalByIdRequest) ([]*record.MonthMethodNominalRecord, error)
	FindYearMethodNominalFailedById(req *requests.YearMethodNominalByIdRequest) ([]*record.YearMethodNominalRecord, error)

	FindMonthAmountNominalSuccessByMerchant(req *requests.MonthAmountNominalByMerchantRequest) ([]*record.MonthAmountNominalSuccessRecord, error)
	FindYearAmountNominalSuccessByMerchant(req *requests.YearAmountNominalByMerchantRequest) ([]*record.YearAmountNominalSuccessRecord, error)
	FindMonthAmountNominalFailedByMerchant(req *requests.MonthAmountNominalByMerchantRequest) ([]*record.MonthAmountNominalFailedRecord, error)
	FindYearAmountNominalFailedByMerchant(req *requests.YearAmountNominalByMerchantRequest) ([]*record.YearAmountNominalFailedRecord, error)
	FindMonthMethodNominalSuccessByMerchant(req *requests.MonthMethodNominalByMerchantRequest) ([]*record.MonthMethodNominalRecord, error)
	FindYearMethodNominalSuccessByMerchant(req *requests.YearMethodNominalByMerchantRequest) ([]*record.YearMethodNominalRecord, error)
	FindMonthMethodNominalFailedByMerchant(req *requests.MonthMethodNominalByMerchantRequest) ([]*record.MonthMethodNominalRecord, error)
	FindYearMethodNominalFailedByMerchant(req *requests.YearMethodNominalByMerchantRequest) ([]*record.YearMethodNominalRecord, error)
}

type TransactionRepository interface {
	FindAllTransactions(req *requests.FindAllTransactions) ([]*record.TransactionRecord, *int, error)
	FindByActive(req *requests.FindAllTransactions) ([]*record.TransactionRecord, *int, error)
	FindByTrashed(req *requests.FindAllTransactions) ([]*record.TransactionRecord, *int, error)
	FindById(user_id int) (*record.TransactionRecord, error)
	CreateTransaction(req *requests.CreateTransactionRequest, status string) (*record.TransactionRecord, error)
	UpdateTransaction(req *requests.UpdateTransactionRequest) (*record.TransactionRecord, error)
	UpdateTransactionStatus(transaction_id int, status string) (bool, error)
	TrashTransaction(id int) (*record.TransactionRecord, error)
	RestoreTransaction(id int) (*record.TransactionRecord, error)
	DeleteTransactionPermanent(transaction_id int) (bool, error)
	RestoreAllTransactions() (bool, error)
	DeleteAllTransactionsPermanent() (bool, error)

	FindMonthAmountTransactionSuccess(req *requests.MonthAmountTransactionRequest) ([]*record.MonthAmountTransactionSuccessRecord, error)
	FindYearAmountTransactionSuccess(year int) ([]*record.YearAmountTransactionSuccessRecord, error)
	FindMonthAmountTransactionFailed(req *requests.MonthAmountTransactionRequest) ([]*record.MonthAmountTransactionFailedRecord, error)
	FindYearAmountTransactionFailed(year int) ([]*record.YearAmountTransactionFailedRecord, error)
	FindMonthMethodTransactionSuccess(year int) ([]*record.MonthMethodTransactionRecord, error)
	FindYearMethodTransactionSuccess(year int) ([]*record.YearMethodTransactionRecord, error)
	FindMonthMethodTransactionFailed(year int) ([]*record.MonthMethodTransactionRecord, error)
	FindYearMethodTransactionFailed(year int) ([]*record.YearMethodTransactionRecord, error)

	FindMonthAmountTransactionSuccessByMerchant(req *requests.MonthAmountTransactionByMerchantRequest) ([]*record.MonthAmountTransactionSuccessRecord, error)
	FindYearAmountTransactionSuccessByMerchant(req *requests.YearAmountTransactionByMerchantRequest) ([]*record.YearAmountTransactionSuccessRecord, error)
	FindMonthAmountTransactionFailedByMerchant(req *requests.MonthAmountTransactionByMerchantRequest) ([]*record.MonthAmountTransactionFailedRecord, error)
	FindYearAmountTransactionFailedByMerchant(req *requests.YearAmountTransactionByMerchantRequest) ([]*record.YearAmountTransactionFailedRecord, error)
	FindMonthMethodTransactionSuccessByMerchant(req *requests.MonthMethodTransactionByMerchantRequest) ([]*record.MonthMethodTransactionRecord, error)
	FindYearMethodTransactionSuccessByMerchant(req *requests.YearMethodTransactionByMerchantRequest) ([]*record.YearMethodTransactionRecord, error)
	FindMonthMethodTransactionFailedByMerchant(req *requests.MonthMethodTransactionByMerchantRequest) ([]*record.MonthMethodTransactionRecord, error)
	FindYearMethodTransactionFailedByMerchant(req *requests.YearMethodTransactionByMerchantRequest) ([]*record.YearMethodTransactionRecord, error)
}

type VoucherRepository interface {
	FindAllVouchers(req *requests.FindAllVouchers) ([]*record.VoucherRecord, *int, error)
	FindByActiveVouchers(req *requests.FindAllVouchers) ([]*record.VoucherRecord, *int, error)
	FindByTrashedVoucher(req *requests.FindAllVouchers) ([]*record.VoucherRecord, *int, error)
	FindById(id int) (*record.VoucherRecord, error)
	CreateVoucher(req *requests.CreateVoucherRequest) (*record.VoucherRecord, error)
	UpdateVoucher(req *requests.UpdateVoucherRequest) (*record.VoucherRecord, error)
	TrashVoucher(id int) (*record.VoucherRecord, error)
	RestoreVoucher(id int) (*record.VoucherRecord, error)
	DeleteVoucherPermanent(category_id int) (bool, error)
	RestoreAllVouchers() (bool, error)
	DeleteAllVouchersPermanent() (bool, error)

	FindMonthAmountVoucherSuccess(req *requests.MonthAmountVoucherRequest) ([]*record.MonthAmountVoucherSuccessRecord, error)
	FindYearAmountVoucherSuccess(year int) ([]*record.YearAmountVoucherSuccessRecord, error)
	FindMonthAmountVoucherFailed(req *requests.MonthAmountVoucherRequest) ([]*record.MonthAmountVoucherFailedRecord, error)
	FindYearAmountVoucherFailed(year int) ([]*record.YearAmountVoucherFailedRecord, error)
	FindMonthMethodVoucherSuccess(year int) ([]*record.MonthMethodVoucherRecord, error)
	FindYearMethodVoucherSuccess(year int) ([]*record.YearMethodVoucherRecord, error)
	FindMonthMethodVoucherFailed(year int) ([]*record.MonthMethodVoucherRecord, error)
	FindYearMethodVoucherFailed(year int) ([]*record.YearMethodVoucherRecord, error)

	FindMonthAmountVoucherSuccessById(req *requests.MonthAmountVoucherByIdRequest) ([]*record.MonthAmountVoucherSuccessRecord, error)
	FindYearAmountVoucherSuccessById(req *requests.YearAmountVoucherByIdRequest) ([]*record.YearAmountVoucherSuccessRecord, error)
	FindMonthAmountVoucherFailedById(req *requests.MonthAmountVoucherByIdRequest) ([]*record.MonthAmountVoucherFailedRecord, error)
	FindYearAmountVoucherFailedById(req *requests.YearAmountVoucherByIdRequest) ([]*record.YearAmountVoucherFailedRecord, error)
	FindMonthMethodVoucherSuccessById(req *requests.MonthMethodVoucherByIdRequest) ([]*record.MonthMethodVoucherRecord, error)
	FindYearMethodVoucherSuccessById(req *requests.YearMethodVoucherByIdRequest) ([]*record.YearMethodVoucherRecord, error)
	FindMonthMethodVoucherFailedById(req *requests.MonthMethodVoucherByIdRequest) ([]*record.MonthMethodVoucherRecord, error)
	FindYearMethodVoucherFailedById(req *requests.YearMethodVoucherByIdRequest) ([]*record.YearMethodVoucherRecord, error)

	FindMonthAmountVoucherSuccessByMerchant(req *requests.MonthAmountVoucherByMerchantRequest) ([]*record.MonthAmountVoucherSuccessRecord, error)
	FindYearAmountVoucherSuccessByMerchant(req *requests.YearAmountVoucherByMerchantRequest) ([]*record.YearAmountVoucherSuccessRecord, error)
	FindMonthAmountVoucherFailedByMerchant(req *requests.MonthAmountVoucherByMerchantRequest) ([]*record.MonthAmountVoucherFailedRecord, error)
	FindYearAmountVoucherFailedByMerchant(req *requests.YearAmountVoucherByMerchantRequest) ([]*record.YearAmountVoucherFailedRecord, error)
	FindMonthMethodVoucherSuccessByMerchant(req *requests.MonthMethodVoucherByMerchantRequest) ([]*record.MonthMethodVoucherRecord, error)
	FindYearMethodVoucherSuccessByMerchant(req *requests.YearMethodVoucherByMerchantRequest) ([]*record.YearMethodVoucherRecord, error)
	FindMonthMethodVoucherFailedByMerchant(req *requests.MonthMethodVoucherByMerchantRequest) ([]*record.MonthMethodVoucherRecord, error)
	FindYearMethodVoucherFailedByMerchant(req *requests.YearMethodVoucherByMerchantRequest) ([]*record.YearMethodVoucherRecord, error)
}
