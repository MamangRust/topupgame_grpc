package response_api

import (
	"topup_game/internal/domain/response"
	"topup_game/internal/pb"
)

type bankResponseMapper struct {
}

func NewBankResponseMapper() *bankResponseMapper {
	return &bankResponseMapper{}
}

func (s *bankResponseMapper) ToApiResponseBankAll(pbResponse *pb.ApiResponseBankAll) *response.ApiResponseBankAll {
	return &response.ApiResponseBankAll{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

func (s *bankResponseMapper) ToApiResponseBankDelete(pbResponse *pb.ApiResponseBankDelete) *response.ApiResponseBankDelete {
	return &response.ApiResponseBankDelete{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

func (s *bankResponseMapper) ToApiResponseBank(pbResponse *pb.ApiResponseBank) *response.ApiResponseBank {
	return &response.ApiResponseBank{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseBank(pbResponse.Data),
	}
}

func (s *bankResponseMapper) ToApiResponseBankDeleteAt(pbResponse *pb.ApiResponseBankDeleteAt) *response.ApiResponseBankDeleteAt {
	return &response.ApiResponseBankDeleteAt{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseBankDeleteAt(pbResponse.Data),
	}
}

func (s *bankResponseMapper) ToApiResponsesBank(pbResponse *pb.ApiResponsesBank) *response.ApiResponsesBank {
	return &response.ApiResponsesBank{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesBank(pbResponse.Data),
	}
}

func (s *bankResponseMapper) ToApiResponsePaginationBank(pbResponse *pb.ApiResponsePaginationBank) *response.ApiResponsePaginationBank {
	return &response.ApiResponsePaginationBank{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.mapResponsesBank(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
	}
}

func (s *bankResponseMapper) ToApiResponsePaginationBankDeleteAt(pbResponse *pb.ApiResponsePaginationBankDeleteAt) *response.ApiResponsePaginationBankDeleteAt {
	return &response.ApiResponsePaginationBankDeleteAt{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.mapResponsesBankDeleteAt(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
	}
}

func (s *bankResponseMapper) ToApiResponsesMonthAmountSuccess(pbResponse *pb.ApiResponseBankMonthAmountSuccess) *response.ApiResponseBankMonthAmountSuccess {
	return &response.ApiResponseBankMonthAmountSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesBankMonthAmountSuccess(pbResponse.Data),
	}
}

func (s *bankResponseMapper) ToApiResponseYearAmountSuccess(pbResponse *pb.ApiResponseBankYearAmountSuccess) *response.ApiResponseBankYearAmountSuccess {
	return &response.ApiResponseBankYearAmountSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesBankYearAmountSuccess(pbResponse.Data),
	}
}

func (s *bankResponseMapper) ToApiResponsesMonthAmountFailed(pbResponse *pb.ApiResponseBankMonthAmountFailed) *response.ApiResponseBankMonthAmountFailed {
	return &response.ApiResponseBankMonthAmountFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesBankMonthAmountFailed(pbResponse.Data),
	}
}

func (s *bankResponseMapper) ToApiResponseYearAmountFailed(pbResponse *pb.ApiResponseBankYearAmountFailed) *response.ApiResponseBankYearAmountFailed {
	return &response.ApiResponseBankYearAmountFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesBankYearAmountFailed(pbResponse.Data),
	}
}

func (s *bankResponseMapper) ToApiResponsesMonthMethod(pbResponse *pb.ApiResponseBankMonthMethod) *response.ApiResponseBankMonthMethod {
	return &response.ApiResponseBankMonthMethod{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesBankMonthMethod(pbResponse.Data),
	}
}

func (s *bankResponseMapper) ToApiResponseYearMethod(pbResponse *pb.ApiResponseBankYearMethod) *response.ApiResponseBankYearMethod {
	return &response.ApiResponseBankYearMethod{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesBankYearMethod(pbResponse.Data),
	}
}

func (s *bankResponseMapper) mapResponseBank(Bank *pb.BankResponse) *response.BankResponse {
	return &response.BankResponse{
		ID:        int(Bank.Id),
		Name:      Bank.Name,
		CreatedAt: Bank.CreatedAt,
		UpdatedAt: Bank.UpdatedAt,
	}
}

func (s *bankResponseMapper) mapResponsesBank(Banks []*pb.BankResponse) []*response.BankResponse {
	var responseBanks []*response.BankResponse

	for _, Bank := range Banks {
		responseBanks = append(responseBanks, s.mapResponseBank(Bank))
	}

	return responseBanks
}

func (s *bankResponseMapper) mapResponseBankDeleteAt(Bank *pb.BankResponseDeleteAt) *response.BankResponseDeleteAt {
	return &response.BankResponseDeleteAt{
		ID:        int(Bank.Id),
		Name:      Bank.Name,
		CreatedAt: Bank.CreatedAt,
		UpdatedAt: Bank.UpdatedAt,
		DeletedAt: Bank.DeletedAt,
	}
}

func (s *bankResponseMapper) mapResponsesBankDeleteAt(Banks []*pb.BankResponseDeleteAt) []*response.BankResponseDeleteAt {
	var responseBanks []*response.BankResponseDeleteAt

	for _, Bank := range Banks {
		responseBanks = append(responseBanks, s.mapResponseBankDeleteAt(Bank))
	}

	return responseBanks
}

func (s *bankResponseMapper) mapResponseBankMonthAmountSuccess(b *pb.MonthAmountBankSuccessResponse) *response.MonthAmountBankSuccessResponse {
	return &response.MonthAmountBankSuccessResponse{
		ID:           int(b.Id),
		BankName:     b.BankName,
		Year:         b.Year,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *bankResponseMapper) mapResponsesBankMonthAmountSuccess(b []*pb.MonthAmountBankSuccessResponse) []*response.MonthAmountBankSuccessResponse {
	var result []*response.MonthAmountBankSuccessResponse

	for _, Bank := range b {
		result = append(result, s.mapResponseBankMonthAmountSuccess(Bank))
	}

	return result
}

func (s *bankResponseMapper) mapResponseBankYearAmountSuccess(b *pb.YearAmountBankSuccessResponse) *response.YearAmountBankSuccessResponse {
	return &response.YearAmountBankSuccessResponse{
		ID:           int(b.Id),
		BankName:     b.BankName,
		Year:         b.Year,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *bankResponseMapper) mapResponsesBankYearAmountSuccess(b []*pb.YearAmountBankSuccessResponse) []*response.YearAmountBankSuccessResponse {
	var result []*response.YearAmountBankSuccessResponse

	for _, Bank := range b {
		result = append(result, s.mapResponseBankYearAmountSuccess(Bank))
	}

	return result
}

func (s *bankResponseMapper) mapResponseBankMonthAmountFailed(b *pb.MonthAmountBankFailedResponse) *response.MonthAmountBankFailedResponse {
	return &response.MonthAmountBankFailedResponse{
		ID:          int(b.Id),
		BankName:    b.BankName,
		Year:        b.Year,
		Month:       b.Month,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *bankResponseMapper) mapResponsesBankMonthAmountFailed(b []*pb.MonthAmountBankFailedResponse) []*response.MonthAmountBankFailedResponse {
	var result []*response.MonthAmountBankFailedResponse

	for _, Bank := range b {
		result = append(result, s.mapResponseBankMonthAmountFailed(Bank))
	}

	return result
}

func (s *bankResponseMapper) mapResponseBankYearAmountFailed(b *pb.YearAmountBankFailedResponse) *response.YearAmountBankFailedResponse {
	return &response.YearAmountBankFailedResponse{
		ID:          int(b.Id),
		BankName:    b.BankName,
		Year:        b.Year,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *bankResponseMapper) mapResponsesBankYearAmountFailed(b []*pb.YearAmountBankFailedResponse) []*response.YearAmountBankFailedResponse {
	var result []*response.YearAmountBankFailedResponse

	for _, Bank := range b {
		result = append(result, s.mapResponseBankYearAmountFailed(Bank))
	}

	return result
}

func (s *bankResponseMapper) mapResponseBankMonthMethod(b *pb.MonthMethodBankResponse) *response.MonthMethodBankResponse {
	return &response.MonthMethodBankResponse{
		ID:                int(b.Id),
		Month:             b.Month,
		BankName:          b.BankName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s bankResponseMapper) mapResponsesBankMonthMethod(b []*pb.MonthMethodBankResponse) []*response.MonthMethodBankResponse {
	var result []*response.MonthMethodBankResponse

	for _, Bank := range b {
		result = append(result, s.mapResponseBankMonthMethod(Bank))
	}

	return result
}

func (s *bankResponseMapper) mapResponseBankYearMethod(b *pb.YearMethodBankResponse) *response.YearMethodBankResponse {
	return &response.YearMethodBankResponse{
		ID:                int(b.Id),
		Year:              b.Year,
		BankName:          b.BankName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s bankResponseMapper) mapResponsesBankYearMethod(b []*pb.YearMethodBankResponse) []*response.YearMethodBankResponse {
	var result []*response.YearMethodBankResponse

	for _, bank := range b {
		result = append(result, s.mapResponseBankYearMethod(bank))
	}

	return result
}
