package category_errors

import (
	"net/http"
	"topup_game/internal/domain/response"

	"github.com/labstack/echo/v4"
)

var (
	ErrApiCategoryNotFound = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "Category not found", http.StatusNotFound)
	}

	ErrApiCategoryInvalidId = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "invalid Category id", http.StatusNotFound)
	}

	ErrApiFailedFindAll = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to fetch Categories", http.StatusInternalServerError)
	}

	ErrApiFailedFindActive = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to fetch active Categories", http.StatusInternalServerError)
	}

	ErrApiFailedFindTrashed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to fetch trashed Categories", http.StatusInternalServerError)
	}

	ErrApiFailedCreateCategory = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to create Category", http.StatusInternalServerError)
	}

	ErrApiFailedUpdateCategory = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to update Category", http.StatusInternalServerError)
	}

	ErrApiValidateCreateCategory = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "validation failed: invalid create Category request", http.StatusBadRequest)
	}

	ErrApiValidateUpdateCategory = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "validation failed: invalid update Category request", http.StatusBadRequest)
	}

	ErrInvalidCategoryId = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "invalid Category id", http.StatusBadRequest)
	}

	ErrApiBindCreateCategory = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "bind failed: invalid create Category request", http.StatusBadRequest)
	}

	ErrApiBindUpdateCategory = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "bind failed: invalid update Category request", http.StatusBadRequest)
	}

	ErrApiFailedTrashedCategory = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to move Category to trash", http.StatusInternalServerError)
	}

	ErrApiFailedRestoreCategory = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to restore Category", http.StatusInternalServerError)
	}

	ErrApiFailedDeletePermanent = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to delete Category permanently", http.StatusInternalServerError)
	}

	ErrApiFailedRestoreAll = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to restore all Categories", http.StatusInternalServerError)
	}

	ErrApiFailedDeleteAll = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to delete all Categories permanently", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountCategorySuccess = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Category success amounts", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountCategorySuccess = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Category success amounts", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountCategoryFailed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Category failed amounts", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountCategoryFailed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Category failed amounts", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodCategorySuccess = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Category success methods", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodCategorySuccess = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Category success methods", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodCategoryFailed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Category failed methods", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodCategoryFailed = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Category failed methods", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountCategorySuccessById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Category success amounts by ID", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountCategorySuccessById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Category success amounts by ID", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountCategoryFailedById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Category failed amounts by ID", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountCategoryFailedById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Category failed amounts by ID", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodCategorySuccessById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Category success methods by ID", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodCategorySuccessById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Category success methods by ID", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodCategoryFailedById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Category failed methods by ID", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodCategoryFailedById = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Category failed methods by ID", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountCategorySuccessByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Category success amounts by merchant", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountCategorySuccessByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Category success amounts by merchant", http.StatusInternalServerError)
	}

	ErrApiFindMonthAmountCategoryFailedByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Category failed amounts by merchant", http.StatusInternalServerError)
	}

	ErrApiFindYearAmountCategoryFailedByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Category failed amounts by merchant", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodCategorySuccessByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Category success methods by merchant", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodCategorySuccessByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Category success methods by merchant", http.StatusInternalServerError)
	}

	ErrApiFindMonthMethodCategoryFailedByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find monthly Category failed methods by merchant", http.StatusInternalServerError)
	}

	ErrApiFindYearMethodCategoryFailedByMerchant = func(c echo.Context) error {
		return response.NewApiErrorResponse(c, "error", "failed to find yearly Category failed methods by merchant", http.StatusInternalServerError)
	}
)
