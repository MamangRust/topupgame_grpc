package voucher_errors

import (
	"net/http"
	"topup_game/internal/domain/response"

	"github.com/labstack/echo/v4"
)

var (
	ErrApiVoucherNotFound = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "Voucher not found", http.StatusNotFound)
	}

	ErrApiVoucherInvalidId = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "invalid Voucher id", http.StatusNotFound)
	}

	ErrApiFailedFindAll = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to fetch Vouchers", http.StatusInternalServerError)
	}

	ErrApiFailedFindActive = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to fetch active Vouchers", http.StatusInternalServerError)
	}

	ErrApiFailedFindTrashed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to fetch trashed Vouchers", http.StatusInternalServerError)
	}

	ErrApiFailedCreateVoucher = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to create Voucher", http.StatusInternalServerError)
	}

	ErrApiFailedUpdateVoucher = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to update Voucher", http.StatusInternalServerError)
	}

	ErrApiValidateCreateVoucher = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "validation failed: invalid create Voucher request", http.StatusBadRequest)
	}

	ErrApiValidateUpdateVoucher = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "validation failed: invalid update Voucher request", http.StatusBadRequest)
	}

	ErrInvalidVoucherId = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "invalid Voucher id", http.StatusBadRequest)
	}

	ErrApiBindCreateVoucher = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "bind failed: invalid create Voucher request", http.StatusBadRequest)
	}

	ErrApiBindUpdateVoucher = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "bind failed: invalid update Voucher request", http.StatusBadRequest)
	}

	ErrApiFailedTrashedVoucher = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to move Voucher to trash", http.StatusInternalServerError)
	}

	ErrApiFailedRestoreVoucher = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to restore Voucher", http.StatusInternalServerError)
	}

	ErrApiFailedDeletePermanent = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to delete Voucher permanently", http.StatusInternalServerError)
	}

	ErrApiFailedRestoreAll = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to restore all Vouchers", http.StatusInternalServerError)
	}

	ErrApiFailedDeleteAll = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to delete all Vouchers permanently", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountVoucherSuccess = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Voucher success amounts", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountVoucherSuccess = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Voucher success amounts", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountVoucherFailed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Voucher failed amounts", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountVoucherFailed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Voucher failed amounts", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodVoucherSuccess = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Voucher success methods", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodVoucherSuccess = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Voucher success methods", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodVoucherFailed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Voucher failed methods", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodVoucherFailed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Voucher failed methods", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountVoucherSuccessById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Voucher success amounts by ID", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountVoucherSuccessById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Voucher success amounts by ID", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountVoucherFailedById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Voucher failed amounts by ID", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountVoucherFailedById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Voucher failed amounts by ID", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodVoucherSuccessById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Voucher success methods by ID", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodVoucherSuccessById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Voucher success methods by ID", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodVoucherFailedById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Voucher failed methods by ID", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodVoucherFailedById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Voucher failed methods by ID", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountVoucherSuccessByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Voucher success amounts by merchant", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountVoucherSuccessByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Voucher success amounts by merchant", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountVoucherFailedByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Voucher failed amounts by merchant", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountVoucherFailedByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Voucher failed amounts by merchant", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodVoucherSuccessByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Voucher success methods by merchant", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodVoucherSuccessByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Voucher success methods by merchant", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodVoucherFailedByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Voucher failed methods by merchant", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodVoucherFailedByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Voucher failed methods by merchant", http.StatusInternalServerError)
	}
)
