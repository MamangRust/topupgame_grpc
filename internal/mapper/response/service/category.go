package response_service

import (
	"topup_game/internal/domain/record"
	"topup_game/internal/domain/response"
)

type categoryResponseMapper struct {
}

func NewCategoryResponseMapper() *categoryResponseMapper {
	return &categoryResponseMapper{}
}

func (s *categoryResponseMapper) ToCategoryResponse(Category *record.CategoryRecord) *response.CategoryResponse {
	return &response.CategoryResponse{
		ID:        Category.ID,
		Name:      Category.Name,
		CreatedAt: Category.CreatedAt,
		UpdatedAt: Category.UpdatedAt,
	}
}

func (s *categoryResponseMapper) ToCategorysResponse(Categorys []*record.CategoryRecord) []*response.CategoryResponse {
	var responseCategorys []*response.CategoryResponse

	for _, Category := range Categorys {
		responseCategorys = append(responseCategorys, s.ToCategoryResponse(Category))
	}

	return responseCategorys
}

func (s *categoryResponseMapper) ToCategoryResponseDeleteAt(Category *record.CategoryRecord) *response.CategoryResponseDeleteAt {
	return &response.CategoryResponseDeleteAt{
		ID:        Category.ID,
		Name:      Category.Name,
		CreatedAt: Category.CreatedAt,
		UpdatedAt: Category.UpdatedAt,
		DeletedAt: *Category.DeletedAt,
	}
}

func (s *categoryResponseMapper) ToCategorysResponseDeleteAt(Categorys []*record.CategoryRecord) []*response.CategoryResponseDeleteAt {
	var responseCategorys []*response.CategoryResponseDeleteAt

	for _, Category := range Categorys {
		responseCategorys = append(responseCategorys, s.ToCategoryResponseDeleteAt(Category))
	}

	return responseCategorys
}
