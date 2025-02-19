package service

import (
	"topup_game/internal/domain/requests"
	"topup_game/internal/domain/response"
	response_service "topup_game/internal/mapper/response/service"
	"topup_game/internal/repository"
	"topup_game/pkg/logger"

	"go.uber.org/zap"
)

type nominalService struct {
	nominalRepository repository.NominalRepository
	voucherRepository repository.VoucherRepository
	logger            logger.LoggerInterface
	mapping           response_service.NominalResponseMapper
}

func NewNominalService(nominalRepository repository.NominalRepository, voucherRepository repository.VoucherRepository, logger logger.LoggerInterface, mapping response_service.NominalResponseMapper) *nominalService {
	return &nominalService{
		nominalRepository: nominalRepository,
		voucherRepository: voucherRepository,
		logger:            logger,
		mapping:           mapping,
	}
}

func (s *nominalService) FindAll(page int, pageSize int, search string) ([]*response.NominalResponse, int, *response.ErrorResponse) {
	s.logger.Debug("Fetching nominals",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	banks, totalRecords, err := s.nominalRepository.FindAllNominals(page, pageSize, search)

	if err != nil {
		s.logger.Error("Failed to fetch nominals",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch nominals",
		}
	}

	nominalResponses := s.mapping.ToNominalsResponse(banks)

	s.logger.Debug("Successfully fetched banks",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return nominalResponses, int(totalRecords), nil
}

func (s *nominalService) FindByID(id int) (*response.NominalResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching nominals by id", zap.Int("nominal_id", id))

	user, err := s.nominalRepository.FindById(id)

	if err != nil {
		s.logger.Error("failed to find nominal by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "nominal not found",
		}
	}

	so := s.mapping.ToNominalResponse(user)

	s.logger.Debug("Successfully fetched nominal", zap.Int("nominal_id", id))

	return so, nil
}

func (s *nominalService) FindByActive(page int, pageSize int, search string) ([]*response.NominalResponseDeleteAt, int, *response.ErrorResponse) {
	s.logger.Debug("Fetching active nominals",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	users, totalRecords, err := s.nominalRepository.FindByActiveNominal(page, pageSize, search)

	if err != nil {
		s.logger.Error("Failed to fetch active nominals",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find active nominals",
		}
	}

	so := s.mapping.ToNominalsResponseDeleteAt(users)

	s.logger.Debug("Successfully fetched active nominals",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return so, totalRecords, nil
}

func (s *nominalService) FindByTrashed(page int, pageSize int, search string) ([]*response.NominalResponseDeleteAt, int, *response.ErrorResponse) {
	s.logger.Debug("Fetching trashed nominals",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	banks, totalRecords, err := s.nominalRepository.FindByTrashedNominal(page, pageSize, search)

	if err != nil {
		s.logger.Error("Failed to find trashed nominals", zap.Error(err))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find trashed nominals",
		}
	}

	so := s.mapping.ToNominalsResponseDeleteAt(banks)

	s.logger.Debug("Successfully fetched trashed nominals",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return so, totalRecords, nil
}

func (s *nominalService) Create(request *requests.CreateNominalRequest) (*response.NominalResponse, *response.ErrorResponse) {
	s.logger.Debug("Starting Create Nominal process",
		zap.String("name", request.Name),
		zap.Int("voucher_id", request.VoucherID),
	)

	_, err := s.voucherRepository.FindById(request.VoucherID)
	if err != nil {
		s.logger.Error("Voucher ID not found",
			zap.Int("voucher_id", request.VoucherID),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Invalid voucher ID",
		}
	}

	res, err := s.nominalRepository.CreateNominal(request)

	if err != nil {
		s.logger.Error("Failed to create nominal",
			zap.String("name", request.Name),
			zap.Int("voucher_id", request.VoucherID),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create nominal record",
		}
	}

	so := s.mapping.ToNominalResponse(res)

	s.logger.Debug("Create Nominal process completed",
		zap.String("name", request.Name),
		zap.Int("id", res.ID),
	)

	return so, nil
}

func (s *nominalService) Update(request *requests.UpdateNominalRequest) (*response.NominalResponse, *response.ErrorResponse) {
	s.logger.Debug("Starting Update Nominal process",
		zap.Int("id", request.ID),
		zap.String("name", request.Name),
		zap.Int("voucher_id", request.VoucherID),
	)

	if request.VoucherID != 0 {
		_, err := s.voucherRepository.FindById(request.VoucherID)
		if err != nil {
			s.logger.Error("Voucher ID not found",
				zap.Int("voucher_id", request.VoucherID),
				zap.Error(err),
			)

			return nil, &response.ErrorResponse{
				Status:  "error",
				Message: "Invalid voucher ID",
			}
		}
	}

	res, err := s.nominalRepository.UpdateNominal(request)

	if err != nil {
		s.logger.Error("Failed to update nominal",
			zap.Int("id", request.ID),
			zap.String("name", request.Name),
			zap.Int("voucher_id", request.VoucherID),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update nominal record",
		}
	}

	so := s.mapping.ToNominalResponse(res)

	s.logger.Debug("Update Nominal process completed",
		zap.Int("id", request.ID),
		zap.String("name", request.Name),
	)

	return so, nil
}

func (s *nominalService) Trashed(id int) (*response.NominalResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Starting Trashed Nominal process",
		zap.Int("nominal_id", id),
	)

	res, err := s.nominalRepository.TrashedNominal(id)

	if err != nil {
		s.logger.Error("Failed to move Nominal to trash",
			zap.Int("nominal_id", id),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trashed nominal record",
		}
	}

	so := s.mapping.ToNominalResponseDeleteAt(res)

	s.logger.Debug("TrashedNominal process completed",
		zap.Int("nominal_id", id),
	)

	return so, nil
}

func (s *nominalService) Restore(id int) (*response.NominalResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Starting RestoreNominal process",
		zap.Int("nominal_id", id),
	)

	res, err := s.nominalRepository.RestoreNominal(id)

	if err != nil {
		s.logger.Error("Failed to restore nominal", zap.Error(err))

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore nominal record",
		}
	}

	so := s.mapping.ToNominalResponseDeleteAt(res)

	s.logger.Debug("RestoreNominal process completed",
		zap.Int("nominal_id", id),
	)

	return so, nil
}

func (s *nominalService) DeletePermanent(id int) (bool, *response.ErrorResponse) {
	s.logger.Debug("Starting DeleteNominalPermanent process",
		zap.Int("nominal_id", id),
	)

	_, err := s.nominalRepository.DeleteNominalPermanent(id)

	if err != nil {
		s.logger.Error("Failed to delete nominal permanently",
			zap.Int("nominal_id", id),
			zap.Error(err),
		)

		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete nominal record",
		}
	}

	s.logger.Debug("DeleteNominalPermanent process completed",
		zap.Int("nominal_id", id),
	)

	return true, nil
}

func (s *nominalService) RestoreAll() (bool, *response.ErrorResponse) {
	s.logger.Debug("Restoring all nominals")

	_, err := s.nominalRepository.RestoreAllNominal()

	if err != nil {
		s.logger.Error("Failed to restore all nominals", zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all nominals: " + err.Error(),
		}
	}

	s.logger.Debug("Successfully restored all nominals")
	return true, nil
}

func (s *nominalService) DeleteAllPermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all nominals")

	_, err := s.nominalRepository.DeleteAllNominalsPermanent()

	if err != nil {
		s.logger.Error("Failed to permanently delete all nominals", zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete all nominals: " + err.Error(),
		}
	}

	s.logger.Debug("Successfully deleted all nominals permanently")
	return true, nil
}
