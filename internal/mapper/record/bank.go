package recordmapper

import (
	"topup_game/internal/domain/record"
	db "topup_game/pkg/database/schema"
)

type bankRecordMapper struct {
}

func NewBankRecordMapper() *bankRecordMapper {
	return &bankRecordMapper{}
}

func (s *bankRecordMapper) ToBankRecord(Bank *db.Bank) *record.BankRecord {
	deletedAt := Bank.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.BankRecord{
		ID:        int(Bank.BankID),
		Name:      Bank.Name,
		CreatedAt: Bank.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt: Bank.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt: &deletedAt,
	}
}

func (s *bankRecordMapper) ToBanksRecord(Banks []*db.Bank) []*record.BankRecord {
	var result []*record.BankRecord

	for _, Bank := range Banks {
		result = append(result, s.ToBankRecord(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordAll(Bank *db.GetBanksRow) *record.BankRecord {
	deletedAt := Bank.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.BankRecord{
		ID:        int(Bank.BankID),
		Name:      Bank.Name,
		CreatedAt: Bank.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt: Bank.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt: &deletedAt,
	}
}

func (s *bankRecordMapper) ToBanksRecordAll(Banks []*db.GetBanksRow) []*record.BankRecord {
	var result []*record.BankRecord

	for _, Bank := range Banks {
		result = append(result, s.ToBankRecordAll(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordActive(Bank *db.GetBanksActiveRow) *record.BankRecord {
	deletedAt := Bank.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.BankRecord{
		ID:        int(Bank.BankID),
		Name:      Bank.Name,
		CreatedAt: Bank.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt: Bank.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt: &deletedAt,
	}
}

func (s *bankRecordMapper) ToBanksRecordActive(Banks []*db.GetBanksActiveRow) []*record.BankRecord {
	var result []*record.BankRecord

	for _, Bank := range Banks {
		result = append(result, s.ToBankRecordActive(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordTrashed(Bank *db.GetBanksTrashedRow) *record.BankRecord {
	deletedAt := Bank.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.BankRecord{
		ID:        int(Bank.BankID),
		Name:      Bank.Name,
		CreatedAt: Bank.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt: Bank.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt: &deletedAt,
	}
}

func (s *bankRecordMapper) ToBanksRecordTrashed(Banks []*db.GetBanksTrashedRow) []*record.BankRecord {
	var result []*record.BankRecord

	for _, Bank := range Banks {
		result = append(result, s.ToBankRecordTrashed(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordMonthAmountSuccess(b *db.GetMonthAmountBankSuccessRow) *record.MonthAmountBankSuccessRecord {
	return &record.MonthAmountBankSuccessRecord{
		ID:           int(b.BankID),
		BankName:     b.BankName,
		Year:         b.MbcYear,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *bankRecordMapper) ToBanksRecordMonthAmountSuccess(b []*db.GetMonthAmountBankSuccessRow) []*record.MonthAmountBankSuccessRecord {
	var result []*record.MonthAmountBankSuccessRecord

	for _, Bank := range b {
		result = append(result, s.ToBankRecordMonthAmountSuccess(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordYearAmountSuccess(b *db.GetYearAmountBankSuccessRow) *record.YearAmountBankSuccessRecord {
	return &record.YearAmountBankSuccessRecord{
		ID:           int(b.BankID),
		BankName:     b.BankName,
		Year:         b.YbcYear,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *bankRecordMapper) ToBanksRecordYearAmountSuccess(b []*db.GetYearAmountBankSuccessRow) []*record.YearAmountBankSuccessRecord {
	var result []*record.YearAmountBankSuccessRecord

	for _, Bank := range b {
		result = append(result, s.ToBankRecordYearAmountSuccess(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordMonthAmountFailed(b *db.GetMonthAmountBankFailedRow) *record.MonthAmountBankFailedRecord {
	return &record.MonthAmountBankFailedRecord{
		ID:          int(b.BankID),
		BankName:    b.BankName,
		Year:        b.MbcYear,
		Month:       b.Month,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *bankRecordMapper) ToBanksRecordMonthAmountFailed(b []*db.GetMonthAmountBankFailedRow) []*record.MonthAmountBankFailedRecord {
	var result []*record.MonthAmountBankFailedRecord

	for _, Bank := range b {
		result = append(result, s.ToBankRecordMonthAmountFailed(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordYearAmountFailed(b *db.GetYearAmountBankFailedRow) *record.YearAmountBankFailedRecord {
	return &record.YearAmountBankFailedRecord{
		ID:          int(b.BankID),
		BankName:    b.BankName,
		Year:        b.YbcYear,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *bankRecordMapper) ToBanksRecordYearAmountFailed(b []*db.GetYearAmountBankFailedRow) []*record.YearAmountBankFailedRecord {
	var result []*record.YearAmountBankFailedRecord

	for _, Bank := range b {
		result = append(result, s.ToBankRecordYearAmountFailed(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordMonthMethodSuccess(b *db.GetMonthBankMethodsSuccessRow) *record.MonthMethodBankRecord {
	return &record.MonthMethodBankRecord{
		ID:                int(b.BankID),
		Month:             b.Month,
		BankName:          b.BankName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s bankRecordMapper) ToBanksRecordMonthMethodSuccess(b []*db.GetMonthBankMethodsSuccessRow) []*record.MonthMethodBankRecord {
	var result []*record.MonthMethodBankRecord

	for _, Bank := range b {
		result = append(result, s.ToBankRecordMonthMethodSuccess(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordMonthMethodFailed(b *db.GetMonthBankMethodsFailedRow) *record.MonthMethodBankRecord {
	return &record.MonthMethodBankRecord{
		ID:                int(b.BankID),
		Month:             b.Month,
		BankName:          b.BankName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s bankRecordMapper) ToBanksRecordMonthMethodFailed(b []*db.GetMonthBankMethodsFailedRow) []*record.MonthMethodBankRecord {
	var result []*record.MonthMethodBankRecord

	for _, Bank := range b {
		result = append(result, s.ToBankRecordMonthMethodFailed(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordYearMethodSuccess(b *db.GetYearBankMethodsSuccessRow) *record.YearMethodBankRecord {
	return &record.YearMethodBankRecord{
		ID:                int(b.BankID),
		Year:              b.Year,
		BankName:          b.BankName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s bankRecordMapper) ToBanksRecordYearMethodSuccess(b []*db.GetYearBankMethodsSuccessRow) []*record.YearMethodBankRecord {
	var result []*record.YearMethodBankRecord

	for _, bank := range b {
		result = append(result, s.ToBankRecordYearMethodSuccess(bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordYearMethodFailed(b *db.GetYearBankMethodsFailedRow) *record.YearMethodBankRecord {
	return &record.YearMethodBankRecord{
		ID:                int(b.BankID),
		Year:              b.Year,
		BankName:          b.BankName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s bankRecordMapper) ToBanksRecordYearMethodFailed(b []*db.GetYearBankMethodsFailedRow) []*record.YearMethodBankRecord {
	var result []*record.YearMethodBankRecord

	for _, bank := range b {
		result = append(result, s.ToBankRecordYearMethodFailed(bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordMonthAmountSuccessById(b *db.GetMonthAmountBankSuccessByIdRow) *record.MonthAmountBankSuccessRecord {
	return &record.MonthAmountBankSuccessRecord{
		ID:           int(b.BankID),
		BankName:     b.BankName,
		Year:         b.Year,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *bankRecordMapper) ToBanksRecordMonthAmountSuccessById(b []*db.GetMonthAmountBankSuccessByIdRow) []*record.MonthAmountBankSuccessRecord {
	var result []*record.MonthAmountBankSuccessRecord

	for _, Bank := range b {
		result = append(result, s.ToBankRecordMonthAmountSuccessById(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordYearAmountSuccessById(b *db.GetYearAmountBankSuccessByIdRow) *record.YearAmountBankSuccessRecord {
	return &record.YearAmountBankSuccessRecord{
		ID:           int(b.BankID),
		BankName:     b.BankName,
		Year:         b.Year,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *bankRecordMapper) ToBanksRecordYearAmountSuccessById(b []*db.GetYearAmountBankSuccessByIdRow) []*record.YearAmountBankSuccessRecord {
	var result []*record.YearAmountBankSuccessRecord

	for _, Bank := range b {
		result = append(result, s.ToBankRecordYearAmountSuccessById(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordMonthAmountFailedById(b *db.GetMonthAmountBankFailedByIdRow) *record.MonthAmountBankFailedRecord {
	return &record.MonthAmountBankFailedRecord{
		ID:          int(b.BankID),
		BankName:    b.BankName,
		Year:        b.Year,
		Month:       b.Month,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *bankRecordMapper) ToBanksRecordMonthAmountFailedById(b []*db.GetMonthAmountBankFailedByIdRow) []*record.MonthAmountBankFailedRecord {
	var result []*record.MonthAmountBankFailedRecord

	for _, Bank := range b {
		result = append(result, s.ToBankRecordMonthAmountFailedById(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordYearAmountFailedById(b *db.GetYearAmountBankFailedByIdRow) *record.YearAmountBankFailedRecord {
	return &record.YearAmountBankFailedRecord{
		ID:          int(b.BankID),
		BankName:    b.BankName,
		Year:        b.Year,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *bankRecordMapper) ToBanksRecordYearAmountFaileById(b []*db.GetYearAmountBankFailedByIdRow) []*record.YearAmountBankFailedRecord {
	var result []*record.YearAmountBankFailedRecord

	for _, Bank := range b {
		result = append(result, s.ToBankRecordYearAmountFailedById(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordMonthMethodSuccessById(b *db.GetMonthBankMethodsSuccessByIdRow) *record.MonthMethodBankRecord {
	return &record.MonthMethodBankRecord{
		ID:                int(b.BankID),
		Month:             b.Month,
		BankName:          b.BankName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s bankRecordMapper) ToBanksRecordMonthMethodSuccessById(b []*db.GetMonthBankMethodsSuccessByIdRow) []*record.MonthMethodBankRecord {
	var result []*record.MonthMethodBankRecord

	for _, Bank := range b {
		result = append(result, s.ToBankRecordMonthMethodSuccessById(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordMonthMethodFailedById(b *db.GetMonthBankMethodsFailedByIdRow) *record.MonthMethodBankRecord {
	return &record.MonthMethodBankRecord{
		ID:                int(b.BankID),
		Month:             b.Month,
		BankName:          b.BankName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s bankRecordMapper) ToBanksRecordMonthMethodFailedById(b []*db.GetMonthBankMethodsFailedByIdRow) []*record.MonthMethodBankRecord {
	var result []*record.MonthMethodBankRecord

	for _, Bank := range b {
		result = append(result, s.ToBankRecordMonthMethodFailedById(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordYearMethodSuccessById(b *db.GetYearBankMethodsSuccessByIdRow) *record.YearMethodBankRecord {
	return &record.YearMethodBankRecord{
		ID:                int(b.BankID),
		Year:              b.Year,
		BankName:          b.BankName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s bankRecordMapper) ToBanksRecordYearMethodSuccessById(b []*db.GetYearBankMethodsSuccessByIdRow) []*record.YearMethodBankRecord {
	var result []*record.YearMethodBankRecord

	for _, bank := range b {
		result = append(result, s.ToBankRecordYearMethodSuccessById(bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordYearMethodFailedById(b *db.GetYearBankMethodsFailedByIdRow) *record.YearMethodBankRecord {
	return &record.YearMethodBankRecord{
		ID:                int(b.BankID),
		Year:              b.Year,
		BankName:          b.BankName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s bankRecordMapper) ToBanksRecordYearMethodFailedById(b []*db.GetYearBankMethodsFailedByIdRow) []*record.YearMethodBankRecord {
	var result []*record.YearMethodBankRecord

	for _, bank := range b {
		result = append(result, s.ToBankRecordYearMethodFailedById(bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordMonthAmountSuccessByMerchant(b *db.GetMonthAmountBankSuccessByMerchantRow) *record.MonthAmountBankSuccessRecord {
	return &record.MonthAmountBankSuccessRecord{
		ID:           int(b.BankID),
		BankName:     b.BankName,
		Year:         b.BmcYear,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *bankRecordMapper) ToBanksRecordMonthAmountSuccessByMerchant(b []*db.GetMonthAmountBankSuccessByMerchantRow) []*record.MonthAmountBankSuccessRecord {
	var result []*record.MonthAmountBankSuccessRecord

	for _, Bank := range b {
		result = append(result, s.ToBankRecordMonthAmountSuccessByMerchant(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordYearAmountSuccessByMerchant(b *db.GetYearAmountBankSuccessByMerchantRow) *record.YearAmountBankSuccessRecord {
	return &record.YearAmountBankSuccessRecord{
		ID:           int(b.BankID),
		BankName:     b.BankName,
		Year:         b.BycYear,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *bankRecordMapper) ToBanksRecordYearAmountSuccessByMerchant(b []*db.GetYearAmountBankSuccessByMerchantRow) []*record.YearAmountBankSuccessRecord {
	var result []*record.YearAmountBankSuccessRecord

	for _, Bank := range b {
		result = append(result, s.ToBankRecordYearAmountSuccessByMerchant(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordMonthAmountFailedByMerchant(b *db.GetMonthAmountBankFailedByMerchantRow) *record.MonthAmountBankFailedRecord {
	return &record.MonthAmountBankFailedRecord{
		ID:          int(b.BankID),
		BankName:    b.BankName,
		Year:        b.BmcYear,
		Month:       b.Month,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *bankRecordMapper) ToBanksRecordMonthAmountFailedByMerchant(b []*db.GetMonthAmountBankFailedByMerchantRow) []*record.MonthAmountBankFailedRecord {
	var result []*record.MonthAmountBankFailedRecord

	for _, Bank := range b {
		result = append(result, s.ToBankRecordMonthAmountFailedByMerchant(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordYearAmountFailedByMerchant(b *db.GetYearAmountBankFailedByMerchantRow) *record.YearAmountBankFailedRecord {
	return &record.YearAmountBankFailedRecord{
		ID:          int(b.BankID),
		BankName:    b.BankName,
		Year:        b.BycYear,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *bankRecordMapper) ToBanksRecordYearAmountFaileByMerchant(b []*db.GetYearAmountBankFailedByMerchantRow) []*record.YearAmountBankFailedRecord {
	var result []*record.YearAmountBankFailedRecord

	for _, Bank := range b {
		result = append(result, s.ToBankRecordYearAmountFailedByMerchant(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordMonthMethodSuccessByMerchant(b *db.GetMonthBankMethodsSuccessByMerchantRow) *record.MonthMethodBankRecord {
	return &record.MonthMethodBankRecord{
		ID:                int(b.BankID),
		Month:             b.Month,
		BankName:          b.BankName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s bankRecordMapper) ToBanksRecordMonthMethodSuccessByMerchant(b []*db.GetMonthBankMethodsSuccessByMerchantRow) []*record.MonthMethodBankRecord {
	var result []*record.MonthMethodBankRecord

	for _, Bank := range b {
		result = append(result, s.ToBankRecordMonthMethodSuccessByMerchant(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordMonthMethodFailedByMerchant(b *db.GetMonthBankMethodsFailedByMerchantRow) *record.MonthMethodBankRecord {
	return &record.MonthMethodBankRecord{
		ID:                int(b.BankID),
		Month:             b.Month,
		BankName:          b.BankName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s bankRecordMapper) ToBanksRecordMonthMethodFailedByMerchant(b []*db.GetMonthBankMethodsFailedByMerchantRow) []*record.MonthMethodBankRecord {
	var result []*record.MonthMethodBankRecord

	for _, Bank := range b {
		result = append(result, s.ToBankRecordMonthMethodFailedByMerchant(Bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordYearMethodSuccessByMerchant(b *db.GetYearBankMethodsSuccessByMerchantRow) *record.YearMethodBankRecord {
	return &record.YearMethodBankRecord{
		ID:                int(b.BankID),
		Year:              b.Year,
		BankName:          b.BankName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s bankRecordMapper) ToBanksRecordYearMethodSuccessByMerchant(b []*db.GetYearBankMethodsSuccessByMerchantRow) []*record.YearMethodBankRecord {
	var result []*record.YearMethodBankRecord

	for _, bank := range b {
		result = append(result, s.ToBankRecordYearMethodSuccessByMerchant(bank))
	}

	return result
}

func (s *bankRecordMapper) ToBankRecordYearMethodFailedByMerchant(b *db.GetYearBankMethodsFailedByMerchantRow) *record.YearMethodBankRecord {
	return &record.YearMethodBankRecord{
		ID:                int(b.BankID),
		Year:              b.Year,
		BankName:          b.BankName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s bankRecordMapper) ToBanksRecordYearMethodFailedByMerchant(b []*db.GetYearBankMethodsFailedByMerchantRow) []*record.YearMethodBankRecord {
	var result []*record.YearMethodBankRecord

	for _, bank := range b {
		result = append(result, s.ToBankRecordYearMethodFailedByMerchant(bank))
	}

	return result
}
