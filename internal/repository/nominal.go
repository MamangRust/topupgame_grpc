package repository

import (
	"context"
	"fmt"
	"topup_game/internal/domain/record"
	"topup_game/internal/domain/requests"
	recordmapper "topup_game/internal/mapper/record"
	db "topup_game/pkg/database/schema"
)

type nominalRepository struct {
	db      *db.Queries
	ctx     context.Context
	mapping recordmapper.NominalRecordMapping
}

func NewNominalRepository(db *db.Queries, ctx context.Context, mapping recordmapper.NominalRecordMapping) *nominalRepository {
	return &nominalRepository{
		db:      db,
		ctx:     ctx,
		mapping: mapping,
	}
}

func (r *nominalRepository) FindAllNominals(page int, pageSize int, search string) ([]*record.NominalRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetNominalsParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetNominals(r.ctx, req)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to find categories: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToNominalRecordsAll(res), totalCount, nil
}

func (r *nominalRepository) FindById(id int) (*record.NominalRecord, error) {
	res, err := r.db.GetNominalByID(r.ctx, int32(id))

	if err != nil {
		return nil, fmt.Errorf("failed to find nominal by id: %w", err)
	}

	return r.mapping.ToNominalRecord(res), nil
}

func (r *nominalRepository) FindByActiveNominal(page int, pageSize int, search string) ([]*record.NominalRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetNominalsActiveParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetNominalsActive(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find active nominals: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToNominalRecordsActive(res), totalCount, nil
}

func (r *nominalRepository) FindByTrashedNominal(page int, pageSize int, search string) ([]*record.NominalRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetNominalsTrashedParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetNominalsTrashed(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find trashed nominals: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToNominalRecordsTrashed(res), totalCount, nil
}

func (r *nominalRepository) CreateNominal(req *requests.CreateNominalRequest) (*record.NominalRecord, error) {
	res, err := r.db.CreateNominal(r.ctx, db.CreateNominalParams{
		VoucherID: int32(req.VoucherID),
		Name:      req.Name,
		Quantity:  int32(req.Quantity),
		Price:     req.Price,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create nominal: %w", err)
	}

	return r.mapping.ToNominalRecord(res), nil
}

func (r *nominalRepository) UpdateNominal(req *requests.UpdateNominalRequest) (*record.NominalRecord, error) {
	res, err := r.db.UpdateNominal(r.ctx, db.UpdateNominalParams{
		NominalID: int32(req.ID),
		VoucherID: int32(req.VoucherID),
		Name:      req.Name,
		Quantity:  int32(req.Quantity),
		Price:     req.Price,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to update nominal: %w", err)
	}

	return r.mapping.ToNominalRecord(res), nil
}

func (r *nominalRepository) UpdateQuantity(nominal int, quantity int) (bool, error) {
	err := r.db.DecreaseNominalQuantity(r.ctx, db.DecreaseNominalQuantityParams{
		Quantity:  int32(quantity),
		NominalID: int32(nominal),
	})

	if err != nil {
		return false, fmt.Errorf("failed to update nominal quantity: %w", err)
	}

	return true, nil
}

func (r *nominalRepository) TrashedNominal(id int) (*record.NominalRecord, error) {
	res, err := r.db.TrashNominal(r.ctx, int32(id))

	if err != nil {
		return nil, fmt.Errorf("failed to trash nominal: %w", err)
	}

	return r.mapping.ToNominalRecord(res), nil
}

func (r *nominalRepository) RestoreNominal(id int) (*record.NominalRecord, error) {
	res, err := r.db.RestoreNominal(r.ctx, int32(id))

	if err != nil {
		return nil, fmt.Errorf("failed to find nominal after restore: %w", err)
	}

	return r.mapping.ToNominalRecord(res), nil
}

func (r *nominalRepository) DeleteNominalPermanent(nominal_id int) (bool, error) {
	err := r.db.DeleteNominalPermanently(r.ctx, int32(nominal_id))

	if err != nil {
		return false, fmt.Errorf("failed to delete nominal: %w", err)
	}

	return true, nil
}

func (r *nominalRepository) RestoreAllNominal() (bool, error) {
	err := r.db.RestoreAllNominals(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to restore all nominal: %w", err)
	}

	return true, nil
}

func (r *nominalRepository) DeleteAllNominalsPermanent() (bool, error) {
	err := r.db.DeleteAllPermanentNominals(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to delete all nominal permanently: %w", err)
	}

	return true, nil
}
