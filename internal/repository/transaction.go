package repository

import (
	"context"
	"database/sql"
	"time"
	"topup_game/internal/domain/record"
	"topup_game/internal/domain/requests"
	recordmapper "topup_game/internal/mapper/record"
	db "topup_game/pkg/database/schema"
	"topup_game/pkg/errors/transaction_errors"
)

type transactionRepository struct {
	db      *db.Queries
	ctx     context.Context
	mapping recordmapper.TransactionRecordMapping
}

func NewTransactionRepository(db *db.Queries, ctx context.Context, mapping recordmapper.TransactionRecordMapping) *transactionRepository {
	return &transactionRepository{
		db:      db,
		ctx:     ctx,
		mapping: mapping,
	}
}

func (r *transactionRepository) FindAllTransactions(request *requests.FindAllTransactions) ([]*record.TransactionRecord, *int, error) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	offset := (page - 1) * pageSize

	req := db.GetTransactionsParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetTransactions(r.ctx, req)

	if err != nil {
		return nil, nil, transaction_errors.ErrFindAllTransactions
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToTransactionRecordsAll(res), &totalCount, nil
}

func (r *transactionRepository) FindMonthAmountTransactionSuccess(req *requests.MonthAmountTransactionRequest) ([]*record.MonthAmountTransactionSuccessRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthlyAmountTransactionSuccess(r.ctx, db.GetMonthlyAmountTransactionSuccessParams{
		Column1: currentDate,
		Column2: lastDayCurrentMonth,
		Column3: prevDate,
		Column4: lastDayPrevMonth,
	})

	if err != nil {
		return nil, transaction_errors.ErrFindMonthAmountTransactionSuccess
	}

	so := r.mapping.ToTransactionsRecordMonthAmountSuccess(res)

	return so, nil
}

func (r *transactionRepository) FindYearAmountTransactionSuccess(year int) ([]*record.YearAmountTransactionSuccessRecord, error) {
	res, err := r.db.GetYearlyAmountTransactionSuccess(r.ctx, int32(year))

	if err != nil {
		return nil, transaction_errors.ErrFindYearAmountTransactionSuccess
	}

	so := r.mapping.ToTransactionsRecordYearAmountSuccess(res)

	return so, nil
}

func (r *transactionRepository) FindMonthAmountTransactionFailed(req *requests.MonthAmountTransactionRequest) ([]*record.MonthAmountTransactionFailedRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)

	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthlyAmountTransactionFailed(r.ctx, db.GetMonthlyAmountTransactionFailedParams{
		Column1: currentDate,
		Column2: lastDayCurrentMonth,
		Column3: prevDate,
		Column4: lastDayPrevMonth,
	})

	if err != nil {
		return nil, transaction_errors.ErrFindMonthAmountTransactionFailed
	}

	so := r.mapping.ToTransactionsRecordMonthAmountFailed(res)

	return so, nil
}

func (r *transactionRepository) FindYearAmountTransactionFailed(year int) ([]*record.YearAmountTransactionFailedRecord, error) {
	res, err := r.db.GetYearlyAmountTransactionFailed(r.ctx, int32(year))

	if err != nil {
		return nil, transaction_errors.ErrFindYearAmountTransactionFailed
	}

	so := r.mapping.ToTransactionsRecordYearAmountFailed(res)

	return so, nil
}

func (r *transactionRepository) FindMonthMethodTransactionSuccess(year int) ([]*record.MonthMethodTransactionRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyTransactionMethodsSuccess(r.ctx, yearStart)

	if err != nil {
		return nil, transaction_errors.ErrFindMonthMethodTransactionSuccess
	}

	so := r.mapping.ToTransactionsRecordMonthMethodSuccess(res)

	return so, nil
}

func (r *transactionRepository) FindYearMethodTransactionSuccess(year int) ([]*record.YearMethodTransactionRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyTransactionMethodsSuccess(r.ctx, yearStart)

	if err != nil {
		return nil, transaction_errors.ErrFindYearMethodTransactionSuccess
	}

	so := r.mapping.ToTransactionsRecordYearMethodSuccess(res)

	return so, nil
}

func (r *transactionRepository) FindMonthMethodTransactionFailed(year int) ([]*record.MonthMethodTransactionRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyTransactionMethodsFailed(r.ctx, yearStart)

	if err != nil {
		return nil, transaction_errors.ErrFindMonthMethodTransactionFailed
	}

	so := r.mapping.ToTransactionsRecordMonthMethodFailed(res)

	return so, nil
}

func (r *transactionRepository) FindYearMethodTransactionFailed(year int) ([]*record.YearMethodTransactionRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyTransactionMethodsFailed(r.ctx, yearStart)

	if err != nil {
		return nil, transaction_errors.ErrFindYearMethodTransactionFailed
	}

	so := r.mapping.ToTransactionsRecordYearMethodFailed(res)

	return so, nil
}

func (r *transactionRepository) FindMonthAmountTransactionSuccessByMerchant(req *requests.MonthAmountTransactionByMerchantRequest) ([]*record.MonthAmountTransactionSuccessRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthlyAmountTransactionSuccessByMerchant(r.ctx, db.GetMonthlyAmountTransactionSuccessByMerchantParams{
		Column1:    currentDate,
		Column2:    lastDayCurrentMonth,
		Column3:    prevDate,
		Column4:    lastDayPrevMonth,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, transaction_errors.ErrFindMonthAmountTransactionSuccessByMerchant
	}

	so := r.mapping.ToTransactionsRecordMonthAmountSuccessByMerchant(res)

	return so, nil
}

func (r *transactionRepository) FindYearAmountTransactionSuccessByMerchant(req *requests.YearAmountTransactionByMerchantRequest) ([]*record.YearAmountTransactionSuccessRecord, error) {
	res, err := r.db.GetYearlyAmountTransactionSuccessByMerchant(r.ctx, db.GetYearlyAmountTransactionSuccessByMerchantParams{
		Column1:    int32(req.Year),
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, transaction_errors.ErrFindYearAmountTransactionSuccessByMerchant
	}

	so := r.mapping.ToTransactionsRecordYearAmountSuccessByMerchant(res)

	return so, nil
}

func (r *transactionRepository) FindMonthAmountTransactionFailedByMerchant(req *requests.MonthAmountTransactionByMerchantRequest) ([]*record.MonthAmountTransactionFailedRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthlyAmountTransactionFailedByMerchant(r.ctx, db.GetMonthlyAmountTransactionFailedByMerchantParams{
		Column1:    currentDate,
		Column2:    lastDayCurrentMonth,
		Column3:    prevDate,
		Column4:    lastDayPrevMonth,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, transaction_errors.ErrFindMonthAmountTransactionFailedByMerchant
	}

	so := r.mapping.ToTransactionsRecordMonthAmountFailedByMerchant(res)

	return so, nil
}

func (r *transactionRepository) FindYearAmountTransactionFailedByMerchant(req *requests.YearAmountTransactionByMerchantRequest) ([]*record.YearAmountTransactionFailedRecord, error) {
	res, err := r.db.GetYearlyAmountTransactionFailedByMerchant(r.ctx, db.GetYearlyAmountTransactionFailedByMerchantParams{
		Column1:    int32(req.Year),
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, transaction_errors.ErrFindYearAmountTransactionFailedByMerchant
	}

	so := r.mapping.ToTransactionsRecordYearAmountFailedByMerchant(res)

	return so, nil
}

func (r *transactionRepository) FindMonthMethodTransactionSuccessByMerchant(req *requests.MonthMethodTransactionByMerchantRequest) ([]*record.MonthMethodTransactionRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyTransactionMethodsSuccessByMerchant(r.ctx, db.GetMonthlyTransactionMethodsSuccessByMerchantParams{
		Column1:    yearStart,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, transaction_errors.ErrFindMonthMethodTransactionSuccessByMerchant
	}

	so := r.mapping.ToTransactionsRecordMonthMethodSuccessByMerchant(res)

	return so, nil
}

func (r *transactionRepository) FindYearMethodTransactionSuccessByMerchant(req *requests.YearMethodTransactionByMerchantRequest) ([]*record.YearMethodTransactionRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyTransactionMethodsSuccessByMerchant(r.ctx, db.GetYearlyTransactionMethodsSuccessByMerchantParams{
		Column1:    yearStart,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, transaction_errors.ErrFindYearMethodTransactionSuccessByMerchant
	}

	so := r.mapping.ToTransactionsRecordYearMethodSuccessByMerchant(res)

	return so, nil
}

func (r *transactionRepository) FindMonthMethodTransactionFailedByMerchant(req *requests.MonthMethodTransactionByMerchantRequest) ([]*record.MonthMethodTransactionRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyTransactionMethodsFailedByMerchant(r.ctx, db.GetMonthlyTransactionMethodsFailedByMerchantParams{
		Column1:    yearStart,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, transaction_errors.ErrFindYearMethodTransactionFailedByMerchant
	}

	so := r.mapping.ToTransactionsRecordMonthMethodFailedByMerchant(res)

	return so, nil
}

func (r *transactionRepository) FindYearMethodTransactionFailedByMerchant(req *requests.YearMethodTransactionByMerchantRequest) ([]*record.YearMethodTransactionRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyTransactionMethodsFailedByMerchant(r.ctx, db.GetYearlyTransactionMethodsFailedByMerchantParams{
		Column1:    yearStart,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, transaction_errors.ErrFindYearMethodTransactionFailedByMerchant
	}

	so := r.mapping.ToTransactionsRecordYearMethodFailedByMerchant(res)

	return so, nil
}

func (r *transactionRepository) FindByActive(request *requests.FindAllTransactions) ([]*record.TransactionRecord, *int, error) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	offset := (page - 1) * pageSize

	req := db.GetTransactionsActiveParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetTransactionsActive(r.ctx, req)

	if err != nil {
		return nil, nil, transaction_errors.ErrFindActiveTransactions
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToTransactionRecordsActive(res), &totalCount, nil
}

func (r *transactionRepository) FindByTrashed(request *requests.FindAllTransactions) ([]*record.TransactionRecord, *int, error) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	offset := (page - 1) * pageSize

	req := db.GetTransactionsTrashedParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetTransactionsTrashed(r.ctx, req)

	if err != nil {
		return nil, nil, transaction_errors.ErrFindTrashedTransactions
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToTransactionRecordsTrashed(res), &totalCount, nil
}

func (r *transactionRepository) FindById(user_id int) (*record.TransactionRecord, error) {
	res, err := r.db.GetTransactionByID(r.ctx, int32(user_id))

	if err != nil {
		return nil, transaction_errors.ErrTransactionNotFound
	}

	return r.mapping.ToTransactionRecord(res), nil
}

func (r *transactionRepository) CreateTransaction(req *requests.CreateTransactionRequest, status string) (*record.TransactionRecord, error) {
	res, err := r.db.CreateTransaction(r.ctx, db.CreateTransactionParams{
		UserID: int32(req.UserID),
		MerchantID: sql.NullInt32{
			Int32: int32(req.MerchantID),
			Valid: true,
		},
		VoucherID: sql.NullInt32{
			Int32: int32(req.VoucherID),
			Valid: req.VoucherID > 0,
		},
		NominalID: sql.NullInt32{
			Int32: int32(req.NominalID),
			Valid: req.NominalID > 0,
		},
		BankID: sql.NullInt32{
			Int32: int32(req.BankID),
			Valid: req.BankID > 0,
		},
		PaymentMethod: req.PaymentMethod,
		Status: sql.NullString{
			String: status,
			Valid:  true,
		},
	})

	if err != nil {
		return nil, transaction_errors.ErrCreateTransaction
	}

	return r.mapping.ToTransactionRecord(res), nil
}

func (r *transactionRepository) UpdateTransaction(req *requests.UpdateTransactionRequest) (*record.TransactionRecord, error) {
	res, err := r.db.UpdateTransaction(r.ctx, db.UpdateTransactionParams{
		TransactionID: int32(req.ID),
		UserID:        int32(req.UserID),
		MerchantID: sql.NullInt32{
			Int32: int32(req.MerchantID),
			Valid: true,
		},
		VoucherID: sql.NullInt32{
			Int32: int32(req.VoucherID),
			Valid: req.VoucherID > 0,
		},
		NominalID: sql.NullInt32{
			Int32: int32(req.NominalID),
			Valid: req.NominalID > 0,
		},
		BankID: sql.NullInt32{
			Int32: int32(req.BankID),
			Valid: req.BankID > 0,
		},
		PaymentMethod: req.PaymentMethod,
		Status: sql.NullString{
			String: *req.Status,
			Valid:  true,
		},
	})

	if err != nil {
		return nil, transaction_errors.ErrUpdateTransaction
	}

	return r.mapping.ToTransactionRecord(res), nil
}

func (r *transactionRepository) UpdateTransactionStatus(transaction_id int, status string) (bool, error) {
	err := r.db.UpdateTransactionStatus(r.ctx, db.UpdateTransactionStatusParams{
		TransactionID: int32(transaction_id),
		Status:        sql.NullString{String: status, Valid: true},
	})

	if err != nil {
		return false, transaction_errors.ErrUpdateTransaction
	}

	return true, nil
}

func (r *transactionRepository) TrashTransaction(id int) (*record.TransactionRecord, error) {
	res, err := r.db.TrashTransaction(r.ctx, int32(id))

	if err != nil {
		return nil, transaction_errors.ErrTrashedTransaction
	}

	return r.mapping.ToTransactionRecord(res), nil
}

func (r *transactionRepository) RestoreTransaction(id int) (*record.TransactionRecord, error) {
	res, err := r.db.RestoreTransaction(r.ctx, int32(id))

	if err != nil {
		return nil, transaction_errors.ErrRestoreTransaction
	}

	return r.mapping.ToTransactionRecord(res), nil
}

func (r *transactionRepository) DeleteTransactionPermanent(transaction_id int) (bool, error) {
	err := r.db.DeleteTransactionPermanently(r.ctx, int32(transaction_id))

	if err != nil {
		return false, transaction_errors.ErrDeleteTransactionPermanent
	}

	return true, nil
}

func (r *transactionRepository) RestoreAllTransactions() (bool, error) {
	err := r.db.RestoreAllTransactions(r.ctx)

	if err != nil {
		return false, transaction_errors.ErrRestoreTransaction
	}

	return true, nil
}

func (r *transactionRepository) DeleteAllTransactionsPermanent() (bool, error) {
	err := r.db.DeleteAllPermanentTransactions(r.ctx)

	if err != nil {
		return false, transaction_errors.ErrDeleteAllTransactions
	}

	return true, nil
}
