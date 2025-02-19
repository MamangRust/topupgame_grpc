package service

import (
	"topup_game/internal/domain/requests"
	"topup_game/internal/domain/response"
	response_service "topup_game/internal/mapper/response/service"
	"topup_game/internal/repository"
	"topup_game/pkg/logger"

	"go.uber.org/zap"
)

type voucherService struct {
	merchantRepository repository.MerchantRepository
	categoryRepository repository.CategoryRepository
	voucherRepository  repository.VoucherRepository
	logger             logger.LoggerInterface
	mapping            response_service.VoucherResponseMapper
}

func NewVoucherService(
	merchantRepository repository.MerchantRepository,
	categoryRepository repository.CategoryRepository,
	voucherRepository repository.VoucherRepository,
	logger logger.LoggerInterface,
	mapping response_service.VoucherResponseMapper,
) *voucherService {
	return &voucherService{
		merchantRepository: merchantRepository,
		categoryRepository: categoryRepository,
		voucherRepository:  voucherRepository,
		logger:             logger,
		mapping:            mapping,
	}
}

func (s *voucherService) FindAll(page int, pageSize int, search string) ([]*response.VoucherResponse, int, *response.ErrorResponse) {
	s.logger.Debug("Fetching voucher",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.voucherRepository.FindAllVouchers(page, pageSize, search)
	if err != nil {
		s.logger.Error("Failed to fetch voucher",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch voucher records",
		}
	}

	s.logger.Debug("Successfully fetched voucher",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	so := s.mapping.ToVouchersResponse(res)

	return so, totalRecords, nil
}

func (s *voucherService) FindById(id int) (*response.VoucherResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching voucher by ID", zap.Int("id", id))

	res, err := s.voucherRepository.FindById(id)

	if err != nil {
		s.logger.Error("Failed to fetch voucher record by ID", zap.Error(err))

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch voucher record by ID",
		}
	}

	s.logger.Debug("Successfully fetched voucher", zap.Int("id", id))

	so := s.mapping.ToVoucherResponse(res)

	return so, nil
}

func (s *voucherService) FindByActive(page int, pageSize int, search string) ([]*response.VoucherResponseDeleteAt, int, *response.ErrorResponse) {
	s.logger.Debug("Fetching active voucher",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.voucherRepository.FindByActiveVouchers(page, pageSize, search)
	if err != nil {
		s.logger.Error("Failed to fetch active voucher",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch voucher records",
		}
	}

	s.logger.Debug("Successfully fetched active voucher",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	so := s.mapping.ToVouchersResponseDeleteAt(res)

	return so, totalRecords, nil
}

func (s *voucherService) FindByTrashed(page int, pageSize int, search string) ([]*response.VoucherResponseDeleteAt, int, *response.ErrorResponse) {
	s.logger.Debug("Fetching trashed voucher",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.voucherRepository.FindByTrashedVoucher(page, pageSize, search)

	if err != nil {
		s.logger.Error("Failed to fetch trashed voucher",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch voucher records",
		}
	}

	s.logger.Debug("Successfully fetched trashed voucher",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	so := s.mapping.ToVouchersResponseDeleteAt(res)

	return so, totalRecords, nil
}

func (s *voucherService) Create(request *requests.CreateVoucherRequest) (*response.VoucherResponse, *response.ErrorResponse) {
	s.logger.Debug("Starting Create Voucher process",
		zap.String("name", request.Name),
		zap.Int("merchant_id", request.MerchantID),
		zap.Int("category_id", request.CategoryID),
	)

	res, err := s.voucherRepository.CreateVoucher(request)

	if err != nil {
		s.logger.Error("Failed to create voucher",
			zap.String("name", request.Name),
			zap.Int("merchant_id", request.MerchantID),
			zap.Int("category_id", request.CategoryID),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create voucher record",
		}
	}

	so := s.mapping.ToVoucherResponse(res)

	s.logger.Debug("Create Voucher process completed",
		zap.String("name", request.Name),
		zap.Int("id", res.ID),
	)

	return so, nil
}

func (s *voucherService) Update(request *requests.UpdateVoucherRequest) (*response.VoucherResponse, *response.ErrorResponse) {
	s.logger.Debug("Starting Update Voucher process",
		zap.Int("id", request.ID),
		zap.String("name", request.Name),
		zap.Int("merchant_id", request.MerchantID),
		zap.Int("category_id", request.CategoryID),
	)

	res, err := s.voucherRepository.UpdateVoucher(request)

	if err != nil {
		s.logger.Error("Failed to update voucher",
			zap.Int("id", request.ID),
			zap.String("name", request.Name),
			zap.Int("merchant_id", request.MerchantID),
			zap.Int("category_id", request.CategoryID),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update voucher record",
		}
	}

	so := s.mapping.ToVoucherResponse(res)

	s.logger.Debug("Update Voucher process completed",
		zap.Int("id", request.ID),
		zap.String("name", request.Name),
	)

	return so, nil
}

func (s *voucherService) Trashed(id int) (*response.VoucherResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Starting Trashed Voucher process",
		zap.Int("voucher_id", id),
	)

	res, err := s.voucherRepository.TrashVoucher(id)

	if err != nil {
		s.logger.Error("Failed to move voucher to trash",
			zap.Int("voucher_id", id),
			zap.Error(err),
		)

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trashed voucher record",
		}
	}

	so := s.mapping.ToVoucherResponseDeleteAt(res)

	s.logger.Debug("TrashedCategory process completed",
		zap.Int("category_id", id),
	)

	return so, nil
}

func (s *voucherService) Restore(id int) (*response.VoucherResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Starting Restore Voucher process",
		zap.Int("voucher_id", id),
	)

	res, err := s.voucherRepository.RestoreVoucher(id)

	if err != nil {
		s.logger.Error("Failed to restore voucher", zap.Error(err))

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore voucher record",
		}
	}

	so := s.mapping.ToVoucherResponseDeleteAt(res)

	s.logger.Debug("RestoreCategory process completed",
		zap.Int("category_id", id),
	)

	return so, nil
}

func (s *voucherService) DeletePermanent(id int) (bool, *response.ErrorResponse) {
	s.logger.Debug("Starting DeleteVoucherPermanent process",
		zap.Int("voucher_id", id),
	)

	_, err := s.voucherRepository.DeleteVoucherPermanent(id)

	if err != nil {
		s.logger.Error("Failed to delete voucher permanently",
			zap.Int("voucher_id", id),
			zap.Error(err),
		)

		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete voucher record",
		}
	}

	s.logger.Debug("DeleteVoucherPermanent process completed",
		zap.Int("voucher_id", id),
	)

	return true, nil
}

func (s *voucherService) RestoreAll() (bool, *response.ErrorResponse) {
	s.logger.Debug("Restoring all Voucher")

	_, err := s.voucherRepository.RestoreAllVouchers()

	if err != nil {
		s.logger.Error("Failed to restore all Voucher", zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all Voucher: " + err.Error(),
		}
	}

	s.logger.Debug("Successfully restored all Voucher")
	return true, nil
}

func (s *voucherService) DeleteAllPermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all Voucher")

	_, err := s.voucherRepository.DeleteAllVouchersPermanent()

	if err != nil {
		s.logger.Error("Failed to permanently delete all Voucher", zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete all Voucher: " + err.Error(),
		}
	}

	s.logger.Debug("Successfully deleted all Voucher permanently")
	return true, nil
}
