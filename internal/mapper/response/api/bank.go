package response_api

import (
	"topup_game/internal/domain/response"
	"topup_game/internal/pb"
)

type bankResponseMapper struct {
}

func NewBankResponseMapper() *bankResponseMapper {
	return &bankResponseMapper{}
}

func (s *bankResponseMapper) ToApiResponseBankAll(pbResponse *pb.ApiResponseBankAll) *response.ApiResponseBankAll {
	return &response.ApiResponseBankAll{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

func (s *bankResponseMapper) ToApiResponseBankDelete(pbResponse *pb.ApiResponseBankDelete) *response.ApiResponseBankDelete {
	return &response.ApiResponseBankDelete{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

func (s *bankResponseMapper) ToApiResponseBank(pbResponse *pb.ApiResponseBank) *response.ApiResponseBank {
	return &response.ApiResponseBank{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseBank(pbResponse.Data),
	}
}

func (s *bankResponseMapper) ToApiResponseBankDeleteAt(pbResponse *pb.ApiResponseBankDeleteAt) *response.ApiResponseBankDeleteAt {
	return &response.ApiResponseBankDeleteAt{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseBankDeleteAt(pbResponse.Data),
	}
}

func (s *bankResponseMapper) ToApiResponsesBank(pbResponse *pb.ApiResponsesBank) *response.ApiResponsesBank {
	return &response.ApiResponsesBank{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesBank(pbResponse.Data),
	}
}

func (s *bankResponseMapper) ToApiResponsePaginationBank(pbResponse *pb.ApiResponsePaginationBank) *response.ApiResponsePaginationBank {
	return &response.ApiResponsePaginationBank{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.mapResponsesBank(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
	}
}

func (s *bankResponseMapper) ToApiResponsePaginationBankDeleteAt(pbResponse *pb.ApiResponsePaginationBankDeleteAt) *response.ApiResponsePaginationBankDeleteAt {
	return &response.ApiResponsePaginationBankDeleteAt{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.mapResponsesBankDeleteAt(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
	}
}

func (s *bankResponseMapper) mapResponseBank(Bank *pb.BankResponse) *response.BankResponse {
	return &response.BankResponse{
		ID:        int(Bank.Id),
		Name:      Bank.Name,
		CreatedAt: Bank.CreatedAt,
		UpdatedAt: Bank.UpdatedAt,
	}
}

func (s *bankResponseMapper) mapResponsesBank(Banks []*pb.BankResponse) []*response.BankResponse {
	var responseBanks []*response.BankResponse

	for _, Bank := range Banks {
		responseBanks = append(responseBanks, s.mapResponseBank(Bank))
	}

	return responseBanks
}

func (s *bankResponseMapper) mapResponseBankDeleteAt(Bank *pb.BankResponseDeleteAt) *response.BankResponseDeleteAt {
	return &response.BankResponseDeleteAt{
		ID:        int(Bank.Id),
		Name:      Bank.Name,
		CreatedAt: Bank.CreatedAt,
		UpdatedAt: Bank.UpdatedAt,
		DeletedAt: Bank.DeletedAt,
	}
}

func (s *bankResponseMapper) mapResponsesBankDeleteAt(Banks []*pb.BankResponseDeleteAt) []*response.BankResponseDeleteAt {
	var responseBanks []*response.BankResponseDeleteAt

	for _, Bank := range Banks {
		responseBanks = append(responseBanks, s.mapResponseBankDeleteAt(Bank))
	}

	return responseBanks
}
