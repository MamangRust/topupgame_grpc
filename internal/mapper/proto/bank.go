package protomapper

import (
	"topup_game/internal/domain/response"
	"topup_game/internal/pb"
)

type bankProtoMapper struct {
}

func NewBankProtoMapper() *bankProtoMapper {
	return &bankProtoMapper{}
}

func (s *bankProtoMapper) ToProtoResponseBankAll(status string, message string) *pb.ApiResponseBankAll {
	return &pb.ApiResponseBankAll{
		Status:  status,
		Message: message,
	}
}

func (s *bankProtoMapper) ToProtoResponseBankDelete(status string, message string) *pb.ApiResponseBankDelete {
	return &pb.ApiResponseBankDelete{
		Status:  status,
		Message: message,
	}
}

func (s *bankProtoMapper) ToProtoResponseBank(status string, message string, pbResponse *response.BankResponse) *pb.ApiResponseBank {
	return &pb.ApiResponseBank{
		Status:  status,
		Message: message,
		Data:    s.mapResponseBank(pbResponse),
	}
}

func (s *bankProtoMapper) ToProtoResponseBankDeleteAt(status string, message string, pbResponse *response.BankResponseDeleteAt) *pb.ApiResponseBankDeleteAt {
	return &pb.ApiResponseBankDeleteAt{
		Status:  status,
		Message: message,
		Data:    s.mapResponseBankDeleteAt(pbResponse),
	}
}

func (s *bankProtoMapper) ToProtoResponsesBank(status string, message string, pbResponse []*response.BankResponse) *pb.ApiResponsesBank {
	return &pb.ApiResponsesBank{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesBank(pbResponse),
	}
}

func (s *bankProtoMapper) ToProtoResponsePaginationBank(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.BankResponse) *pb.ApiResponsePaginationBank {
	return &pb.ApiResponsePaginationBank{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesBank(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *bankProtoMapper) ToProtoResponsePaginationBankDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.BankResponseDeleteAt) *pb.ApiResponsePaginationBankDeleteAt {
	return &pb.ApiResponsePaginationBankDeleteAt{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesBankDeleteAt(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *bankProtoMapper) mapResponseBank(Bank *response.BankResponse) *pb.BankResponse {
	return &pb.BankResponse{
		Id:        int32(Bank.ID),
		Name:      Bank.Name,
		CreatedAt: Bank.CreatedAt,
		UpdatedAt: Bank.UpdatedAt,
	}
}

func (s *bankProtoMapper) mapResponsesBank(Banks []*response.BankResponse) []*pb.BankResponse {
	var responseBanks []*pb.BankResponse

	for _, Bank := range Banks {
		responseBanks = append(responseBanks, s.mapResponseBank(Bank))
	}

	return responseBanks
}

func (s *bankProtoMapper) mapResponseBankDeleteAt(Bank *response.BankResponseDeleteAt) *pb.BankResponseDeleteAt {
	return &pb.BankResponseDeleteAt{
		Id:        int32(Bank.ID),
		Name:      Bank.Name,
		CreatedAt: Bank.CreatedAt,
		UpdatedAt: Bank.UpdatedAt,
		DeletedAt: Bank.DeletedAt,
	}
}

func (s *bankProtoMapper) mapResponsesBankDeleteAt(Banks []*response.BankResponseDeleteAt) []*pb.BankResponseDeleteAt {
	var responseBanks []*pb.BankResponseDeleteAt

	for _, Bank := range Banks {
		responseBanks = append(responseBanks, s.mapResponseBankDeleteAt(Bank))
	}

	return responseBanks
}
