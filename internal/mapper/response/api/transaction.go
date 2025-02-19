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

func (s *transactionResponseMapper) mapResponseTransaction(tx *pb.TransactionResponse) *response.TransactionResponse {
	return &response.TransactionResponse{
		ID:            int(tx.Id),
		UserID:        int(tx.UserId),
		MerchantID:    int(tx.MerchantId),
		VoucherID:     int(tx.VoucherId),
		NominalID:     int(tx.NominalId),
		CategoryID:    int(tx.CategoryId),
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
		NominalID:     int(tx.NominalId),
		CategoryID:    int(tx.CategoryId),
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
