package merchant_errors

import (
	"net/http"
	"topup_game/internal/domain/response"
)

var (
	ErrMerchantNotFoundRes = response.NewErrorResponse("Merchant not found", http.StatusNotFound)
	ErrFailedFindAll       = response.NewErrorResponse("Failed to fetch Merchants", http.StatusInternalServerError)
	ErrFailedFindActive    = response.NewErrorResponse("Failed to fetch active Merchants", http.StatusInternalServerError)
	ErrFailedFindTrashed   = response.NewErrorResponse("Failed to fetch trashed Merchants", http.StatusInternalServerError)

	ErrFailedCreateMerchant = response.NewErrorResponse("Failed to create Merchant", http.StatusInternalServerError)
	ErrFailedUpdateMerchant = response.NewErrorResponse("Failed to update Merchant", http.StatusInternalServerError)

	ErrFailedTrashedMerchant = response.NewErrorResponse("Failed to move Merchant to trash", http.StatusInternalServerError)
	ErrFailedRestoreMerchant = response.NewErrorResponse("Failed to restore Merchant", http.StatusInternalServerError)
	ErrFailedDeletePermanent = response.NewErrorResponse("Failed to delete Merchant permanently", http.StatusInternalServerError)

	ErrFailedRestoreAll = response.NewErrorResponse("Failed to restore all Merchants", http.StatusInternalServerError)
	ErrFailedDeleteAll  = response.NewErrorResponse("Failed to delete all Merchants permanently", http.StatusInternalServerError)
)
