package response_service

import (
	"topup_game/internal/domain/record"
	"topup_game/internal/domain/response"
)

type UserResponseMapper interface {
	ToUserResponse(user *record.UserRecord) *response.UserResponse
	ToUsersResponse(users []*record.UserRecord) []*response.UserResponse

	ToUserResponseDeleteAt(user *record.UserRecord) *response.UserResponseDeleteAt
	ToUsersResponseDeleteAt(users []*record.UserRecord) []*response.UserResponseDeleteAt
}

type RoleResponseMapper interface {
	ToRoleResponse(role *record.RoleRecord) *response.RoleResponse
	ToRolesResponse(roles []*record.RoleRecord) []*response.RoleResponse

	ToRoleResponseDeleteAt(role *record.RoleRecord) *response.RoleResponseDeleteAt
	ToRolesResponseDeleteAt(roles []*record.RoleRecord) []*response.RoleResponseDeleteAt
}

type RefreshTokenResponseMapper interface {
	ToRefreshTokenResponse(refresh *record.RefreshTokenRecord) *response.RefreshTokenResponse
	ToRefreshTokenResponses(refreshs []*record.RefreshTokenRecord) []*response.RefreshTokenResponse
}

type CategoryResponseMapper interface {
	ToCategoryResponse(Category *record.CategoryRecord) *response.CategoryResponse
	ToCategorysResponse(Categorys []*record.CategoryRecord) []*response.CategoryResponse
	ToCategoryResponseDeleteAt(Category *record.CategoryRecord) *response.CategoryResponseDeleteAt
	ToCategorysResponseDeleteAt(Categorys []*record.CategoryRecord) []*response.CategoryResponseDeleteAt
}

type BankResponseMapper interface {
	ToBankResponse(Bank *record.BankRecord) *response.BankResponse
	ToBanksResponse(Banks []*record.BankRecord) []*response.BankResponse
	ToBankResponseDeleteAt(Bank *record.BankRecord) *response.BankResponseDeleteAt
	ToBanksResponseDeleteAt(Banks []*record.BankRecord) []*response.BankResponseDeleteAt
}

type MerchantResponseMapper interface {
	ToMerchantResponse(merchant *record.MerchantRecord) *response.MerchantResponse
	ToMerchantsResponse(merchants []*record.MerchantRecord) []*response.MerchantResponse
	ToMerchantResponseDeleteAt(merchant *record.MerchantRecord) *response.MerchantResponseDeleteAt
	ToMerchantsResponseDeleteAt(merchants []*record.MerchantRecord) []*response.MerchantResponseDeleteAt
}

type NominalResponseMapper interface {
	ToNominalResponse(nominal *record.NominalRecord) *response.NominalResponse
	ToNominalsResponse(nominals []*record.NominalRecord) []*response.NominalResponse
	ToNominalResponseDeleteAt(nominal *record.NominalRecord) *response.NominalResponseDeleteAt
	ToNominalsResponseDeleteAt(nominals []*record.NominalRecord) []*response.NominalResponseDeleteAt
}

type TransactionResponseMapper interface {
	ToTransactionResponse(transaction *record.TransactionRecord) *response.TransactionResponse
	ToTransactionsResponse(transactions []*record.TransactionRecord) []*response.TransactionResponse
	ToTransactionResponseDeleteAt(transaction *record.TransactionRecord) *response.TransactionResponseDeleteAt
	ToTransactionsResponseDeleteAt(transactions []*record.TransactionRecord) []*response.TransactionResponseDeleteAt
}

type VoucherResponseMapper interface {
	ToVoucherResponse(voucher *record.VoucherRecord) *response.VoucherResponse
	ToVouchersResponse(vouchers []*record.VoucherRecord) []*response.VoucherResponse
	ToVoucherResponseDeleteAt(voucher *record.VoucherRecord) *response.VoucherResponseDeleteAt
	ToVouchersResponseDeleteAt(vouchers []*record.VoucherRecord) []*response.VoucherResponseDeleteAt
}
