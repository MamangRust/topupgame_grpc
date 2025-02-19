package repository

import (
	"context"
	"fmt"
	"topup_game/internal/domain/record"
	"topup_game/internal/domain/requests"
	recordmapper "topup_game/internal/mapper/record"
	db "topup_game/pkg/database/schema"
)

type bankRepository struct {
	db      *db.Queries
	ctx     context.Context
	mapping recordmapper.BankRecordMapping
}

func NewBankRepository(db *db.Queries, ctx context.Context, mapping recordmapper.BankRecordMapping) *bankRepository {
	return &bankRepository{
		db:      db,
		ctx:     ctx,
		mapping: mapping,
	}
}

func (r *bankRepository) FindAllBanks(page int, pageSize int, search string) ([]*record.BankRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetBanksParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetBanks(r.ctx, req)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to find banks: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToBanksRecordAll(res), totalCount, nil
}

func (r *bankRepository) FindById(id int) (*record.BankRecord, error) {
	res, err := r.db.GetBankByID(r.ctx, int32(id))

	if err != nil {
		return nil, fmt.Errorf("failed to find bank by id: %w", err)
	}

	return r.mapping.ToBankRecord(res), nil
}

func (r *bankRepository) FindByActiveBanks(page int, pageSize int, search string) ([]*record.BankRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetBanksActiveParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetBanksActive(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find active banks: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToBanksRecordActive(res), totalCount, nil
}

func (r *bankRepository) FindByTrashedBanks(page int, pageSize int, search string) ([]*record.BankRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetBanksTrashedParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetBanksTrashed(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find trashed banks: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToBanksRecordTrashed(res), totalCount, nil
}

func (r *bankRepository) CreateBank(req *requests.CreateBankRequest) (*record.BankRecord, error) {
	res, err := r.db.CreateBank(r.ctx, req.Name)

	if err != nil {
		return nil, fmt.Errorf("failed to create bank: %w", err)
	}

	return r.mapping.ToBankRecord(res), nil
}

func (r *bankRepository) UpdateBank(req *requests.UpdateBankRequest) (*record.BankRecord, error) {
	res, err := r.db.UpdateBank(r.ctx, db.UpdateBankParams{
		BankID: int32(req.ID),
		Name:   req.Name,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to update bank: %w", err)
	}

	return r.mapping.ToBankRecord(res), nil
}

func (r *bankRepository) TrashedBank(id int) (*record.BankRecord, error) {
	res, err := r.db.TrashBank(r.ctx, int32(id))

	if err != nil {
		return nil, fmt.Errorf("failed to trash banks: %w", err)
	}

	return r.mapping.ToBankRecord(res), nil
}

func (r *bankRepository) RestoreBank(id int) (*record.BankRecord, error) {
	res, err := r.db.RestoreBank(r.ctx, int32(id))

	if err != nil {
		return nil, fmt.Errorf("failed to find banks after restore: %w", err)
	}

	return r.mapping.ToBankRecord(res), nil
}

func (r *bankRepository) DeleteBankPermanent(bank_id int) (bool, error) {
	err := r.db.DeleteBankPermanently(r.ctx, int32(bank_id))

	if err != nil {
		return false, fmt.Errorf("failed to delete bank: %w", err)
	}

	return true, nil
}

func (r *bankRepository) RestoreAllBanks() (bool, error) {
	err := r.db.RestoreAllBanks(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to restore all banks: %w", err)
	}

	return true, nil
}

func (r *bankRepository) DeleteAllBanksPermanent() (bool, error) {
	err := r.db.DeleteAllPermanentBanks(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to delete all banks permanently: %w", err)
	}

	return true, nil
}
