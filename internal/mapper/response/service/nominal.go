package response_service

import (
	"topup_game/internal/domain/record"
	"topup_game/internal/domain/response"
)

type nominalResponseMapper struct {
}

func NewNominalResponseMapper() *nominalResponseMapper {
	return &nominalResponseMapper{}
}

func (s *nominalResponseMapper) ToNominalResponse(nominal *record.NominalRecord) *response.NominalResponse {
	return &response.NominalResponse{
		ID:        nominal.ID,
		VoucherID: nominal.VoucherID,
		Name:      nominal.Name,
		Quantity:  nominal.Quantity,
		Price:     nominal.Price,
		CreatedAt: nominal.CreatedAt,
		UpdatedAt: nominal.UpdatedAt,
		DeletedAt: nominal.DeletedAt,
	}
}

func (s *nominalResponseMapper) ToNominalsResponse(nominals []*record.NominalRecord) []*response.NominalResponse {
	var responseNominals []*response.NominalResponse

	for _, nominal := range nominals {
		responseNominals = append(responseNominals, s.ToNominalResponse(nominal))
	}

	return responseNominals
}

func (s *nominalResponseMapper) ToNominalResponseDeleteAt(nominal *record.NominalRecord) *response.NominalResponseDeleteAt {
	return &response.NominalResponseDeleteAt{
		ID:        nominal.ID,
		VoucherID: nominal.VoucherID,
		Name:      nominal.Name,
		Quantity:  nominal.Quantity,
		Price:     nominal.Price,
		CreatedAt: nominal.CreatedAt,
		UpdatedAt: nominal.UpdatedAt,
		DeletedAt: *nominal.DeletedAt,
	}
}

func (s *nominalResponseMapper) ToNominalsResponseDeleteAt(nominals []*record.NominalRecord) []*response.NominalResponseDeleteAt {
	var responseNominals []*response.NominalResponseDeleteAt

	for _, nominal := range nominals {
		responseNominals = append(responseNominals, s.ToNominalResponseDeleteAt(nominal))
	}

	return responseNominals
}

func (s *nominalResponseMapper) ToNominalResponseMonthAmountSuccess(b *record.MonthAmountNominalSuccessRecord) *response.MonthAmountNominalSuccessResponse {
	return &response.MonthAmountNominalSuccessResponse{
		ID:           int(b.ID),
		NominalName:  b.NominalName,
		Year:         b.Year,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *nominalResponseMapper) ToNominalsResponseMonthAmountSuccess(b []*record.MonthAmountNominalSuccessRecord) []*response.MonthAmountNominalSuccessResponse {
	var result []*response.MonthAmountNominalSuccessResponse

	for _, nominal := range b {
		result = append(result, s.ToNominalResponseMonthAmountSuccess(nominal))
	}

	return result
}

func (s *nominalResponseMapper) ToNominalResponseYearAmountSuccess(b *record.YearAmountNominalSuccessRecord) *response.YearAmountNominalSuccessResponse {
	return &response.YearAmountNominalSuccessResponse{
		ID:           int(b.ID),
		NominalName:  b.NominalName,
		Year:         b.Year,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *nominalResponseMapper) ToNominalsResponseYearAmountSuccess(b []*record.YearAmountNominalSuccessRecord) []*response.YearAmountNominalSuccessResponse {
	var result []*response.YearAmountNominalSuccessResponse

	for _, nominal := range b {
		result = append(result, s.ToNominalResponseYearAmountSuccess(nominal))
	}

	return result
}

func (s *nominalResponseMapper) ToNominalResponseMonthAmountFailed(b *record.MonthAmountNominalFailedRecord) *response.MonthAmountNominalFailedResponse {
	return &response.MonthAmountNominalFailedResponse{
		ID:          int(b.ID),
		NominalName: b.NominalName,
		Year:        b.Year,
		Month:       b.Month,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *nominalResponseMapper) ToNominalsResponseMonthAmountFailed(b []*record.MonthAmountNominalFailedRecord) []*response.MonthAmountNominalFailedResponse {
	var result []*response.MonthAmountNominalFailedResponse

	for _, Bank := range b {
		result = append(result, s.ToNominalResponseMonthAmountFailed(Bank))
	}

	return result
}

func (s *nominalResponseMapper) ToNominalResponseYearAmountFailed(b *record.YearAmountNominalFailedRecord) *response.YearAmountNominalFailedResponse {
	return &response.YearAmountNominalFailedResponse{
		ID:          int(b.ID),
		NominalName: b.NominalName,
		Year:        b.Year,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *nominalResponseMapper) ToNominalsResponseYearAmountFailed(b []*record.YearAmountNominalFailedRecord) []*response.YearAmountNominalFailedResponse {
	var result []*response.YearAmountNominalFailedResponse

	for _, Bank := range b {
		result = append(result, s.ToNominalResponseYearAmountFailed(Bank))
	}

	return result
}

func (s *nominalResponseMapper) ToNominalResponseMonthMethodSuccess(b *record.MonthMethodNominalRecord) *response.MonthMethodNominalResponse {
	return &response.MonthMethodNominalResponse{
		ID:                int(b.ID),
		Month:             b.Month,
		NominalName:       b.NominalName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s nominalResponseMapper) ToNominalsResponseMonthMethodSuccess(b []*record.MonthMethodNominalRecord) []*response.MonthMethodNominalResponse {
	var result []*response.MonthMethodNominalResponse

	for _, nominal := range b {
		result = append(result, s.ToNominalResponseMonthMethodSuccess(nominal))
	}

	return result
}

func (s *nominalResponseMapper) ToNominalResponseMonthMethodFailed(b *record.MonthMethodNominalRecord) *response.MonthMethodNominalResponse {
	return &response.MonthMethodNominalResponse{
		ID:                int(b.ID),
		Month:             b.Month,
		NominalName:       b.NominalName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s nominalResponseMapper) ToNominalsResponseMonthMethodFailed(b []*record.MonthMethodNominalRecord) []*response.MonthMethodNominalResponse {
	var result []*response.MonthMethodNominalResponse

	for _, nominal := range b {
		result = append(result, s.ToNominalResponseMonthMethodFailed(nominal))
	}

	return result
}

func (s *nominalResponseMapper) ToNominalResponseYearMethod(b *record.YearMethodNominalRecord) *response.YearMethodNominalResponse {
	return &response.YearMethodNominalResponse{
		ID:                int(b.ID),
		Year:              b.Year,
		NominalName:       b.NominalName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s nominalResponseMapper) ToNominalsResponseYearMethod(b []*record.YearMethodNominalRecord) []*response.YearMethodNominalResponse {
	var result []*response.YearMethodNominalResponse

	for _, nominal := range b {
		result = append(result, s.ToNominalResponseYearMethod(nominal))
	}

	return result
}
