package category_errors

import (
	"net/http"
	"topup_game/internal/domain/response"
)

var (
	ErrCategoryNotFoundRes = response.NewErrorResponse("Category not found", http.StatusNotFound)
	ErrFailedFindAll       = response.NewErrorResponse("Failed to fetch Categorys", http.StatusInternalServerError)
	ErrFailedFindActive    = response.NewErrorResponse("Failed to fetch active Categorys", http.StatusInternalServerError)
	ErrFailedFindTrashed   = response.NewErrorResponse("Failed to fetch trashed Categorys", http.StatusInternalServerError)

	ErrFailedCreateCategory = response.NewErrorResponse("Failed to create Category", http.StatusInternalServerError)
	ErrFailedUpdateCategory = response.NewErrorResponse("Failed to update Category", http.StatusInternalServerError)

	ErrFailedTrashedCategory = response.NewErrorResponse("Failed to move Category to trash", http.StatusInternalServerError)
	ErrFailedRestoreCategory = response.NewErrorResponse("Failed to restore Category", http.StatusInternalServerError)
	ErrFailedDeletePermanent = response.NewErrorResponse("Failed to delete Category permanently", http.StatusInternalServerError)

	ErrFailedRestoreAll = response.NewErrorResponse("Failed to restore all Categorys", http.StatusInternalServerError)
	ErrFailedDeleteAll  = response.NewErrorResponse("Failed to delete all Categorys permanently", http.StatusInternalServerError)

	ErrFailedFindMonthAmountCategorySuccess = response.NewErrorResponse("failed to find monthly Category success amounts", http.StatusInternalServerError)
	ErrFailedFindYearAmountCategorySuccess  = response.NewErrorResponse("failed to find yearly Category success amounts", http.StatusInternalServerError)
	ErrFailedFindMonthAmountCategoryFailed  = response.NewErrorResponse("failed to find monthly Category failed amounts", http.StatusInternalServerError)
	ErrFailedFindYearAmountCategoryFailed   = response.NewErrorResponse("failed to find yearly Category failed amounts", http.StatusInternalServerError)

	ErrFailedFindMonthMethodCategorySuccess = response.NewErrorResponse("failed to find monthly Category success methods", http.StatusInternalServerError)
	ErrFailedFindYearMethodCategorySuccess  = response.NewErrorResponse("failed to find yearly Category success methods", http.StatusInternalServerError)
	ErrFailedFindMonthMethodCategoryFailed  = response.NewErrorResponse("failed to find monthly Category failed methods", http.StatusInternalServerError)
	ErrFailedFindYearMethodCategoryFailed   = response.NewErrorResponse("failed to find yearly Category failed methods", http.StatusInternalServerError)

	ErrFailedFindMonthAmountCategorySuccessById = response.NewErrorResponse("failed to find monthly Category success amounts by ID", http.StatusInternalServerError)
	ErrFailedFindYearAmountCategorySuccessById  = response.NewErrorResponse("failed to find yearly Category success amounts by ID", http.StatusInternalServerError)
	ErrFailedFindMonthAmountCategoryFailedById  = response.NewErrorResponse("failed to find monthly Category failed amounts by ID", http.StatusInternalServerError)
	ErrFailedFindYearAmountCategoryFailedById   = response.NewErrorResponse("failed to find yearly Category failed amounts by ID", http.StatusInternalServerError)

	ErrFailedFindMonthMethodCategorySuccessById = response.NewErrorResponse("failed to find monthly Category success methods by ID", http.StatusInternalServerError)
	ErrFailedFindYearMethodCategorySuccessById  = response.NewErrorResponse("failed to find yearly Category success methods by ID", http.StatusInternalServerError)
	ErrFailedFindMonthMethodCategoryFailedById  = response.NewErrorResponse("failed to find monthly Category failed methods by ID", http.StatusInternalServerError)
	ErrFailedFindYearMethodCategoryFailedById   = response.NewErrorResponse("failed to find yearly Category failed methods by ID", http.StatusInternalServerError)

	ErrFailedFindMonthAmountCategorySuccessByMerchant = response.NewErrorResponse("failed to find monthly Category success amounts by merchant", http.StatusInternalServerError)
	ErrFailedFindYearAmountCategorySuccessByMerchant  = response.NewErrorResponse("failed to find yearly Category success amounts by merchant", http.StatusInternalServerError)
	ErrFailedFindMonthAmountCategoryFailedByMerchant  = response.NewErrorResponse("failed to find monthly Category failed amounts by merchant", http.StatusInternalServerError)
	ErrFailedFindYearAmountCategoryFailedByMerchant   = response.NewErrorResponse("failed to find yearly Category failed amounts by merchant", http.StatusInternalServerError)

	ErrFailedFindMonthMethodCategorySuccessByMerchant = response.NewErrorResponse("failed to find monthly Category success methods by merchant", http.StatusInternalServerError)
	ErrFailedFindYearMethodCategorySuccessByMerchant  = response.NewErrorResponse("failed to find yearly Category success methods by merchant", http.StatusInternalServerError)
	ErrFailedFindMonthMethodCategoryFailedByMerchant  = response.NewErrorResponse("failed to find monthly Category failed methods by merchant", http.StatusInternalServerError)
	ErrFailedFindYearMethodCategoryFailedByMerchant   = response.NewErrorResponse("failed to find yearly Category failed methods by merchant", http.StatusInternalServerError)
)
