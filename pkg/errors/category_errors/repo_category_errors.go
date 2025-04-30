package category_errors

import "errors"

var (
	ErrCategoryNotFound      = errors.New("category not found")
	ErrFindAllCategories     = errors.New("failed to find all Categories")
	ErrFindActiveCategories  = errors.New("failed to find active Categories")
	ErrFindTrashedCategories = errors.New("failed to find trashed Categories")
	ErrCategoryConflict      = errors.New("failed Category already exists")

	ErrCreateCategory = errors.New("failed to create Category")
	ErrUpdateCategory = errors.New("failed to update Category")

	ErrTrashedCategory         = errors.New("failed to move Category to trash")
	ErrRestoreCategory         = errors.New("failed to restore Category from trash")
	ErrDeleteCategoryPermanent = errors.New("failed to permanently delete Category")

	ErrRestoreAllCategories = errors.New("failed to restore all Categories")
	ErrDeleteAllCategories  = errors.New("failed to permanently delete all Categories")

	ErrFindMonthAmountCategorySuccess = errors.New("failed to find monthly Category success amounts")
	ErrFindYearAmountCategorySuccess  = errors.New("failed to find yearly Category success amounts")
	ErrFindMonthAmountCategoryFailed  = errors.New("failed to find monthly Category failed amounts")
	ErrFindYearAmountCategoryFailed   = errors.New("failed to find yearly Category failed amounts")

	ErrFindMonthMethodCategorySuccess = errors.New("failed to find monthly Category success methods")
	ErrFindYearMethodCategorySuccess  = errors.New("failed to find yearly Category success methods")
	ErrFindMonthMethodCategoryFailed  = errors.New("failed to find monthly Category failed methods")
	ErrFindYearMethodCategoryFailed   = errors.New("failed to find yearly Category failed methods")

	ErrFindMonthAmountCategorySuccessById = errors.New("failed to find monthly Category success amounts by ID")
	ErrFindYearAmountCategorySuccessById  = errors.New("failed to find yearly Category success amounts by ID")
	ErrFindMonthAmountCategoryFailedById  = errors.New("failed to find monthly Category failed amounts by ID")
	ErrFindYearAmountCategoryFailedById   = errors.New("failed to find yearly Category failed amounts by ID")

	ErrFindMonthMethodCategorySuccessById = errors.New("failed to find monthly Category success methods by ID")
	ErrFindYearMethodCategorySuccessById  = errors.New("failed to find yearly Category success methods by ID")
	ErrFindMonthMethodCategoryFailedById  = errors.New("failed to find monthly Category failed methods by ID")
	ErrFindYearMethodCategoryFailedById   = errors.New("failed to find yearly Category failed methods by ID")

	ErrFindMonthAmountCategorySuccessByMerchant = errors.New("failed to find monthly Category success amounts by merchant")
	ErrFindYearAmountCategorySuccessByMerchant  = errors.New("failed to find yearly Category success amounts by merchant")
	ErrFindMonthAmountCategoryFailedByMerchant  = errors.New("failed to find monthly Category failed amounts by merchant")
	ErrFindYearAmountCategoryFailedByMerchant   = errors.New("failed to find yearly Category failed amounts by merchant")

	ErrFindMonthMethodCategorySuccessByMerchant = errors.New("failed to find monthly Category success methods by merchant")
	ErrFindYearMethodCategorySuccessByMerchant  = errors.New("failed to find yearly Category success methods by merchant")
	ErrFindMonthMethodCategoryFailedByMerchant  = errors.New("failed to find monthly Category failed methods by merchant")
	ErrFindYearMethodCategoryFailedByMerchant   = errors.New("failed to find yearly Category failed methods by merchant")
)
