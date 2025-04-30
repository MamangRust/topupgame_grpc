package nominal_errors

import (
	"net/http"
	"topup_game/internal/domain/response"

	"github.com/labstack/echo/v4"
)

var (
	ErrApiNominalNotFound = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "Nominal not found", http.StatusNotFound)
	}

	ErrApiNominalInvalidId = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "invalid Nominal id", http.StatusNotFound)
	}

	ErrApiFailedFindAll = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to fetch Nominals", http.StatusInternalServerError)
	}

	ErrApiFailedFindActive = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to fetch active Nominals", http.StatusInternalServerError)
	}

	ErrApiFailedFindTrashed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to fetch trashed Nominals", http.StatusInternalServerError)
	}

	ErrApiFailedCreateNominal = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to create Nominal", http.StatusInternalServerError)
	}

	ErrApiFailedUpdateNominal = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to update Nominal", http.StatusInternalServerError)
	}

	ErrApiValidateCreateNominal = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "validation failed: invalid create Nominal request", http.StatusBadRequest)
	}

	ErrApiValidateUpdateNominal = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "validation failed: invalid update Nominal request", http.StatusBadRequest)
	}

	ErrInvalidNominalId = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "invalid Nominal id", http.StatusBadRequest)
	}

	ErrApiBindCreateNominal = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "bind failed: invalid create Nominal request", http.StatusBadRequest)
	}

	ErrApiBindUpdateNominal = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "bind failed: invalid update Nominal request", http.StatusBadRequest)
	}

	ErrApiFailedTrashedNominal = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to move Nominal to trash", http.StatusInternalServerError)
	}

	ErrApiFailedRestoreNominal = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to restore Nominal", http.StatusInternalServerError)
	}

	ErrApiFailedDeletePermanent = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to delete Nominal permanently", http.StatusInternalServerError)
	}

	ErrApiFailedRestoreAll = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to restore all Nominals", http.StatusInternalServerError)
	}

	ErrApiFailedDeleteAll = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to delete all Nominals permanently", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountNominalSuccess = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Nominal success amounts", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountNominalSuccess = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Nominal success amounts", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountNominalFailed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Nominal failed amounts", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountNominalFailed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Nominal failed amounts", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodNominalSuccess = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Nominal success methods", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodNominalSuccess = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Nominal success methods", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodNominalFailed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Nominal failed methods", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodNominalFailed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Nominal failed methods", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountNominalSuccessById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Nominal success amounts by ID", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountNominalSuccessById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Nominal success amounts by ID", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountNominalFailedById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Nominal failed amounts by ID", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountNominalFailedById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Nominal failed amounts by ID", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodNominalSuccessById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Nominal success methods by ID", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodNominalSuccessById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Nominal success methods by ID", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodNominalFailedById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Nominal failed methods by ID", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodNominalFailedById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Nominal failed methods by ID", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountNominalSuccessByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Nominal success amounts by merchant", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountNominalSuccessByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Nominal success amounts by merchant", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountNominalFailedByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Nominal failed amounts by merchant", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountNominalFailedByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Nominal failed amounts by merchant", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodNominalSuccessByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Nominal success methods by merchant", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodNominalSuccessByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Nominal success methods by merchant", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodNominalFailedByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Nominal failed methods by merchant", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodNominalFailedByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Nominal failed methods by merchant", http.StatusInternalServerError)
	}
)
