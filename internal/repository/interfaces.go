package repository

import (
	"topup_game/internal/domain/record"
	"topup_game/internal/domain/requests"
)

type UserRepository interface {
	FindAllUsers(search string, page, pageSize int) ([]*record.UserRecord, int, error)
	FindById(user_id int) (*record.UserRecord, error)
	FindByEmail(email string) (*record.UserRecord, error)
	FindByActive(search string, page, pageSize int) ([]*record.UserRecord, int, error)
	FindByTrashed(search string, page, pageSize int) ([]*record.UserRecord, int, error)
	CreateUser(request *requests.CreateUserRequest) (*record.UserRecord, error)
	UpdateUser(request *requests.UpdateUserRequest) (*record.UserRecord, error)
	TrashedUser(user_id int) (*record.UserRecord, error)
	RestoreUser(user_id int) (*record.UserRecord, error)
	DeleteUserPermanent(user_id int) (bool, error)
	RestoreAllUser() (bool, error)
	DeleteAllUserPermanent() (bool, error)
}

type RoleRepository interface {
	FindAllRoles(page int, pageSize int, search string) ([]*record.RoleRecord, int, error)
	FindById(role_id int) (*record.RoleRecord, error)
	FindByName(name string) (*record.RoleRecord, error)
	FindByUserId(user_id int) ([]*record.RoleRecord, error)
	FindByActiveRole(page int, pageSize int, search string) ([]*record.RoleRecord, int, error)
	FindByTrashedRole(page int, pageSize int, search string) ([]*record.RoleRecord, int, error)
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
	FindAllBanks(page int, pageSize int, search string) ([]*record.BankRecord, int, error)
	FindById(id int) (*record.BankRecord, error)
	FindByActiveBanks(page int, pageSize int, search string) ([]*record.BankRecord, int, error)
	FindByTrashedBanks(page int, pageSize int, search string) ([]*record.BankRecord, int, error)
	CreateBank(req *requests.CreateBankRequest) (*record.BankRecord, error)
	UpdateBank(req *requests.UpdateBankRequest) (*record.BankRecord, error)
	TrashedBank(id int) (*record.BankRecord, error)
	RestoreBank(id int) (*record.BankRecord, error)
	DeleteBankPermanent(bank_id int) (bool, error)
	RestoreAllBanks() (bool, error)
	DeleteAllBanksPermanent() (bool, error)
}

type CategoryRepository interface {
	FindAllCategories(page int, pageSize int, search string) ([]*record.CategoryRecord, int, error)
	FindById(id int) (*record.CategoryRecord, error)
	FindByActiveCategories(page int, pageSize int, search string) ([]*record.CategoryRecord, int, error)
	FindByTrashedCategory(page int, pageSize int, search string) ([]*record.CategoryRecord, int, error)
	CreateCategory(req *requests.CreateCategoryRequest) (*record.CategoryRecord, error)
	UpdateCategory(req *requests.UpdateCategoryRequest) (*record.CategoryRecord, error)
	TrashedCategory(id int) (*record.CategoryRecord, error)
	RestoreCategory(id int) (*record.CategoryRecord, error)
	DeleteCategoryPermanent(category_id int) (bool, error)
	RestoreAllCategories() (bool, error)
	DeleteAllCategoriesPermanent() (bool, error)
}

type MerchantRepository interface {
	FindAllMerchants(search string, page, pageSize int) ([]*record.MerchantRecord, int, error)
	FindByActive(search string, page, pageSize int) ([]*record.MerchantRecord, int, error)
	FindByTrashed(search string, page, pageSize int) ([]*record.MerchantRecord, int, error)
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
	FindAllNominals(page int, pageSize int, search string) ([]*record.NominalRecord, int, error)
	FindById(id int) (*record.NominalRecord, error)
	FindByActiveNominal(page int, pageSize int, search string) ([]*record.NominalRecord, int, error)
	FindByTrashedNominal(page int, pageSize int, search string) ([]*record.NominalRecord, int, error)
	CreateNominal(req *requests.CreateNominalRequest) (*record.NominalRecord, error)
	UpdateNominal(req *requests.UpdateNominalRequest) (*record.NominalRecord, error)
	UpdateQuantity(nominal int, quantity int) (bool, error)
	TrashedNominal(id int) (*record.NominalRecord, error)
	RestoreNominal(id int) (*record.NominalRecord, error)
	DeleteNominalPermanent(nominal_id int) (bool, error)
	RestoreAllNominal() (bool, error)
	DeleteAllNominalsPermanent() (bool, error)
}

type TransactionRepository interface {
	FindAllTransactions(search string, page, pageSize int) ([]*record.TransactionRecord, int, error)
	FindByActive(search string, page, pageSize int) ([]*record.TransactionRecord, int, error)
	FindByTrashed(search string, page, pageSize int) ([]*record.TransactionRecord, int, error)
	FindById(user_id int) (*record.TransactionRecord, error)
	CreateTransaction(req *requests.CreateTransactionRequest, status string) (*record.TransactionRecord, error)
	UpdateTransaction(req *requests.UpdateTransactionRequest) (*record.TransactionRecord, error)
	UpdateTransactionStatus(transaction_id int, status string) (bool, error)
	TrashTransaction(id int) (*record.TransactionRecord, error)
	RestoreTransaction(id int) (*record.TransactionRecord, error)
	DeleteTransactionPermanent(transaction_id int) (bool, error)
	RestoreAllTransactions() (bool, error)
	DeleteAllTransactionsPermanent() (bool, error)
}

type VoucherRepository interface {
	FindAllVouchers(page int, pageSize int, search string) ([]*record.VoucherRecord, int, error)
	FindById(id int) (*record.VoucherRecord, error)
	FindByActiveVouchers(page int, pageSize int, search string) ([]*record.VoucherRecord, int, error)
	FindByTrashedVoucher(page int, pageSize int, search string) ([]*record.VoucherRecord, int, error)
	CreateVoucher(req *requests.CreateVoucherRequest) (*record.VoucherRecord, error)
	UpdateVoucher(req *requests.UpdateVoucherRequest) (*record.VoucherRecord, error)
	TrashVoucher(id int) (*record.VoucherRecord, error)
	RestoreVoucher(id int) (*record.VoucherRecord, error)
	DeleteVoucherPermanent(category_id int) (bool, error)
	RestoreAllVouchers() (bool, error)
	DeleteAllVouchersPermanent() (bool, error)
}
