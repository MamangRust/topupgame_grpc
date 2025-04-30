package response_api

import (
	"topup_game/internal/domain/response"
	"topup_game/internal/pb"
)

type transactionResponseMapper struct {
}

func NewTransactionResponseMapper() *transactionResponseMapper {
	return &transactionResponseMapper{}
}

func (s *transactionResponseMapper) ToApiResponseTransaction(pbResponse *pb.ApiResponseTransaction) *response.ApiResponseTransaction {
	return &response.ApiResponseTransaction{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseTransaction(pbResponse.Data),
	}
}

func (s *transactionResponseMapper) ToApiResponseTransactionDelete(pbResponse *pb.ApiResponseTransactionDelete) *response.ApiResponseTransactionDelete {
	return &response.ApiResponseTransactionDelete{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

func (s *transactionResponseMapper) ToApiResponseTransactionAll(pbResponse *pb.ApiResponseTransactionAll) *response.ApiResponseTransactionAll {
	return &response.ApiResponseTransactionAll{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

func (s *transactionResponseMapper) ToApiResponseTransactionDeleteAt(pbResponse *pb.ApiResponseTransactionDeleteAt) *response.ApiResponseTransactionDeleteAt {
	return &response.ApiResponseTransactionDeleteAt{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseTransactionDeleteAt(pbResponse.Data),
	}
}

func (s *transactionResponseMapper) ToApiResponsesTransaction(pbResponse *pb.ApiResponsesTransaction) *response.ApiResponsesTransaction {
	return &response.ApiResponsesTransaction{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesTransaction(pbResponse.Data),
	}
}

func (s *transactionResponseMapper) ToApiResponsePaginationTransaction(pbResponse *pb.ApiResponsePaginationTransaction) *response.ApiResponsePaginationTransaction {
	return &response.ApiResponsePaginationTransaction{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.mapResponsesTransaction(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
	}
}

func (s *transactionResponseMapper) ToApiResponsePaginationTransactionDeleteAt(pbResponse *pb.ApiResponsePaginationTransactionDeleteAt) *response.ApiResponsePaginationTransactionDeleteAt {
	return &response.ApiResponsePaginationTransactionDeleteAt{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.mapResponsesTransactionDeleteAt(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
	}
}

func (s *transactionResponseMapper) ToApiResponsesMonthAmountSuccess(pbResponse *pb.ApiResponseTransactionMonthAmountSuccess) *response.ApiResponsesTransactionMonthSuccess {
	return &response.ApiResponsesTransactionMonthSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesMonthAmountSuccess(pbResponse.Data),
	}
}

func (s *transactionResponseMapper) ToApiResponseYearAmountSuccess(pbResponse *pb.ApiResponseTransactionYearAmountSuccess) *response.ApiResponsesTransactionYearSuccess {
	return &response.ApiResponsesTransactionYearSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesYearAmountSuccess(pbResponse.Data),
	}
}

func (s *transactionResponseMapper) ToApiResponsesMonthAmountFailed(pbResponse *pb.ApiResponseTransactionMonthAmountFailed) *response.ApiResponsesTransactionMonthFailed {
	return &response.ApiResponsesTransactionMonthFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesMonthAmountFailed(pbResponse.Data),
	}
}

func (s *transactionResponseMapper) ToApiResponseYearAmountFailed(pbResponse *pb.ApiResponseTransactionYearAmountFailed) *response.ApiResponsesTransactionYearFailed {
	return &response.ApiResponsesTransactionYearFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesYearAmountFailed(pbResponse.Data),
	}
}

func (s *transactionResponseMapper) ToApiResponsesMonthMethod(pbResponse *pb.ApiResponseTransactionMonthMethod) *response.ApiResponsesTransactionMonthMethod {
	return &response.ApiResponsesTransactionMonthMethod{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesMonthMethod(pbResponse.Data),
	}
}

func (s *transactionResponseMapper) ToApiResponseYearMethod(pbResponse *pb.ApiResponseTransactionYearMethod) *response.ApiResponsesTransactionYearMethod {
	return &response.ApiResponsesTransactionYearMethod{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesYearMethod(pbResponse.Data),
	}
}

func (s *transactionResponseMapper) mapResponseTransaction(tx *pb.TransactionResponse) *response.TransactionResponse {
	return &response.TransactionResponse{
		ID:            int(tx.Id),
		UserID:        int(tx.UserId),
		MerchantID:    int(tx.MerchantId),
		VoucherID:     int(tx.VoucherId),
		BankID:        int(tx.BankId),
		PaymentMethod: tx.PaymentMethod,
		Status:        tx.Status,
		CreatedAt:     tx.CreatedAt,
		UpdatedAt:     tx.UpdatedAt,
	}
}

func (s *transactionResponseMapper) mapResponsesTransaction(txs []*pb.TransactionResponse) []*response.TransactionResponse {
	var responseTxs []*response.TransactionResponse

	for _, tx := range txs {
		responseTxs = append(responseTxs, s.mapResponseTransaction(tx))
	}

	return responseTxs
}

func (s *transactionResponseMapper) mapResponseTransactionDeleteAt(tx *pb.TransactionResponseDeleteAt) *response.TransactionResponseDeleteAt {
	return &response.TransactionResponseDeleteAt{
		ID:            int(tx.Id),
		UserID:        int(tx.UserId),
		MerchantID:    int(tx.MerchantId),
		VoucherID:     int(tx.VoucherId),
		BankID:        int(tx.BankId),
		PaymentMethod: tx.PaymentMethod,
		Status:        tx.Status,
		CreatedAt:     tx.CreatedAt,
		UpdatedAt:     tx.UpdatedAt,
		DeletedAt:     tx.DeletedAt,
	}
}

func (s *transactionResponseMapper) mapResponsesTransactionDeleteAt(txs []*pb.TransactionResponseDeleteAt) []*response.TransactionResponseDeleteAt {
	var responseTxs []*response.TransactionResponseDeleteAt

	for _, tx := range txs {
		responseTxs = append(responseTxs, s.mapResponseTransactionDeleteAt(tx))
	}

	return responseTxs
}

func (s *transactionResponseMapper) mapResponseMonthAmountSuccess(b *pb.MonthAmountTransactionSuccessResponse) *response.MonthAmountTransactionSuccessResponse {
	return &response.MonthAmountTransactionSuccessResponse{
		Year:         b.Year,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *transactionResponseMapper) mapResponsesMonthAmountSuccess(b []*pb.MonthAmountTransactionSuccessResponse) []*response.MonthAmountTransactionSuccessResponse {
	var result []*response.MonthAmountTransactionSuccessResponse

	for _, transaction := range b {
		result = append(result, s.mapResponseMonthAmountSuccess(transaction))
	}

	return result
}

func (s *transactionResponseMapper) mapResponseYearAmountSuccess(b *pb.YearAmountTransactionSuccessResponse) *response.YearAmountTransactionSuccessResponse {
	return &response.YearAmountTransactionSuccessResponse{
		Year:         b.Year,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *transactionResponseMapper) mapResponsesYearAmountSuccess(b []*pb.YearAmountTransactionSuccessResponse) []*response.YearAmountTransactionSuccessResponse {
	var result []*response.YearAmountTransactionSuccessResponse

	for _, transaction := range b {
		result = append(result, s.mapResponseYearAmountSuccess(transaction))
	}

	return result
}

func (s *transactionResponseMapper) mapResponseMonthAmountFailed(b *pb.MonthAmountTransactionFailedResponse) *response.MonthAmountTransactionFailedResponse {
	return &response.MonthAmountTransactionFailedResponse{
		Year:        b.Year,
		Month:       b.Month,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *transactionResponseMapper) mapResponsesMonthAmountFailed(b []*pb.MonthAmountTransactionFailedResponse) []*response.MonthAmountTransactionFailedResponse {
	var result []*response.MonthAmountTransactionFailedResponse

	for _, transaction := range b {
		result = append(result, s.mapResponseMonthAmountFailed(transaction))
	}

	return result
}

func (s *transactionResponseMapper) mapResponseYearAmountFailed(b *pb.YearAmountTransactionFailedResponse) *response.YearAmountTransactionFailedResponse {
	return &response.YearAmountTransactionFailedResponse{
		Year:        b.Year,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *transactionResponseMapper) mapResponsesYearAmountFailed(b []*pb.YearAmountTransactionFailedResponse) []*response.YearAmountTransactionFailedResponse {
	var result []*response.YearAmountTransactionFailedResponse

	for _, transaction := range b {
		result = append(result, s.mapResponseYearAmountFailed(transaction))
	}

	return result
}

func (s *transactionResponseMapper) mapResponseMonthMethod(b *pb.MonthMethodTransactionResponse) *response.MonthMethodTransactionResponse {
	return &response.MonthMethodTransactionResponse{
		Month:             b.Month,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *transactionResponseMapper) mapResponsesMonthMethod(b []*pb.MonthMethodTransactionResponse) []*response.MonthMethodTransactionResponse {
	var result []*response.MonthMethodTransactionResponse

	for _, transaction := range b {
		result = append(result, s.mapResponseMonthMethod(transaction))
	}

	return result
}

func (s *transactionResponseMapper) mapResponseYearMethod(b *pb.YearMethodTransactionResponse) *response.YearMethodTransactionResponse {
	return &response.YearMethodTransactionResponse{
		Year:              b.Year,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *transactionResponseMapper) mapResponsesYearMethod(b []*pb.YearMethodTransactionResponse) []*response.YearMethodTransactionResponse {
	var result []*response.YearMethodTransactionResponse

	for _, transaction := range b {
		result = append(result, s.mapResponseYearMethod(transaction))
	}

	return result
}
