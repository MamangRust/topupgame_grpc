package response_api

import (
	"topup_game/internal/domain/response"
	"topup_game/internal/pb"
)

type voucherResponseMapper struct {
}

func NewVoucherResponseMapper() *voucherResponseMapper {
	return &voucherResponseMapper{}
}

func (s *voucherResponseMapper) ToApiResponseVoucher(pbResponse *pb.ApiResponseVoucher) *response.ApiResponseVoucher {
	return &response.ApiResponseVoucher{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseVoucher(pbResponse.Data),
	}
}

func (s *voucherResponseMapper) ToApiResponseVoucherAll(pbResponse *pb.ApiResponseVoucherAll) *response.ApiResponseVoucherAll {
	return &response.ApiResponseVoucherAll{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

func (s *voucherResponseMapper) ToApiResponseVoucherDelete(pbResponse *pb.ApiResponseVoucherDelete) *response.ApiResponseVoucherDelete {
	return &response.ApiResponseVoucherDelete{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

func (s *voucherResponseMapper) ToApiResponseVoucherDeleteAt(pbResponse *pb.ApiResponseVoucherDeleteAt) *response.ApiResponseVoucherDeleteAt {
	return &response.ApiResponseVoucherDeleteAt{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseVoucherDeleteAt(pbResponse.Data),
	}
}

func (s *voucherResponseMapper) ToApiResponsesVoucher(pbResponse *pb.ApiResponsesVoucher) *response.ApiResponsesVoucher {
	return &response.ApiResponsesVoucher{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesVoucher(pbResponse.Data),
	}
}

func (s *voucherResponseMapper) ToApiResponsePaginationVoucher(pbResponse *pb.ApiResponsePaginationVoucher) *response.ApiResponsePaginationVoucher {
	return &response.ApiResponsePaginationVoucher{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.mapResponsesVoucher(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
	}
}

func (s *voucherResponseMapper) ToApiResponsePaginationVoucherDeleteAt(pbResponse *pb.ApiResponsePaginationVoucherDeleteAt) *response.ApiResponsePaginationVoucherDeleteAt {
	return &response.ApiResponsePaginationVoucherDeleteAt{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.mapResponsesVoucherDeleteAt(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
	}
}

func (s *voucherResponseMapper) mapResponseVoucher(voucher *pb.VoucherResponse) *response.VoucherResponse {
	return &response.VoucherResponse{
		ID:         int(voucher.Id),
		MerchantID: int(voucher.MerchantId),
		CategoryID: int(voucher.CategoryId),
		Name:       voucher.Name,
		ImageName:  voucher.ImageName,
		CreatedAt:  voucher.CreatedAt,
		UpdatedAt:  voucher.UpdatedAt,
	}
}

func (s *voucherResponseMapper) mapResponsesVoucher(vouchers []*pb.VoucherResponse) []*response.VoucherResponse {
	var responseVouchers []*response.VoucherResponse

	for _, voucher := range vouchers {
		responseVouchers = append(responseVouchers, s.mapResponseVoucher(voucher))
	}

	return responseVouchers
}

func (s *voucherResponseMapper) mapResponseVoucherDeleteAt(voucher *pb.VoucherResponseDeleteAt) *response.VoucherResponseDeleteAt {
	return &response.VoucherResponseDeleteAt{
		ID:         int(voucher.Id),
		MerchantID: int(voucher.MerchantId),
		CategoryID: int(voucher.CategoryId),
		Name:       voucher.Name,
		ImageName:  voucher.ImageName,
		CreatedAt:  voucher.CreatedAt,
		UpdatedAt:  voucher.UpdatedAt,
		DeletedAt:  voucher.DeletedAt,
	}
}

func (s *voucherResponseMapper) mapResponsesVoucherDeleteAt(vouchers []*pb.VoucherResponseDeleteAt) []*response.VoucherResponseDeleteAt {
	var responseVouchers []*response.VoucherResponseDeleteAt

	for _, voucher := range vouchers {
		responseVouchers = append(responseVouchers, s.mapResponseVoucherDeleteAt(voucher))
	}

	return responseVouchers
}
