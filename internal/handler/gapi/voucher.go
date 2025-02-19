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

type voucherHandleGrpc struct {
	pb.UnimplementedVoucherServiceServer
	voucherService service.VoucherService
	mapping        protomapper.VoucherProtoMapper
}

func NewVoucherHandleGrpc(voucherService service.VoucherService,
	mapping protomapper.VoucherProtoMapper) *voucherHandleGrpc {
	return &voucherHandleGrpc{
		voucherService: voucherService,
		mapping:        mapping,
	}
}

func (s *voucherHandleGrpc) FindAll(ctx context.Context, req *pb.FindAllVoucherRequest) (*pb.ApiResponsePaginationVoucher, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	role, totalRecords, err := s.voucherService.FindAll(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch Voucher records: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationVoucher(paginationMeta, "success", "Successfully fetched Voucher records", role)

	return so, nil
}

func (s *voucherHandleGrpc) FindByIdRole(ctx context.Context, req *pb.FindByIdVoucherRequest) (*pb.ApiResponseVoucher, error) {
	Voucher_id := int(req.GetVoucherId())

	Voucher, err := s.voucherService.FindById(Voucher_id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch Voucher: " + err.Message,
		})
	}

	VoucherResponse := s.mapping.ToProtoResponseVoucher("success", "Successfully fetched Voucher", Voucher)

	return VoucherResponse, nil
}

func (s *voucherHandleGrpc) FindByActive(ctx context.Context, req *pb.FindAllVoucherRequest) (*pb.ApiResponsePaginationVoucherDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	Vouchers, totalRecords, err := s.voucherService.FindByActive(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch active Vouchers: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}
	so := s.mapping.ToProtoResponsePaginationVoucherDeleteAt(paginationMeta, "success", "Successfully fetched active Vouchers", Vouchers)

	return so, nil
}

func (s *voucherHandleGrpc) FindByTrashed(ctx context.Context, req *pb.FindAllVoucherRequest) (*pb.ApiResponsePaginationVoucherDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	roles, totalRecords, err := s.voucherService.FindByTrashed(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch trashed Vouchers: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}
	so := s.mapping.ToProtoResponsePaginationVoucherDeleteAt(paginationMeta, "success", "Successfully fetched trashed Vouchers", roles)

	return so, nil
}

func (s *voucherHandleGrpc) Create(ctx context.Context, req *pb.CreateVoucherRequest) (*pb.ApiResponseVoucher, error) {
	name := req.GetName()

	request := &requests.CreateVoucherRequest{
		Name: name,
	}

	if err := request.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid create Voucher request: " + err.Error(),
		})
	}

	Voucher, err := s.voucherService.Create(request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create Voucher: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseVoucher("success", "Successfully created Voucher", Voucher)

	return so, nil
}

func (s *voucherHandleGrpc) Update(ctx context.Context, req *pb.UpdateVoucherRequest) (*pb.ApiResponseVoucher, error) {
	Voucher_id := int(req.GetId())
	name := req.GetName()

	request := &requests.UpdateVoucherRequest{
		ID:   Voucher_id,
		Name: name,
	}

	if err := request.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid update Voucher request: " + err.Error(),
		})
	}

	role, err := s.voucherService.Update(request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update Voucher: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseVoucher("success", "Successfully updated Voucher", role)

	return so, nil
}

func (s *voucherHandleGrpc) Trashed(ctx context.Context, req *pb.FindByIdVoucherRequest) (*pb.ApiResponseVoucherDeleteAt, error) {
	Voucher_id := req.GetVoucherId()

	Voucher, err := s.voucherService.Trashed(int(Voucher_id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash Voucher: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseVoucherDeleteAt("success", "Successfully trashed Voucher", Voucher)

	return so, nil
}

func (s *voucherHandleGrpc) Restore(ctx context.Context, req *pb.FindByIdVoucherRequest) (*pb.ApiResponseVoucherDeleteAt, error) {
	Voucher_id := req.GetVoucherId()

	role, err := s.voucherService.Restore(int(Voucher_id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore Voucher: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseVoucherDeleteAt("success", "Successfully restored Voucher", role)

	return so, nil
}

func (s *voucherHandleGrpc) DeletePermanent(ctx context.Context, req *pb.FindByIdVoucherRequest) (*pb.ApiResponseVoucherDelete, error) {
	Voucher_id := req.GetVoucherId()

	_, err := s.voucherService.DeletePermanent(int(Voucher_id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete Voucher permanently: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseVoucherDelete("success", "Successfully deleted Voucher permanently")

	return so, nil
}

func (s *voucherHandleGrpc) RestoreAll(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseVoucherAll, error) {
	_, err := s.voucherService.RestoreAll()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all Vouchers: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseVoucherAll("success", "Successfully restored all Vouchers")

	return so, nil
}

func (s *voucherHandleGrpc) DeleteAllPermanent(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseVoucherAll, error) {
	_, err := s.voucherService.DeleteAllPermanent()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete all Vouchers permanently: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseVoucherAll("success", "Successfully deleted all Vouchers")

	return so, nil
}
