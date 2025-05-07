package bank_errors

import (
	"topup_game/internal/domain/response"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcInvalidMerchantId = response.NewGrpcError("error", "Invalid merchant ID", int(codes.InvalidArgument))
	ErrGrpcInvalidId         = response.NewGrpcError("error", "Invalid ID", int(codes.InvalidArgument))
	ErrGrpcInvalidYear       = response.NewGrpcError("error", "Invalid year", int(codes.InvalidArgument))
	ErrGrpcInvalidMonth      = response.NewGrpcError("error", "Invalid month", int(codes.InvalidArgument))

	ErrGrpcBankNotFound      = response.NewGrpcError("error", "Bank not found", int(codes.NotFound))
	ErrGrpcFailedFindAll     = response.NewGrpcError("error", "Failed to fetch Banks", int(codes.Internal))
	ErrGrpcFailedFindActive  = response.NewGrpcError("error", "Failed to fetch active Banks", int(codes.Internal))
	ErrGrpcFailedFindTrashed = response.NewGrpcError("error", "Failed to fetch trashed Banks", int(codes.Internal))

	ErrGrpcFailedCreateBank   = response.NewGrpcError("error", "Failed to create Bank", int(codes.Internal))
	ErrGrpcFailedUpdateBank   = response.NewGrpcError("error", "Failed to update Bank", int(codes.Internal))
	ErrGrpcValidateCreateBank = response.NewGrpcError("error", "validation failed: invalid create bank request", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateBank = response.NewGrpcError("error", "validation failed: invalid update bank request", int(codes.InvalidArgument))

	ErrGrpcFailedTrashedBank     = response.NewGrpcError("error", "Failed to move Bank to trash", int(codes.Internal))
	ErrGrpcFailedRestoreBank     = response.NewGrpcError("error", "Failed to restore Bank", int(codes.Internal))
	ErrGrpcFailedDeletePermanent = response.NewGrpcError("error", "Failed to delete Bank permanently", int(codes.Internal))

	ErrGrpcFailedRestoreAll = response.NewGrpcError("error", "Failed to restore all Banks", int(codes.Internal))
	ErrGrpcFailedDeleteAll  = response.NewGrpcError("error", "Failed to delete all Banks permanently", int(codes.Internal))

	ErrGrpcFindMonthAmountBankSuccess = response.NewGrpcError("error", "failed to find monthly bank success amounts", int(codes.Internal))
	ErrGrpcFindYearAmountBankSuccess  = response.NewGrpcError("error", "failed to FailedFind yearly bank success amounts", int(codes.Internal))
	ErrGrpcFindMonthAmountBankFailed  = response.NewGrpcError("error", "failed to FailedFind monthly bank failed amounts", int(codes.Internal))
	ErrGrpcFindYearAmountBankFailed   = response.NewGrpcError("error", "failed to FailedFind yearly bank failed amounts", int(codes.Internal))

	ErrGrpcFindMonthMethodBankSuccess = response.NewGrpcError("error", "failed to FailedFind monthly bank success methods", int(codes.Internal))
	ErrGrpcFindYearMethodBankSuccess  = response.NewGrpcError("error", "failed to FailedFind yearly bank success methods", int(codes.Internal))
	ErrGrpcFindMonthMethodBankFailed  = response.NewGrpcError("error", "failed to FailedFind monthly bank failed methods", int(codes.Internal))
	ErrGrpcFindYearMethodBankFailed   = response.NewGrpcError("error", "failed to FailedFind yearly bank failed methods", int(codes.Internal))

	ErrGrpcFindMonthAmountBankSuccessById = response.NewGrpcError("error", "failed to FailedFind monthly bank success amounts by ID", int(codes.Internal))
	ErrGrpcFindYearAmountBankSuccessById  = response.NewGrpcError("error", "failed to FailedFind yearly bank success amounts by ID", int(codes.Internal))
	ErrGrpcFindMonthAmountBankFailedById  = response.NewGrpcError("error", "failed to FailedFind monthly bank failed amounts by ID", int(codes.Internal))
	ErrGrpcFindYearAmountBankFailedById   = response.NewGrpcError("error", "failed to FailedFind yearly bank failed amounts by ID", int(codes.Internal))

	ErrGrpcFindMonthMethodBankSuccessById = response.NewGrpcError("error", "failed to FailedFind monthly bank success methods by ID", int(codes.Internal))
	ErrGrpcFindYearMethodBankSuccessById  = response.NewGrpcError("error", "failed to FailedFind yearly bank success methods by ID", int(codes.Internal))
	ErrGrpcFindMonthMethodBankFailedById  = response.NewGrpcError("error", "failed to FailedFind monthly bank failed methods by ID", int(codes.Internal))
	ErrGrpcFindYearMethodBankFailedById   = response.NewGrpcError("error", "failed to FailedFind yearly bank failed methods by ID", int(codes.Internal))

	ErrGrpcFindMonthAmountBankSuccessByMerchant = response.NewGrpcError("error", "failed to FailedFind monthly bank success amounts by merchant", int(codes.Internal))
	ErrGrpcFindYearAmountBankSuccessByMerchant  = response.NewGrpcError("error", "failed to FailedFind yearly bank success amounts by merchant", int(codes.Internal))
	ErrGrpcFindMonthAmountBankFailedByMerchant  = response.NewGrpcError("error", "failed to FailedFind monthly bank failed amounts by merchant", int(codes.Internal))
	ErrGrpcFindYearAmountBankFailedByMerchant   = response.NewGrpcError("error", "failed to FailedFind yearly bank failed amounts by merchant", int(codes.Internal))

	ErrGrpcFindMonthMethodBankSuccessByMerchant = response.NewGrpcError("error", "failed to FailedFind monthly bank success methods by merchant", int(codes.Internal))
	ErrGrpcFindYearMethodBankSuccessByMerchant  = response.NewGrpcError("error", "failed to FailedFind yearly bank success methods by merchant", int(codes.Internal))
	ErrGrpcFindMonthMethodBankFailedByMerchant  = response.NewGrpcError("error", "failed to FailedFind monthly bank failed methods by merchant", int(codes.Internal))
	ErrGrpcFindYearMethodBankFailedByMerchant   = response.NewGrpcError("error", "failed to FailedFind yearly bank failed methods by merchant", int(codes.Internal))
)
