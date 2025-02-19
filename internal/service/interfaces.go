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
	FindAll(page int, pageSize int, search string) ([]*response.RoleResponse, int, *response.ErrorResponse)
	FindByActive(page int, pageSize int, search string) ([]*response.RoleResponseDeleteAt, int, *response.ErrorResponse)
	FindByTrashed(page int, pageSize int, search string) ([]*response.RoleResponseDeleteAt, int, *response.ErrorResponse)
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
	FindAll(page int, pageSize int, search string) ([]*response.UserResponse, int, *response.ErrorResponse)
	FindByID(id int) (*response.UserResponse, *response.ErrorResponse)
	FindByActive(page int, pageSize int, search string) ([]*response.UserResponseDeleteAt, int, *response.ErrorResponse)
	FindByTrashed(page int, pageSize int, search string) ([]*response.UserResponseDeleteAt, int, *response.ErrorResponse)
	Create(request *requests.CreateUserRequest) (*response.UserResponse, *response.ErrorResponse)
	Update(request *requests.UpdateUserRequest) (*response.UserResponse, *response.ErrorResponse)
	Trashed(user_id int) (*response.UserResponseDeleteAt, *response.ErrorResponse)
	Restore(user_id int) (*response.UserResponseDeleteAt, *response.ErrorResponse)
	DeletePermanent(user_id int) (bool, *response.ErrorResponse)

	RestoreAll() (bool, *response.ErrorResponse)
	DeleteAllPermanent() (bool, *response.ErrorResponse)
}

type BankService interface {
	FindAll(page int, pageSize int, search string) ([]*response.BankResponse, int, *response.ErrorResponse)
	FindByID(id int) (*response.BankResponse, *response.ErrorResponse)
	FindByActive(page int, pageSize int, search string) ([]*response.BankResponseDeleteAt, int, *response.ErrorResponse)
	FindByTrashed(page int, pageSize int, search string) ([]*response.BankResponseDeleteAt, int, *response.ErrorResponse)
	Create(request *requests.CreateBankRequest) (*response.BankResponse, *response.ErrorResponse)
	Update(request *requests.UpdateBankRequest) (*response.BankResponse, *response.ErrorResponse)
	Trashed(id int) (*response.BankResponseDeleteAt, *response.ErrorResponse)
	Restore(id int) (*response.BankResponseDeleteAt, *response.ErrorResponse)
	DeletePermanent(id int) (bool, *response.ErrorResponse)
	RestoreAll() (bool, *response.ErrorResponse)
	DeleteAllPermanent() (bool, *response.ErrorResponse)
}

type CategoryService interface {
	FindAll(page int, pageSize int, search string) ([]*response.CategoryResponse, int, *response.ErrorResponse)
	FindById(id int) (*response.CategoryResponse, *response.ErrorResponse)
	FindByActive(page int, pageSize int, search string) ([]*response.CategoryResponseDeleteAt, int, *response.ErrorResponse)
	FindByTrashed(page int, pageSize int, search string) ([]*response.CategoryResponseDeleteAt, int, *response.ErrorResponse)
	Create(request *requests.CreateCategoryRequest) (*response.CategoryResponse, *response.ErrorResponse)
	Update(request *requests.UpdateCategoryRequest) (*response.CategoryResponse, *response.ErrorResponse)
	Trashed(id int) (*response.CategoryResponseDeleteAt, *response.ErrorResponse)
	Restore(id int) (*response.CategoryResponseDeleteAt, *response.ErrorResponse)
	DeletePermanent(id int) (bool, *response.ErrorResponse)
	RestoreAll() (bool, *response.ErrorResponse)
	DeleteAllPermanent() (bool, *response.ErrorResponse)
}

type MerchantService interface {
	FindAll(page, pageSize int, search string) ([]*response.MerchantResponse, int, *response.ErrorResponse)
	FindByActive(search string, page, pageSize int) ([]*response.MerchantResponseDeleteAt, int, *response.ErrorResponse)
	FindByTrashed(search string, page, pageSize int) ([]*response.MerchantResponseDeleteAt, int, *response.ErrorResponse)
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
	FindAll(page int, pageSize int, search string) ([]*response.NominalResponse, int, *response.ErrorResponse)
	FindByID(id int) (*response.NominalResponse, *response.ErrorResponse)
	FindByActive(page int, pageSize int, search string) ([]*response.NominalResponseDeleteAt, int, *response.ErrorResponse)
	FindByTrashed(page int, pageSize int, search string) ([]*response.NominalResponseDeleteAt, int, *response.ErrorResponse)
	Create(request *requests.CreateNominalRequest) (*response.NominalResponse, *response.ErrorResponse)
	Update(request *requests.UpdateNominalRequest) (*response.NominalResponse, *response.ErrorResponse)
	Trashed(id int) (*response.NominalResponseDeleteAt, *response.ErrorResponse)
	Restore(id int) (*response.NominalResponseDeleteAt, *response.ErrorResponse)
	DeletePermanent(id int) (bool, *response.ErrorResponse)
	RestoreAll() (bool, *response.ErrorResponse)
	DeleteAllPermanent() (bool, *response.ErrorResponse)
}

type TransactionService interface {
	FindAll(page int, pageSize int, search string) ([]*response.TransactionResponse, int, *response.ErrorResponse)
	FindById(id int) (*response.TransactionResponse, *response.ErrorResponse)
	FindByActive(page int, pageSize int, search string) ([]*response.TransactionResponseDeleteAt, int, *response.ErrorResponse)
	FindByTrashed(page int, pageSize int, search string) ([]*response.TransactionResponseDeleteAt, int, *response.ErrorResponse)
	Create(request *requests.CreateTransactionRequest) (*response.TransactionResponse, *response.ErrorResponse)
	Update(request *requests.UpdateTransactionRequest) (*response.TransactionResponse, *response.ErrorResponse)
	Trashed(id int) (*response.TransactionResponseDeleteAt, *response.ErrorResponse)
	Restore(id int) (*response.TransactionResponseDeleteAt, *response.ErrorResponse)
	DeletePermanent(id int) (bool, *response.ErrorResponse)
	RestoreAll() (bool, *response.ErrorResponse)
	DeleteAllPermanent() (bool, *response.ErrorResponse)
}

type VoucherService interface {
	FindAll(page int, pageSize int, search string) ([]*response.VoucherResponse, int, *response.ErrorResponse)
	FindById(id int) (*response.VoucherResponse, *response.ErrorResponse)
	FindByActive(page int, pageSize int, search string) ([]*response.VoucherResponseDeleteAt, int, *response.ErrorResponse)
	FindByTrashed(page int, pageSize int, search string) ([]*response.VoucherResponseDeleteAt, int, *response.ErrorResponse)
	Create(request *requests.CreateVoucherRequest) (*response.VoucherResponse, *response.ErrorResponse)
	Update(request *requests.UpdateVoucherRequest) (*response.VoucherResponse, *response.ErrorResponse)
	Trashed(id int) (*response.VoucherResponseDeleteAt, *response.ErrorResponse)
	Restore(id int) (*response.VoucherResponseDeleteAt, *response.ErrorResponse)
	DeletePermanent(id int) (bool, *response.ErrorResponse)
	RestoreAll() (bool, *response.ErrorResponse)
	DeleteAllPermanent() (bool, *response.ErrorResponse)
}
