package transaction_errors

import (
	"net/http"
	"topup_game/internal/domain/response"

	"github.com/labstack/echo/v4"
)

var (
	ErrApiTransactionNotFound = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "Transaction not found", http.StatusNotFound)
	}

	ErrApiTransactionInvalidId = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "invalid Transaction id", http.StatusNotFound)
	}

	ErrApiFailedFindAll = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to fetch Transactions", http.StatusInternalServerError)
	}

	ErrApiFailedFindActive = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to fetch active Transactions", http.StatusInternalServerError)
	}

	ErrApiFailedFindTrashed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to fetch trashed Transactions", http.StatusInternalServerError)
	}

	ErrApiFailedCreateTransaction = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to create Transaction", http.StatusInternalServerError)
	}

	ErrApiFailedUpdateTransaction = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to update Transaction", http.StatusInternalServerError)
	}

	ErrApiValidateCreateTransaction = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "validation failed: invalid create Transaction request", http.StatusBadRequest)
	}

	ErrApiValidateUpdateTransaction = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "validation failed: invalid update Transaction request", http.StatusBadRequest)
	}

	ErrApiBindCreateTransaction = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "bind failed: invalid create Transaction request", http.StatusBadRequest)
	}

	ErrApiBindUpdateTransaction = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "bind failed: invalid update Transaction request", http.StatusBadRequest)
	}

	ErrInvalidTransactionId = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "invalid Transaction id", http.StatusBadRequest)
	}

	ErrApiFailedTrashedTransaction = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to move Transaction to trash", http.StatusInternalServerError)
	}

	ErrApiFailedRestoreTransaction = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to restore Transaction", http.StatusInternalServerError)
	}

	ErrApiFailedDeletePermanent = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to delete Transaction permanently", http.StatusInternalServerError)
	}

	ErrApiFailedRestoreAll = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to restore all Transactions", http.StatusInternalServerError)
	}

	ErrApiFailedDeleteAll = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to delete all Transactions permanently", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountTransactionSuccess = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Transaction success amounts", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountTransactionSuccess = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Transaction success amounts", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountTransactionFailed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Transaction failed amounts", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountTransactionFailed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Transaction failed amounts", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodTransactionSuccess = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Transaction success methods", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodTransactionSuccess = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Transaction success methods", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodTransactionFailed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Transaction failed methods", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodTransactionFailed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Transaction failed methods", http.StatusInternalServerError)
	}
	ErrApiFindMonthAmountTransactionSuccessByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Transaction success amounts by merchant", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountTransactionSuccessByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Transaction success amounts by merchant", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountTransactionFailedByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Transaction failed amounts by merchant", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountTransactionFailedByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Transaction failed amounts by merchant", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodTransactionSuccessByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Transaction success methods by merchant", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodTransactionSuccessByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Transaction success methods by merchant", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodTransactionFailedByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Transaction failed methods by merchant", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodTransactionFailedByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Transaction failed methods by merchant", http.StatusInternalServerError)
	}
)
