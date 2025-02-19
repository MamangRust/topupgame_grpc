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

	role, totalRecords, err := s.categoryService.FindAll(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch Category records: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationCategory(paginationMeta, "success", "Successfully fetched Category records", role)

	return so, nil
}

func (s *categoryHandleGrpc) FindByIdRole(ctx context.Context, req *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategory, error) {
	Category_id := int(req.GetId())

	Category, err := s.categoryService.FindById(Category_id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch Category: " + err.Message,
		})
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

	Categorys, totalRecords, err := s.categoryService.FindByActive(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch active Categorys: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
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

	roles, totalRecords, err := s.categoryService.FindByTrashed(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch trashed Categorys: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
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
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid create category request: " + err.Error(),
		})
	}

	Category, err := s.categoryService.Create(request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create category: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseCategory("success", "Successfully created Category", Category)

	return so, nil
}

func (s *categoryHandleGrpc) Update(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.ApiResponseCategory, error) {
	category_id := int(req.GetCategoryId())
	name := req.GetName()

	request := &requests.UpdateCategoryRequest{
		ID:   category_id,
		Name: name,
	}

	if err := request.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid update category request: " + err.Error(),
		})
	}

	role, err := s.categoryService.Update(request)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update Category: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseCategory("success", "Successfully updated Category", role)

	return so, nil
}

func (s *categoryHandleGrpc) Trashed(ctx context.Context, req *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategoryDeleteAt, error) {
	category_id := req.Id

	Category, err := s.categoryService.Trashed(int(category_id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash Category: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseCategoryDeleteAt("success", "Successfully trashed Category", Category)

	return so, nil
}

func (s *categoryHandleGrpc) Restore(ctx context.Context, req *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategoryDeleteAt, error) {
	category_id := req.Id

	role, err := s.categoryService.Restore(int(category_id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore Category: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseCategoryDeleteAt("success", "Successfully restored Category", role)

	return so, nil
}

func (s *categoryHandleGrpc) DeletePermanent(ctx context.Context, req *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategoryDelete, error) {
	category_id := req.Id

	_, err := s.categoryService.DeletePermanent(int(category_id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete Category permanently: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseCategoryDelete("success", "Successfully deleted Category permanently")

	return so, nil
}

func (s *categoryHandleGrpc) RestoreAll(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseCategoryAll, error) {
	_, err := s.categoryService.RestoreAll()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all Categorys: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseCategoryAll("success", "Successfully restored all Categorys")

	return so, nil
}

func (s *categoryHandleGrpc) DeleteAllPermanent(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseCategoryAll, error) {
	_, err := s.categoryService.DeleteAllPermanent()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete all Categorys permanently: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseCategoryAll("success", "Successfully deleted all Categorys")

	return so, nil
}
