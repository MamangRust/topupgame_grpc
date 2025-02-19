package response_service

import (
	"topup_game/internal/domain/record"
	"topup_game/internal/domain/response"
)

type bankResponseMapper struct {
}

func NewBankResponseMapper() *bankResponseMapper {
	return &bankResponseMapper{}
}

func (s *bankResponseMapper) ToBankResponse(Bank *record.BankRecord) *response.BankResponse {
	return &response.BankResponse{
		ID:        Bank.ID,
		Name:      Bank.Name,
		CreatedAt: Bank.CreatedAt,
		UpdatedAt: Bank.UpdatedAt,
	}
}

func (s *bankResponseMapper) ToBanksResponse(Banks []*record.BankRecord) []*response.BankResponse {
	var responseBanks []*response.BankResponse

	for _, Bank := range Banks {
		responseBanks = append(responseBanks, s.ToBankResponse(Bank))
	}

	return responseBanks
}

func (s *bankResponseMapper) ToBankResponseDeleteAt(Bank *record.BankRecord) *response.BankResponseDeleteAt {
	return &response.BankResponseDeleteAt{
		ID:        Bank.ID,
		Name:      Bank.Name,
		CreatedAt: Bank.CreatedAt,
		UpdatedAt: Bank.UpdatedAt,
		DeletedAt: *Bank.DeletedAt,
	}
}

func (s *bankResponseMapper) ToBanksResponseDeleteAt(Banks []*record.BankRecord) []*response.BankResponseDeleteAt {
	var responseBanks []*response.BankResponseDeleteAt

	for _, Bank := range Banks {
		responseBanks = append(responseBanks, s.ToBankResponseDeleteAt(Bank))
	}

	return responseBanks
}
