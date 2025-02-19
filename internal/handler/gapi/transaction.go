package gapi

import (
	"context"
	"math"
	"topup_game/internal/domain/requests"
	protomapper "topup_game/internal/mapper/proto"
	"topup_game/internal/pb"
	"topup_game/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	role, totalRecords, err := s.transactionService.FindAll(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch Transaction records: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationTransaction(paginationMeta, "success", "Successfully fetched Transaction records", role)

	return so, nil
}

func (s *transactionHandleGrpc) FindById(ctx context.Context, req *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransaction, error) {
	Transaction_id := int(req.GetId())

	Transaction, err := s.transactionService.FindById(Transaction_id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch Transaction: " + err.Message,
		})
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

	Transactions, totalRecords, err := s.transactionService.FindByActive(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch active Transactions: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
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

	roles, totalRecords, err := s.transactionService.FindByTrashed(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch trashed Transactions: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
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
		CategoryID:    int(req.GetCategoryId()),
		BankID:        int(req.GetBankId()),
		PaymentMethod: req.GetPaymentMethod(),
	}

	if err := createReq.Validate(); err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create transaction: ",
		})
	}

	transaction, err := s.transactionService.Create(createReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create transaction: ",
		})
	}

	response := s.mapping.ToProtoResponseTransaction("success", "Successfully created transaction", transaction)
	return response, nil
}

func (s *transactionHandleGrpc) Update(ctx context.Context, req *pb.UpdateTransactionRequest) (*pb.ApiResponseTransaction, error) {
	updateReq := &requests.UpdateTransactionRequest{
		ID:            int(req.GetId()),
		UserID:        int(req.GetUserId()),
		MerchantID:    int(req.GetMerchantId()),
		VoucherID:     int(req.GetVoucherId()),
		NominalID:     int(req.GetNominalId()),
		CategoryID:    int(req.GetCategoryId()),
		BankID:        int(req.GetBankId()),
		PaymentMethod: req.GetPaymentMethod(),
		Status:        req.GetStatus(),
	}

	if err := updateReq.Validate(); err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update transaction: ",
		})
	}

	transaction, err := s.transactionService.Update(updateReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update transaction: ",
		})
	}

	response := s.mapping.ToProtoResponseTransaction("success", "Successfully updated transaction", transaction)
	return response, nil
}

func (s *transactionHandleGrpc) Trashed(ctx context.Context, req *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransactionDeleteAt, error) {
	Transaction_id := req.GetId()
	Transaction, err := s.transactionService.Trashed(int(Transaction_id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash Transaction: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseTransactionDeleteAt("success", "Successfully trashed Transaction", Transaction)

	return so, nil
}

func (s *transactionHandleGrpc) Restore(ctx context.Context, req *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransactionDeleteAt, error) {
	Transaction_id := req.GetId()

	role, err := s.transactionService.Restore(int(Transaction_id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore Transaction: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseTransactionDeleteAt("success", "Successfully restored Transaction", role)

	return so, nil
}

func (s *transactionHandleGrpc) DeletePermanent(ctx context.Context, req *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransactionDelete, error) {
	Transaction_id := req.GetId()

	_, err := s.transactionService.DeletePermanent(int(Transaction_id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete Transaction permanently: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseTransactionDelete("success", "Successfully deleted Transaction permanently")

	return so, nil
}

func (s *transactionHandleGrpc) RestoreAll(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseTransactionAll, error) {
	_, err := s.transactionService.RestoreAll()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all Transactions: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseTransactionAll("success", "Successfully restored all Transactions")

	return so, nil
}

func (s *transactionHandleGrpc) DeleteAllPermanent(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseTransactionAll, error) {
	_, err := s.transactionService.DeleteAllPermanent()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete all Transactions permanently: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseTransactionAll("success", "Successfully deleted all Transactions")

	return so, nil
}
