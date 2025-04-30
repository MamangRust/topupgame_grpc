package gapi

import (
	"context"
	"math"
	"topup_game/internal/domain/requests"
	protomapper "topup_game/internal/mapper/proto"
	"topup_game/internal/pb"
	"topup_game/internal/service"
	"topup_game/pkg/errors/transaction_errors"

	"google.golang.org/protobuf/types/known/emptypb"
)

type transactionHandleGrpc struct {
	pb.UnimplementedTransactionServiceServer
	transactionService service.TransactionService
	mapping            protomapper.TransactionProtoMapper
}

func NewTransactionHandleGrpc(transactionService service.TransactionService,
	mapping protomapper.TransactionProtoMapper) *transactionHandleGrpc {
	return &transactionHandleGrpc{
		transactionService: transactionService,
		mapping:            mapping,
	}
}

func (s *transactionHandleGrpc) FindAll(ctx context.Context, req *pb.FindAllTransactionRequest) (*pb.ApiResponsePaginationTransaction, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := &requests.FindAllTransactions{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	role, totalRecords, err := s.transactionService.FindAll(reqService)

	if err != nil {
		return nil, transaction_errors.ErrGrpcFailedFindAll
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationTransaction(paginationMeta, "success", "Successfully fetched Transaction records", role)

	return so, nil
}

func (s *transactionHandleGrpc) FindMonthAmountTransactionSuccess(ctx context.Context, req *pb.MonthAmountTransactionRequest) (*pb.ApiResponseTransactionMonthAmountSuccess, error) {
	results, err := s.transactionService.FindMonthAmountTransactionSuccess(&requests.MonthAmountTransactionRequest{
		Year:  int(req.GetYear()),
		Month: int(req.GetMonth()),
	})
	if err != nil {
		return nil, transaction_errors.ErrGrpcFindMonthAmountTransactionSuccess
	}
	return s.mapping.ToProtoResponsesMonthAmountSuccess("success", "Monthly transaction success amounts", results), nil
}

func (s *transactionHandleGrpc) FindYearAmountTransactionSuccess(ctx context.Context, req *pb.YearAmountTransactionRequest) (*pb.ApiResponseTransactionYearAmountSuccess, error) {
	results, err := s.transactionService.FindYearAmountTransactionSuccess(int(req.GetYear()))
	if err != nil {
		return nil, transaction_errors.ErrGrpcFindYearAmountTransactionSuccess
	}
	return s.mapping.ToProtoResponseYearAmountSuccess("success", "Yearly transaction success amounts", results), nil
}

func (s *transactionHandleGrpc) FindMonthAmountTransactionFailed(ctx context.Context, req *pb.MonthAmountTransactionRequest) (*pb.ApiResponseTransactionMonthAmountFailed, error) {
	results, err := s.transactionService.FindMonthAmountTransactionFailed(&requests.MonthAmountTransactionRequest{
		Year:  int(req.GetYear()),
		Month: int(req.GetMonth()),
	})
	if err != nil {
		return nil, transaction_errors.ErrGrpcFindMonthAmountTransactionFailed
	}
	return s.mapping.ToProtoResponsesMonthAmountFailed("success", "Monthly transaction failed amounts", results), nil
}

func (s *transactionHandleGrpc) FindYearAmountTransactionFailed(ctx context.Context, req *pb.YearAmountTransactionRequest) (*pb.ApiResponseTransactionYearAmountFailed, error) {
	results, err := s.transactionService.FindYearAmountTransactionFailed(int(req.GetYear()))
	if err != nil {
		return nil, transaction_errors.ErrGrpcFindYearAmountTransactionFailed
	}
	return s.mapping.ToProtoResponseYearAmountFailed("success", "Yearly transaction failed amounts", results), nil
}

func (s *transactionHandleGrpc) FindMonthMethodTransactionSuccess(ctx context.Context, req *pb.YearAmountTransactionRequest) (*pb.ApiResponseTransactionMonthMethod, error) {
	results, err := s.transactionService.FindMonthMethodTransactionSuccess(int(req.GetYear()))
	if err != nil {
		return nil, transaction_errors.ErrGrpcFindMonthMethodTransactionSuccess
	}
	return s.mapping.ToProtoResponsesMonthMethod("success", "Monthly transaction success methods", results), nil
}

func (s *transactionHandleGrpc) FindYearMethodTransactionSuccess(ctx context.Context, req *pb.YearAmountTransactionRequest) (*pb.ApiResponseTransactionYearMethod, error) {
	results, err := s.transactionService.FindYearMethodTransactionSuccess(int(req.GetYear()))
	if err != nil {
		return nil, transaction_errors.ErrGrpcFindYearMethodTransactionSuccess
	}
	return s.mapping.ToProtoResponseYearMethod("success", "Yearly transaction success methods", results), nil
}

func (s *transactionHandleGrpc) FindMonthMethodTransactionFailed(ctx context.Context, req *pb.YearAmountTransactionRequest) (*pb.ApiResponseTransactionMonthMethod, error) {
	results, err := s.transactionService.FindMonthMethodTransactionFailed(int(req.GetYear()))
	if err != nil {
		return nil, transaction_errors.ErrGrpcFindMonthMethodTransactionFailed
	}
	return s.mapping.ToProtoResponsesMonthMethod("success", "Monthly transaction failed methods", results), nil
}

func (s *transactionHandleGrpc) FindYearMethodTransactionFailed(ctx context.Context, req *pb.YearAmountTransactionRequest) (*pb.ApiResponseTransactionYearMethod, error) {
	results, err := s.transactionService.FindYearMethodTransactionFailed(int(req.GetYear()))
	if err != nil {
		return nil, transaction_errors.ErrGrpcFindYearMethodTransactionFailed
	}
	return s.mapping.ToProtoResponseYearMethod("success", "Yearly transaction failed methods", results), nil
}

func (s *transactionHandleGrpc) FindMonthAmountTransactionSuccessByMerchant(ctx context.Context, req *pb.MonthAmountTransactionByMerchantRequest) (*pb.ApiResponseTransactionMonthAmountSuccess, error) {
	request := &requests.MonthAmountTransactionByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
		Month:      int(req.GetMonth()),
	}

	results, err := s.transactionService.FindMonthAmountTransactionSuccessByMerchant(request)
	if err != nil {
		return nil, transaction_errors.ErrGrpcFindMonthAmountTransactionSuccessByMerchant
	}

	return s.mapping.ToProtoResponsesMonthAmountSuccess(
		"success",
		"Successfully fetched monthly Transaction success amounts by merchant",
		results,
	), nil
}

func (s *transactionHandleGrpc) FindYearAmountTransactionSuccessByMerchant(ctx context.Context, req *pb.YearAmountTransactionByMerchantRequest) (*pb.ApiResponseTransactionYearAmountSuccess, error) {
	request := &requests.YearAmountTransactionByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.transactionService.FindYearAmountTransactionSuccessByMerchant(request)
	if err != nil {
		return nil, transaction_errors.ErrGrpcFindYearAmountTransactionSuccessByMerchant
	}

	return s.mapping.ToProtoResponseYearAmountSuccess(
		"success",
		"Successfully fetched yearly Transaction success amounts by merchant",
		results,
	), nil
}

func (s *transactionHandleGrpc) FindMonthAmountTransactionFailedByMerchant(ctx context.Context, req *pb.MonthAmountTransactionByMerchantRequest) (*pb.ApiResponseTransactionMonthAmountFailed, error) {
	request := &requests.MonthAmountTransactionByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
		Month:      int(req.GetMonth()),
	}

	results, err := s.transactionService.FindMonthAmountTransactionFailedByMerchant(request)
	if err != nil {
		return nil, transaction_errors.ErrGrpcFindMonthAmountTransactionFailedByMerchant
	}

	return s.mapping.ToProtoResponsesMonthAmountFailed(
		"success",
		"Successfully fetched monthly Transaction failed amounts by merchant",
		results,
	), nil
}

func (s *transactionHandleGrpc) FindYearAmountTransactionFailedByMerchant(ctx context.Context, req *pb.YearAmountTransactionByMerchantRequest) (*pb.ApiResponseTransactionYearAmountFailed, error) {
	request := &requests.YearAmountTransactionByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.transactionService.FindYearAmountTransactionFailedByMerchant(request)
	if err != nil {
		return nil, transaction_errors.ErrGrpcFindYearAmountTransactionFailedByMerchant
	}

	return s.mapping.ToProtoResponseYearAmountFailed(
		"success",
		"Successfully fetched yearly Transaction failed amounts by merchant",
		results,
	), nil
}

func (s *transactionHandleGrpc) FindMonthMethodTransactionSuccessByMerchant(ctx context.Context, req *pb.MonthMethodTransactionByMerchantRequest) (*pb.ApiResponseTransactionMonthMethod, error) {
	request := &requests.MonthMethodTransactionByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.transactionService.FindMonthMethodTransactionSuccessByMerchant(request)
	if err != nil {
		return nil, transaction_errors.ErrGrpcFindMonthMethodTransactionSuccessByMerchant
	}

	return s.mapping.ToProtoResponsesMonthMethod(
		"success",
		"Successfully fetched monthly Transaction success methods by merchant",
		results,
	), nil
}

func (s *transactionHandleGrpc) FindYearMethodTransactionSuccessByMerchant(ctx context.Context, req *pb.YearMethodTransactionByMerchantRequest) (*pb.ApiResponseTransactionYearMethod, error) {
	request := &requests.YearMethodTransactionByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.transactionService.FindYearMethodTransactionSuccessByMerchant(request)
	if err != nil {
		return nil, transaction_errors.ErrGrpcFindYearMethodTransactionSuccessByMerchant
	}

	return s.mapping.ToProtoResponseYearMethod(
		"success",
		"Successfully fetched yearly Transaction success methods by merchant",
		results,
	), nil
}

func (s *transactionHandleGrpc) FindMonthMethodTransactionFailedByMerchant(ctx context.Context, req *pb.MonthMethodTransactionByMerchantRequest) (*pb.ApiResponseTransactionMonthMethod, error) {
	request := &requests.MonthMethodTransactionByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.transactionService.FindMonthMethodTransactionFailedByMerchant(request)
	if err != nil {
		return nil, transaction_errors.ErrGrpcFindMonthMethodTransactionFailedByMerchant
	}

	return s.mapping.ToProtoResponsesMonthMethod(
		"success",
		"Successfully fetched monthly Transaction failed methods by merchant",
		results,
	), nil
}

func (s *transactionHandleGrpc) FindYearMethodTransactionFailedByMerchant(ctx context.Context, req *pb.YearMethodTransactionByMerchantRequest) (*pb.ApiResponseTransactionYearMethod, error) {
	request := &requests.YearMethodTransactionByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.transactionService.FindYearMethodTransactionFailedByMerchant(request)
	if err != nil {
		return nil, transaction_errors.ErrGrpcFindYearMethodTransactionFailedByMerchant
	}

	return s.mapping.ToProtoResponseYearMethod(
		"success",
		"Successfully fetched yearly Transaction failed methods by merchant",
		results,
	), nil
}

func (s *transactionHandleGrpc) FindById(ctx context.Context, req *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransaction, error) {
	id := int(req.GetId())

	if id == 0 {
		return nil, transaction_errors.ErrGrpcTransactionInvalidId
	}

	Transaction, err := s.transactionService.FindById(id)

	if err != nil {
		return nil, transaction_errors.ErrTransactionNotFound
	}

	TransactionResponse := s.mapping.ToProtoResponseTransaction("success", "Successfully fetched Transaction", Transaction)

	return TransactionResponse, nil
}

func (s *transactionHandleGrpc) FindByActive(ctx context.Context, req *pb.FindAllTransactionRequest) (*pb.ApiResponsePaginationTransactionDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := &requests.FindAllTransactions{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	Transactions, totalRecords, err := s.transactionService.FindByActive(reqService)

	if err != nil {
		return nil, transaction_errors.ErrGrpcFailedFindActive
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}
	so := s.mapping.ToProtoResponsePaginationTransactionDeleteAt(paginationMeta, "success", "Successfully fetched active Transactions", Transactions)

	return so, nil
}

func (s *transactionHandleGrpc) FindByTrashed(ctx context.Context, req *pb.FindAllTransactionRequest) (*pb.ApiResponsePaginationTransactionDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := &requests.FindAllTransactions{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	roles, totalRecords, err := s.transactionService.FindByTrashed(reqService)

	if err != nil {
		return nil, transaction_errors.ErrGrpcFailedFindTrashed
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}
	so := s.mapping.ToProtoResponsePaginationTransactionDeleteAt(paginationMeta, "success", "Successfully fetched trashed Transactions", roles)

	return so, nil
}

func (s *transactionHandleGrpc) Create(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.ApiResponseTransaction, error) {
	createReq := &requests.CreateTransactionRequest{
		UserID:        int(req.GetUserId()),
		MerchantID:    int(req.GetMerchantId()),
		VoucherID:     int(req.GetVoucherId()),
		NominalID:     int(req.GetNominalId()),
		BankID:        int(req.GetBankId()),
		PaymentMethod: req.GetPaymentMethod(),
	}

	if err := createReq.Validate(); err != nil {
		return nil, transaction_errors.ErrGrpcValidateCreateTransaction
	}

	transaction, err := s.transactionService.Create(createReq)
	if err != nil {
		return nil, transaction_errors.ErrGrpcFailedCreateTransaction
	}

	response := s.mapping.ToProtoResponseTransaction("success", "Successfully created transaction", transaction)
	return response, nil
}

func (s *transactionHandleGrpc) Update(ctx context.Context, req *pb.UpdateTransactionRequest) (*pb.ApiResponseTransaction, error) {
	id := int(req.GetId())

	if id == 0 {
		return nil, transaction_errors.ErrGrpcTransactionInvalidId
	}

	updateReq := &requests.UpdateTransactionRequest{
		ID:            id,
		UserID:        int(req.GetUserId()),
		MerchantID:    int(req.GetMerchantId()),
		VoucherID:     int(req.GetVoucherId()),
		NominalID:     int(req.GetNominalId()),
		BankID:        int(req.GetBankId()),
		PaymentMethod: req.GetPaymentMethod(),
	}

	if err := updateReq.Validate(); err != nil {
		return nil, transaction_errors.ErrGrpcValidateUpdateTransaction
	}

	transaction, err := s.transactionService.Update(updateReq)
	if err != nil {
		return nil, transaction_errors.ErrGrpcFailedUpdateTransaction
	}

	response := s.mapping.ToProtoResponseTransaction("success", "Successfully updated transaction", transaction)
	return response, nil
}

func (s *transactionHandleGrpc) Trashed(ctx context.Context, req *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransactionDeleteAt, error) {
	id := int(req.GetId())

	if id == 0 {
		return nil, transaction_errors.ErrGrpcTransactionInvalidId
	}

	Transaction, err := s.transactionService.Trashed(id)

	if err != nil {
		return nil, transaction_errors.ErrGrpcFailedTrashedTransaction
	}

	so := s.mapping.ToProtoResponseTransactionDeleteAt("success", "Successfully trashed Transaction", Transaction)

	return so, nil
}

func (s *transactionHandleGrpc) Restore(ctx context.Context, req *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransactionDeleteAt, error) {
	id := int(req.GetId())

	if id == 0 {
		return nil, transaction_errors.ErrGrpcTransactionInvalidId
	}

	role, err := s.transactionService.Restore(id)

	if err != nil {
		return nil, transaction_errors.ErrGrpcFailedRestoreTransaction
	}

	so := s.mapping.ToProtoResponseTransactionDeleteAt("success", "Successfully restored Transaction", role)

	return so, nil
}

func (s *transactionHandleGrpc) DeletePermanent(ctx context.Context, req *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransactionDelete, error) {
	id := int(req.GetId())

	if id == 0 {
		return nil, transaction_errors.ErrGrpcTransactionInvalidId
	}

	_, err := s.transactionService.DeletePermanent(id)

	if err != nil {
		return nil, transaction_errors.ErrGrpcFailedDeletePermanent
	}

	so := s.mapping.ToProtoResponseTransactionDelete("success", "Successfully deleted Transaction permanently")

	return so, nil
}

func (s *transactionHandleGrpc) RestoreAll(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseTransactionAll, error) {
	_, err := s.transactionService.RestoreAll()

	if err != nil {
		return nil, transaction_errors.ErrGrpcFailedRestoreAll
	}

	so := s.mapping.ToProtoResponseTransactionAll("success", "Successfully restored all Transactions")

	return so, nil
}

func (s *transactionHandleGrpc) DeleteAllPermanent(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseTransactionAll, error) {
	_, err := s.transactionService.DeleteAllPermanent()

	if err != nil {
		return nil, transaction_errors.ErrGrpcFailedDeletePermanent
	}

	so := s.mapping.ToProtoResponseTransactionAll("success", "Successfully deleted all Transactions")

	return so, nil
}
