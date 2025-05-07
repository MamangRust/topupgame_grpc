package transaction_errors

import (
	"topup_game/internal/domain/response"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcTransactionInvalidYear       = response.NewGrpcError("error", "Invalid Transaction year", int(codes.InvalidArgument))
	ErrGrpcTransactionInvalidMonth      = response.NewGrpcError("error", "Invalid Transaction month", int(codes.InvalidArgument))
	ErrGrpcTransactionInvalidMerchantId = response.NewGrpcError("error", "Invalid Transaction merchant ID", int(codes.InvalidArgument))

	ErrGrpcTransactionNotFound  = response.NewGrpcError("error", "Transaction not found", int(codes.NotFound))
	ErrGrpcTransactionInvalidId = response.NewGrpcError("error", "Invalid Transaction ID", int(codes.NotFound))

	ErrGrpcFailedCreateTransaction   = response.NewGrpcError("error", "Failed to create Transaction", int(codes.Internal))
	ErrGrpcFailedUpdateTransaction   = response.NewGrpcError("error", "Failed to update Transaction", int(codes.Internal))
	ErrGrpcValidateCreateTransaction = response.NewGrpcError("error", "validation failed: invalid create Transaction request", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateTransaction = response.NewGrpcError("error", "validation failed: invalid update Transaction request", int(codes.InvalidArgument))
)
