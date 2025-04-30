package merchant_errors

import (
	"topup_game/internal/domain/response"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcMerchantNotFound  = response.NewGrpcError("error", "Merchant not found", int(codes.NotFound))
	ErrGrpcMerchantInvalidId = response.NewGrpcError("error", "Invalid Merchant ID", int(codes.NotFound))

	ErrGrpcFailedFindAll     = response.NewGrpcError("error", "Failed to fetch Merchants", int(codes.Internal))
	ErrGrpcFailedFindActive  = response.NewGrpcError("error", "Failed to fetch active Merchants", int(codes.Internal))
	ErrGrpcFailedFindTrashed = response.NewGrpcError("error", "Failed to fetch trashed Merchants", int(codes.Internal))

	ErrGrpcFailedCreateMerchant   = response.NewGrpcError("error", "Failed to create Merchant", int(codes.Internal))
	ErrGrpcFailedUpdateMerchant   = response.NewGrpcError("error", "Failed to update Merchant", int(codes.Internal))
	ErrGrpcValidateCreateMerchant = response.NewGrpcError("error", "validation failed: invalid create Merchant request", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateMerchant = response.NewGrpcError("error", "validation failed: invalid update Merchant request", int(codes.InvalidArgument))

	ErrGrpcFailedTrashedMerchant = response.NewGrpcError("error", "Failed to move Merchant to trash", int(codes.Internal))
	ErrGrpcFailedRestoreMerchant = response.NewGrpcError("error", "Failed to restore Merchant", int(codes.Internal))
	ErrGrpcFailedDeletePermanent = response.NewGrpcError("error", "Failed to delete Merchant permanently", int(codes.Internal))

	ErrGrpcFailedRestoreAll = response.NewGrpcError("error", "Failed to restore all Merchants", int(codes.Internal))
	ErrGrpcFailedDeleteAll  = response.NewGrpcError("error", "Failed to delete all Merchants permanently", int(codes.Internal))
)
