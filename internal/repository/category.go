package repository

import (
	"context"
	"fmt"
	"topup_game/internal/domain/record"
	"topup_game/internal/domain/requests"
	recordmapper "topup_game/internal/mapper/record"
	db "topup_game/pkg/database/schema"
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

func (r *categoryRepository) FindAllCategories(page int, pageSize int, search string) ([]*record.CategoryRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetCategoriesParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetCategories(r.ctx, req)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to find categories: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToCategorysRecordAll(res), totalCount, nil
}

func (r *categoryRepository) FindById(id int) (*record.CategoryRecord, error) {
	res, err := r.db.GetCategoryByID(r.ctx, int32(id))

	if err != nil {
		return nil, fmt.Errorf("failed to find category by id: %w", err)
	}

	return r.mapping.ToCategoryRecord(res), nil
}

func (r *categoryRepository) FindByActiveCategories(page int, pageSize int, search string) ([]*record.CategoryRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetCategoriesActiveParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetCategoriesActive(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find active categories: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToCategorysRecordActive(res), totalCount, nil
}

func (r *categoryRepository) FindByTrashedCategory(page int, pageSize int, search string) ([]*record.CategoryRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetCategoriesTrashedParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetCategoriesTrashed(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find trashed categories: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToCategorysRecordTrashed(res), totalCount, nil
}

func (r *categoryRepository) CreateCategory(req *requests.CreateCategoryRequest) (*record.CategoryRecord, error) {
	res, err := r.db.CreateCategory(r.ctx, req.Name)

	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	return r.mapping.ToCategoryRecord(res), nil
}

func (r *categoryRepository) UpdateCategory(req *requests.UpdateCategoryRequest) (*record.CategoryRecord, error) {
	res, err := r.db.UpdateCategory(r.ctx, db.UpdateCategoryParams{
		CategoryID: int32(req.ID),
		Name:       req.Name,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to update category: %w", err)
	}

	return r.mapping.ToCategoryRecord(res), nil
}

func (r *categoryRepository) TrashedCategory(id int) (*record.CategoryRecord, error) {
	res, err := r.db.TrashCategory(r.ctx, int32(id))

	if err != nil {
		return nil, fmt.Errorf("failed to trash category: %w", err)
	}

	return r.mapping.ToCategoryRecord(res), nil
}

func (r *categoryRepository) RestoreCategory(id int) (*record.CategoryRecord, error) {
	res, err := r.db.RestoreCategory(r.ctx, int32(id))

	if err != nil {
		return nil, fmt.Errorf("failed to find category after restore: %w", err)
	}

	return r.mapping.ToCategoryRecord(res), nil
}

func (r *categoryRepository) DeleteCategoryPermanent(category_id int) (bool, error) {
	err := r.db.DeleteCategoryPermanently(r.ctx, int32(category_id))

	if err != nil {
		return false, fmt.Errorf("failed to delete category: %w", err)
	}

	return true, nil
}

func (r *categoryRepository) RestoreAllCategories() (bool, error) {
	err := r.db.RestoreAllCategories(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to restore all categories: %w", err)
	}

	return true, nil
}

func (r *categoryRepository) DeleteAllCategoriesPermanent() (bool, error) {
	err := r.db.DeleteAllPermanentCategories(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to delete all categories permanently: %w", err)
	}

	return true, nil
}
