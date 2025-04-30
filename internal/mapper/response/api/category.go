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

func (s *categoryResponseMapper) ToApiResponsesMonthAmountSuccess(pbResponse *pb.ApiResponseCategoryMonthAmountSuccess) *response.ApiResponseCategoryMonthAmountSuccess {
	return &response.ApiResponseCategoryMonthAmountSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesCategoryMonthAmountSuccess(pbResponse.Data),
	}
}

func (s *categoryResponseMapper) ToApiResponseYearAmountSuccess(pbResponse *pb.ApiResponseCategoryYearAmountSuccess) *response.ApiResponseCategoryYearAmountSuccess {
	return &response.ApiResponseCategoryYearAmountSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesCategoryYearAmountSuccess(pbResponse.Data),
	}
}

func (s *categoryResponseMapper) ToApiResponsesMonthAmountFailed(pbResponse *pb.ApiResponseCategoryMonthAmountFailed) *response.ApiResponseCategoryMonthAmountFailed {
	return &response.ApiResponseCategoryMonthAmountFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesCategoryMonthAmountFailed(pbResponse.Data),
	}
}

func (s *categoryResponseMapper) ToApiResponseYearAmountFailed(pbResponse *pb.ApiResponseCategoryYearAmountFailed) *response.ApiResponseCategoryYearAmountFailed {
	return &response.ApiResponseCategoryYearAmountFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesCategoryYearAmountFailed(pbResponse.Data),
	}
}

func (s *categoryResponseMapper) ToApiResponsesMonthMethod(pbResponse *pb.ApiResponseCategoryMonthMethod) *response.ApiResponseCategoryMonthMethod {
	return &response.ApiResponseCategoryMonthMethod{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesCategoryMonthMethod(pbResponse.Data),
	}
}

func (s *categoryResponseMapper) ToApiResponseYearMethod(pbResponse *pb.ApiResponseCategoryYearMethod) *response.ApiResponseCategoryYearMethod {
	return &response.ApiResponseCategoryYearMethod{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesCategoryYearMethod(pbResponse.Data),
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

func (s *categoryResponseMapper) mapResponseCategoryMonthAmountSuccess(b *pb.MonthAmountCategorySuccessResponse) *response.MonthAmountCategorySuccessResponse {
	return &response.MonthAmountCategorySuccessResponse{
		ID:           int(b.Id),
		CategoryName: b.CategoryName,
		Year:         b.Year,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *categoryResponseMapper) mapResponsesCategoryMonthAmountSuccess(b []*pb.MonthAmountCategorySuccessResponse) []*response.MonthAmountCategorySuccessResponse {
	var result []*response.MonthAmountCategorySuccessResponse

	for _, Category := range b {
		result = append(result, s.mapResponseCategoryMonthAmountSuccess(Category))
	}

	return result
}

func (s *categoryResponseMapper) mapResponseCategoryYearAmountSuccess(b *pb.YearAmountCategorySuccessResponse) *response.YearAmountCategorySuccessResponse {
	return &response.YearAmountCategorySuccessResponse{
		ID:           int(b.Id),
		CategoryName: b.CategoryName,
		Year:         b.Year,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *categoryResponseMapper) mapResponsesCategoryYearAmountSuccess(b []*pb.YearAmountCategorySuccessResponse) []*response.YearAmountCategorySuccessResponse {
	var result []*response.YearAmountCategorySuccessResponse

	for _, Category := range b {
		result = append(result, s.mapResponseCategoryYearAmountSuccess(Category))
	}

	return result
}

func (s *categoryResponseMapper) mapResponseCategoryMonthAmountFailed(b *pb.MonthAmountCategoryFailedResponse) *response.MonthAmountCategoryFailedResponse {
	return &response.MonthAmountCategoryFailedResponse{
		ID:           int(b.Id),
		CategoryName: b.CategoryName,
		Year:         b.Year,
		Month:        b.Month,
		TotalFailed:  int(b.TotalFailed),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *categoryResponseMapper) mapResponsesCategoryMonthAmountFailed(b []*pb.MonthAmountCategoryFailedResponse) []*response.MonthAmountCategoryFailedResponse {
	var result []*response.MonthAmountCategoryFailedResponse

	for _, Category := range b {
		result = append(result, s.mapResponseCategoryMonthAmountFailed(Category))
	}

	return result
}

func (s *categoryResponseMapper) mapResponseCategoryYearAmountFailed(b *pb.YearAmountCategoryFailedResponse) *response.YearAmountCategoryFailedResponse {
	return &response.YearAmountCategoryFailedResponse{
		ID:           int(b.Id),
		CategoryName: b.CategoryName,
		Year:         b.Year,
		TotalFailed:  int(b.TotalFailed),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *categoryResponseMapper) mapResponsesCategoryYearAmountFailed(b []*pb.YearAmountCategoryFailedResponse) []*response.YearAmountCategoryFailedResponse {
	var result []*response.YearAmountCategoryFailedResponse

	for _, Category := range b {
		result = append(result, s.mapResponseCategoryYearAmountFailed(Category))
	}

	return result
}

func (s *categoryResponseMapper) mapResponseCategoryMonthMethod(b *pb.MonthMethodCategoryResponse) *response.MonthMethodCategoryResponse {
	return &response.MonthMethodCategoryResponse{
		ID:                int(b.Id),
		Month:             b.Month,
		CategoryName:      b.CategoryName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s categoryResponseMapper) mapResponsesCategoryMonthMethod(b []*pb.MonthMethodCategoryResponse) []*response.MonthMethodCategoryResponse {
	var result []*response.MonthMethodCategoryResponse

	for _, category := range b {
		result = append(result, s.mapResponseCategoryMonthMethod(category))
	}

	return result
}

func (s *categoryResponseMapper) mapResponseCategoryYearMethod(b *pb.YearMethodCategoryResponse) *response.YearMethodCategoryResponse {
	return &response.YearMethodCategoryResponse{
		ID:                int(b.Id),
		Year:              b.Year,
		CategoryName:      b.CategoryName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s categoryResponseMapper) mapResponsesCategoryYearMethod(b []*pb.YearMethodCategoryResponse) []*response.YearMethodCategoryResponse {
	var result []*response.YearMethodCategoryResponse

	for _, category := range b {
		result = append(result, s.mapResponseCategoryYearMethod(category))
	}

	return result
}
