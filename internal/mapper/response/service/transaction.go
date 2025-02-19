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
