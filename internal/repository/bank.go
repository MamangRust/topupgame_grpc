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
	"topup_game/pkg/errors/bank_errors"
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

func (r *bankRepository) FindAllBanks(request *requests.FindAllBanks) ([]*record.BankRecord, *int, error) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	offset := (page - 1) * pageSize

	req := db.GetBanksParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetBanks(r.ctx, req)
	if err != nil {
		return nil, nil, bank_errors.ErrFindAllBanks
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToBanksRecordAll(res), &totalCount, nil
}

func (r *bankRepository) FindMonthAmountBankSuccess(req *requests.MonthAmountBankRequest) ([]*record.MonthAmountBankSuccessRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountBankSuccess(r.ctx, db.GetMonthAmountBankSuccessParams{
		Column1: currentDate,
		Column2: lastDayCurrentMonth,
		Column3: prevDate,
		Column4: lastDayPrevMonth,
	})

	if err != nil {
		return nil, bank_errors.ErrFindMonthAmountBankSuccess
	}

	so := r.mapping.ToBanksRecordMonthAmountSuccess(res)

	return so, nil
}

func (r *bankRepository) FindYearAmountBankSuccess(year int) ([]*record.YearAmountBankSuccessRecord, error) {
	res, err := r.db.GetYearAmountBankSuccess(r.ctx, int32(year))

	if err != nil {
		return nil, bank_errors.ErrFindYearAmountBankSuccess
	}

	so := r.mapping.ToBanksRecordYearAmountSuccess(res)

	return so, nil
}

func (r *bankRepository) FindMonthAmountBankFailed(req *requests.MonthAmountBankRequest) ([]*record.MonthAmountBankFailedRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)

	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountBankFailed(r.ctx, db.GetMonthAmountBankFailedParams{
		Column1: currentDate,
		Column2: lastDayCurrentMonth,
		Column3: prevDate,
		Column4: lastDayPrevMonth,
	})

	if err != nil {
		return nil, bank_errors.ErrFindMonthAmountBankFailed
	}

	so := r.mapping.ToBanksRecordMonthAmountFailed(res)

	return so, nil
}

func (r *bankRepository) FindYearAmountBankFailed(year int) ([]*record.YearAmountBankFailedRecord, error) {
	res, err := r.db.GetYearAmountBankFailed(r.ctx, int32(year))

	if err != nil {
		return nil, bank_errors.ErrFindYearAmountBankFailed
	}

	so := r.mapping.ToBanksRecordYearAmountFailed(res)

	return so, nil
}

func (r *bankRepository) FindMonthMethodBankSuccess(year int) ([]*record.MonthMethodBankRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthBankMethodsSuccess(r.ctx, yearStart)

	if err != nil {
		return nil, bank_errors.ErrFindMonthMethodBankSuccess
	}

	so := r.mapping.ToBanksRecordMonthMethodSuccess(res)

	return so, nil
}

func (r *bankRepository) FindYearMethodBankSuccess(year int) ([]*record.YearMethodBankRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearBankMethodsSuccess(r.ctx, yearStart)

	if err != nil {
		return nil, bank_errors.ErrFindYearMethodBankSuccess
	}

	so := r.mapping.ToBanksRecordYearMethodSuccess(res)

	return so, nil
}

func (r *bankRepository) FindMonthMethodBankFailed(year int) ([]*record.MonthMethodBankRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthBankMethodsFailed(r.ctx, yearStart)

	if err != nil {
		return nil, bank_errors.ErrFindMonthMethodBankFailed
	}

	so := r.mapping.ToBanksRecordMonthMethodFailed(res)

	return so, nil
}

func (r *bankRepository) FindYearMethodBankFailed(year int) ([]*record.YearMethodBankRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearBankMethodsFailed(r.ctx, yearStart)

	if err != nil {
		return nil, bank_errors.ErrFindYearMethodBankFailed
	}

	so := r.mapping.ToBanksRecordYearMethodFailed(res)

	return so, nil
}

func (r *bankRepository) FindMonthAmountBankSuccessById(req *requests.MonthAmountBankByIdRequest) ([]*record.MonthAmountBankSuccessRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountBankSuccessById(r.ctx, db.GetMonthAmountBankSuccessByIdParams{
		Column1: currentDate,
		Column2: lastDayCurrentMonth,
		Column3: prevDate,
		Column4: lastDayPrevMonth,
		BankID:  int32(req.ID),
	})

	if err != nil {
		return nil, bank_errors.ErrFindMonthAmountBankSuccessById
	}

	so := r.mapping.ToBanksRecordMonthAmountSuccessById(res)

	return so, nil
}

func (r *bankRepository) FindYearAmountBankSuccessById(req *requests.YearAmountBankByIdRequest) ([]*record.YearAmountBankSuccessRecord, error) {
	res, err := r.db.GetYearAmountBankSuccessById(r.ctx, db.GetYearAmountBankSuccessByIdParams{
		Column1: int32(req.Year),
		BankID:  int32(req.ID),
	})

	if err != nil {
		return nil, bank_errors.ErrFindYearAmountBankSuccessById
	}

	so := r.mapping.ToBanksRecordYearAmountSuccessById(res)

	return so, nil
}

func (r *bankRepository) FindMonthAmountBankFailedById(req *requests.MonthAmountBankByIdRequest) ([]*record.MonthAmountBankFailedRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountBankFailedById(r.ctx, db.GetMonthAmountBankFailedByIdParams{
		Column1: currentDate,
		Column2: lastDayCurrentMonth,
		Column3: prevDate,
		Column4: lastDayPrevMonth,
		BankID:  int32(req.ID),
	})

	if err != nil {
		return nil, bank_errors.ErrFindMonthAmountBankFailedById
	}

	so := r.mapping.ToBanksRecordMonthAmountFailedById(res)

	return so, nil
}

func (r *bankRepository) FindYearAmountBankFailedById(req *requests.YearAmountBankByIdRequest) ([]*record.YearAmountBankFailedRecord, error) {
	res, err := r.db.GetYearAmountBankFailedById(r.ctx, db.GetYearAmountBankFailedByIdParams{
		Column1: int32(req.Year),
		BankID:  int32(req.ID),
	})

	if err != nil {
		return nil, bank_errors.ErrFindYearAmountBankFailedById
	}

	so := r.mapping.ToBanksRecordYearAmountFaileById(res)

	return so, nil
}

func (r *bankRepository) FindMonthMethodBankSuccessById(req *requests.MonthMethodBankByIdRequest) ([]*record.MonthMethodBankRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthBankMethodsSuccessById(r.ctx, db.GetMonthBankMethodsSuccessByIdParams{
		Column1: yearStart,
		BankID:  int32(req.ID),
	})

	if err != nil {
		return nil, bank_errors.ErrFindMonthMethodBankSuccessById
	}

	so := r.mapping.ToBanksRecordMonthMethodSuccessById(res)

	return so, nil
}

func (r *bankRepository) FindYearMethodBankSuccessById(req *requests.YearMethodBankByIdRequest) ([]*record.YearMethodBankRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearBankMethodsSuccessById(r.ctx, db.GetYearBankMethodsSuccessByIdParams{
		Column1: yearStart,
		BankID:  int32(req.ID),
	})

	if err != nil {
		return nil, bank_errors.ErrFindYearMethodBankSuccessById
	}

	so := r.mapping.ToBanksRecordYearMethodSuccessById(res)

	return so, nil
}

func (r *bankRepository) FindMonthMethodBankFailedById(req *requests.MonthMethodBankByIdRequest) ([]*record.MonthMethodBankRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthBankMethodsFailedById(r.ctx, db.GetMonthBankMethodsFailedByIdParams{
		Column1: yearStart,
		BankID:  int32(req.ID),
	})

	if err != nil {
		return nil, bank_errors.ErrFindYearMethodBankFailedById
	}

	so := r.mapping.ToBanksRecordMonthMethodFailedById(res)

	return so, nil
}

func (r *bankRepository) FindYearMethodBankFailedById(req *requests.YearMethodBankByIdRequest) ([]*record.YearMethodBankRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearBankMethodsFailedById(r.ctx, db.GetYearBankMethodsFailedByIdParams{
		Column1: yearStart,
		BankID:  int32(req.ID),
	})

	if err != nil {
		return nil, bank_errors.ErrFindYearMethodBankFailedById
	}

	so := r.mapping.ToBanksRecordYearMethodFailedById(res)

	return so, nil
}

func (r *bankRepository) FindMonthAmountBankSuccessByMerchant(req *requests.MonthAmountBankByMerchantRequest) ([]*record.MonthAmountBankSuccessRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountBankSuccessByMerchant(r.ctx, db.GetMonthAmountBankSuccessByMerchantParams{
		Column1:    currentDate,
		Column2:    lastDayCurrentMonth,
		Column3:    prevDate,
		Column4:    lastDayPrevMonth,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, bank_errors.ErrFindMonthAmountBankSuccessByMerchant
	}

	so := r.mapping.ToBanksRecordMonthAmountSuccessByMerchant(res)

	return so, nil
}

func (r *bankRepository) FindYearAmountBankSuccessByMerchant(req *requests.YearAmountBankByMerchantRequest) ([]*record.YearAmountBankSuccessRecord, error) {
	res, err := r.db.GetYearAmountBankSuccessByMerchant(r.ctx, db.GetYearAmountBankSuccessByMerchantParams{
		Column1:    int32(req.Year),
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, bank_errors.ErrFindYearAmountBankSuccessByMerchant
	}

	so := r.mapping.ToBanksRecordYearAmountSuccessByMerchant(res)

	return so, nil
}

func (r *bankRepository) FindMonthAmountBankFailedByMerchant(req *requests.MonthAmountBankByMerchantRequest) ([]*record.MonthAmountBankFailedRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountBankFailedByMerchant(r.ctx, db.GetMonthAmountBankFailedByMerchantParams{
		Column1:    currentDate,
		Column2:    lastDayCurrentMonth,
		Column3:    prevDate,
		Column4:    lastDayPrevMonth,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, bank_errors.ErrFindMonthAmountBankFailedByMerchant
	}

	so := r.mapping.ToBanksRecordMonthAmountFailedByMerchant(res)

	return so, nil
}

func (r *bankRepository) FindYearAmountBankFailedByMerchant(req *requests.YearAmountBankByMerchantRequest) ([]*record.YearAmountBankFailedRecord, error) {
	res, err := r.db.GetYearAmountBankFailedByMerchant(r.ctx, db.GetYearAmountBankFailedByMerchantParams{
		Column1:    int32(req.Year),
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, bank_errors.ErrFindYearAmountBankFailedByMerchant
	}

	so := r.mapping.ToBanksRecordYearAmountFaileByMerchant(res)

	return so, nil
}

func (r *bankRepository) FindMonthMethodBankSuccessByMerchant(req *requests.MonthMethodBankByMerchantRequest) ([]*record.MonthMethodBankRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthBankMethodsSuccessByMerchant(r.ctx, db.GetMonthBankMethodsSuccessByMerchantParams{
		Column1:    yearStart,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, bank_errors.ErrFindMonthMethodBankSuccessByMerchant
	}

	so := r.mapping.ToBanksRecordMonthMethodSuccessByMerchant(res)

	return so, nil
}

func (r *bankRepository) FindYearMethodBankSuccessByMerchant(req *requests.YearMethodBankByMerchantRequest) ([]*record.YearMethodBankRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearBankMethodsSuccessByMerchant(r.ctx, db.GetYearBankMethodsSuccessByMerchantParams{
		Column1:    yearStart,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, bank_errors.ErrFindYearMethodBankSuccessByMerchant
	}

	so := r.mapping.ToBanksRecordYearMethodSuccessByMerchant(res)

	return so, nil
}

func (r *bankRepository) FindMonthMethodBankFailedByMerchant(req *requests.MonthMethodBankByMerchantRequest) ([]*record.MonthMethodBankRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthBankMethodsFailedByMerchant(r.ctx, db.GetMonthBankMethodsFailedByMerchantParams{
		Column1:    yearStart,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, bank_errors.ErrFindYearMethodBankFailedByMerchant
	}

	so := r.mapping.ToBanksRecordMonthMethodFailedByMerchant(res)

	return so, nil
}

func (r *bankRepository) FindYearMethodBankFailedByMerchant(req *requests.YearMethodBankByMerchantRequest) ([]*record.YearMethodBankRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearBankMethodsFailedByMerchant(r.ctx, db.GetYearBankMethodsFailedByMerchantParams{
		Column1:    yearStart,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, bank_errors.ErrFindYearMethodBankFailedByMerchant
	}

	so := r.mapping.ToBanksRecordYearMethodFailedByMerchant(res)

	return so, nil
}

func (r *bankRepository) FindById(id int) (*record.BankRecord, error) {
	res, err := r.db.GetBankByID(r.ctx, int32(id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, bank_errors.ErrBankNotFound
		}

		return nil, bank_errors.ErrBankNotFound
	}

	return r.mapping.ToBankRecord(res), nil
}

func (r *bankRepository) FindByActiveBanks(request *requests.FindAllBanks) ([]*record.BankRecord, *int, error) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	offset := (page - 1) * pageSize

	req := db.GetBanksActiveParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetBanksActive(r.ctx, req)

	if err != nil {
		return nil, nil, bank_errors.ErrFindActiveBanks
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToBanksRecordActive(res), &totalCount, nil
}

func (r *bankRepository) FindByTrashedBanks(request *requests.FindAllBanks) ([]*record.BankRecord, *int, error) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	offset := (page - 1) * pageSize

	req := db.GetBanksTrashedParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetBanksTrashed(r.ctx, req)

	if err != nil {
		return nil, nil, bank_errors.ErrFindTrashedBanks
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToBanksRecordTrashed(res), &totalCount, nil
}

func (r *bankRepository) CreateBank(req *requests.CreateBankRequest) (*record.BankRecord, error) {
	res, err := r.db.CreateBank(r.ctx, req.Name)

	if err != nil {
		return nil, bank_errors.ErrCreateBank
	}

	return r.mapping.ToBankRecord(res), nil
}

func (r *bankRepository) UpdateBank(req *requests.UpdateBankRequest) (*record.BankRecord, error) {
	res, err := r.db.UpdateBank(r.ctx, db.UpdateBankParams{
		BankID: int32(req.ID),
		Name:   req.Name,
	})

	if err != nil {
		return nil, bank_errors.ErrUpdateBank
	}

	return r.mapping.ToBankRecord(res), nil
}

func (r *bankRepository) TrashedBank(id int) (*record.BankRecord, error) {
	res, err := r.db.TrashBank(r.ctx, int32(id))

	if err != nil {
		return nil, bank_errors.ErrTrashedBank
	}

	return r.mapping.ToBankRecord(res), nil
}

func (r *bankRepository) RestoreBank(id int) (*record.BankRecord, error) {
	res, err := r.db.RestoreBank(r.ctx, int32(id))

	if err != nil {
		return nil, bank_errors.ErrRestoreBank
	}

	return r.mapping.ToBankRecord(res), nil
}

func (r *bankRepository) DeleteBankPermanent(bank_id int) (bool, error) {
	err := r.db.DeleteBankPermanently(r.ctx, int32(bank_id))

	if err != nil {
		return false, bank_errors.ErrDeleteBankPermanent
	}

	return true, nil
}

func (r *bankRepository) RestoreAllBanks() (bool, error) {
	err := r.db.RestoreAllBanks(r.ctx)

	if err != nil {
		return false, bank_errors.ErrRestoreAllBanks
	}

	return true, nil
}

func (r *bankRepository) DeleteAllBanksPermanent() (bool, error) {
	err := r.db.DeleteAllPermanentBanks(r.ctx)

	if err != nil {
		return false, bank_errors.ErrDeleteAllBanks
	}

	return true, nil
}
