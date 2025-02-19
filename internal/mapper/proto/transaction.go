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

func (s *transactionProtoMapper) mapResponseTransaction(transaction *response.TransactionResponse) *pb.TransactionResponse {
	return &pb.TransactionResponse{
		Id:            int32(transaction.ID),
		UserId:        int32(transaction.UserID),
		MerchantId:    int32(transaction.MerchantID),
		VoucherId:     int32(transaction.VoucherID),
		NominalId:     int32(transaction.NominalID),
		CategoryId:    int32(transaction.CategoryID),
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
		CategoryId:    int32(transaction.CategoryID),
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
