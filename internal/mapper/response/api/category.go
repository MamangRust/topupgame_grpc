package response_api

import (
	"topup_game/internal/domain/response"
	"topup_game/internal/pb"
)

type categoryResponseMapper struct {
}

func NewCategoryResponseMapper() *categoryResponseMapper {
	return &categoryResponseMapper{}
}

func (s *categoryResponseMapper) ToApiResponseCategoryAll(pbResponse *pb.ApiResponseCategoryAll) *response.ApiResponseCategoryAll {
	return &response.ApiResponseCategoryAll{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

func (s *categoryResponseMapper) ToApiResponseCategoryDelete(pbResponse *pb.ApiResponseCategoryDelete) *response.ApiResponseCategoryDelete {
	return &response.ApiResponseCategoryDelete{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

func (s *categoryResponseMapper) ToApiResponseCategory(pbResponse *pb.ApiResponseCategory) *response.ApiResponseCategory {
	return &response.ApiResponseCategory{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseCategory(pbResponse.Data),
	}
}

func (s *categoryResponseMapper) ToApiResponseCategoryDeleteAt(pbResponse *pb.ApiResponseCategoryDeleteAt) *response.ApiResponseCategoryDeleteAt {
	return &response.ApiResponseCategoryDeleteAt{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseCategoryDeleteAt(pbResponse.Data),
	}
}

func (s *categoryResponseMapper) ToApiResponsesCategory(pbResponse *pb.ApiResponsesCategory) *response.ApiResponsesCategory {
	return &response.ApiResponsesCategory{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesCategory(pbResponse.Data),
	}
}

func (s *categoryResponseMapper) ToApiResponsePaginationCategory(pbResponse *pb.ApiResponsePaginationCategory) *response.ApiResponsePaginationCategory {
	return &response.ApiResponsePaginationCategory{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.mapResponsesCategory(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
	}
}

func (s *categoryResponseMapper) ToApiResponsePaginationCategoryDeleteAt(pbResponse *pb.ApiResponsePaginationCategoryDeleteAt) *response.ApiResponsePaginationCategoryDeleteAt {
	return &response.ApiResponsePaginationCategoryDeleteAt{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.mapResponsesCategoryDeleteAt(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
	}
}

func (s *categoryResponseMapper) mapResponseCategory(Category *pb.CategoryResponse) *response.CategoryResponse {
	return &response.CategoryResponse{
		ID:        int(Category.Id),
		Name:      Category.Name,
		CreatedAt: Category.CreatedAt,
		UpdatedAt: Category.UpdatedAt,
	}
}

func (s *categoryResponseMapper) mapResponsesCategory(Categorys []*pb.CategoryResponse) []*response.CategoryResponse {
	var responseCategorys []*response.CategoryResponse

	for _, Category := range Categorys {
		responseCategorys = append(responseCategorys, s.mapResponseCategory(Category))
	}

	return responseCategorys
}

func (s *categoryResponseMapper) mapResponseCategoryDeleteAt(Category *pb.CategoryResponseDeleteAt) *response.CategoryResponseDeleteAt {
	return &response.CategoryResponseDeleteAt{
		ID:        int(Category.Id),
		Name:      Category.Name,
		CreatedAt: Category.CreatedAt,
		UpdatedAt: Category.UpdatedAt,
		DeletedAt: Category.DeletedAt,
	}
}

func (s *categoryResponseMapper) mapResponsesCategoryDeleteAt(Categorys []*pb.CategoryResponseDeleteAt) []*response.CategoryResponseDeleteAt {
	var responseCategorys []*response.CategoryResponseDeleteAt

	for _, Category := range Categorys {
		responseCategorys = append(responseCategorys, s.mapResponseCategoryDeleteAt(Category))
	}

	return responseCategorys
}
