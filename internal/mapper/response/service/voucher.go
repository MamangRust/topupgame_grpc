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
