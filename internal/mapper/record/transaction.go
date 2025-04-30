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

func (s *transactionRecordMapper) ToTransactionRecordMonthAmountSuccess(b *db.GetMonthlyAmountTransactionSuccessRow) *record.MonthAmountTransactionSuccessRecord {
	return &record.MonthAmountTransactionSuccessRecord{
		Year:         b.RmYear,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *transactionRecordMapper) ToTransactionsRecordMonthAmountSuccess(b []*db.GetMonthlyAmountTransactionSuccessRow) []*record.MonthAmountTransactionSuccessRecord {
	var result []*record.MonthAmountTransactionSuccessRecord

	for _, transaction := range b {
		result = append(result, s.ToTransactionRecordMonthAmountSuccess(transaction))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionRecordYearAmountSuccess(b *db.GetYearlyAmountTransactionSuccessRow) *record.YearAmountTransactionSuccessRecord {
	return &record.YearAmountTransactionSuccessRecord{
		Year:         b.RyYear,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *transactionRecordMapper) ToTransactionsRecordYearAmountSuccess(b []*db.GetYearlyAmountTransactionSuccessRow) []*record.YearAmountTransactionSuccessRecord {
	var result []*record.YearAmountTransactionSuccessRecord

	for _, transaction := range b {
		result = append(result, s.ToTransactionRecordYearAmountSuccess(transaction))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionRecordMonthAmountFailed(b *db.GetMonthlyAmountTransactionFailedRow) *record.MonthAmountTransactionFailedRecord {
	return &record.MonthAmountTransactionFailedRecord{
		Year:        b.RmYear,
		Month:       b.Month,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *transactionRecordMapper) ToTransactionsRecordMonthAmountFailed(b []*db.GetMonthlyAmountTransactionFailedRow) []*record.MonthAmountTransactionFailedRecord {
	var result []*record.MonthAmountTransactionFailedRecord

	for _, transaction := range b {
		result = append(result, s.ToTransactionRecordMonthAmountFailed(transaction))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionRecordYearAmountFailed(b *db.GetYearlyAmountTransactionFailedRow) *record.YearAmountTransactionFailedRecord {
	return &record.YearAmountTransactionFailedRecord{
		Year:        b.RyYear,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *transactionRecordMapper) ToTransactionsRecordYearAmountFailed(b []*db.GetYearlyAmountTransactionFailedRow) []*record.YearAmountTransactionFailedRecord {
	var result []*record.YearAmountTransactionFailedRecord

	for _, transaction := range b {
		result = append(result, s.ToTransactionRecordYearAmountFailed(transaction))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionRecordMonthMethodSuccess(b *db.GetMonthlyTransactionMethodsSuccessRow) *record.MonthMethodTransactionRecord {
	return &record.MonthMethodTransactionRecord{
		Month:             b.Month,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *transactionRecordMapper) ToTransactionsRecordMonthMethodSuccess(b []*db.GetMonthlyTransactionMethodsSuccessRow) []*record.MonthMethodTransactionRecord {
	var result []*record.MonthMethodTransactionRecord

	for _, transaction := range b {
		result = append(result, s.ToTransactionRecordMonthMethodSuccess(transaction))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionRecordMonthMethodFailed(b *db.GetMonthlyTransactionMethodsFailedRow) *record.MonthMethodTransactionRecord {
	return &record.MonthMethodTransactionRecord{
		Month:             b.Month,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *transactionRecordMapper) ToTransactionsRecordMonthMethodFailed(b []*db.GetMonthlyTransactionMethodsFailedRow) []*record.MonthMethodTransactionRecord {
	var result []*record.MonthMethodTransactionRecord

	for _, transaction := range b {
		result = append(result, s.ToTransactionRecordMonthMethodFailed(transaction))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionRecordYearMethodSuccess(b *db.GetYearlyTransactionMethodsSuccessRow) *record.YearMethodTransactionRecord {
	return &record.YearMethodTransactionRecord{
		Year:              b.Year,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *transactionRecordMapper) ToTransactionsRecordYearMethodSuccess(b []*db.GetYearlyTransactionMethodsSuccessRow) []*record.YearMethodTransactionRecord {
	var result []*record.YearMethodTransactionRecord

	for _, transaction := range b {
		result = append(result, s.ToTransactionRecordYearMethodSuccess(transaction))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionRecordYearMethodFailed(b *db.GetYearlyTransactionMethodsFailedRow) *record.YearMethodTransactionRecord {
	return &record.YearMethodTransactionRecord{
		Year:              b.Year,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *transactionRecordMapper) ToTransactionsRecordYearMethodFailed(b []*db.GetYearlyTransactionMethodsFailedRow) []*record.YearMethodTransactionRecord {
	var result []*record.YearMethodTransactionRecord

	for _, transaction := range b {
		result = append(result, s.ToTransactionRecordYearMethodFailed(transaction))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionRecordMonthAmountSuccessByMerchant(b *db.GetMonthlyAmountTransactionSuccessByMerchantRow) *record.MonthAmountTransactionSuccessRecord {
	return &record.MonthAmountTransactionSuccessRecord{
		Year:         b.Year,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *transactionRecordMapper) ToTransactionsRecordMonthAmountSuccessByMerchant(b []*db.GetMonthlyAmountTransactionSuccessByMerchantRow) []*record.MonthAmountTransactionSuccessRecord {
	var result []*record.MonthAmountTransactionSuccessRecord

	for _, transaction := range b {
		result = append(result, s.ToTransactionRecordMonthAmountSuccessByMerchant(transaction))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionRecordYearAmountSuccessByMerchant(b *db.GetYearlyAmountTransactionSuccessByMerchantRow) *record.YearAmountTransactionSuccessRecord {
	return &record.YearAmountTransactionSuccessRecord{
		Year:         b.YsYear,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *transactionRecordMapper) ToTransactionsRecordYearAmountSuccessByMerchant(b []*db.GetYearlyAmountTransactionSuccessByMerchantRow) []*record.YearAmountTransactionSuccessRecord {
	var result []*record.YearAmountTransactionSuccessRecord

	for _, transaction := range b {
		result = append(result, s.ToTransactionRecordYearAmountSuccessByMerchant(transaction))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionRecordMonthAmountFailedByMerchant(b *db.GetMonthlyAmountTransactionFailedByMerchantRow) *record.MonthAmountTransactionFailedRecord {
	return &record.MonthAmountTransactionFailedRecord{
		Year:        b.Year,
		Month:       b.Month,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *transactionRecordMapper) ToTransactionsRecordMonthAmountFailedByMerchant(b []*db.GetMonthlyAmountTransactionFailedByMerchantRow) []*record.MonthAmountTransactionFailedRecord {
	var result []*record.MonthAmountTransactionFailedRecord

	for _, transaction := range b {
		result = append(result, s.ToTransactionRecordMonthAmountFailedByMerchant(transaction))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionRecordYearAmountFailedByMerchant(b *db.GetYearlyAmountTransactionFailedByMerchantRow) *record.YearAmountTransactionFailedRecord {
	return &record.YearAmountTransactionFailedRecord{
		Year:        b.YsYear,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *transactionRecordMapper) ToTransactionsRecordYearAmountFailedByMerchant(b []*db.GetYearlyAmountTransactionFailedByMerchantRow) []*record.YearAmountTransactionFailedRecord {
	var result []*record.YearAmountTransactionFailedRecord

	for _, transaction := range b {
		result = append(result, s.ToTransactionRecordYearAmountFailedByMerchant(transaction))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionRecordMonthMethodSuccessByMerchant(b *db.GetMonthlyTransactionMethodsSuccessByMerchantRow) *record.MonthMethodTransactionRecord {
	return &record.MonthMethodTransactionRecord{
		Month:             b.Month,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *transactionRecordMapper) ToTransactionsRecordMonthMethodSuccessByMerchant(b []*db.GetMonthlyTransactionMethodsSuccessByMerchantRow) []*record.MonthMethodTransactionRecord {
	var result []*record.MonthMethodTransactionRecord

	for _, transaction := range b {
		result = append(result, s.ToTransactionRecordMonthMethodSuccessByMerchant(transaction))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionRecordMonthMethodFailedByMerchant(b *db.GetMonthlyTransactionMethodsFailedByMerchantRow) *record.MonthMethodTransactionRecord {
	return &record.MonthMethodTransactionRecord{
		Month:             b.Month,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *transactionRecordMapper) ToTransactionsRecordMonthMethodFailedByMerchant(b []*db.GetMonthlyTransactionMethodsFailedByMerchantRow) []*record.MonthMethodTransactionRecord {
	var result []*record.MonthMethodTransactionRecord

	for _, transaction := range b {
		result = append(result, s.ToTransactionRecordMonthMethodFailedByMerchant(transaction))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionRecordYearMethodSuccessByMerchant(b *db.GetYearlyTransactionMethodsSuccessByMerchantRow) *record.YearMethodTransactionRecord {
	return &record.YearMethodTransactionRecord{
		Year:              b.Year,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *transactionRecordMapper) ToTransactionsRecordYearMethodSuccessByMerchant(b []*db.GetYearlyTransactionMethodsSuccessByMerchantRow) []*record.YearMethodTransactionRecord {
	var result []*record.YearMethodTransactionRecord

	for _, transaction := range b {
		result = append(result, s.ToTransactionRecordYearMethodSuccessByMerchant(transaction))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionRecordYearMethodFailedByMerchant(b *db.GetYearlyTransactionMethodsFailedByMerchantRow) *record.YearMethodTransactionRecord {
	return &record.YearMethodTransactionRecord{
		Year:              b.Year,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *transactionRecordMapper) ToTransactionsRecordYearMethodFailedByMerchant(b []*db.GetYearlyTransactionMethodsFailedByMerchantRow) []*record.YearMethodTransactionRecord {
	var result []*record.YearMethodTransactionRecord

	for _, transaction := range b {
		result = append(result, s.ToTransactionRecordYearMethodFailedByMerchant(transaction))
	}

	return result
}
