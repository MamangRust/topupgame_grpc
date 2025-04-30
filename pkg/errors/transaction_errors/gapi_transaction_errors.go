package transaction_errors

import (
	"topup_game/internal/domain/response"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcTransactionNotFound  = response.NewGrpcError("error", "Transaction not found", int(codes.NotFound))
	ErrGrpcTransactionInvalidId = response.NewGrpcError("error", "Invalid Transaction ID", int(codes.NotFound))

	ErrGrpcFailedFindAll     = response.NewGrpcError("error", "Failed to fetch Transactions", int(codes.Internal))
	ErrGrpcFailedFindActive  = response.NewGrpcError("error", "Failed to fetch active Transactions", int(codes.Internal))
	ErrGrpcFailedFindTrashed = response.NewGrpcError("error", "Failed to fetch trashed Transactions", int(codes.Internal))

	ErrGrpcFailedCreateTransaction   = response.NewGrpcError("error", "Failed to create Transaction", int(codes.Internal))
	ErrGrpcFailedUpdateTransaction   = response.NewGrpcError("error", "Failed to update Transaction", int(codes.Internal))
	ErrGrpcValidateCreateTransaction = response.NewGrpcError("error", "validation failed: invalid create Transaction request", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateTransaction = response.NewGrpcError("error", "validation failed: invalid update Transaction request", int(codes.InvalidArgument))

	ErrGrpcFailedTrashedTransaction = response.NewGrpcError("error", "Failed to move Transaction to trash", int(codes.Internal))
	ErrGrpcFailedRestoreTransaction = response.NewGrpcError("error", "Failed to restore Transaction", int(codes.Internal))
	ErrGrpcFailedDeletePermanent    = response.NewGrpcError("error", "Failed to delete Transaction permanently", int(codes.Internal))

	ErrGrpcFailedRestoreAll = response.NewGrpcError("error", "Failed to restore all Transactions", int(codes.Internal))
	ErrGrpcFailedDeleteAll  = response.NewGrpcError("error", "Failed to delete all Transactions permanently", int(codes.Internal))

	ErrGrpcFindMonthAmountTransactionSuccess = response.NewGrpcError("error", "failed to find monthly Transaction success amounts", int(codes.Internal))
	ErrGrpcFindYearAmountTransactionSuccess  = response.NewGrpcError("error", "failed to find yearly Transaction success amounts", int(codes.Internal))
	ErrGrpcFindMonthAmountTransactionFailed  = response.NewGrpcError("error", "failed to find monthly Transaction failed amounts", int(codes.Internal))
	ErrGrpcFindYearAmountTransactionFailed   = response.NewGrpcError("error", "failed to find yearly Transaction failed amounts", int(codes.Internal))

	ErrGrpcFindMonthMethodTransactionSuccess = response.NewGrpcError("error", "failed to find monthly Transaction success methods", int(codes.Internal))
	ErrGrpcFindYearMethodTransactionSuccess  = response.NewGrpcError("error", "failed to find yearly Transaction success methods", int(codes.Internal))
	ErrGrpcFindMonthMethodTransactionFailed  = response.NewGrpcError("error", "failed to find monthly Transaction failed methods", int(codes.Internal))
	ErrGrpcFindYearMethodTransactionFailed   = response.NewGrpcError("error", "failed to find yearly Transaction failed methods", int(codes.Internal))

	ErrGrpcFindMonthAmountTransactionSuccessByMerchant = response.NewGrpcError("error", "failed to find monthly Transaction success amounts by merchant", int(codes.Internal))
	ErrGrpcFindYearAmountTransactionSuccessByMerchant  = response.NewGrpcError("error", "failed to find yearly Transaction success amounts by merchant", int(codes.Internal))
	ErrGrpcFindMonthAmountTransactionFailedByMerchant  = response.NewGrpcError("error", "failed to find monthly Transaction failed amounts by merchant", int(codes.Internal))
	ErrGrpcFindYearAmountTransactionFailedByMerchant   = response.NewGrpcError("error", "failed to find yearly Transaction failed amounts by merchant", int(codes.Internal))

	ErrGrpcFindMonthMethodTransactionSuccessByMerchant = response.NewGrpcError("error", "failed to find monthly Transaction success methods by merchant", int(codes.Internal))
	ErrGrpcFindYearMethodTransactionSuccessByMerchant  = response.NewGrpcError("error", "failed to find yearly Transaction success methods by merchant", int(codes.Internal))
	ErrGrpcFindMonthMethodTransactionFailedByMerchant  = response.NewGrpcError("error", "failed to find monthly Transaction failed methods by merchant", int(codes.Internal))
	ErrGrpcFindYearMethodTransactionFailedByMerchant   = response.NewGrpcError("error", "failed to find yearly Transaction failed methods by merchant", int(codes.Internal))
)
