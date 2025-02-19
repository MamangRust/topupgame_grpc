package protomapper

import (
	"topup_game/internal/domain/response"
	"topup_game/internal/pb"
)

type nominalProtoMapper struct{}

func NewNominalProtoMapper() *nominalProtoMapper {
	return &nominalProtoMapper{}
}

func (s *nominalProtoMapper) ToProtoResponseNominalAll(status string, message string) *pb.ApiResponseNominalAll {
	return &pb.ApiResponseNominalAll{
		Status:  status,
		Message: message,
	}
}

func (s *nominalProtoMapper) ToProtoResponseNominalDelete(status string, message string) *pb.ApiResponseNominalDelete {
	return &pb.ApiResponseNominalDelete{
		Status:  status,
		Message: message,
	}
}

func (s *nominalProtoMapper) ToProtoResponseNominal(status string, message string, pbResponse *response.NominalResponse) *pb.ApiResponseNominal {
	return &pb.ApiResponseNominal{
		Status:  status,
		Message: message,
		Data:    s.mapResponseNominal(pbResponse),
	}
}

func (s *nominalProtoMapper) ToProtoResponseNominalDeleteAt(status string, message string, pbResponse *response.NominalResponseDeleteAt) *pb.ApiResponseNominalDeleteAt {
	return &pb.ApiResponseNominalDeleteAt{
		Status:  status,
		Message: message,
		Data:    s.mapResponseNominalDeleteAt(pbResponse),
	}
}

func (s *nominalProtoMapper) ToProtoResponsesNominal(status string, message string, pbResponse []*response.NominalResponse) *pb.ApiResponsesNominal {
	return &pb.ApiResponsesNominal{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesNominal(pbResponse),
	}
}

func (s *nominalProtoMapper) ToProtoResponsePaginationNominal(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.NominalResponse) *pb.ApiResponsePaginationNominal {
	return &pb.ApiResponsePaginationNominal{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesNominal(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *nominalProtoMapper) ToProtoResponsePaginationNominalDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.NominalResponseDeleteAt) *pb.ApiResponsePaginationNominalDeleteAt {
	return &pb.ApiResponsePaginationNominalDeleteAt{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesNominalDeleteAt(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (s *nominalProtoMapper) mapResponseNominal(nominal *response.NominalResponse) *pb.NominalResponse {
	return &pb.NominalResponse{
		Id:        int32(nominal.ID),
		Name:      nominal.Name,
		Quantity:  int32(nominal.Quantity),
		Price:     nominal.Price,
		CreatedAt: nominal.CreatedAt,
		UpdatedAt: nominal.UpdatedAt,
	}
}

func (s *nominalProtoMapper) mapResponsesNominal(nominals []*response.NominalResponse) []*pb.NominalResponse {
	var responseNominals []*pb.NominalResponse

	for _, nominal := range nominals {
		responseNominals = append(responseNominals, s.mapResponseNominal(nominal))
	}

	return responseNominals
}

func (s *nominalProtoMapper) mapResponseNominalDeleteAt(nominal *response.NominalResponseDeleteAt) *pb.NominalResponseDeleteAt {
	return &pb.NominalResponseDeleteAt{
		Id:        int32(nominal.ID),
		Name:      nominal.Name,
		Quantity:  int32(nominal.Quantity),
		Price:     nominal.Price,
		CreatedAt: nominal.CreatedAt,
		UpdatedAt: nominal.UpdatedAt,
		DeletedAt: nominal.DeletedAt,
	}
}

func (s *nominalProtoMapper) mapResponsesNominalDeleteAt(nominals []*response.NominalResponseDeleteAt) []*pb.NominalResponseDeleteAt {
	var responseNominals []*pb.NominalResponseDeleteAt

	for _, nominal := range nominals {
		responseNominals = append(responseNominals, s.mapResponseNominalDeleteAt(nominal))
	}

	return responseNominals
}
