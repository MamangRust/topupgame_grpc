package repository

import (
	"context"
	"database/sql"
	"time"
	"topup_game/internal/domain/record"
	"topup_game/internal/domain/requests"
	recordmapper "topup_game/internal/mapper/record"
	db "topup_game/pkg/database/schema"
	"topup_game/pkg/errors/voucher_errors"
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

func (r *voucherRepository) FindAllVouchers(request *requests.FindAllVouchers) ([]*record.VoucherRecord, *int, error) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	offset := (page - 1) * pageSize

	req := db.GetVouchersParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetVouchers(r.ctx, req)
	if err != nil {
		return nil, nil, voucher_errors.ErrFindAllVouchers
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToVouchersRecordAll(res), &totalCount, nil
}

func (r *voucherRepository) FindMonthAmountVoucherSuccess(req *requests.MonthAmountVoucherRequest) ([]*record.MonthAmountVoucherSuccessRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountVouchersSuccess(r.ctx, db.GetMonthAmountVouchersSuccessParams{
		Column1: currentDate,
		Column2: lastDayCurrentMonth,
		Column3: prevDate,
		Column4: lastDayPrevMonth,
	})

	if err != nil {
		return nil, voucher_errors.ErrFindMonthAmountVoucherSuccess
	}

	so := r.mapping.ToVouchersRecordMonthAmountSuccess(res)

	return so, nil
}

func (r *voucherRepository) FindYearAmountVoucherSuccess(year int) ([]*record.YearAmountVoucherSuccessRecord, error) {
	res, err := r.db.GetYearAmountVouchersSuccess(r.ctx, int32(year))

	if err != nil {
		return nil, voucher_errors.ErrFindYearAmountVoucherSuccess
	}

	so := r.mapping.ToVouchersRecordYearAmountSuccess(res)

	return so, nil
}

func (r *voucherRepository) FindMonthAmountVoucherFailed(req *requests.MonthAmountVoucherRequest) ([]*record.MonthAmountVoucherFailedRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)

	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountVouchersFailed(r.ctx, db.GetMonthAmountVouchersFailedParams{
		Column1: currentDate,
		Column2: lastDayCurrentMonth,
		Column3: prevDate,
		Column4: lastDayPrevMonth,
	})

	if err != nil {
		return nil, voucher_errors.ErrFindMonthAmountVoucherFailed
	}

	so := r.mapping.ToVouchersRecordMonthAmountFailed(res)

	return so, nil
}

func (r *voucherRepository) FindYearAmountVoucherFailed(year int) ([]*record.YearAmountVoucherFailedRecord, error) {
	res, err := r.db.GetYearAmountVouchersFailed(r.ctx, int32(year))

	if err != nil {
		return nil, voucher_errors.ErrFindYearAmountVoucherFailed
	}

	so := r.mapping.ToVouchersRecordYearAmountFailed(res)

	return so, nil
}

func (r *voucherRepository) FindMonthMethodVoucherSuccess(year int) ([]*record.MonthMethodVoucherRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthMethodVouchersSuccess(r.ctx, yearStart)

	if err != nil {
		return nil, voucher_errors.ErrFindMonthMethodVoucherSuccess
	}

	so := r.mapping.ToVouchersRecordMonthMethodSuccess(res)

	return so, nil
}

func (r *voucherRepository) FindYearMethodVoucherSuccess(year int) ([]*record.YearMethodVoucherRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearMethodVouchersSuccess(r.ctx, yearStart)

	if err != nil {
		return nil, voucher_errors.ErrFindYearMethodVoucherSuccess
	}

	so := r.mapping.ToVouchersRecordYearMethodSuccess(res)

	return so, nil
}

func (r *voucherRepository) FindMonthMethodVoucherFailed(year int) ([]*record.MonthMethodVoucherRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthMethodVouchersFailed(r.ctx, yearStart)

	if err != nil {
		return nil, voucher_errors.ErrFindMonthMethodVoucherFailed
	}

	so := r.mapping.ToVouchersRecordMonthMethodFailed(res)

	return so, nil
}

func (r *voucherRepository) FindYearMethodVoucherFailed(year int) ([]*record.YearMethodVoucherRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearMethodVouchersFailed(r.ctx, yearStart)

	if err != nil {
		return nil, voucher_errors.ErrFindYearMethodVoucherFailed
	}

	so := r.mapping.ToVouchersRecordYearMethodFailed(res)

	return so, nil
}

func (r *voucherRepository) FindMonthAmountVoucherSuccessById(req *requests.MonthAmountVoucherByIdRequest) ([]*record.MonthAmountVoucherSuccessRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountVouchersSuccessById(r.ctx, db.GetMonthAmountVouchersSuccessByIdParams{
		Column1:   currentDate,
		Column2:   lastDayCurrentMonth,
		Column3:   prevDate,
		Column4:   lastDayPrevMonth,
		VoucherID: int32(req.ID),
	})

	if err != nil {
		return nil, voucher_errors.ErrFindMonthAmountVoucherSuccessById
	}

	so := r.mapping.ToVouchersRecordMonthAmountSuccessById(res)

	return so, nil
}

func (r *voucherRepository) FindYearAmountVoucherSuccessById(req *requests.YearAmountVoucherByIdRequest) ([]*record.YearAmountVoucherSuccessRecord, error) {
	res, err := r.db.GetYearAmountVouchersSuccessById(r.ctx, db.GetYearAmountVouchersSuccessByIdParams{
		Column1:   int32(req.Year),
		VoucherID: int32(req.ID),
	})

	if err != nil {
		return nil, voucher_errors.ErrFindYearAmountVoucherSuccessById
	}

	so := r.mapping.ToVouchersRecordYearAmountSuccessById(res)

	return so, nil
}

func (r *voucherRepository) FindMonthAmountVoucherFailedById(req *requests.MonthAmountVoucherByIdRequest) ([]*record.MonthAmountVoucherFailedRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountVouchersFailedById(r.ctx, db.GetMonthAmountVouchersFailedByIdParams{
		Column1:   currentDate,
		Column2:   lastDayCurrentMonth,
		Column3:   prevDate,
		Column4:   lastDayPrevMonth,
		VoucherID: int32(req.ID),
	})

	if err != nil {
		return nil, voucher_errors.ErrFindMonthAmountVoucherFailedById
	}

	so := r.mapping.ToVouchersRecordMonthAmountFailedById(res)

	return so, nil
}

func (r *voucherRepository) FindYearAmountVoucherFailedById(req *requests.YearAmountVoucherByIdRequest) ([]*record.YearAmountVoucherFailedRecord, error) {
	res, err := r.db.GetYearAmountVouchersFailedById(r.ctx, db.GetYearAmountVouchersFailedByIdParams{
		Column1:   int32(req.Year),
		VoucherID: int32(req.ID),
	})

	if err != nil {
		return nil, voucher_errors.ErrFindYearAmountVoucherFailedById
	}

	so := r.mapping.ToVouchersRecordYearAmountFailedById(res)

	return so, nil
}

func (r *voucherRepository) FindMonthMethodVoucherSuccessById(req *requests.MonthMethodVoucherByIdRequest) ([]*record.MonthMethodVoucherRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthMethodVouchersSuccessById(r.ctx, db.GetMonthMethodVouchersSuccessByIdParams{
		Column1:   yearStart,
		VoucherID: int32(req.ID),
	})

	if err != nil {
		return nil, voucher_errors.ErrFindMonthMethodVoucherSuccessById
	}

	so := r.mapping.ToVouchersRecordMonthMethodSuccessById(res)

	return so, nil
}

func (r *voucherRepository) FindYearMethodVoucherSuccessById(req *requests.YearMethodVoucherByIdRequest) ([]*record.YearMethodVoucherRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearMethodVouchersSuccessById(r.ctx, db.GetYearMethodVouchersSuccessByIdParams{
		Column1:   yearStart,
		VoucherID: int32(req.ID),
	})

	if err != nil {
		return nil, voucher_errors.ErrFindYearMethodVoucherSuccessById
	}

	so := r.mapping.ToVouchersRecordYearMethodSuccessById(res)

	return so, nil
}

func (r *voucherRepository) FindMonthMethodVoucherFailedById(req *requests.MonthMethodVoucherByIdRequest) ([]*record.MonthMethodVoucherRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthMethodVouchersFailedById(r.ctx, db.GetMonthMethodVouchersFailedByIdParams{
		Column1:   yearStart,
		VoucherID: int32(req.ID),
	})

	if err != nil {
		return nil, voucher_errors.ErrFindYearMethodVoucherFailedById
	}

	so := r.mapping.ToVouchersRecordMonthMethodFailedById(res)

	return so, nil
}

func (r *voucherRepository) FindYearMethodVoucherFailedById(req *requests.YearMethodVoucherByIdRequest) ([]*record.YearMethodVoucherRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearMethodVouchersFailedById(r.ctx, db.GetYearMethodVouchersFailedByIdParams{
		Column1:   yearStart,
		VoucherID: int32(req.ID),
	})

	if err != nil {
		return nil, voucher_errors.ErrFindYearMethodVoucherFailedById
	}

	so := r.mapping.ToVouchersRecordYearMethodFailedById(res)

	return so, nil
}

func (r *voucherRepository) FindMonthAmountVoucherSuccessByMerchant(req *requests.MonthAmountVoucherByMerchantRequest) ([]*record.MonthAmountVoucherSuccessRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountVouchersSuccessByMerchant(r.ctx, db.GetMonthAmountVouchersSuccessByMerchantParams{
		Column1:    currentDate,
		Column2:    lastDayCurrentMonth,
		Column3:    prevDate,
		Column4:    lastDayPrevMonth,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, voucher_errors.ErrFindMonthAmountVoucherSuccessByMerchant
	}

	so := r.mapping.ToVouchersRecordMonthAmountSuccessByMerchant(res)

	return so, nil
}

func (r *voucherRepository) FindYearAmountVoucherSuccessByMerchant(req *requests.YearAmountVoucherByMerchantRequest) ([]*record.YearAmountVoucherSuccessRecord, error) {
	res, err := r.db.GetYearAmountVouchersSuccessByMerchant(r.ctx, db.GetYearAmountVouchersSuccessByMerchantParams{
		Column1:    int32(req.Year),
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, voucher_errors.ErrFindYearAmountVoucherSuccessByMerchant
	}

	so := r.mapping.ToVouchersRecordYearAmountSuccessByMerchant(res)

	return so, nil
}

func (r *voucherRepository) FindMonthAmountVoucherFailedByMerchant(req *requests.MonthAmountVoucherByMerchantRequest) ([]*record.MonthAmountVoucherFailedRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountVouchersFailedByMerchant(r.ctx, db.GetMonthAmountVouchersFailedByMerchantParams{
		Column1:    currentDate,
		Column2:    lastDayCurrentMonth,
		Column3:    prevDate,
		Column4:    lastDayPrevMonth,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, voucher_errors.ErrFindMonthAmountVoucherFailedByMerchant
	}

	so := r.mapping.ToVouchersRecordMonthAmountFailedByMerchant(res)

	return so, nil
}

func (r *voucherRepository) FindYearAmountVoucherFailedByMerchant(req *requests.YearAmountVoucherByMerchantRequest) ([]*record.YearAmountVoucherFailedRecord, error) {
	res, err := r.db.GetYearAmountVouchersFailedByMerchant(r.ctx, db.GetYearAmountVouchersFailedByMerchantParams{
		Column1:    int32(req.Year),
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, voucher_errors.ErrFindYearAmountVoucherFailedByMerchant
	}

	so := r.mapping.ToVouchersRecordYearAmountFailedByMerchant(res)

	return so, nil
}

func (r *voucherRepository) FindMonthMethodVoucherSuccessByMerchant(req *requests.MonthMethodVoucherByMerchantRequest) ([]*record.MonthMethodVoucherRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthMethodVouchersSuccessByMerchant(r.ctx, db.GetMonthMethodVouchersSuccessByMerchantParams{
		Column1:    yearStart,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, voucher_errors.ErrFindMonthMethodVoucherSuccessByMerchant
	}

	so := r.mapping.ToVouchersRecordMonthMethodSuccessByMerchant(res)

	return so, nil
}

func (r *voucherRepository) FindYearMethodVoucherSuccessByMerchant(req *requests.YearMethodVoucherByMerchantRequest) ([]*record.YearMethodVoucherRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearMethodVouchersSuccessByMerchant(r.ctx, db.GetYearMethodVouchersSuccessByMerchantParams{
		Column1:    yearStart,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, voucher_errors.ErrFindYearMethodVoucherSuccessByMerchant
	}

	so := r.mapping.ToVouchersRecordYearMethodSuccessByMerchant(res)

	return so, nil
}

func (r *voucherRepository) FindMonthMethodVoucherFailedByMerchant(req *requests.MonthMethodVoucherByMerchantRequest) ([]*record.MonthMethodVoucherRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthMethodVouchersFailedByMerchant(r.ctx, db.GetMonthMethodVouchersFailedByMerchantParams{
		Column1:    yearStart,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, voucher_errors.ErrFindYearMethodVoucherFailedByMerchant
	}

	so := r.mapping.ToVouchersRecordMonthMethodFailedByMerchant(res)

	return so, nil
}

func (r *voucherRepository) FindYearMethodVoucherFailedByMerchant(req *requests.YearMethodVoucherByMerchantRequest) ([]*record.YearMethodVoucherRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearMethodVouchersFailedByMerchant(r.ctx, db.GetYearMethodVouchersFailedByMerchantParams{
		Column1:    yearStart,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, voucher_errors.ErrFindYearMethodVoucherFailedByMerchant
	}

	so := r.mapping.ToVouchersRecordYearMethodFailedByMerchant(res)

	return so, nil
}

func (r *voucherRepository) FindById(id int) (*record.VoucherRecord, error) {
	res, err := r.db.GetVoucherByID(r.ctx, int32(id))

	if err != nil {
		return nil, voucher_errors.ErrVoucherNotFound
	}

	return r.mapping.ToVoucherRecord(res), nil
}

func (r *voucherRepository) FindByActiveVouchers(request *requests.FindAllVouchers) ([]*record.VoucherRecord, *int, error) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	offset := (page - 1) * pageSize

	req := db.GetVouchersActiveParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetVouchersActive(r.ctx, req)

	if err != nil {
		return nil, nil, voucher_errors.ErrFindActiveVouchers
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToVouchersRecordActive(res), &totalCount, nil
}

func (r *voucherRepository) FindByTrashedVoucher(request *requests.FindAllVouchers) ([]*record.VoucherRecord, *int, error) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	offset := (page - 1) * pageSize

	req := db.GetVouchersTrashedParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetVouchersTrashed(r.ctx, req)

	if err != nil {
		return nil, nil, voucher_errors.ErrFindTrashedVouchers
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToVouchersRecordTrashed(res), &totalCount, nil
}

func (r *voucherRepository) CreateVoucher(req *requests.CreateVoucherRequest) (*record.VoucherRecord, error) {
	res, err := r.db.CreateVoucher(r.ctx, db.CreateVoucherParams{
		MerchantID: int32(req.MerchantID),
		CategoryID: int32(req.CategoryID),
		Name:       req.Name,
		ImageName:  req.ImageName,
	})

	if err != nil {
		return nil, voucher_errors.ErrCreateVoucher
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
		return nil, voucher_errors.ErrUpdateVoucher
	}

	return r.mapping.ToVoucherRecord(res), nil
}

func (r *voucherRepository) TrashVoucher(id int) (*record.VoucherRecord, error) {
	res, err := r.db.TrashVoucher(r.ctx, int32(id))

	if err != nil {
		return nil, voucher_errors.ErrTrashedVoucher
	}

	return r.mapping.ToVoucherRecord(res), nil
}

func (r *voucherRepository) RestoreVoucher(id int) (*record.VoucherRecord, error) {
	res, err := r.db.RestoreVoucher(r.ctx, int32(id))

	if err != nil {
		return nil, voucher_errors.ErrRestoreVoucher
	}

	return r.mapping.ToVoucherRecord(res), nil
}

func (r *voucherRepository) DeleteVoucherPermanent(category_id int) (bool, error) {
	err := r.db.DeleteVoucherPermanently(r.ctx, int32(category_id))

	if err != nil {
		return false, voucher_errors.ErrDeleteVoucherPermanent
	}

	return true, nil
}

func (r *voucherRepository) RestoreAllVouchers() (bool, error) {
	err := r.db.RestoreAllVouchers(r.ctx)

	if err != nil {
		return false, voucher_errors.ErrRestoreAllVouchers
	}

	return true, nil
}

func (r *voucherRepository) DeleteAllVouchersPermanent() (bool, error) {
	err := r.db.DeleteAllPermanentVouchers(r.ctx)

	if err != nil {
		return false, voucher_errors.ErrDeleteAllVouchers
	}

	return true, nil
}
