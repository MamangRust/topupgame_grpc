package service

import (
	"topup_game/internal/domain/requests"
	"topup_game/internal/domain/response"
	response_service "topup_game/internal/mapper/response/service"
	"topup_game/internal/repository"
	"topup_game/pkg/errors/bank_errors"
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

func (s *bankService) FindAll(request *requests.FindAllBanks) ([]*response.BankResponse, *int, *response.ErrorResponse) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

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

	banks, totalRecords, err := s.bankRepository.FindAllBanks(request)

	if err != nil {
		s.logger.Error("Failed to fetch bank",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, nil, bank_errors.ErrFailedFailedFindAll
	}

	bankResponses := s.mapping.ToBanksResponse(banks)

	s.logger.Debug("Successfully fetched banks",
		zap.Int("totalRecords", *totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return bankResponses, totalRecords, nil
}

func (s *bankService) FindMonthAmountBankSuccess(req *requests.MonthAmountBankRequest) ([]*response.MonthAmountBankSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly bank success amounts", zap.Any("request", req))

	records, err := s.bankRepository.FindMonthAmountBankSuccess(req)
	if err != nil {
		s.logger.Error("failed to find monthly bank success amounts", zap.Error(err))
		return nil, bank_errors.ErrFailedFindMonthAmountBankSuccess
	}

	responses := s.mapping.ToBanksResponseMonthAmountSuccess(records)

	s.logger.Debug("Successfully fetched monthly bank success amounts", zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindYearAmountBankSuccess(year int) ([]*response.YearAmountBankSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly bank success amounts", zap.Int("year", year))

	records, err := s.bankRepository.FindYearAmountBankSuccess(year)
	if err != nil {
		s.logger.Error("failed to find yearly bank success amounts", zap.Error(err))
		return nil, bank_errors.ErrFailedFindYearAmountBankSuccess
	}

	responses := s.mapping.ToBanksResponseYearAmountSuccess(records)

	s.logger.Debug("Successfully fetched yearly bank success amounts", zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindMonthAmountBankFailed(req *requests.MonthAmountBankRequest) ([]*response.MonthAmountBankFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly bank failed amounts", zap.Any("request", req))

	records, err := s.bankRepository.FindMonthAmountBankFailed(req)
	if err != nil {
		s.logger.Error("failed to find monthly bank failed amounts", zap.Error(err))
		return nil, bank_errors.ErrFailedFindMonthAmountBankFailed
	}

	responses := s.mapping.ToBanksResponseMonthAmountFailed(records)

	s.logger.Debug("Successfully fetched monthly bank failed amounts", zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindYearAmountBankFailed(year int) ([]*response.YearAmountBankFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly bank failed amounts", zap.Int("year", year))

	records, err := s.bankRepository.FindYearAmountBankFailed(year)
	if err != nil {
		s.logger.Error("failed to find yearly bank failed amounts", zap.Error(err))
		return nil, bank_errors.ErrFailedFindYearAmountBankFailed
	}

	responses := s.mapping.ToBanksResponseYearAmountFailed(records)

	s.logger.Debug("Successfully fetched yearly bank failed amounts", zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindMonthMethodBankSuccess(year int) ([]*response.MonthMethodBankResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly bank success methods", zap.Int("year", year))

	records, err := s.bankRepository.FindMonthMethodBankSuccess(year)
	if err != nil {
		s.logger.Error("failed to find monthly bank success methods", zap.Error(err))
		return nil, bank_errors.ErrFailedFindMonthMethodBankSuccess
	}

	responses := s.mapping.ToBanksResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly bank success methods", zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindYearMethodBankSuccess(year int) ([]*response.YearMethodBankResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly bank success methods", zap.Int("year", year))

	records, err := s.bankRepository.FindYearMethodBankSuccess(year)
	if err != nil {
		s.logger.Error("failed to find yearly bank success methods", zap.Error(err))
		return nil, bank_errors.ErrFailedFindYearMethodBankSuccess
	}

	responses := s.mapping.ToBanksResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly bank success methods", zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindMonthMethodBankFailed(year int) ([]*response.MonthMethodBankResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly bank failed methods", zap.Int("year", year))

	records, err := s.bankRepository.FindMonthMethodBankFailed(year)
	if err != nil {
		s.logger.Error("failed to find monthly bank failed methods", zap.Error(err))
		return nil, bank_errors.ErrFailedFindMonthMethodBankFailed
	}

	responses := s.mapping.ToBanksResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly bank failed methods", zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindYearMethodBankFailed(year int) ([]*response.YearMethodBankResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly bank failed methods", zap.Int("year", year))

	records, err := s.bankRepository.FindYearMethodBankFailed(year)
	if err != nil {
		s.logger.Error("failed to find yearly bank failed methods", zap.Error(err))
		return nil, bank_errors.ErrFailedFindYearMethodBankFailed
	}

	responses := s.mapping.ToBanksResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly bank failed methods", zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindMonthAmountBankSuccessById(req *requests.MonthAmountBankByIdRequest) ([]*response.MonthAmountBankSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly bank success amounts by ID", zap.Any("request", req))

	records, err := s.bankRepository.FindMonthAmountBankSuccessById(req)
	if err != nil {
		s.logger.Error("failed to find monthly bank success amounts by ID", zap.Error(err))
		return nil, bank_errors.ErrFailedFindMonthAmountBankSuccessById
	}

	responses := s.mapping.ToBanksResponseMonthAmountSuccess(records)

	s.logger.Debug("Successfully fetched monthly bank success amounts by ID",
		zap.Int("bank_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindYearAmountBankSuccessById(req *requests.YearAmountBankByIdRequest) ([]*response.YearAmountBankSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly bank success amounts by ID", zap.Any("request", req))

	records, err := s.bankRepository.FindYearAmountBankSuccessById(req)
	if err != nil {
		s.logger.Error("failed to find yearly bank success amounts by ID", zap.Error(err))
		return nil, bank_errors.ErrFailedFindYearAmountBankSuccessById
	}

	responses := s.mapping.ToBanksResponseYearAmountSuccess(records)

	s.logger.Debug("Successfully fetched yearly bank success amounts by ID",
		zap.Int("bank_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindMonthAmountBankFailedById(req *requests.MonthAmountBankByIdRequest) ([]*response.MonthAmountBankFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly bank failed amounts by ID", zap.Any("request", req))

	records, err := s.bankRepository.FindMonthAmountBankFailedById(req)
	if err != nil {
		s.logger.Error("failed to find monthly bank failed amounts by ID", zap.Error(err))
		return nil, bank_errors.ErrFailedFindMonthAmountBankFailedById
	}

	responses := s.mapping.ToBanksResponseMonthAmountFailed(records)

	s.logger.Debug("Successfully fetched monthly bank failed amounts by ID",
		zap.Int("bank_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindYearAmountBankFailedById(req *requests.YearAmountBankByIdRequest) ([]*response.YearAmountBankFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly bank failed amounts by ID", zap.Any("request", req))

	records, err := s.bankRepository.FindYearAmountBankFailedById(req)
	if err != nil {
		s.logger.Error("failed to find yearly bank failed amounts by ID", zap.Error(err))
		return nil, bank_errors.ErrFailedFindYearAmountBankFailedById
	}

	responses := s.mapping.ToBanksResponseYearAmountFailed(records)

	s.logger.Debug("Successfully fetched yearly bank failed amounts by ID",
		zap.Int("bank_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindMonthMethodBankSuccessById(req *requests.MonthMethodBankByIdRequest) ([]*response.MonthMethodBankResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly bank success methods by ID", zap.Any("request", req))

	records, err := s.bankRepository.FindMonthMethodBankSuccessById(req)
	if err != nil {
		s.logger.Error("failed to find monthly bank success methods by ID", zap.Error(err))
		return nil, bank_errors.ErrFailedFindMonthMethodBankSuccessById
	}

	responses := s.mapping.ToBanksResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly bank success methods by ID",
		zap.Int("bank_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindYearMethodBankSuccessById(req *requests.YearMethodBankByIdRequest) ([]*response.YearMethodBankResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly bank success methods by ID", zap.Any("request", req))

	records, err := s.bankRepository.FindYearMethodBankSuccessById(req)
	if err != nil {
		s.logger.Error("failed to find yearly bank success methods by ID", zap.Error(err))
		return nil, bank_errors.ErrFailedFindYearMethodBankSuccessById
	}

	responses := s.mapping.ToBanksResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly bank success methods by ID",
		zap.Int("bank_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindMonthMethodBankFailedById(req *requests.MonthMethodBankByIdRequest) ([]*response.MonthMethodBankResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly bank failed methods by ID", zap.Any("request", req))

	records, err := s.bankRepository.FindMonthMethodBankFailedById(req)
	if err != nil {
		s.logger.Error("failed to find monthly bank failed methods by ID", zap.Error(err))
		return nil, bank_errors.ErrFailedFindMonthMethodBankFailedById
	}

	responses := s.mapping.ToBanksResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly bank failed methods by ID",
		zap.Int("bank_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindYearMethodBankFailedById(req *requests.YearMethodBankByIdRequest) ([]*response.YearMethodBankResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly bank failed methods by ID", zap.Any("request", req))

	records, err := s.bankRepository.FindYearMethodBankFailedById(req)
	if err != nil {
		s.logger.Error("failed to find yearly bank failed methods by ID", zap.Error(err))
		return nil, bank_errors.ErrFailedFindYearMethodBankFailedById
	}

	responses := s.mapping.ToBanksResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly bank failed methods by ID",
		zap.Int("bank_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindMonthAmountBankSuccessByMerchant(req *requests.MonthAmountBankByMerchantRequest) ([]*response.MonthAmountBankSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly bank success amounts by ID", zap.Any("request", req))

	records, err := s.bankRepository.FindMonthAmountBankSuccessByMerchant(req)
	if err != nil {
		s.logger.Error("failed to find monthly bank success amounts by ID", zap.Error(err))
		return nil, bank_errors.ErrFailedFindMonthAmountBankSuccessByMerchant
	}

	responses := s.mapping.ToBanksResponseMonthAmountSuccess(records)

	s.logger.Debug("Successfully fetched monthly bank success amounts by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindYearAmountBankSuccessByMerchant(req *requests.YearAmountBankByMerchantRequest) ([]*response.YearAmountBankSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly bank success amounts by ID", zap.Any("request", req))

	records, err := s.bankRepository.FindYearAmountBankSuccessByMerchant(req)
	if err != nil {
		s.logger.Error("failed to find yearly bank success amounts by ID", zap.Error(err))
		return nil, bank_errors.ErrFailedFindYearAmountBankSuccessByMerchant
	}

	responses := s.mapping.ToBanksResponseYearAmountSuccess(records)

	s.logger.Debug("Successfully fetched yearly bank success amounts by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindMonthAmountBankFailedByMerchant(req *requests.MonthAmountBankByMerchantRequest) ([]*response.MonthAmountBankFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly bank failed amounts by ID", zap.Any("request", req))

	records, err := s.bankRepository.FindMonthAmountBankFailedByMerchant(req)
	if err != nil {
		s.logger.Error("failed to find monthly bank failed amounts by ID", zap.Error(err))
		return nil, bank_errors.ErrFailedFindMonthAmountBankFailedByMerchant
	}

	responses := s.mapping.ToBanksResponseMonthAmountFailed(records)

	s.logger.Debug("Successfully fetched monthly bank failed amounts by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindYearAmountBankFailedByMerchant(req *requests.YearAmountBankByMerchantRequest) ([]*response.YearAmountBankFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly bank failed amounts by ID", zap.Any("request", req))

	records, err := s.bankRepository.FindYearAmountBankFailedByMerchant(req)
	if err != nil {
		s.logger.Error("failed to find yearly bank failed amounts by ID", zap.Error(err))
		return nil, bank_errors.ErrFailedFindYearAmountBankFailedByMerchant
	}

	responses := s.mapping.ToBanksResponseYearAmountFailed(records)

	s.logger.Debug("Successfully fetched yearly bank failed amounts by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindMonthMethodBankSuccessByMerchant(req *requests.MonthMethodBankByMerchantRequest) ([]*response.MonthMethodBankResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly bank success methods by ID", zap.Any("request", req))

	records, err := s.bankRepository.FindMonthMethodBankSuccessByMerchant(req)
	if err != nil {
		s.logger.Error("failed to find monthly bank success methods by ID", zap.Error(err))
		return nil, bank_errors.ErrFailedFindMonthMethodBankSuccessByMerchant
	}

	responses := s.mapping.ToBanksResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly bank success methods by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindYearMethodBankSuccessByMerchant(req *requests.YearMethodBankByMerchantRequest) ([]*response.YearMethodBankResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly bank success methods by ID", zap.Any("request", req))

	records, err := s.bankRepository.FindYearMethodBankSuccessByMerchant(req)
	if err != nil {
		s.logger.Error("failed to find yearly bank success methods by ID", zap.Error(err))
		return nil, bank_errors.ErrFailedFindYearMethodBankSuccessByMerchant
	}

	responses := s.mapping.ToBanksResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly bank success methods by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindMonthMethodBankFailedByMerchant(req *requests.MonthMethodBankByMerchantRequest) ([]*response.MonthMethodBankResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly bank failed methods by ID", zap.Any("request", req))

	records, err := s.bankRepository.FindMonthMethodBankFailedByMerchant(req)
	if err != nil {
		s.logger.Error("failed to find monthly bank failed methods by ID", zap.Error(err))
		return nil, bank_errors.ErrFailedFindMonthMethodBankFailedByMerchant
	}

	responses := s.mapping.ToBanksResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly bank failed methods by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindYearMethodBankFailedByMerchant(req *requests.YearMethodBankByMerchantRequest) ([]*response.YearMethodBankResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly bank failed methods by ID", zap.Any("request", req))

	records, err := s.bankRepository.FindYearMethodBankFailedByMerchant(req)
	if err != nil {
		s.logger.Error("failed to find yearly bank failed methods by ID", zap.Error(err))
		return nil, bank_errors.ErrFailedFindYearMethodBankFailedByMerchant
	}

	responses := s.mapping.ToBanksResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly bank failed methods by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *bankService) FindByID(id int) (*response.BankResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching banks by id", zap.Int("user_id", id))

	user, err := s.bankRepository.FindById(id)

	if err != nil {
		s.logger.Error("failed to find bank by ID", zap.Error(err))
		return nil, bank_errors.ErrBankNotFoundRes
	}

	so := s.mapping.ToBankResponse(user)

	s.logger.Debug("Successfully fetched bank", zap.Int("bank_id", id))

	return so, nil
}

func (s *bankService) FindByActive(request *requests.FindAllBanks) ([]*response.BankResponseDeleteAt, *int, *response.ErrorResponse) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

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

	users, totalRecords, err := s.bankRepository.FindByActiveBanks(request)

	if err != nil {
		s.logger.Error("Failed to fetch active banks",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, nil, bank_errors.ErrFailedFailedFindActive
	}

	so := s.mapping.ToBanksResponseDeleteAt(users)

	s.logger.Debug("Successfully fetched active banks",
		zap.Int("totalRecords", *totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return so, totalRecords, nil
}

func (s *bankService) FindByTrashed(request *requests.FindAllBanks) ([]*response.BankResponseDeleteAt, *int, *response.ErrorResponse) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

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

	banks, totalRecords, err := s.bankRepository.FindByTrashedBanks(request)

	if err != nil {
		s.logger.Error("Failed to find trashed banks", zap.Error(err))

		return nil, nil, bank_errors.ErrFailedFailedFindTrashed
	}

	so := s.mapping.ToBanksResponseDeleteAt(banks)

	s.logger.Debug("Successfully fetched trashed banks",
		zap.Int("totalRecords", *totalRecords),
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

		return nil, bank_errors.ErrFailedCreateBank
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

		return nil, bank_errors.ErrFailedUpdateBank
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

		return nil, bank_errors.ErrFailedTrashedBank
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

		return nil, bank_errors.ErrFailedRestoreBank
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

		return false, bank_errors.ErrFailedDeletePermanent
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
		return false, bank_errors.ErrFailedRestoreAll
	}

	s.logger.Debug("Successfully restored all banks")
	return true, nil
}

func (s *bankService) DeleteAllPermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all banks")

	_, err := s.bankRepository.DeleteAllBanksPermanent()

	if err != nil {
		s.logger.Error("Failed to permanently delete all banks", zap.Error(err))
		return false, bank_errors.ErrFailedDeleteAll
	}

	s.logger.Debug("Successfully deleted all banks permanently")
	return true, nil
}
