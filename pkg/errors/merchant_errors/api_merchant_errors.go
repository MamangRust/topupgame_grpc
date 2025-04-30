package merchant_errors

import (
	"net/http"
	"topup_game/internal/domain/response"

	"github.com/labstack/echo/v4"
)

var (
	ErrApiMerchantNotFound = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "Merchant not found", http.StatusNotFound)
	}

	ErrApiMerchantInvalidId = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "invalid Merchant id", http.StatusNotFound)
	}

	ErrApiFailedFindAll = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to fetch Merchants", http.StatusInternalServerError)
	}

	ErrApiFailedFindActive = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to fetch active Merchants", http.StatusInternalServerError)
	}

	ErrApiFailedFindTrashed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to fetch trashed Merchants", http.StatusInternalServerError)
	}

	ErrApiFailedCreateMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to create Merchant", http.StatusInternalServerError)
	}

	ErrApiFailedUpdateMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to update Merchant", http.StatusInternalServerError)
	}

	ErrApiValidateCreateMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "validation failed: invalid create Merchant request", http.StatusBadRequest)
	}

	ErrApiValidateUpdateMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "validation failed: invalid update Merchant request", http.StatusBadRequest)
	}

	ErrInvalidMerchantId = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "invalid Merchant id", http.StatusBadRequest)
	}

	ErrApiBindCreateMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "bind failed: invalid create Merchant request", http.StatusBadRequest)
	}

	ErrApiBindUpdateMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "bind failed: invalid update Merchant request", http.StatusBadRequest)
	}

	ErrApiFailedTrashedMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to move Merchant to trash", http.StatusInternalServerError)
	}

	ErrApiFailedRestoreMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to restore Merchant", http.StatusInternalServerError)
	}

	ErrApiFailedDeletePermanent = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to delete Merchant permanently", http.StatusInternalServerError)
	}

	ErrApiFailedRestoreAll = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to restore all Merchants", http.StatusInternalServerError)
	}

	ErrApiFailedDeleteAll = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to delete all Merchants permanently", http.StatusInternalServerError)
	}
)
