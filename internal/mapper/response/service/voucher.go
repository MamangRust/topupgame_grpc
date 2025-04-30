package response_service

import (
	"topup_game/internal/domain/record"
	"topup_game/internal/domain/response"
)

type voucherResponseMapper struct {
}

func NewVoucherResponseMapper() *voucherResponseMapper {
	return &voucherResponseMapper{}
}

func (s *voucherResponseMapper) ToVoucherResponse(voucher *record.VoucherRecord) *response.VoucherResponse {
	return &response.VoucherResponse{
		ID:         voucher.ID,
		MerchantID: voucher.MerchantID,
		CategoryID: voucher.CategoryID,
		Name:       voucher.Name,
		ImageName:  voucher.ImageName,
		CreatedAt:  voucher.CreatedAt,
		UpdatedAt:  voucher.UpdatedAt,
		DeletedAt:  voucher.DeletedAt,
	}
}

func (s *voucherResponseMapper) ToVouchersResponse(vouchers []*record.VoucherRecord) []*response.VoucherResponse {
	var responseVouchers []*response.VoucherResponse

	for _, voucher := range vouchers {
		responseVouchers = append(responseVouchers, s.ToVoucherResponse(voucher))
	}

	return responseVouchers
}

func (s *voucherResponseMapper) ToVoucherResponseDeleteAt(voucher *record.VoucherRecord) *response.VoucherResponseDeleteAt {
	return &response.VoucherResponseDeleteAt{
		ID:         voucher.ID,
		MerchantID: voucher.MerchantID,
		CategoryID: voucher.CategoryID,
		Name:       voucher.Name,
		ImageName:  voucher.ImageName,
		CreatedAt:  voucher.CreatedAt,
		UpdatedAt:  voucher.UpdatedAt,
		DeletedAt:  *voucher.DeletedAt,
	}
}

func (s *voucherResponseMapper) ToVouchersResponseDeleteAt(vouchers []*record.VoucherRecord) []*response.VoucherResponseDeleteAt {
	var responseVouchers []*response.VoucherResponseDeleteAt

	for _, voucher := range vouchers {
		responseVouchers = append(responseVouchers, s.ToVoucherResponseDeleteAt(voucher))
	}

	return responseVouchers
}

func (s *voucherResponseMapper) ToVoucherResponseMonthAmountSuccess(b *record.MonthAmountVoucherSuccessRecord) *response.MonthAmountVoucherSuccessResponse {
	return &response.MonthAmountVoucherSuccessResponse{
		ID:           int(b.ID),
		VoucherName:  b.VoucherName,
		Year:         b.Year,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *voucherResponseMapper) ToVouchersResponseMonthAmountSuccess(b []*record.MonthAmountVoucherSuccessRecord) []*response.MonthAmountVoucherSuccessResponse {
	var result []*response.MonthAmountVoucherSuccessResponse

	for _, voucher := range b {
		result = append(result, s.ToVoucherResponseMonthAmountSuccess(voucher))
	}

	return result
}

func (s *voucherResponseMapper) ToVoucherResponseYearAmountSuccess(b *record.YearAmountVoucherSuccessRecord) *response.YearAmountVoucherSuccessResponse {
	return &response.YearAmountVoucherSuccessResponse{
		ID:           int(b.ID),
		VoucherName:  b.VoucherName,
		Year:         b.Year,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *voucherResponseMapper) ToVouchersResponseYearAmountSuccess(b []*record.YearAmountVoucherSuccessRecord) []*response.YearAmountVoucherSuccessResponse {
	var result []*response.YearAmountVoucherSuccessResponse

	for _, voucher := range b {
		result = append(result, s.ToVoucherResponseYearAmountSuccess(voucher))
	}

	return result
}

func (s *voucherResponseMapper) ToVoucherResponseMonthAmountFailed(b *record.MonthAmountVoucherFailedRecord) *response.MonthAmountVoucherFailedResponse {
	return &response.MonthAmountVoucherFailedResponse{
		ID:          int(b.ID),
		VoucherName: b.VoucherName,
		Year:        b.Year,
		Month:       b.Month,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *voucherResponseMapper) ToVouchersResponseMonthAmountFailed(b []*record.MonthAmountVoucherFailedRecord) []*response.MonthAmountVoucherFailedResponse {
	var result []*response.MonthAmountVoucherFailedResponse

	for _, voucher := range b {
		result = append(result, s.ToVoucherResponseMonthAmountFailed(voucher))
	}

	return result
}

func (s *voucherResponseMapper) ToVoucherResponseYearAmountFailed(b *record.YearAmountVoucherFailedRecord) *response.YearAmountVoucherFailedResponse {
	return &response.YearAmountVoucherFailedResponse{
		ID:          int(b.ID),
		VoucherName: b.VoucherName,
		Year:        b.Year,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *voucherResponseMapper) ToVouchersResponseYearAmountFailed(b []*record.YearAmountVoucherFailedRecord) []*response.YearAmountVoucherFailedResponse {
	var result []*response.YearAmountVoucherFailedResponse

	for _, voucher := range b {
		result = append(result, s.ToVoucherResponseYearAmountFailed(voucher))
	}

	return result
}

func (s *voucherResponseMapper) ToVoucherResponseMonthMethod(b *record.MonthMethodVoucherRecord) *response.MonthMethodVoucherResponse {
	return &response.MonthMethodVoucherResponse{
		ID:                int(b.ID),
		Month:             b.Month,
		VoucherName:       b.VoucherName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *voucherResponseMapper) ToVouchersResponseMonthMethod(b []*record.MonthMethodVoucherRecord) []*response.MonthMethodVoucherResponse {
	var result []*response.MonthMethodVoucherResponse

	for _, voucher := range b {
		result = append(result, s.ToVoucherResponseMonthMethod(voucher))
	}

	return result
}

func (s *voucherResponseMapper) ToVoucherResponseYearMethod(b *record.YearMethodVoucherRecord) *response.YearMethodVoucherResponse {
	return &response.YearMethodVoucherResponse{
		ID:                int(b.ID),
		Year:              b.Year,
		VoucherName:       b.VoucherName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *voucherResponseMapper) ToVouchersResponseYearMethod(b []*record.YearMethodVoucherRecord) []*response.YearMethodVoucherResponse {
	var result []*response.YearMethodVoucherResponse

	for _, voucher := range b {
		result = append(result, s.ToVoucherResponseYearMethod(voucher))
	}

	return result
}
