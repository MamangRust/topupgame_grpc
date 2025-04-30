package gapi

import (
	"context"
	"math"
	"topup_game/internal/domain/requests"
	protomapper "topup_game/internal/mapper/proto"
	"topup_game/internal/pb"
	"topup_game/internal/service"
	"topup_game/pkg/errors/merchant_errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type merchantHandleGrpc struct {
	pb.UnimplementedMerchantServiceServer
	merchantService service.MerchantService
	mapping         protomapper.MerchantProtoMapper
}

func NewMerchantHandleGrpc(
	merchantService service.MerchantService,
	mapping protomapper.MerchantProtoMapper,
) *merchantHandleGrpc {
	return &merchantHandleGrpc{
		merchantService: merchantService,
		mapping:         mapping,
	}
}

func (s *merchantHandleGrpc) FindAll(ctx context.Context, request *pb.FindAllMerchantRequest) (*pb.ApiResponsePaginationMerchant, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllMerchants{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	merchant, totalRecords, err := s.merchantService.FindAll(&reqService)

	if err != nil {
		return nil, merchant_errors.ErrGrpcFailedFindAll
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationMerchant(paginationMeta, "success", "Successfully fetched merchant", merchant)
	return so, nil
}

func (s *merchantHandleGrpc) FindById(ctx context.Context, request *pb.FindByIdMerchantRequest) (*pb.ApiResponseMerchant, error) {
	if request.GetId() == 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidId
	}

	merchant, err := s.merchantService.FindById(int(request.GetId()))

	if err != nil {
		return nil, merchant_errors.ErrGrpcMerchantNotFound
	}

	so := s.mapping.ToProtoResponseMerchant("success", "Successfully fetched categories", merchant)

	return so, nil

}

func (s *merchantHandleGrpc) FindByActive(ctx context.Context, request *pb.FindAllMerchantRequest) (*pb.ApiResponsePaginationMerchantDeleteAt, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllMerchants{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	merchant, totalRecords, err := s.merchantService.FindByActive(&reqService)

	if err != nil {
		return nil, merchant_errors.ErrGrpcFailedFindActive
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}
	so := s.mapping.ToProtoResponsePaginationMerchantDeleteAt(paginationMeta, "success", "Successfully fetched active merchant", merchant)

	return so, nil
}

func (s *merchantHandleGrpc) FindByTrashed(ctx context.Context, request *pb.FindAllMerchantRequest) (*pb.ApiResponsePaginationMerchantDeleteAt, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllMerchants{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	users, totalRecords, err := s.merchantService.FindByTrashed(&reqService)

	if err != nil {
		return nil, merchant_errors.ErrGrpcFailedFindTrashed
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationMerchantDeleteAt(paginationMeta, "success", "Successfully fetched trashed merchant", users)

	return so, nil
}

func (s *merchantHandleGrpc) Create(ctx context.Context, request *pb.CreateMerchantRequest) (*pb.ApiResponseMerchant, error) {
	req := &requests.CreateMerchantRequest{
		UserID:       int(request.GetUserId()),
		Name:         request.GetName(),
		Description:  request.GetDescription(),
		Address:      request.GetAddress(),
		ContactEmail: request.GetContactEmail(),
		ContactPhone: request.GetContactPhone(),
		Status:       request.GetStatus(),
	}

	if err := req.Validate(); err != nil {
		return nil, merchant_errors.ErrGrpcValidateCreateMerchant
	}

	merchant, err := s.merchantService.Create(req)
	if err != nil {
		return nil, merchant_errors.ErrGrpcFailedCreateMerchant
	}

	so := s.mapping.ToProtoResponseMerchant("success", "Successfully created merchant", merchant)
	return so, nil
}

func (s *merchantHandleGrpc) Update(ctx context.Context, request *pb.UpdateMerchantRequest) (*pb.ApiResponseMerchant, error) {
	id := int(request.GetMerchantId())

	if id == 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidId
	}

	req := &requests.UpdateMerchantRequest{
		MerchantID:   id,
		UserID:       int(request.GetUserId()),
		Name:         request.GetName(),
		Description:  request.GetDescription(),
		Address:      request.GetAddress(),
		ContactEmail: request.GetContactEmail(),
		ContactPhone: request.GetContactPhone(),
		Status:       request.GetStatus(),
	}

	if err := req.Validate(); err != nil {
		return nil, merchant_errors.ErrGrpcValidateUpdateMerchant
	}

	merchant, err := s.merchantService.Update(req)
	if err != nil {
		return nil, merchant_errors.ErrGrpcFailedUpdateMerchant
	}

	so := s.mapping.ToProtoResponseMerchant("success", "Successfully updated merchant", merchant)
	return so, nil
}

func (s *merchantHandleGrpc) Trashed(ctx context.Context, request *pb.FindByIdMerchantRequest) (*pb.ApiResponseMerchantDeleteAt, error) {
	if request.GetId() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid merchant id",
		})
	}

	merchant, err := s.merchantService.Trashed(int(request.GetId()))

	if err != nil {
		return nil, merchant_errors.ErrGrpcFailedTrashedMerchant
	}

	so := s.mapping.ToProtoResponseMerchantDeleteAt("success", "Successfully trashed merchant", merchant)

	return so, nil
}

func (s *merchantHandleGrpc) Restore(ctx context.Context, request *pb.FindByIdMerchantRequest) (*pb.ApiResponseMerchantDeleteAt, error) {
	if request.GetId() == 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidId
	}

	merchant, err := s.merchantService.Restore(int(request.GetId()))

	if err != nil {
		return nil, merchant_errors.ErrGrpcFailedRestoreMerchant
	}

	so := s.mapping.ToProtoResponseMerchantDeleteAt("success", "Successfully restored merchant", merchant)

	return so, nil
}

func (s *merchantHandleGrpc) DeletePermanent(ctx context.Context, request *pb.FindByIdMerchantRequest) (*pb.ApiResponseMerchantDelete, error) {
	if request.GetId() == 0 {
		return nil, merchant_errors.ErrGrpcMerchantInvalidId
	}

	_, err := s.merchantService.DeletePermanent(int(request.GetId()))

	if err != nil {
		return nil, merchant_errors.ErrGrpcFailedDeletePermanent
	}

	so := s.mapping.ToProtoResponseMerchantDelete("success", "Successfully deleted merchant permanently")

	return so, nil
}

func (s *merchantHandleGrpc) RestoreAll(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseMerchantAll, error) {
	_, err := s.merchantService.RestoreAll()

	if err != nil {
		return nil, merchant_errors.ErrGrpcFailedRestoreAll
	}

	so := s.mapping.ToProtoResponseMerchantAll("success", "Successfully restore all merchant")

	return so, nil
}

func (s *merchantHandleGrpc) DeleteAllPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseMerchantAll, error) {
	_, err := s.merchantService.DeleteAllPermanent()

	if err != nil {
		return nil, merchant_errors.ErrGrpcFailedDeleteAll
	}

	so := s.mapping.ToProtoResponseMerchantAll("success", "Successfully delete merchant permanen")

	return so, nil
}
