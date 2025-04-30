package protomapper

import (
	"topup_game/internal/domain/response"
	"topup_game/internal/pb"
)

type AuthProtoMapper interface {
	ToProtoResponseLogin(status string, message string, response *response.TokenResponse) *pb.ApiResponseLogin
	ToProtoResponseRegister(status string, message string, response *response.UserResponse) *pb.ApiResponseRegister
	ToProtoResponseRefreshToken(status string, message string, response *response.TokenResponse) *pb.ApiResponseRefreshToken
	ToProtoResponseGetMe(status string, message string, response *response.UserResponse) *pb.ApiResponseGetMe
}

type UserProtoMapper interface {
	ToProtoResponseUserDeleteAt(status string, message string, pbResponse *response.UserResponseDeleteAt) *pb.ApiResponseUserDeleteAt
	ToProtoResponsesUser(status string, message string, pbResponse []*response.UserResponse) *pb.ApiResponsesUser
	ToProtoResponseUser(status string, message string, pbResponse *response.UserResponse) *pb.ApiResponseUser
	ToProtoResponseUserDelete(status string, message string) *pb.ApiResponseUserDelete
	ToProtoResponseUserAll(status string, message string) *pb.ApiResponseUserAll
	ToProtoResponsePaginationUserDeleteAt(pagination *pb.PaginationMeta, status string, message string, users []*response.UserResponseDeleteAt) *pb.ApiResponsePaginationUserDeleteAt
	ToProtoResponsePaginationUser(pagination *pb.PaginationMeta, status string, message string, users []*response.UserResponse) *pb.ApiResponsePaginationUser
}

type RoleProtoMapper interface {
	ToProtoResponseRoleAll(status string, message string) *pb.ApiResponseRoleAll
	ToProtoResponseRoleDelete(status string, message string) *pb.ApiResponseRoleDelete
	ToProtoResponseRole(status string, message string, pbResponse *response.RoleResponse) *pb.ApiResponseRole
	ToProtoResponsesRole(status string, message string, pbResponse []*response.RoleResponse) *pb.ApiResponsesRole
	ToProtoResponsePaginationRole(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.RoleResponse) *pb.ApiResponsePaginationRole
	ToProtoResponsePaginationRoleDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.RoleResponseDeleteAt) *pb.ApiResponsePaginationRoleDeleteAt
}

type BankProtoMapper interface {
	ToProtoResponseBankAll(status string, message string) *pb.ApiResponseBankAll
	ToProtoResponseBankDelete(status string, message string) *pb.ApiResponseBankDelete
	ToProtoResponseBankDeleteAt(status string, message string, pbResponse *response.BankResponseDeleteAt) *pb.ApiResponseBankDeleteAt
	ToProtoResponseBank(status string, message string, pbResponse *response.BankResponse) *pb.ApiResponseBank
	ToProtoResponsesBank(status string, message string, pbResponse []*response.BankResponse) *pb.ApiResponsesBank
	ToProtoResponsePaginationBank(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.BankResponse) *pb.ApiResponsePaginationBank
	ToProtoResponsePaginationBankDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.BankResponseDeleteAt) *pb.ApiResponsePaginationBankDeleteAt
	ToProtoResponsesMonthAmountSuccess(status string, message string, pbResponse []*response.MonthAmountBankSuccessResponse) *pb.ApiResponseBankMonthAmountSuccess
	ToProtoResponseYearAmountSuccess(status string, message string, pbResponse []*response.YearAmountBankSuccessResponse) *pb.ApiResponseBankYearAmountSuccess
	ToProtoResponsesMonthAmountFailed(status string, message string, pbResponse []*response.MonthAmountBankFailedResponse) *pb.ApiResponseBankMonthAmountFailed
	ToProtoResponseYearAmountFailed(status string, message string, pbResponse []*response.YearAmountBankFailedResponse) *pb.ApiResponseBankYearAmountFailed
	ToProtoResponsesMonthMethod(status string, message string, pbResponse []*response.MonthMethodBankResponse) *pb.ApiResponseBankMonthMethod
	ToProtoResponseYearMethod(status string, message string, pbResponse []*response.YearMethodBankResponse) *pb.ApiResponseBankYearMethod
}

type CategoryProtoMapper interface {
	ToProtoResponseCategoryAll(status string, message string) *pb.ApiResponseCategoryAll
	ToProtoResponseCategoryDelete(status string, message string) *pb.ApiResponseCategoryDelete
	ToProtoResponseCategoryDeleteAt(status string, message string, pbResponse *response.CategoryResponseDeleteAt) *pb.ApiResponseCategoryDeleteAt
	ToProtoResponseCategory(status string, message string, pbResponse *response.CategoryResponse) *pb.ApiResponseCategory
	ToProtoResponsesCategory(status string, message string, pbResponse []*response.CategoryResponse) *pb.ApiResponsesCategory
	ToProtoResponsePaginationCategory(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.CategoryResponse) *pb.ApiResponsePaginationCategory
	ToProtoResponsePaginationCategoryDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.CategoryResponseDeleteAt) *pb.ApiResponsePaginationCategoryDeleteAt
	ToProtoResponsesMonthAmountSuccess(status string, message string, pbResponse []*response.MonthAmountCategorySuccessResponse) *pb.ApiResponseCategoryMonthAmountSuccess
	ToProtoResponseYearAmountSuccess(status string, message string, pbResponse []*response.YearAmountCategorySuccessResponse) *pb.ApiResponseCategoryYearAmountSuccess
	ToProtoResponsesMonthAmountFailed(status string, message string, pbResponse []*response.MonthAmountCategoryFailedResponse) *pb.ApiResponseCategoryMonthAmountFailed
	ToProtoResponseYearAmountFailed(status string, message string, pbResponse []*response.YearAmountCategoryFailedResponse) *pb.ApiResponseCategoryYearAmountFailed
	ToProtoResponsesMonthMethod(status string, message string, pbResponse []*response.MonthMethodCategoryResponse) *pb.ApiResponseCategoryMonthMethod
	ToProtoResponseYearMethod(status string, message string, pbResponse []*response.YearMethodCategoryResponse) *pb.ApiResponseCategoryYearMethod
}

type MerchantProtoMapper interface {
	ToProtoResponseMerchant(status string, message string, pbResponse *response.MerchantResponse) *pb.ApiResponseMerchant
	ToProtoResponseMerchantDeleteAt(status string, message string, pbResponse *response.MerchantResponseDeleteAt) *pb.ApiResponseMerchantDeleteAt

	ToProtoResponsesMerchant(status string, message string, pbResponse []*response.MerchantResponse) *pb.ApiResponsesMerchant
	ToProtoResponseMerchantDelete(status string, message string) *pb.ApiResponseMerchantDelete
	ToProtoResponseMerchantAll(status string, message string) *pb.ApiResponseMerchantAll
	ToProtoResponsePaginationMerchantDeleteAt(pagination *pb.PaginationMeta, status string, message string, merchants []*response.MerchantResponseDeleteAt) *pb.ApiResponsePaginationMerchantDeleteAt
	ToProtoResponsePaginationMerchant(pagination *pb.PaginationMeta, status string, message string, merchants []*response.MerchantResponse) *pb.ApiResponsePaginationMerchant
}

type NominalProtoMapper interface {
	ToProtoResponseNominalAll(status string, message string) *pb.ApiResponseNominalAll
	ToProtoResponseNominalDelete(status string, message string) *pb.ApiResponseNominalDelete
	ToProtoResponseNominal(status string, message string, pbResponse *response.NominalResponse) *pb.ApiResponseNominal
	ToProtoResponseNominalDeleteAt(status string, message string, pbResponse *response.NominalResponseDeleteAt) *pb.ApiResponseNominalDeleteAt
	ToProtoResponsesNominal(status string, message string, pbResponse []*response.NominalResponse) *pb.ApiResponsesNominal
	ToProtoResponsePaginationNominal(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.NominalResponse) *pb.ApiResponsePaginationNominal
	ToProtoResponsePaginationNominalDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.NominalResponseDeleteAt) *pb.ApiResponsePaginationNominalDeleteAt
	ToProtoResponsesMonthAmountSuccess(status string, message string, pbResponse []*response.MonthAmountNominalSuccessResponse) *pb.ApiResponseNominalMonthAmountSuccess
	ToProtoResponseYearAmountSuccess(status string, message string, pbResponse []*response.YearAmountNominalSuccessResponse) *pb.ApiResponseNominalYearAmountSuccess
	ToProtoResponsesMonthAmountFailed(status string, message string, pbResponse []*response.MonthAmountNominalFailedResponse) *pb.ApiResponseNominalMonthAmountFailed
	ToProtoResponseYearAmountFailed(status string, message string, pbResponse []*response.YearAmountNominalFailedResponse) *pb.ApiResponseNominalYearAmountFailed
	ToProtoResponsesMonthMethod(status string, message string, pbResponse []*response.MonthMethodNominalResponse) *pb.ApiResponseNominalMonthMethod
	ToProtoResponseYearMethod(status string, message string, pbResponse []*response.YearMethodNominalResponse) *pb.ApiResponseNominalYearMethod
}

type TransactionProtoMapper interface {
	ToProtoResponseTransactionAll(status string, message string) *pb.ApiResponseTransactionAll
	ToProtoResponseTransactionDelete(status string, message string) *pb.ApiResponseTransactionDelete
	ToProtoResponseTransaction(status string, message string, pbResponse *response.TransactionResponse) *pb.ApiResponseTransaction
	ToProtoResponsesTransaction(status string, message string, pbResponse []*response.TransactionResponse) *pb.ApiResponsesTransaction
	ToProtoResponseTransactionDeleteAt(status string, message string, pbResponse *response.TransactionResponseDeleteAt) *pb.ApiResponseTransactionDeleteAt
	ToProtoResponsePaginationTransaction(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.TransactionResponse) *pb.ApiResponsePaginationTransaction
	ToProtoResponsePaginationTransactionDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.TransactionResponseDeleteAt) *pb.ApiResponsePaginationTransactionDeleteAt

	ToProtoResponsesMonthAmountSuccess(status string, message string, pbResponse []*response.MonthAmountTransactionSuccessResponse) *pb.ApiResponseTransactionMonthAmountSuccess
	ToProtoResponseYearAmountSuccess(status string, message string, pbResponse []*response.YearAmountTransactionSuccessResponse) *pb.ApiResponseTransactionYearAmountSuccess
	ToProtoResponsesMonthAmountFailed(status string, message string, pbResponse []*response.MonthAmountTransactionFailedResponse) *pb.ApiResponseTransactionMonthAmountFailed
	ToProtoResponseYearAmountFailed(status string, message string, pbResponse []*response.YearAmountTransactionFailedResponse) *pb.ApiResponseTransactionYearAmountFailed
	ToProtoResponsesMonthMethod(status string, message string, pbResponse []*response.MonthMethodTransactionResponse) *pb.ApiResponseTransactionMonthMethod
	ToProtoResponseYearMethod(status string, message string, pbResponse []*response.YearMethodTransactionResponse) *pb.ApiResponseTransactionYearMethod
}

type VoucherProtoMapper interface {
	ToProtoResponseVoucherAll(status string, message string) *pb.ApiResponseVoucherAll
	ToProtoResponseVoucherDelete(status string, message string) *pb.ApiResponseVoucherDelete
	ToProtoResponseVoucher(status string, message string, pbResponse *response.VoucherResponse) *pb.ApiResponseVoucher
	ToProtoResponseVoucherDeleteAt(status string, message string, pbResponse *response.VoucherResponseDeleteAt) *pb.ApiResponseVoucherDeleteAt
	ToProtoResponsesVoucher(status string, message string, pbResponse []*response.VoucherResponse) *pb.ApiResponsesVoucher
	ToProtoResponsePaginationVoucher(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.VoucherResponse) *pb.ApiResponsePaginationVoucher
	ToProtoResponsePaginationVoucherDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.VoucherResponseDeleteAt) *pb.ApiResponsePaginationVoucherDeleteAt
	ToProtoResponsesMonthAmountSuccess(status string, message string, pbResponse []*response.MonthAmountVoucherSuccessResponse) *pb.ApiResponseVoucherMonthAmountSuccess
	ToProtoResponsesYearAmountSuccess(status string, message string, pbResponse []*response.YearAmountVoucherSuccessResponse) *pb.ApiResponseVoucherYearAmountSuccess
	ToProtoResponsesMonthAmountFailed(status string, message string, pbResponse []*response.MonthAmountVoucherFailedResponse) *pb.ApiResponseVoucherMonthAmountFailed
	ToProtoResponsesYearAmountFailed(status string, message string, pbResponse []*response.YearAmountVoucherFailedResponse) *pb.ApiResponseVoucherYearAmountFailed
	ToProtoResponsesMonthMethod(status string, message string, pbResponse []*response.MonthMethodVoucherResponse) *pb.ApiResponseVoucherMonthMethod
	ToProtoResponsesYearMethod(status string, message string, pbResponse []*response.YearMethodVoucherResponse) *pb.ApiResponseVoucherYearMethod
}
