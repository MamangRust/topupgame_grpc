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

func (s *categoryResponseMapper) ToCategoriesResponse(Categorys []*record.CategoryRecord) []*response.CategoryResponse {
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

func (s *categoryResponseMapper) ToCategoriesResponseDeleteAt(Categorys []*record.CategoryRecord) []*response.CategoryResponseDeleteAt {
	var responseCategorys []*response.CategoryResponseDeleteAt

	for _, Category := range Categorys {
		responseCategorys = append(responseCategorys, s.ToCategoryResponseDeleteAt(Category))
	}

	return responseCategorys
}

func (s *categoryResponseMapper) ToCategoryResponseMonthAmountSuccess(b *record.MonthAmountCategorySuccessRecord) *response.MonthAmountCategorySuccessResponse {
	return &response.MonthAmountCategorySuccessResponse{
		ID:           int(b.ID),
		CategoryName: b.CategoryName,
		Year:         b.Year,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *categoryResponseMapper) ToCategoriesResponseMonthAmountSuccess(b []*record.MonthAmountCategorySuccessRecord) []*response.MonthAmountCategorySuccessResponse {
	var result []*response.MonthAmountCategorySuccessResponse

	for _, Bank := range b {
		result = append(result, s.ToCategoryResponseMonthAmountSuccess(Bank))
	}

	return result
}

func (s *categoryResponseMapper) ToCategoryResponseYearAmountSuccess(b *record.YearAmountCategorySuccessRecord) *response.YearAmountCategorySuccessResponse {
	return &response.YearAmountCategorySuccessResponse{
		ID:           int(b.ID),
		CategoryName: b.CategoryName,
		Year:         b.Year,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *categoryResponseMapper) ToCategoriesResponseYearAmountSuccess(b []*record.YearAmountCategorySuccessRecord) []*response.YearAmountCategorySuccessResponse {
	var result []*response.YearAmountCategorySuccessResponse

	for _, Bank := range b {
		result = append(result, s.ToCategoryResponseYearAmountSuccess(Bank))
	}

	return result
}

func (s *categoryResponseMapper) ToCategoryResponseMonthAmountFailed(b *record.MonthAmountCategoryFailedRecord) *response.MonthAmountCategoryFailedResponse {
	return &response.MonthAmountCategoryFailedResponse{
		ID:           int(b.ID),
		CategoryName: b.CategoryName,
		Year:         b.Year,
		Month:        b.Month,
		TotalFailed:  int(b.TotalFailed),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *categoryResponseMapper) ToCategoriesResponseMonthAmountFailed(b []*record.MonthAmountCategoryFailedRecord) []*response.MonthAmountCategoryFailedResponse {
	var result []*response.MonthAmountCategoryFailedResponse

	for _, Bank := range b {
		result = append(result, s.ToCategoryResponseMonthAmountFailed(Bank))
	}

	return result
}

func (s *categoryResponseMapper) ToCategoryResponseYearAmountFailed(b *record.YearAmountCategoryFailedRecord) *response.YearAmountCategoryFailedResponse {
	return &response.YearAmountCategoryFailedResponse{
		ID:           int(b.ID),
		CategoryName: b.CategoryName,
		Year:         b.Year,
		TotalFailed:  int(b.TotalFailed),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *categoryResponseMapper) ToCategoriesResponseYearAmountFailed(b []*record.YearAmountCategoryFailedRecord) []*response.YearAmountCategoryFailedResponse {
	var result []*response.YearAmountCategoryFailedResponse

	for _, Bank := range b {
		result = append(result, s.ToCategoryResponseYearAmountFailed(Bank))
	}

	return result
}

func (s *categoryResponseMapper) ToCategoryResponseMonthMethod(b *record.MonthMethodCategoryRecord) *response.MonthMethodCategoryResponse {
	return &response.MonthMethodCategoryResponse{
		ID:                int(b.ID),
		Month:             b.Month,
		CategoryName:      b.CategoryName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s categoryResponseMapper) ToCategoriesResponseMonthMethod(b []*record.MonthMethodCategoryRecord) []*response.MonthMethodCategoryResponse {
	var result []*response.MonthMethodCategoryResponse

	for _, category := range b {
		result = append(result, s.ToCategoryResponseMonthMethod(category))
	}

	return result
}

func (s *categoryResponseMapper) ToCategoryResponseYearMethod(b *record.YearMethodCategoryRecord) *response.YearMethodCategoryResponse {
	return &response.YearMethodCategoryResponse{
		ID:                int(b.ID),
		Year:              b.Year,
		CategoryName:      b.CategoryName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s categoryResponseMapper) ToCategoriesResponseYearMethod(b []*record.YearMethodCategoryRecord) []*response.YearMethodCategoryResponse {
	var result []*response.YearMethodCategoryResponse

	for _, category := range b {
		result = append(result, s.ToCategoryResponseYearMethod(category))
	}

	return result
}
