package response_service

import (
	"topup_game/internal/domain/record"
	"topup_game/internal/domain/response"
)

type bankResponseMapper struct {
}

func NewBankResponseMapper() *bankResponseMapper {
	return &bankResponseMapper{}
}

func (s *bankResponseMapper) ToBankResponse(Bank *record.BankRecord) *response.BankResponse {
	return &response.BankResponse{
		ID:        Bank.ID,
		Name:      Bank.Name,
		CreatedAt: Bank.CreatedAt,
		UpdatedAt: Bank.UpdatedAt,
	}
}

func (s *bankResponseMapper) ToBanksResponse(Banks []*record.BankRecord) []*response.BankResponse {
	var responseBanks []*response.BankResponse

	for _, Bank := range Banks {
		responseBanks = append(responseBanks, s.ToBankResponse(Bank))
	}

	return responseBanks
}

func (s *bankResponseMapper) ToBankResponseDeleteAt(Bank *record.BankRecord) *response.BankResponseDeleteAt {
	return &response.BankResponseDeleteAt{
		ID:        Bank.ID,
		Name:      Bank.Name,
		CreatedAt: Bank.CreatedAt,
		UpdatedAt: Bank.UpdatedAt,
		DeletedAt: *Bank.DeletedAt,
	}
}

func (s *bankResponseMapper) ToBanksResponseDeleteAt(Banks []*record.BankRecord) []*response.BankResponseDeleteAt {
	var responseBanks []*response.BankResponseDeleteAt

	for _, Bank := range Banks {
		responseBanks = append(responseBanks, s.ToBankResponseDeleteAt(Bank))
	}

	return responseBanks
}

func (s *bankResponseMapper) ToBankResponseMonthAmountSuccess(b *record.MonthAmountBankSuccessRecord) *response.MonthAmountBankSuccessResponse {
	return &response.MonthAmountBankSuccessResponse{
		ID:           int(b.ID),
		BankName:     b.BankName,
		Year:         b.Year,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *bankResponseMapper) ToBanksResponseMonthAmountSuccess(b []*record.MonthAmountBankSuccessRecord) []*response.MonthAmountBankSuccessResponse {
	var result []*response.MonthAmountBankSuccessResponse

	for _, Bank := range b {
		result = append(result, s.ToBankResponseMonthAmountSuccess(Bank))
	}

	return result
}

func (s *bankResponseMapper) ToBankResponseYearAmountSuccess(b *record.YearAmountBankSuccessRecord) *response.YearAmountBankSuccessResponse {
	return &response.YearAmountBankSuccessResponse{
		ID:           int(b.ID),
		BankName:     b.BankName,
		Year:         b.Year,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *bankResponseMapper) ToBanksResponseYearAmountSuccess(b []*record.YearAmountBankSuccessRecord) []*response.YearAmountBankSuccessResponse {
	var result []*response.YearAmountBankSuccessResponse

	for _, Bank := range b {
		result = append(result, s.ToBankResponseYearAmountSuccess(Bank))
	}

	return result
}

func (s *bankResponseMapper) ToBankResponseMonthAmountFailed(b *record.MonthAmountBankFailedRecord) *response.MonthAmountBankFailedResponse {
	return &response.MonthAmountBankFailedResponse{
		ID:          int(b.ID),
		BankName:    b.BankName,
		Year:        b.Year,
		Month:       b.Month,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *bankResponseMapper) ToBanksResponseMonthAmountFailed(b []*record.MonthAmountBankFailedRecord) []*response.MonthAmountBankFailedResponse {
	var result []*response.MonthAmountBankFailedResponse

	for _, Bank := range b {
		result = append(result, s.ToBankResponseMonthAmountFailed(Bank))
	}

	return result
}

func (s *bankResponseMapper) ToBankResponseYearAmountFailed(b *record.YearAmountBankFailedRecord) *response.YearAmountBankFailedResponse {
	return &response.YearAmountBankFailedResponse{
		ID:          int(b.ID),
		BankName:    b.BankName,
		Year:        b.Year,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *bankResponseMapper) ToBanksResponseYearAmountFailed(b []*record.YearAmountBankFailedRecord) []*response.YearAmountBankFailedResponse {
	var result []*response.YearAmountBankFailedResponse

	for _, Bank := range b {
		result = append(result, s.ToBankResponseYearAmountFailed(Bank))
	}

	return result
}

func (s *bankResponseMapper) ToBankResponseMonthMethod(b *record.MonthMethodBankRecord) *response.MonthMethodBankResponse {
	return &response.MonthMethodBankResponse{
		ID:                int(b.ID),
		Month:             b.Month,
		BankName:          b.BankName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s bankResponseMapper) ToBanksResponseMonthMethod(b []*record.MonthMethodBankRecord) []*response.MonthMethodBankResponse {
	var result []*response.MonthMethodBankResponse

	for _, Bank := range b {
		result = append(result, s.ToBankResponseMonthMethod(Bank))
	}

	return result
}

func (s *bankResponseMapper) ToBankResponseYearMethod(b *record.YearMethodBankRecord) *response.YearMethodBankResponse {
	return &response.YearMethodBankResponse{
		ID:                int(b.ID),
		Year:              b.Year,
		BankName:          b.BankName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s bankResponseMapper) ToBanksResponseYearMethod(b []*record.YearMethodBankRecord) []*response.YearMethodBankResponse {
	var result []*response.YearMethodBankResponse

	for _, bank := range b {
		result = append(result, s.ToBankResponseYearMethod(bank))
	}

	return result
}
