package bank_errors

import (
	"net/http"
	"topup_game/internal/domain/response"
)

var (
	ErrBankNotFoundRes         = response.NewErrorResponse("Bank not found", http.StatusNotFound)
	ErrFailedFailedFindAll     = response.NewErrorResponse("Failed to fetch Banks", http.StatusInternalServerError)
	ErrFailedFailedFindActive  = response.NewErrorResponse("Failed to fetch active Banks", http.StatusInternalServerError)
	ErrFailedFailedFindTrashed = response.NewErrorResponse("Failed to fetch trashed Banks", http.StatusInternalServerError)

	ErrFailedCreateBank = response.NewErrorResponse("Failed to create Bank", http.StatusInternalServerError)
	ErrFailedUpdateBank = response.NewErrorResponse("Failed to update Bank", http.StatusInternalServerError)

	ErrFailedTrashedBank     = response.NewErrorResponse("Failed to move Bank to trash", http.StatusInternalServerError)
	ErrFailedRestoreBank     = response.NewErrorResponse("Failed to restore Bank", http.StatusInternalServerError)
	ErrFailedDeletePermanent = response.NewErrorResponse("Failed to delete Bank permanently", http.StatusInternalServerError)

	ErrFailedRestoreAll = response.NewErrorResponse("Failed to restore all Banks", http.StatusInternalServerError)
	ErrFailedDeleteAll  = response.NewErrorResponse("Failed to delete all Banks permanently", http.StatusInternalServerError)

	ErrFailedFindMonthAmountBankSuccess = response.NewErrorResponse("failed to find monthly bank success amounts", http.StatusInternalServerError)
	ErrFailedFindYearAmountBankSuccess  = response.NewErrorResponse("failed to FailedFind yearly bank success amounts", http.StatusInternalServerError)
	ErrFailedFindMonthAmountBankFailed  = response.NewErrorResponse("failed to FailedFind monthly bank failed amounts", http.StatusInternalServerError)
	ErrFailedFindYearAmountBankFailed   = response.NewErrorResponse("failed to FailedFind yearly bank failed amounts", http.StatusInternalServerError)

	ErrFailedFindMonthMethodBankSuccess = response.NewErrorResponse("failed to FailedFind monthly bank success methods", http.StatusInternalServerError)
	ErrFailedFindYearMethodBankSuccess  = response.NewErrorResponse("failed to FailedFind yearly bank success methods", http.StatusInternalServerError)
	ErrFailedFindMonthMethodBankFailed  = response.NewErrorResponse("failed to FailedFind monthly bank failed methods", http.StatusInternalServerError)
	ErrFailedFindYearMethodBankFailed   = response.NewErrorResponse("failed to FailedFind yearly bank failed methods", http.StatusInternalServerError)

	ErrFailedFindMonthAmountBankSuccessById = response.NewErrorResponse("failed to FailedFind monthly bank success amounts by ID", http.StatusInternalServerError)
	ErrFailedFindYearAmountBankSuccessById  = response.NewErrorResponse("failed to FailedFind yearly bank success amounts by ID", http.StatusInternalServerError)
	ErrFailedFindMonthAmountBankFailedById  = response.NewErrorResponse("failed to FailedFind monthly bank failed amounts by ID", http.StatusInternalServerError)
	ErrFailedFindYearAmountBankFailedById   = response.NewErrorResponse("failed to FailedFind yearly bank failed amounts by ID", http.StatusInternalServerError)

	ErrFailedFindMonthMethodBankSuccessById = response.NewErrorResponse("failed to FailedFind monthly bank success methods by ID", http.StatusInternalServerError)
	ErrFailedFindYearMethodBankSuccessById  = response.NewErrorResponse("failed to FailedFind yearly bank success methods by ID", http.StatusInternalServerError)
	ErrFailedFindMonthMethodBankFailedById  = response.NewErrorResponse("failed to FailedFind monthly bank failed methods by ID", http.StatusInternalServerError)
	ErrFailedFindYearMethodBankFailedById   = response.NewErrorResponse("failed to FailedFind yearly bank failed methods by ID", http.StatusInternalServerError)

	ErrFailedFindMonthAmountBankSuccessByMerchant = response.NewErrorResponse("failed to FailedFind monthly bank success amounts by merchant", http.StatusInternalServerError)
	ErrFailedFindYearAmountBankSuccessByMerchant  = response.NewErrorResponse("failed to FailedFind yearly bank success amounts by merchant", http.StatusInternalServerError)
	ErrFailedFindMonthAmountBankFailedByMerchant  = response.NewErrorResponse("failed to FailedFind monthly bank failed amounts by merchant", http.StatusInternalServerError)
	ErrFailedFindYearAmountBankFailedByMerchant   = response.NewErrorResponse("failed to FailedFind yearly bank failed amounts by merchant", http.StatusInternalServerError)

	ErrFailedFindMonthMethodBankSuccessByMerchant = response.NewErrorResponse("failed to FailedFind monthly bank success methods by merchant", http.StatusInternalServerError)
	ErrFailedFindYearMethodBankSuccessByMerchant  = response.NewErrorResponse("failed to FailedFind yearly bank success methods by merchant", http.StatusInternalServerError)
	ErrFailedFindMonthMethodBankFailedByMerchant  = response.NewErrorResponse("failed to FailedFind monthly bank failed methods by merchant", http.StatusInternalServerError)
	ErrFailedFindYearMethodBankFailedByMerchant   = response.NewErrorResponse("failed to FailedFind yearly bank failed methods by merchant", http.StatusInternalServerError)
)
