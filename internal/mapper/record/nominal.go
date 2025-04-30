package recordmapper

import (
	"topup_game/internal/domain/record"
	db "topup_game/pkg/database/schema"
)

type nominalRecordMapper struct{}

func NewNominalRecordMapper() *nominalRecordMapper {
	return &nominalRecordMapper{}
}

func (s *nominalRecordMapper) ToNominalRecord(nominal *db.Nominal) *record.NominalRecord {
	deletedAt := nominal.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.NominalRecord{
		ID:        int(nominal.NominalID),
		VoucherID: int(nominal.VoucherID),
		Name:      nominal.Name,
		Quantity:  int(nominal.Quantity),
		Price:     nominal.Price,
		CreatedAt: nominal.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt: nominal.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt: &deletedAt,
	}
}

func (s *nominalRecordMapper) ToNominalRecords(nominals []*db.Nominal) []*record.NominalRecord {
	var result []*record.NominalRecord

	for _, nominal := range nominals {
		result = append(result, s.ToNominalRecord(nominal))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordAll(nominal *db.GetNominalsRow) *record.NominalRecord {
	deletedAt := nominal.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.NominalRecord{
		ID:        int(nominal.NominalID),
		VoucherID: int(nominal.VoucherID),
		Name:      nominal.Name,
		Quantity:  int(nominal.Quantity),
		Price:     nominal.Price,
		CreatedAt: nominal.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt: nominal.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt: &deletedAt,
	}
}

func (s *nominalRecordMapper) ToNominalRecordsAll(nominals []*db.GetNominalsRow) []*record.NominalRecord {
	var result []*record.NominalRecord

	for _, nominal := range nominals {
		result = append(result, s.ToNominalRecordAll(nominal))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordActive(nominal *db.GetNominalsActiveRow) *record.NominalRecord {
	deletedAt := nominal.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.NominalRecord{
		ID:        int(nominal.NominalID),
		VoucherID: int(nominal.VoucherID),
		Name:      nominal.Name,
		Quantity:  int(nominal.Quantity),
		Price:     nominal.Price,
		CreatedAt: nominal.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt: nominal.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt: &deletedAt,
	}
}

func (s *nominalRecordMapper) ToNominalRecordsActive(nominals []*db.GetNominalsActiveRow) []*record.NominalRecord {
	var result []*record.NominalRecord

	for _, nominal := range nominals {
		result = append(result, s.ToNominalRecordActive(nominal))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordTrashed(nominal *db.GetNominalsTrashedRow) *record.NominalRecord {
	deletedAt := nominal.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.NominalRecord{
		ID:        int(nominal.NominalID),
		VoucherID: int(nominal.VoucherID),
		Name:      nominal.Name,
		Quantity:  int(nominal.Quantity),
		Price:     nominal.Price,
		CreatedAt: nominal.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt: nominal.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt: &deletedAt,
	}
}

func (s *nominalRecordMapper) ToNominalRecordsTrashed(nominals []*db.GetNominalsTrashedRow) []*record.NominalRecord {
	var result []*record.NominalRecord

	for _, nominal := range nominals {
		result = append(result, s.ToNominalRecordTrashed(nominal))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordMonthAmountSuccess(b *db.GetMonthAmountNominalsSuccessRow) *record.MonthAmountNominalSuccessRecord {
	return &record.MonthAmountNominalSuccessRecord{
		ID:           int(b.NominalID),
		NominalName:  b.NominalName,
		Year:         b.MncYear,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *nominalRecordMapper) ToNominalsRecordMonthAmountSuccess(b []*db.GetMonthAmountNominalsSuccessRow) []*record.MonthAmountNominalSuccessRecord {
	var result []*record.MonthAmountNominalSuccessRecord

	for _, nominal := range b {
		result = append(result, s.ToNominalRecordMonthAmountSuccess(nominal))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordYearAmountSuccess(b *db.GetYearAmountNominalsSuccessRow) *record.YearAmountNominalSuccessRecord {
	return &record.YearAmountNominalSuccessRecord{
		ID:           int(b.NominalID),
		NominalName:  b.NominalName,
		Year:         b.YncYear,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *nominalRecordMapper) ToNominalsRecordYearAmountSuccess(b []*db.GetYearAmountNominalsSuccessRow) []*record.YearAmountNominalSuccessRecord {
	var result []*record.YearAmountNominalSuccessRecord

	for _, nominal := range b {
		result = append(result, s.ToNominalRecordYearAmountSuccess(nominal))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordMonthAmountFailed(b *db.GetMonthAmountNominalsFailedRow) *record.MonthAmountNominalFailedRecord {
	return &record.MonthAmountNominalFailedRecord{
		ID:          int(b.NominalID),
		NominalName: b.NominalName,
		Year:        b.MncYear,
		Month:       b.Month,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *nominalRecordMapper) ToNominalsRecordMonthAmountFailed(b []*db.GetMonthAmountNominalsFailedRow) []*record.MonthAmountNominalFailedRecord {
	var result []*record.MonthAmountNominalFailedRecord

	for _, Bank := range b {
		result = append(result, s.ToNominalRecordMonthAmountFailed(Bank))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordYearAmountFailed(b *db.GetYearAmountNominalsFailedRow) *record.YearAmountNominalFailedRecord {
	return &record.YearAmountNominalFailedRecord{
		ID:          int(b.NominalID),
		NominalName: b.NominalName,
		Year:        b.YncYear,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *nominalRecordMapper) ToNominalsRecordYearAmountFailed(b []*db.GetYearAmountNominalsFailedRow) []*record.YearAmountNominalFailedRecord {
	var result []*record.YearAmountNominalFailedRecord

	for _, Bank := range b {
		result = append(result, s.ToNominalRecordYearAmountFailed(Bank))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordMonthMethodSuccess(b *db.GetMonthMethodNominalsSuccessRow) *record.MonthMethodNominalRecord {
	return &record.MonthMethodNominalRecord{
		ID:                int(b.NominalID),
		Month:             b.Month,
		NominalName:       b.NominalName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s nominalRecordMapper) ToNominalsRecordMonthMethodSuccess(b []*db.GetMonthMethodNominalsSuccessRow) []*record.MonthMethodNominalRecord {
	var result []*record.MonthMethodNominalRecord

	for _, nominal := range b {
		result = append(result, s.ToNominalRecordMonthMethodSuccess(nominal))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordMonthMethodFailed(b *db.GetMonthMethodNominalsFailedRow) *record.MonthMethodNominalRecord {
	return &record.MonthMethodNominalRecord{
		ID:                int(b.NominalID),
		Month:             b.Month,
		NominalName:       b.NominalName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s nominalRecordMapper) ToNominalsRecordMonthMethodFailed(b []*db.GetMonthMethodNominalsFailedRow) []*record.MonthMethodNominalRecord {
	var result []*record.MonthMethodNominalRecord

	for _, nominal := range b {
		result = append(result, s.ToNominalRecordMonthMethodFailed(nominal))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordYearMethodSuccess(b *db.GetYearMethodNominalsSuccessRow) *record.YearMethodNominalRecord {
	return &record.YearMethodNominalRecord{
		ID:                int(b.NominalID),
		Year:              b.Year,
		NominalName:       b.NominalName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s nominalRecordMapper) ToNominalsRecordYearMethodSuccess(b []*db.GetYearMethodNominalsSuccessRow) []*record.YearMethodNominalRecord {
	var result []*record.YearMethodNominalRecord

	for _, nominal := range b {
		result = append(result, s.ToNominalRecordYearMethodSuccess(nominal))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordYearMethodFailed(b *db.GetYearMethodNominalsFailedRow) *record.YearMethodNominalRecord {
	return &record.YearMethodNominalRecord{
		ID:                int(b.NominalID),
		Year:              b.Year,
		NominalName:       b.NominalName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s nominalRecordMapper) ToNominalsRecordYearMethodFailed(b []*db.GetYearMethodNominalsFailedRow) []*record.YearMethodNominalRecord {
	var result []*record.YearMethodNominalRecord

	for _, nominal := range b {
		result = append(result, s.ToNominalRecordYearMethodFailed(nominal))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordMonthAmountSuccessById(b *db.GetMonthAmountNominalsSuccessByIdRow) *record.MonthAmountNominalSuccessRecord {
	return &record.MonthAmountNominalSuccessRecord{
		ID:           int(b.NominalID),
		NominalName:  b.NominalName,
		Year:         b.MncYear,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *nominalRecordMapper) ToNominalsRecordMonthAmountSuccessById(b []*db.GetMonthAmountNominalsSuccessByIdRow) []*record.MonthAmountNominalSuccessRecord {
	var result []*record.MonthAmountNominalSuccessRecord

	for _, Bank := range b {
		result = append(result, s.ToNominalRecordMonthAmountSuccessById(Bank))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordYearAmountSuccessById(b *db.GetYearAmountNominalsSuccessByIdRow) *record.YearAmountNominalSuccessRecord {
	return &record.YearAmountNominalSuccessRecord{
		ID:           int(b.NominalID),
		NominalName:  b.NominalName,
		Year:         b.YncYear,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *nominalRecordMapper) ToNominalsRecordYearAmountSuccessById(b []*db.GetYearAmountNominalsSuccessByIdRow) []*record.YearAmountNominalSuccessRecord {
	var result []*record.YearAmountNominalSuccessRecord

	for _, Bank := range b {
		result = append(result, s.ToNominalRecordYearAmountSuccessById(Bank))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordMonthAmountFailedById(b *db.GetMonthAmountNominalsFailedByIdRow) *record.MonthAmountNominalFailedRecord {
	return &record.MonthAmountNominalFailedRecord{
		ID:          int(b.NominalID),
		NominalName: b.NominalName,
		Year:        b.MncYear,
		Month:       b.Month,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *nominalRecordMapper) ToNominalsRecordMonthAmountFailedById(b []*db.GetMonthAmountNominalsFailedByIdRow) []*record.MonthAmountNominalFailedRecord {
	var result []*record.MonthAmountNominalFailedRecord

	for _, Bank := range b {
		result = append(result, s.ToNominalRecordMonthAmountFailedById(Bank))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordYearAmountFailedById(b *db.GetYearAmountNominalsFailedByIdRow) *record.YearAmountNominalFailedRecord {
	return &record.YearAmountNominalFailedRecord{
		ID:          int(b.NominalID),
		NominalName: b.NominalName,
		Year:        b.YncYear,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *nominalRecordMapper) ToNominalsRecordYearAmountFaileById(b []*db.GetYearAmountNominalsFailedByIdRow) []*record.YearAmountNominalFailedRecord {
	var result []*record.YearAmountNominalFailedRecord

	for _, Bank := range b {
		result = append(result, s.ToNominalRecordYearAmountFailedById(Bank))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordMonthMethodSuccessById(b *db.GetMonthMethodNominalsSuccessByIdRow) *record.MonthMethodNominalRecord {
	return &record.MonthMethodNominalRecord{
		ID:                int(b.NominalID),
		Month:             b.Month,
		NominalName:       b.NominalName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s nominalRecordMapper) ToNominalsRecordMonthMethodSuccessById(b []*db.GetMonthMethodNominalsSuccessByIdRow) []*record.MonthMethodNominalRecord {
	var result []*record.MonthMethodNominalRecord

	for _, nominal := range b {
		result = append(result, s.ToNominalRecordMonthMethodSuccessById(nominal))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordMonthMethodFailedById(b *db.GetMonthMethodNominalsFailedByIdRow) *record.MonthMethodNominalRecord {
	return &record.MonthMethodNominalRecord{
		ID:                int(b.NominalID),
		Month:             b.Month,
		NominalName:       b.NominalName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s nominalRecordMapper) ToNominalsRecordMonthMethodFailedById(b []*db.GetMonthMethodNominalsFailedByIdRow) []*record.MonthMethodNominalRecord {
	var result []*record.MonthMethodNominalRecord

	for _, nominal := range b {
		result = append(result, s.ToNominalRecordMonthMethodFailedById(nominal))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordYearMethodSuccessById(b *db.GetYearMethodNominalsSuccessByIdRow) *record.YearMethodNominalRecord {
	return &record.YearMethodNominalRecord{
		ID:                int(b.NominalID),
		Year:              b.Year,
		NominalName:       b.NominalName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s nominalRecordMapper) ToNominalsRecordYearMethodSuccessById(b []*db.GetYearMethodNominalsSuccessByIdRow) []*record.YearMethodNominalRecord {
	var result []*record.YearMethodNominalRecord

	for _, nominal := range b {
		result = append(result, s.ToNominalRecordYearMethodSuccessById(nominal))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordYearMethodFailedById(b *db.GetYearMethodNominalsFailedByIdRow) *record.YearMethodNominalRecord {
	return &record.YearMethodNominalRecord{
		ID:                int(b.NominalID),
		Year:              b.Year,
		NominalName:       b.NominalName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s nominalRecordMapper) ToNominalsRecordYearMethodFailedById(b []*db.GetYearMethodNominalsFailedByIdRow) []*record.YearMethodNominalRecord {
	var result []*record.YearMethodNominalRecord

	for _, nominal := range b {
		result = append(result, s.ToNominalRecordYearMethodFailedById(nominal))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordMonthAmountSuccessByMerchant(b *db.GetMonthAmountNominalsSuccessByMerchantRow) *record.MonthAmountNominalSuccessRecord {
	return &record.MonthAmountNominalSuccessRecord{
		ID:           int(b.NominalID),
		NominalName:  b.NominalName,
		Year:         b.Year,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *nominalRecordMapper) ToNominalsRecordMonthAmountSuccessByMerchant(b []*db.GetMonthAmountNominalsSuccessByMerchantRow) []*record.MonthAmountNominalSuccessRecord {
	var result []*record.MonthAmountNominalSuccessRecord

	for _, Bank := range b {
		result = append(result, s.ToNominalRecordMonthAmountSuccessByMerchant(Bank))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordYearAmountSuccessByMerchant(b *db.GetYearAmountNominalsSuccessByMerchantRow) *record.YearAmountNominalSuccessRecord {
	return &record.YearAmountNominalSuccessRecord{
		ID:           int(b.NominalID),
		NominalName:  b.NominalName,
		Year:         b.Year,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *nominalRecordMapper) ToNominalsRecordYearAmountSuccessByMerchant(b []*db.GetYearAmountNominalsSuccessByMerchantRow) []*record.YearAmountNominalSuccessRecord {
	var result []*record.YearAmountNominalSuccessRecord

	for _, Bank := range b {
		result = append(result, s.ToNominalRecordYearAmountSuccessByMerchant(Bank))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordMonthAmountFailedByMerchant(b *db.GetMonthAmountNominalsFailedByMerchantRow) *record.MonthAmountNominalFailedRecord {
	return &record.MonthAmountNominalFailedRecord{
		ID:          int(b.NominalID),
		NominalName: b.NominalName,
		Year:        b.Year,
		Month:       b.Month,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *nominalRecordMapper) ToNominalsRecordMonthAmountFailedByMerchant(b []*db.GetMonthAmountNominalsFailedByMerchantRow) []*record.MonthAmountNominalFailedRecord {
	var result []*record.MonthAmountNominalFailedRecord

	for _, Bank := range b {
		result = append(result, s.ToNominalRecordMonthAmountFailedByMerchant(Bank))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordYearAmountFailedByMerchant(b *db.GetYearAmountNominalsFailedByMerchantRow) *record.YearAmountNominalFailedRecord {
	return &record.YearAmountNominalFailedRecord{
		ID:          int(b.NominalID),
		NominalName: b.NominalName,
		Year:        b.Year,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *nominalRecordMapper) ToNominalsRecordYearAmountFaileByMerchant(b []*db.GetYearAmountNominalsFailedByMerchantRow) []*record.YearAmountNominalFailedRecord {
	var result []*record.YearAmountNominalFailedRecord

	for _, Bank := range b {
		result = append(result, s.ToNominalRecordYearAmountFailedByMerchant(Bank))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordMonthMethodSuccessByMerchant(b *db.GetMonthMethodNominalsSuccessByMerchantRow) *record.MonthMethodNominalRecord {
	return &record.MonthMethodNominalRecord{
		ID:                int(b.NominalID),
		Month:             b.Month,
		NominalName:       b.NominalName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s nominalRecordMapper) ToNominalsRecordMonthMethodSuccessByMerchant(b []*db.GetMonthMethodNominalsSuccessByMerchantRow) []*record.MonthMethodNominalRecord {
	var result []*record.MonthMethodNominalRecord

	for _, nominal := range b {
		result = append(result, s.ToNominalRecordMonthMethodSuccessByMerchant(nominal))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordMonthMethodFailedByMerchant(b *db.GetMonthMethodNominalsFailedByMerchantRow) *record.MonthMethodNominalRecord {
	return &record.MonthMethodNominalRecord{
		ID:                int(b.NominalID),
		Month:             b.Month,
		NominalName:       b.NominalName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s nominalRecordMapper) ToNominalsRecordMonthMethodFailedByMerchant(b []*db.GetMonthMethodNominalsFailedByMerchantRow) []*record.MonthMethodNominalRecord {
	var result []*record.MonthMethodNominalRecord

	for _, nominal := range b {
		result = append(result, s.ToNominalRecordMonthMethodFailedByMerchant(nominal))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordYearMethodSuccessByMerchant(b *db.GetYearMethodNominalsSuccessByMerchantRow) *record.YearMethodNominalRecord {
	return &record.YearMethodNominalRecord{
		ID:                int(b.NominalID),
		Year:              b.Year,
		NominalName:       b.NominalName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s nominalRecordMapper) ToNominalsRecordYearMethodSuccessByMerchant(b []*db.GetYearMethodNominalsSuccessByMerchantRow) []*record.YearMethodNominalRecord {
	var result []*record.YearMethodNominalRecord

	for _, nominal := range b {
		result = append(result, s.ToNominalRecordYearMethodSuccessByMerchant(nominal))
	}

	return result
}

func (s *nominalRecordMapper) ToNominalRecordYearMethodFailedByMerchant(b *db.GetYearMethodNominalsFailedByMerchantRow) *record.YearMethodNominalRecord {
	return &record.YearMethodNominalRecord{
		ID:                int(b.NominalID),
		Year:              b.Year,
		NominalName:       b.NominalName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s nominalRecordMapper) ToNominalsRecordYearMethodFailedByMerchant(b []*db.GetYearMethodNominalsFailedByMerchantRow) []*record.YearMethodNominalRecord {
	var result []*record.YearMethodNominalRecord

	for _, nominal := range b {
		result = append(result, s.ToNominalRecordYearMethodFailedByMerchant(nominal))
	}

	return result
}
