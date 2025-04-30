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
	ToCategoriesResponse(Categorys []*record.CategoryRecord) []*response.CategoryResponse
	ToCategoryResponseDeleteAt(Category *record.CategoryRecord) *response.CategoryResponseDeleteAt
	ToCategoriesResponseDeleteAt(Categorys []*record.CategoryRecord) []*response.CategoryResponseDeleteAt
	ToCategoryResponseMonthAmountSuccess(b *record.MonthAmountCategorySuccessRecord) *response.MonthAmountCategorySuccessResponse
	ToCategoriesResponseMonthAmountSuccess(b []*record.MonthAmountCategorySuccessRecord) []*response.MonthAmountCategorySuccessResponse
	ToCategoryResponseYearAmountSuccess(b *record.YearAmountCategorySuccessRecord) *response.YearAmountCategorySuccessResponse
	ToCategoriesResponseYearAmountSuccess(b []*record.YearAmountCategorySuccessRecord) []*response.YearAmountCategorySuccessResponse
	ToCategoryResponseMonthAmountFailed(b *record.MonthAmountCategoryFailedRecord) *response.MonthAmountCategoryFailedResponse
	ToCategoriesResponseMonthAmountFailed(b []*record.MonthAmountCategoryFailedRecord) []*response.MonthAmountCategoryFailedResponse
	ToCategoryResponseYearAmountFailed(b *record.YearAmountCategoryFailedRecord) *response.YearAmountCategoryFailedResponse
	ToCategoriesResponseYearAmountFailed(b []*record.YearAmountCategoryFailedRecord) []*response.YearAmountCategoryFailedResponse

	ToCategoryResponseMonthMethod(b *record.MonthMethodCategoryRecord) *response.MonthMethodCategoryResponse
	ToCategoriesResponseMonthMethod(b []*record.MonthMethodCategoryRecord) []*response.MonthMethodCategoryResponse
	ToCategoryResponseYearMethod(b *record.YearMethodCategoryRecord) *response.YearMethodCategoryResponse
	ToCategoriesResponseYearMethod(b []*record.YearMethodCategoryRecord) []*response.YearMethodCategoryResponse
}

type BankResponseMapper interface {
	ToBankResponse(Bank *record.BankRecord) *response.BankResponse
	ToBanksResponse(Banks []*record.BankRecord) []*response.BankResponse
	ToBankResponseDeleteAt(Bank *record.BankRecord) *response.BankResponseDeleteAt
	ToBanksResponseDeleteAt(Banks []*record.BankRecord) []*response.BankResponseDeleteAt
	ToBanksResponseMonthAmountSuccess(b []*record.MonthAmountBankSuccessRecord) []*response.MonthAmountBankSuccessResponse
	ToBankResponseYearAmountSuccess(b *record.YearAmountBankSuccessRecord) *response.YearAmountBankSuccessResponse
	ToBanksResponseYearAmountSuccess(b []*record.YearAmountBankSuccessRecord) []*response.YearAmountBankSuccessResponse
	ToBankResponseMonthAmountFailed(b *record.MonthAmountBankFailedRecord) *response.MonthAmountBankFailedResponse
	ToBanksResponseMonthAmountFailed(b []*record.MonthAmountBankFailedRecord) []*response.MonthAmountBankFailedResponse
	ToBankResponseYearAmountFailed(b *record.YearAmountBankFailedRecord) *response.YearAmountBankFailedResponse
	ToBanksResponseYearAmountFailed(b []*record.YearAmountBankFailedRecord) []*response.YearAmountBankFailedResponse
	ToBankResponseMonthMethod(b *record.MonthMethodBankRecord) *response.MonthMethodBankResponse
	ToBanksResponseMonthMethod(b []*record.MonthMethodBankRecord) []*response.MonthMethodBankResponse
	ToBankResponseYearMethod(b *record.YearMethodBankRecord) *response.YearMethodBankResponse
	ToBanksResponseYearMethod(b []*record.YearMethodBankRecord) []*response.YearMethodBankResponse
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
	ToNominalResponseMonthAmountSuccess(b *record.MonthAmountNominalSuccessRecord) *response.MonthAmountNominalSuccessResponse
	ToNominalsResponseMonthAmountSuccess(b []*record.MonthAmountNominalSuccessRecord) []*response.MonthAmountNominalSuccessResponse
	ToNominalResponseYearAmountSuccess(b *record.YearAmountNominalSuccessRecord) *response.YearAmountNominalSuccessResponse
	ToNominalsResponseYearAmountSuccess(b []*record.YearAmountNominalSuccessRecord) []*response.YearAmountNominalSuccessResponse
	ToNominalResponseMonthAmountFailed(b *record.MonthAmountNominalFailedRecord) *response.MonthAmountNominalFailedResponse
	ToNominalsResponseMonthAmountFailed(b []*record.MonthAmountNominalFailedRecord) []*response.MonthAmountNominalFailedResponse
	ToNominalResponseYearAmountFailed(b *record.YearAmountNominalFailedRecord) *response.YearAmountNominalFailedResponse
	ToNominalsResponseYearAmountFailed(b []*record.YearAmountNominalFailedRecord) []*response.YearAmountNominalFailedResponse
	ToNominalResponseMonthMethodSuccess(b *record.MonthMethodNominalRecord) *response.MonthMethodNominalResponse
	ToNominalsResponseMonthMethodSuccess(b []*record.MonthMethodNominalRecord) []*response.MonthMethodNominalResponse
	ToNominalResponseMonthMethodFailed(b *record.MonthMethodNominalRecord) *response.MonthMethodNominalResponse
	ToNominalsResponseMonthMethodFailed(b []*record.MonthMethodNominalRecord) []*response.MonthMethodNominalResponse
	ToNominalResponseYearMethod(b *record.YearMethodNominalRecord) *response.YearMethodNominalResponse
	ToNominalsResponseYearMethod(b []*record.YearMethodNominalRecord) []*response.YearMethodNominalResponse
}

type TransactionResponseMapper interface {
	ToTransactionResponse(transaction *record.TransactionRecord) *response.TransactionResponse
	ToTransactionsResponse(transactions []*record.TransactionRecord) []*response.TransactionResponse
	ToTransactionResponseDeleteAt(transaction *record.TransactionRecord) *response.TransactionResponseDeleteAt
	ToTransactionsResponseDeleteAt(transactions []*record.TransactionRecord) []*response.TransactionResponseDeleteAt
	ToTransactionResponseMonthAmountSuccess(b *record.MonthAmountTransactionSuccessRecord) *response.MonthAmountTransactionSuccessResponse
	ToTransactionsResponseMonthAmountSuccess(b []*record.MonthAmountTransactionSuccessRecord) []*response.MonthAmountTransactionSuccessResponse
	ToTransactionResponseYearAmountSuccess(b *record.YearAmountTransactionSuccessRecord) *response.YearAmountTransactionSuccessResponse
	ToTransactionsResponseYearAmountSuccess(b []*record.YearAmountTransactionSuccessRecord) []*response.YearAmountTransactionSuccessResponse
	ToTransactionResponseMonthAmountFailed(b *record.MonthAmountTransactionFailedRecord) *response.MonthAmountTransactionFailedResponse
	ToTransactionsResponseMonthAmountFailed(b []*record.MonthAmountTransactionFailedRecord) []*response.MonthAmountTransactionFailedResponse
	ToTransactionResponseYearAmountFailed(b *record.YearAmountTransactionFailedRecord) *response.YearAmountTransactionFailedResponse
	ToTransactionsResponseYearAmountFailed(b []*record.YearAmountTransactionFailedRecord) []*response.YearAmountTransactionFailedResponse
	ToTransactionResponseMonthMethod(b *record.MonthMethodTransactionRecord) *response.MonthMethodTransactionResponse
	ToTransactionsResponseMonthMethod(b []*record.MonthMethodTransactionRecord) []*response.MonthMethodTransactionResponse
	ToTransactionResponseYearMethod(b *record.YearMethodTransactionRecord) *response.YearMethodTransactionResponse
	ToTransactionsResponseYearMethod(b []*record.YearMethodTransactionRecord) []*response.YearMethodTransactionResponse
}

type VoucherResponseMapper interface {
	ToVoucherResponse(voucher *record.VoucherRecord) *response.VoucherResponse
	ToVouchersResponse(vouchers []*record.VoucherRecord) []*response.VoucherResponse
	ToVoucherResponseDeleteAt(voucher *record.VoucherRecord) *response.VoucherResponseDeleteAt
	ToVouchersResponseDeleteAt(vouchers []*record.VoucherRecord) []*response.VoucherResponseDeleteAt
	ToVoucherResponseMonthAmountSuccess(b *record.MonthAmountVoucherSuccessRecord) *response.MonthAmountVoucherSuccessResponse
	ToVouchersResponseMonthAmountSuccess(b []*record.MonthAmountVoucherSuccessRecord) []*response.MonthAmountVoucherSuccessResponse
	ToVoucherResponseYearAmountSuccess(b *record.YearAmountVoucherSuccessRecord) *response.YearAmountVoucherSuccessResponse
	ToVouchersResponseYearAmountSuccess(b []*record.YearAmountVoucherSuccessRecord) []*response.YearAmountVoucherSuccessResponse
	ToVoucherResponseMonthAmountFailed(b *record.MonthAmountVoucherFailedRecord) *response.MonthAmountVoucherFailedResponse
	ToVouchersResponseMonthAmountFailed(b []*record.MonthAmountVoucherFailedRecord) []*response.MonthAmountVoucherFailedResponse
	ToVoucherResponseYearAmountFailed(b *record.YearAmountVoucherFailedRecord) *response.YearAmountVoucherFailedResponse
	ToVouchersResponseYearAmountFailed(b []*record.YearAmountVoucherFailedRecord) []*response.YearAmountVoucherFailedResponse
	ToVoucherResponseMonthMethod(b *record.MonthMethodVoucherRecord) *response.MonthMethodVoucherResponse
	ToVouchersResponseMonthMethod(b []*record.MonthMethodVoucherRecord) []*response.MonthMethodVoucherResponse
	ToVoucherResponseYearMethod(b *record.YearMethodVoucherRecord) *response.YearMethodVoucherResponse
	ToVouchersResponseYearMethod(b []*record.YearMethodVoucherRecord) []*response.YearMethodVoucherResponse
}
