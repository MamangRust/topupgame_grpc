package protomapper

import (
	"topup_game/internal/domain/response"
	"topup_game/internal/pb"
)

type categoryProtoMapper struct {
}

func NewCategoryProtoMapper() *categoryProtoMapper {
	return &categoryProtoMapper{}
}

func (s *categoryProtoMapper) ToProtoResponseCategoryAll(status string, message string) *pb.ApiResponseCategoryAll {
	return &pb.ApiResponseCategoryAll{
		Status:  status,
		Message: message,
	}
}

func (s *categoryProtoMapper) ToProtoResponseCategoryDelete(status string, message string) *pb.ApiResponseCategoryDelete {
	return &pb.ApiResponseCategoryDelete{
		Status:  status,
		Message: message,
	}
}

func (s *categoryProtoMapper) ToProtoResponseCategory(status string, message string, pbResponse *response.CategoryResponse) *pb.ApiResponseCategory {
	return &pb.ApiResponseCategory{
		Status:  status,
		Message: message,
		Data:    s.mapResponseCategory(pbResponse),
	}
}

func (s *categoryProtoMapper) ToProtoResponseCategoryDeleteAt(status string, message string, pbResponse *response.CategoryResponseDeleteAt) *pb.ApiResponseCategoryDeleteAt {
	return &pb.ApiResponseCategoryDeleteAt{
		Status:  status,
		Message: message,
		Data:    s.mapResponseCategoryDeleteAt(pbResponse),
	}
}

func (s *categoryProtoMapper) ToProtoResponsesCategory(status string, message string, pbResponse []*response.CategoryResponse) *pb.ApiResponsesCategory {
	return &pb.ApiResponsesCategory{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesCategory(pbResponse),
	}
}

func (s *categoryProtoMapper) ToProtoResponsePaginationCategory(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.CategoryResponse) *pb.ApiResponsePaginationCategory {
	return &pb.ApiResponsePaginationCategory{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesCategory(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *categoryProtoMapper) ToProtoResponsePaginationCategoryDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.CategoryResponseDeleteAt) *pb.ApiResponsePaginationCategoryDeleteAt {
	return &pb.ApiResponsePaginationCategoryDeleteAt{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesCategoryDeleteAt(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *categoryProtoMapper) mapResponseCategory(Category *response.CategoryResponse) *pb.CategoryResponse {
	return &pb.CategoryResponse{
		Id:        int32(Category.ID),
		Name:      Category.Name,
		CreatedAt: Category.CreatedAt,
		UpdatedAt: Category.UpdatedAt,
	}
}

func (s *categoryProtoMapper) mapResponsesCategory(Categorys []*response.CategoryResponse) []*pb.CategoryResponse {
	var responseCategorys []*pb.CategoryResponse

	for _, Category := range Categorys {
		responseCategorys = append(responseCategorys, s.mapResponseCategory(Category))
	}

	return responseCategorys
}

func (s *categoryProtoMapper) mapResponseCategoryDeleteAt(Category *response.CategoryResponseDeleteAt) *pb.CategoryResponseDeleteAt {
	return &pb.CategoryResponseDeleteAt{
		Id:        int32(Category.ID),
		Name:      Category.Name,
		CreatedAt: Category.CreatedAt,
		UpdatedAt: Category.UpdatedAt,
		DeletedAt: Category.DeletedAt,
	}
}

func (s *categoryProtoMapper) mapResponsesCategoryDeleteAt(Categorys []*response.CategoryResponseDeleteAt) []*pb.CategoryResponseDeleteAt {
	var responseCategorys []*pb.CategoryResponseDeleteAt

	for _, Category := range Categorys {
		responseCategorys = append(responseCategorys, s.mapResponseCategoryDeleteAt(Category))
	}

	return responseCategorys
}
