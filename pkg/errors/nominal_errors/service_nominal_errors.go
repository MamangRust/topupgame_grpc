package nominal_errors

import (
	"net/http"
	"topup_game/internal/domain/response"
)

var (
	ErrNominalNotFoundRes = response.NewErrorResponse("Nominal not found", http.StatusNotFound)
	ErrFailedFindAll      = response.NewErrorResponse("Failed to fetch Nominals", http.StatusInternalServerError)
	ErrFailedFindActive   = response.NewErrorResponse("Failed to fetch active Nominals", http.StatusInternalServerError)
	ErrFailedFindTrashed  = response.NewErrorResponse("Failed to fetch trashed Nominals", http.StatusInternalServerError)

	ErrFailedCreateNominal = response.NewErrorResponse("Failed to create Nominal", http.StatusInternalServerError)
	ErrFailedUpdateNominal = response.NewErrorResponse("Failed to update Nominal", http.StatusInternalServerError)

	ErrFailedTrashedNominal  = response.NewErrorResponse("Failed to move Nominal to trash", http.StatusInternalServerError)
	ErrFailedRestoreNominal  = response.NewErrorResponse("Failed to restore Nominal", http.StatusInternalServerError)
	ErrFailedDeletePermanent = response.NewErrorResponse("Failed to delete Nominal permanently", http.StatusInternalServerError)

	ErrFailedRestoreAll = response.NewErrorResponse("Failed to restore all Nominals", http.StatusInternalServerError)
	ErrFailedDeleteAll  = response.NewErrorResponse("Failed to delete all Nominals permanently", http.StatusInternalServerError)

	ErrFailedFindMonthAmountNominalSuccess = response.NewErrorResponse("failed to find monthly Nominal success amounts", http.StatusInternalServerError)
	ErrFailedFindYearAmountNominalSuccess  = response.NewErrorResponse("failed to find yearly Nominal success amounts", http.StatusInternalServerError)
	ErrFailedFindMonthAmountNominalFailed  = response.NewErrorResponse("failed to find monthly Nominal failed amounts", http.StatusInternalServerError)
	ErrFailedFindYearAmountNominalFailed   = response.NewErrorResponse("failed to find yearly Nominal failed amounts", http.StatusInternalServerError)

	ErrFailedFindMonthMethodNominalSuccess = response.NewErrorResponse("failed to find monthly Nominal success methods", http.StatusInternalServerError)
	ErrFailedFindYearMethodNominalSuccess  = response.NewErrorResponse("failed to find yearly Nominal success methods", http.StatusInternalServerError)
	ErrFailedFindMonthMethodNominalFailed  = response.NewErrorResponse("failed to find monthly Nominal failed methods", http.StatusInternalServerError)
	ErrFailedFindYearMethodNominalFailed   = response.NewErrorResponse("failed to find yearly Nominal failed methods", http.StatusInternalServerError)

	ErrFailedFindMonthAmountNominalSuccessById = response.NewErrorResponse("failed to find monthly Nominal success amounts by ID", http.StatusInternalServerError)
	ErrFailedFindYearAmountNominalSuccessById  = response.NewErrorResponse("failed to find yearly Nominal success amounts by ID", http.StatusInternalServerError)
	ErrFailedFindMonthAmountNominalFailedById  = response.NewErrorResponse("failed to find monthly Nominal failed amounts by ID", http.StatusInternalServerError)
	ErrFailedFindYearAmountNominalFailedById   = response.NewErrorResponse("failed to find yearly Nominal failed amounts by ID", http.StatusInternalServerError)

	ErrFailedFindMonthMethodNominalSuccessById = response.NewErrorResponse("failed to find monthly Nominal success methods by ID", http.StatusInternalServerError)
	ErrFailedFindYearMethodNominalSuccessById  = response.NewErrorResponse("failed to find yearly Nominal success methods by ID", http.StatusInternalServerError)
	ErrFailedFindMonthMethodNominalFailedById  = response.NewErrorResponse("failed to find monthly Nominal failed methods by ID", http.StatusInternalServerError)
	ErrFailedFindYearMethodNominalFailedById   = response.NewErrorResponse("failed to find yearly Nominal failed methods by ID", http.StatusInternalServerError)

	ErrFailedFindMonthAmountNominalSuccessByMerchant = response.NewErrorResponse("failed to find monthly Nominal success amounts by merchant", http.StatusInternalServerError)
	ErrFailedFindYearAmountNominalSuccessByMerchant  = response.NewErrorResponse("failed to find yearly Nominal success amounts by merchant", http.StatusInternalServerError)
	ErrFailedFindMonthAmountNominalFailedByMerchant  = response.NewErrorResponse("failed to find monthly Nominal failed amounts by merchant", http.StatusInternalServerError)
	ErrFailedFindYearAmountNominalFailedByMerchant   = response.NewErrorResponse("failed to find yearly Nominal failed amounts by merchant", http.StatusInternalServerError)

	ErrFailedFindMonthMethodNominalSuccessByMerchant = response.NewErrorResponse("failed to find monthly Nominal success methods by merchant", http.StatusInternalServerError)
	ErrFailedFindYearMethodNominalSuccessByMerchant  = response.NewErrorResponse("failed to find yearly Nominal success methods by merchant", http.StatusInternalServerError)
	ErrFailedFindMonthMethodNominalFailedByMerchant  = response.NewErrorResponse("failed to find monthly Nominal failed methods by merchant", http.StatusInternalServerError)
	ErrFailedFindYearMethodNominalFailedByMerchant   = response.NewErrorResponse("failed to find yearly Nominal failed methods by merchant", http.StatusInternalServerError)
)
