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
