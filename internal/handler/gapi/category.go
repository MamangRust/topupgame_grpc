package gapi

import (
	"context"
	"math"
	"topup_game/internal/domain/requests"
	protomapper "topup_game/internal/mapper/proto"
	"topup_game/internal/pb"
	"topup_game/internal/service"
	"topup_game/pkg/errors/category_errors"

	"google.golang.org/protobuf/types/known/emptypb"
)

type categoryHandleGrpc struct {
	pb.UnimplementedCategoryServiceServer
	categoryService service.CategoryService
	mapping         protomapper.CategoryProtoMapper
}

func NewCategoryHandleGrpc(Category service.CategoryService, mapping protomapper.CategoryProtoMapper) *categoryHandleGrpc {
	return &categoryHandleGrpc{
		categoryService: Category,
		mapping:         mapping,
	}
}

func (s *categoryHandleGrpc) FindAll(ctx context.Context, req *pb.FindAllCategoryRequest) (*pb.ApiResponsePaginationCategory, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllCategory{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	role, totalRecords, err := s.categoryService.FindAll(&reqService)

	if err != nil {
		return nil, category_errors.ErrGrpcFailedFindAll
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationCategory(paginationMeta, "success", "Successfully fetched Category records", role)

	return so, nil
}

func (s *categoryHandleGrpc) FindMonthAmountCategorySuccess(ctx context.Context, req *pb.MonthAmountCategoryRequest) (*pb.ApiResponseCategoryMonthAmountSuccess, error) {
	request := &requests.MonthAmountCategoryRequest{
		Year:  int(req.GetYear()),
		Month: int(req.GetMonth()),
	}

	results, err := s.categoryService.FindMonthAmountCategorySuccess(request)
	if err != nil {
		return nil, category_errors.ErrGrpcFindMonthAmountCategorySuccess
	}

	response := s.mapping.ToProtoResponsesMonthAmountSuccess("success", "Successfully fetched monthly Category success amounts", results)
	return response, nil
}

func (s *categoryHandleGrpc) FindYearAmountCategorySuccess(ctx context.Context, req *pb.YearAmountCategoryRequest) (*pb.ApiResponseCategoryYearAmountSuccess, error) {
	year := int(req.GetYear())

	results, err := s.categoryService.FindYearAmountCategorySuccess(year)
	if err != nil {
		return nil, category_errors.ErrGrpcFindYearAmountCategorySuccess
	}

	response := s.mapping.ToProtoResponseYearAmountSuccess("success", "Successfully fetched yearly Category success amounts", results)
	return response, nil
}

func (s *categoryHandleGrpc) FindMonthAmountCategoryFailed(ctx context.Context, req *pb.MonthAmountCategoryRequest) (*pb.ApiResponseCategoryMonthAmountFailed, error) {
	request := &requests.MonthAmountCategoryRequest{
		Year:  int(req.GetYear()),
		Month: int(req.GetMonth()),
	}

	results, err := s.categoryService.FindMonthAmountCategoryFailed(request)
	if err != nil {
		return nil, category_errors.ErrGrpcFindMonthAmountCategoryFailed
	}

	response := s.mapping.ToProtoResponsesMonthAmountFailed("success", "Successfully fetched monthly Category failed amounts", results)
	return response, nil
}

func (s *categoryHandleGrpc) FindYearAmountCategoryFailed(ctx context.Context, req *pb.YearAmountCategoryRequest) (*pb.ApiResponseCategoryYearAmountFailed, error) {
	year := int(req.GetYear())

	results, err := s.categoryService.FindYearAmountCategoryFailed(year)
	if err != nil {
		return nil, category_errors.ErrGrpcFindYearAmountCategoryFailed
	}

	response := s.mapping.ToProtoResponseYearAmountFailed("success", "Successfully fetched yearly Category failed amounts", results)
	return response, nil
}

func (s *categoryHandleGrpc) FindMonthMethodCategorySuccess(ctx context.Context, req *pb.YearAmountCategoryRequest) (*pb.ApiResponseCategoryMonthMethod, error) {
	year := int(req.GetYear())

	results, err := s.categoryService.FindMonthMethodCategorySuccess(year)
	if err != nil {
		return nil, category_errors.ErrGrpcFindMonthMethodCategorySuccess
	}

	response := s.mapping.ToProtoResponsesMonthMethod("success", "Successfully fetched monthly Category success methods", results)
	return response, nil
}

func (s *categoryHandleGrpc) FindYearMethodCategorySuccess(ctx context.Context, req *pb.YearAmountCategoryRequest) (*pb.ApiResponseCategoryYearMethod, error) {
	year := int(req.GetYear())

	results, err := s.categoryService.FindYearMethodCategorySuccess(year)
	if err != nil {
		return nil, category_errors.ErrGrpcFindYearMethodCategorySuccess
	}

	response := s.mapping.ToProtoResponseYearMethod("success", "Successfully fetched yearly Category success methods", results)
	return response, nil
}

func (s *categoryHandleGrpc) FindMonthMethodCategoryFailed(ctx context.Context, req *pb.YearAmountCategoryRequest) (*pb.ApiResponseCategoryMonthMethod, error) {
	year := int(req.GetYear())

	results, err := s.categoryService.FindMonthMethodCategoryFailed(year)
	if err != nil {
		return nil, category_errors.ErrGrpcFindMonthMethodCategoryFailed
	}

	response := s.mapping.ToProtoResponsesMonthMethod("success", "Successfully fetched monthly Category failed methods", results)
	return response, nil
}

func (s *categoryHandleGrpc) FindMonthAmountCategorySuccessById(ctx context.Context, req *pb.MonthAmountCategoryByIdRequest) (*pb.ApiResponseCategoryMonthAmountSuccess, error) {
	request := &requests.MonthAmountCategoryByIdRequest{
		ID:    int(req.GetId()),
		Year:  int(req.GetYear()),
		Month: int(req.GetMonth()),
	}

	results, err := s.categoryService.FindMonthAmountCategorySuccessById(request)
	if err != nil {
		return nil, category_errors.ErrGrpcFindMonthAmountCategorySuccessById
	}

	response := s.mapping.ToProtoResponsesMonthAmountSuccess(
		"success",
		"Successfully fetched monthly Category success amounts by ID",
		results,
	)
	return response, nil
}

func (s *categoryHandleGrpc) FindYearAmountCategorySuccessById(ctx context.Context, req *pb.YearAmountCategoryByIdRequest) (*pb.ApiResponseCategoryYearAmountSuccess, error) {
	request := &requests.YearAmountCategoryByIdRequest{
		ID:   int(req.GetId()),
		Year: int(req.GetYear()),
	}

	results, err := s.categoryService.FindYearAmountCategorySuccessById(request)
	if err != nil {
		return nil, category_errors.ErrGrpcFindYearAmountCategorySuccessById
	}

	response := s.mapping.ToProtoResponseYearAmountSuccess(
		"success",
		"Successfully fetched yearly Category success amounts by ID",
		results,
	)
	return response, nil
}

func (s *categoryHandleGrpc) FindMonthAmountCategoryFailedById(ctx context.Context, req *pb.MonthAmountCategoryByIdRequest) (*pb.ApiResponseCategoryMonthAmountFailed, error) {
	request := &requests.MonthAmountCategoryByIdRequest{
		ID:    int(req.GetId()),
		Year:  int(req.GetYear()),
		Month: int(req.GetMonth()),
	}

	results, err := s.categoryService.FindMonthAmountCategoryFailedById(request)
	if err != nil {
		return nil, category_errors.ErrGrpcFindMonthAmountCategoryFailedById
	}

	response := s.mapping.ToProtoResponsesMonthAmountFailed(
		"success",
		"Successfully fetched monthly Category failed amounts by ID",
		results,
	)
	return response, nil
}

func (s *categoryHandleGrpc) FindYearAmountCategoryFailedById(ctx context.Context, req *pb.YearAmountCategoryByIdRequest) (*pb.ApiResponseCategoryYearAmountFailed, error) {
	request := &requests.YearAmountCategoryByIdRequest{
		ID:   int(req.GetId()),
		Year: int(req.GetYear()),
	}

	results, err := s.categoryService.FindYearAmountCategoryFailedById(request)
	if err != nil {
		return nil, category_errors.ErrGrpcFindYearAmountCategoryFailedById
	}

	response := s.mapping.ToProtoResponseYearAmountFailed(
		"success",
		"Successfully fetched yearly Category failed amounts by ID",
		results,
	)
	return response, nil
}

func (s *categoryHandleGrpc) FindMonthMethodCategorySuccessById(ctx context.Context, req *pb.MonthMethodCategoryByIdRequest) (*pb.ApiResponseCategoryMonthMethod, error) {
	request := &requests.MonthMethodCategoryByIdRequest{
		ID:   int(req.GetId()),
		Year: int(req.GetYear()),
	}

	results, err := s.categoryService.FindMonthMethodCategorySuccessById(request)
	if err != nil {
		return nil, category_errors.ErrGrpcFindMonthMethodCategorySuccessById
	}

	response := s.mapping.ToProtoResponsesMonthMethod(
		"success",
		"Successfully fetched monthly Category success methods by ID",
		results,
	)
	return response, nil
}

func (s *categoryHandleGrpc) FindYearMethodCategorySuccessById(ctx context.Context, req *pb.YearMethodCategoryByIdRequest) (*pb.ApiResponseCategoryYearMethod, error) {
	request := &requests.YearMethodCategoryByIdRequest{
		ID:   int(req.GetId()),
		Year: int(req.GetYear()),
	}

	results, err := s.categoryService.FindYearMethodCategorySuccessById(request)
	if err != nil {
		return nil, category_errors.ErrGrpcFindYearMethodCategorySuccessById
	}

	response := s.mapping.ToProtoResponseYearMethod(
		"success",
		"Successfully fetched yearly Category success methods by ID",
		results,
	)
	return response, nil
}

func (s *categoryHandleGrpc) FindMonthMethodCategoryFailedById(ctx context.Context, req *pb.MonthMethodCategoryByIdRequest) (*pb.ApiResponseCategoryMonthMethod, error) {
	request := &requests.MonthMethodCategoryByIdRequest{
		ID:   int(req.GetId()),
		Year: int(req.GetYear()),
	}

	results, err := s.categoryService.FindMonthMethodCategoryFailedById(request)
	if err != nil {
		return nil, category_errors.ErrGrpcFindMonthMethodCategoryFailedById
	}

	response := s.mapping.ToProtoResponsesMonthMethod(
		"success",
		"Successfully fetched monthly Category failed methods by ID",
		results,
	)
	return response, nil
}

func (s *categoryHandleGrpc) FindYearMethodCategoryFailedById(ctx context.Context, req *pb.YearMethodCategoryByIdRequest) (*pb.ApiResponseCategoryYearMethod, error) {
	request := &requests.YearMethodCategoryByIdRequest{
		ID:   int(req.GetId()),
		Year: int(req.GetYear()),
	}

	results, err := s.categoryService.FindYearMethodCategoryFailedById(request)
	if err != nil {
		return nil, category_errors.ErrGrpcFindYearMethodCategoryFailedById
	}

	response := s.mapping.ToProtoResponseYearMethod(
		"success",
		"Successfully fetched yearly Category failed methods by ID",
		results,
	)
	return response, nil
}

func (s *categoryHandleGrpc) FindMonthAmountCategorySuccessByMerchant(ctx context.Context, req *pb.MonthAmountCategoryByMerchantRequest) (*pb.ApiResponseCategoryMonthAmountSuccess, error) {
	request := &requests.MonthAmountCategoryByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
		Month:      int(req.GetMonth()),
	}

	results, err := s.categoryService.FindMonthAmountCategorySuccessByMerchant(request)
	if err != nil {
		return nil, category_errors.ErrGrpcFindMonthAmountCategorySuccessByMerchant
	}

	response := s.mapping.ToProtoResponsesMonthAmountSuccess(
		"success",
		"Successfully fetched monthly Category success amounts by ID",
		results,
	)
	return response, nil
}

func (s *categoryHandleGrpc) FindYearAmountCategorySuccessByMerchant(ctx context.Context, req *pb.YearAmountCategoryByMerchantRequest) (*pb.ApiResponseCategoryYearAmountSuccess, error) {
	request := &requests.YearAmountCategoryByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.categoryService.FindYearAmountCategorySuccessByMerchant(request)
	if err != nil {
		return nil, category_errors.ErrGrpcFindYearAmountCategorySuccessByMerchant
	}

	response := s.mapping.ToProtoResponseYearAmountSuccess(
		"success",
		"Successfully fetched yearly Category success amounts by ID",
		results,
	)
	return response, nil
}

func (s *categoryHandleGrpc) FindMonthAmountCategoryFailedByMerchant(ctx context.Context, req *pb.MonthAmountCategoryByMerchantRequest) (*pb.ApiResponseCategoryMonthAmountFailed, error) {
	request := &requests.MonthAmountCategoryByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
		Month:      int(req.GetMonth()),
	}

	results, err := s.categoryService.FindMonthAmountCategoryFailedByMerchant(request)
	if err != nil {
		return nil, category_errors.ErrGrpcFindMonthAmountCategoryFailedByMerchant
	}

	response := s.mapping.ToProtoResponsesMonthAmountFailed(
		"success",
		"Successfully fetched monthly Category failed amounts by ID",
		results,
	)
	return response, nil
}

func (s *categoryHandleGrpc) FindYearAmountCategoryFailedByMerchant(ctx context.Context, req *pb.YearAmountCategoryByMerchantRequest) (*pb.ApiResponseCategoryYearAmountFailed, error) {
	request := &requests.YearAmountCategoryByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.categoryService.FindYearAmountCategoryFailedByMerchant(request)
	if err != nil {
		return nil, category_errors.ErrGrpcFindYearAmountCategoryFailedByMerchant
	}

	response := s.mapping.ToProtoResponseYearAmountFailed(
		"success",
		"Successfully fetched yearly Category failed amounts by ID",
		results,
	)
	return response, nil
}

func (s *categoryHandleGrpc) FindMonthMethodCategorySuccessByMerchant(ctx context.Context, req *pb.MonthMethodCategoryByMerchantRequest) (*pb.ApiResponseCategoryMonthMethod, error) {
	request := &requests.MonthMethodCategoryByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.categoryService.FindMonthMethodCategorySuccessByMerchant(request)
	if err != nil {
		return nil, category_errors.ErrGrpcFindMonthMethodCategorySuccessByMerchant
	}

	response := s.mapping.ToProtoResponsesMonthMethod(
		"success",
		"Successfully fetched monthly Category success methods by ID",
		results,
	)
	return response, nil
}

func (s *categoryHandleGrpc) FindYearMethodCategorySuccessByMerchant(ctx context.Context, req *pb.YearMethodCategoryByMerchantRequest) (*pb.ApiResponseCategoryYearMethod, error) {
	request := &requests.YearMethodCategoryByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.categoryService.FindYearMethodCategorySuccessByMerchant(request)
	if err != nil {
		return nil, category_errors.ErrGrpcFindYearMethodCategorySuccessByMerchant
	}

	response := s.mapping.ToProtoResponseYearMethod(
		"success",
		"Successfully fetched yearly Category success methods by ID",
		results,
	)
	return response, nil
}

func (s *categoryHandleGrpc) FindMonthMethodCategoryFailedByMerchant(ctx context.Context, req *pb.MonthMethodCategoryByMerchantRequest) (*pb.ApiResponseCategoryMonthMethod, error) {
	request := &requests.MonthMethodCategoryByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.categoryService.FindMonthMethodCategoryFailedByMerchant(request)
	if err != nil {
		return nil, category_errors.ErrGrpcFindMonthMethodCategoryFailedByMerchant
	}

	response := s.mapping.ToProtoResponsesMonthMethod(
		"success",
		"Successfully fetched monthly Category failed methods by ID",
		results,
	)
	return response, nil
}

func (s *categoryHandleGrpc) FindYearMethodCategoryFailedByMerchant(ctx context.Context, req *pb.YearMethodCategoryByMerchantRequest) (*pb.ApiResponseCategoryYearMethod, error) {
	request := &requests.YearMethodCategoryByMerchantRequest{
		MerchantID: int(req.GetMerchantId()),
		Year:       int(req.GetYear()),
	}

	results, err := s.categoryService.FindYearMethodCategoryFailedByMerchant(request)
	if err != nil {
		return nil, category_errors.ErrGrpcFindYearMethodCategoryFailedByMerchant
	}

	response := s.mapping.ToProtoResponseYearMethod(
		"success",
		"Successfully fetched yearly Category failed methods by ID",
		results,
	)
	return response, nil
}

func (s *categoryHandleGrpc) FindYearMethodCategoryFailed(ctx context.Context, req *pb.YearAmountCategoryRequest) (*pb.ApiResponseCategoryYearMethod, error) {
	year := int(req.GetYear())

	results, err := s.categoryService.FindYearMethodCategoryFailed(year)
	if err != nil {
		return nil, category_errors.ErrGrpcFindYearMethodCategoryFailed
	}

	response := s.mapping.ToProtoResponseYearMethod("success", "Successfully fetched yearly Category failed methods", results)
	return response, nil
}

func (s *categoryHandleGrpc) FindById(ctx context.Context, req *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategory, error) {
	id := int(req.GetId())

	if id == 0 {
		return nil, category_errors.ErrGrpcCategoryInvalidId
	}

	Category, err := s.categoryService.FindById(id)

	if err != nil {
		return nil, category_errors.ErrGrpcCategoryNotFound
	}

	CategoryResponse := s.mapping.ToProtoResponseCategory("success", "Successfully fetched Category", Category)

	return CategoryResponse, nil
}

func (s *categoryHandleGrpc) FindByActive(ctx context.Context, req *pb.FindAllCategoryRequest) (*pb.ApiResponsePaginationCategoryDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllCategory{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	Categorys, totalRecords, err := s.categoryService.FindByActive(&reqService)

	if err != nil {
		return nil, category_errors.ErrGrpcFailedFindActive
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}
	so := s.mapping.ToProtoResponsePaginationCategoryDeleteAt(paginationMeta, "success", "Successfully fetched active Categorys", Categorys)

	return so, nil
}

func (s *categoryHandleGrpc) FindByTrashed(ctx context.Context, req *pb.FindAllCategoryRequest) (*pb.ApiResponsePaginationCategoryDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllCategory{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	roles, totalRecords, err := s.categoryService.FindByTrashed(&reqService)

	if err != nil {
		return nil, category_errors.ErrGrpcFailedFindTrashed
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}
	so := s.mapping.ToProtoResponsePaginationCategoryDeleteAt(paginationMeta, "success", "Successfully fetched trashed Categorys", roles)

	return so, nil
}

func (s *categoryHandleGrpc) Create(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.ApiResponseCategory, error) {
	name := req.GetName()

	request := &requests.CreateCategoryRequest{
		Name: name,
	}

	if err := request.Validate(); err != nil {
		return nil, category_errors.ErrGrpcValidateCreateCategory
	}

	Category, err := s.categoryService.Create(request)

	if err != nil {
		return nil, category_errors.ErrGrpcFailedCreateCategory
	}

	so := s.mapping.ToProtoResponseCategory("success", "Successfully created Category", Category)

	return so, nil
}

func (s *categoryHandleGrpc) Update(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.ApiResponseCategory, error) {
	id := int(req.GetCategoryId())

	if id == 0 {
		return nil, category_errors.ErrGrpcCategoryInvalidId
	}

	name := req.GetName()

	request := &requests.UpdateCategoryRequest{
		ID:   id,
		Name: name,
	}

	if err := request.Validate(); err != nil {
		return nil, category_errors.ErrGrpcValidateUpdateCategory
	}

	role, err := s.categoryService.Update(request)

	if err != nil {
		return nil, category_errors.ErrGrpcFailedUpdateCategory
	}

	so := s.mapping.ToProtoResponseCategory("success", "Successfully updated Category", role)

	return so, nil
}

func (s *categoryHandleGrpc) Trashed(ctx context.Context, req *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategoryDeleteAt, error) {
	id := int(req.GetId())

	if id == 0 {
		return nil, category_errors.ErrGrpcCategoryInvalidId
	}

	Category, err := s.categoryService.Trashed(id)

	if err != nil {
		return nil, category_errors.ErrGrpcFailedTrashedCategory
	}

	so := s.mapping.ToProtoResponseCategoryDeleteAt("success", "Successfully trashed Category", Category)

	return so, nil
}

func (s *categoryHandleGrpc) Restore(ctx context.Context, req *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategoryDeleteAt, error) {
	id := int(req.GetId())

	if id == 0 {
		return nil, category_errors.ErrGrpcCategoryInvalidId
	}

	role, err := s.categoryService.Restore(id)

	if err != nil {
		return nil, category_errors.ErrGrpcFailedRestoreCategory
	}

	so := s.mapping.ToProtoResponseCategoryDeleteAt("success", "Successfully restored Category", role)

	return so, nil
}

func (s *categoryHandleGrpc) DeletePermanent(ctx context.Context, req *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategoryDelete, error) {
	id := int(req.GetId())

	if id == 0 {
		return nil, category_errors.ErrGrpcCategoryInvalidId
	}

	_, err := s.categoryService.DeletePermanent(id)

	if err != nil {
		return nil, category_errors.ErrGrpcFailedDeletePermanent
	}

	so := s.mapping.ToProtoResponseCategoryDelete("success", "Successfully deleted Category permanently")

	return so, nil
}

func (s *categoryHandleGrpc) RestoreAll(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseCategoryAll, error) {
	_, err := s.categoryService.RestoreAll()

	if err != nil {
		return nil, category_errors.ErrGrpcFailedRestoreAll
	}

	so := s.mapping.ToProtoResponseCategoryAll("success", "Successfully restored all Categorys")

	return so, nil
}

func (s *categoryHandleGrpc) DeleteAllPermanent(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseCategoryAll, error) {
	_, err := s.categoryService.DeleteAllPermanent()

	if err != nil {
		return nil, category_errors.ErrGrpcFailedDeleteAll
	}

	so := s.mapping.ToProtoResponseCategoryAll("success", "Successfully deleted all Categorys")

	return so, nil
}
