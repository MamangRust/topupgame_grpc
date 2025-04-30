package response_api

import (
	"topup_game/internal/domain/response"
	"topup_game/internal/pb"
)

type nominalResponseMapper struct {
}

func NewNominalResponseMapper() *nominalResponseMapper {
	return &nominalResponseMapper{}
}

func (s *nominalResponseMapper) ToApiResponseNominal(pbResponse *pb.ApiResponseNominal) *response.ApiResponseNominal {
	return &response.ApiResponseNominal{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseNominal(pbResponse.Data),
	}
}

func (s *nominalResponseMapper) ToApiResponseNominalAll(pbResponse *pb.ApiResponseNominalAll) *response.ApiResponseNominalAll {
	return &response.ApiResponseNominalAll{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

func (s *nominalResponseMapper) ToApiResponseNominalDelete(pbResponse *pb.ApiResponseNominalDelete) *response.ApiResponseNominalDelete {
	return &response.ApiResponseNominalDelete{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

func (s *nominalResponseMapper) ToApiResponseNominalDeleteAt(pbResponse *pb.ApiResponseNominalDeleteAt) *response.ApiResponseNominalDeleteAt {
	return &response.ApiResponseNominalDeleteAt{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseNominalDeleteAt(pbResponse.Data),
	}
}

func (s *nominalResponseMapper) ToApiResponsesNominal(pbResponse *pb.ApiResponsesNominal) *response.ApiResponsesNominal {
	return &response.ApiResponsesNominal{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesNominal(pbResponse.Data),
	}
}

func (s *nominalResponseMapper) ToApiResponsePaginationNominal(pbResponse *pb.ApiResponsePaginationNominal) *response.ApiResponsePaginationNominal {
	return &response.ApiResponsePaginationNominal{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.mapResponsesNominal(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
	}
}

func (s *nominalResponseMapper) ToApiResponsePaginationNominalDeleteAt(pbResponse *pb.ApiResponsePaginationNominalDeleteAt) *response.ApiResponsePaginationNominalDeleteAt {
	return &response.ApiResponsePaginationNominalDeleteAt{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.mapResponsesNominalDeleteAt(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
	}
}

func (s *nominalResponseMapper) ToApiResponsesMonthAmountSuccess(pbResponse *pb.ApiResponseNominalMonthAmountSuccess) *response.ApiResponseNominalMonthAmountSuccess {
	return &response.ApiResponseNominalMonthAmountSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesNominalMonthAmountSuccess(pbResponse.Data),
	}
}

func (s *nominalResponseMapper) ToApiResponseYearAmountSuccess(pbResponse *pb.ApiResponseNominalYearAmountSuccess) *response.ApiResponseNominalYearAmountSuccess {
	return &response.ApiResponseNominalYearAmountSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesNominalYearAmountSuccess(pbResponse.Data),
	}
}

func (s *nominalResponseMapper) ToApiResponsesMonthAmountFailed(pbResponse *pb.ApiResponseNominalMonthAmountFailed) *response.ApiResponseNominalMonthAmountFailed {
	return &response.ApiResponseNominalMonthAmountFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesNominalMonthAmountFailed(pbResponse.Data),
	}
}

func (s *nominalResponseMapper) ToApiResponseYearAmountFailed(pbResponse *pb.ApiResponseNominalYearAmountFailed) *response.ApiResponseNominalYearAmountFailed {
	return &response.ApiResponseNominalYearAmountFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesNominalYearAmountFailed(pbResponse.Data),
	}
}

func (s *nominalResponseMapper) ToApiResponsesMonthMethod(pbResponse *pb.ApiResponseNominalMonthMethod) *response.ApiResponseNominalMonthMethod {
	return &response.ApiResponseNominalMonthMethod{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesNominalMonthMethod(pbResponse.Data),
	}
}

func (s *nominalResponseMapper) ToApiResponseYearMethod(pbResponse *pb.ApiResponseNominalYearMethod) *response.ApiResponseNominalYearMethod {
	return &response.ApiResponseNominalYearMethod{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesNominalYearMethod(pbResponse.Data),
	}
}

func (s *nominalResponseMapper) mapResponseNominal(nominal *pb.NominalResponse) *response.NominalResponse {
	return &response.NominalResponse{
		ID:        int(nominal.Id),
		Name:      nominal.Name,
		Quantity:  int(nominal.Quantity),
		Price:     nominal.Price,
		CreatedAt: nominal.CreatedAt,
		UpdatedAt: nominal.UpdatedAt,
	}
}

func (s *nominalResponseMapper) mapResponsesNominal(nominals []*pb.NominalResponse) []*response.NominalResponse {
	var responseNominals []*response.NominalResponse

	for _, nominal := range nominals {
		responseNominals = append(responseNominals, s.mapResponseNominal(nominal))
	}

	return responseNominals
}

func (s *nominalResponseMapper) mapResponseNominalDeleteAt(nominal *pb.NominalResponseDeleteAt) *response.NominalResponseDeleteAt {
	return &response.NominalResponseDeleteAt{
		ID:        int(nominal.Id),
		Name:      nominal.Name,
		Quantity:  int(nominal.Quantity),
		Price:     nominal.Price,
		CreatedAt: nominal.CreatedAt,
		UpdatedAt: nominal.UpdatedAt,
		DeletedAt: nominal.DeletedAt,
	}
}

func (s *nominalResponseMapper) mapResponsesNominalDeleteAt(nominals []*pb.NominalResponseDeleteAt) []*response.NominalResponseDeleteAt {
	var responseNominals []*response.NominalResponseDeleteAt

	for _, nominal := range nominals {
		responseNominals = append(responseNominals, s.mapResponseNominalDeleteAt(nominal))
	}

	return responseNominals
}

func (s *nominalResponseMapper) mapResponseNominalMonthAmountSuccess(b *pb.MonthAmountNominalSuccessResponse) *response.MonthAmountNominalSuccessResponse {
	return &response.MonthAmountNominalSuccessResponse{
		ID:           int(b.Id),
		NominalName:  b.NominalName,
		Year:         b.Year,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *nominalResponseMapper) mapResponsesNominalMonthAmountSuccess(b []*pb.MonthAmountNominalSuccessResponse) []*response.MonthAmountNominalSuccessResponse {
	var result []*response.MonthAmountNominalSuccessResponse

	for _, nominal := range b {
		result = append(result, s.mapResponseNominalMonthAmountSuccess(nominal))
	}

	return result
}

func (s *nominalResponseMapper) mapResponseNominalYearAmountSuccess(b *pb.YearAmountNominalSuccessResponse) *response.YearAmountNominalSuccessResponse {
	return &response.YearAmountNominalSuccessResponse{
		ID:           int(b.Id),
		NominalName:  b.NominalName,
		Year:         b.Year,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *nominalResponseMapper) mapResponsesNominalYearAmountSuccess(b []*pb.YearAmountNominalSuccessResponse) []*response.YearAmountNominalSuccessResponse {
	var result []*response.YearAmountNominalSuccessResponse

	for _, nominal := range b {
		result = append(result, s.mapResponseNominalYearAmountSuccess(nominal))
	}

	return result
}

func (s *nominalResponseMapper) mapResponseNominalMonthAmountFailed(b *pb.MonthAmountNominalFailedResponse) *response.MonthAmountNominalFailedResponse {
	return &response.MonthAmountNominalFailedResponse{
		ID:          int(b.Id),
		NominalName: b.NominalName,
		Year:        b.Year,
		Month:       b.Month,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *nominalResponseMapper) mapResponsesNominalMonthAmountFailed(b []*pb.MonthAmountNominalFailedResponse) []*response.MonthAmountNominalFailedResponse {
	var result []*response.MonthAmountNominalFailedResponse

	for _, Bank := range b {
		result = append(result, s.mapResponseNominalMonthAmountFailed(Bank))
	}

	return result
}

func (s *nominalResponseMapper) mapResponseNominalYearAmountFailed(b *pb.YearAmountNominalFailedResponse) *response.YearAmountNominalFailedResponse {
	return &response.YearAmountNominalFailedResponse{
		ID:          int(b.Id),
		NominalName: b.NominalName,
		Year:        b.Year,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *nominalResponseMapper) mapResponsesNominalYearAmountFailed(b []*pb.YearAmountNominalFailedResponse) []*response.YearAmountNominalFailedResponse {
	var result []*response.YearAmountNominalFailedResponse

	for _, Bank := range b {
		result = append(result, s.mapResponseNominalYearAmountFailed(Bank))
	}

	return result
}

func (s *nominalResponseMapper) mapResponseNominalMonthMethod(b *pb.MonthMethodNominalResponse) *response.MonthMethodNominalResponse {
	return &response.MonthMethodNominalResponse{
		ID:                int(b.Id),
		Month:             b.Month,
		NominalName:       b.NominalName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s nominalResponseMapper) mapResponsesNominalMonthMethod(b []*pb.MonthMethodNominalResponse) []*response.MonthMethodNominalResponse {
	var result []*response.MonthMethodNominalResponse

	for _, nominal := range b {
		result = append(result, s.mapResponseNominalMonthMethod(nominal))
	}

	return result
}

func (s *nominalResponseMapper) mapResponseNominalYearMethod(b *pb.YearMethodNominalResponse) *response.YearMethodNominalResponse {
	return &response.YearMethodNominalResponse{
		ID:                int(b.Id),
		Year:              b.Year,
		NominalName:       b.NominalName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s nominalResponseMapper) mapResponsesNominalYearMethod(b []*pb.YearMethodNominalResponse) []*response.YearMethodNominalResponse {
	var result []*response.YearMethodNominalResponse

	for _, nominal := range b {
		result = append(result, s.mapResponseNominalYearMethod(nominal))
	}

	return result
}
