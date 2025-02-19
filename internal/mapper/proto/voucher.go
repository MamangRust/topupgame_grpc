package protomapper

import (
	"topup_game/internal/domain/response"
	"topup_game/internal/pb"
)

type voucherProtoMapper struct{}

func NewVoucherProtoMapper() *voucherProtoMapper {
	return &voucherProtoMapper{}
}

func (s *voucherProtoMapper) ToProtoResponseVoucherAll(status string, message string) *pb.ApiResponseVoucherAll {
	return &pb.ApiResponseVoucherAll{
		Status:  status,
		Message: message,
	}
}

func (s *voucherProtoMapper) ToProtoResponseVoucherDelete(status string, message string) *pb.ApiResponseVoucherDelete {
	return &pb.ApiResponseVoucherDelete{
		Status:  status,
		Message: message,
	}
}

func (s *voucherProtoMapper) ToProtoResponseVoucher(status string, message string, pbResponse *response.VoucherResponse) *pb.ApiResponseVoucher {
	return &pb.ApiResponseVoucher{
		Status:  status,
		Message: message,
		Data:    s.mapResponseVoucher(pbResponse),
	}
}

func (s *voucherProtoMapper) ToProtoResponseVoucherDeleteAt(status string, message string, pbResponse *response.VoucherResponseDeleteAt) *pb.ApiResponseVoucherDeleteAt {
	return &pb.ApiResponseVoucherDeleteAt{
		Status:  status,
		Message: message,
		Data:    s.mapResponseVoucherDeleteAt(pbResponse),
	}
}

func (s *voucherProtoMapper) ToProtoResponsesVoucher(status string, message string, pbResponse []*response.VoucherResponse) *pb.ApiResponsesVoucher {
	return &pb.ApiResponsesVoucher{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesVoucher(pbResponse),
	}
}

func (s *voucherProtoMapper) ToProtoResponsePaginationVoucher(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.VoucherResponse) *pb.ApiResponsePaginationVoucher {
	return &pb.ApiResponsePaginationVoucher{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesVoucher(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *voucherProtoMapper) ToProtoResponsePaginationVoucherDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.VoucherResponseDeleteAt) *pb.ApiResponsePaginationVoucherDeleteAt {
	return &pb.ApiResponsePaginationVoucherDeleteAt{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesVoucherDeleteAt(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *voucherProtoMapper) mapResponseVoucher(voucher *response.VoucherResponse) *pb.VoucherResponse {
	return &pb.VoucherResponse{
		Id:         int32(voucher.ID),
		MerchantId: int32(voucher.MerchantID),
		CategoryId: int32(voucher.CategoryID),
		Name:       voucher.Name,
		ImageName:  voucher.ImageName,
		CreatedAt:  voucher.CreatedAt,
		UpdatedAt:  voucher.UpdatedAt,
	}
}

func (s *voucherProtoMapper) mapResponsesVoucher(vouchers []*response.VoucherResponse) []*pb.VoucherResponse {
	var responseVouchers []*pb.VoucherResponse

	for _, voucher := range vouchers {
		responseVouchers = append(responseVouchers, s.mapResponseVoucher(voucher))
	}

	return responseVouchers
}

func (s *voucherProtoMapper) mapResponseVoucherDeleteAt(voucher *response.VoucherResponseDeleteAt) *pb.VoucherResponseDeleteAt {
	return &pb.VoucherResponseDeleteAt{
		Id:         int32(voucher.ID),
		MerchantId: int32(voucher.MerchantID),
		CategoryId: int32(voucher.CategoryID),
		Name:       voucher.Name,
		ImageName:  voucher.ImageName,
		CreatedAt:  voucher.CreatedAt,
		UpdatedAt:  voucher.UpdatedAt,
		DeletedAt:  voucher.DeletedAt,
	}
}

func (s *voucherProtoMapper) mapResponsesVoucherDeleteAt(vouchers []*response.VoucherResponseDeleteAt) []*pb.VoucherResponseDeleteAt {
	var responseVouchers []*pb.VoucherResponseDeleteAt

	for _, voucher := range vouchers {
		responseVouchers = append(responseVouchers, s.mapResponseVoucherDeleteAt(voucher))
	}

	return responseVouchers
}
