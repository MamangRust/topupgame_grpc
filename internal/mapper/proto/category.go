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

func (s *categoryProtoMapper) ToProtoResponsesMonthAmountSuccess(status string, message string, pbResponse []*response.MonthAmountCategorySuccessResponse) *pb.ApiResponseCategoryMonthAmountSuccess {
	return &pb.ApiResponseCategoryMonthAmountSuccess{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesCategoryMonthAmountSuccess(pbResponse),
	}
}

func (s *categoryProtoMapper) ToProtoResponseYearAmountSuccess(status string, message string, pbResponse []*response.YearAmountCategorySuccessResponse) *pb.ApiResponseCategoryYearAmountSuccess {
	return &pb.ApiResponseCategoryYearAmountSuccess{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesCategoryYearAmountSuccess(pbResponse),
	}
}

func (s *categoryProtoMapper) ToProtoResponsesMonthAmountFailed(status string, message string, pbResponse []*response.MonthAmountCategoryFailedResponse) *pb.ApiResponseCategoryMonthAmountFailed {
	return &pb.ApiResponseCategoryMonthAmountFailed{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesCategoryMonthAmountFailed(pbResponse),
	}
}

func (s *categoryProtoMapper) ToProtoResponseYearAmountFailed(status string, message string, pbResponse []*response.YearAmountCategoryFailedResponse) *pb.ApiResponseCategoryYearAmountFailed {
	return &pb.ApiResponseCategoryYearAmountFailed{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesCategoryYearAmountFailed(pbResponse),
	}
}

func (s *categoryProtoMapper) ToProtoResponsesMonthMethod(status string, message string, pbResponse []*response.MonthMethodCategoryResponse) *pb.ApiResponseCategoryMonthMethod {
	return &pb.ApiResponseCategoryMonthMethod{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesCategoryMonthMethod(pbResponse),
	}
}

func (s *categoryProtoMapper) ToProtoResponseYearMethod(status string, message string, pbResponse []*response.YearMethodCategoryResponse) *pb.ApiResponseCategoryYearMethod {
	return &pb.ApiResponseCategoryYearMethod{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesCategoryYearMethod(pbResponse),
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

func (s *categoryProtoMapper) mapResponseCategoryMonthAmountSuccess(b *response.MonthAmountCategorySuccessResponse) *pb.MonthAmountCategorySuccessResponse {
	return &pb.MonthAmountCategorySuccessResponse{
		Id:           int32(b.ID),
		CategoryName: b.CategoryName,
		Year:         b.Year,
		Month:        b.Month,
		TotalSuccess: int32(b.TotalSuccess),
		TotalAmount:  int32(b.TotalAmount),
	}
}

func (s *categoryProtoMapper) mapResponsesCategoryMonthAmountSuccess(b []*response.MonthAmountCategorySuccessResponse) []*pb.MonthAmountCategorySuccessResponse {
	var result []*pb.MonthAmountCategorySuccessResponse

	for _, Bank := range b {
		result = append(result, s.mapResponseCategoryMonthAmountSuccess(Bank))
	}

	return result
}

func (s *categoryProtoMapper) mapResponseCategoryYearAmountSuccess(b *response.YearAmountCategorySuccessResponse) *pb.YearAmountCategorySuccessResponse {
	return &pb.YearAmountCategorySuccessResponse{
		Id:           int32(b.ID),
		CategoryName: b.CategoryName,
		Year:         b.Year,
		TotalSuccess: int32(b.TotalSuccess),
		TotalAmount:  int32(b.TotalAmount),
	}
}

func (s *categoryProtoMapper) mapResponsesCategoryYearAmountSuccess(b []*response.YearAmountCategorySuccessResponse) []*pb.YearAmountCategorySuccessResponse {
	var result []*pb.YearAmountCategorySuccessResponse

	for _, Bank := range b {
		result = append(result, s.mapResponseCategoryYearAmountSuccess(Bank))
	}

	return result
}

func (s *categoryProtoMapper) mapResponseCategoryMonthAmountFailed(b *response.MonthAmountCategoryFailedResponse) *pb.MonthAmountCategoryFailedResponse {
	return &pb.MonthAmountCategoryFailedResponse{
		Id:           int32(b.ID),
		CategoryName: b.CategoryName,
		Year:         b.Year,
		Month:        b.Month,
		TotalFailed:  int32(b.TotalFailed),
		TotalAmount:  int32(b.TotalAmount),
	}
}

func (s *categoryProtoMapper) mapResponsesCategoryMonthAmountFailed(b []*response.MonthAmountCategoryFailedResponse) []*pb.MonthAmountCategoryFailedResponse {
	var result []*pb.MonthAmountCategoryFailedResponse

	for _, Bank := range b {
		result = append(result, s.mapResponseCategoryMonthAmountFailed(Bank))
	}

	return result
}

func (s *categoryProtoMapper) mapResponseCategoryYearAmountFailed(b *response.YearAmountCategoryFailedResponse) *pb.YearAmountCategoryFailedResponse {
	return &pb.YearAmountCategoryFailedResponse{
		Id:           int32(b.ID),
		CategoryName: b.CategoryName,
		Year:         b.Year,
		TotalFailed:  int32(b.TotalFailed),
		TotalAmount:  int32(b.TotalAmount),
	}
}

func (s *categoryProtoMapper) mapResponsesCategoryYearAmountFailed(b []*response.YearAmountCategoryFailedResponse) []*pb.YearAmountCategoryFailedResponse {
	var result []*pb.YearAmountCategoryFailedResponse

	for _, Bank := range b {
		result = append(result, s.mapResponseCategoryYearAmountFailed(Bank))
	}

	return result
}

func (s *categoryProtoMapper) mapResponseCategoryMonthMethod(b *response.MonthMethodCategoryResponse) *pb.MonthMethodCategoryResponse {
	return &pb.MonthMethodCategoryResponse{
		Id:                int32(b.ID),
		Month:             b.Month,
		CategoryName:      b.CategoryName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int32(b.TotalAmount),
		TotalTransactions: int32(b.TotalTransactions),
	}
}

func (s categoryProtoMapper) mapResponsesCategoryMonthMethod(b []*response.MonthMethodCategoryResponse) []*pb.MonthMethodCategoryResponse {
	var result []*pb.MonthMethodCategoryResponse

	for _, category := range b {
		result = append(result, s.mapResponseCategoryMonthMethod(category))
	}

	return result
}

func (s *categoryProtoMapper) mapResponseCategoryYearMethod(b *response.YearMethodCategoryResponse) *pb.YearMethodCategoryResponse {
	return &pb.YearMethodCategoryResponse{
		Id:                int32(b.ID),
		Year:              b.Year,
		CategoryName:      b.CategoryName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int32(b.TotalAmount),
		TotalTransactions: int32(b.TotalTransactions),
	}
}

func (s categoryProtoMapper) mapResponsesCategoryYearMethod(b []*response.YearMethodCategoryResponse) []*pb.YearMethodCategoryResponse {
	var result []*pb.YearMethodCategoryResponse

	for _, category := range b {
		result = append(result, s.mapResponseCategoryYearMethod(category))
	}

	return result
}
