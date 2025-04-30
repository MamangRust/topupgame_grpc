package service

import (
	"topup_game/internal/domain/requests"
	"topup_game/internal/domain/response"
	response_service "topup_game/internal/mapper/response/service"
	"topup_game/internal/repository"
	"topup_game/pkg/errors/voucher_errors"
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

func (s *voucherService) FindAll(request *requests.FindAllVouchers) ([]*response.VoucherResponse, *int, *response.ErrorResponse) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

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

	res, totalRecords, err := s.voucherRepository.FindAllVouchers(request)
	if err != nil {
		s.logger.Error("Failed to fetch voucher",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, nil, voucher_errors.ErrFailedFindAll
	}

	s.logger.Debug("Successfully fetched voucher",
		zap.Int("totalRecords", *totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	so := s.mapping.ToVouchersResponse(res)

	return so, totalRecords, nil
}

func (s *voucherService) FindMonthAmountVoucherSuccess(req *requests.MonthAmountVoucherRequest) ([]*response.MonthAmountVoucherSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly voucher success amounts", zap.Any("request", req))

	records, err := s.voucherRepository.FindMonthAmountVoucherSuccess(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly voucher success amounts", zap.Error(err))
		return nil, voucher_errors.ErrFailedFindMonthAmountVoucherSuccess
	}

	responses := s.mapping.ToVouchersResponseMonthAmountSuccess(records)

	s.logger.Debug("Successfully fetched monthly voucher success amounts",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindYearAmountVoucherSuccess(year int) ([]*response.YearAmountVoucherSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly voucher success amounts", zap.Int("year", year))

	records, err := s.voucherRepository.FindYearAmountVoucherSuccess(year)
	if err != nil {
		s.logger.Error("Failed to fetch yearly voucher success amounts", zap.Error(err))
		return nil, voucher_errors.ErrFailedFindYearAmountVoucherSuccess
	}

	responses := s.mapping.ToVouchersResponseYearAmountSuccess(records)

	s.logger.Debug("Successfully fetched yearly voucher success amounts",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindMonthAmountVoucherFailed(req *requests.MonthAmountVoucherRequest) ([]*response.MonthAmountVoucherFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly voucher failed amounts", zap.Any("request", req))

	records, err := s.voucherRepository.FindMonthAmountVoucherFailed(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly voucher failed amounts", zap.Error(err))
		return nil, voucher_errors.ErrFailedFindMonthAmountVoucherFailed
	}

	responses := s.mapping.ToVouchersResponseMonthAmountFailed(records)

	s.logger.Debug("Successfully fetched monthly voucher failed amounts",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindYearAmountVoucherFailed(year int) ([]*response.YearAmountVoucherFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly voucher failed amounts", zap.Int("year", year))

	records, err := s.voucherRepository.FindYearAmountVoucherFailed(year)
	if err != nil {
		s.logger.Error("Failed to fetch yearly voucher failed amounts", zap.Error(err))
		return nil, voucher_errors.ErrFailedFindYearAmountVoucherFailed
	}

	responses := s.mapping.ToVouchersResponseYearAmountFailed(records)

	s.logger.Debug("Successfully fetched yearly voucher failed amounts",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindMonthMethodVoucherSuccess(year int) ([]*response.MonthMethodVoucherResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly voucher success methods", zap.Int("year", year))

	records, err := s.voucherRepository.FindMonthMethodVoucherSuccess(year)
	if err != nil {
		s.logger.Error("Failed to fetch monthly voucher success methods", zap.Error(err))
		return nil, voucher_errors.ErrFailedFindMonthMethodVoucherSuccess
	}

	responses := s.mapping.ToVouchersResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly voucher success methods",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindYearMethodVoucherSuccess(year int) ([]*response.YearMethodVoucherResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly voucher success methods", zap.Int("year", year))

	records, err := s.voucherRepository.FindYearMethodVoucherSuccess(year)
	if err != nil {
		s.logger.Error("Failed to fetch yearly voucher success methods", zap.Error(err))
		return nil, voucher_errors.ErrFailedFindYearMethodVoucherSuccess
	}

	responses := s.mapping.ToVouchersResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly voucher success methods",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindMonthMethodVoucherFailed(year int) ([]*response.MonthMethodVoucherResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly voucher failed methods", zap.Int("year", year))

	records, err := s.voucherRepository.FindMonthMethodVoucherFailed(year)
	if err != nil {
		s.logger.Error("Failed to fetch monthly voucher failed methods", zap.Error(err))
		return nil, voucher_errors.ErrFailedFindMonthMethodVoucherFailed
	}

	responses := s.mapping.ToVouchersResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly voucher failed methods",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindYearMethodVoucherFailed(year int) ([]*response.YearMethodVoucherResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly voucher failed methods", zap.Int("year", year))

	records, err := s.voucherRepository.FindYearMethodVoucherFailed(year)
	if err != nil {
		s.logger.Error("Failed to fetch yearly voucher failed methods", zap.Error(err))
		return nil, voucher_errors.ErrFailedFindYearMethodVoucherFailed
	}

	responses := s.mapping.ToVouchersResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly voucher failed methods",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindMonthAmountVoucherSuccessById(req *requests.MonthAmountVoucherByIdRequest) ([]*response.MonthAmountVoucherSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly voucher success amounts by ID",
		zap.Any("request", req),
		zap.Int("voucher_id", req.ID))

	records, err := s.voucherRepository.FindMonthAmountVoucherSuccessById(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly voucher success amounts by ID",
			zap.Error(err),
			zap.Int("voucher_id", req.ID))
		return nil, voucher_errors.ErrFailedFindMonthAmountVoucherSuccessById
	}

	responses := s.mapping.ToVouchersResponseMonthAmountSuccess(records)

	s.logger.Debug("Successfully fetched monthly voucher success amounts by ID",
		zap.Int("voucher_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindYearAmountVoucherSuccessById(req *requests.YearAmountVoucherByIdRequest) ([]*response.YearAmountVoucherSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly voucher success amounts by ID",
		zap.Any("request", req),
		zap.Int("voucher_id", req.ID))

	records, err := s.voucherRepository.FindYearAmountVoucherSuccessById(req)
	if err != nil {
		s.logger.Error("Failed to fetch yearly voucher success amounts by ID",
			zap.Error(err),
			zap.Int("voucher_id", req.ID))
		return nil, voucher_errors.ErrFailedFindYearAmountVoucherSuccessById
	}

	responses := s.mapping.ToVouchersResponseYearAmountSuccess(records)

	s.logger.Debug("Successfully fetched yearly voucher success amounts by ID",
		zap.Int("voucher_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindMonthAmountVoucherFailedById(req *requests.MonthAmountVoucherByIdRequest) ([]*response.MonthAmountVoucherFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly voucher failed amounts by ID",
		zap.Any("request", req),
		zap.Int("voucher_id", req.ID))

	records, err := s.voucherRepository.FindMonthAmountVoucherFailedById(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly voucher failed amounts by ID",
			zap.Error(err),
			zap.Int("voucher_id", req.ID))
		return nil, voucher_errors.ErrFailedFindMonthAmountVoucherFailedById
	}

	responses := s.mapping.ToVouchersResponseMonthAmountFailed(records)

	s.logger.Debug("Successfully fetched monthly voucher failed amounts by ID",
		zap.Int("voucher_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindYearAmountVoucherFailedById(req *requests.YearAmountVoucherByIdRequest) ([]*response.YearAmountVoucherFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly voucher failed amounts by ID",
		zap.Any("request", req),
		zap.Int("voucher_id", req.ID))

	records, err := s.voucherRepository.FindYearAmountVoucherFailedById(req)
	if err != nil {
		s.logger.Error("Failed to fetch yearly voucher failed amounts by ID",
			zap.Error(err),
			zap.Int("voucher_id", req.ID))
		return nil, voucher_errors.ErrFailedFindYearAmountVoucherFailedById
	}

	responses := s.mapping.ToVouchersResponseYearAmountFailed(records)

	s.logger.Debug("Successfully fetched yearly voucher failed amounts by ID",
		zap.Int("voucher_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindMonthMethodVoucherSuccessById(req *requests.MonthMethodVoucherByIdRequest) ([]*response.MonthMethodVoucherResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly voucher success methods by ID",
		zap.Any("request", req),
		zap.Int("voucher_id", req.ID))

	records, err := s.voucherRepository.FindMonthMethodVoucherSuccessById(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly voucher success methods by ID",
			zap.Error(err),
			zap.Int("voucher_id", req.ID))
		return nil, voucher_errors.ErrFailedFindMonthMethodVoucherSuccessById
	}

	responses := s.mapping.ToVouchersResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly voucher success methods by ID",
		zap.Int("voucher_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindYearMethodVoucherSuccessById(req *requests.YearMethodVoucherByIdRequest) ([]*response.YearMethodVoucherResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly voucher success methods by ID",
		zap.Any("request", req),
		zap.Int("voucher_id", req.ID))

	records, err := s.voucherRepository.FindYearMethodVoucherSuccessById(req)
	if err != nil {
		s.logger.Error("Failed to fetch yearly voucher success methods by ID",
			zap.Error(err),
			zap.Int("voucher_id", req.ID))
		return nil, voucher_errors.ErrFailedFindYearMethodVoucherSuccessById
	}

	responses := s.mapping.ToVouchersResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly voucher success methods by ID",
		zap.Int("voucher_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindMonthMethodVoucherFailedById(req *requests.MonthMethodVoucherByIdRequest) ([]*response.MonthMethodVoucherResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly voucher failed methods by ID",
		zap.Any("request", req),
		zap.Int("voucher_id", req.ID))

	records, err := s.voucherRepository.FindMonthMethodVoucherFailedById(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly voucher failed methods by ID",
			zap.Error(err),
			zap.Int("voucher_id", req.ID))
		return nil, voucher_errors.ErrFailedFindMonthMethodVoucherFailedById
	}

	responses := s.mapping.ToVouchersResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly voucher failed methods by ID",
		zap.Int("voucher_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindYearMethodVoucherFailedById(req *requests.YearMethodVoucherByIdRequest) ([]*response.YearMethodVoucherResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly voucher failed methods by ID",
		zap.Any("request", req),
		zap.Int("voucher_id", req.ID))

	records, err := s.voucherRepository.FindYearMethodVoucherFailedById(req)
	if err != nil {
		s.logger.Error("Failed to fetch yearly voucher failed methods by ID",
			zap.Error(err),
			zap.Int("voucher_id", req.ID))
		return nil, voucher_errors.ErrFailedFindYearMethodVoucherFailedById
	}

	responses := s.mapping.ToVouchersResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly voucher failed methods by ID",
		zap.Int("voucher_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindMonthAmountVoucherSuccessByMerchant(req *requests.MonthAmountVoucherByMerchantRequest) ([]*response.MonthAmountVoucherSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly voucher success amounts by ID",
		zap.Any("request", req),
		zap.Int("merchant_id", req.MerchantID))

	records, err := s.voucherRepository.FindMonthAmountVoucherSuccessByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly voucher success amounts by ID",
			zap.Error(err),
			zap.Int("merchant_id", req.MerchantID))
		return nil, voucher_errors.ErrFailedFindMonthAmountVoucherSuccessByMerchant
	}

	responses := s.mapping.ToVouchersResponseMonthAmountSuccess(records)

	s.logger.Debug("Successfully fetched monthly voucher success amounts by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindYearAmountVoucherSuccessByMerchant(req *requests.YearAmountVoucherByMerchantRequest) ([]*response.YearAmountVoucherSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly voucher success amounts by ID",
		zap.Any("request", req),
		zap.Int("merchant_id", req.MerchantID))

	records, err := s.voucherRepository.FindYearAmountVoucherSuccessByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch yearly voucher success amounts by ID",
			zap.Error(err),
			zap.Int("merchant_id", req.MerchantID))
		return nil, voucher_errors.ErrFailedFindYearAmountVoucherSuccessByMerchant
	}

	responses := s.mapping.ToVouchersResponseYearAmountSuccess(records)

	s.logger.Debug("Successfully fetched yearly voucher success amounts by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindMonthAmountVoucherFailedByMerchant(req *requests.MonthAmountVoucherByMerchantRequest) ([]*response.MonthAmountVoucherFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly voucher failed amounts by ID",
		zap.Any("request", req),
		zap.Int("merchant_id", req.MerchantID))

	records, err := s.voucherRepository.FindMonthAmountVoucherFailedByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly voucher failed amounts by ID",
			zap.Error(err),
			zap.Int("merchant_id", req.MerchantID))
		return nil, voucher_errors.ErrFailedFindMonthAmountVoucherFailedByMerchant
	}

	responses := s.mapping.ToVouchersResponseMonthAmountFailed(records)

	s.logger.Debug("Successfully fetched monthly voucher failed amounts by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindYearAmountVoucherFailedByMerchant(req *requests.YearAmountVoucherByMerchantRequest) ([]*response.YearAmountVoucherFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly voucher failed amounts by ID",
		zap.Any("request", req),
		zap.Int("merchant_id", req.MerchantID))

	records, err := s.voucherRepository.FindYearAmountVoucherFailedByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch yearly voucher failed amounts by ID",
			zap.Error(err),
			zap.Int("merchant_id", req.MerchantID))
		return nil, voucher_errors.ErrFailedFindYearAmountVoucherFailedByMerchant
	}

	responses := s.mapping.ToVouchersResponseYearAmountFailed(records)

	s.logger.Debug("Successfully fetched yearly voucher failed amounts by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindMonthMethodVoucherSuccessByMerchant(req *requests.MonthMethodVoucherByMerchantRequest) ([]*response.MonthMethodVoucherResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly voucher success methods by ID",
		zap.Any("request", req),
		zap.Int("merchant_id", req.MerchantID))

	records, err := s.voucherRepository.FindMonthMethodVoucherSuccessByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly voucher success methods by ID",
			zap.Error(err),
			zap.Int("merchant_id", req.MerchantID))
		return nil, voucher_errors.ErrFailedFindMonthMethodVoucherSuccessByMerchant
	}

	responses := s.mapping.ToVouchersResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly voucher success methods by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindYearMethodVoucherSuccessByMerchant(req *requests.YearMethodVoucherByMerchantRequest) ([]*response.YearMethodVoucherResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly voucher success methods by ID",
		zap.Any("request", req),
		zap.Int("merchant_id", req.MerchantID))

	records, err := s.voucherRepository.FindYearMethodVoucherSuccessByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch yearly voucher success methods by ID",
			zap.Error(err),
			zap.Int("merchant_id", req.MerchantID))
		return nil, voucher_errors.ErrFailedFindYearMethodVoucherSuccessByMerchant
	}

	responses := s.mapping.ToVouchersResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly voucher success methods by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindMonthMethodVoucherFailedByMerchant(req *requests.MonthMethodVoucherByMerchantRequest) ([]*response.MonthMethodVoucherResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly voucher failed methods by ID",
		zap.Any("request", req),
		zap.Int("merchant_id", req.MerchantID))

	records, err := s.voucherRepository.FindMonthMethodVoucherFailedByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly voucher failed methods by ID",
			zap.Error(err),
			zap.Int("merchant_id", req.MerchantID))
		return nil, voucher_errors.ErrFailedFindMonthMethodVoucherFailedByMerchant
	}

	responses := s.mapping.ToVouchersResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly voucher failed methods by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindYearMethodVoucherFailedByMerchant(req *requests.YearMethodVoucherByMerchantRequest) ([]*response.YearMethodVoucherResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly voucher failed methods by ID",
		zap.Any("request", req),
		zap.Int("merchant_id", req.MerchantID))

	records, err := s.voucherRepository.FindYearMethodVoucherFailedByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch yearly voucher failed methods by ID",
			zap.Error(err),
			zap.Int("merchant_id", req.MerchantID))
		return nil, voucher_errors.ErrFailedFindYearMethodVoucherFailedByMerchant
	}

	responses := s.mapping.ToVouchersResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly voucher failed methods by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *voucherService) FindById(id int) (*response.VoucherResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching voucher by ID", zap.Int("id", id))

	res, err := s.voucherRepository.FindById(id)

	if err != nil {
		s.logger.Error("Failed to fetch voucher record by ID", zap.Error(err))

		return nil, voucher_errors.ErrVoucherNotFoundRes
	}

	s.logger.Debug("Successfully fetched voucher", zap.Int("id", id))

	so := s.mapping.ToVoucherResponse(res)

	return so, nil
}

func (s *voucherService) FindByActive(request *requests.FindAllVouchers) ([]*response.VoucherResponseDeleteAt, *int, *response.ErrorResponse) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

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

	res, totalRecords, err := s.voucherRepository.FindByActiveVouchers(request)
	if err != nil {
		s.logger.Error("Failed to fetch active voucher",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, nil, voucher_errors.ErrFailedFindActive
	}

	s.logger.Debug("Successfully fetched active voucher",
		zap.Int("totalRecords", *totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	so := s.mapping.ToVouchersResponseDeleteAt(res)

	return so, totalRecords, nil
}

func (s *voucherService) FindByTrashed(request *requests.FindAllVouchers) ([]*response.VoucherResponseDeleteAt, *int, *response.ErrorResponse) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

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

	res, totalRecords, err := s.voucherRepository.FindByTrashedVoucher(request)

	if err != nil {
		s.logger.Error("Failed to fetch trashed voucher",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, nil, voucher_errors.ErrFailedFindTrashed
	}

	s.logger.Debug("Successfully fetched trashed voucher",
		zap.Int("totalRecords", *totalRecords),
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

		return nil, voucher_errors.ErrFailedCreateVoucher
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

		return nil, voucher_errors.ErrFailedUpdateVoucher
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

		return nil, voucher_errors.ErrFailedTrashedVoucher
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

		return nil, voucher_errors.ErrFailedRestoreVoucher
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

		return false, voucher_errors.ErrFailedDeletePermanent
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
		return false, voucher_errors.ErrFailedRestoreAll
	}

	s.logger.Debug("Successfully restored all Voucher")
	return true, nil
}

func (s *voucherService) DeleteAllPermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all Voucher")

	_, err := s.voucherRepository.DeleteAllVouchersPermanent()

	if err != nil {
		s.logger.Error("Failed to permanently delete all Voucher", zap.Error(err))
		return false, voucher_errors.ErrFailedDeleteAll
	}

	s.logger.Debug("Successfully deleted all Voucher permanently")
	return true, nil
}
