package response_api

import (
	"topup_game/internal/domain/response"
	"topup_game/internal/pb"
)

type nominalResponseMapper struct {
}

func NewNominalResponseMapper() *nominalResponseMapper {
	return &nominalResponseMapper{}
}

func (s *nominalResponseMapper) ToApiResponseNominal(pbResponse *pb.ApiResponseNominal) *response.ApiResponseNominal {
	return &response.ApiResponseNominal{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseNominal(pbResponse.Data),
	}
}

func (s *nominalResponseMapper) ToApiResponseNominalAll(pbResponse *pb.ApiResponseNominalAll) *response.ApiResponseNominalAll {
	return &response.ApiResponseNominalAll{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

func (s *nominalResponseMapper) ToApiResponseNominalDelete(pbResponse *pb.ApiResponseNominalDelete) *response.ApiResponseNominalDelete {
	return &response.ApiResponseNominalDelete{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

func (s *nominalResponseMapper) ToApiResponseNominalDeleteAt(pbResponse *pb.ApiResponseNominalDeleteAt) *response.ApiResponseNominalDeleteAt {
	return &response.ApiResponseNominalDeleteAt{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseNominalDeleteAt(pbResponse.Data),
	}
}

func (s *nominalResponseMapper) ToApiResponsesNominal(pbResponse *pb.ApiResponsesNominal) *response.ApiResponsesNominal {
	return &response.ApiResponsesNominal{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesNominal(pbResponse.Data),
	}
}

func (s *nominalResponseMapper) ToApiResponsePaginationNominal(pbResponse *pb.ApiResponsePaginationNominal) *response.ApiResponsePaginationNominal {
	return &response.ApiResponsePaginationNominal{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.mapResponsesNominal(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
	}
}

func (s *nominalResponseMapper) ToApiResponsePaginationNominalDeleteAt(pbResponse *pb.ApiResponsePaginationNominalDeleteAt) *response.ApiResponsePaginationNominalDeleteAt {
	return &response.ApiResponsePaginationNominalDeleteAt{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.mapResponsesNominalDeleteAt(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
	}
}

func (s *nominalResponseMapper) mapResponseNominal(nominal *pb.NominalResponse) *response.NominalResponse {
	return &response.NominalResponse{
		ID:        int(nominal.Id),
		Name:      nominal.Name,
		Quantity:  int(nominal.Quantity),
		Price:     nominal.Price,
		CreatedAt: nominal.CreatedAt,
		UpdatedAt: nominal.UpdatedAt,
	}
}

func (s *nominalResponseMapper) mapResponsesNominal(nominals []*pb.NominalResponse) []*response.NominalResponse {
	var responseNominals []*response.NominalResponse

	for _, nominal := range nominals {
		responseNominals = append(responseNominals, s.mapResponseNominal(nominal))
	}

	return responseNominals
}

func (s *nominalResponseMapper) mapResponseNominalDeleteAt(nominal *pb.NominalResponseDeleteAt) *response.NominalResponseDeleteAt {
	return &response.NominalResponseDeleteAt{
		ID:        int(nominal.Id),
		Name:      nominal.Name,
		Quantity:  int(nominal.Quantity),
		Price:     nominal.Price,
		CreatedAt: nominal.CreatedAt,
		UpdatedAt: nominal.UpdatedAt,
		DeletedAt: nominal.DeletedAt,
	}
}

func (s *nominalResponseMapper) mapResponsesNominalDeleteAt(nominals []*pb.NominalResponseDeleteAt) []*response.NominalResponseDeleteAt {
	var responseNominals []*response.NominalResponseDeleteAt

	for _, nominal := range nominals {
		responseNominals = append(responseNominals, s.mapResponseNominalDeleteAt(nominal))
	}

	return responseNominals
}
