package response_service

import (
	"topup_game/internal/domain/record"
	"topup_game/internal/domain/response"
)

type transactionResponseMapper struct {
}

func NewTransactionResponseMapper() *transactionResponseMapper {
	return &transactionResponseMapper{}
}

func (s *transactionResponseMapper) ToTransactionResponse(transaction *record.TransactionRecord) *response.TransactionResponse {
	return &response.TransactionResponse{
		ID:            transaction.ID,
		UserID:        transaction.UserID,
		MerchantID:    transaction.MerchantID,
		VoucherID:     transaction.VoucherID,
		NominalID:     transaction.NominalID,
		CategoryID:    transaction.CategoryID,
		BankID:        transaction.BankID,
		PaymentMethod: transaction.PaymentMethod,
		Status:        transaction.Status,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	}
}

func (s *transactionResponseMapper) ToTransactionsResponse(transactions []*record.TransactionRecord) []*response.TransactionResponse {
	var responseTransactions []*response.TransactionResponse

	for _, transaction := range transactions {
		responseTransactions = append(responseTransactions, s.ToTransactionResponse(transaction))
	}

	return responseTransactions
}

func (s *transactionResponseMapper) ToTransactionResponseDeleteAt(transaction *record.TransactionRecord) *response.TransactionResponseDeleteAt {
	return &response.TransactionResponseDeleteAt{
		ID:            transaction.ID,
		UserID:        transaction.UserID,
		MerchantID:    transaction.MerchantID,
		VoucherID:     transaction.VoucherID,
		NominalID:     transaction.NominalID,
		CategoryID:    transaction.CategoryID,
		BankID:        transaction.BankID,
		PaymentMethod: transaction.PaymentMethod,
		Status:        transaction.Status,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
		DeletedAt:     *transaction.DeletedAt,
	}
}

func (s *transactionResponseMapper) ToTransactionsResponseDeleteAt(transactions []*record.TransactionRecord) []*response.TransactionResponseDeleteAt {
	var responseTransactions []*response.TransactionResponseDeleteAt

	for _, transaction := range transactions {
		responseTransactions = append(responseTransactions, s.ToTransactionResponseDeleteAt(transaction))
	}

	return responseTransactions
}

func (s *transactionResponseMapper) ToTransactionResponseMonthAmountSuccess(b *record.MonthAmountTransactionSuccessRecord) *response.MonthAmountTransactionSuccessResponse {
	return &response.MonthAmountTransactionSuccessResponse{
		Year:         b.Year,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *transactionResponseMapper) ToTransactionsResponseMonthAmountSuccess(b []*record.MonthAmountTransactionSuccessRecord) []*response.MonthAmountTransactionSuccessResponse {
	var result []*response.MonthAmountTransactionSuccessResponse

	for _, transaction := range b {
		result = append(result, s.ToTransactionResponseMonthAmountSuccess(transaction))
	}

	return result
}

func (s *transactionResponseMapper) ToTransactionResponseYearAmountSuccess(b *record.YearAmountTransactionSuccessRecord) *response.YearAmountTransactionSuccessResponse {
	return &response.YearAmountTransactionSuccessResponse{
		Year:         b.Year,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *transactionResponseMapper) ToTransactionsResponseYearAmountSuccess(b []*record.YearAmountTransactionSuccessRecord) []*response.YearAmountTransactionSuccessResponse {
	var result []*response.YearAmountTransactionSuccessResponse

	for _, transaction := range b {
		result = append(result, s.ToTransactionResponseYearAmountSuccess(transaction))
	}

	return result
}

func (s *transactionResponseMapper) ToTransactionResponseMonthAmountFailed(b *record.MonthAmountTransactionFailedRecord) *response.MonthAmountTransactionFailedResponse {
	return &response.MonthAmountTransactionFailedResponse{
		Year:        b.Year,
		Month:       b.Month,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *transactionResponseMapper) ToTransactionsResponseMonthAmountFailed(b []*record.MonthAmountTransactionFailedRecord) []*response.MonthAmountTransactionFailedResponse {
	var result []*response.MonthAmountTransactionFailedResponse

	for _, transaction := range b {
		result = append(result, s.ToTransactionResponseMonthAmountFailed(transaction))
	}

	return result
}

func (s *transactionResponseMapper) ToTransactionResponseYearAmountFailed(b *record.YearAmountTransactionFailedRecord) *response.YearAmountTransactionFailedResponse {
	return &response.YearAmountTransactionFailedResponse{
		Year:        b.Year,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *transactionResponseMapper) ToTransactionsResponseYearAmountFailed(b []*record.YearAmountTransactionFailedRecord) []*response.YearAmountTransactionFailedResponse {
	var result []*response.YearAmountTransactionFailedResponse

	for _, transaction := range b {
		result = append(result, s.ToTransactionResponseYearAmountFailed(transaction))
	}

	return result
}

func (s *transactionResponseMapper) ToTransactionResponseMonthMethod(b *record.MonthMethodTransactionRecord) *response.MonthMethodTransactionResponse {
	return &response.MonthMethodTransactionResponse{
		Month:             b.Month,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *transactionResponseMapper) ToTransactionsResponseMonthMethod(b []*record.MonthMethodTransactionRecord) []*response.MonthMethodTransactionResponse {
	var result []*response.MonthMethodTransactionResponse

	for _, transaction := range b {
		result = append(result, s.ToTransactionResponseMonthMethod(transaction))
	}

	return result
}

func (s *transactionResponseMapper) ToTransactionResponseYearMethod(b *record.YearMethodTransactionRecord) *response.YearMethodTransactionResponse {
	return &response.YearMethodTransactionResponse{
		Year:              b.Year,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *transactionResponseMapper) ToTransactionsResponseYearMethod(b []*record.YearMethodTransactionRecord) []*response.YearMethodTransactionResponse {
	var result []*response.YearMethodTransactionResponse

	for _, transaction := range b {
		result = append(result, s.ToTransactionResponseYearMethod(transaction))
	}

	return result
}
