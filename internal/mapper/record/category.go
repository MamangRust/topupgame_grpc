package recordmapper

import (
	"topup_game/internal/domain/record"
	db "topup_game/pkg/database/schema"
)

type categoryRecordMapper struct {
}

func NewCategoryRecordMapper() *categoryRecordMapper {
	return &categoryRecordMapper{}
}

func (s *categoryRecordMapper) ToCategoryRecord(Category *db.Category) *record.CategoryRecord {
	deletedAt := Category.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.CategoryRecord{
		ID:        int(Category.CategoryID),
		Name:      Category.Name,
		CreatedAt: Category.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt: Category.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt: &deletedAt,
	}
}

func (s *categoryRecordMapper) ToCategorysRecord(Categorys []*db.Category) []*record.CategoryRecord {
	var result []*record.CategoryRecord

	for _, Category := range Categorys {
		result = append(result, s.ToCategoryRecord(Category))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordAll(Category *db.GetCategoriesRow) *record.CategoryRecord {
	deletedAt := Category.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.CategoryRecord{
		ID:        int(Category.CategoryID),
		Name:      Category.Name,
		CreatedAt: Category.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt: Category.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt: &deletedAt,
	}
}

func (s *categoryRecordMapper) ToCategorysRecordAll(Categorys []*db.GetCategoriesRow) []*record.CategoryRecord {
	var result []*record.CategoryRecord

	for _, Category := range Categorys {
		result = append(result, s.ToCategoryRecordAll(Category))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordActive(Category *db.GetCategoriesActiveRow) *record.CategoryRecord {
	deletedAt := Category.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.CategoryRecord{
		ID:        int(Category.CategoryID),
		Name:      Category.Name,
		CreatedAt: Category.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt: Category.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt: &deletedAt,
	}
}

func (s *categoryRecordMapper) ToCategorysRecordActive(Categorys []*db.GetCategoriesActiveRow) []*record.CategoryRecord {
	var result []*record.CategoryRecord

	for _, Category := range Categorys {
		result = append(result, s.ToCategoryRecordActive(Category))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordTrashed(Category *db.GetCategoriesTrashedRow) *record.CategoryRecord {
	deletedAt := Category.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.CategoryRecord{
		ID:        int(Category.CategoryID),
		Name:      Category.Name,
		CreatedAt: Category.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt: Category.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt: &deletedAt,
	}
}

func (s *categoryRecordMapper) ToCategorysRecordTrashed(Categorys []*db.GetCategoriesTrashedRow) []*record.CategoryRecord {
	var result []*record.CategoryRecord

	for _, Category := range Categorys {
		result = append(result, s.ToCategoryRecordTrashed(Category))
	}

	return result
}
