package service

import (
	response_service "topup_game/internal/mapper/response/service"
	"topup_game/internal/repository"
	"topup_game/pkg/auth"
	"topup_game/pkg/hash"
	"topup_game/pkg/logger"
)

type Service struct {
	Auth        AuthService
	User        UserService
	Role        RoleService
	Bank        BankService
	Category    CategoryService
	Merchant    MerchantService
	Nominal     NominalService
	Transaction TransactionService
	Voucher     VoucherService
}

type Deps struct {
	Repositories *repository.Repositories
	Token        auth.TokenManager
	Hash         hash.HashPassword
	Logger       logger.LoggerInterface
	Mapper       response_service.ResponseServiceMapper
}

func NewService(deps Deps) *Service {
	return &Service{
		Auth:        NewAuthService(deps.Repositories.User, deps.Repositories.RefreshToken, deps.Repositories.Role, deps.Repositories.UserRole, deps.Hash, deps.Token, deps.Logger, deps.Mapper.UserResponseMapper),
		User:        NewUserService(deps.Repositories.User, deps.Logger, deps.Mapper.UserResponseMapper, deps.Hash),
		Role:        NewRoleService(deps.Repositories.Role, deps.Logger, deps.Mapper.RoleResponseMapper),
		Bank:        NewBankService(deps.Repositories.Bank, deps.Logger, deps.Mapper.BankResponseMapper),
		Category:    NewCategoryService(deps.Repositories.Category, deps.Logger, deps.Mapper.CategoryResponseMapper),
		Merchant:    NewMerchantService(deps.Repositories.Merchant, deps.Logger, deps.Mapper.MerchantResponseMapper),
		Nominal:     NewNominalService(deps.Repositories.Nominal, deps.Repositories.Voucher, deps.Logger, deps.Mapper.NominalResponseMapper),
		Transaction: NewTransactionService(deps.Repositories.User, deps.Repositories.Merchant, deps.Repositories.Voucher, deps.Repositories.Nominal, deps.Repositories.Category, deps.Repositories.Bank, deps.Repositories.Transaction, deps.Logger, deps.Mapper.TransactionResponseMapper),
	}
}
