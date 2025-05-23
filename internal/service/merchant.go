package service

import (
	"topup_game/internal/domain/requests"
	"topup_game/internal/domain/response"
	response_service "topup_game/internal/mapper/response/service"
	"topup_game/internal/repository"
	"topup_game/pkg/errors/merchant_errors"
	"topup_game/pkg/logger"

	"go.uber.org/zap"
)

type merchantService struct {
	merchantRepository repository.MerchantRepository
	logger             logger.LoggerInterface
	mapping            response_service.MerchantResponseMapper
}

func NewMerchantService(
	merchantRepository repository.MerchantRepository,
	logger logger.LoggerInterface,
	mapping response_service.MerchantResponseMapper,
) *merchantService {
	return &merchantService{
		merchantRepository: merchantRepository,
		logger:             logger,
		mapping:            mapping,
	}
}

func (s *merchantService) FindAll(request *requests.FindAllMerchants) ([]*response.MerchantResponse, *int, *response.ErrorResponse) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	s.logger.Debug("Fetching all merchants",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	merchants, totalRecords, err := s.merchantRepository.FindAllMerchants(request)
	if err != nil {
		s.logger.Error("Failed to fetch merchants", zap.Error(err))
		return nil, nil, merchant_errors.ErrFailedFindAll
	}

	return s.mapping.ToMerchantsResponse(merchants), totalRecords, nil
}

func (s *merchantService) FindByActive(request *requests.FindAllMerchants) ([]*response.MerchantResponseDeleteAt, *int, *response.ErrorResponse) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	s.logger.Debug("Fetching active merchants", zap.String("search", search), zap.Int("page", page), zap.Int("pageSize", pageSize))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	merchants, totalRecords, err := s.merchantRepository.FindByActive(request)
	if err != nil {
		s.logger.Error("Failed to fetch active merchants", zap.Error(err))
		return nil, nil, merchant_errors.ErrFailedFindActive
	}

	return s.mapping.ToMerchantsResponseDeleteAt(merchants), totalRecords, nil
}

func (s *merchantService) FindByTrashed(request *requests.FindAllMerchants) ([]*response.MerchantResponseDeleteAt, *int, *response.ErrorResponse) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	s.logger.Debug("Fetching trashed merchants", zap.String("search", search), zap.Int("page", page), zap.Int("pageSize", pageSize))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	merchants, totalRecords, err := s.merchantRepository.FindByTrashed(request)
	if err != nil {
		s.logger.Error("Failed to fetch trashed merchants", zap.Error(err))
		return nil, nil, merchant_errors.ErrFailedFindTrashed
	}

	return s.mapping.ToMerchantsResponseDeleteAt(merchants), totalRecords, nil
}

func (s *merchantService) FindById(merchantID int) (*response.MerchantResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching merchant by ID", zap.Int("merchantID", merchantID))

	merchant, err := s.merchantRepository.FindById(merchantID)
	if err != nil {
		s.logger.Error("Merchant not found", zap.Error(err))
		return nil, merchant_errors.ErrMerchantNotFoundRes
	}

	return s.mapping.ToMerchantResponse(merchant), nil
}

func (s *merchantService) Create(req *requests.CreateMerchantRequest) (*response.MerchantResponse, *response.ErrorResponse) {
	s.logger.Debug("Creating new merchant")

	merchant, err := s.merchantRepository.CreateMerchant(req)
	if err != nil {
		s.logger.Error("Failed to create merchant", zap.Error(err))
		return nil, merchant_errors.ErrFailedCreateMerchant
	}

	return s.mapping.ToMerchantResponse(merchant), nil
}

func (s *merchantService) Update(req *requests.UpdateMerchantRequest) (*response.MerchantResponse, *response.ErrorResponse) {
	s.logger.Debug("Updating merchant", zap.Int("merchantID", req.MerchantID))

	merchant, err := s.merchantRepository.UpdateMerchant(req)
	if err != nil {
		s.logger.Error("Failed to update merchant", zap.Error(err))
		return nil, merchant_errors.ErrFailedUpdateMerchant
	}

	return s.mapping.ToMerchantResponse(merchant), nil
}

func (s *merchantService) Trashed(merchantID int) (*response.MerchantResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Trashing merchant", zap.Int("merchantID", merchantID))

	merchant, err := s.merchantRepository.TrashedMerchant(merchantID)
	if err != nil {
		s.logger.Error("Failed to trash merchant", zap.Error(err))
		return nil, merchant_errors.ErrFailedTrashedMerchant
	}

	return s.mapping.ToMerchantResponseDeleteAt(merchant), nil
}

func (s *merchantService) Restore(merchantID int) (*response.MerchantResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Restoring merchant", zap.Int("merchantID", merchantID))

	merchant, err := s.merchantRepository.RestoreMerchant(merchantID)
	if err != nil {
		s.logger.Error("Failed to restore merchant", zap.Error(err))
		return nil, merchant_errors.ErrFailedRestoreMerchant
	}

	return s.mapping.ToMerchantResponseDeleteAt(merchant), nil
}

func (s *merchantService) DeletePermanent(merchantID int) (bool, *response.ErrorResponse) {
	s.logger.Debug("Deleting merchant permanently", zap.Int("merchantID", merchantID))

	success, err := s.merchantRepository.DeleteMerchantPermanent(merchantID)
	if err != nil {
		s.logger.Error("Failed to delete merchant permanently", zap.Error(err))
		return false, merchant_errors.ErrFailedDeletePermanent
	}

	return success, nil
}

func (s *merchantService) RestoreAll() (bool, *response.ErrorResponse) {
	s.logger.Debug("Restoring all trashed merchants")

	success, err := s.merchantRepository.RestoreAllMerchant()
	if err != nil {
		s.logger.Error("Failed to restore all merchants", zap.Error(err))
		return false, merchant_errors.ErrFailedRestoreAll
	}

	return success, nil
}

func (s *merchantService) DeleteAllPermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all merchants")

	success, err := s.merchantRepository.DeleteAllMerchantPermanent()
	if err != nil {
		s.logger.Error("Failed to permanently delete all merchants", zap.Error(err))
		return false, merchant_errors.ErrFailedDeleteAll
	}

	return success, nil
}
