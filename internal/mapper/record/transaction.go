package recordmapper

import (
	"topup_game/internal/domain/record"
	db "topup_game/pkg/database/schema"
)

type transactionRecordMapper struct{}

func NewTransactionRecordMapper() *transactionRecordMapper {
	return &transactionRecordMapper{}
}

func (s *transactionRecordMapper) ToTransactionRecord(transaction *db.Transaction) *record.TransactionRecord {
	deletedAt := transaction.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.TransactionRecord{
		ID:            int(transaction.TransactionID),
		UserID:        int(transaction.UserID),
		MerchantID:    int(transaction.MerchantID.Int32),
		VoucherID:     int(transaction.VoucherID.Int32),
		NominalID:     int(transaction.NominalID.Int32),
		CategoryID:    int(transaction.CategoryID.Int32),
		BankID:        int(transaction.BankID.Int32),
		PaymentMethod: transaction.PaymentMethod,
		Status:        transaction.Status.String,
		CreatedAt:     transaction.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:     transaction.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:     &deletedAt,
	}
}

func (s *transactionRecordMapper) ToTransactionRecords(transactions []*db.Transaction) []*record.TransactionRecord {
	var result []*record.TransactionRecord

	for _, transaction := range transactions {
		result = append(result, s.ToTransactionRecord(transaction))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionRecordAll(transaction *db.GetTransactionsRow) *record.TransactionRecord {
	deletedAt := transaction.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.TransactionRecord{
		ID:            int(transaction.TransactionID),
		UserID:        int(transaction.UserID),
		MerchantID:    int(transaction.MerchantID.Int32),
		VoucherID:     int(transaction.VoucherID.Int32),
		NominalID:     int(transaction.NominalID.Int32),
		CategoryID:    int(transaction.CategoryID.Int32),
		BankID:        int(transaction.BankID.Int32),
		PaymentMethod: transaction.PaymentMethod,
		Status:        transaction.Status.String,
		CreatedAt:     transaction.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:     transaction.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:     &deletedAt,
	}
}

func (s *transactionRecordMapper) ToTransactionRecordsAll(transactions []*db.GetTransactionsRow) []*record.TransactionRecord {
	var result []*record.TransactionRecord

	for _, transaction := range transactions {
		result = append(result, s.ToTransactionRecordAll(transaction))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionRecordActive(transaction *db.GetTransactionsActiveRow) *record.TransactionRecord {
	deletedAt := transaction.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.TransactionRecord{
		ID:            int(transaction.TransactionID),
		UserID:        int(transaction.UserID),
		MerchantID:    int(transaction.MerchantID.Int32),
		VoucherID:     int(transaction.VoucherID.Int32),
		NominalID:     int(transaction.NominalID.Int32),
		CategoryID:    int(transaction.CategoryID.Int32),
		BankID:        int(transaction.BankID.Int32),
		PaymentMethod: transaction.PaymentMethod,
		Status:        transaction.Status.String,
		CreatedAt:     transaction.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:     transaction.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:     &deletedAt,
	}
}

func (s *transactionRecordMapper) ToTransactionRecordsActive(transactions []*db.GetTransactionsActiveRow) []*record.TransactionRecord {
	var result []*record.TransactionRecord

	for _, transaction := range transactions {
		result = append(result, s.ToTransactionRecordActive(transaction))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionRecordTrashed(transaction *db.GetTransactionsTrashedRow) *record.TransactionRecord {
	deletedAt := transaction.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.TransactionRecord{
		ID:            int(transaction.TransactionID),
		UserID:        int(transaction.UserID),
		MerchantID:    int(transaction.MerchantID.Int32),
		VoucherID:     int(transaction.VoucherID.Int32),
		NominalID:     int(transaction.NominalID.Int32),
		CategoryID:    int(transaction.CategoryID.Int32),
		BankID:        int(transaction.BankID.Int32),
		PaymentMethod: transaction.PaymentMethod,
		Status:        transaction.Status.String,
		CreatedAt:     transaction.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:     transaction.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:     &deletedAt,
	}
}

func (s *transactionRecordMapper) ToTransactionRecordsTrashed(transactions []*db.GetTransactionsTrashedRow) []*record.TransactionRecord {
	var result []*record.TransactionRecord

	for _, transaction := range transactions {
		result = append(result, s.ToTransactionRecordTrashed(transaction))
	}

	return result
}
