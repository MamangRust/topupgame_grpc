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
	"topup_game/pkg/errors/category_errors"
)

type categoryRepository struct {
	db      *db.Queries
	ctx     context.Context
	mapping recordmapper.CategoryRecordMapping
}

func NewCategoryRepository(db *db.Queries, ctx context.Context, mapping recordmapper.CategoryRecordMapping) *categoryRepository {
	return &categoryRepository{
		db:      db,
		ctx:     ctx,
		mapping: mapping,
	}
}

func (r *categoryRepository) FindAllCategories(request *requests.FindAllCategory) ([]*record.CategoryRecord, *int, error) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	offset := (page - 1) * pageSize

	req := db.GetCategoriesParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetCategories(r.ctx, req)
	if err != nil {
		return nil, nil, category_errors.ErrFindAllCategories
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToCategoriesRecordAll(res), &totalCount, nil
}

func (r *categoryRepository) FindMonthAmountCategorySuccess(req *requests.MonthAmountCategoryRequest) ([]*record.MonthAmountCategorySuccessRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountCategorySuccess(r.ctx, db.GetMonthAmountCategorySuccessParams{
		Column1: currentDate,
		Column2: lastDayCurrentMonth,
		Column3: prevDate,
		Column4: lastDayPrevMonth,
	})

	if err != nil {
		return nil, category_errors.ErrFindMonthAmountCategorySuccess
	}

	so := r.mapping.ToCategoriesRecordMonthAmountSuccess(res)

	return so, nil
}

func (r *categoryRepository) FindYearAmountCategorySuccess(year int) ([]*record.YearAmountCategorySuccessRecord, error) {
	res, err := r.db.GetYearAmountCategorySuccess(r.ctx, int32(year))

	if err != nil {
		return nil, category_errors.ErrFindYearAmountCategorySuccess
	}

	so := r.mapping.ToCategoriesRecordYearAmountSuccess(res)

	return so, nil
}

func (r *categoryRepository) FindMonthAmountCategoryFailed(req *requests.MonthAmountCategoryRequest) ([]*record.MonthAmountCategoryFailedRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)

	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountCategoryFailed(r.ctx, db.GetMonthAmountCategoryFailedParams{
		Column1: currentDate,
		Column2: lastDayCurrentMonth,
		Column3: prevDate,
		Column4: lastDayPrevMonth,
	})

	if err != nil {
		return nil, category_errors.ErrFindMonthAmountCategoryFailed
	}

	so := r.mapping.ToCategoriesRecordMonthAmountFailed(res)

	return so, nil
}

func (r *categoryRepository) FindYearAmountCategoryFailed(year int) ([]*record.YearAmountCategoryFailedRecord, error) {
	res, err := r.db.GetYearAmountCategoryFailed(r.ctx, int32(year))

	if err != nil {
		return nil, category_errors.ErrFindYearAmountCategoryFailed
	}

	so := r.mapping.ToCategoriesRecordYearAmountFailed(res)

	return so, nil
}

func (r *categoryRepository) FindMonthMethodCategorySuccess(year int) ([]*record.MonthMethodCategoryRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthMethodCategoriesSuccess(r.ctx, yearStart)

	if err != nil {
		return nil, category_errors.ErrFindMonthMethodCategorySuccess
	}

	so := r.mapping.ToCategoriesRecordMonthMethodSuccess(res)

	return so, nil
}

func (r *categoryRepository) FindYearMethodCategorySuccess(year int) ([]*record.YearMethodCategoryRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearMethodCategoriesSuccess(r.ctx, yearStart)

	if err != nil {
		return nil, category_errors.ErrFindYearMethodCategorySuccess
	}

	so := r.mapping.ToCategoriesRecordYearMethodSuccess(res)

	return so, nil
}

func (r *categoryRepository) FindMonthMethodCategoryFailed(year int) ([]*record.MonthMethodCategoryRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthMethodCategoriesFailed(r.ctx, yearStart)

	if err != nil {
		return nil, category_errors.ErrFindMonthMethodCategoryFailed
	}

	so := r.mapping.ToCategoriesRecordMonthMethodFailed(res)

	return so, nil
}

func (r *categoryRepository) FindYearMethodCategoryFailed(year int) ([]*record.YearMethodCategoryRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearMethodCategoriesFailed(r.ctx, yearStart)

	if err != nil {
		return nil, category_errors.ErrFindYearMethodCategoryFailed
	}

	so := r.mapping.ToCategoriesRecordYearMethodFailed(res)

	return so, nil
}

func (r *categoryRepository) FindMonthAmountCategorySuccessById(req *requests.MonthAmountCategoryByIdRequest) ([]*record.MonthAmountCategorySuccessRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountCategorySuccessById(r.ctx, db.GetMonthAmountCategorySuccessByIdParams{
		Column1:    currentDate,
		Column2:    lastDayCurrentMonth,
		Column3:    prevDate,
		Column4:    lastDayPrevMonth,
		CategoryID: int32(req.ID),
	})

	if err != nil {
		return nil, category_errors.ErrFindMonthAmountCategorySuccessById
	}

	so := r.mapping.ToCategoriesRecordMonthAmountSuccessById(res)

	return so, nil
}

func (r *categoryRepository) FindYearAmountCategorySuccessById(req *requests.YearAmountCategoryByIdRequest) ([]*record.YearAmountCategorySuccessRecord, error) {
	res, err := r.db.GetYearAmountCategorySuccessById(r.ctx, db.GetYearAmountCategorySuccessByIdParams{
		Column1:    int32(req.Year),
		CategoryID: int32(req.ID),
	})

	if err != nil {
		return nil, category_errors.ErrFindYearAmountCategorySuccessById
	}

	so := r.mapping.ToCategoriesRecordYearAmountSuccessById(res)

	return so, nil
}

func (r *categoryRepository) FindMonthAmountCategoryFailedById(req *requests.MonthAmountCategoryByIdRequest) ([]*record.MonthAmountCategoryFailedRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountCategoryFailedById(r.ctx, db.GetMonthAmountCategoryFailedByIdParams{
		Column1:    currentDate,
		Column2:    lastDayCurrentMonth,
		Column3:    prevDate,
		Column4:    lastDayPrevMonth,
		CategoryID: int32(req.ID),
	})

	if err != nil {
		return nil, category_errors.ErrFindMonthAmountCategoryFailedById
	}

	so := r.mapping.ToCategoriesRecordMonthAmountFailedById(res)

	return so, nil
}

func (r *categoryRepository) FindYearAmountCategoryFailedById(req *requests.YearAmountCategoryByIdRequest) ([]*record.YearAmountCategoryFailedRecord, error) {
	res, err := r.db.GetYearAmountCategoryFailedById(r.ctx, db.GetYearAmountCategoryFailedByIdParams{
		Column1:    int32(req.Year),
		CategoryID: int32(req.ID),
	})

	if err != nil {
		return nil, category_errors.ErrFindYearAmountCategoryFailedById
	}

	so := r.mapping.ToCategoriesRecordYearAmountFaileById(res)

	return so, nil
}

func (r *categoryRepository) FindMonthMethodCategorySuccessById(req *requests.MonthMethodCategoryByIdRequest) ([]*record.MonthMethodCategoryRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthMethodCategoriesSuccessById(r.ctx, db.GetMonthMethodCategoriesSuccessByIdParams{
		Column1:    yearStart,
		CategoryID: int32(req.ID),
	})

	if err != nil {
		return nil, category_errors.ErrFindMonthMethodCategorySuccessById
	}

	so := r.mapping.ToCategoriesRecordMonthMethodSuccessById(res)

	return so, nil
}

func (r *categoryRepository) FindYearMethodCategorySuccessById(req *requests.YearMethodCategoryByIdRequest) ([]*record.YearMethodCategoryRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearMethodCategoriesSuccessById(r.ctx, db.GetYearMethodCategoriesSuccessByIdParams{
		Column1:    yearStart,
		CategoryID: int32(req.ID),
	})

	if err != nil {
		return nil, category_errors.ErrFindYearMethodCategorySuccessById
	}

	so := r.mapping.ToCategoriesRecordYearMethodSuccessById(res)

	return so, nil
}

func (r *categoryRepository) FindMonthMethodCategoryFailedById(req *requests.MonthMethodCategoryByIdRequest) ([]*record.MonthMethodCategoryRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthMethodCategoriesFailedById(r.ctx, db.GetMonthMethodCategoriesFailedByIdParams{
		Column1:    yearStart,
		CategoryID: int32(req.ID),
	})

	if err != nil {
		return nil, category_errors.ErrFindYearMethodCategoryFailedById
	}

	so := r.mapping.ToCategoriesRecordMonthMethodFailedById(res)

	return so, nil
}

func (r *categoryRepository) FindYearMethodCategoryFailedById(req *requests.YearMethodCategoryByIdRequest) ([]*record.YearMethodCategoryRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearMethodCategoriesFailedById(r.ctx, db.GetYearMethodCategoriesFailedByIdParams{
		Column1:    yearStart,
		CategoryID: int32(req.ID),
	})

	if err != nil {
		return nil, category_errors.ErrFindYearMethodCategoryFailedById
	}

	so := r.mapping.ToCategoriesRecordYearMethodFailedById(res)

	return so, nil
}

func (r *categoryRepository) FindMonthAmountCategorySuccessByMerchant(req *requests.MonthAmountCategoryByMerchantRequest) ([]*record.MonthAmountCategorySuccessRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountCategorySuccessByMerchant(r.ctx, db.GetMonthAmountCategorySuccessByMerchantParams{
		Column1:    currentDate,
		Column2:    lastDayCurrentMonth,
		Column3:    prevDate,
		Column4:    lastDayPrevMonth,
		MerchantID: int32(req.MerchantID),
	})

	if err != nil {
		return nil, category_errors.ErrFindMonthAmountCategorySuccessByMerchant
	}

	so := r.mapping.ToCategoriesRecordMonthAmountSuccessByMerchant(res)

	return so, nil
}

func (r *categoryRepository) FindYearAmountCategorySuccessByMerchant(req *requests.YearAmountCategoryByMerchantRequest) ([]*record.YearAmountCategorySuccessRecord, error) {
	res, err := r.db.GetYearAmountCategorySuccessByMerchant(r.ctx, db.GetYearAmountCategorySuccessByMerchantParams{
		Column1:    int32(req.Year),
		MerchantID: int32(req.MerchantID),
	})

	if err != nil {
		return nil, category_errors.ErrFindYearAmountCategorySuccessByMerchant
	}

	so := r.mapping.ToCategoriesRecordYearAmountSuccessByMerchant(res)

	return so, nil
}

func (r *categoryRepository) FindMonthAmountCategoryFailedByMerchant(req *requests.MonthAmountCategoryByMerchantRequest) ([]*record.MonthAmountCategoryFailedRecord, error) {
	currentDate := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)
	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthAmountCategoryFailedByMerchant(r.ctx, db.GetMonthAmountCategoryFailedByMerchantParams{
		Column1:    currentDate,
		Column2:    lastDayCurrentMonth,
		Column3:    prevDate,
		Column4:    lastDayPrevMonth,
		MerchantID: int32(req.MerchantID),
	})

	if err != nil {
		return nil, category_errors.ErrFindMonthAmountCategoryFailedByMerchant
	}

	so := r.mapping.ToCategoriesRecordMonthAmountFailedByMerchant(res)

	return so, nil
}

func (r *categoryRepository) FindYearAmountCategoryFailedByMerchant(req *requests.YearAmountCategoryByMerchantRequest) ([]*record.YearAmountCategoryFailedRecord, error) {
	res, err := r.db.GetYearAmountCategoryFailedByMerchant(r.ctx, db.GetYearAmountCategoryFailedByMerchantParams{
		Column1:    int32(req.Year),
		MerchantID: int32(req.MerchantID),
	})

	if err != nil {
		return nil, category_errors.ErrFindYearAmountCategoryFailedByMerchant
	}

	so := r.mapping.ToCategoriesRecordYearAmountFaileByMerchant(res)

	return so, nil
}

func (r *categoryRepository) FindMonthMethodCategorySuccessByMerchant(req *requests.MonthMethodCategoryByMerchantRequest) ([]*record.MonthMethodCategoryRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthMethodCategoriesSuccessByMerchant(r.ctx, db.GetMonthMethodCategoriesSuccessByMerchantParams{
		Column1:    yearStart,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, category_errors.ErrFindMonthMethodCategorySuccessByMerchant
	}

	so := r.mapping.ToCategoriesRecordMonthMethodSuccessByMerchant(res)

	return so, nil
}

func (r *categoryRepository) FindYearMethodCategorySuccessByMerchant(req *requests.YearMethodCategoryByMerchantRequest) ([]*record.YearMethodCategoryRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearMethodCategoriesSuccessByMerchant(r.ctx, db.GetYearMethodCategoriesSuccessByMerchantParams{
		Column1:    yearStart,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, category_errors.ErrFindYearMethodCategorySuccessByMerchant
	}

	so := r.mapping.ToCategoriesRecordYearMethodSuccessByMerchant(res)

	return so, nil
}

func (r *categoryRepository) FindMonthMethodCategoryFailedByMerchant(req *requests.MonthMethodCategoryByMerchantRequest) ([]*record.MonthMethodCategoryRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthMethodCategoriesFailedByMerchant(r.ctx, db.GetMonthMethodCategoriesFailedByMerchantParams{
		Column1:    yearStart,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, category_errors.ErrFindYearMethodCategoryFailedByMerchant
	}

	so := r.mapping.ToCategoriesRecordMonthMethodFailedByMerchant(res)

	return so, nil
}

func (r *categoryRepository) FindYearMethodCategoryFailedByMerchant(req *requests.YearMethodCategoryByMerchantRequest) ([]*record.YearMethodCategoryRecord, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearMethodCategoriesFailedByMerchant(r.ctx, db.GetYearMethodCategoriesFailedByMerchantParams{
		Column1:    yearStart,
		MerchantID: sql.NullInt32{Int32: int32(req.MerchantID), Valid: true},
	})

	if err != nil {
		return nil, category_errors.ErrFindYearMethodCategoryFailedByMerchant
	}

	so := r.mapping.ToCategoriesRecordYearMethodFailedByMerchant(res)

	return so, nil
}

func (r *categoryRepository) FindById(id int) (*record.CategoryRecord, error) {
	res, err := r.db.GetCategoryByID(r.ctx, int32(id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, category_errors.ErrCategoryNotFound
		}

		return nil, category_errors.ErrCategoryNotFound
	}

	return r.mapping.ToCategoryRecord(res), nil
}

func (r *categoryRepository) FindByActiveCategories(request *requests.FindAllCategory) ([]*record.CategoryRecord, *int, error) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	offset := (page - 1) * pageSize

	req := db.GetCategoriesActiveParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetCategoriesActive(r.ctx, req)

	if err != nil {
		return nil, nil, category_errors.ErrFindActiveCategories
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToCategoriesRecordActive(res), &totalCount, nil
}

func (r *categoryRepository) FindByTrashedCategories(request *requests.FindAllCategory) ([]*record.CategoryRecord, *int, error) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	offset := (page - 1) * pageSize

	req := db.GetCategoriesTrashedParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetCategoriesTrashed(r.ctx, req)

	if err != nil {
		return nil, nil, category_errors.ErrFindTrashedCategories
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToCategoriesRecordTrashed(res), &totalCount, nil
}

func (r *categoryRepository) CreateCategory(req *requests.CreateCategoryRequest) (*record.CategoryRecord, error) {
	res, err := r.db.CreateCategory(r.ctx, req.Name)

	if err != nil {
		return nil, category_errors.ErrCreateCategory
	}

	return r.mapping.ToCategoryRecord(res), nil
}

func (r *categoryRepository) UpdateCategory(req *requests.UpdateCategoryRequest) (*record.CategoryRecord, error) {
	res, err := r.db.UpdateCategory(r.ctx, db.UpdateCategoryParams{
		CategoryID: int32(req.ID),
		Name:       req.Name,
	})

	if err != nil {
		return nil, category_errors.ErrUpdateCategory
	}

	return r.mapping.ToCategoryRecord(res), nil
}

func (r *categoryRepository) TrashedCategory(id int) (*record.CategoryRecord, error) {
	res, err := r.db.TrashCategory(r.ctx, int32(id))

	if err != nil {
		return nil, category_errors.ErrTrashedCategory
	}

	return r.mapping.ToCategoryRecord(res), nil
}

func (r *categoryRepository) RestoreCategory(id int) (*record.CategoryRecord, error) {
	res, err := r.db.RestoreCategory(r.ctx, int32(id))

	if err != nil {
		return nil, category_errors.ErrRestoreCategory
	}

	return r.mapping.ToCategoryRecord(res), nil
}

func (r *categoryRepository) DeleteCategoryPermanent(category_id int) (bool, error) {
	err := r.db.DeleteCategoryPermanently(r.ctx, int32(category_id))

	if err != nil {
		return false, category_errors.ErrDeleteCategoryPermanent
	}

	return true, nil
}

func (r *categoryRepository) RestoreAllCategories() (bool, error) {
	err := r.db.RestoreAllCategories(r.ctx)

	if err != nil {
		return false, category_errors.ErrRestoreAllCategories
	}

	return true, nil
}

func (r *categoryRepository) DeleteAllCategoriesPermanent() (bool, error) {
	err := r.db.DeleteAllPermanentCategories(r.ctx)

	if err != nil {
		return false, category_errors.ErrDeleteAllCategories
	}

	return true, nil
}
