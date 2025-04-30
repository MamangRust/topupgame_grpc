package category_errors

import (
	"topup_game/internal/domain/response"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcCategoryNotFound  = response.NewGrpcError("error", "Category not found", int(codes.NotFound))
	ErrGrpcCategoryInvalidId = response.NewGrpcError("error", "Invalid Category ID", int(codes.NotFound))

	ErrGrpcFailedFindAll     = response.NewGrpcError("error", "Failed to fetch Categories", int(codes.Internal))
	ErrGrpcFailedFindActive  = response.NewGrpcError("error", "Failed to fetch active Categories", int(codes.Internal))
	ErrGrpcFailedFindTrashed = response.NewGrpcError("error", "Failed to fetch trashed Categories", int(codes.Internal))

	ErrGrpcFailedCreateCategory   = response.NewGrpcError("error", "Failed to create Category", int(codes.Internal))
	ErrGrpcFailedUpdateCategory   = response.NewGrpcError("error", "Failed to update Category", int(codes.Internal))
	ErrGrpcValidateCreateCategory = response.NewGrpcError("error", "validation failed: invalid create Category request", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateCategory = response.NewGrpcError("error", "validation failed: invalid update Category request", int(codes.InvalidArgument))

	ErrGrpcFailedTrashedCategory = response.NewGrpcError("error", "Failed to move Category to trash", int(codes.Internal))
	ErrGrpcFailedRestoreCategory = response.NewGrpcError("error", "Failed to restore Category", int(codes.Internal))
	ErrGrpcFailedDeletePermanent = response.NewGrpcError("error", "Failed to delete Category permanently", int(codes.Internal))

	ErrGrpcFailedRestoreAll = response.NewGrpcError("error", "Failed to restore all Categories", int(codes.Internal))
	ErrGrpcFailedDeleteAll  = response.NewGrpcError("error", "Failed to delete all Categories permanently", int(codes.Internal))

	ErrGrpcFindMonthAmountCategorySuccess = response.NewGrpcError("error", "failed to find monthly Category success amounts", int(codes.Internal))
	ErrGrpcFindYearAmountCategorySuccess  = response.NewGrpcError("error", "failed to find yearly Category success amounts", int(codes.Internal))
	ErrGrpcFindMonthAmountCategoryFailed  = response.NewGrpcError("error", "failed to find monthly Category failed amounts", int(codes.Internal))
	ErrGrpcFindYearAmountCategoryFailed   = response.NewGrpcError("error", "failed to find yearly Category failed amounts", int(codes.Internal))

	ErrGrpcFindMonthMethodCategorySuccess = response.NewGrpcError("error", "failed to find monthly Category success methods", int(codes.Internal))
	ErrGrpcFindYearMethodCategorySuccess  = response.NewGrpcError("error", "failed to find yearly Category success methods", int(codes.Internal))
	ErrGrpcFindMonthMethodCategoryFailed  = response.NewGrpcError("error", "failed to find monthly Category failed methods", int(codes.Internal))
	ErrGrpcFindYearMethodCategoryFailed   = response.NewGrpcError("error", "failed to find yearly Category failed methods", int(codes.Internal))

	ErrGrpcFindMonthAmountCategorySuccessById = response.NewGrpcError("error", "failed to find monthly Category success amounts by ID", int(codes.Internal))
	ErrGrpcFindYearAmountCategorySuccessById  = response.NewGrpcError("error", "failed to find yearly Category success amounts by ID", int(codes.Internal))
	ErrGrpcFindMonthAmountCategoryFailedById  = response.NewGrpcError("error", "failed to find monthly Category failed amounts by ID", int(codes.Internal))
	ErrGrpcFindYearAmountCategoryFailedById   = response.NewGrpcError("error", "failed to find yearly Category failed amounts by ID", int(codes.Internal))

	ErrGrpcFindMonthMethodCategorySuccessById = response.NewGrpcError("error", "failed to find monthly Category success methods by ID", int(codes.Internal))
	ErrGrpcFindYearMethodCategorySuccessById  = response.NewGrpcError("error", "failed to find yearly Category success methods by ID", int(codes.Internal))
	ErrGrpcFindMonthMethodCategoryFailedById  = response.NewGrpcError("error", "failed to find monthly Category failed methods by ID", int(codes.Internal))
	ErrGrpcFindYearMethodCategoryFailedById   = response.NewGrpcError("error", "failed to find yearly Category failed methods by ID", int(codes.Internal))

	ErrGrpcFindMonthAmountCategorySuccessByMerchant = response.NewGrpcError("error", "failed to find monthly Category success amounts by merchant", int(codes.Internal))
	ErrGrpcFindYearAmountCategorySuccessByMerchant  = response.NewGrpcError("error", "failed to find yearly Category success amounts by merchant", int(codes.Internal))
	ErrGrpcFindMonthAmountCategoryFailedByMerchant  = response.NewGrpcError("error", "failed to find monthly Category failed amounts by merchant", int(codes.Internal))
	ErrGrpcFindYearAmountCategoryFailedByMerchant   = response.NewGrpcError("error", "failed to find yearly Category failed amounts by merchant", int(codes.Internal))

	ErrGrpcFindMonthMethodCategorySuccessByMerchant = response.NewGrpcError("error", "failed to find monthly Category success methods by merchant", int(codes.Internal))
	ErrGrpcFindYearMethodCategorySuccessByMerchant  = response.NewGrpcError("error", "failed to find yearly Category success methods by merchant", int(codes.Internal))
	ErrGrpcFindMonthMethodCategoryFailedByMerchant  = response.NewGrpcError("error", "failed to find monthly Category failed methods by merchant", int(codes.Internal))
	ErrGrpcFindYearMethodCategoryFailedByMerchant   = response.NewGrpcError("error", "failed to find yearly Category failed methods by merchant", int(codes.Internal))
)
