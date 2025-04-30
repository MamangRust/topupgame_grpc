package gapi

import (
	"context"
	"math"
	"topup_game/internal/domain/requests"
	protomapper "topup_game/internal/mapper/proto"
	"topup_game/internal/pb"
	"topup_game/internal/service"
	"topup_game/pkg/errors/bank_errors"

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

	reqService := requests.FindAllBanks{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	role, totalRecords, err := s.bankService.FindAll(&reqService)

	if err != nil {
		return nil, bank_errors.ErrGrpcFailedFindAll
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationBank(paginationMeta, "success", "Successfully fetched bank records", role)

	return so, nil
}

func (s *bankHandleGrpc) FindMonthAmountBankSuccess(ctx context.Context, req *pb.MonthAmountBankRequest) (*pb.ApiResponseBankMonthAmountSuccess, error) {
	request := &requests.MonthAmountBankRequest{
		Year:  int(req.GetYear()),
		Month: int(req.GetMonth()),
	}

	results, err := s.bankService.FindMonthAmountBankSuccess(request)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindMonthAmountBankSuccess
	}

	response := s.mapping.ToProtoResponsesMonthAmountSuccess("success", "Successfully fetched monthly bank success amounts", results)
	return response, nil
}

func (s *bankHandleGrpc) FindYearAmountBankSuccess(ctx context.Context, req *pb.YearAmountBankRequest) (*pb.ApiResponseBankYearAmountSuccess, error) {
	year := int(req.GetYear())

	results, err := s.bankService.FindYearAmountBankSuccess(year)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindYearAmountBankSuccess
	}

	response := s.mapping.ToProtoResponseYearAmountSuccess("success", "Successfully fetched yearly bank success amounts", results)
	return response, nil
}

func (s *bankHandleGrpc) FindMonthAmountBankFailed(ctx context.Context, req *pb.MonthAmountBankRequest) (*pb.ApiResponseBankMonthAmountFailed, error) {
	request := &requests.MonthAmountBankRequest{
		Year:  int(req.GetYear()),
		Month: int(req.GetMonth()),
	}

	results, err := s.bankService.FindMonthAmountBankFailed(request)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindMonthAmountBankFailed
	}

	response := s.mapping.ToProtoResponsesMonthAmountFailed("success", "Successfully fetched monthly bank failed amounts", results)
	return response, nil
}

func (s *bankHandleGrpc) FindYearAmountBankFailed(ctx context.Context, req *pb.YearAmountBankRequest) (*pb.ApiResponseBankYearAmountFailed, error) {
	year := int(req.GetYear())

	results, err := s.bankService.FindYearAmountBankFailed(year)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindYearAmountBankFailed
	}

	response := s.mapping.ToProtoResponseYearAmountFailed("success", "Successfully fetched yearly bank failed amounts", results)
	return response, nil
}

func (s *bankHandleGrpc) FindMonthMethodBankSuccess(ctx context.Context, req *pb.YearAmountBankRequest) (*pb.ApiResponseBankMonthMethod, error) {
	year := int(req.GetYear())

	results, err := s.bankService.FindMonthMethodBankSuccess(year)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindMonthMethodBankSuccess
	}

	response := s.mapping.ToProtoResponsesMonthMethod("success", "Successfully fetched monthly bank success methods", results)
	return response, nil
}

func (s *bankHandleGrpc) FindYearMethodBankSuccess(ctx context.Context, req *pb.YearAmountBankRequest) (*pb.ApiResponseBankYearMethod, error) {
	year := int(req.GetYear())

	results, err := s.bankService.FindYearMethodBankSuccess(year)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindYearMethodBankSuccess
	}

	response := s.mapping.ToProtoResponseYearMethod("success", "Successfully fetched yearly bank success methods", results)
	return response, nil
}

func (s *bankHandleGrpc) FindMonthMethodBankFailed(ctx context.Context, req *pb.YearAmountBankRequest) (*pb.ApiResponseBankMonthMethod, error) {
	year := int(req.GetYear())

	results, err := s.bankService.FindMonthMethodBankFailed(year)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindMonthMethodBankFailed
	}

	response := s.mapping.ToProtoResponsesMonthMethod("success", "Successfully fetched monthly bank failed methods", results)
	return response, nil
}

func (s *bankHandleGrpc) FindYearMethodBankFailed(ctx context.Context, req *pb.YearAmountBankRequest) (*pb.ApiResponseBankYearMethod, error) {
	year := int(req.GetYear())

	results, err := s.bankService.FindYearMethodBankFailed(year)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindYearMethodBankFailed
	}

	response := s.mapping.ToProtoResponseYearMethod("success", "Successfully fetched yearly bank failed methods", results)
	return response, nil
}

func (s *bankHandleGrpc) FindMonthAmountBankSuccessById(ctx context.Context, req *pb.MonthAmountBankByIdRequest) (*pb.ApiResponseBankMonthAmountSuccess, error) {
	request := &requests.MonthAmountBankByIdRequest{
		ID:    int(req.GetId()),
		Year:  int(req.GetYear()),
		Month: int(req.GetMonth()),
	}

	results, err := s.bankService.FindMonthAmountBankSuccessById(request)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindMonthAmountBankSuccessById
	}

	response := s.mapping.ToProtoResponsesMonthAmountSuccess(
		"success",
		"Successfully fetched monthly bank success amounts by ID",
		results,
	)
	return response, nil
}

func (s *bankHandleGrpc) FindYearAmountBankSuccessById(ctx context.Context, req *pb.YearAmountBankByIdRequest) (*pb.ApiResponseBankYearAmountSuccess, error) {
	request := &requests.YearAmountBankByIdRequest{
		ID:   int(req.GetId()),
		Year: int(req.GetYear()),
	}

	results, err := s.bankService.FindYearAmountBankSuccessById(request)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindYearAmountBankSuccessById
	}

	response := s.mapping.ToProtoResponseYearAmountSuccess(
		"success",
		"Successfully fetched yearly bank success amounts by ID",
		results,
	)
	return response, nil
}

func (s *bankHandleGrpc) FindMonthAmountBankFailedById(ctx context.Context, req *pb.MonthAmountBankByIdRequest) (*pb.ApiResponseBankMonthAmountFailed, error) {
	request := &requests.MonthAmountBankByIdRequest{
		ID:    int(req.GetId()),
		Year:  int(req.GetYear()),
		Month: int(req.GetMonth()),
	}

	results, err := s.bankService.FindMonthAmountBankFailedById(request)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindMonthAmountBankFailedById
	}

	response := s.mapping.ToProtoResponsesMonthAmountFailed(
		"success",
		"Successfully fetched monthly bank failed amounts by ID",
		results,
	)
	return response, nil
}

func (s *bankHandleGrpc) FindYearAmountBankFailedById(ctx context.Context, req *pb.YearAmountBankByIdRequest) (*pb.ApiResponseBankYearAmountFailed, error) {
	request := &requests.YearAmountBankByIdRequest{
		ID:   int(req.GetId()),
		Year: int(req.GetYear()),
	}

	results, err := s.bankService.FindYearAmountBankFailedById(request)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindYearAmountBankFailedById
	}

	response := s.mapping.ToProtoResponseYearAmountFailed(
		"success",
		"Successfully fetched yearly bank failed amounts by ID",
		results,
	)
	return response, nil
}

func (s *bankHandleGrpc) FindMonthMethodBankSuccessById(ctx context.Context, req *pb.MonthMethodBankByIdRequest) (*pb.ApiResponseBankMonthMethod, error) {
	request := &requests.MonthMethodBankByIdRequest{
		ID:   int(req.GetId()),
		Year: int(req.GetYear()),
	}

	results, err := s.bankService.FindMonthMethodBankSuccessById(request)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindMonthMethodBankSuccessById
	}

	response := s.mapping.ToProtoResponsesMonthMethod(
		"success",
		"Successfully fetched monthly bank success methods by ID",
		results,
	)
	return response, nil
}

func (s *bankHandleGrpc) FindYearMethodBankSuccessById(ctx context.Context, req *pb.YearMethodBankByIdRequest) (*pb.ApiResponseBankYearMethod, error) {
	request := &requests.YearMethodBankByIdRequest{
		ID:   int(req.GetId()),
		Year: int(req.GetYear()),
	}

	results, err := s.bankService.FindYearMethodBankSuccessById(request)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindYearMethodBankSuccessById
	}

	response := s.mapping.ToProtoResponseYearMethod(
		"success",
		"Successfully fetched yearly bank success methods by ID",
		results,
	)
	return response, nil
}

func (s *bankHandleGrpc) FindMonthMethodBankFailedById(ctx context.Context, req *pb.MonthMethodBankByIdRequest) (*pb.ApiResponseBankMonthMethod, error) {
	request := &requests.MonthMethodBankByIdRequest{
		ID:   int(req.GetId()),
		Year: int(req.GetYear()),
	}

	results, err := s.bankService.FindMonthMethodBankFailedById(request)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindMonthMethodBankFailedById
	}

	response := s.mapping.ToProtoResponsesMonthMethod(
		"success",
		"Successfully fetched monthly bank failed methods by ID",
		results,
	)
	return response, nil
}

func (s *bankHandleGrpc) FindYearMethodBankFailedById(ctx context.Context, req *pb.YearMethodBankByIdRequest) (*pb.ApiResponseBankYearMethod, error) {
	request := &requests.YearMethodBankByIdRequest{
		ID:   int(req.GetId()),
		Year: int(req.GetYear()),
	}

	results, err := s.bankService.FindYearMethodBankFailedById(request)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindYearMethodBankFailedById
	}

	response := s.mapping.ToProtoResponseYearMethod(
		"success",
		"Successfully fetched yearly bank failed methods by ID",
		results,
	)
	return response, nil
}

func (s *bankHandleGrpc) FindMonthAmountBankSuccessByMerchant(ctx context.Context, req *pb.MonthAmountBankByMerchantRequest) (*pb.ApiResponseBankMonthAmountSuccess, error) {
	request := &requests.MonthAmountBankByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
		Month:      int(req.GetMonth()),
	}

	results, err := s.bankService.FindMonthAmountBankSuccessByMerchant(request)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindMonthAmountBankSuccessByMerchant
	}

	response := s.mapping.ToProtoResponsesMonthAmountSuccess(
		"success",
		"Successfully fetched monthly bank success amounts by ID",
		results,
	)
	return response, nil
}

func (s *bankHandleGrpc) FindYearAmountBankSuccessByMerchant(ctx context.Context, req *pb.YearAmountBankByMerchantRequest) (*pb.ApiResponseBankYearAmountSuccess, error) {
	request := &requests.YearAmountBankByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.bankService.FindYearAmountBankSuccessByMerchant(request)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindYearAmountBankSuccessByMerchant
	}

	response := s.mapping.ToProtoResponseYearAmountSuccess(
		"success",
		"Successfully fetched yearly bank success amounts by ID",
		results,
	)
	return response, nil
}

func (s *bankHandleGrpc) FindMonthAmountBankFailedByMerchant(ctx context.Context, req *pb.MonthAmountBankByMerchantRequest) (*pb.ApiResponseBankMonthAmountFailed, error) {
	request := &requests.MonthAmountBankByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
		Month:      int(req.GetMonth()),
	}

	results, err := s.bankService.FindMonthAmountBankFailedByMerchant(request)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindMonthAmountBankFailedByMerchant
	}

	response := s.mapping.ToProtoResponsesMonthAmountFailed(
		"success",
		"Successfully fetched monthly bank failed amounts by ID",
		results,
	)
	return response, nil
}

func (s *bankHandleGrpc) FindYearAmountBankFailedByMerchant(ctx context.Context, req *pb.YearAmountBankByMerchantRequest) (*pb.ApiResponseBankYearAmountFailed, error) {
	request := &requests.YearAmountBankByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.bankService.FindYearAmountBankFailedByMerchant(request)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindYearAmountBankFailedByMerchant
	}

	response := s.mapping.ToProtoResponseYearAmountFailed(
		"success",
		"Successfully fetched yearly bank failed amounts by ID",
		results,
	)
	return response, nil
}

func (s *bankHandleGrpc) FindMonthMethodBankSuccessByMerchant(ctx context.Context, req *pb.MonthMethodBankByMerchantRequest) (*pb.ApiResponseBankMonthMethod, error) {
	request := &requests.MonthMethodBankByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.bankService.FindMonthMethodBankSuccessByMerchant(request)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindMonthMethodBankSuccessByMerchant
	}

	response := s.mapping.ToProtoResponsesMonthMethod(
		"success",
		"Successfully fetched monthly bank success methods by ID",
		results,
	)
	return response, nil
}

func (s *bankHandleGrpc) FindYearMethodBankSuccessByMerchant(ctx context.Context, req *pb.YearMethodBankByMerchantRequest) (*pb.ApiResponseBankYearMethod, error) {
	request := &requests.YearMethodBankByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.bankService.FindYearMethodBankSuccessByMerchant(request)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindYearMethodBankSuccessByMerchant
	}

	response := s.mapping.ToProtoResponseYearMethod(
		"success",
		"Successfully fetched yearly bank success methods by ID",
		results,
	)
	return response, nil
}

func (s *bankHandleGrpc) FindMonthMethodBankFailedByMerchant(ctx context.Context, req *pb.MonthMethodBankByMerchantRequest) (*pb.ApiResponseBankMonthMethod, error) {
	request := &requests.MonthMethodBankByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.bankService.FindMonthMethodBankFailedByMerchant(request)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindMonthMethodBankFailedByMerchant
	}

	response := s.mapping.ToProtoResponsesMonthMethod(
		"success",
		"Successfully fetched monthly bank failed methods by ID",
		results,
	)
	return response, nil
}

func (s *bankHandleGrpc) FindYearMethodBankFailedByMerchant(ctx context.Context, req *pb.YearMethodBankByMerchantRequest) (*pb.ApiResponseBankYearMethod, error) {
	request := &requests.YearMethodBankByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.bankService.FindYearMethodBankFailedByMerchant(request)
	if err != nil {
		return nil, bank_errors.ErrGrpcFindYearMethodBankFailedByMerchant
	}

	response := s.mapping.ToProtoResponseYearMethod(
		"success",
		"Successfully fetched yearly bank failed methods by ID",
		results,
	)
	return response, nil
}

func (s *bankHandleGrpc) FindById(ctx context.Context, req *pb.FindByIdBankRequest) (*pb.ApiResponseBank, error) {
	id := int(req.GetBankId())

	if id == 0 {
		return nil, bank_errors.ErrGrpcBankNotFound
	}

	bank, err := s.bankService.FindByID(id)

	if err != nil {
		return nil, bank_errors.ErrGrpcBankNotFound
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

	reqService := requests.FindAllBanks{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	banks, totalRecords, err := s.bankService.FindByActive(&reqService)

	if err != nil {
		return nil, bank_errors.ErrGrpcFailedFindActive
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
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

	reqService := requests.FindAllBanks{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	roles, totalRecords, err := s.bankService.FindByTrashed(&reqService)

	if err != nil {
		return nil, bank_errors.ErrGrpcFailedFindTrashed
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
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
		return nil, bank_errors.ErrGrpcValidateCreateBank
	}

	bank, err := s.bankService.Create(request)

	if err != nil {
		return nil, bank_errors.ErrGrpcFailedCreateBank
	}

	so := s.mapping.ToProtoResponseBank("success", "Successfully created bank", bank)

	return so, nil
}

func (s *bankHandleGrpc) Update(ctx context.Context, req *pb.UpdateBankRequest) (*pb.ApiResponseBank, error) {
	id := int(req.GetId())

	if id == 0 {
		return nil, bank_errors.ErrGrpcBankNotFound
	}

	name := req.GetName()

	request := &requests.UpdateBankRequest{
		ID:   id,
		Name: name,
	}

	if err := request.Validate(); err != nil {
		return nil, bank_errors.ErrGrpcValidateUpdateBank
	}

	role, err := s.bankService.Update(request)

	if err != nil {
		return nil, bank_errors.ErrGrpcFailedUpdateBank
	}

	so := s.mapping.ToProtoResponseBank("success", "Successfully updated bank", role)

	return so, nil
}

func (s *bankHandleGrpc) Trashed(ctx context.Context, req *pb.FindByIdBankRequest) (*pb.ApiResponseBankDeleteAt, error) {
	id := int(req.GetBankId())

	if id == 0 {
		return nil, bank_errors.ErrGrpcBankNotFound
	}

	bank, err := s.bankService.Trashed(id)

	if err != nil {
		return nil, bank_errors.ErrGrpcFailedTrashedBank
	}

	so := s.mapping.ToProtoResponseBankDeleteAt("success", "Successfully trashed bank", bank)

	return so, nil
}

func (s *bankHandleGrpc) Restore(ctx context.Context, req *pb.FindByIdBankRequest) (*pb.ApiResponseBankDeleteAt, error) {
	id := int(req.GetBankId())

	if id == 0 {
		return nil, bank_errors.ErrGrpcBankNotFound
	}

	role, err := s.bankService.Restore(id)

	if err != nil {
		return nil, bank_errors.ErrGrpcFailedRestoreBank
	}

	so := s.mapping.ToProtoResponseBankDeleteAt("success", "Successfully restored bank", role)

	return so, nil
}

func (s *bankHandleGrpc) DeletePermanent(ctx context.Context, req *pb.FindByIdBankRequest) (*pb.ApiResponseBankDelete, error) {
	id := int(req.GetBankId())

	if id == 0 {
		return nil, bank_errors.ErrGrpcBankNotFound
	}

	_, err := s.bankService.DeletePermanent(id)

	if err != nil {
		return nil, bank_errors.ErrGrpcFailedDeletePermanent
	}

	so := s.mapping.ToProtoResponseBankDelete("success", "Successfully deleted bank permanently")

	return so, nil
}

func (s *bankHandleGrpc) RestoreAll(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseBankAll, error) {
	_, err := s.bankService.RestoreAll()

	if err != nil {
		return nil, bank_errors.ErrGrpcFailedRestoreAll
	}

	so := s.mapping.ToProtoResponseBankAll("success", "Successfully restored all banks")

	return so, nil
}

func (s *bankHandleGrpc) DeleteAllPermanent(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseBankAll, error) {
	_, err := s.bankService.DeleteAllPermanent()

	if err != nil {
		return nil, bank_errors.ErrGrpcFailedDeleteAll
	}

	so := s.mapping.ToProtoResponseBankAll("success", "Successfully deleted all banks")

	return so, nil
}
