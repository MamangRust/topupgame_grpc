package gapi

import (
	"context"
	"math"
	"topup_game/internal/domain/requests"
	protomapper "topup_game/internal/mapper/proto"
	"topup_game/internal/pb"
	"topup_game/internal/service"
	"topup_game/pkg/errors/voucher_errors"

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

	reqService := &requests.FindAllVouchers{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	role, totalRecords, err := s.voucherService.FindAll(reqService)

	if err != nil {
		return nil, voucher_errors.ErrGrpcFailedFindAll
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationVoucher(paginationMeta, "success", "Successfully fetched Voucher records", role)

	return so, nil
}

func (s *voucherHandleGrpc) FindMonthAmountVoucherSuccess(ctx context.Context, req *pb.MonthAmountVoucherRequest) (*pb.ApiResponseVoucherMonthAmountSuccess, error) {
	request := &requests.MonthAmountVoucherRequest{
		Year:  int(req.GetYear()),
		Month: int(req.GetMonth()),
	}

	results, err := s.voucherService.FindMonthAmountVoucherSuccess(request)
	if err != nil {
		return nil, voucher_errors.ErrGrpcFindMonthAmountVoucherSuccess
	}

	return s.mapping.ToProtoResponsesMonthAmountSuccess(
		"success",
		"Successfully fetched monthly Voucher success amounts",
		results,
	), nil
}

func (s *voucherHandleGrpc) FindYearAmountVoucherSuccess(ctx context.Context, req *pb.YearAmountVoucherRequest) (*pb.ApiResponseVoucherYearAmountSuccess, error) {
	year := int(req.GetYear())

	results, err := s.voucherService.FindYearAmountVoucherSuccess(year)
	if err != nil {
		return nil, voucher_errors.ErrGrpcFindYearAmountVoucherSuccess
	}

	return s.mapping.ToProtoResponsesYearAmountSuccess(
		"success",
		"Successfully fetched yearly Voucher success amounts",
		results,
	), nil
}

func (s *voucherHandleGrpc) FindMonthAmountVoucherFailed(ctx context.Context, req *pb.MonthAmountVoucherRequest) (*pb.ApiResponseVoucherMonthAmountFailed, error) {
	request := &requests.MonthAmountVoucherRequest{
		Year:  int(req.GetYear()),
		Month: int(req.GetMonth()),
	}

	results, err := s.voucherService.FindMonthAmountVoucherFailed(request)
	if err != nil {
		return nil, voucher_errors.ErrGrpcFindMonthAmountVoucherFailed
	}

	return s.mapping.ToProtoResponsesMonthAmountFailed(
		"success",
		"Successfully fetched monthly Voucher failed amounts",
		results,
	), nil
}

func (s *voucherHandleGrpc) FindYearAmountVoucherFailed(ctx context.Context, req *pb.YearAmountVoucherRequest) (*pb.ApiResponseVoucherYearAmountFailed, error) {
	year := int(req.GetYear())

	results, err := s.voucherService.FindYearAmountVoucherFailed(year)
	if err != nil {
		return nil, voucher_errors.ErrGrpcFindYearAmountVoucherFailed
	}

	return s.mapping.ToProtoResponsesYearAmountFailed(
		"success",
		"Successfully fetched yearly Voucher failed amounts",
		results,
	), nil
}

func (s *voucherHandleGrpc) FindMonthMethodVoucherSuccess(ctx context.Context, req *pb.YearAmountVoucherRequest) (*pb.ApiResponseVoucherMonthMethod, error) {
	year := int(req.GetYear())

	results, err := s.voucherService.FindMonthMethodVoucherSuccess(year)
	if err != nil {
		return nil, voucher_errors.ErrGrpcFindMonthMethodVoucherSuccess
	}

	return s.mapping.ToProtoResponsesMonthMethod(
		"success",
		"Successfully fetched monthly Voucher success methods",
		results,
	), nil
}

func (s *voucherHandleGrpc) FindYearMethodVoucherSuccess(ctx context.Context, req *pb.YearAmountVoucherRequest) (*pb.ApiResponseVoucherYearMethod, error) {
	year := int(req.GetYear())

	results, err := s.voucherService.FindYearMethodVoucherSuccess(year)
	if err != nil {
		return nil, voucher_errors.ErrGrpcFindYearMethodVoucherSuccess
	}

	return s.mapping.ToProtoResponsesYearMethod(
		"success",
		"Successfully fetched yearly Voucher success methods",
		results,
	), nil
}

func (s *voucherHandleGrpc) FindMonthMethodVoucherFailed(ctx context.Context, req *pb.YearAmountVoucherRequest) (*pb.ApiResponseVoucherMonthMethod, error) {
	year := int(req.GetYear())

	results, err := s.voucherService.FindMonthMethodVoucherFailed(year)
	if err != nil {
		return nil, voucher_errors.ErrGrpcFindMonthMethodVoucherFailed
	}

	return s.mapping.ToProtoResponsesMonthMethod(
		"success",
		"Successfully fetched monthly Voucher failed methods",
		results,
	), nil
}

func (s *voucherHandleGrpc) FindYearMethodVoucherFailed(ctx context.Context, req *pb.YearAmountVoucherRequest) (*pb.ApiResponseVoucherYearMethod, error) {
	year := int(req.GetYear())

	results, err := s.voucherService.FindYearMethodVoucherFailed(year)
	if err != nil {
		return nil, voucher_errors.ErrGrpcFindYearMethodVoucherFailed
	}

	return s.mapping.ToProtoResponsesYearMethod(
		"success",
		"Successfully fetched yearly Voucher failed methods",
		results,
	), nil
}

func (s *voucherHandleGrpc) FindMonthAmountVoucherSuccessById(ctx context.Context, req *pb.MonthAmountVoucherByIdRequest) (*pb.ApiResponseVoucherMonthAmountSuccess, error) {
	id := int(req.GetId())

	if id == 0 {
		return nil, voucher_errors.ErrGrpcVoucherInvalidId
	}

	result, err := s.voucherService.FindMonthAmountVoucherSuccessById(&requests.MonthAmountVoucherByIdRequest{
		ID: id,
	})

	if err != nil {
		return nil, voucher_errors.ErrGrpcFindMonthAmountVoucherSuccessById
	}

	response := s.mapping.ToProtoResponsesMonthAmountSuccess("success", "Successfully fetched monthly voucher success amounts", result)
	return response, nil
}

func (s *voucherHandleGrpc) FindYearAmountVoucherSuccessById(ctx context.Context, req *pb.YearAmountVoucherByIdRequest) (*pb.ApiResponseVoucherYearAmountSuccess, error) {
	id := int(req.GetId())

	if id == 0 {
		return nil, voucher_errors.ErrGrpcVoucherInvalidId
	}

	result, err := s.voucherService.FindYearAmountVoucherSuccessById(&requests.YearAmountVoucherByIdRequest{
		ID: id,
	})

	if err != nil {
		return nil, voucher_errors.ErrGrpcFindYearAmountVoucherSuccessById
	}

	response := s.mapping.ToProtoResponsesYearAmountSuccess("success", "Successfully fetched yearly voucher success amounts", result)
	return response, nil
}

func (s *voucherHandleGrpc) FindMonthAmountVoucherFailedById(ctx context.Context, req *pb.MonthAmountVoucherByIdRequest) (*pb.ApiResponseVoucherMonthAmountFailed, error) {
	id := int(req.GetId())

	if id == 0 {
		return nil, voucher_errors.ErrGrpcVoucherInvalidId
	}

	result, err := s.voucherService.FindMonthAmountVoucherFailedById(&requests.MonthAmountVoucherByIdRequest{
		ID: id,
	})

	if err != nil {
		return nil, voucher_errors.ErrGrpcFindMonthAmountVoucherFailedById
	}

	response := s.mapping.ToProtoResponsesMonthAmountFailed("success", "Successfully fetched monthly voucher failed amounts", result)
	return response, nil
}

func (s *voucherHandleGrpc) FindYearAmountVoucherFailedById(ctx context.Context, req *pb.YearAmountVoucherByIdRequest) (*pb.ApiResponseVoucherYearAmountFailed, error) {
	id := int(req.GetId())

	if id == 0 {
		return nil, voucher_errors.ErrGrpcVoucherInvalidId
	}

	result, err := s.voucherService.FindYearAmountVoucherFailedById(&requests.YearAmountVoucherByIdRequest{
		ID: id,
	})

	if err != nil {
		return nil, voucher_errors.ErrGrpcFindYearAmountVoucherFailedById
	}

	response := s.mapping.ToProtoResponsesYearAmountFailed("success", "Successfully fetched yearly voucher failed amounts", result)
	return response, nil
}

func (s *voucherHandleGrpc) FindMonthMethodVoucherSuccessById(ctx context.Context, req *pb.MonthMethodVoucherByIdRequest) (*pb.ApiResponseVoucherMonthMethod, error) {
	id := int(req.GetId())

	if id == 0 {
		return nil, voucher_errors.ErrGrpcVoucherInvalidId
	}

	result, err := s.voucherService.FindMonthMethodVoucherSuccessById(&requests.MonthMethodVoucherByIdRequest{
		ID: id,
	})

	if err != nil {
		return nil, voucher_errors.ErrGrpcFindMonthMethodVoucherSuccessById
	}

	response := s.mapping.ToProtoResponsesMonthMethod("success", "Successfully fetched monthly voucher success methods", result)
	return response, nil
}

func (s *voucherHandleGrpc) FindYearMethodVoucherSuccessById(ctx context.Context, req *pb.YearMethodVoucherByIdRequest) (*pb.ApiResponseVoucherYearMethod, error) {
	id := int(req.GetId())

	if id == 0 {
		return nil, voucher_errors.ErrGrpcVoucherInvalidId
	}

	result, err := s.voucherService.FindYearMethodVoucherSuccessById(&requests.YearMethodVoucherByIdRequest{
		ID: id,
	})

	if err != nil {
		return nil, voucher_errors.ErrGrpcFindYearMethodVoucherSuccessById
	}

	response := s.mapping.ToProtoResponsesYearMethod("success", "Successfully fetched yearly voucher success methods", result)
	return response, nil
}

func (s *voucherHandleGrpc) FindMonthMethodVoucherFailedById(ctx context.Context, req *pb.MonthMethodVoucherByIdRequest) (*pb.ApiResponseVoucherMonthMethod, error) {
	id := int(req.GetId())

	if id == 0 {
		return nil, voucher_errors.ErrGrpcVoucherInvalidId
	}

	result, err := s.voucherService.FindMonthMethodVoucherFailedById(&requests.MonthMethodVoucherByIdRequest{
		ID: id,
	})

	if err != nil {
		return nil, voucher_errors.ErrGrpcFindMonthMethodVoucherFailedById
	}

	response := s.mapping.ToProtoResponsesMonthMethod("success", "Successfully fetched monthly voucher failed methods", result)
	return response, nil
}

func (s *voucherHandleGrpc) FindYearMethodVoucherFailedById(ctx context.Context, req *pb.YearMethodVoucherByIdRequest) (*pb.ApiResponseVoucherYearMethod, error) {
	id := int(req.GetId())

	if id == 0 {
		return nil, voucher_errors.ErrGrpcVoucherInvalidId
	}

	result, err := s.voucherService.FindYearMethodVoucherFailedById(&requests.YearMethodVoucherByIdRequest{
		ID: id,
	})

	if err != nil {
		return nil, voucher_errors.ErrGrpcFindYearMethodVoucherFailedById
	}

	response := s.mapping.ToProtoResponsesYearMethod("success", "Successfully fetched yearly voucher failed methods", result)
	return response, nil
}

func (s *voucherHandleGrpc) FindMonthAmountVoucherSuccessByMerchant(ctx context.Context, req *pb.MonthAmountVoucherByMerchantRequest) (*pb.ApiResponseVoucherMonthAmountSuccess, error) {
	id := int(req.GetMerchantId())

	if id == 0 {
		return nil, voucher_errors.ErrGrpcVoucherInvalidId
	}

	result, err := s.voucherService.FindMonthAmountVoucherSuccessByMerchant(&requests.MonthAmountVoucherByMerchantRequest{
		MerchantID: id,
	})

	if err != nil {
		return nil, voucher_errors.ErrGrpcFindMonthAmountVoucherSuccessByMerchant
	}

	response := s.mapping.ToProtoResponsesMonthAmountSuccess("success", "Successfully fetched monthly voucher success amounts", result)
	return response, nil
}

func (s *voucherHandleGrpc) FindYearAmountVoucherSuccessByMerchant(ctx context.Context, req *pb.YearAmountVoucherByMerchantRequest) (*pb.ApiResponseVoucherYearAmountSuccess, error) {
	id := int(req.GetMerchantId())

	if id == 0 {
		return nil, voucher_errors.ErrGrpcVoucherInvalidId
	}

	result, err := s.voucherService.FindYearAmountVoucherSuccessByMerchant(&requests.YearAmountVoucherByMerchantRequest{
		MerchantID: id,
	})

	if err != nil {
		return nil, voucher_errors.ErrGrpcFindYearAmountVoucherSuccessByMerchant
	}

	response := s.mapping.ToProtoResponsesYearAmountSuccess("success", "Successfully fetched yearly voucher success amounts", result)
	return response, nil
}

func (s *voucherHandleGrpc) FindMonthAmountVoucherFailedByMerchant(ctx context.Context, req *pb.MonthAmountVoucherByMerchantRequest) (*pb.ApiResponseVoucherMonthAmountFailed, error) {
	id := int(req.GetMerchantId())

	if id == 0 {
		return nil, voucher_errors.ErrGrpcVoucherInvalidId
	}

	result, err := s.voucherService.FindMonthAmountVoucherFailedByMerchant(&requests.MonthAmountVoucherByMerchantRequest{
		MerchantID: id,
	})

	if err != nil {
		return nil, voucher_errors.ErrGrpcFindMonthAmountVoucherFailedByMerchant
	}

	response := s.mapping.ToProtoResponsesMonthAmountFailed("success", "Successfully fetched monthly voucher failed amounts", result)
	return response, nil
}

func (s *voucherHandleGrpc) FindYearAmountVoucherFailedByMerchant(ctx context.Context, req *pb.YearAmountVoucherByMerchantRequest) (*pb.ApiResponseVoucherYearAmountFailed, error) {
	id := int(req.GetMerchantId())

	if id == 0 {
		return nil, voucher_errors.ErrGrpcVoucherInvalidId
	}

	result, err := s.voucherService.FindYearAmountVoucherFailedByMerchant(&requests.YearAmountVoucherByMerchantRequest{
		MerchantID: id,
	})

	if err != nil {
		return nil, voucher_errors.ErrGrpcFindYearAmountVoucherFailedByMerchant
	}

	response := s.mapping.ToProtoResponsesYearAmountFailed("success", "Successfully fetched yearly voucher failed amounts", result)
	return response, nil
}

func (s *voucherHandleGrpc) FindMonthMethodVoucherSuccessByMerchant(ctx context.Context, req *pb.MonthMethodVoucherByMerchantRequest) (*pb.ApiResponseVoucherMonthMethod, error) {
	id := int(req.GetMerchantId())

	if id == 0 {
		return nil, voucher_errors.ErrGrpcVoucherInvalidId
	}

	result, err := s.voucherService.FindMonthMethodVoucherSuccessByMerchant(&requests.MonthMethodVoucherByMerchantRequest{
		MerchantID: id,
	})

	if err != nil {
		return nil, voucher_errors.ErrGrpcFindMonthMethodVoucherSuccessByMerchant
	}

	response := s.mapping.ToProtoResponsesMonthMethod("success", "Successfully fetched monthly voucher success methods", result)
	return response, nil
}

func (s *voucherHandleGrpc) FindYearMethodVoucherSuccessByMerchant(ctx context.Context, req *pb.YearMethodVoucherByMerchantRequest) (*pb.ApiResponseVoucherYearMethod, error) {
	id := int(req.GetMerchantId())

	if id == 0 {
		return nil, voucher_errors.ErrGrpcVoucherInvalidId
	}

	result, err := s.voucherService.FindYearMethodVoucherSuccessByMerchant(&requests.YearMethodVoucherByMerchantRequest{
		MerchantID: id,
	})

	if err != nil {
		return nil, voucher_errors.ErrGrpcFindYearMethodVoucherSuccessByMerchant
	}

	response := s.mapping.ToProtoResponsesYearMethod("success", "Successfully fetched yearly voucher success methods", result)
	return response, nil
}

func (s *voucherHandleGrpc) FindMonthMethodVoucherFailedByMerchant(ctx context.Context, req *pb.MonthMethodVoucherByMerchantRequest) (*pb.ApiResponseVoucherMonthMethod, error) {
	id := int(req.GetMerchantId())

	if id == 0 {
		return nil, voucher_errors.ErrGrpcVoucherInvalidId
	}

	result, err := s.voucherService.FindMonthMethodVoucherFailedByMerchant(&requests.MonthMethodVoucherByMerchantRequest{
		MerchantID: id,
	})

	if err != nil {
		return nil, voucher_errors.ErrGrpcFindMonthMethodVoucherFailedByMerchant
	}

	response := s.mapping.ToProtoResponsesMonthMethod("success", "Successfully fetched monthly voucher failed methods", result)
	return response, nil
}

func (s *voucherHandleGrpc) FindYearMethodVoucherFailedByMerchant(ctx context.Context, req *pb.YearMethodVoucherByMerchantRequest) (*pb.ApiResponseVoucherYearMethod, error) {
	id := int(req.GetMerchantId())

	if id == 0 {
		return nil, voucher_errors.ErrGrpcVoucherInvalidId
	}

	result, err := s.voucherService.FindYearMethodVoucherFailedByMerchant(&requests.YearMethodVoucherByMerchantRequest{
		MerchantID: id,
	})

	if err != nil {
		return nil, voucher_errors.ErrGrpcFindYearMethodVoucherFailedByMerchant
	}

	response := s.mapping.ToProtoResponsesYearMethod("success", "Successfully fetched yearly voucher failed methods", result)
	return response, nil
}

func (s *voucherHandleGrpc) FindByID(ctx context.Context, req *pb.FindByIdVoucherRequest) (*pb.ApiResponseVoucher, error) {
	id := int(req.GetVoucherId())

	if id == 0 {
		return nil, voucher_errors.ErrGrpcVoucherInvalidId
	}

	Voucher, err := s.voucherService.FindById(id)

	if err != nil {
		return nil, voucher_errors.ErrGrpcVoucherNotFound
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

	reqService := &requests.FindAllVouchers{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	Vouchers, totalRecords, err := s.voucherService.FindByActive(reqService)

	if err != nil {
		return nil, voucher_errors.ErrGrpcFailedFindActive
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
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

	reqService := &requests.FindAllVouchers{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	roles, totalRecords, err := s.voucherService.FindByTrashed(reqService)

	if err != nil {
		return nil, voucher_errors.ErrGrpcFailedFindTrashed
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
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
		return nil, voucher_errors.ErrGrpcValidateCreateVoucher
	}

	Voucher, err := s.voucherService.Create(request)

	if err != nil {
		return nil, voucher_errors.ErrGrpcFailedCreateVoucher
	}

	so := s.mapping.ToProtoResponseVoucher("success", "Successfully created Voucher", Voucher)

	return so, nil
}

func (s *voucherHandleGrpc) Update(ctx context.Context, req *pb.UpdateVoucherRequest) (*pb.ApiResponseVoucher, error) {
	id := int(req.GetId())

	if id == 0 {
		return nil, voucher_errors.ErrGrpcVoucherInvalidId
	}

	name := req.GetName()

	request := &requests.UpdateVoucherRequest{
		ID:   id,
		Name: name,
	}

	if err := request.Validate(); err != nil {
		return nil, voucher_errors.ErrGrpcValidateUpdateVoucher
	}

	role, err := s.voucherService.Update(request)

	if err != nil {
		return nil, voucher_errors.ErrGrpcFailedUpdateVoucher
	}

	so := s.mapping.ToProtoResponseVoucher("success", "Successfully updated Voucher", role)

	return so, nil
}

func (s *voucherHandleGrpc) Trashed(ctx context.Context, req *pb.FindByIdVoucherRequest) (*pb.ApiResponseVoucherDeleteAt, error) {
	id := int(req.GetVoucherId())

	if id == 0 {
		return nil, voucher_errors.ErrGrpcVoucherInvalidId
	}

	Voucher, err := s.voucherService.Trashed(id)

	if err != nil {
		return nil, voucher_errors.ErrGrpcFailedTrashedVoucher
	}

	so := s.mapping.ToProtoResponseVoucherDeleteAt("success", "Successfully trashed Voucher", Voucher)

	return so, nil
}

func (s *voucherHandleGrpc) Restore(ctx context.Context, req *pb.FindByIdVoucherRequest) (*pb.ApiResponseVoucherDeleteAt, error) {
	id := int(req.GetVoucherId())

	if id == 0 {
		return nil, voucher_errors.ErrGrpcVoucherInvalidId
	}

	role, err := s.voucherService.Restore(id)

	if err != nil {
		return nil, voucher_errors.ErrGrpcFailedRestoreVoucher
	}

	so := s.mapping.ToProtoResponseVoucherDeleteAt("success", "Successfully restored Voucher", role)

	return so, nil
}

func (s *voucherHandleGrpc) DeletePermanent(ctx context.Context, req *pb.FindByIdVoucherRequest) (*pb.ApiResponseVoucherDelete, error) {
	id := int(req.GetVoucherId())

	if id == 0 {
		return nil, voucher_errors.ErrGrpcVoucherInvalidId
	}

	_, err := s.voucherService.DeletePermanent(id)

	if err != nil {
		return nil, voucher_errors.ErrGrpcFailedDeletePermanent
	}

	so := s.mapping.ToProtoResponseVoucherDelete("success", "Successfully deleted Voucher permanently")

	return so, nil
}

func (s *voucherHandleGrpc) RestoreAll(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseVoucherAll, error) {
	_, err := s.voucherService.RestoreAll()

	if err != nil {
		return nil, voucher_errors.ErrGrpcFailedRestoreAll
	}

	so := s.mapping.ToProtoResponseVoucherAll("success", "Successfully restored all Vouchers")

	return so, nil
}

func (s *voucherHandleGrpc) DeleteAllPermanent(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseVoucherAll, error) {
	_, err := s.voucherService.DeleteAllPermanent()

	if err != nil {
		return nil, voucher_errors.ErrGrpcFailedDeleteAll
	}

	so := s.mapping.ToProtoResponseVoucherAll("success", "Successfully deleted all Vouchers")

	return so, nil
}
