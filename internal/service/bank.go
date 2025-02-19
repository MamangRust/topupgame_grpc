package service

import (
	"topup_game/internal/domain/requests"
	"topup_game/internal/domain/response"
	response_service "topup_game/internal/mapper/response/service"
	"topup_game/internal/repository"
	"topup_game/pkg/logger"

	"go.uber.org/zap"
)

type bankService struct {
	bankRepository repository.BankRepository
	logger         logger.LoggerInterface
	mapping        response_service.BankResponseMapper
}

func NewBankService(
	bankRepository repository.BankRepository,
	logger logger.LoggerInterface,
	mapping response_service.BankResponseMapper,
) *bankService {
	return &bankService{
		bankRepository: bankRepository,
		logger:         logger,
		mapping:        mapping,
	}
}

func (s *bankService) FindAll(page int, pageSize int, search string) ([]*response.BankResponse, int, *response.ErrorResponse) {
	s.logger.Debug("Fetching banks",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	banks, totalRecords, err := s.bankRepository.FindAllBanks(page, pageSize, search)

	if err != nil {
		s.logger.Error("Failed to fetch bank",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch banks",
		}
	}

	bankResponses := s.mapping.ToBanksResponse(banks)

	s.logger.Debug("Successfully fetched banks",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return bankResponses, int(totalRecords), nil
}

func (s *bankService) FindByID(id int) (*response.BankResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching banks by id", zap.Int("user_id", id))

	user, err := s.bankRepository.FindById(id)

	if err != nil {
		s.logger.Error("failed to find bank by ID", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Bank not found",
		}
	}

	so := s.mapping.ToBankResponse(user)

	s.logger.Debug("Successfully fetched bank", zap.Int("bank_id", id))

	return so, nil
}

func (s *bankService) FindByActive(page int, pageSize int, search string) ([]*response.BankResponseDeleteAt, int, *response.ErrorResponse) {
	s.logger.Debug("Fetching active banks",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	users, totalRecords, err := s.bankRepository.FindByActiveBanks(page, pageSize, search)

	if err != nil {
		s.logger.Error("Failed to fetch active banks",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find active banks",
		}
	}

	so := s.mapping.ToBanksResponseDeleteAt(users)

	s.logger.Debug("Successfully fetched active banks",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return so, totalRecords, nil
}

func (s *bankService) FindByTrashed(page int, pageSize int, search string) ([]*response.BankResponseDeleteAt, int, *response.ErrorResponse) {
	s.logger.Debug("Fetching trashed banks",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	banks, totalRecords, err := s.bankRepository.FindByTrashedBanks(page, pageSize, search)

	if err != nil {
		s.logger.Error("Failed to find trashed banks", zap.Error(err))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to find trashed banks",
		}
	}

	so := s.mapping.ToBanksResponseDeleteAt(banks)

	s.logger.Debug("Successfully fetched trashed banks",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return so, totalRecords, nil
}

func (s *bankService) Create(request *requests.CreateBankRequest) (*response.BankResponse, *response.ErrorResponse) {
	s.logger.Debug("Starting Create Bank process",
		zap.String("name", request.Name),
	)

	res, err := s.bankRepository.CreateBank(request)

	if err != nil {
		s.logger.Error("Failed to create bank",
			zap.String("name", request.Name),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create bank record",
		}
	}

	so := s.mapping.ToBankResponse(res)

	s.logger.Debug("Create Bank process completed",
		zap.String("name", request.Name),
		zap.Int("id", res.ID),
	)

	return so, nil
}

func (s *bankService) Update(request *requests.UpdateBankRequest) (*response.BankResponse, *response.ErrorResponse) {
	s.logger.Debug("Starting Update bank process",
		zap.Int("id", request.ID),
		zap.String("name", request.Name),
	)

	res, err := s.bankRepository.UpdateBank(request)

	if err != nil {
		s.logger.Error("Failed to update bank",
			zap.Int("id", request.ID),
			zap.String("name", request.Name),
			zap.Error(err))

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update bank record",
		}
	}

	so := s.mapping.ToBankResponse(res)

	s.logger.Debug("UpdateBank process completed",
		zap.Int("id", request.ID),
		zap.String("name", request.Name),
	)

	return so, nil
}

func (s *bankService) Trashed(id int) (*response.BankResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Starting Trashed Bank process",
		zap.Int("roleID", id),
	)

	res, err := s.bankRepository.TrashedBank(id)

	if err != nil {
		s.logger.Error("Failed to move Bank to trash",
			zap.Int("bank_id", id),
			zap.Error(err))

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trashed bank record",
		}
	}

	so := s.mapping.ToBankResponseDeleteAt(res)

	s.logger.Debug("TrashedBank process completed",
		zap.Int("bank_id", id),
	)

	return so, nil
}

func (s *bankService) Restore(id int) (*response.BankResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Starting RestoreBank process",
		zap.Int("bank_id", id),
	)

	res, err := s.bankRepository.RestoreBank(id)

	if err != nil {
		s.logger.Error("Failed to restore bank", zap.Error(err))

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore role record",
		}
	}

	so := s.mapping.ToBankResponseDeleteAt(res)

	s.logger.Debug("RestoreBank process completed",
		zap.Int("bank_id", id),
	)

	return so, nil
}

func (s *bankService) DeletePermanent(id int) (bool, *response.ErrorResponse) {
	s.logger.Debug("Starting DeleteBankPermanent process",
		zap.Int("bank_id", id),
	)

	_, err := s.bankRepository.DeleteBankPermanent(id)

	if err != nil {
		s.logger.Error("Failed to delete bank permanently",
			zap.Int("bank_id", id),
			zap.Error(err),
		)

		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete bank record",
		}
	}

	s.logger.Debug("DeleteBankPermanent process completed",
		zap.Int("bank_id", id),
	)

	return true, nil
}

func (s *bankService) RestoreAll() (bool, *response.ErrorResponse) {
	s.logger.Debug("Restoring all banks")

	_, err := s.bankRepository.RestoreAllBanks()

	if err != nil {
		s.logger.Error("Failed to restore all banks", zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all banks: " + err.Error(),
		}
	}

	s.logger.Debug("Successfully restored all banks")
	return true, nil
}

func (s *bankService) DeleteAllPermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all banks")

	_, err := s.bankRepository.DeleteAllBanksPermanent()

	if err != nil {
		s.logger.Error("Failed to permanently delete all banks", zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete all banks: " + err.Error(),
		}
	}

	s.logger.Debug("Successfully deleted all banks permanently")
	return true, nil
}
