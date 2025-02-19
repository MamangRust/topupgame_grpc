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

type bankHandleGrpc struct {
	pb.UnimplementedBankServiceServer
	bankService service.BankService
	mapping     protomapper.BankProtoMapper
}

func NewBankHandleGrpc(bank service.BankService, mapping protomapper.BankProtoMapper) *bankHandleGrpc {
	return &bankHandleGrpc{
		bankService: bank,
		mapping:     mapping,
	}
}

func (s *bankHandleGrpc) FindAll(ctx context.Context, req *pb.FindAllBankRequest) (*pb.ApiResponsePaginationBank, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	role, totalRecords, err := s.bankService.FindAll(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch bank records: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationBank(paginationMeta, "success", "Successfully fetched bank records", role)

	return so, nil
}

func (s *bankHandleGrpc) FindByIdRole(ctx context.Context, req *pb.FindByIdBankRequest) (*pb.ApiResponseBank, error) {
	bank_id := int(req.GetBankId())

	bank, err := s.bankService.FindByID(bank_id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch bank: " + err.Message,
		})
	}

	bankResponse := s.mapping.ToProtoResponseBank("success", "Successfully fetched bank", bank)

	return bankResponse, nil
}

func (s *bankHandleGrpc) FindByActive(ctx context.Context, req *pb.FindAllBankRequest) (*pb.ApiResponsePaginationBankDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	banks, totalRecords, err := s.bankService.FindByActive(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch active banks: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}
	so := s.mapping.ToProtoResponsePaginationBankDeleteAt(paginationMeta, "success", "Successfully fetched active banks", banks)

	return so, nil
}

func (s *bankHandleGrpc) FindByTrashed(ctx context.Context, req *pb.FindAllBankRequest) (*pb.ApiResponsePaginationBankDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	roles, totalRecords, err := s.bankService.FindByTrashed(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch trashed banks: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}
	so := s.mapping.ToProtoResponsePaginationBankDeleteAt(paginationMeta, "success", "Successfully fetched trashed banks", roles)

	return so, nil
}

func (s *bankHandleGrpc) Create(ctx context.Context, req *pb.CreateBankRequest) (*pb.ApiResponseBank, error) {
	name := req.GetName()

	request := &requests.CreateBankRequest{
		Name: name,
	}

	if err := request.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid create bank request: " + err.Error(),
		})
	}

	bank, err := s.bankService.Create(request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create bank: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseBank("success", "Successfully created bank", bank)

	return so, nil
}

func (s *bankHandleGrpc) Update(ctx context.Context, req *pb.UpdateBankRequest) (*pb.ApiResponseBank, error) {
	roleID := int(req.GetId())
	name := req.GetName()

	request := &requests.UpdateBankRequest{
		ID:   roleID,
		Name: name,
	}

	if err := request.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid update bank request: " + err.Error(),
		})
	}

	role, err := s.bankService.Update(request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update bank: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseBank("success", "Successfully updated bank", role)

	return so, nil
}

func (s *bankHandleGrpc) Trashed(ctx context.Context, req *pb.FindByIdBankRequest) (*pb.ApiResponseBankDeleteAt, error) {
	bank_id := int(req.GetBankId())

	bank, err := s.bankService.Trashed(bank_id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash bank: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseBankDeleteAt("success", "Successfully trashed bank", bank)

	return so, nil
}

func (s *bankHandleGrpc) Restore(ctx context.Context, req *pb.FindByIdBankRequest) (*pb.ApiResponseBankDeleteAt, error) {
	bank_id := int(req.GetBankId())

	role, err := s.bankService.Restore(bank_id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore bank: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseBankDeleteAt("success", "Successfully restored bank", role)

	return so, nil
}

func (s *bankHandleGrpc) DeletePermanent(ctx context.Context, req *pb.FindByIdBankRequest) (*pb.ApiResponseBankDelete, error) {
	bank_id := int(req.GetBankId())

	_, err := s.bankService.DeletePermanent(bank_id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete bank permanently: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseBankDelete("success", "Successfully deleted bank permanently")

	return so, nil
}

func (s *bankHandleGrpc) RestoreAll(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseBankAll, error) {
	_, err := s.bankService.RestoreAll()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all banks: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseBankAll("success", "Successfully restored all banks")

	return so, nil
}

func (s *bankHandleGrpc) DeleteAllPermanent(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseBankAll, error) {
	_, err := s.bankService.DeleteAllPermanent()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete all banks permanently: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseBankAll("success", "Successfully deleted all banks")

	return so, nil
}
