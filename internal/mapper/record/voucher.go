package recordmapper

import (
	"topup_game/internal/domain/record"
	db "topup_game/pkg/database/schema"
)

type voucherRecordMapper struct{}

func NewVoucherRecordMapper() *voucherRecordMapper {
	return &voucherRecordMapper{}
}

func (s *voucherRecordMapper) ToVoucherRecord(voucher *db.Voucher) *record.VoucherRecord {
	deletedAt := voucher.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.VoucherRecord{
		ID:         int(voucher.VoucherID),
		MerchantID: int(voucher.MerchantID),
		CategoryID: int(voucher.CategoryID),
		Name:       voucher.Name,
		ImageName:  voucher.ImageName,
		CreatedAt:  voucher.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:  voucher.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:  &deletedAt,
	}
}

func (s *voucherRecordMapper) ToVouchersRecord(vouchers []*db.Voucher) []*record.VoucherRecord {
	var result []*record.VoucherRecord

	for _, voucher := range vouchers {
		result = append(result, s.ToVoucherRecord(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordAll(voucher *db.GetVouchersRow) *record.VoucherRecord {
	deletedAt := voucher.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.VoucherRecord{
		ID:         int(voucher.VoucherID),
		MerchantID: int(voucher.MerchantID),
		CategoryID: int(voucher.CategoryID),
		Name:       voucher.Name,
		ImageName:  voucher.ImageName,
		CreatedAt:  voucher.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:  voucher.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:  &deletedAt,
	}
}

func (s *voucherRecordMapper) ToVouchersRecordAll(vouchers []*db.GetVouchersRow) []*record.VoucherRecord {
	var result []*record.VoucherRecord

	for _, voucher := range vouchers {
		result = append(result, s.ToVoucherRecordAll(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordActive(voucher *db.GetVouchersActiveRow) *record.VoucherRecord {
	deletedAt := voucher.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.VoucherRecord{
		ID:         int(voucher.VoucherID),
		MerchantID: int(voucher.MerchantID),
		CategoryID: int(voucher.CategoryID),
		Name:       voucher.Name,
		ImageName:  voucher.ImageName,
		CreatedAt:  voucher.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:  voucher.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:  &deletedAt,
	}
}

func (s *voucherRecordMapper) ToVouchersRecordActive(vouchers []*db.GetVouchersActiveRow) []*record.VoucherRecord {
	var result []*record.VoucherRecord

	for _, voucher := range vouchers {
		result = append(result, s.ToVoucherRecordActive(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordTrashed(voucher *db.GetVouchersTrashedRow) *record.VoucherRecord {
	deletedAt := voucher.DeletedAt.Time.Format("2006-01-02 15:04:05.000")

	return &record.VoucherRecord{
		ID:         int(voucher.VoucherID),
		MerchantID: int(voucher.MerchantID),
		CategoryID: int(voucher.CategoryID),
		Name:       voucher.Name,
		ImageName:  voucher.ImageName,
		CreatedAt:  voucher.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:  voucher.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:  &deletedAt,
	}
}

func (s *voucherRecordMapper) ToVouchersRecordTrashed(vouchers []*db.GetVouchersTrashedRow) []*record.VoucherRecord {
	var result []*record.VoucherRecord

	for _, voucher := range vouchers {
		result = append(result, s.ToVoucherRecordTrashed(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordMonthAmountSuccess(b *db.GetMonthAmountVouchersSuccessRow) *record.MonthAmountVoucherSuccessRecord {
	return &record.MonthAmountVoucherSuccessRecord{
		ID:           int(b.VoucherID),
		VoucherName:  b.VoucherName,
		Year:         b.MvcYear,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordMonthAmountSuccess(b []*db.GetMonthAmountVouchersSuccessRow) []*record.MonthAmountVoucherSuccessRecord {
	var result []*record.MonthAmountVoucherSuccessRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordMonthAmountSuccess(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordYearAmountSuccess(b *db.GetYearAmountVouchersSuccessRow) *record.YearAmountVoucherSuccessRecord {
	return &record.YearAmountVoucherSuccessRecord{
		ID:           int(b.VoucherID),
		VoucherName:  b.VoucherName,
		Year:         b.YvcYear,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordYearAmountSuccess(b []*db.GetYearAmountVouchersSuccessRow) []*record.YearAmountVoucherSuccessRecord {
	var result []*record.YearAmountVoucherSuccessRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordYearAmountSuccess(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordMonthAmountFailed(b *db.GetMonthAmountVouchersFailedRow) *record.MonthAmountVoucherFailedRecord {
	return &record.MonthAmountVoucherFailedRecord{
		ID:          int(b.VoucherID),
		VoucherName: b.VoucherName,
		Year:        b.MvcYear,
		Month:       b.Month,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordMonthAmountFailed(b []*db.GetMonthAmountVouchersFailedRow) []*record.MonthAmountVoucherFailedRecord {
	var result []*record.MonthAmountVoucherFailedRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordMonthAmountFailed(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordYearAmountFailed(b *db.GetYearAmountVouchersFailedRow) *record.YearAmountVoucherFailedRecord {
	return &record.YearAmountVoucherFailedRecord{
		ID:          int(b.VoucherID),
		VoucherName: b.VoucherName,
		Year:        b.YvcYear,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordYearAmountFailed(b []*db.GetYearAmountVouchersFailedRow) []*record.YearAmountVoucherFailedRecord {
	var result []*record.YearAmountVoucherFailedRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordYearAmountFailed(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordMonthMethodSuccess(b *db.GetMonthMethodVouchersSuccessRow) *record.MonthMethodVoucherRecord {
	return &record.MonthMethodVoucherRecord{
		ID:                int(b.VoucherID),
		Month:             b.Month,
		VoucherName:       b.VoucherName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordMonthMethodSuccess(b []*db.GetMonthMethodVouchersSuccessRow) []*record.MonthMethodVoucherRecord {
	var result []*record.MonthMethodVoucherRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordMonthMethodSuccess(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordMonthMethodFailed(b *db.GetMonthMethodVouchersFailedRow) *record.MonthMethodVoucherRecord {
	return &record.MonthMethodVoucherRecord{
		ID:                int(b.VoucherID),
		Month:             b.Month,
		VoucherName:       b.VoucherName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordMonthMethodFailed(b []*db.GetMonthMethodVouchersFailedRow) []*record.MonthMethodVoucherRecord {
	var result []*record.MonthMethodVoucherRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordMonthMethodFailed(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordYearMethodSuccess(b *db.GetYearMethodVouchersSuccessRow) *record.YearMethodVoucherRecord {
	return &record.YearMethodVoucherRecord{
		ID:                int(b.VoucherID),
		Year:              b.Year,
		VoucherName:       b.VoucherName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordYearMethodSuccess(b []*db.GetYearMethodVouchersSuccessRow) []*record.YearMethodVoucherRecord {
	var result []*record.YearMethodVoucherRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordYearMethodSuccess(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordYearMethodFailed(b *db.GetYearMethodVouchersFailedRow) *record.YearMethodVoucherRecord {
	return &record.YearMethodVoucherRecord{
		ID:                int(b.VoucherID),
		Year:              b.Year,
		VoucherName:       b.VoucherName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordYearMethodFailed(b []*db.GetYearMethodVouchersFailedRow) []*record.YearMethodVoucherRecord {
	var result []*record.YearMethodVoucherRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordYearMethodFailed(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordMonthAmountSuccessById(b *db.GetMonthAmountVouchersSuccessByIdRow) *record.MonthAmountVoucherSuccessRecord {
	return &record.MonthAmountVoucherSuccessRecord{
		ID:           int(b.VoucherID),
		VoucherName:  b.VoucherName,
		Year:         b.MvcYear,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordMonthAmountSuccessById(b []*db.GetMonthAmountVouchersSuccessByIdRow) []*record.MonthAmountVoucherSuccessRecord {
	var result []*record.MonthAmountVoucherSuccessRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordMonthAmountSuccessById(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordYearAmountSuccessById(b *db.GetYearAmountVouchersSuccessByIdRow) *record.YearAmountVoucherSuccessRecord {
	return &record.YearAmountVoucherSuccessRecord{
		ID:           int(b.VoucherID),
		VoucherName:  b.VoucherName,
		Year:         b.YvcYear,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordYearAmountSuccessById(b []*db.GetYearAmountVouchersSuccessByIdRow) []*record.YearAmountVoucherSuccessRecord {
	var result []*record.YearAmountVoucherSuccessRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordYearAmountSuccessById(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordMonthAmountFailedById(b *db.GetMonthAmountVouchersFailedByIdRow) *record.MonthAmountVoucherFailedRecord {
	return &record.MonthAmountVoucherFailedRecord{
		ID:          int(b.VoucherID),
		VoucherName: b.VoucherName,
		Year:        b.MvcYear,
		Month:       b.Month,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordMonthAmountFailedById(b []*db.GetMonthAmountVouchersFailedByIdRow) []*record.MonthAmountVoucherFailedRecord {
	var result []*record.MonthAmountVoucherFailedRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordMonthAmountFailedById(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordYearAmountFailedById(b *db.GetYearAmountVouchersFailedByIdRow) *record.YearAmountVoucherFailedRecord {
	return &record.YearAmountVoucherFailedRecord{
		ID:          int(b.VoucherID),
		VoucherName: b.VoucherName,
		Year:        b.YvcYear,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordYearAmountFailedById(b []*db.GetYearAmountVouchersFailedByIdRow) []*record.YearAmountVoucherFailedRecord {
	var result []*record.YearAmountVoucherFailedRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordYearAmountFailedById(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordMonthMethodSuccessById(b *db.GetMonthMethodVouchersSuccessByIdRow) *record.MonthMethodVoucherRecord {
	return &record.MonthMethodVoucherRecord{
		ID:                int(b.VoucherID),
		Month:             b.Month,
		VoucherName:       b.VoucherName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordMonthMethodSuccessById(b []*db.GetMonthMethodVouchersSuccessByIdRow) []*record.MonthMethodVoucherRecord {
	var result []*record.MonthMethodVoucherRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordMonthMethodSuccessById(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordMonthMethodFailedById(b *db.GetMonthMethodVouchersFailedByIdRow) *record.MonthMethodVoucherRecord {
	return &record.MonthMethodVoucherRecord{
		ID:                int(b.VoucherID),
		Month:             b.Month,
		VoucherName:       b.VoucherName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordMonthMethodFailedById(b []*db.GetMonthMethodVouchersFailedByIdRow) []*record.MonthMethodVoucherRecord {
	var result []*record.MonthMethodVoucherRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordMonthMethodFailedById(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordYearMethodSuccessById(b *db.GetYearMethodVouchersSuccessByIdRow) *record.YearMethodVoucherRecord {
	return &record.YearMethodVoucherRecord{
		ID:                int(b.VoucherID),
		Year:              b.Year,
		VoucherName:       b.VoucherName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordYearMethodSuccessById(b []*db.GetYearMethodVouchersSuccessByIdRow) []*record.YearMethodVoucherRecord {
	var result []*record.YearMethodVoucherRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordYearMethodSuccessById(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordYearMethodFailedById(b *db.GetYearMethodVouchersFailedByIdRow) *record.YearMethodVoucherRecord {
	return &record.YearMethodVoucherRecord{
		ID:                int(b.VoucherID),
		Year:              b.Year,
		VoucherName:       b.VoucherName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordYearMethodFailedById(b []*db.GetYearMethodVouchersFailedByIdRow) []*record.YearMethodVoucherRecord {
	var result []*record.YearMethodVoucherRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordYearMethodFailedById(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordMonthAmountSuccessByMerchant(b *db.GetMonthAmountVouchersSuccessByMerchantRow) *record.MonthAmountVoucherSuccessRecord {
	return &record.MonthAmountVoucherSuccessRecord{
		ID:           int(b.VoucherID),
		VoucherName:  b.VoucherName,
		Year:         b.Year,
		Month:        b.Month,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordMonthAmountSuccessByMerchant(b []*db.GetMonthAmountVouchersSuccessByMerchantRow) []*record.MonthAmountVoucherSuccessRecord {
	var result []*record.MonthAmountVoucherSuccessRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordMonthAmountSuccessByMerchant(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordYearAmountSuccessByMerchant(b *db.GetYearAmountVouchersSuccessByMerchantRow) *record.YearAmountVoucherSuccessRecord {
	return &record.YearAmountVoucherSuccessRecord{
		ID:           int(b.VoucherID),
		VoucherName:  b.VoucherName,
		Year:         b.Year,
		TotalSuccess: int(b.TotalSuccess),
		TotalAmount:  int(b.TotalAmount),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordYearAmountSuccessByMerchant(b []*db.GetYearAmountVouchersSuccessByMerchantRow) []*record.YearAmountVoucherSuccessRecord {
	var result []*record.YearAmountVoucherSuccessRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordYearAmountSuccessByMerchant(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordMonthAmountFailedByMerchant(b *db.GetMonthAmountVouchersFailedByMerchantRow) *record.MonthAmountVoucherFailedRecord {
	return &record.MonthAmountVoucherFailedRecord{
		ID:          int(b.VoucherID),
		VoucherName: b.VoucherName,
		Year:        b.Year,
		Month:       b.Month,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordMonthAmountFailedByMerchant(b []*db.GetMonthAmountVouchersFailedByMerchantRow) []*record.MonthAmountVoucherFailedRecord {
	var result []*record.MonthAmountVoucherFailedRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordMonthAmountFailedByMerchant(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordYearAmountFailedByMerchant(b *db.GetYearAmountVouchersFailedByMerchantRow) *record.YearAmountVoucherFailedRecord {
	return &record.YearAmountVoucherFailedRecord{
		ID:          int(b.VoucherID),
		VoucherName: b.VoucherName,
		Year:        b.Year,
		TotalFailed: int(b.TotalFailed),
		TotalAmount: int(b.TotalAmount),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordYearAmountFailedByMerchant(b []*db.GetYearAmountVouchersFailedByMerchantRow) []*record.YearAmountVoucherFailedRecord {
	var result []*record.YearAmountVoucherFailedRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordYearAmountFailedByMerchant(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordMonthMethodSuccessByMerchant(b *db.GetMonthMethodVouchersSuccessByMerchantRow) *record.MonthMethodVoucherRecord {
	return &record.MonthMethodVoucherRecord{
		ID:                int(b.VoucherID),
		Month:             b.Month,
		VoucherName:       b.VoucherName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordMonthMethodSuccessByMerchant(b []*db.GetMonthMethodVouchersSuccessByMerchantRow) []*record.MonthMethodVoucherRecord {
	var result []*record.MonthMethodVoucherRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordMonthMethodSuccessByMerchant(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordMonthMethodFailedByMerchant(b *db.GetMonthMethodVouchersFailedByMerchantRow) *record.MonthMethodVoucherRecord {
	return &record.MonthMethodVoucherRecord{
		ID:                int(b.VoucherID),
		Month:             b.Month,
		VoucherName:       b.VoucherName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordMonthMethodFailedByMerchant(b []*db.GetMonthMethodVouchersFailedByMerchantRow) []*record.MonthMethodVoucherRecord {
	var result []*record.MonthMethodVoucherRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordMonthMethodFailedByMerchant(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordYearMethodSuccessByMerchant(b *db.GetYearMethodVouchersSuccessByMerchantRow) *record.YearMethodVoucherRecord {
	return &record.YearMethodVoucherRecord{
		ID:                int(b.VoucherID),
		Year:              b.Year,
		VoucherName:       b.VoucherName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordYearMethodSuccessByMerchant(b []*db.GetYearMethodVouchersSuccessByMerchantRow) []*record.YearMethodVoucherRecord {
	var result []*record.YearMethodVoucherRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordYearMethodSuccessByMerchant(voucher))
	}

	return result
}

func (s *voucherRecordMapper) ToVoucherRecordYearMethodFailedByMerchant(b *db.GetYearMethodVouchersFailedByMerchantRow) *record.YearMethodVoucherRecord {
	return &record.YearMethodVoucherRecord{
		ID:                int(b.VoucherID),
		Year:              b.Year,
		VoucherName:       b.VoucherName,
		PaymentMethod:     b.PaymentMethod,
		TotalAmount:       int(b.TotalAmount),
		TotalTransactions: int(b.TotalTransactions),
	}
}

func (s *voucherRecordMapper) ToVouchersRecordYearMethodFailedByMerchant(b []*db.GetYearMethodVouchersFailedByMerchantRow) []*record.YearMethodVoucherRecord {
	var result []*record.YearMethodVoucherRecord

	for _, voucher := range b {
		result = append(result, s.ToVoucherRecordYearMethodFailedByMerchant(voucher))
	}

	return result
}
