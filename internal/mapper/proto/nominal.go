package protomapper

import (
	"topup_game/internal/domain/response"
	"topup_game/internal/pb"
)

type nominalProtoMapper struct{}

func NewNominalProtoMapper() *nominalProtoMapper {
	return &nominalProtoMapper{}
}

func (s *nominalProtoMapper) ToProtoResponseNominalAll(status string, message string) *pb.ApiResponseNominalAll {
	return &pb.ApiResponseNominalAll{
		Status:  status,
		Message: message,
	}
}

func (s *nominalProtoMapper) ToProtoResponseNominalDelete(status string, message string) *pb.ApiResponseNominalDelete {
	return &pb.ApiResponseNominalDelete{
		Status:  status,
		Message: message,
	}
}

func (s *nominalProtoMapper) ToProtoResponseNominal(status string, message string, pbResponse *response.NominalResponse) *pb.ApiResponseNominal {
	return &pb.ApiResponseNominal{
		Status:  status,
		Message: message,
		Data:    s.mapResponseNominal(pbResponse),
	}
}

func (s *nominalProtoMapper) ToProtoResponseNominalDeleteAt(status string, message string, pbResponse *response.NominalResponseDeleteAt) *pb.ApiResponseNominalDeleteAt {
	return &pb.ApiResponseNominalDeleteAt{
		Status:  status,
		Message: message,
		Data:    s.mapResponseNominalDeleteAt(pbResponse),
	}
}

func (s *nominalProtoMapper) ToProtoResponsesNominal(status string, message string, pbResponse []*response.NominalResponse) *pb.ApiResponsesNominal {
	return &pb.ApiResponsesNominal{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesNominal(pbResponse),
	}
}

func (s *nominalProtoMapper) ToProtoResponsePaginationNominal(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.NominalResponse) *pb.ApiResponsePaginationNominal {
	return &pb.ApiResponsePaginationNominal{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesNominal(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *nominalProtoMapper) ToProtoResponsePaginationNominalDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.NominalResponseDeleteAt) *pb.ApiResponsePaginationNominalDeleteAt {
	return &pb.ApiResponsePaginationNominalDeleteAt{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesNominalDeleteAt(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *nominalProtoMapper) ToProtoResponsesMonthAmountSuccess(status string, message string, pbResponse []*response.MonthAmountNominalSuccessResponse) *pb.ApiResponseNominalMonthAmountSuccess {
	return &pb.ApiResponseNominalMonthAmountSuccess{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesNominalMonthAmountSuccess(pbResponse),
	}
}

func (s *nominalProtoMapper) ToProtoResponseYearAmountSuccess(status string, message string, pbResponse []*response.YearAmountNominalSuccessResponse) *pb.ApiResponseNominalYearAmountSuccess {
	return &pb.ApiResponseNominalYearAmountSuccess{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesNominalYearAmountSuccess(pbResponse),
	}
}

func (s *nominalProtoMapper) ToProtoResponsesMonthAmountFailed(status string, message string, pbResponse []*response.MonthAmountNominalFailedResponse) *pb.ApiResponseNominalMonthAmountFailed {
	return &pb.ApiResponseNominalMonthAmountFailed{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesNominalMonthAmountFailed(pbResponse),
	}
}

func (s *nominalProtoMapper) ToProtoResponseYearAmountFailed(status string, message string, pbResponse []*response.YearAmountNominalFailedResponse) *pb.ApiResponseNominalYearAmountFailed {
	return &pb.ApiResponseNominalYearAmountFailed{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesNominalYearAmountFailed(pbResponse),
	}
}

func (s *nominalProtoMapper) ToProtoResponsesMonthMethod(status string, message string, pbResponse []*response.MonthMethodNominalResponse) *pb.ApiResponseNominalMonthMethod {
	return &pb.ApiResponseNominalMonthMethod{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesNominalMonthMethod(pbResponse),
	}
}

func (s *nominalProtoMapper) ToProtoResponseYearMethod(status string, message string, pbResponse []*response.YearMethodNominalResponse) *pb.ApiResponseNominalYearMethod {
	return &pb.ApiResponseNominalYearMethod{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesNominalYearMethod(pbResponse),
	}
}

func (s *nominalProtoMapper) mapResponseNominal(nominal *response.NominalResponse) *pb.NominalResponse {
	return &pb.NominalResponse{
		Id:        int32(nominal.ID),
		Name:      nominal.Name,
		Quantity:  int32(nominal.Quantity),
		Price:     nominal.Price,
		CreatedAt: nominal.CreatedAt,
		UpdatedAt: nominal.UpdatedAt,
	}
}

func (s *nominalProtoMapper) mapResponsesNominal(nominals []*response.NominalResponse) []*pb.NominalResponse {
	var responseNominals []*pb.NominalResponse

	for _, nominal := range nominals {
		responseNominals = append(responseNominals, s.mapResponseNominal(nominal))
	}

	return responseNominals
}

func (s *nominalProtoMapper) mapResponseNominalDeleteAt(nominal *response.NominalResponseDeleteAt) *pb.NominalResponseDeleteAt {
	return &pb.NominalResponseDeleteAt{
		Id:        int32(nominal.ID),
		Name:      nominal.Name,
		Quantity:  int32(nominal.Quantity),
		Price:     nominal.Price,
		CreatedAt: nominal.CreatedAt,
		UpdatedAt: nominal.UpdatedAt,
		DeletedAt: nominal.DeletedAt,
	}
}

func (s *nominalProtoMapper) mapResponsesNominalDeleteAt(nominals []*response.NominalResponseDeleteAt) []*pb.NominalResponseDeleteAt {
	var responseNominals []*pb.NominalResponseDeleteAt

	for _, nominal := range nominals {
		responseNominals = append(responseNominals, s.mapResponseNominalDeleteAt(nominal))
	}

	return responseNominals
}

func (s *nominalProtoMapper) mapResponseNominalMonthAmountSuccess(b *response.MonthAmountNominalSuccessResponse) *pb.MonthAmountNominalSuccessResponse {
	return &pb.MonthAmountNominalSuccessResponse{
		Id:           int32(b.ID),
		NominalName:  b.NominalName,
		Year:         b.Year,
		Month:        b.Month,
		TotalSuccess: int32(b.TotalSuccess),
		TotalAmount:  int32(b.TotalAmount),
	}
}

func (s *nominalProtoMapper) mapResponsesNominalMonthAmountSuccess(b []*response.MonthAmountNominalSuccessResponse) []*pb.MonthAmountNominalSuccessResponse {
	var result []*pb.MonthAmountNominalSuccessResponse

	for _, nominal := range b {
		result = append(result, s.mapResponseNominalMonthAmountSuccess(nominal))
	}

	return result
}

func (s *nominalProtoMapper) mapResponseNominalYearAmountSuccess(b *response.YearAmountNominalSuccessResponse) *pb.YearAmountNominalSuccessResponse {
	return &pb.YearAmountNominalSuccessResponse{
		Id:           int32(b.ID),
		NominalName:  b.NominalName,
		Year:         b.Year,
		TotalSuccess: int32(b.TotalSuccess),
		TotalAmount:  int32(b.TotalAmount),
	}
}

func (s *nominalProtoMapper) mapResponsesNominalYearAmountSuccess(b []*response.YearAmountNominalSuccessResponse) []*pb.YearAmountNominalSuccessResponse {
	var result []*pb.YearAmountNominalSuccessResponse

	for _, nominal := range b {
		result = append(result, s.mapResponseNominalYearAmountSuccess(nominal))
	}

	return result
}

func (s *nominalProtoMapper) mapResponseNominalMonthAmountFailed(b *response.MonthAmountNominalFailedResponse) *pb.MonthAmountNominalFailedResponse {
	return &pb.MonthAmountNominalFailedResponse{
		Id:          int32(b.ID),
		NominalName: b.NominalName,
		Year:        b.Year,
		Month:       b.Month,
		TotalFailed: int32(b.TotalFailed),
		TotalAmount: int32(b.TotalAmount),
	}
}

func (s *nominalProtoMapper) mapResponsesNominalMonthAmountFailed(b []*response.MonthAmountNominalFailedResponse) []*pb.MonthAmountNominalFailedResponse {
	var result []*pb.MonthAmountNominalFailedResponse

	for _, Bank := range b {
		result = append(result, s.mapResponseNominalMonthAmountFailed(Bank))
	}

	return result
}

func (s *nominalProtoMapper) mapResponseNominalYearAmountFailed(b *response.YearAmountNominalFailedResponse) *pb.YearAmountNominalFailedResponse {
	return &pb.YearAmountNominalFailedResponse{
		Id:          int32(b.ID),
		NominalName: b.NominalName,
		Year:        b.Year,
		TotalFailed: int32(b.TotalFailed),
		TotalAmount: int32(b.TotalAmount),
	}
}

func (s *nominalProtoMapper) mapResponsesNominalYearAmountFailed(b []*response.YearAmountNominalFailedResponse) []*pb.YearAmountNominalFailedResponse {
	var result []*pb.YearAmountNominalFailedResponse

	for _, Bank := range b {
		result = append(result, s.mapResponseNominalYearAmountFailed(Bank))
	}

	return result
}

func (s *nominalProtoMapper) mapResponseNominalMonthMethod(b *response.MonthMethodNominalResponse) *pb.MonthMethodNominalResponse {
	return &pb.MonthMethodNominalResponse{
		Id:                int32(b.ID),
		Month:             b.Month,
		NominalName:       b.NominalName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int32(b.TotalAmount),
		TotalTransactions: int32(b.TotalTransactions),
	}
}

func (s nominalProtoMapper) mapResponsesNominalMonthMethod(b []*response.MonthMethodNominalResponse) []*pb.MonthMethodNominalResponse {
	var result []*pb.MonthMethodNominalResponse

	for _, nominal := range b {
		result = append(result, s.mapResponseNominalMonthMethod(nominal))
	}

	return result
}

func (s *nominalProtoMapper) mapResponseNominalYearMethod(b *response.YearMethodNominalResponse) *pb.YearMethodNominalResponse {
	return &pb.YearMethodNominalResponse{
		Id:                int32(b.ID),
		Year:              b.Year,
		NominalName:       b.NominalName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int32(b.TotalAmount),
		TotalTransactions: int32(b.TotalTransactions),
	}
}

func (s nominalProtoMapper) mapResponsesNominalYearMethod(b []*response.YearMethodNominalResponse) []*pb.YearMethodNominalResponse {
	var result []*pb.YearMethodNominalResponse

	for _, nominal := range b {
		result = append(result, s.mapResponseNominalYearMethod(nominal))
	}

	return result
}
