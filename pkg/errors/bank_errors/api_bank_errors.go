package bank_errors

import (
	"net/http"
	"topup_game/internal/domain/response"

	"github.com/labstack/echo/v4"
)

var (
	ErrApiBankNotFound = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "bank not found", http.StatusNotFound)
	}

	ErrApiBankInvalidId = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "invalid bank id", http.StatusNotFound)
	}

	ErrApiFailedFindAll = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to fetch banks", http.StatusInternalServerError)
	}

	ErrApiFailedFindActive = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to fetch active banks", http.StatusInternalServerError)
	}

	ErrApiFailedFindTrashed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to fetch trashed banks", http.StatusInternalServerError)
	}

	ErrApiFailedCreateBank = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to create bank", http.StatusInternalServerError)
	}

	ErrApiFailedUpdateBank = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to update bank", http.StatusInternalServerError)
	}

	ErrApiValidateCreateBank = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "validation failed: invalid create bank request", http.StatusBadRequest)
	}

	ErrApiValidateUpdateBank = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "validation failed: invalid update bank request", http.StatusBadRequest)
	}

	ErrInvalidBankId = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "invalid bank id", http.StatusBadRequest)
	}

	ErrApiBindCreateBank = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "bind failed: invalid create bank request", http.StatusBadRequest)
	}

	ErrApiBindUpdateBank = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "bind failed: invalid update bank request", http.StatusBadRequest)
	}

	ErrApiFailedTrashedBank = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to move bank to trash", http.StatusInternalServerError)
	}

	ErrApiFailedRestoreBank = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to restore bank", http.StatusInternalServerError)
	}

	ErrApiFailedDeletePermanent = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to delete bank permanently", http.StatusInternalServerError)
	}

	ErrApiFailedRestoreAll = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to restore all banks", http.StatusInternalServerError)
	}

	ErrApiFailedDeleteAll = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to delete all banks permanently", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountBankSuccess = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly bank success amounts", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountBankSuccess = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly bank success amounts", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountBankFailed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly bank failed amounts", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountBankFailed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly bank failed amounts", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodBankSuccess = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly bank success methods", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodBankSuccess = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly bank success methods", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodBankFailed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly bank failed methods", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodBankFailed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly bank failed methods", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountBankSuccessById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly bank success amounts by ID", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountBankSuccessById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly bank success amounts by ID", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountBankFailedById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly bank failed amounts by ID", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountBankFailedById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly bank failed amounts by ID", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodBankSuccessById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly bank success methods by ID", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodBankSuccessById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly bank success methods by ID", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodBankFailedById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly bank failed methods by ID", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodBankFailedById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly bank failed methods by ID", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountBankSuccessByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly bank success amounts by merchant", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountBankSuccessByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly bank success amounts by merchant", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountBankFailedByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly bank failed amounts by merchant", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountBankFailedByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly bank failed amounts by merchant", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodBankSuccessByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly bank success methods by merchant", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodBankSuccessByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly bank success methods by merchant", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodBankFailedByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly bank failed methods by merchant", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodBankFailedByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly bank failed methods by merchant", http.StatusInternalServerError)
	}
)
