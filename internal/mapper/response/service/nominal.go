package response_service

import (
	"topup_game/internal/domain/record"
	"topup_game/internal/domain/response"
)

type nominalResponseMapper struct {
}

func NewNominalResponseMapper() *nominalResponseMapper {
	return &nominalResponseMapper{}
}

func (s *nominalResponseMapper) ToNominalResponse(nominal *record.NominalRecord) *response.NominalResponse {
	return &response.NominalResponse{
		ID:        nominal.ID,
		VoucherID: nominal.VoucherID,
		Name:      nominal.Name,
		Quantity:  nominal.Quantity,
		Price:     nominal.Price,
		CreatedAt: nominal.CreatedAt,
		UpdatedAt: nominal.UpdatedAt,
		DeletedAt: nominal.DeletedAt,
	}
}

func (s *nominalResponseMapper) ToNominalsResponse(nominals []*record.NominalRecord) []*response.NominalResponse {
	var responseNominals []*response.NominalResponse

	for _, nominal := range nominals {
		responseNominals = append(responseNominals, s.ToNominalResponse(nominal))
	}

	return responseNominals
}

func (s *nominalResponseMapper) ToNominalResponseDeleteAt(nominal *record.NominalRecord) *response.NominalResponseDeleteAt {
	return &response.NominalResponseDeleteAt{
		ID:        nominal.ID,
		VoucherID: nominal.VoucherID,
		Name:      nominal.Name,
		Quantity:  nominal.Quantity,
		Price:     nominal.Price,
		CreatedAt: nominal.CreatedAt,
		UpdatedAt: nominal.UpdatedAt,
		DeletedAt: *nominal.DeletedAt,
	}
}

func (s *nominalResponseMapper) ToNominalsResponseDeleteAt(nominals []*record.NominalRecord) []*response.NominalResponseDeleteAt {
	var responseNominals []*response.NominalResponseDeleteAt

	for _, nominal := range nominals {
		responseNominals = append(responseNominals, s.ToNominalResponseDeleteAt(nominal))
	}

	return responseNominals
}
