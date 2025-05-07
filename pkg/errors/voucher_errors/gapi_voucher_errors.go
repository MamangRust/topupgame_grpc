package voucher_errors

import (
	"topup_game/internal/domain/response"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcVoucherInvalidYear       = response.NewGrpcError("error", "Invalid Voucher year", int(codes.InvalidArgument))
	ErrGrpcVoucherInvalidMonth      = response.NewGrpcError("error", "Invalid Voucher month", int(codes.InvalidArgument))
	ErrGrpcVoucherInvalidMerchantId = response.NewGrpcError("error", "Invalid Voucher merchant ID", int(codes.InvalidArgument))

	ErrGrpcVoucherNotFound  = response.NewGrpcError("error", "Voucher not found", int(codes.NotFound))
	ErrGrpcVoucherInvalidId = response.NewGrpcError("error", "Invalid Voucher ID", int(codes.NotFound))

	ErrGrpcValidateCreateVoucher = response.NewGrpcError("error", "validation failed: invalid create Voucher request", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateVoucher = response.NewGrpcError("error", "validation failed: invalid update Voucher request", int(codes.InvalidArgument))
)
