package protomapper

import (
	"topup_game/internal/domain/response"
	"topup_game/internal/pb"
)

type bankProtoMapper struct {
}

func NewBankProtoMapper() *bankProtoMapper {
	return &bankProtoMapper{}
}

func (s *bankProtoMapper) ToProtoResponseBankAll(status string, message string) *pb.ApiResponseBankAll {
	return &pb.ApiResponseBankAll{
		Status:  status,
		Message: message,
	}
}

func (s *bankProtoMapper) ToProtoResponseBankDelete(status string, message string) *pb.ApiResponseBankDelete {
	return &pb.ApiResponseBankDelete{
		Status:  status,
		Message: message,
	}
}

func (s *bankProtoMapper) ToProtoResponseBank(status string, message string, pbResponse *response.BankResponse) *pb.ApiResponseBank {
	return &pb.ApiResponseBank{
		Status:  status,
		Message: message,
		Data:    s.mapResponseBank(pbResponse),
	}
}

func (s *bankProtoMapper) ToProtoResponseBankDeleteAt(status string, message string, pbResponse *response.BankResponseDeleteAt) *pb.ApiResponseBankDeleteAt {
	return &pb.ApiResponseBankDeleteAt{
		Status:  status,
		Message: message,
		Data:    s.mapResponseBankDeleteAt(pbResponse),
	}
}

func (s *bankProtoMapper) ToProtoResponsesBank(status string, message string, pbResponse []*response.BankResponse) *pb.ApiResponsesBank {
	return &pb.ApiResponsesBank{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesBank(pbResponse),
	}
}

func (s *bankProtoMapper) ToProtoResponsePaginationBank(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.BankResponse) *pb.ApiResponsePaginationBank {
	return &pb.ApiResponsePaginationBank{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesBank(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *bankProtoMapper) ToProtoResponsePaginationBankDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.BankResponseDeleteAt) *pb.ApiResponsePaginationBankDeleteAt {
	return &pb.ApiResponsePaginationBankDeleteAt{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesBankDeleteAt(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *bankProtoMapper) ToProtoResponsesMonthAmountSuccess(status string, message string, pbResponse []*response.MonthAmountBankSuccessResponse) *pb.ApiResponseBankMonthAmountSuccess {
	return &pb.ApiResponseBankMonthAmountSuccess{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesBankMonthAmountSuccess(pbResponse),
	}
}

func (s *bankProtoMapper) ToProtoResponseYearAmountSuccess(status string, message string, pbResponse []*response.YearAmountBankSuccessResponse) *pb.ApiResponseBankYearAmountSuccess {
	return &pb.ApiResponseBankYearAmountSuccess{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesBankYearAmountSuccess(pbResponse),
	}
}

func (s *bankProtoMapper) ToProtoResponsesMonthAmountFailed(status string, message string, pbResponse []*response.MonthAmountBankFailedResponse) *pb.ApiResponseBankMonthAmountFailed {
	return &pb.ApiResponseBankMonthAmountFailed{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesBankMonthAmountFailed(pbResponse),
	}
}

func (s *bankProtoMapper) ToProtoResponseYearAmountFailed(status string, message string, pbResponse []*response.YearAmountBankFailedResponse) *pb.ApiResponseBankYearAmountFailed {
	return &pb.ApiResponseBankYearAmountFailed{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesBankYearAmountFailed(pbResponse),
	}
}

func (s *bankProtoMapper) ToProtoResponsesMonthMethod(status string, message string, pbResponse []*response.MonthMethodBankResponse) *pb.ApiResponseBankMonthMethod {
	return &pb.ApiResponseBankMonthMethod{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesBankMonthMethod(pbResponse),
	}
}

func (s *bankProtoMapper) ToProtoResponseYearMethod(status string, message string, pbResponse []*response.YearMethodBankResponse) *pb.ApiResponseBankYearMethod {
	return &pb.ApiResponseBankYearMethod{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesBankYearMethod(pbResponse),
	}
}

func (s *bankProtoMapper) mapResponseBank(Bank *response.BankResponse) *pb.BankResponse {
	return &pb.BankResponse{
		Id:        int32(Bank.ID),
		Name:      Bank.Name,
		CreatedAt: Bank.CreatedAt,
		UpdatedAt: Bank.UpdatedAt,
	}
}

func (s *bankProtoMapper) mapResponsesBank(Banks []*response.BankResponse) []*pb.BankResponse {
	var responseBanks []*pb.BankResponse

	for _, Bank := range Banks {
		responseBanks = append(responseBanks, s.mapResponseBank(Bank))
	}

	return responseBanks
}

func (s *bankProtoMapper) mapResponseBankDeleteAt(Bank *response.BankResponseDeleteAt) *pb.BankResponseDeleteAt {
	return &pb.BankResponseDeleteAt{
		Id:        int32(Bank.ID),
		Name:      Bank.Name,
		CreatedAt: Bank.CreatedAt,
		UpdatedAt: Bank.UpdatedAt,
		DeletedAt: Bank.DeletedAt,
	}
}

func (s *bankProtoMapper) mapResponsesBankDeleteAt(Banks []*response.BankResponseDeleteAt) []*pb.BankResponseDeleteAt {
	var responseBanks []*pb.BankResponseDeleteAt

	for _, Bank := range Banks {
		responseBanks = append(responseBanks, s.mapResponseBankDeleteAt(Bank))
	}

	return responseBanks
}

func (s *bankProtoMapper) mapResponseBankMonthAmountSuccess(b *response.MonthAmountBankSuccessResponse) *pb.MonthAmountBankSuccessResponse {
	return &pb.MonthAmountBankSuccessResponse{
		Id:           int32(b.ID),
		BankName:     b.BankName,
		Year:         b.Year,
		Month:        b.Month,
		TotalSuccess: int32(b.TotalSuccess),
		TotalAmount:  int32(b.TotalAmount),
	}
}

func (s *bankProtoMapper) mapResponsesBankMonthAmountSuccess(b []*response.MonthAmountBankSuccessResponse) []*pb.MonthAmountBankSuccessResponse {
	var result []*pb.MonthAmountBankSuccessResponse

	for _, Bank := range b {
		result = append(result, s.mapResponseBankMonthAmountSuccess(Bank))
	}

	return result
}

func (s *bankProtoMapper) mapResponseBankYearAmountSuccess(b *response.YearAmountBankSuccessResponse) *pb.YearAmountBankSuccessResponse {
	return &pb.YearAmountBankSuccessResponse{
		Id:           int32(b.ID),
		BankName:     b.BankName,
		Year:         b.Year,
		TotalSuccess: int32(b.TotalSuccess),
		TotalAmount:  int32(b.TotalAmount),
	}
}

func (s *bankProtoMapper) mapResponsesBankYearAmountSuccess(b []*response.YearAmountBankSuccessResponse) []*pb.YearAmountBankSuccessResponse {
	var result []*pb.YearAmountBankSuccessResponse

	for _, Bank := range b {
		result = append(result, s.mapResponseBankYearAmountSuccess(Bank))
	}

	return result
}

func (s *bankProtoMapper) mapResponseBankMonthAmountFailed(b *response.MonthAmountBankFailedResponse) *pb.MonthAmountBankFailedResponse {
	return &pb.MonthAmountBankFailedResponse{
		Id:          int32(b.ID),
		BankName:    b.BankName,
		Year:        b.Year,
		Month:       b.Month,
		TotalFailed: int32(b.TotalFailed),
		TotalAmount: int32(b.TotalAmount),
	}
}

func (s *bankProtoMapper) mapResponsesBankMonthAmountFailed(b []*response.MonthAmountBankFailedResponse) []*pb.MonthAmountBankFailedResponse {
	var result []*pb.MonthAmountBankFailedResponse

	for _, Bank := range b {
		result = append(result, s.mapResponseBankMonthAmountFailed(Bank))
	}

	return result
}

func (s *bankProtoMapper) mapResponseBankYearAmountFailed(b *response.YearAmountBankFailedResponse) *pb.YearAmountBankFailedResponse {
	return &pb.YearAmountBankFailedResponse{
		Id:          int32(b.ID),
		BankName:    b.BankName,
		Year:        b.Year,
		TotalFailed: int32(b.TotalFailed),
		TotalAmount: int32(b.TotalAmount),
	}
}

func (s *bankProtoMapper) mapResponsesBankYearAmountFailed(b []*response.YearAmountBankFailedResponse) []*pb.YearAmountBankFailedResponse {
	var result []*pb.YearAmountBankFailedResponse

	for _, Bank := range b {
		result = append(result, s.mapResponseBankYearAmountFailed(Bank))
	}

	return result
}

func (s *bankProtoMapper) mapResponseBankMonthMethod(b *response.MonthMethodBankResponse) *pb.MonthMethodBankResponse {
	return &pb.MonthMethodBankResponse{
		Id:                int32(b.ID),
		Month:             b.Month,
		BankName:          b.BankName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int32(b.TotalAmount),
		TotalTransactions: int32(b.TotalTransactions),
	}
}

func (s bankProtoMapper) mapResponsesBankMonthMethod(b []*response.MonthMethodBankResponse) []*pb.MonthMethodBankResponse {
	var result []*pb.MonthMethodBankResponse

	for _, Bank := range b {
		result = append(result, s.mapResponseBankMonthMethod(Bank))
	}

	return result
}

func (s *bankProtoMapper) mapResponseBankYearMethod(b *response.YearMethodBankResponse) *pb.YearMethodBankResponse {
	return &pb.YearMethodBankResponse{
		Id:                int32(b.ID),
		Year:              b.Year,
		BankName:          b.BankName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int32(b.TotalAmount),
		TotalTransactions: int32(b.TotalTransactions),
	}
}

func (s bankProtoMapper) mapResponsesBankYearMethod(b []*response.YearMethodBankResponse) []*pb.YearMethodBankResponse {
	var result []*pb.YearMethodBankResponse

	for _, bank := range b {
		result = append(result, s.mapResponseBankYearMethod(bank))
	}

	return result
}
