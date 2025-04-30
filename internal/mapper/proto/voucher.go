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

func (s *voucherProtoMapper) ToProtoResponsesMonthAmountSuccess(status string, message string, pbResponse []*response.MonthAmountVoucherSuccessResponse) *pb.ApiResponseVoucherMonthAmountSuccess {
	return &pb.ApiResponseVoucherMonthAmountSuccess{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesVoucherMonthAmountSuccess(pbResponse),
	}
}

func (s *voucherProtoMapper) ToProtoResponsesYearAmountSuccess(status string, message string, pbResponse []*response.YearAmountVoucherSuccessResponse) *pb.ApiResponseVoucherYearAmountSuccess {
	return &pb.ApiResponseVoucherYearAmountSuccess{
		Status:  status,
		Message: message,
		Data:    s.mapResponseVouchersYearAmountSuccess(pbResponse),
	}
}

func (s *voucherProtoMapper) ToProtoResponsesMonthAmountFailed(status string, message string, pbResponse []*response.MonthAmountVoucherFailedResponse) *pb.ApiResponseVoucherMonthAmountFailed {
	return &pb.ApiResponseVoucherMonthAmountFailed{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesVoucherMonthAmountFailed(pbResponse),
	}
}

func (s *voucherProtoMapper) ToProtoResponsesYearAmountFailed(status string, message string, pbResponse []*response.YearAmountVoucherFailedResponse) *pb.ApiResponseVoucherYearAmountFailed {
	return &pb.ApiResponseVoucherYearAmountFailed{
		Status:  status,
		Message: message,
		Data:    s.mapResponseVouchersYearAmountFailed(pbResponse),
	}
}

func (s *voucherProtoMapper) ToProtoResponsesMonthMethod(status string, message string, pbResponse []*response.MonthMethodVoucherResponse) *pb.ApiResponseVoucherMonthMethod {
	return &pb.ApiResponseVoucherMonthMethod{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesVoucherMonthMethod(pbResponse),
	}
}

func (s *voucherProtoMapper) ToProtoResponsesYearMethod(status string, message string, pbResponse []*response.YearMethodVoucherResponse) *pb.ApiResponseVoucherYearMethod {
	return &pb.ApiResponseVoucherYearMethod{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesVoucherYearMethod(pbResponse),
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

func (s *voucherProtoMapper) mapResponseVoucherMonthAmountSuccess(b *response.MonthAmountVoucherSuccessResponse) *pb.MonthAmountVoucherSuccessResponse {
	return &pb.MonthAmountVoucherSuccessResponse{
		Id:           int32(b.ID),
		VoucherName:  b.VoucherName,
		Year:         b.Year,
		Month:        b.Month,
		TotalSuccess: int32(b.TotalSuccess),
		TotalAmount:  int32(b.TotalAmount),
	}
}

func (s *voucherProtoMapper) mapResponsesVoucherMonthAmountSuccess(b []*response.MonthAmountVoucherSuccessResponse) []*pb.MonthAmountVoucherSuccessResponse {
	var result []*pb.MonthAmountVoucherSuccessResponse

	for _, voucher := range b {
		result = append(result, s.mapResponseVoucherMonthAmountSuccess(voucher))
	}

	return result
}

func (s *voucherProtoMapper) mapResponseVoucherYearAmountSuccess(b *response.YearAmountVoucherSuccessResponse) *pb.YearAmountVoucherSuccessResponse {
	return &pb.YearAmountVoucherSuccessResponse{
		Id:           int32(b.ID),
		VoucherName:  b.VoucherName,
		Year:         b.Year,
		TotalSuccess: int32(b.TotalSuccess),
		TotalAmount:  int32(b.TotalAmount),
	}
}

func (s *voucherProtoMapper) mapResponseVouchersYearAmountSuccess(b []*response.YearAmountVoucherSuccessResponse) []*pb.YearAmountVoucherSuccessResponse {
	var result []*pb.YearAmountVoucherSuccessResponse

	for _, voucher := range b {
		result = append(result, s.mapResponseVoucherYearAmountSuccess(voucher))
	}

	return result
}

func (s *voucherProtoMapper) mapResponseVoucherMonthAmountFailed(b *response.MonthAmountVoucherFailedResponse) *pb.MonthAmountVoucherFailedResponse {
	return &pb.MonthAmountVoucherFailedResponse{
		Id:          int32(b.ID),
		VoucherName: b.VoucherName,
		Year:        b.Year,
		Month:       b.Month,
		TotalFailed: int32(b.TotalFailed),
		TotalAmount: int32(b.TotalAmount),
	}
}

func (s *voucherProtoMapper) mapResponsesVoucherMonthAmountFailed(b []*response.MonthAmountVoucherFailedResponse) []*pb.MonthAmountVoucherFailedResponse {
	var result []*pb.MonthAmountVoucherFailedResponse

	for _, voucher := range b {
		result = append(result, s.mapResponseVoucherMonthAmountFailed(voucher))
	}

	return result
}

func (s *voucherProtoMapper) mapResponseVoucherYearAmountFailed(b *response.YearAmountVoucherFailedResponse) *pb.YearAmountVoucherFailedResponse {
	return &pb.YearAmountVoucherFailedResponse{
		Id:          int32(b.ID),
		VoucherName: b.VoucherName,
		Year:        b.Year,
		TotalFailed: int32(b.TotalFailed),
		TotalAmount: int32(b.TotalAmount),
	}
}

func (s *voucherProtoMapper) mapResponseVouchersYearAmountFailed(b []*response.YearAmountVoucherFailedResponse) []*pb.YearAmountVoucherFailedResponse {
	var result []*pb.YearAmountVoucherFailedResponse

	for _, voucher := range b {
		result = append(result, s.mapResponseVoucherYearAmountFailed(voucher))
	}

	return result
}

func (s *voucherProtoMapper) mapResponseVoucherMonthMethod(b *response.MonthMethodVoucherResponse) *pb.MonthMethodVoucherResponse {
	return &pb.MonthMethodVoucherResponse{
		Id:                int32(b.ID),
		Month:             b.Month,
		VoucherName:       b.VoucherName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int32(b.TotalAmount),
		TotalTransactions: int32(b.TotalTransactions),
	}
}

func (s *voucherProtoMapper) mapResponsesVoucherMonthMethod(b []*response.MonthMethodVoucherResponse) []*pb.MonthMethodVoucherResponse {
	var result []*pb.MonthMethodVoucherResponse

	for _, voucher := range b {
		result = append(result, s.mapResponseVoucherMonthMethod(voucher))
	}

	return result
}

func (s *voucherProtoMapper) mapResponseVoucherYearMethod(b *response.YearMethodVoucherResponse) *pb.YearMethodVoucherResponse {
	return &pb.YearMethodVoucherResponse{
		Id:                int32(b.ID),
		Year:              b.Year,
		VoucherName:       b.VoucherName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int32(b.TotalAmount),
		TotalTransactions: int32(b.TotalTransactions),
	}
}

func (s *voucherProtoMapper) mapResponsesVoucherYearMethod(b []*response.YearMethodVoucherResponse) []*pb.YearMethodVoucherResponse {
	var result []*pb.YearMethodVoucherResponse

	for _, voucher := range b {
		result = append(result, s.mapResponseVoucherYearMethod(voucher))
	}

	return result
}
