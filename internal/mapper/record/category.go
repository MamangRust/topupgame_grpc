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

func (s *categoryRecordMapper) ToCategoriesRecord(Categorys []*db.Category) []*record.CategoryRecord {
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

func (s *categoryRecordMapper) ToCategoriesRecordAll(Categorys []*db.GetCategoriesRow) []*record.CategoryRecord {
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

func (s *categoryRecordMapper) ToCategoriesRecordActive(Categorys []*db.GetCategoriesActiveRow) []*record.CategoryRecord {
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

func (s *categoryRecordMapper) ToCategoriesRecordTrashed(Categorys []*db.GetCategoriesTrashedRow) []*record.CategoryRecord {
	var result []*record.CategoryRecord

	for _, Category := range Categorys {
		result = append(result, s.ToCategoryRecordTrashed(Category))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordMonthAmountSuccess(b *db.GetMonthAmountCategorySuccessRow) *record.MonthAmountCategorySuccessRecord {
	return &record.MonthAmountCategorySuccessRecord{
		ID:           int(b.CategoryID),
		CategoryName: b.CategoryName,
		Year:         b.MccYear,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *categoryRecordMapper) ToCategoriesRecordMonthAmountSuccess(b []*db.GetMonthAmountCategorySuccessRow) []*record.MonthAmountCategorySuccessRecord {
	var result []*record.MonthAmountCategorySuccessRecord

	for _, Bank := range b {
		result = append(result, s.ToCategoryRecordMonthAmountSuccess(Bank))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordYearAmountSuccess(b *db.GetYearAmountCategorySuccessRow) *record.YearAmountCategorySuccessRecord {
	return &record.YearAmountCategorySuccessRecord{
		ID:           int(b.CategoryID),
		CategoryName: b.CategoryName,
		Year:         b.YccYear,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *categoryRecordMapper) ToCategoriesRecordYearAmountSuccess(b []*db.GetYearAmountCategorySuccessRow) []*record.YearAmountCategorySuccessRecord {
	var result []*record.YearAmountCategorySuccessRecord

	for _, Bank := range b {
		result = append(result, s.ToCategoryRecordYearAmountSuccess(Bank))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordMonthAmountFailed(b *db.GetMonthAmountCategoryFailedRow) *record.MonthAmountCategoryFailedRecord {
	return &record.MonthAmountCategoryFailedRecord{
		ID:           int(b.CategoryID),
		CategoryName: b.CategoryName,
		Year:         b.MccYear,
		Month:        b.Month,
		TotalFailed:  int(b.TotalFailed),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *categoryRecordMapper) ToCategoriesRecordMonthAmountFailed(b []*db.GetMonthAmountCategoryFailedRow) []*record.MonthAmountCategoryFailedRecord {
	var result []*record.MonthAmountCategoryFailedRecord

	for _, Bank := range b {
		result = append(result, s.ToCategoryRecordMonthAmountFailed(Bank))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordYearAmountFailed(b *db.GetYearAmountCategoryFailedRow) *record.YearAmountCategoryFailedRecord {
	return &record.YearAmountCategoryFailedRecord{
		ID:           int(b.CategoryID),
		CategoryName: b.CategoryName,
		Year:         b.YccYear,
		TotalFailed:  int(b.TotalFailed),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *categoryRecordMapper) ToCategoriesRecordYearAmountFailed(b []*db.GetYearAmountCategoryFailedRow) []*record.YearAmountCategoryFailedRecord {
	var result []*record.YearAmountCategoryFailedRecord

	for _, Bank := range b {
		result = append(result, s.ToCategoryRecordYearAmountFailed(Bank))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordMonthMethodSuccess(b *db.GetMonthMethodCategoriesSuccessRow) *record.MonthMethodCategoryRecord {
	return &record.MonthMethodCategoryRecord{
		ID:                int(b.CategoryID),
		Month:             b.Month,
		CategoryName:      b.CategoryName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s categoryRecordMapper) ToCategoriesRecordMonthMethodSuccess(b []*db.GetMonthMethodCategoriesSuccessRow) []*record.MonthMethodCategoryRecord {
	var result []*record.MonthMethodCategoryRecord

	for _, category := range b {
		result = append(result, s.ToCategoryRecordMonthMethodSuccess(category))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordMonthMethodFailed(b *db.GetMonthMethodCategoriesFailedRow) *record.MonthMethodCategoryRecord {
	return &record.MonthMethodCategoryRecord{
		ID:                int(b.CategoryID),
		Month:             b.Month,
		CategoryName:      b.CategoryName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s categoryRecordMapper) ToCategoriesRecordMonthMethodFailed(b []*db.GetMonthMethodCategoriesFailedRow) []*record.MonthMethodCategoryRecord {
	var result []*record.MonthMethodCategoryRecord

	for _, category := range b {
		result = append(result, s.ToCategoryRecordMonthMethodFailed(category))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordYearMethodSuccess(b *db.GetYearMethodCategoriesSuccessRow) *record.YearMethodCategoryRecord {
	return &record.YearMethodCategoryRecord{
		ID:                int(b.CategoryID),
		Year:              b.Year,
		CategoryName:      b.CategoryName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s categoryRecordMapper) ToCategoriesRecordYearMethodSuccess(b []*db.GetYearMethodCategoriesSuccessRow) []*record.YearMethodCategoryRecord {
	var result []*record.YearMethodCategoryRecord

	for _, category := range b {
		result = append(result, s.ToCategoryRecordYearMethodSuccess(category))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordYearMethodFailed(b *db.GetYearMethodCategoriesFailedRow) *record.YearMethodCategoryRecord {
	return &record.YearMethodCategoryRecord{
		ID:                int(b.CategoryID),
		Year:              b.Year,
		CategoryName:      b.CategoryName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s categoryRecordMapper) ToCategoriesRecordYearMethodFailed(b []*db.GetYearMethodCategoriesFailedRow) []*record.YearMethodCategoryRecord {
	var result []*record.YearMethodCategoryRecord

	for _, category := range b {
		result = append(result, s.ToCategoryRecordYearMethodFailed(category))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordMonthAmountSuccessById(b *db.GetMonthAmountCategorySuccessByIdRow) *record.MonthAmountCategorySuccessRecord {
	return &record.MonthAmountCategorySuccessRecord{
		ID:           int(b.CategoryID),
		CategoryName: b.CategoryName,
		Year:         b.MccYear,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *categoryRecordMapper) ToCategoriesRecordMonthAmountSuccessById(b []*db.GetMonthAmountCategorySuccessByIdRow) []*record.MonthAmountCategorySuccessRecord {
	var result []*record.MonthAmountCategorySuccessRecord

	for _, Bank := range b {
		result = append(result, s.ToCategoryRecordMonthAmountSuccessById(Bank))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordYearAmountSuccessById(b *db.GetYearAmountCategorySuccessByIdRow) *record.YearAmountCategorySuccessRecord {
	return &record.YearAmountCategorySuccessRecord{
		ID:           int(b.CategoryID),
		CategoryName: b.CategoryName,
		Year:         b.YccYear,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *categoryRecordMapper) ToCategoriesRecordYearAmountSuccessById(b []*db.GetYearAmountCategorySuccessByIdRow) []*record.YearAmountCategorySuccessRecord {
	var result []*record.YearAmountCategorySuccessRecord

	for _, Bank := range b {
		result = append(result, s.ToCategoryRecordYearAmountSuccessById(Bank))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordMonthAmountFailedById(b *db.GetMonthAmountCategoryFailedByIdRow) *record.MonthAmountCategoryFailedRecord {
	return &record.MonthAmountCategoryFailedRecord{
		ID:           int(b.CategoryID),
		CategoryName: b.CategoryName,
		Year:         b.MccYear,
		Month:        b.Month,
		TotalFailed:  int(b.TotalFailed),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *categoryRecordMapper) ToCategoriesRecordMonthAmountFailedById(b []*db.GetMonthAmountCategoryFailedByIdRow) []*record.MonthAmountCategoryFailedRecord {
	var result []*record.MonthAmountCategoryFailedRecord

	for _, Bank := range b {
		result = append(result, s.ToCategoryRecordMonthAmountFailedById(Bank))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordYearAmountFailedById(b *db.GetYearAmountCategoryFailedByIdRow) *record.YearAmountCategoryFailedRecord {
	return &record.YearAmountCategoryFailedRecord{
		ID:           int(b.CategoryID),
		CategoryName: b.CategoryName,
		Year:         b.YccYear,
		TotalFailed:  int(b.TotalFailed),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *categoryRecordMapper) ToCategoriesRecordYearAmountFaileById(b []*db.GetYearAmountCategoryFailedByIdRow) []*record.YearAmountCategoryFailedRecord {
	var result []*record.YearAmountCategoryFailedRecord

	for _, Bank := range b {
		result = append(result, s.ToCategoryRecordYearAmountFailedById(Bank))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordMonthMethodSuccessById(b *db.GetMonthMethodCategoriesSuccessByIdRow) *record.MonthMethodCategoryRecord {
	return &record.MonthMethodCategoryRecord{
		ID:                int(b.CategoryID),
		Month:             b.Month,
		CategoryName:      b.CategoryName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s categoryRecordMapper) ToCategoriesRecordMonthMethodSuccessById(b []*db.GetMonthMethodCategoriesSuccessByIdRow) []*record.MonthMethodCategoryRecord {
	var result []*record.MonthMethodCategoryRecord

	for _, category := range b {
		result = append(result, s.ToCategoryRecordMonthMethodSuccessById(category))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordMonthMethodFailedById(b *db.GetMonthMethodCategoriesFailedByIdRow) *record.MonthMethodCategoryRecord {
	return &record.MonthMethodCategoryRecord{
		ID:                int(b.CategoryID),
		Month:             b.Month,
		CategoryName:      b.CategoryName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s categoryRecordMapper) ToCategoriesRecordMonthMethodFailedById(b []*db.GetMonthMethodCategoriesFailedByIdRow) []*record.MonthMethodCategoryRecord {
	var result []*record.MonthMethodCategoryRecord

	for _, category := range b {
		result = append(result, s.ToCategoryRecordMonthMethodFailedById(category))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordYearMethodSuccessById(b *db.GetYearMethodCategoriesSuccessByIdRow) *record.YearMethodCategoryRecord {
	return &record.YearMethodCategoryRecord{
		ID:                int(b.CategoryID),
		Year:              b.Year,
		CategoryName:      b.CategoryName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s categoryRecordMapper) ToCategoriesRecordYearMethodSuccessById(b []*db.GetYearMethodCategoriesSuccessByIdRow) []*record.YearMethodCategoryRecord {
	var result []*record.YearMethodCategoryRecord

	for _, category := range b {
		result = append(result, s.ToCategoryRecordYearMethodSuccessById(category))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordYearMethodFailedById(b *db.GetYearMethodCategoriesFailedByIdRow) *record.YearMethodCategoryRecord {
	return &record.YearMethodCategoryRecord{
		ID:                int(b.CategoryID),
		Year:              b.Year,
		CategoryName:      b.CategoryName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s categoryRecordMapper) ToCategoriesRecordYearMethodFailedById(b []*db.GetYearMethodCategoriesFailedByIdRow) []*record.YearMethodCategoryRecord {
	var result []*record.YearMethodCategoryRecord

	for _, category := range b {
		result = append(result, s.ToCategoryRecordYearMethodFailedById(category))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordMonthAmountSuccessByMerchant(b *db.GetMonthAmountCategorySuccessByMerchantRow) *record.MonthAmountCategorySuccessRecord {
	return &record.MonthAmountCategorySuccessRecord{
		ID:           int(b.CategoryID),
		CategoryName: b.CategoryName,
		Year:         b.MccYear,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *categoryRecordMapper) ToCategoriesRecordMonthAmountSuccessByMerchant(b []*db.GetMonthAmountCategorySuccessByMerchantRow) []*record.MonthAmountCategorySuccessRecord {
	var result []*record.MonthAmountCategorySuccessRecord

	for _, Bank := range b {
		result = append(result, s.ToCategoryRecordMonthAmountSuccessByMerchant(Bank))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordYearAmountSuccessByMerchant(b *db.GetYearAmountCategorySuccessByMerchantRow) *record.YearAmountCategorySuccessRecord {
	return &record.YearAmountCategorySuccessRecord{
		ID:           int(b.CategoryID),
		CategoryName: b.CategoryName,
		Year:         b.YccYear,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *categoryRecordMapper) ToCategoriesRecordYearAmountSuccessByMerchant(b []*db.GetYearAmountCategorySuccessByMerchantRow) []*record.YearAmountCategorySuccessRecord {
	var result []*record.YearAmountCategorySuccessRecord

	for _, Bank := range b {
		result = append(result, s.ToCategoryRecordYearAmountSuccessByMerchant(Bank))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordMonthAmountFailedByMerchant(b *db.GetMonthAmountCategoryFailedByMerchantRow) *record.MonthAmountCategoryFailedRecord {
	return &record.MonthAmountCategoryFailedRecord{
		ID:           int(b.CategoryID),
		CategoryName: b.CategoryName,
		Year:         b.MccYear,
		Month:        b.Month,
		TotalFailed:  int(b.TotalFailed),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *categoryRecordMapper) ToCategoriesRecordMonthAmountFailedByMerchant(b []*db.GetMonthAmountCategoryFailedByMerchantRow) []*record.MonthAmountCategoryFailedRecord {
	var result []*record.MonthAmountCategoryFailedRecord

	for _, Bank := range b {
		result = append(result, s.ToCategoryRecordMonthAmountFailedByMerchant(Bank))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordYearAmountFailedByMerchant(b *db.GetYearAmountCategoryFailedByMerchantRow) *record.YearAmountCategoryFailedRecord {
	return &record.YearAmountCategoryFailedRecord{
		ID:           int(b.CategoryID),
		CategoryName: b.CategoryName,
		Year:         b.YccYear,
		TotalFailed:  int(b.TotalFailed),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *categoryRecordMapper) ToCategoriesRecordYearAmountFaileByMerchant(b []*db.GetYearAmountCategoryFailedByMerchantRow) []*record.YearAmountCategoryFailedRecord {
	var result []*record.YearAmountCategoryFailedRecord

	for _, Bank := range b {
		result = append(result, s.ToCategoryRecordYearAmountFailedByMerchant(Bank))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordMonthMethodSuccessByMerchant(b *db.GetMonthMethodCategoriesSuccessByMerchantRow) *record.MonthMethodCategoryRecord {
	return &record.MonthMethodCategoryRecord{
		ID:                int(b.CategoryID),
		Month:             b.Month,
		CategoryName:      b.CategoryName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s categoryRecordMapper) ToCategoriesRecordMonthMethodSuccessByMerchant(b []*db.GetMonthMethodCategoriesSuccessByMerchantRow) []*record.MonthMethodCategoryRecord {
	var result []*record.MonthMethodCategoryRecord

	for _, category := range b {
		result = append(result, s.ToCategoryRecordMonthMethodSuccessByMerchant(category))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordMonthMethodFailedByMerchant(b *db.GetMonthMethodCategoriesFailedByMerchantRow) *record.MonthMethodCategoryRecord {
	return &record.MonthMethodCategoryRecord{
		ID:                int(b.CategoryID),
		Month:             b.Month,
		CategoryName:      b.CategoryName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s categoryRecordMapper) ToCategoriesRecordMonthMethodFailedByMerchant(b []*db.GetMonthMethodCategoriesFailedByMerchantRow) []*record.MonthMethodCategoryRecord {
	var result []*record.MonthMethodCategoryRecord

	for _, category := range b {
		result = append(result, s.ToCategoryRecordMonthMethodFailedByMerchant(category))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordYearMethodSuccessByMerchant(b *db.GetYearMethodCategoriesSuccessByMerchantRow) *record.YearMethodCategoryRecord {
	return &record.YearMethodCategoryRecord{
		ID:                int(b.CategoryID),
		Year:              b.Year,
		CategoryName:      b.CategoryName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s categoryRecordMapper) ToCategoriesRecordYearMethodSuccessByMerchant(b []*db.GetYearMethodCategoriesSuccessByMerchantRow) []*record.YearMethodCategoryRecord {
	var result []*record.YearMethodCategoryRecord

	for _, category := range b {
		result = append(result, s.ToCategoryRecordYearMethodSuccessByMerchant(category))
	}

	return result
}

func (s *categoryRecordMapper) ToCategoryRecordYearMethodFailedByMerchant(b *db.GetYearMethodCategoriesFailedByMerchantRow) *record.YearMethodCategoryRecord {
	return &record.YearMethodCategoryRecord{
		ID:                int(b.CategoryID),
		Year:              b.Year,
		CategoryName:      b.CategoryName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s categoryRecordMapper) ToCategoriesRecordYearMethodFailedByMerchant(b []*db.GetYearMethodCategoriesFailedByMerchantRow) []*record.YearMethodCategoryRecord {
	var result []*record.YearMethodCategoryRecord

	for _, category := range b {
		result = append(result, s.ToCategoryRecordYearMethodFailedByMerchant(category))
	}

	return result
}
