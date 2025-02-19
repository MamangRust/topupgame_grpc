package repository

import (
	"context"
	"fmt"
	"topup_game/internal/domain/record"
	"topup_game/internal/domain/requests"
	recordmapper "topup_game/internal/mapper/record"
	db "topup_game/pkg/database/schema"
)

type voucherRepository struct {
	db      *db.Queries
	ctx     context.Context
	mapping recordmapper.VoucherRecordMapping
}

func NewVoucherRepository(db *db.Queries, ctx context.Context, mapping recordmapper.VoucherRecordMapping) *voucherRepository {
	return &voucherRepository{
		db:      db,
		ctx:     ctx,
		mapping: mapping,
	}
}

func (r *voucherRepository) FindAllVouchers(page int, pageSize int, search string) ([]*record.VoucherRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetVouchersParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetVouchers(r.ctx, req)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to find vouchers: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToVouchersRecordAll(res), totalCount, nil
}

func (r *voucherRepository) FindById(id int) (*record.VoucherRecord, error) {
	res, err := r.db.GetVoucherByID(r.ctx, int32(id))

	if err != nil {
		return nil, fmt.Errorf("failed to find voucher by id: %w", err)
	}

	return r.mapping.ToVoucherRecord(res), nil
}

func (r *voucherRepository) FindByActiveVouchers(page int, pageSize int, search string) ([]*record.VoucherRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetVouchersActiveParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetVouchersActive(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find active voucher: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToVouchersRecordActive(res), totalCount, nil
}

func (r *voucherRepository) FindByTrashedVoucher(page int, pageSize int, search string) ([]*record.VoucherRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetVouchersTrashedParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetVouchersTrashed(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find trashed vouchers: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToVouchersRecordTrashed(res), totalCount, nil
}

func (r *voucherRepository) CreateVoucher(req *requests.CreateVoucherRequest) (*record.VoucherRecord, error) {
	res, err := r.db.CreateVoucher(r.ctx, db.CreateVoucherParams{
		MerchantID: int32(req.MerchantID),
		CategoryID: int32(req.CategoryID),
		Name:       req.Name,
		ImageName:  req.ImageName,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create voucher: %w", err)
	}

	return r.mapping.ToVoucherRecord(res), nil
}

func (r *voucherRepository) UpdateVoucher(req *requests.UpdateVoucherRequest) (*record.VoucherRecord, error) {
	res, err := r.db.UpdateVoucher(r.ctx, db.UpdateVoucherParams{
		VoucherID:  int32(req.ID),
		CategoryID: int32(req.CategoryID),
		Name:       req.Name,
		ImageName:  req.ImageName,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create voucher: %w", err)
	}

	return r.mapping.ToVoucherRecord(res), nil
}

func (r *voucherRepository) TrashVoucher(id int) (*record.VoucherRecord, error) {
	res, err := r.db.TrashVoucher(r.ctx, int32(id))

	if err != nil {
		return nil, fmt.Errorf("failed to trash voucher: %w", err)
	}

	return r.mapping.ToVoucherRecord(res), nil
}

func (r *voucherRepository) RestoreVoucher(id int) (*record.VoucherRecord, error) {
	res, err := r.db.RestoreVoucher(r.ctx, int32(id))

	if err != nil {
		return nil, fmt.Errorf("failed to find category after restore: %w", err)
	}

	return r.mapping.ToVoucherRecord(res), nil
}

func (r *voucherRepository) DeleteVoucherPermanent(category_id int) (bool, error) {
	err := r.db.DeleteVoucherPermanently(r.ctx, int32(category_id))

	if err != nil {
		return false, fmt.Errorf("failed to delete category: %w", err)
	}

	return true, nil
}

func (r *voucherRepository) RestoreAllVouchers() (bool, error) {
	err := r.db.RestoreAllVouchers(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to restore all vouchers: %w", err)
	}

	return true, nil
}

func (r *voucherRepository) DeleteAllVouchersPermanent() (bool, error) {
	err := r.db.DeleteAllPermanentVouchers(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to delete all vouchers permanently: %w", err)
	}

	return true, nil
}
