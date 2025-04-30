package nominal_errors

import (
	"topup_game/internal/domain/response"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcNominalNotFound  = response.NewGrpcError("error", "Nominal not found", int(codes.NotFound))
	ErrGrpcNominalInvalidId = response.NewGrpcError("error", "Invalid Nominal ID", int(codes.NotFound))

	ErrGrpcFailedFindAll    = response.NewGrpcError("error", "Grpc to fetch Nominals", int(codes.Internal))
	ErrGrpcFailedFindActive = response.NewGrpcError("error", "Grpc to fetch active Nominals", int(codes.Internal))
	ErrGrpcGrpcFindTrashed  = response.NewGrpcError("error", "Grpc to fetch trashed Nominals", int(codes.Internal))

	ErrGrpcFailedCreateNominal   = response.NewGrpcError("error", "Grpc to create Nominal", int(codes.Internal))
	ErrGrpcFailedUpdateNominal   = response.NewGrpcError("error", "Grpc to update Nominal", int(codes.Internal))
	ErrGrpcValidateCreateNominal = response.NewGrpcError("error", "validation Grpc: invalid create Nominal request", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateNominal = response.NewGrpcError("error", "validation Grpc: invalid update Nominal request", int(codes.InvalidArgument))

	ErrGrpcFailedTrashedNominal  = response.NewGrpcError("error", "Grpc to move Nominal to trash", int(codes.Internal))
	ErrGrpcFailedRestoreNominal  = response.NewGrpcError("error", "Grpc to restore Nominal", int(codes.Internal))
	ErrGrpcFailedDeletePermanent = response.NewGrpcError("error", "Grpc to delete Nominal permanently", int(codes.Internal))

	ErrGrpcFailedRestoreAll = response.NewGrpcError("error", "Grpc to restore all Nominals", int(codes.Internal))
	ErrGrpcFailedDeleteAll  = response.NewGrpcError("error", "Grpc to delete all Nominals permanently", int(codes.Internal))

	ErrGrpcFindMonthAmountNominalSuccess = response.NewGrpcError("error", "Grpc to find monthly Nominal success amounts", int(codes.Internal))
	ErrGrpcFindYearAmountNominalSuccess  = response.NewGrpcError("error", "Grpc to find yearly Nominal success amounts", int(codes.Internal))
	ErrGrpcFindMonthAmountNominalGrpc    = response.NewGrpcError("error", "Grpc to find monthly Nominal Grpc amounts", int(codes.Internal))
	ErrGrpcFindYearAmountNominalGrpc     = response.NewGrpcError("error", "Grpc to find yearly Nominal Grpc amounts", int(codes.Internal))

	ErrGrpcFindMonthMethodNominalSuccess = response.NewGrpcError("error", "Grpc to find monthly Nominal success methods", int(codes.Internal))
	ErrGrpcFindYearMethodNominalSuccess  = response.NewGrpcError("error", "Grpc to find yearly Nominal success methods", int(codes.Internal))
	ErrGrpcFindMonthMethodNominalGrpc    = response.NewGrpcError("error", "Grpc to find monthly Nominal Grpc methods", int(codes.Internal))
	ErrGrpcFindYearMethodNominalGrpc     = response.NewGrpcError("error", "Grpc to find yearly Nominal Grpc methods", int(codes.Internal))

	ErrGrpcFindMonthAmountNominalSuccessById = response.NewGrpcError("error", "Grpc to find monthly Nominal success amounts by ID", int(codes.Internal))
	ErrGrpcFindYearAmountNominalSuccessById  = response.NewGrpcError("error", "Grpc to find yearly Nominal success amounts by ID", int(codes.Internal))
	ErrGrpcFindMonthAmountNominalGrpcById    = response.NewGrpcError("error", "Grpc to find monthly Nominal Grpc amounts by ID", int(codes.Internal))
	ErrGrpcFindYearAmountNominalGrpcById     = response.NewGrpcError("error", "Grpc to find yearly Nominal Grpc amounts by ID", int(codes.Internal))

	ErrGrpcFindMonthMethodNominalSuccessById = response.NewGrpcError("error", "Grpc to find monthly Nominal success methods by ID", int(codes.Internal))
	ErrGrpcFindYearMethodNominalSuccessById  = response.NewGrpcError("error", "Grpc to find yearly Nominal success methods by ID", int(codes.Internal))
	ErrGrpcFindMonthMethodNominalGrpcById    = response.NewGrpcError("error", "Grpc to find monthly Nominal Grpc methods by ID", int(codes.Internal))
	ErrGrpcFindYearMethodNominalGrpcById     = response.NewGrpcError("error", "Grpc to find yearly Nominal Grpc methods by ID", int(codes.Internal))

	ErrGrpcFindMonthAmountNominalSuccessByMerchant = response.NewGrpcError("error", "Grpc to find monthly Nominal success amounts by merchant", int(codes.Internal))
	ErrGrpcFindYearAmountNominalSuccessByMerchant  = response.NewGrpcError("error", "Grpc to find yearly Nominal success amounts by merchant", int(codes.Internal))
	ErrGrpcFindMonthAmountNominalGrpcByMerchant    = response.NewGrpcError("error", "Grpc to find monthly Nominal Grpc amounts by merchant", int(codes.Internal))
	ErrGrpcFindYearAmountNominalGrpcByMerchant     = response.NewGrpcError("error", "Grpc to find yearly Nominal Grpc amounts by merchant", int(codes.Internal))

	ErrGrpcFindMonthMethodNominalSuccessByMerchant = response.NewGrpcError("error", "Grpc to find monthly Nominal success methods by merchant", int(codes.Internal))
	ErrGrpcFindYearMethodNominalSuccessByMerchant  = response.NewGrpcError("error", "Grpc to find yearly Nominal success methods by merchant", int(codes.Internal))
	ErrGrpcFindMonthMethodNominalGrpcByMerchant    = response.NewGrpcError("error", "Grpc to find monthly Nominal Grpc methods by merchant", int(codes.Internal))
	ErrGrpcFindYearMethodNominalGrpcByMerchant     = response.NewGrpcError("error", "Grpc to find yearly Nominal Grpc methods by merchant", int(codes.Internal))
)
