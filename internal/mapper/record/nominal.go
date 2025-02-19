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
