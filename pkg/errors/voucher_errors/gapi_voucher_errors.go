package voucher_errors

import (
	"topup_game/internal/domain/response"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcVoucherNotFound  = response.NewGrpcError("error", "Voucher not found", int(codes.NotFound))
	ErrGrpcVoucherInvalidId = response.NewGrpcError("error", "Invalid Voucher ID", int(codes.NotFound))

	ErrGrpcFailedFindAll     = response.NewGrpcError("error", "Failed to fetch Vouchers", int(codes.Internal))
	ErrGrpcFailedFindActive  = response.NewGrpcError("error", "Failed to fetch active Vouchers", int(codes.Internal))
	ErrGrpcFailedFindTrashed = response.NewGrpcError("error", "Failed to fetch trashed Vouchers", int(codes.Internal))

	ErrGrpcFailedCreateVoucher   = response.NewGrpcError("error", "Failed to create Voucher", int(codes.Internal))
	ErrGrpcFailedUpdateVoucher   = response.NewGrpcError("error", "Failed to update Voucher", int(codes.Internal))
	ErrGrpcValidateCreateVoucher = response.NewGrpcError("error", "validation failed: invalid create Voucher request", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateVoucher = response.NewGrpcError("error", "validation failed: invalid update Voucher request", int(codes.InvalidArgument))

	ErrGrpcFailedTrashedVoucher  = response.NewGrpcError("error", "Failed to move Voucher to trash", int(codes.Internal))
	ErrGrpcFailedRestoreVoucher  = response.NewGrpcError("error", "Failed to restore Voucher", int(codes.Internal))
	ErrGrpcFailedDeletePermanent = response.NewGrpcError("error", "Failed to delete Voucher permanently", int(codes.Internal))

	ErrGrpcFailedRestoreAll = response.NewGrpcError("error", "Failed to restore all Vouchers", int(codes.Internal))
	ErrGrpcFailedDeleteAll  = response.NewGrpcError("error", "Failed to delete all Vouchers permanently", int(codes.Internal))

	ErrGrpcFindMonthAmountVoucherSuccess = response.NewGrpcError("error", "failed to find monthly Voucher success amounts", int(codes.Internal))
	ErrGrpcFindYearAmountVoucherSuccess  = response.NewGrpcError("error", "failed to find yearly Voucher success amounts", int(codes.Internal))
	ErrGrpcFindMonthAmountVoucherFailed  = response.NewGrpcError("error", "failed to find monthly Voucher failed amounts", int(codes.Internal))
	ErrGrpcFindYearAmountVoucherFailed   = response.NewGrpcError("error", "failed to find yearly Voucher failed amounts", int(codes.Internal))

	ErrGrpcFindMonthMethodVoucherSuccess = response.NewGrpcError("error", "failed to find monthly Voucher success methods", int(codes.Internal))
	ErrGrpcFindYearMethodVoucherSuccess  = response.NewGrpcError("error", "failed to find yearly Voucher success methods", int(codes.Internal))
	ErrGrpcFindMonthMethodVoucherFailed  = response.NewGrpcError("error", "failed to find monthly Voucher failed methods", int(codes.Internal))
	ErrGrpcFindYearMethodVoucherFailed   = response.NewGrpcError("error", "failed to find yearly Voucher failed methods", int(codes.Internal))

	ErrGrpcFindMonthAmountVoucherSuccessById = response.NewGrpcError("error", "failed to find monthly Voucher success amounts by ID", int(codes.Internal))
	ErrGrpcFindYearAmountVoucherSuccessById  = response.NewGrpcError("error", "failed to find yearly Voucher success amounts by ID", int(codes.Internal))
	ErrGrpcFindMonthAmountVoucherFailedById  = response.NewGrpcError("error", "failed to find monthly Voucher failed amounts by ID", int(codes.Internal))
	ErrGrpcFindYearAmountVoucherFailedById   = response.NewGrpcError("error", "failed to find yearly Voucher failed amounts by ID", int(codes.Internal))

	ErrGrpcFindMonthMethodVoucherSuccessById = response.NewGrpcError("error", "failed to find monthly Voucher success methods by ID", int(codes.Internal))
	ErrGrpcFindYearMethodVoucherSuccessById  = response.NewGrpcError("error", "failed to find yearly Voucher success methods by ID", int(codes.Internal))
	ErrGrpcFindMonthMethodVoucherFailedById  = response.NewGrpcError("error", "failed to find monthly Voucher failed methods by ID", int(codes.Internal))
	ErrGrpcFindYearMethodVoucherFailedById   = response.NewGrpcError("error", "failed to find yearly Voucher failed methods by ID", int(codes.Internal))

	ErrGrpcFindMonthAmountVoucherSuccessByMerchant = response.NewGrpcError("error", "failed to find monthly Voucher success amounts by merchant", int(codes.Internal))
	ErrGrpcFindYearAmountVoucherSuccessByMerchant  = response.NewGrpcError("error", "failed to find yearly Voucher success amounts by merchant", int(codes.Internal))
	ErrGrpcFindMonthAmountVoucherFailedByMerchant  = response.NewGrpcError("error", "failed to find monthly Voucher failed amounts by merchant", int(codes.Internal))
	ErrGrpcFindYearAmountVoucherFailedByMerchant   = response.NewGrpcError("error", "failed to find yearly Voucher failed amounts by merchant", int(codes.Internal))

	ErrGrpcFindMonthMethodVoucherSuccessByMerchant = response.NewGrpcError("error", "failed to find monthly Voucher success methods by merchant", int(codes.Internal))
	ErrGrpcFindYearMethodVoucherSuccessByMerchant  = response.NewGrpcError("error", "failed to find yearly Voucher success methods by merchant", int(codes.Internal))
	ErrGrpcFindMonthMethodVoucherFailedByMerchant  = response.NewGrpcError("error", "failed to find monthly Voucher failed methods by merchant", int(codes.Internal))
	ErrGrpcFindYearMethodVoucherFailedByMerchant   = response.NewGrpcError("error", "failed to find yearly Voucher failed methods by merchant", int(codes.Internal))
)
