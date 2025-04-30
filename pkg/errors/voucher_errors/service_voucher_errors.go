package voucher_errors

import (
	"net/http"
	"topup_game/internal/domain/response"
)

var (
	ErrVoucherNotFoundRes = response.NewErrorResponse("Voucher not found", http.StatusNotFound)
	ErrFailedFindAll      = response.NewErrorResponse("Failed to fetch Vouchers", http.StatusInternalServerError)
	ErrFailedFindActive   = response.NewErrorResponse("Failed to fetch active Vouchers", http.StatusInternalServerError)
	ErrFailedFindTrashed  = response.NewErrorResponse("Failed to fetch trashed Vouchers", http.StatusInternalServerError)

	ErrFailedCreateVoucher = response.NewErrorResponse("Failed to create Voucher", http.StatusInternalServerError)
	ErrFailedUpdateVoucher = response.NewErrorResponse("Failed to update Voucher", http.StatusInternalServerError)

	ErrFailedTrashedVoucher  = response.NewErrorResponse("Failed to move Voucher to trash", http.StatusInternalServerError)
	ErrFailedRestoreVoucher  = response.NewErrorResponse("Failed to restore Voucher", http.StatusInternalServerError)
	ErrFailedDeletePermanent = response.NewErrorResponse("Failed to delete Voucher permanently", http.StatusInternalServerError)

	ErrFailedRestoreAll = response.NewErrorResponse("Failed to restore all Vouchers", http.StatusInternalServerError)
	ErrFailedDeleteAll  = response.NewErrorResponse("Failed to delete all Vouchers permanently", http.StatusInternalServerError)

	ErrFailedFindMonthAmountVoucherSuccess = response.NewErrorResponse("failed to find monthly Voucher success amounts", http.StatusInternalServerError)
	ErrFailedFindYearAmountVoucherSuccess  = response.NewErrorResponse("failed to find yearly Voucher success amounts", http.StatusInternalServerError)
	ErrFailedFindMonthAmountVoucherFailed  = response.NewErrorResponse("failed to find monthly Voucher failed amounts", http.StatusInternalServerError)
	ErrFailedFindYearAmountVoucherFailed   = response.NewErrorResponse("failed to find yearly Voucher failed amounts", http.StatusInternalServerError)

	ErrFailedFindMonthMethodVoucherSuccess = response.NewErrorResponse("failed to find monthly Voucher success methods", http.StatusInternalServerError)
	ErrFailedFindYearMethodVoucherSuccess  = response.NewErrorResponse("failed to find yearly Voucher success methods", http.StatusInternalServerError)
	ErrFailedFindMonthMethodVoucherFailed  = response.NewErrorResponse("failed to find monthly Voucher failed methods", http.StatusInternalServerError)
	ErrFailedFindYearMethodVoucherFailed   = response.NewErrorResponse("failed to find yearly Voucher failed methods", http.StatusInternalServerError)

	ErrFailedFindMonthAmountVoucherSuccessById = response.NewErrorResponse("failed to find monthly Voucher success amounts by ID", http.StatusInternalServerError)
	ErrFailedFindYearAmountVoucherSuccessById  = response.NewErrorResponse("failed to find yearly Voucher success amounts by ID", http.StatusInternalServerError)
	ErrFailedFindMonthAmountVoucherFailedById  = response.NewErrorResponse("failed to find monthly Voucher failed amounts by ID", http.StatusInternalServerError)
	ErrFailedFindYearAmountVoucherFailedById   = response.NewErrorResponse("failed to find yearly Voucher failed amounts by ID", http.StatusInternalServerError)

	ErrFailedFindMonthMethodVoucherSuccessById = response.NewErrorResponse("failed to find monthly Voucher success methods by ID", http.StatusInternalServerError)
	ErrFailedFindYearMethodVoucherSuccessById  = response.NewErrorResponse("failed to find yearly Voucher success methods by ID", http.StatusInternalServerError)
	ErrFailedFindMonthMethodVoucherFailedById  = response.NewErrorResponse("failed to find monthly Voucher failed methods by ID", http.StatusInternalServerError)
	ErrFailedFindYearMethodVoucherFailedById   = response.NewErrorResponse("failed to find yearly Voucher failed methods by ID", http.StatusInternalServerError)

	ErrFailedFindMonthAmountVoucherSuccessByMerchant = response.NewErrorResponse("failed to find monthly Voucher success amounts by merchant", http.StatusInternalServerError)
	ErrFailedFindYearAmountVoucherSuccessByMerchant  = response.NewErrorResponse("failed to find yearly Voucher success amounts by merchant", http.StatusInternalServerError)
	ErrFailedFindMonthAmountVoucherFailedByMerchant  = response.NewErrorResponse("failed to find monthly Voucher failed amounts by merchant", http.StatusInternalServerError)
	ErrFailedFindYearAmountVoucherFailedByMerchant   = response.NewErrorResponse("failed to find yearly Voucher failed amounts by merchant", http.StatusInternalServerError)

	ErrFailedFindMonthMethodVoucherSuccessByMerchant = response.NewErrorResponse("failed to find monthly Voucher success methods by merchant", http.StatusInternalServerError)
	ErrFailedFindYearMethodVoucherSuccessByMerchant  = response.NewErrorResponse("failed to find yearly Voucher success methods by merchant", http.StatusInternalServerError)
	ErrFailedFindMonthMethodVoucherFailedByMerchant  = response.NewErrorResponse("failed to find monthly Voucher failed methods by merchant", http.StatusInternalServerError)
	ErrFailedFindYearMethodVoucherFailedByMerchant   = response.NewErrorResponse("failed to find yearly Voucher failed methods by merchant", http.StatusInternalServerError)
)
