package transaction_errors

import (
	"net/http"
	"topup_game/internal/domain/response"
)

var (
	ErrTransactionNotFoundRes = response.NewErrorResponse("transaction not found", http.StatusNotFound)
	ErrFailedFindAll          = response.NewErrorResponse("Failed to fetch Transactions", http.StatusInternalServerError)
	ErrFailedFindActive       = response.NewErrorResponse("Failed to fetch active Transactions", http.StatusInternalServerError)
	ErrFailedFindTrashed      = response.NewErrorResponse("Failed to fetch trashed Transactions", http.StatusInternalServerError)

	ErrFailedCreateTransaction = response.NewErrorResponse("Failed to create Transaction", http.StatusInternalServerError)
	ErrFailedUpdateTransaction = response.NewErrorResponse("Failed to update Transaction", http.StatusInternalServerError)

	ErrFailedTrashedTransaction = response.NewErrorResponse("Failed to move Transaction to trash", http.StatusInternalServerError)
	ErrFailedRestoreTransaction = response.NewErrorResponse("Failed to restore Transaction", http.StatusInternalServerError)
	ErrFailedDeletePermanent    = response.NewErrorResponse("Failed to delete Transaction permanently", http.StatusInternalServerError)

	ErrFailedRestoreAll = response.NewErrorResponse("Failed to restore all Transactions", http.StatusInternalServerError)
	ErrFailedDeleteAll  = response.NewErrorResponse("Failed to delete all Transactions permanently", http.StatusInternalServerError)

	ErrFailedFindMonthAmountTransactionSuccess = response.NewErrorResponse("failed to find monthly Transaction success amounts", http.StatusInternalServerError)
	ErrFailedFindYearAmountTransactionSuccess  = response.NewErrorResponse("failed to find yearly Transaction success amounts", http.StatusInternalServerError)
	ErrFailedFindMonthAmountTransactionFailed  = response.NewErrorResponse("failed to find monthly Transaction failed amounts", http.StatusInternalServerError)
	ErrFailedFindYearAmountTransactionFailed   = response.NewErrorResponse("failed to find yearly Transaction failed amounts", http.StatusInternalServerError)

	ErrFailedFindMonthMethodTransactionSuccess = response.NewErrorResponse("failed to find monthly Transaction success methods", http.StatusInternalServerError)
	ErrFailedFindYearMethodTransactionSuccess  = response.NewErrorResponse("failed to find yearly Transaction success methods", http.StatusInternalServerError)
	ErrFailedFindMonthMethodTransactionFailed  = response.NewErrorResponse("failed to find monthly Transaction failed methods", http.StatusInternalServerError)
	ErrFailedFindYearMethodTransactionFailed   = response.NewErrorResponse("failed to find yearly Transaction failed methods", http.StatusInternalServerError)

	ErrFailedFindMonthAmountTransactionSuccessById = response.NewErrorResponse("failed to find monthly Transaction success amounts by ID", http.StatusInternalServerError)
	ErrFailedFindYearAmountTransactionSuccessById  = response.NewErrorResponse("failed to find yearly Transaction success amounts by ID", http.StatusInternalServerError)
	ErrFailedFindMonthAmountTransactionFailedById  = response.NewErrorResponse("failed to find monthly Transaction failed amounts by ID", http.StatusInternalServerError)
	ErrFailedFindYearAmountTransactionFailedById   = response.NewErrorResponse("failed to find yearly Transaction failed amounts by ID", http.StatusInternalServerError)

	ErrFailedFindMonthMethodTransactionSuccessById = response.NewErrorResponse("failed to find monthly Transaction success methods by ID", http.StatusInternalServerError)
	ErrFailedFindYearMethodTransactionSuccessById  = response.NewErrorResponse("failed to find yearly Transaction success methods by ID", http.StatusInternalServerError)
	ErrFailedFindMonthMethodTransactionFailedById  = response.NewErrorResponse("failed to find monthly Transaction failed methods by ID", http.StatusInternalServerError)
	ErrFailedFindYearMethodTransactionFailedById   = response.NewErrorResponse("failed to find yearly Transaction failed methods by ID", http.StatusInternalServerError)

	ErrFailedFindMonthAmountTransactionSuccessByMerchant = response.NewErrorResponse("failed to find monthly Transaction success amounts by merchant", http.StatusInternalServerError)
	ErrFailedFindYearAmountTransactionSuccessByMerchant  = response.NewErrorResponse("failed to find yearly Transaction success amounts by merchant", http.StatusInternalServerError)
	ErrFailedFindMonthAmountTransactionFailedByMerchant  = response.NewErrorResponse("failed to find monthly Transaction failed amounts by merchant", http.StatusInternalServerError)
	ErrFailedFindYearAmountTransactionFailedByMerchant   = response.NewErrorResponse("failed to find yearly Transaction failed amounts by merchant", http.StatusInternalServerError)

	ErrFailedFindMonthMethodTransactionSuccessByMerchant = response.NewErrorResponse("failed to find monthly Transaction success methods by merchant", http.StatusInternalServerError)
	ErrFailedFindYearMethodTransactionSuccessByMerchant  = response.NewErrorResponse("failed to find yearly Transaction success methods by merchant", http.StatusInternalServerError)
	ErrFailedFindMonthMethodTransactionFailedByMerchant  = response.NewErrorResponse("failed to find monthly Transaction failed methods by merchant", http.StatusInternalServerError)
	ErrFailedFindYearMethodTransactionFailedByMerchant   = response.NewErrorResponse("failed to find yearly Transaction failed methods by merchant", http.StatusInternalServerError)
)
