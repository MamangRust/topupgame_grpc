package repository

import (
	"context"
	"database/sql"
	"fmt"
	"topup_game/internal/domain/record"
	"topup_game/internal/domain/requests"
	recordmapper "topup_game/internal/mapper/record"
	db "topup_game/pkg/database/schema"
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

func (r *transactionRepository) FindAllTransactions(search string, page, pageSize int) ([]*record.TransactionRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetTransactionsParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetTransactions(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find Transactions: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToTransactionRecordsAll(res), totalCount, nil
}

func (r *transactionRepository) FindByActive(search string, page, pageSize int) ([]*record.TransactionRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetTransactionsActiveParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetTransactionsActive(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find Transaction: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToTransactionRecordsActive(res), totalCount, nil
}

func (r *transactionRepository) FindByTrashed(search string, page, pageSize int) ([]*record.TransactionRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetTransactionsTrashedParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetTransactionsTrashed(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find Transaction: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToTransactionRecordsTrashed(res), totalCount, nil
}

func (r *transactionRepository) FindById(user_id int) (*record.TransactionRecord, error) {
	res, err := r.db.GetTransactionByID(r.ctx, int32(user_id))

	if err != nil {
		fmt.Printf("Error fetching user: %v\n", err)

		return nil, fmt.Errorf("failed to find users: %w", err)
	}

	return r.mapping.ToTransactionRecord(res), nil
}

func (r *transactionRepository) CreateTransaction(req *requests.CreateTransactionRequest, status string) (*record.TransactionRecord, error) {
	res, err := r.db.CreateTransaction(r.ctx, db.CreateTransactionParams{
		UserID: int32(req.UserID),
		MerchantID: sql.NullInt32{
			Int32: int32(req.MerchantID),
		},
		VoucherID:     sql.NullInt32{Int32: int32(req.VoucherID)},
		NominalID:     sql.NullInt32{Int32: int32(req.NominalID)},
		CategoryID:    sql.NullInt32{Int32: int32(req.CategoryID)},
		BankID:        sql.NullInt32{Int32: int32(req.BankID)},
		PaymentMethod: req.PaymentMethod,
		Status:        sql.NullString{String: status},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	return r.mapping.ToTransactionRecord(res), nil
}

func (r *transactionRepository) UpdateTransaction(req *requests.UpdateTransactionRequest) (*record.TransactionRecord, error) {
	res, err := r.db.UpdateTransaction(r.ctx, db.UpdateTransactionParams{
		TransactionID: int32(req.ID),
		UserID:        int32(req.UserID),
		MerchantID: sql.NullInt32{
			Int32: int32(req.MerchantID),
		},
		VoucherID:     sql.NullInt32{Int32: int32(req.VoucherID)},
		NominalID:     sql.NullInt32{Int32: int32(req.NominalID)},
		CategoryID:    sql.NullInt32{Int32: int32(req.CategoryID)},
		BankID:        sql.NullInt32{Int32: int32(req.BankID)},
		PaymentMethod: req.PaymentMethod,
		Status:        sql.NullString{String: req.Status},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to update transaction: %w", err)
	}

	return r.mapping.ToTransactionRecord(res), nil
}

func (r *transactionRepository) UpdateTransactionStatus(transaction_id int, status string) (bool, error) {
	err := r.db.UpdateTransactionStatus(r.ctx, db.UpdateTransactionStatusParams{
		TransactionID: int32(transaction_id),
		Status:        sql.NullString{String: status},
	})

	if err != nil {
		return false, fmt.Errorf("failed to update transaction status: %w", err)
	}

	return true, nil
}

func (r *transactionRepository) TrashTransaction(id int) (*record.TransactionRecord, error) {
	res, err := r.db.TrashTransaction(r.ctx, int32(id))

	if err != nil {
		return nil, fmt.Errorf("failed to trash transaction: %w", err)
	}

	return r.mapping.ToTransactionRecord(res), nil
}

func (r *transactionRepository) RestoreTransaction(id int) (*record.TransactionRecord, error) {
	res, err := r.db.RestoreTransaction(r.ctx, int32(id))

	if err != nil {
		return nil, fmt.Errorf("failed to find transaction after restore: %w", err)
	}

	return r.mapping.ToTransactionRecord(res), nil
}

func (r *transactionRepository) DeleteTransactionPermanent(transaction_id int) (bool, error) {
	err := r.db.DeleteTransactionPermanently(r.ctx, int32(transaction_id))

	if err != nil {
		return false, fmt.Errorf("failed to delete transaction: %w", err)
	}

	return true, nil
}

func (r *transactionRepository) RestoreAllTransactions() (bool, error) {
	err := r.db.RestoreAllTransactions(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to restore all transaction: %w", err)
	}

	return true, nil
}

func (r *transactionRepository) DeleteAllTransactionsPermanent() (bool, error) {
	err := r.db.DeleteAllPermanentTransactions(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to delete all transaction permanently: %w", err)
	}

	return true, nil
}
