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

type nominalHandleGrpc struct {
	pb.UnimplementedNominalServiceServer
	nominalService service.NominalService
	mapping        protomapper.NominalProtoMapper
}

func NewNominalHandleGrpc(nominalService service.NominalService,
	mapping protomapper.NominalProtoMapper) *nominalHandleGrpc {
	return &nominalHandleGrpc{
		nominalService: nominalService,
		mapping:        mapping,
	}
}

func (s *nominalHandleGrpc) FindAll(ctx context.Context, req *pb.FindAllNominalRequest) (*pb.ApiResponsePaginationNominal, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	role, totalRecords, err := s.nominalService.FindAll(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch Nominal records: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationNominal(paginationMeta, "success", "Successfully fetched Nominal records", role)

	return so, nil
}

func (s *nominalHandleGrpc) FindById(ctx context.Context, req *pb.FindByIdNominalRequest) (*pb.ApiResponseNominal, error) {
	Nominal_id := int(req.GetNominalId())

	Nominal, err := s.nominalService.FindByID(Nominal_id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch Nominal: " + err.Message,
		})
	}

	NominalResponse := s.mapping.ToProtoResponseNominal("success", "Successfully fetched Nominal", Nominal)

	return NominalResponse, nil
}

func (s *nominalHandleGrpc) FindByActive(ctx context.Context, req *pb.FindAllNominalRequest) (*pb.ApiResponsePaginationNominalDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	Nominals, totalRecords, err := s.nominalService.FindByActive(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch active Nominals: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}
	so := s.mapping.ToProtoResponsePaginationNominalDeleteAt(paginationMeta, "success", "Successfully fetched active Nominals", Nominals)

	return so, nil
}

func (s *nominalHandleGrpc) FindByTrashed(ctx context.Context, req *pb.FindAllNominalRequest) (*pb.ApiResponsePaginationNominalDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	roles, totalRecords, err := s.nominalService.FindByTrashed(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch trashed Nominals: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}
	so := s.mapping.ToProtoResponsePaginationNominalDeleteAt(paginationMeta, "success", "Successfully fetched trashed Nominals", roles)

	return so, nil
}

func (s *nominalHandleGrpc) Create(ctx context.Context, req *pb.CreateNominalRequest) (*pb.ApiResponseNominal, error) {
	createReq := &requests.CreateNominalRequest{
		VoucherID: int(req.GetVoucherId()),
		Name:      req.GetName(),
		Quantity:  int(req.GetQuantity()),
		Price:     req.GetPrice(),
	}

	if err := createReq.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid create nominal request: " + err.Error(),
		})
	}

	nominal, err := s.nominalService.Create(createReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create nominal: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseNominal("success", "Successfully created Nominal", nominal)

	return so, nil
}

func (s *nominalHandleGrpc) Update(ctx context.Context, req *pb.UpdateNominalRequest) (*pb.ApiResponseNominal, error) {
	updateReq := &requests.UpdateNominalRequest{
		ID:        int(req.GetId()),
		VoucherID: int(req.GetVoucherId()),
		Name:      req.GetName(),
		Quantity:  int(req.GetQuantity()),
		Price:     req.GetPrice(),
	}

	if err := updateReq.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid update nominal request: " + err.Error(),
		})
	}

	nominal, err := s.nominalService.Update(updateReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update nominal: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseNominal("success", "Successfully updated Nominal", nominal)

	return so, nil
}

func (s *nominalHandleGrpc) Trashed(ctx context.Context, req *pb.FindByIdNominalRequest) (*pb.ApiResponseNominalDeleteAt, error) {
	nominal_id := req.NominalId

	Nominal, err := s.nominalService.Trashed(int(nominal_id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash Nominal: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseNominalDeleteAt("success", "Successfully trashed Nominal", Nominal)

	return so, nil
}

func (s *nominalHandleGrpc) Restore(ctx context.Context, req *pb.FindByIdNominalRequest) (*pb.ApiResponseNominalDeleteAt, error) {
	nominal_id := req.NominalId

	role, err := s.nominalService.Restore(int(nominal_id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore Nominal: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseNominalDeleteAt("success", "Successfully restored Nominal", role)

	return so, nil
}

func (s *nominalHandleGrpc) DeletePermanent(ctx context.Context, req *pb.FindByIdNominalRequest) (*pb.ApiResponseNominalDelete, error) {
	nominal_id := req.NominalId

	_, err := s.nominalService.DeletePermanent(int(nominal_id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete Nominal permanently: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseNominalDelete("success", "Successfully deleted Nominal permanently")

	return so, nil
}

func (s *nominalHandleGrpc) RestoreAll(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseNominalAll, error) {
	_, err := s.nominalService.RestoreAll()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all Nominals: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseNominalAll("success", "Successfully restored all Nominals")

	return so, nil
}

func (s *nominalHandleGrpc) DeleteAllPermanent(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseNominalAll, error) {
	_, err := s.nominalService.DeleteAllPermanent()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete all Nominals permanently: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseNominalAll("success", "Successfully deleted all Nominals")

	return so, nil
}
