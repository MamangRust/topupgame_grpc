package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"topup_game/internal/domain/record"
	"topup_game/internal/domain/requests"
	recordmapper "topup_game/internal/mapper/record"
	db "topup_game/pkg/database/schema"
	"topup_game/pkg/errors/nominal_errors"
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

func (r *nominalRepository) FindAllNominals(request *requests.FindAllNominals) ([]*record.NominalRecord, *int, error) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	offset := (page - 1) * pageSize

	req := db.GetNominalsParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetNominals(r.ctx, req)
	if err != nil {
		return nil, nil, nominal_errors.ErrFindAllNominals
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToNominalRecordsAll(res), &totalCount, nil
}

func (r *nominalRepository) FindMonthAmountNominalSuccess(req *requests.MonthAmountNominalRequest) ([]*record.MonthAmountNominalSuccessRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountNominalsSuccess(r.ctx, db.GetMonthAmountNominalsSuccessParams{
		Column1: currentDate,
		Column2: lastDayCurrentMonth,
		Column3: prevDate,
		Column4: lastDayPrevMonth,
	})

	if err != nil {
		return nil, nominal_errors.ErrFindMonthAmountNominalSuccess
	}

	so := r.mapping.ToNominalsRecordMonthAmountSuccess(res)

	return so, nil
}

func (r *nominalRepository) FindYearAmountNominalSuccess(year int) ([]*record.YearAmountNominalSuccessRecord, error) {
	res, err := r.db.GetYearAmountNominalsSuccess(r.ctx, int32(year))

	if err != nil {
		return nil, nominal_errors.ErrFindYearAmountNominalSuccess
	}

	so := r.mapping.ToNominalsRecordYearAmountSuccess(res)

	return so, nil
}

func (r *nominalRepository) FindMonthAmountNominalFailed(req *requests.MonthAmountNominalRequest) ([]*record.MonthAmountNominalFailedRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)

	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountNominalsFailed(r.ctx, db.GetMonthAmountNominalsFailedParams{
		Column1: currentDate,
		Column2: lastDayCurrentMonth,
		Column3: prevDate,
		Column4: lastDayPrevMonth,
	})

	if err != nil {
		return nil, nominal_errors.ErrFindMonthAmountNominalFailed
	}

	so := r.mapping.ToNominalsRecordMonthAmountFailed(res)

	return so, nil
}

func (r *nominalRepository) FindYearAmountNominalFailed(year int) ([]*record.YearAmountNominalFailedRecord, error) {
	res, err := r.db.GetYearAmountNominalsFailed(r.ctx, int32(year))

	if err != nil {
		return nil, nominal_errors.ErrFindYearAmountNominalFailed
	}

	so := r.mapping.ToNominalsRecordYearAmountFailed(res)

	return so, nil
}

func (r *nominalRepository) FindMonthMethodNominalSuccess(year int) ([]*record.MonthMethodNominalRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthMethodNominalsSuccess(r.ctx, yearStart)

	if err != nil {
		return nil, nominal_errors.ErrFindMonthMethodNominalSuccess
	}

	so := r.mapping.ToNominalsRecordMonthMethodSuccess(res)

	return so, nil
}

func (r *nominalRepository) FindYearMethodNominalSuccess(year int) ([]*record.YearMethodNominalRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearMethodNominalsSuccess(r.ctx, yearStart)

	if err != nil {
		return nil, nominal_errors.ErrFindYearMethodNominalSuccess
	}

	so := r.mapping.ToNominalsRecordYearMethodSuccess(res)

	return so, nil
}

func (r *nominalRepository) FindMonthMethodNominalFailed(year int) ([]*record.MonthMethodNominalRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthMethodNominalsFailed(r.ctx, yearStart)

	if err != nil {
		return nil, nominal_errors.ErrFindMonthMethodNominalFailed
	}

	so := r.mapping.ToNominalsRecordMonthMethodFailed(res)

	return so, nil
}

func (r *nominalRepository) FindYearMethodNominalFailed(year int) ([]*record.YearMethodNominalRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearMethodNominalsFailed(r.ctx, yearStart)

	if err != nil {
		return nil, nominal_errors.ErrFindYearMethodNominalFailed
	}

	so := r.mapping.ToNominalsRecordYearMethodFailed(res)

	return so, nil
}

func (r *nominalRepository) FindMonthAmountNominalSuccessById(req *requests.MonthAmountNominalByIdRequest) ([]*record.MonthAmountNominalSuccessRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountNominalsSuccessById(r.ctx, db.GetMonthAmountNominalsSuccessByIdParams{
		Column1:   currentDate,
		Column2:   lastDayCurrentMonth,
		Column3:   prevDate,
		Column4:   lastDayPrevMonth,
		NominalID: int32(req.ID),
	})

	if err != nil {
		return nil, nominal_errors.ErrFindMonthAmountNominalSuccessById
	}

	so := r.mapping.ToNominalsRecordMonthAmountSuccessById(res)

	return so, nil
}

func (r *nominalRepository) FindYearAmountNominalSuccessById(req *requests.YearAmountNominalByIdRequest) ([]*record.YearAmountNominalSuccessRecord, error) {
	res, err := r.db.GetYearAmountNominalsSuccessById(r.ctx, db.GetYearAmountNominalsSuccessByIdParams{
		Column1:   int32(req.Year),
		NominalID: int32(req.ID),
	})

	if err != nil {
		return nil, nominal_errors.ErrFindYearAmountNominalSuccessById
	}

	so := r.mapping.ToNominalsRecordYearAmountSuccessById(res)

	return so, nil
}

func (r *nominalRepository) FindMonthAmountNominalFailedById(req *requests.MonthAmountNominalByIdRequest) ([]*record.MonthAmountNominalFailedRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountNominalsFailedById(r.ctx, db.GetMonthAmountNominalsFailedByIdParams{
		Column1:   currentDate,
		Column2:   lastDayCurrentMonth,
		Column3:   prevDate,
		Column4:   lastDayPrevMonth,
		NominalID: int32(req.ID),
	})

	if err != nil {
		return nil, nominal_errors.ErrFindMonthAmountNominalFailedById
	}

	so := r.mapping.ToNominalsRecordMonthAmountFailedById(res)

	return so, nil
}

func (r *nominalRepository) FindYearAmountNominalFailedById(req *requests.YearAmountNominalByIdRequest) ([]*record.YearAmountNominalFailedRecord, error) {
	res, err := r.db.GetYearAmountNominalsFailedById(r.ctx, db.GetYearAmountNominalsFailedByIdParams{
		Column1:   int32(req.Year),
		NominalID: int32(req.ID),
	})

	if err != nil {
		return nil, nominal_errors.ErrFindYearAmountNominalFailedById
	}

	so := r.mapping.ToNominalsRecordYearAmountFaileById(res)

	return so, nil
}

func (r *nominalRepository) FindMonthMethodNominalSuccessById(req *requests.MonthMethodNominalByIdRequest) ([]*record.MonthMethodNominalRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthMethodNominalsSuccessById(r.ctx, db.GetMonthMethodNominalsSuccessByIdParams{
		Column1:   yearStart,
		NominalID: int32(req.ID),
	})

	if err != nil {
		return nil, nominal_errors.ErrFindMonthMethodNominalSuccessById
	}

	so := r.mapping.ToNominalsRecordMonthMethodSuccessById(res)

	return so, nil
}

func (r *nominalRepository) FindYearMethodNominalSuccessById(req *requests.YearMethodNominalByIdRequest) ([]*record.YearMethodNominalRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearMethodNominalsSuccessById(r.ctx, db.GetYearMethodNominalsSuccessByIdParams{
		Column1:   yearStart,
		NominalID: int32(req.ID),
	})

	if err != nil {
		return nil, nominal_errors.ErrFindYearMethodNominalSuccessById
	}

	so := r.mapping.ToNominalsRecordYearMethodSuccessById(res)

	return so, nil
}

func (r *nominalRepository) FindMonthMethodNominalFailedById(req *requests.MonthMethodNominalByIdRequest) ([]*record.MonthMethodNominalRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthMethodNominalsFailedById(r.ctx, db.GetMonthMethodNominalsFailedByIdParams{
		Column1:   yearStart,
		NominalID: int32(req.ID),
	})

	if err != nil {
		return nil, nominal_errors.ErrFindYearMethodNominalFailedById
	}

	so := r.mapping.ToNominalsRecordMonthMethodFailedById(res)

	return so, nil
}

func (r *nominalRepository) FindYearMethodNominalFailedById(req *requests.YearMethodNominalByIdRequest) ([]*record.YearMethodNominalRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearMethodNominalsFailedById(r.ctx, db.GetYearMethodNominalsFailedByIdParams{
		Column1:   yearStart,
		NominalID: int32(req.ID),
	})

	if err != nil {
		return nil, nominal_errors.ErrFindYearMethodNominalFailedById
	}

	so := r.mapping.ToNominalsRecordYearMethodFailedById(res)

	return so, nil
}

func (r *nominalRepository) FindMonthAmountNominalSuccessByMerchant(req *requests.MonthAmountNominalByMerchantRequest) ([]*record.MonthAmountNominalSuccessRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountNominalsSuccessByMerchant(r.ctx, db.GetMonthAmountNominalsSuccessByMerchantParams{
		Column1:    currentDate,
		Column2:    lastDayCurrentMonth,
		Column3:    prevDate,
		Column4:    lastDayPrevMonth,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, nominal_errors.ErrFindMonthAmountNominalSuccessByMerchant
	}

	so := r.mapping.ToNominalsRecordMonthAmountSuccessByMerchant(res)

	return so, nil
}

func (r *nominalRepository) FindYearAmountNominalSuccessByMerchant(req *requests.YearAmountNominalByMerchantRequest) ([]*record.YearAmountNominalSuccessRecord, error) {
	res, err := r.db.GetYearAmountNominalsSuccessByMerchant(r.ctx, db.GetYearAmountNominalsSuccessByMerchantParams{
		Column1:    int32(req.Year),
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, nominal_errors.ErrFindYearAmountNominalSuccessByMerchant
	}

	so := r.mapping.ToNominalsRecordYearAmountSuccessByMerchant(res)

	return so, nil
}

func (r *nominalRepository) FindMonthAmountNominalFailedByMerchant(req *requests.MonthAmountNominalByMerchantRequest) ([]*record.MonthAmountNominalFailedRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountNominalsFailedByMerchant(r.ctx, db.GetMonthAmountNominalsFailedByMerchantParams{
		Column1:    currentDate,
		Column2:    lastDayCurrentMonth,
		Column3:    prevDate,
		Column4:    lastDayPrevMonth,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, nominal_errors.ErrFindMonthAmountNominalFailedByMerchant
	}

	so := r.mapping.ToNominalsRecordMonthAmountFailedByMerchant(res)

	return so, nil
}

func (r *nominalRepository) FindYearAmountNominalFailedByMerchant(req *requests.YearAmountNominalByMerchantRequest) ([]*record.YearAmountNominalFailedRecord, error) {
	res, err := r.db.GetYearAmountNominalsFailedByMerchant(r.ctx, db.GetYearAmountNominalsFailedByMerchantParams{
		Column1:    int32(req.Year),
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, nominal_errors.ErrFindYearAmountNominalFailedByMerchant
	}

	so := r.mapping.ToNominalsRecordYearAmountFaileByMerchant(res)

	return so, nil
}

func (r *nominalRepository) FindMonthMethodNominalSuccessByMerchant(req *requests.MonthMethodNominalByMerchantRequest) ([]*record.MonthMethodNominalRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthMethodNominalsSuccessByMerchant(r.ctx, db.GetMonthMethodNominalsSuccessByMerchantParams{
		Column1:    yearStart,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, nominal_errors.ErrFindMonthMethodNominalSuccessByMerchant
	}

	so := r.mapping.ToNominalsRecordMonthMethodSuccessByMerchant(res)

	return so, nil
}

func (r *nominalRepository) FindYearMethodNominalSuccessByMerchant(req *requests.YearMethodNominalByMerchantRequest) ([]*record.YearMethodNominalRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearMethodNominalsSuccessByMerchant(r.ctx, db.GetYearMethodNominalsSuccessByMerchantParams{
		Column1:    yearStart,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, nominal_errors.ErrFindYearMethodNominalSuccessByMerchant
	}

	so := r.mapping.ToNominalsRecordYearMethodSuccessByMerchant(res)

	return so, nil
}

func (r *nominalRepository) FindMonthMethodNominalFailedByMerchant(req *requests.MonthMethodNominalByMerchantRequest) ([]*record.MonthMethodNominalRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthMethodNominalsFailedByMerchant(r.ctx, db.GetMonthMethodNominalsFailedByMerchantParams{
		Column1:    yearStart,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, nominal_errors.ErrFindYearMethodNominalFailedByMerchant
	}

	so := r.mapping.ToNominalsRecordMonthMethodFailedByMerchant(res)

	return so, nil
}

func (r *nominalRepository) FindYearMethodNominalFailedByMerchant(req *requests.YearMethodNominalByMerchantRequest) ([]*record.YearMethodNominalRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearMethodNominalsFailedByMerchant(r.ctx, db.GetYearMethodNominalsFailedByMerchantParams{
		Column1:    yearStart,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, nominal_errors.ErrFindYearMethodNominalFailedByMerchant
	}

	so := r.mapping.ToNominalsRecordYearMethodFailedByMerchant(res)

	return so, nil
}

func (r *nominalRepository) FindById(id int) (*record.NominalRecord, error) {
	res, err := r.db.GetNominalByID(r.ctx, int32(id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nominal_errors.ErrNominalNotFound
		}

		return nil, nominal_errors.ErrNominalNotFound
	}

	return r.mapping.ToNominalRecord(res), nil
}

func (r *nominalRepository) FindByActiveNominal(request *requests.FindAllNominals) ([]*record.NominalRecord, *int, error) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	offset := (page - 1) * pageSize

	req := db.GetNominalsActiveParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetNominalsActive(r.ctx, req)

	if err != nil {
		return nil, nil, nominal_errors.ErrFindActiveNominals
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToNominalRecordsActive(res), &totalCount, nil
}

func (r *nominalRepository) FindByTrashedNominal(request *requests.FindAllNominals) ([]*record.NominalRecord, *int, error) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	offset := (page - 1) * pageSize

	req := db.GetNominalsTrashedParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetNominalsTrashed(r.ctx, req)

	if err != nil {
		return nil, nil, nominal_errors.ErrFindTrashedNominals
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToNominalRecordsTrashed(res), &totalCount, nil
}

func (r *nominalRepository) CreateNominal(req *requests.CreateNominalRequest) (*record.NominalRecord, error) {
	res, err := r.db.CreateNominal(r.ctx, db.CreateNominalParams{
		VoucherID: int32(req.VoucherID),
		Name:      req.Name,
		Quantity:  int32(req.Quantity),
		Price:     req.Price,
	})

	if err != nil {
		return nil, nominal_errors.ErrCreateNominal
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
		return nil, nominal_errors.ErrUpdateNominal
	}

	return r.mapping.ToNominalRecord(res), nil
}

func (r *nominalRepository) UpdateQuantity(nominal int, quantity int) (bool, error) {
	err := r.db.DecreaseNominalQuantity(r.ctx, db.DecreaseNominalQuantityParams{
		Quantity:  int32(quantity),
		NominalID: int32(nominal),
	})

	if err != nil {
		return false, nominal_errors.ErrUpdateNominal
	}

	return true, nil
}

func (r *nominalRepository) TrashedNominal(id int) (*record.NominalRecord, error) {
	res, err := r.db.TrashNominal(r.ctx, int32(id))

	if err != nil {
		return nil, nominal_errors.ErrTrashedNominal
	}

	return r.mapping.ToNominalRecord(res), nil
}

func (r *nominalRepository) RestoreNominal(id int) (*record.NominalRecord, error) {
	res, err := r.db.RestoreNominal(r.ctx, int32(id))

	if err != nil {
		return nil, nominal_errors.ErrRestoreNominal
	}

	return r.mapping.ToNominalRecord(res), nil
}

func (r *nominalRepository) DeleteNominalPermanent(nominal_id int) (bool, error) {
	err := r.db.DeleteNominalPermanently(r.ctx, int32(nominal_id))

	if err != nil {
		return false, nominal_errors.ErrDeleteNominalPermanent
	}

	return true, nil
}

func (r *nominalRepository) RestoreAllNominal() (bool, error) {
	err := r.db.RestoreAllNominals(r.ctx)

	if err != nil {
		return false, nominal_errors.ErrRestoreAllNominals
	}

	return true, nil
}

func (r *nominalRepository) DeleteAllNominalsPermanent() (bool, error) {
	err := r.db.DeleteAllPermanentNominals(r.ctx)

	if err != nil {
		return false, nominal_errors.ErrDeleteAllNominals
	}

	return true, nil
}
