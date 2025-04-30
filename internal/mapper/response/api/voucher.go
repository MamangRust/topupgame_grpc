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

func (s *voucherResponseMapper) ToApiResponsesMonthAmountSuccess(pbResponse *pb.ApiResponseVoucherMonthAmountSuccess) *response.ApiResponsesVoucherMonthSuccess {
	return &response.ApiResponsesVoucherMonthSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesVoucherMonthAmountSuccess(pbResponse.Data),
	}
}

func (s *voucherResponseMapper) ToApiResponseYearAmountSuccess(pbResponse *pb.ApiResponseVoucherYearAmountSuccess) *response.ApiResponsesVoucherYearSuccess {
	return &response.ApiResponsesVoucherYearSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseVouchersYearAmountSuccess(pbResponse.Data),
	}
}

func (s *voucherResponseMapper) ToApiResponsesMonthAmountFailed(pbResponse *pb.ApiResponseVoucherMonthAmountFailed) *response.ApiResponsesVoucherMonthFailed {
	return &response.ApiResponsesVoucherMonthFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesVoucherMonthAmountFailed(pbResponse.Data),
	}
}

func (s *voucherResponseMapper) ToApiResponseYearAmountFailed(pbResponse *pb.ApiResponseVoucherYearAmountFailed) *response.ApiResponsesVoucherYearFailed {
	return &response.ApiResponsesVoucherYearFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseVouchersYearAmountFailed(pbResponse.Data),
	}
}

func (s *voucherResponseMapper) ToApiResponsesMonthMethod(pbResponse *pb.ApiResponseVoucherMonthMethod) *response.ApiResponsesVoucherMonthMethod {
	return &response.ApiResponsesVoucherMonthMethod{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesVoucherMonthMethod(pbResponse.Data),
	}
}

func (s *voucherResponseMapper) ToApiResponseYearMethod(pbResponse *pb.ApiResponseVoucherYearMethod) *response.ApiResponsesVoucherYearMethod {
	return &response.ApiResponsesVoucherYearMethod{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesVoucherYearMethod(pbResponse.Data),
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

func (s *voucherResponseMapper) mapResponseVoucherMonthAmountSuccess(b *pb.MonthAmountVoucherSuccessResponse) *response.MonthAmountVoucherSuccessResponse {
	return &response.MonthAmountVoucherSuccessResponse{
		ID:           int(b.Id),
		VoucherName:  b.VoucherName,
		Year:         b.Year,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *voucherResponseMapper) mapResponsesVoucherMonthAmountSuccess(b []*pb.MonthAmountVoucherSuccessResponse) []*response.MonthAmountVoucherSuccessResponse {
	var result []*response.MonthAmountVoucherSuccessResponse

	for _, voucher := range b {
		result = append(result, s.mapResponseVoucherMonthAmountSuccess(voucher))
	}

	return result
}

func (s *voucherResponseMapper) mapResponseVoucherYearAmountSuccess(b *pb.YearAmountVoucherSuccessResponse) *response.YearAmountVoucherSuccessResponse {
	return &response.YearAmountVoucherSuccessResponse{
		ID:           int(b.Id),
		VoucherName:  b.VoucherName,
		Year:         b.Year,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *voucherResponseMapper) mapResponseVouchersYearAmountSuccess(b []*pb.YearAmountVoucherSuccessResponse) []*response.YearAmountVoucherSuccessResponse {
	var result []*response.YearAmountVoucherSuccessResponse

	for _, voucher := range b {
		result = append(result, s.mapResponseVoucherYearAmountSuccess(voucher))
	}

	return result
}

func (s *voucherResponseMapper) mapResponseVoucherMonthAmountFailed(b *pb.MonthAmountVoucherFailedResponse) *response.MonthAmountVoucherFailedResponse {
	return &response.MonthAmountVoucherFailedResponse{
		ID:          int(b.Id),
		VoucherName: b.VoucherName,
		Year:        b.Year,
		Month:       b.Month,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *voucherResponseMapper) mapResponsesVoucherMonthAmountFailed(b []*pb.MonthAmountVoucherFailedResponse) []*response.MonthAmountVoucherFailedResponse {
	var result []*response.MonthAmountVoucherFailedResponse

	for _, voucher := range b {
		result = append(result, s.mapResponseVoucherMonthAmountFailed(voucher))
	}

	return result
}

func (s *voucherResponseMapper) mapResponseVoucherYearAmountFailed(b *pb.YearAmountVoucherFailedResponse) *response.YearAmountVoucherFailedResponse {
	return &response.YearAmountVoucherFailedResponse{
		ID:          int(b.Id),
		VoucherName: b.VoucherName,
		Year:        b.Year,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *voucherResponseMapper) mapResponseVouchersYearAmountFailed(b []*pb.YearAmountVoucherFailedResponse) []*response.YearAmountVoucherFailedResponse {
	var result []*response.YearAmountVoucherFailedResponse

	for _, voucher := range b {
		result = append(result, s.mapResponseVoucherYearAmountFailed(voucher))
	}

	return result
}

func (s *voucherResponseMapper) mapResponseVoucherMonthMethod(b *pb.MonthMethodVoucherResponse) *response.MonthMethodVoucherResponse {
	return &response.MonthMethodVoucherResponse{
		ID:                int(b.Id),
		Month:             b.Month,
		VoucherName:       b.VoucherName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *voucherResponseMapper) mapResponsesVoucherMonthMethod(b []*pb.MonthMethodVoucherResponse) []*response.MonthMethodVoucherResponse {
	var result []*response.MonthMethodVoucherResponse

	for _, voucher := range b {
		result = append(result, s.mapResponseVoucherMonthMethod(voucher))
	}

	return result
}

func (s *voucherResponseMapper) mapResponseVoucherYearMethod(b *pb.YearMethodVoucherResponse) *response.YearMethodVoucherResponse {
	return &response.YearMethodVoucherResponse{
		ID:                int(b.Id),
		Year:              b.Year,
		VoucherName:       b.VoucherName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *voucherResponseMapper) mapResponsesVoucherYearMethod(b []*pb.YearMethodVoucherResponse) []*response.YearMethodVoucherResponse {
	var result []*response.YearMethodVoucherResponse

	for _, voucher := range b {
		result = append(result, s.mapResponseVoucherYearMethod(voucher))
	}

	return result
}
