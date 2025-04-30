package protomapper

import (
	"topup_game/internal/domain/response"
	"topup_game/internal/pb"
)

type transactionProtoMapper struct{}

func NewTransactionProtoMapper() *transactionProtoMapper {
	return &transactionProtoMapper{}
}

func (s *transactionProtoMapper) ToProtoResponseTransactionAll(status string, message string) *pb.ApiResponseTransactionAll {
	return &pb.ApiResponseTransactionAll{
		Status:  status,
		Message: message,
	}
}

func (s *transactionProtoMapper) ToProtoResponseTransactionDelete(status string, message string) *pb.ApiResponseTransactionDelete {
	return &pb.ApiResponseTransactionDelete{
		Status:  status,
		Message: message,
	}
}

func (s *transactionProtoMapper) ToProtoResponseTransaction(status string, message string, pbResponse *response.TransactionResponse) *pb.ApiResponseTransaction {
	return &pb.ApiResponseTransaction{
		Status:  status,
		Message: message,
		Data:    s.mapResponseTransaction(pbResponse),
	}
}

func (s *transactionProtoMapper) ToProtoResponseTransactionDeleteAt(status string, message string, pbResponse *response.TransactionResponseDeleteAt) *pb.ApiResponseTransactionDeleteAt {
	return &pb.ApiResponseTransactionDeleteAt{
		Status:  status,
		Message: message,
		Data:    s.mapResponseTransactionDeleteAt(pbResponse),
	}
}

func (s *transactionProtoMapper) ToProtoResponsesTransaction(status string, message string, pbResponse []*response.TransactionResponse) *pb.ApiResponsesTransaction {
	return &pb.ApiResponsesTransaction{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesTransaction(pbResponse),
	}
}

func (s *transactionProtoMapper) ToProtoResponsePaginationTransaction(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.TransactionResponse) *pb.ApiResponsePaginationTransaction {
	return &pb.ApiResponsePaginationTransaction{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesTransaction(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *transactionProtoMapper) ToProtoResponsePaginationTransactionDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.TransactionResponseDeleteAt) *pb.ApiResponsePaginationTransactionDeleteAt {
	return &pb.ApiResponsePaginationTransactionDeleteAt{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesTransactionDeleteAt(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *transactionProtoMapper) ToProtoResponsesMonthAmountSuccess(status string, message string, pbResponse []*response.MonthAmountTransactionSuccessResponse) *pb.ApiResponseTransactionMonthAmountSuccess {
	return &pb.ApiResponseTransactionMonthAmountSuccess{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesMonthAmountSuccess(pbResponse),
	}
}

func (s *transactionProtoMapper) ToProtoResponseYearAmountSuccess(status string, message string, pbResponse []*response.YearAmountTransactionSuccessResponse) *pb.ApiResponseTransactionYearAmountSuccess {
	return &pb.ApiResponseTransactionYearAmountSuccess{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesYearAmountSuccess(pbResponse),
	}
}

func (s *transactionProtoMapper) ToProtoResponsesMonthAmountFailed(status string, message string, pbResponse []*response.MonthAmountTransactionFailedResponse) *pb.ApiResponseTransactionMonthAmountFailed {
	return &pb.ApiResponseTransactionMonthAmountFailed{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesMonthAmountFailed(pbResponse),
	}
}

func (s *transactionProtoMapper) ToProtoResponseYearAmountFailed(status string, message string, pbResponse []*response.YearAmountTransactionFailedResponse) *pb.ApiResponseTransactionYearAmountFailed {
	return &pb.ApiResponseTransactionYearAmountFailed{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesYearAmountFailed(pbResponse),
	}
}

func (s *transactionProtoMapper) ToProtoResponsesMonthMethod(status string, message string, pbResponse []*response.MonthMethodTransactionResponse) *pb.ApiResponseTransactionMonthMethod {
	return &pb.ApiResponseTransactionMonthMethod{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesMonthMethod(pbResponse),
	}
}

func (s *transactionProtoMapper) ToProtoResponseYearMethod(status string, message string, pbResponse []*response.YearMethodTransactionResponse) *pb.ApiResponseTransactionYearMethod {
	return &pb.ApiResponseTransactionYearMethod{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesYearMethod(pbResponse),
	}
}

func (s *transactionProtoMapper) mapResponseTransaction(transaction *response.TransactionResponse) *pb.TransactionResponse {
	return &pb.TransactionResponse{
		Id:            int32(transaction.ID),
		UserId:        int32(transaction.UserID),
		MerchantId:    int32(transaction.MerchantID),
		VoucherId:     int32(transaction.VoucherID),
		NominalId:     int32(transaction.NominalID),
		BankId:        int32(transaction.BankID),
		PaymentMethod: transaction.PaymentMethod,
		Status:        transaction.Status,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	}
}

func (s *transactionProtoMapper) mapResponsesTransaction(transactions []*response.TransactionResponse) []*pb.TransactionResponse {
	var responseTransactions []*pb.TransactionResponse

	for _, transaction := range transactions {
		responseTransactions = append(responseTransactions, s.mapResponseTransaction(transaction))
	}

	return responseTransactions
}

func (s *transactionProtoMapper) mapResponseTransactionDeleteAt(transaction *response.TransactionResponseDeleteAt) *pb.TransactionResponseDeleteAt {
	return &pb.TransactionResponseDeleteAt{
		Id:            int32(transaction.ID),
		UserId:        int32(transaction.UserID),
		MerchantId:    int32(transaction.MerchantID),
		VoucherId:     int32(transaction.VoucherID),
		NominalId:     int32(transaction.NominalID),
		BankId:        int32(transaction.BankID),
		PaymentMethod: transaction.PaymentMethod,
		Status:        transaction.Status,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
		DeletedAt:     transaction.DeletedAt,
	}
}

func (s *transactionProtoMapper) mapResponsesTransactionDeleteAt(transactions []*response.TransactionResponseDeleteAt) []*pb.TransactionResponseDeleteAt {
	var responseTransactions []*pb.TransactionResponseDeleteAt

	for _, transaction := range transactions {
		responseTransactions = append(responseTransactions, s.mapResponseTransactionDeleteAt(transaction))
	}

	return responseTransactions
}

func (s *transactionProtoMapper) mapResponseMonthAmountSuccess(b *response.MonthAmountTransactionSuccessResponse) *pb.MonthAmountTransactionSuccessResponse {
	return &pb.MonthAmountTransactionSuccessResponse{
		Year:         b.Year,
		Month:        b.Month,
		TotalSuccess: int32(b.TotalSuccess),
		TotalAmount:  int32(b.TotalAmount),
	}
}

func (s *transactionProtoMapper) mapResponsesMonthAmountSuccess(b []*response.MonthAmountTransactionSuccessResponse) []*pb.MonthAmountTransactionSuccessResponse {
	var result []*pb.MonthAmountTransactionSuccessResponse

	for _, transaction := range b {
		result = append(result, s.mapResponseMonthAmountSuccess(transaction))
	}

	return result
}

func (s *transactionProtoMapper) mapResponseYearAmountSuccess(b *response.YearAmountTransactionSuccessResponse) *pb.YearAmountTransactionSuccessResponse {
	return &pb.YearAmountTransactionSuccessResponse{
		Year:         b.Year,
		TotalSuccess: int32(b.TotalSuccess),
		TotalAmount:  int32(b.TotalAmount),
	}
}

func (s *transactionProtoMapper) mapResponsesYearAmountSuccess(b []*response.YearAmountTransactionSuccessResponse) []*pb.YearAmountTransactionSuccessResponse {
	var result []*pb.YearAmountTransactionSuccessResponse

	for _, transaction := range b {
		result = append(result, s.mapResponseYearAmountSuccess(transaction))
	}

	return result
}

func (s *transactionProtoMapper) mapResponseMonthAmountFailed(b *response.MonthAmountTransactionFailedResponse) *pb.MonthAmountTransactionFailedResponse {
	return &pb.MonthAmountTransactionFailedResponse{
		Year:        b.Year,
		Month:       b.Month,
		TotalFailed: int32(b.TotalFailed),
		TotalAmount: int32(b.TotalAmount),
	}
}

func (s *transactionProtoMapper) mapResponsesMonthAmountFailed(b []*response.MonthAmountTransactionFailedResponse) []*pb.MonthAmountTransactionFailedResponse {
	var result []*pb.MonthAmountTransactionFailedResponse

	for _, transaction := range b {
		result = append(result, s.mapResponseMonthAmountFailed(transaction))
	}

	return result
}

func (s *transactionProtoMapper) mapResponseYearAmountFailed(b *response.YearAmountTransactionFailedResponse) *pb.YearAmountTransactionFailedResponse {
	return &pb.YearAmountTransactionFailedResponse{
		Year:        b.Year,
		TotalFailed: int32(b.TotalFailed),
		TotalAmount: int32(b.TotalAmount),
	}
}

func (s *transactionProtoMapper) mapResponsesYearAmountFailed(b []*response.YearAmountTransactionFailedResponse) []*pb.YearAmountTransactionFailedResponse {
	var result []*pb.YearAmountTransactionFailedResponse

	for _, transaction := range b {
		result = append(result, s.mapResponseYearAmountFailed(transaction))
	}

	return result
}

func (s *transactionProtoMapper) mapResponseMonthMethod(b *response.MonthMethodTransactionResponse) *pb.MonthMethodTransactionResponse {
	return &pb.MonthMethodTransactionResponse{
		Month:             b.Month,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int32(b.TotalAmount),
		TotalTransactions: int32(b.TotalTransactions),
	}
}

func (s *transactionProtoMapper) mapResponsesMonthMethod(b []*response.MonthMethodTransactionResponse) []*pb.MonthMethodTransactionResponse {
	var result []*pb.MonthMethodTransactionResponse

	for _, transaction := range b {
		result = append(result, s.mapResponseMonthMethod(transaction))
	}

	return result
}

func (s *transactionProtoMapper) mapResponseYearMethod(b *response.YearMethodTransactionResponse) *pb.YearMethodTransactionResponse {
	return &pb.YearMethodTransactionResponse{
		Year:              b.Year,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int32(b.TotalAmount),
		TotalTransactions: int32(b.TotalTransactions),
	}
}

func (s *transactionProtoMapper) mapResponsesYearMethod(b []*response.YearMethodTransactionResponse) []*pb.YearMethodTransactionResponse {
	var result []*pb.YearMethodTransactionResponse

	for _, transaction := range b {
		result = append(result, s.mapResponseYearMethod(transaction))
	}

	return result
}
