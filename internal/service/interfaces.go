package service

import (
	"topup_game/internal/domain/requests"
	"topup_game/internal/domain/response"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go
type AuthService interface {
	Register(request *requests.CreateUserRequest) (*response.UserResponse, *response.ErrorResponse)
	Login(request *requests.AuthRequest) (*response.TokenResponse, *response.ErrorResponse)
	RefreshToken(token string) (*response.TokenResponse, *response.ErrorResponse)
	GetMe(token string) (*response.UserResponse, *response.ErrorResponse)
}

type RoleService interface {
	FindAll(req *requests.FindAllRoles) ([]*response.RoleResponse, *int, *response.ErrorResponse)
	FindByActive(req *requests.FindAllRoles) ([]*response.RoleResponseDeleteAt, *int, *response.ErrorResponse)
	FindByTrashed(req *requests.FindAllRoles) ([]*response.RoleResponseDeleteAt, *int, *response.ErrorResponse)
	FindById(role_id int) (*response.RoleResponse, *response.ErrorResponse)
	FindByUserId(id int) ([]*response.RoleResponse, *response.ErrorResponse)
	Create(request *requests.CreateRoleRequest) (*response.RoleResponse, *response.ErrorResponse)
	Update(request *requests.UpdateRoleRequest) (*response.RoleResponse, *response.ErrorResponse)
	Trashed(role_id int) (*response.RoleResponse, *response.ErrorResponse)
	Restore(role_id int) (*response.RoleResponse, *response.ErrorResponse)
	DeletePermanent(role_id int) (bool, *response.ErrorResponse)

	RestoreAll() (bool, *response.ErrorResponse)
	DeleteAllPermanent() (bool, *response.ErrorResponse)
}

type UserService interface {
	FindAll(req *requests.FindAllUsers) ([]*response.UserResponse, *int, *response.ErrorResponse)
	FindByActive(req *requests.FindAllUsers) ([]*response.UserResponseDeleteAt, *int, *response.ErrorResponse)
	FindByTrashed(req *requests.FindAllUsers) ([]*response.UserResponseDeleteAt, *int, *response.ErrorResponse)
	FindByID(id int) (*response.UserResponse, *response.ErrorResponse)
	Create(request *requests.CreateUserRequest) (*response.UserResponse, *response.ErrorResponse)
	Update(request *requests.UpdateUserRequest) (*response.UserResponse, *response.ErrorResponse)
	Trashed(user_id int) (*response.UserResponseDeleteAt, *response.ErrorResponse)
	Restore(user_id int) (*response.UserResponseDeleteAt, *response.ErrorResponse)
	DeletePermanent(user_id int) (bool, *response.ErrorResponse)

	RestoreAll() (bool, *response.ErrorResponse)
	DeleteAllPermanent() (bool, *response.ErrorResponse)
}

type BankService interface {
	FindAll(req *requests.FindAllBanks) ([]*response.BankResponse, *int, *response.ErrorResponse)
	FindByActive(req *requests.FindAllBanks) ([]*response.BankResponseDeleteAt, *int, *response.ErrorResponse)
	FindByTrashed(req *requests.FindAllBanks) ([]*response.BankResponseDeleteAt, *int, *response.ErrorResponse)
	FindByID(id int) (*response.BankResponse, *response.ErrorResponse)
	Create(request *requests.CreateBankRequest) (*response.BankResponse, *response.ErrorResponse)
	Update(request *requests.UpdateBankRequest) (*response.BankResponse, *response.ErrorResponse)
	Trashed(id int) (*response.BankResponseDeleteAt, *response.ErrorResponse)
	Restore(id int) (*response.BankResponseDeleteAt, *response.ErrorResponse)
	DeletePermanent(id int) (bool, *response.ErrorResponse)
	RestoreAll() (bool, *response.ErrorResponse)
	DeleteAllPermanent() (bool, *response.ErrorResponse)

	FindMonthAmountBankSuccess(req *requests.MonthAmountBankRequest) ([]*response.MonthAmountBankSuccessResponse, *response.ErrorResponse)
	FindYearAmountBankSuccess(year int) ([]*response.YearAmountBankSuccessResponse, *response.ErrorResponse)
	FindMonthAmountBankFailed(req *requests.MonthAmountBankRequest) ([]*response.MonthAmountBankFailedResponse, *response.ErrorResponse)
	FindYearAmountBankFailed(year int) ([]*response.YearAmountBankFailedResponse, *response.ErrorResponse)
	FindMonthMethodBankSuccess(year int) ([]*response.MonthMethodBankResponse, *response.ErrorResponse)
	FindYearMethodBankSuccess(year int) ([]*response.YearMethodBankResponse, *response.ErrorResponse)
	FindMonthMethodBankFailed(year int) ([]*response.MonthMethodBankResponse, *response.ErrorResponse)
	FindYearMethodBankFailed(year int) ([]*response.YearMethodBankResponse, *response.ErrorResponse)

	FindMonthAmountBankSuccessById(req *requests.MonthAmountBankByIdRequest) ([]*response.MonthAmountBankSuccessResponse, *response.ErrorResponse)
	FindYearAmountBankSuccessById(req *requests.YearAmountBankByIdRequest) ([]*response.YearAmountBankSuccessResponse, *response.ErrorResponse)
	FindMonthAmountBankFailedById(req *requests.MonthAmountBankByIdRequest) ([]*response.MonthAmountBankFailedResponse, *response.ErrorResponse)
	FindYearAmountBankFailedById(req *requests.YearAmountBankByIdRequest) ([]*response.YearAmountBankFailedResponse, *response.ErrorResponse)
	FindMonthMethodBankSuccessById(req *requests.MonthMethodBankByIdRequest) ([]*response.MonthMethodBankResponse, *response.ErrorResponse)
	FindYearMethodBankSuccessById(req *requests.YearMethodBankByIdRequest) ([]*response.YearMethodBankResponse, *response.ErrorResponse)
	FindMonthMethodBankFailedById(req *requests.MonthMethodBankByIdRequest) ([]*response.MonthMethodBankResponse, *response.ErrorResponse)
	FindYearMethodBankFailedById(req *requests.YearMethodBankByIdRequest) ([]*response.YearMethodBankResponse, *response.ErrorResponse)

	FindMonthAmountBankSuccessByMerchant(req *requests.MonthAmountBankByMerchantRequest) ([]*response.MonthAmountBankSuccessResponse, *response.ErrorResponse)
	FindYearAmountBankSuccessByMerchant(req *requests.YearAmountBankByMerchantRequest) ([]*response.YearAmountBankSuccessResponse, *response.ErrorResponse)
	FindMonthAmountBankFailedByMerchant(req *requests.MonthAmountBankByMerchantRequest) ([]*response.MonthAmountBankFailedResponse, *response.ErrorResponse)
	FindYearAmountBankFailedByMerchant(req *requests.YearAmountBankByMerchantRequest) ([]*response.YearAmountBankFailedResponse, *response.ErrorResponse)
	FindMonthMethodBankSuccessByMerchant(req *requests.MonthMethodBankByMerchantRequest) ([]*response.MonthMethodBankResponse, *response.ErrorResponse)
	FindYearMethodBankSuccessByMerchant(req *requests.YearMethodBankByMerchantRequest) ([]*response.YearMethodBankResponse, *response.ErrorResponse)
	FindMonthMethodBankFailedByMerchant(req *requests.MonthMethodBankByMerchantRequest) ([]*response.MonthMethodBankResponse, *response.ErrorResponse)
	FindYearMethodBankFailedByMerchant(req *requests.YearMethodBankByMerchantRequest) ([]*response.YearMethodBankResponse, *response.ErrorResponse)
}

type CategoryService interface {
	FindAll(req *requests.FindAllCategory) ([]*response.CategoryResponse, *int, *response.ErrorResponse)
	FindByActive(req *requests.FindAllCategory) ([]*response.CategoryResponseDeleteAt, *int, *response.ErrorResponse)
	FindByTrashed(req *requests.FindAllCategory) ([]*response.CategoryResponseDeleteAt, *int, *response.ErrorResponse)
	FindById(id int) (*response.CategoryResponse, *response.ErrorResponse)
	Create(request *requests.CreateCategoryRequest) (*response.CategoryResponse, *response.ErrorResponse)
	Update(request *requests.UpdateCategoryRequest) (*response.CategoryResponse, *response.ErrorResponse)
	Trashed(id int) (*response.CategoryResponseDeleteAt, *response.ErrorResponse)
	Restore(id int) (*response.CategoryResponseDeleteAt, *response.ErrorResponse)
	DeletePermanent(id int) (bool, *response.ErrorResponse)
	RestoreAll() (bool, *response.ErrorResponse)
	DeleteAllPermanent() (bool, *response.ErrorResponse)

	FindMonthAmountCategorySuccess(req *requests.MonthAmountCategoryRequest) ([]*response.MonthAmountCategorySuccessResponse, *response.ErrorResponse)
	FindYearAmountCategorySuccess(year int) ([]*response.YearAmountCategorySuccessResponse, *response.ErrorResponse)
	FindMonthAmountCategoryFailed(req *requests.MonthAmountCategoryRequest) ([]*response.MonthAmountCategoryFailedResponse, *response.ErrorResponse)
	FindYearAmountCategoryFailed(year int) ([]*response.YearAmountCategoryFailedResponse, *response.ErrorResponse)
	FindMonthMethodCategorySuccess(year int) ([]*response.MonthMethodCategoryResponse, *response.ErrorResponse)
	FindYearMethodCategorySuccess(year int) ([]*response.YearMethodCategoryResponse, *response.ErrorResponse)
	FindMonthMethodCategoryFailed(year int) ([]*response.MonthMethodCategoryResponse, *response.ErrorResponse)
	FindYearMethodCategoryFailed(year int) ([]*response.YearMethodCategoryResponse, *response.ErrorResponse)

	FindMonthAmountCategorySuccessById(req *requests.MonthAmountCategoryByIdRequest) ([]*response.MonthAmountCategorySuccessResponse, *response.ErrorResponse)
	FindYearAmountCategorySuccessById(req *requests.YearAmountCategoryByIdRequest) ([]*response.YearAmountCategorySuccessResponse, *response.ErrorResponse)
	FindMonthAmountCategoryFailedById(req *requests.MonthAmountCategoryByIdRequest) ([]*response.MonthAmountCategoryFailedResponse, *response.ErrorResponse)
	FindYearAmountCategoryFailedById(req *requests.YearAmountCategoryByIdRequest) ([]*response.YearAmountCategoryFailedResponse, *response.ErrorResponse)
	FindMonthMethodCategorySuccessById(req *requests.MonthMethodCategoryByIdRequest) ([]*response.MonthMethodCategoryResponse, *response.ErrorResponse)
	FindYearMethodCategorySuccessById(req *requests.YearMethodCategoryByIdRequest) ([]*response.YearMethodCategoryResponse, *response.ErrorResponse)
	FindMonthMethodCategoryFailedById(req *requests.MonthMethodCategoryByIdRequest) ([]*response.MonthMethodCategoryResponse, *response.ErrorResponse)
	FindYearMethodCategoryFailedById(req *requests.YearMethodCategoryByIdRequest) ([]*response.YearMethodCategoryResponse, *response.ErrorResponse)

	FindMonthAmountCategorySuccessByMerchant(req *requests.MonthAmountCategoryByMerchantRequest) ([]*response.MonthAmountCategorySuccessResponse, *response.ErrorResponse)
	FindYearAmountCategorySuccessByMerchant(req *requests.YearAmountCategoryByMerchantRequest) ([]*response.YearAmountCategorySuccessResponse, *response.ErrorResponse)
	FindMonthAmountCategoryFailedByMerchant(req *requests.MonthAmountCategoryByMerchantRequest) ([]*response.MonthAmountCategoryFailedResponse, *response.ErrorResponse)
	FindYearAmountCategoryFailedByMerchant(req *requests.YearAmountCategoryByMerchantRequest) ([]*response.YearAmountCategoryFailedResponse, *response.ErrorResponse)
	FindMonthMethodCategorySuccessByMerchant(req *requests.MonthMethodCategoryByMerchantRequest) ([]*response.MonthMethodCategoryResponse, *response.ErrorResponse)
	FindYearMethodCategorySuccessByMerchant(req *requests.YearMethodCategoryByMerchantRequest) ([]*response.YearMethodCategoryResponse, *response.ErrorResponse)
	FindMonthMethodCategoryFailedByMerchant(req *requests.MonthMethodCategoryByMerchantRequest) ([]*response.MonthMethodCategoryResponse, *response.ErrorResponse)
	FindYearMethodCategoryFailedByMerchant(req *requests.YearMethodCategoryByMerchantRequest) ([]*response.YearMethodCategoryResponse, *response.ErrorResponse)
}

type MerchantService interface {
	FindAll(req *requests.FindAllMerchants) ([]*response.MerchantResponse, *int, *response.ErrorResponse)
	FindByActive(req *requests.FindAllMerchants) ([]*response.MerchantResponseDeleteAt, *int, *response.ErrorResponse)
	FindByTrashed(req *requests.FindAllMerchants) ([]*response.MerchantResponseDeleteAt, *int, *response.ErrorResponse)
	FindById(merchantID int) (*response.MerchantResponse, *response.ErrorResponse)
	Create(req *requests.CreateMerchantRequest) (*response.MerchantResponse, *response.ErrorResponse)
	Update(req *requests.UpdateMerchantRequest) (*response.MerchantResponse, *response.ErrorResponse)
	Trashed(merchantID int) (*response.MerchantResponseDeleteAt, *response.ErrorResponse)
	Restore(merchantID int) (*response.MerchantResponseDeleteAt, *response.ErrorResponse)
	DeletePermanent(merchantID int) (bool, *response.ErrorResponse)
	RestoreAll() (bool, *response.ErrorResponse)
	DeleteAllPermanent() (bool, *response.ErrorResponse)
}

type NominalService interface {
	FindAll(req *requests.FindAllNominals) ([]*response.NominalResponse, *int, *response.ErrorResponse)
	FindByActive(req *requests.FindAllNominals) ([]*response.NominalResponseDeleteAt, *int, *response.ErrorResponse)
	FindByTrashed(req *requests.FindAllNominals) ([]*response.NominalResponseDeleteAt, *int, *response.ErrorResponse)
	FindByID(id int) (*response.NominalResponse, *response.ErrorResponse)
	Create(request *requests.CreateNominalRequest) (*response.NominalResponse, *response.ErrorResponse)
	Update(request *requests.UpdateNominalRequest) (*response.NominalResponse, *response.ErrorResponse)
	Trashed(id int) (*response.NominalResponseDeleteAt, *response.ErrorResponse)
	Restore(id int) (*response.NominalResponseDeleteAt, *response.ErrorResponse)
	DeletePermanent(id int) (bool, *response.ErrorResponse)
	RestoreAll() (bool, *response.ErrorResponse)
	DeleteAllPermanent() (bool, *response.ErrorResponse)

	FindMonthAmountNominalSuccess(req *requests.MonthAmountNominalRequest) ([]*response.MonthAmountNominalSuccessResponse, *response.ErrorResponse)
	FindYearAmountNominalSuccess(year int) ([]*response.YearAmountNominalSuccessResponse, *response.ErrorResponse)
	FindMonthAmountNominalFailed(req *requests.MonthAmountNominalRequest) ([]*response.MonthAmountNominalFailedResponse, *response.ErrorResponse)
	FindYearAmountNominalFailed(year int) ([]*response.YearAmountNominalFailedResponse, *response.ErrorResponse)
	FindMonthMethodNominalSuccess(year int) ([]*response.MonthMethodNominalResponse, *response.ErrorResponse)
	FindYearMethodNominalSuccess(year int) ([]*response.YearMethodNominalResponse, *response.ErrorResponse)
	FindMonthMethodNominalFailed(year int) ([]*response.MonthMethodNominalResponse, *response.ErrorResponse)
	FindYearMethodNominalFailed(year int) ([]*response.YearMethodNominalResponse, *response.ErrorResponse)

	FindMonthAmountNominalSuccessById(req *requests.MonthAmountNominalByIdRequest) ([]*response.MonthAmountNominalSuccessResponse, *response.ErrorResponse)
	FindYearAmountNominalSuccessById(req *requests.YearAmountNominalByIdRequest) ([]*response.YearAmountNominalSuccessResponse, *response.ErrorResponse)
	FindMonthAmountNominalFailedById(req *requests.MonthAmountNominalByIdRequest) ([]*response.MonthAmountNominalFailedResponse, *response.ErrorResponse)
	FindYearAmountNominalFailedById(req *requests.YearAmountNominalByIdRequest) ([]*response.YearAmountNominalFailedResponse, *response.ErrorResponse)
	FindMonthMethodNominalSuccessById(req *requests.MonthMethodNominalByIdRequest) ([]*response.MonthMethodNominalResponse, *response.ErrorResponse)
	FindYearMethodNominalSuccessById(req *requests.YearMethodNominalByIdRequest) ([]*response.YearMethodNominalResponse, *response.ErrorResponse)
	FindMonthMethodNominalFailedById(req *requests.MonthMethodNominalByIdRequest) ([]*response.MonthMethodNominalResponse, *response.ErrorResponse)
	FindYearMethodNominalFailedById(req *requests.YearMethodNominalByIdRequest) ([]*response.YearMethodNominalResponse, *response.ErrorResponse)

	FindMonthAmountNominalSuccessByMerchant(req *requests.MonthAmountNominalByMerchantRequest) ([]*response.MonthAmountNominalSuccessResponse, *response.ErrorResponse)
	FindYearAmountNominalSuccessByMerchant(req *requests.YearAmountNominalByMerchantRequest) ([]*response.YearAmountNominalSuccessResponse, *response.ErrorResponse)
	FindMonthAmountNominalFailedByMerchant(req *requests.MonthAmountNominalByMerchantRequest) ([]*response.MonthAmountNominalFailedResponse, *response.ErrorResponse)
	FindYearAmountNominalFailedByMerchant(req *requests.YearAmountNominalByMerchantRequest) ([]*response.YearAmountNominalFailedResponse, *response.ErrorResponse)
	FindMonthMethodNominalSuccessByMerchant(req *requests.MonthMethodNominalByMerchantRequest) ([]*response.MonthMethodNominalResponse, *response.ErrorResponse)
	FindYearMethodNominalSuccessByMerchant(req *requests.YearMethodNominalByMerchantRequest) ([]*response.YearMethodNominalResponse, *response.ErrorResponse)
	FindMonthMethodNominalFailedByMerchant(req *requests.MonthMethodNominalByMerchantRequest) ([]*response.MonthMethodNominalResponse, *response.ErrorResponse)
	FindYearMethodNominalFailedByMerchant(req *requests.YearMethodNominalByMerchantRequest) ([]*response.YearMethodNominalResponse, *response.ErrorResponse)
}

type TransactionService interface {
	FindAll(req *requests.FindAllTransactions) ([]*response.TransactionResponse, *int, *response.ErrorResponse)
	FindByActive(req *requests.FindAllTransactions) ([]*response.TransactionResponseDeleteAt, *int, *response.ErrorResponse)
	FindByTrashed(req *requests.FindAllTransactions) ([]*response.TransactionResponseDeleteAt, *int, *response.ErrorResponse)
	FindById(id int) (*response.TransactionResponse, *response.ErrorResponse)
	Create(request *requests.CreateTransactionRequest) (*response.TransactionResponse, *response.ErrorResponse)
	Update(request *requests.UpdateTransactionRequest) (*response.TransactionResponse, *response.ErrorResponse)
	Trashed(id int) (*response.TransactionResponseDeleteAt, *response.ErrorResponse)
	Restore(id int) (*response.TransactionResponseDeleteAt, *response.ErrorResponse)
	DeletePermanent(id int) (bool, *response.ErrorResponse)
	RestoreAll() (bool, *response.ErrorResponse)
	DeleteAllPermanent() (bool, *response.ErrorResponse)

	FindMonthAmountTransactionSuccess(req *requests.MonthAmountTransactionRequest) ([]*response.MonthAmountTransactionSuccessResponse, *response.ErrorResponse)
	FindYearAmountTransactionSuccess(year int) ([]*response.YearAmountTransactionSuccessResponse, *response.ErrorResponse)
	FindMonthAmountTransactionFailed(req *requests.MonthAmountTransactionRequest) ([]*response.MonthAmountTransactionFailedResponse, *response.ErrorResponse)
	FindYearAmountTransactionFailed(year int) ([]*response.YearAmountTransactionFailedResponse, *response.ErrorResponse)
	FindMonthMethodTransactionSuccess(year int) ([]*response.MonthMethodTransactionResponse, *response.ErrorResponse)
	FindYearMethodTransactionSuccess(year int) ([]*response.YearMethodTransactionResponse, *response.ErrorResponse)
	FindMonthMethodTransactionFailed(year int) ([]*response.MonthMethodTransactionResponse, *response.ErrorResponse)
	FindYearMethodTransactionFailed(year int) ([]*response.YearMethodTransactionResponse, *response.ErrorResponse)

	FindMonthAmountTransactionSuccessByMerchant(req *requests.MonthAmountTransactionByMerchantRequest) ([]*response.MonthAmountTransactionSuccessResponse, *response.ErrorResponse)
	FindYearAmountTransactionSuccessByMerchant(req *requests.YearAmountTransactionByMerchantRequest) ([]*response.YearAmountTransactionSuccessResponse, *response.ErrorResponse)
	FindMonthAmountTransactionFailedByMerchant(req *requests.MonthAmountTransactionByMerchantRequest) ([]*response.MonthAmountTransactionFailedResponse, *response.ErrorResponse)
	FindYearAmountTransactionFailedByMerchant(req *requests.YearAmountTransactionByMerchantRequest) ([]*response.YearAmountTransactionFailedResponse, *response.ErrorResponse)
	FindMonthMethodTransactionSuccessByMerchant(req *requests.MonthMethodTransactionByMerchantRequest) ([]*response.MonthMethodTransactionResponse, *response.ErrorResponse)
	FindYearMethodTransactionSuccessByMerchant(req *requests.YearMethodTransactionByMerchantRequest) ([]*response.YearMethodTransactionResponse, *response.ErrorResponse)
	FindMonthMethodTransactionFailedByMerchant(req *requests.MonthMethodTransactionByMerchantRequest) ([]*response.MonthMethodTransactionResponse, *response.ErrorResponse)
	FindYearMethodTransactionFailedByMerchant(req *requests.YearMethodTransactionByMerchantRequest) ([]*response.YearMethodTransactionResponse, *response.ErrorResponse)
}

type VoucherService interface {
	FindAll(req *requests.FindAllVouchers) ([]*response.VoucherResponse, *int, *response.ErrorResponse)
	FindByActive(req *requests.FindAllVouchers) ([]*response.VoucherResponseDeleteAt, *int, *response.ErrorResponse)
	FindByTrashed(req *requests.FindAllVouchers) ([]*response.VoucherResponseDeleteAt, *int, *response.ErrorResponse)
	FindById(id int) (*response.VoucherResponse, *response.ErrorResponse)
	Create(request *requests.CreateVoucherRequest) (*response.VoucherResponse, *response.ErrorResponse)
	Update(request *requests.UpdateVoucherRequest) (*response.VoucherResponse, *response.ErrorResponse)
	Trashed(id int) (*response.VoucherResponseDeleteAt, *response.ErrorResponse)
	Restore(id int) (*response.VoucherResponseDeleteAt, *response.ErrorResponse)
	DeletePermanent(id int) (bool, *response.ErrorResponse)
	RestoreAll() (bool, *response.ErrorResponse)
	DeleteAllPermanent() (bool, *response.ErrorResponse)

	FindMonthAmountVoucherSuccess(req *requests.MonthAmountVoucherRequest) ([]*response.MonthAmountVoucherSuccessResponse, *response.ErrorResponse)
	FindYearAmountVoucherSuccess(year int) ([]*response.YearAmountVoucherSuccessResponse, *response.ErrorResponse)
	FindMonthAmountVoucherFailed(req *requests.MonthAmountVoucherRequest) ([]*response.MonthAmountVoucherFailedResponse, *response.ErrorResponse)
	FindYearAmountVoucherFailed(year int) ([]*response.YearAmountVoucherFailedResponse, *response.ErrorResponse)
	FindMonthMethodVoucherSuccess(year int) ([]*response.MonthMethodVoucherResponse, *response.ErrorResponse)
	FindYearMethodVoucherSuccess(year int) ([]*response.YearMethodVoucherResponse, *response.ErrorResponse)
	FindMonthMethodVoucherFailed(year int) ([]*response.MonthMethodVoucherResponse, *response.ErrorResponse)
	FindYearMethodVoucherFailed(year int) ([]*response.YearMethodVoucherResponse, *response.ErrorResponse)

	FindMonthAmountVoucherSuccessById(req *requests.MonthAmountVoucherByIdRequest) ([]*response.MonthAmountVoucherSuccessResponse, *response.ErrorResponse)
	FindYearAmountVoucherSuccessById(req *requests.YearAmountVoucherByIdRequest) ([]*response.YearAmountVoucherSuccessResponse, *response.ErrorResponse)
	FindMonthAmountVoucherFailedById(req *requests.MonthAmountVoucherByIdRequest) ([]*response.MonthAmountVoucherFailedResponse, *response.ErrorResponse)
	FindYearAmountVoucherFailedById(req *requests.YearAmountVoucherByIdRequest) ([]*response.YearAmountVoucherFailedResponse, *response.ErrorResponse)
	FindMonthMethodVoucherSuccessById(req *requests.MonthMethodVoucherByIdRequest) ([]*response.MonthMethodVoucherResponse, *response.ErrorResponse)
	FindYearMethodVoucherSuccessById(req *requests.YearMethodVoucherByIdRequest) ([]*response.YearMethodVoucherResponse, *response.ErrorResponse)
	FindMonthMethodVoucherFailedById(req *requests.MonthMethodVoucherByIdRequest) ([]*response.MonthMethodVoucherResponse, *response.ErrorResponse)
	FindYearMethodVoucherFailedById(req *requests.YearMethodVoucherByIdRequest) ([]*response.YearMethodVoucherResponse, *response.ErrorResponse)

	FindMonthAmountVoucherSuccessByMerchant(req *requests.MonthAmountVoucherByMerchantRequest) ([]*response.MonthAmountVoucherSuccessResponse, *response.ErrorResponse)
	FindYearAmountVoucherSuccessByMerchant(req *requests.YearAmountVoucherByMerchantRequest) ([]*response.YearAmountVoucherSuccessResponse, *response.ErrorResponse)
	FindMonthAmountVoucherFailedByMerchant(req *requests.MonthAmountVoucherByMerchantRequest) ([]*response.MonthAmountVoucherFailedResponse, *response.ErrorResponse)
	FindYearAmountVoucherFailedByMerchant(req *requests.YearAmountVoucherByMerchantRequest) ([]*response.YearAmountVoucherFailedResponse, *response.ErrorResponse)
	FindMonthMethodVoucherSuccessByMerchant(req *requests.MonthMethodVoucherByMerchantRequest) ([]*response.MonthMethodVoucherResponse, *response.ErrorResponse)
	FindYearMethodVoucherSuccessByMerchant(req *requests.YearMethodVoucherByMerchantRequest) ([]*response.YearMethodVoucherResponse, *response.ErrorResponse)
	FindMonthMethodVoucherFailedByMerchant(req *requests.MonthMethodVoucherByMerchantRequest) ([]*response.MonthMethodVoucherResponse, *response.ErrorResponse)
	FindYearMethodVoucherFailedByMerchant(req *requests.YearMethodVoucherByMerchantRequest) ([]*response.YearMethodVoucherResponse, *response.ErrorResponse)
}
