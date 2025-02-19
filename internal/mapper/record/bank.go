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
