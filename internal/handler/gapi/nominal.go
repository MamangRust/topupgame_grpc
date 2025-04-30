package gapi

import (
	"context"
	"math"
	"topup_game/internal/domain/requests"
	protomapper "topup_game/internal/mapper/proto"
	"topup_game/internal/pb"
	"topup_game/internal/service"
	"topup_game/pkg/errors/nominal_errors"

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

	reqService := requests.FindAllNominals{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	role, totalRecords, err := s.nominalService.FindAll(&reqService)

	if err != nil {
		return nil, nominal_errors.ErrGrpcFailedFindAll
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationNominal(paginationMeta, "success", "Successfully fetched Nominal records", role)

	return so, nil
}

func (s *nominalHandleGrpc) FindMonthAmountNominalSuccess(ctx context.Context, req *pb.MonthAmountNominalRequest) (*pb.ApiResponseNominalMonthAmountSuccess, error) {
	request := &requests.MonthAmountNominalRequest{
		Year:  int(req.GetYear()),
		Month: int(req.GetMonth()),
	}

	results, err := s.nominalService.FindMonthAmountNominalSuccess(request)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindMonthAmountNominalSuccess
	}

	response := s.mapping.ToProtoResponsesMonthAmountSuccess(
		"success",
		"Successfully fetched monthly Nominal success amounts",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindYearAmountNominalSuccess(ctx context.Context, req *pb.YearAmountNominalRequest) (*pb.ApiResponseNominalYearAmountSuccess, error) {
	year := int(req.GetYear())

	results, err := s.nominalService.FindYearAmountNominalSuccess(year)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindYearAmountNominalSuccess
	}

	response := s.mapping.ToProtoResponseYearAmountSuccess(
		"success",
		"Successfully fetched yearly Nominal success amounts",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindMonthAmountNominalFailed(ctx context.Context, req *pb.MonthAmountNominalRequest) (*pb.ApiResponseNominalMonthAmountFailed, error) {
	request := &requests.MonthAmountNominalRequest{
		Year:  int(req.GetYear()),
		Month: int(req.GetMonth()),
	}

	results, err := s.nominalService.FindMonthAmountNominalFailed(request)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindMonthAmountNominalGrpc
	}

	response := s.mapping.ToProtoResponsesMonthAmountFailed(
		"success",
		"Successfully fetched monthly Nominal failed amounts",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindYearAmountNominalFailed(ctx context.Context, req *pb.YearAmountNominalRequest) (*pb.ApiResponseNominalYearAmountFailed, error) {
	year := int(req.GetYear())

	results, err := s.nominalService.FindYearAmountNominalFailed(year)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindYearAmountNominalGrpc
	}

	response := s.mapping.ToProtoResponseYearAmountFailed(
		"success",
		"Successfully fetched yearly Nominal failed amounts",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindMonthMethodNominalSuccess(ctx context.Context, req *pb.YearAmountNominalRequest) (*pb.ApiResponseNominalMonthMethod, error) {
	year := int(req.GetYear())

	results, err := s.nominalService.FindMonthMethodNominalSuccess(year)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindMonthMethodNominalSuccess
	}

	response := s.mapping.ToProtoResponsesMonthMethod(
		"success",
		"Successfully fetched monthly Nominal success methods",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindYearMethodNominalSuccess(ctx context.Context, req *pb.YearAmountNominalRequest) (*pb.ApiResponseNominalYearMethod, error) {
	year := int(req.GetYear())

	results, err := s.nominalService.FindYearMethodNominalSuccess(year)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindYearMethodNominalSuccess
	}

	response := s.mapping.ToProtoResponseYearMethod(
		"success",
		"Successfully fetched yearly Nominal success methods",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindMonthMethodNominalFailed(ctx context.Context, req *pb.YearAmountNominalRequest) (*pb.ApiResponseNominalMonthMethod, error) {
	year := int(req.GetYear())

	results, err := s.nominalService.FindMonthMethodNominalFailed(year)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindMonthMethodNominalGrpc
	}

	response := s.mapping.ToProtoResponsesMonthMethod(
		"success",
		"Successfully fetched monthly Nominal failed methods",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindYearMethodNominalFailed(ctx context.Context, req *pb.YearAmountNominalRequest) (*pb.ApiResponseNominalYearMethod, error) {
	year := int(req.GetYear())

	results, err := s.nominalService.FindYearMethodNominalFailed(year)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindYearMethodNominalGrpc
	}

	response := s.mapping.ToProtoResponseYearMethod(
		"success",
		"Successfully fetched yearly Nominal failed methods",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindMonthAmountNominalSuccessById(ctx context.Context, req *pb.MonthAmountNominalByIdRequest) (*pb.ApiResponseNominalMonthAmountSuccess, error) {
	request := &requests.MonthAmountNominalByIdRequest{
		ID:    int(req.GetId()),
		Year:  int(req.GetYear()),
		Month: int(req.GetMonth()),
	}

	results, err := s.nominalService.FindMonthAmountNominalSuccessById(request)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindMonthAmountNominalSuccessById
	}

	response := s.mapping.ToProtoResponsesMonthAmountSuccess(
		"success",
		"Successfully fetched monthly Nominal success amounts by ID",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindYearAmountNominalSuccessById(ctx context.Context, req *pb.YearAmountNominalByIdRequest) (*pb.ApiResponseNominalYearAmountSuccess, error) {
	request := &requests.YearAmountNominalByIdRequest{
		ID:   int(req.GetId()),
		Year: int(req.GetYear()),
	}

	results, err := s.nominalService.FindYearAmountNominalSuccessById(request)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindYearAmountNominalSuccessById
	}

	response := s.mapping.ToProtoResponseYearAmountSuccess(
		"success",
		"Successfully fetched yearly Nominal success amounts by ID",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindMonthAmountNominalFailedById(ctx context.Context, req *pb.MonthAmountNominalByIdRequest) (*pb.ApiResponseNominalMonthAmountFailed, error) {
	request := &requests.MonthAmountNominalByIdRequest{
		ID:    int(req.GetId()),
		Year:  int(req.GetYear()),
		Month: int(req.GetMonth()),
	}

	results, err := s.nominalService.FindMonthAmountNominalFailedById(request)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindMonthAmountNominalGrpcById
	}

	response := s.mapping.ToProtoResponsesMonthAmountFailed(
		"success",
		"Successfully fetched monthly Nominal failed amounts by ID",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindYearAmountNominalFailedById(ctx context.Context, req *pb.YearAmountNominalByIdRequest) (*pb.ApiResponseNominalYearAmountFailed, error) {
	request := &requests.YearAmountNominalByIdRequest{
		ID:   int(req.GetId()),
		Year: int(req.GetYear()),
	}

	results, err := s.nominalService.FindYearAmountNominalFailedById(request)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindYearAmountNominalGrpcById
	}

	response := s.mapping.ToProtoResponseYearAmountFailed(
		"success",
		"Successfully fetched yearly Nominal failed amounts by ID",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindMonthMethodNominalSuccessById(ctx context.Context, req *pb.MonthMethodNominalByIdRequest) (*pb.ApiResponseNominalMonthMethod, error) {
	request := &requests.MonthMethodNominalByIdRequest{
		ID:   int(req.GetId()),
		Year: int(req.GetYear()),
	}

	results, err := s.nominalService.FindMonthMethodNominalSuccessById(request)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindMonthMethodNominalSuccessById
	}

	response := s.mapping.ToProtoResponsesMonthMethod(
		"success",
		"Successfully fetched monthly Nominal success methods by ID",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindYearMethodNominalSuccessById(ctx context.Context, req *pb.YearMethodNominalByIdRequest) (*pb.ApiResponseNominalYearMethod, error) {
	request := &requests.YearMethodNominalByIdRequest{
		ID:   int(req.GetId()),
		Year: int(req.GetYear()),
	}

	results, err := s.nominalService.FindYearMethodNominalSuccessById(request)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindYearMethodNominalSuccessById
	}

	response := s.mapping.ToProtoResponseYearMethod(
		"success",
		"Successfully fetched yearly Nominal success methods by ID",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindMonthMethodNominalFailedById(ctx context.Context, req *pb.MonthMethodNominalByIdRequest) (*pb.ApiResponseNominalMonthMethod, error) {
	request := &requests.MonthMethodNominalByIdRequest{
		ID:   int(req.GetId()),
		Year: int(req.GetYear()),
	}

	results, err := s.nominalService.FindMonthMethodNominalFailedById(request)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindMonthMethodNominalGrpcById
	}

	response := s.mapping.ToProtoResponsesMonthMethod(
		"success",
		"Successfully fetched monthly Nominal failed methods by ID",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindYearMethodNominalFailedById(ctx context.Context, req *pb.YearMethodNominalByIdRequest) (*pb.ApiResponseNominalYearMethod, error) {
	request := &requests.YearMethodNominalByIdRequest{
		ID:   int(req.GetId()),
		Year: int(req.GetYear()),
	}

	results, err := s.nominalService.FindYearMethodNominalFailedById(request)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindYearMethodNominalGrpcById
	}

	response := s.mapping.ToProtoResponseYearMethod(
		"success",
		"Successfully fetched yearly Nominal failed methods by ID",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindMonthAmountNominalSuccessByMerchant(ctx context.Context, req *pb.MonthAmountNominalByMerchantRequest) (*pb.ApiResponseNominalMonthAmountSuccess, error) {
	request := &requests.MonthAmountNominalByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
		Month:      int(req.GetMonth()),
	}

	results, err := s.nominalService.FindMonthAmountNominalSuccessByMerchant(request)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindMonthAmountNominalSuccessByMerchant
	}

	response := s.mapping.ToProtoResponsesMonthAmountSuccess(
		"success",
		"Successfully fetched monthly Nominal success amounts by ID",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindYearAmountNominalSuccessByMerchant(ctx context.Context, req *pb.YearAmountNominalByMerchantRequest) (*pb.ApiResponseNominalYearAmountSuccess, error) {
	request := &requests.YearAmountNominalByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.nominalService.FindYearAmountNominalSuccessByMerchant(request)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindYearAmountNominalSuccessByMerchant
	}

	response := s.mapping.ToProtoResponseYearAmountSuccess(
		"success",
		"Successfully fetched yearly Nominal success amounts by ID",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindMonthAmountNominalFailedByMerchant(ctx context.Context, req *pb.MonthAmountNominalByMerchantRequest) (*pb.ApiResponseNominalMonthAmountFailed, error) {
	request := &requests.MonthAmountNominalByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
		Month:      int(req.GetMonth()),
	}

	results, err := s.nominalService.FindMonthAmountNominalFailedByMerchant(request)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindMonthAmountNominalGrpcByMerchant
	}

	response := s.mapping.ToProtoResponsesMonthAmountFailed(
		"success",
		"Successfully fetched monthly Nominal failed amounts by ID",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindYearAmountNominalFailedByMerchant(ctx context.Context, req *pb.YearAmountNominalByMerchantRequest) (*pb.ApiResponseNominalYearAmountFailed, error) {
	request := &requests.YearAmountNominalByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.nominalService.FindYearAmountNominalFailedByMerchant(request)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindYearAmountNominalGrpcByMerchant
	}

	response := s.mapping.ToProtoResponseYearAmountFailed(
		"success",
		"Successfully fetched yearly Nominal failed amounts by ID",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindMonthMethodNominalSuccessByMerchant(ctx context.Context, req *pb.MonthMethodNominalByMerchantRequest) (*pb.ApiResponseNominalMonthMethod, error) {
	request := &requests.MonthMethodNominalByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.nominalService.FindMonthMethodNominalSuccessByMerchant(request)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindMonthMethodNominalSuccessByMerchant
	}

	response := s.mapping.ToProtoResponsesMonthMethod(
		"success",
		"Successfully fetched monthly Nominal success methods by ID",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindYearMethodNominalSuccessByMerchant(ctx context.Context, req *pb.YearMethodNominalByMerchantRequest) (*pb.ApiResponseNominalYearMethod, error) {
	request := &requests.YearMethodNominalByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.nominalService.FindYearMethodNominalSuccessByMerchant(request)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindYearMethodNominalSuccessByMerchant
	}

	response := s.mapping.ToProtoResponseYearMethod(
		"success",
		"Successfully fetched yearly Nominal success methods by ID",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindMonthMethodNominalFailedByMerchant(ctx context.Context, req *pb.MonthMethodNominalByMerchantRequest) (*pb.ApiResponseNominalMonthMethod, error) {
	request := &requests.MonthMethodNominalByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.nominalService.FindMonthMethodNominalFailedByMerchant(request)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindMonthMethodNominalGrpcByMerchant
	}

	response := s.mapping.ToProtoResponsesMonthMethod(
		"success",
		"Successfully fetched monthly Nominal failed methods by ID",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindYearMethodNominalFailedByMerchant(ctx context.Context, req *pb.YearMethodNominalByMerchantRequest) (*pb.ApiResponseNominalYearMethod, error) {
	request := &requests.YearMethodNominalByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.nominalService.FindYearMethodNominalFailedByMerchant(request)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFindYearMethodNominalGrpcByMerchant
	}

	response := s.mapping.ToProtoResponseYearMethod(
		"success",
		"Successfully fetched yearly Nominal failed methods by ID",
		results,
	)
	return response, nil
}

func (s *nominalHandleGrpc) FindById(ctx context.Context, req *pb.FindByIdNominalRequest) (*pb.ApiResponseNominal, error) {
	id := int(req.GetNominalId())

	if id == 0 {
		return nil, nominal_errors.ErrGrpcNominalNotFound
	}

	Nominal, err := s.nominalService.FindByID(id)

	if err != nil {
		return nil, nominal_errors.ErrGrpcNominalNotFound
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

	reqService := requests.FindAllNominals{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	nominals, totalRecords, err := s.nominalService.FindByActive(&reqService)

	if err != nil {
		return nil, nominal_errors.ErrGrpcFailedFindActive
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}
	so := s.mapping.ToProtoResponsePaginationNominalDeleteAt(paginationMeta, "success", "Successfully fetched active Nominals", nominals)

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
	reqService := requests.FindAllNominals{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	roles, totalRecords, err := s.nominalService.FindByTrashed(&reqService)

	if err != nil {
		return nil, nominal_errors.ErrFindTrashedNominals
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
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
		return nil, nominal_errors.ErrGrpcValidateCreateNominal
	}

	nominal, err := s.nominalService.Create(createReq)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFailedCreateNominal
	}

	so := s.mapping.ToProtoResponseNominal("success", "Successfully created Nominal", nominal)

	return so, nil
}

func (s *nominalHandleGrpc) Update(ctx context.Context, req *pb.UpdateNominalRequest) (*pb.ApiResponseNominal, error) {
	id := int(req.GetId())

	if id == 0 {
		return nil, nominal_errors.ErrGrpcNominalNotFound
	}

	updateReq := &requests.UpdateNominalRequest{
		ID:        id,
		VoucherID: int(req.GetVoucherId()),
		Name:      req.GetName(),
		Quantity:  int(req.GetQuantity()),
		Price:     req.GetPrice(),
	}

	if err := updateReq.Validate(); err != nil {
		return nil, nominal_errors.ErrGrpcValidateUpdateNominal
	}

	nominal, err := s.nominalService.Update(updateReq)
	if err != nil {
		return nil, nominal_errors.ErrGrpcFailedUpdateNominal
	}

	so := s.mapping.ToProtoResponseNominal("success", "Successfully updated Nominal", nominal)

	return so, nil
}

func (s *nominalHandleGrpc) Trashed(ctx context.Context, req *pb.FindByIdNominalRequest) (*pb.ApiResponseNominalDeleteAt, error) {
	id := int(req.GetNominalId())

	if id == 0 {
		return nil, nominal_errors.ErrGrpcNominalNotFound
	}

	Nominal, err := s.nominalService.Trashed(id)

	if err != nil {
		return nil, nominal_errors.ErrGrpcFailedTrashedNominal
	}

	so := s.mapping.ToProtoResponseNominalDeleteAt("success", "Successfully trashed Nominal", Nominal)

	return so, nil
}

func (s *nominalHandleGrpc) Restore(ctx context.Context, req *pb.FindByIdNominalRequest) (*pb.ApiResponseNominalDeleteAt, error) {
	id := int(req.GetNominalId())

	if id == 0 {
		return nil, nominal_errors.ErrGrpcNominalNotFound
	}

	role, err := s.nominalService.Restore(id)

	if err != nil {
		return nil, nominal_errors.ErrGrpcFailedRestoreNominal
	}

	so := s.mapping.ToProtoResponseNominalDeleteAt("success", "Successfully restored Nominal", role)

	return so, nil
}

func (s *nominalHandleGrpc) DeletePermanent(ctx context.Context, req *pb.FindByIdNominalRequest) (*pb.ApiResponseNominalDelete, error) {
	nominal_id := req.NominalId

	_, err := s.nominalService.DeletePermanent(int(nominal_id))

	if err != nil {
		return nil, nominal_errors.ErrGrpcFailedDeletePermanent
	}

	so := s.mapping.ToProtoResponseNominalDelete("success", "Successfully deleted Nominal permanently")

	return so, nil
}

func (s *nominalHandleGrpc) RestoreAll(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseNominalAll, error) {
	_, err := s.nominalService.RestoreAll()

	if err != nil {
		return nil, nominal_errors.ErrGrpcFailedRestoreAll
	}

	so := s.mapping.ToProtoResponseNominalAll("success", "Successfully restored all Nominals")

	return so, nil
}

func (s *nominalHandleGrpc) DeleteAllPermanent(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseNominalAll, error) {
	_, err := s.nominalService.DeleteAllPermanent()

	if err != nil {
		return nil, nominal_errors.ErrGrpcFailedDeleteAll
	}

	so := s.mapping.ToProtoResponseNominalAll("success", "Successfully deleted all Nominals")

	return so, nil
}
