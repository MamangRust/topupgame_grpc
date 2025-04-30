package service

import (
	"topup_game/internal/domain/requests"
	"topup_game/internal/domain/response"
	response_service "topup_game/internal/mapper/response/service"
	"topup_game/internal/repository"
	"topup_game/pkg/errors/nominal_errors"
	"topup_game/pkg/errors/voucher_errors"
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

func (s *nominalService) FindAll(request *requests.FindAllNominals) ([]*response.NominalResponse, *int, *response.ErrorResponse) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

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

	banks, totalRecords, err := s.nominalRepository.FindAllNominals(request)

	if err != nil {
		s.logger.Error("Failed to fetch nominals",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, nil, nominal_errors.ErrFailedFindAll
	}

	nominalResponses := s.mapping.ToNominalsResponse(banks)

	s.logger.Debug("Successfully fetched banks",
		zap.Int("totalRecords", *totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return nominalResponses, totalRecords, nil
}

func (s *nominalService) FindMonthAmountNominalSuccess(req *requests.MonthAmountNominalRequest) ([]*response.MonthAmountNominalSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly nominal success amounts", zap.Any("request", req))

	records, err := s.nominalRepository.FindMonthAmountNominalSuccess(req)
	if err != nil {
		s.logger.Error("failed to find monthly nominal success amounts", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindMonthAmountNominalSuccess
	}

	responses := s.mapping.ToNominalsResponseMonthAmountSuccess(records)

	s.logger.Debug("Successfully fetched monthly nominal success amounts",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindYearAmountNominalSuccess(year int) ([]*response.YearAmountNominalSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly nominal success amounts", zap.Int("year", year))

	records, err := s.nominalRepository.FindYearAmountNominalSuccess(year)
	if err != nil {
		s.logger.Error("failed to find yearly nominal success amounts", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindYearAmountNominalSuccess
	}

	responses := s.mapping.ToNominalsResponseYearAmountSuccess(records)

	s.logger.Debug("Successfully fetched yearly nominal success amounts",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindMonthAmountNominalFailed(req *requests.MonthAmountNominalRequest) ([]*response.MonthAmountNominalFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly nominal failed amounts", zap.Any("request", req))

	records, err := s.nominalRepository.FindMonthAmountNominalFailed(req)
	if err != nil {
		s.logger.Error("failed to find monthly nominal failed amounts", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindMonthAmountNominalFailed
	}

	responses := s.mapping.ToNominalsResponseMonthAmountFailed(records)

	s.logger.Debug("Successfully fetched monthly nominal failed amounts",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindYearAmountNominalFailed(year int) ([]*response.YearAmountNominalFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly nominal failed amounts", zap.Int("year", year))

	records, err := s.nominalRepository.FindYearAmountNominalFailed(year)
	if err != nil {
		s.logger.Error("failed to find yearly nominal failed amounts", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindYearAmountNominalFailed
	}

	responses := s.mapping.ToNominalsResponseYearAmountFailed(records)

	s.logger.Debug("Successfully fetched yearly nominal failed amounts",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindMonthMethodNominalSuccess(year int) ([]*response.MonthMethodNominalResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly nominal success methods", zap.Int("year", year))

	records, err := s.nominalRepository.FindMonthMethodNominalSuccess(year)
	if err != nil {
		s.logger.Error("failed to find monthly nominal success methods", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindMonthMethodNominalSuccess
	}

	responses := s.mapping.ToNominalsResponseMonthMethodSuccess(records)

	s.logger.Debug("Successfully fetched monthly nominal success methods",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindYearMethodNominalSuccess(year int) ([]*response.YearMethodNominalResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly nominal success methods", zap.Int("year", year))

	records, err := s.nominalRepository.FindYearMethodNominalSuccess(year)
	if err != nil {
		s.logger.Error("failed to find yearly nominal success methods", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindYearMethodNominalSuccess
	}

	responses := s.mapping.ToNominalsResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly nominal success methods",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindMonthMethodNominalFailed(year int) ([]*response.MonthMethodNominalResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly nominal failed methods", zap.Int("year", year))

	records, err := s.nominalRepository.FindMonthMethodNominalFailed(year)
	if err != nil {
		s.logger.Error("failed to find monthly nominal failed methods", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindMonthMethodNominalFailed
	}

	responses := s.mapping.ToNominalsResponseMonthMethodFailed(records)

	s.logger.Debug("Successfully fetched monthly nominal failed methods",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindYearMethodNominalFailed(year int) ([]*response.YearMethodNominalResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly nominal failed methods", zap.Int("year", year))

	records, err := s.nominalRepository.FindYearMethodNominalFailed(year)
	if err != nil {
		s.logger.Error("failed to find yearly nominal failed methods", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindYearMethodNominalFailed
	}

	responses := s.mapping.ToNominalsResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly nominal failed methods",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindMonthAmountNominalSuccessById(req *requests.MonthAmountNominalByIdRequest) ([]*response.MonthAmountNominalSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly nominal success amounts by ID", zap.Any("request", req))

	records, err := s.nominalRepository.FindMonthAmountNominalSuccessById(req)
	if err != nil {
		s.logger.Error("failed to find monthly nominal success amounts by ID", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindMonthAmountNominalSuccessById
	}

	responses := s.mapping.ToNominalsResponseMonthAmountSuccess(records)

	s.logger.Debug("Successfully fetched monthly nominal success amounts by ID",
		zap.Int("nominal_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindYearAmountNominalSuccessById(req *requests.YearAmountNominalByIdRequest) ([]*response.YearAmountNominalSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly nominal success amounts by ID", zap.Any("request", req))

	records, err := s.nominalRepository.FindYearAmountNominalSuccessById(req)
	if err != nil {
		s.logger.Error("failed to find yearly nominal success amounts by ID", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindYearAmountNominalSuccessById
	}

	responses := s.mapping.ToNominalsResponseYearAmountSuccess(records)

	s.logger.Debug("Successfully fetched yearly nominal success amounts by ID",
		zap.Int("nominal_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindMonthAmountNominalFailedById(req *requests.MonthAmountNominalByIdRequest) ([]*response.MonthAmountNominalFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly nominal failed amounts by ID", zap.Any("request", req))

	records, err := s.nominalRepository.FindMonthAmountNominalFailedById(req)
	if err != nil {
		s.logger.Error("failed to find monthly nominal failed amounts by ID", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindMonthAmountNominalFailedById
	}

	responses := s.mapping.ToNominalsResponseMonthAmountFailed(records)

	s.logger.Debug("Successfully fetched monthly nominal failed amounts by ID",
		zap.Int("nominal_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindYearAmountNominalFailedById(req *requests.YearAmountNominalByIdRequest) ([]*response.YearAmountNominalFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly nominal failed amounts by ID", zap.Any("request", req))

	records, err := s.nominalRepository.FindYearAmountNominalFailedById(req)
	if err != nil {
		s.logger.Error("failed to find yearly nominal failed amounts by ID", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindYearAmountNominalFailedById
	}

	responses := s.mapping.ToNominalsResponseYearAmountFailed(records)

	s.logger.Debug("Successfully fetched yearly nominal failed amounts by ID",
		zap.Int("nominal_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindMonthMethodNominalSuccessById(req *requests.MonthMethodNominalByIdRequest) ([]*response.MonthMethodNominalResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly nominal success methods by ID", zap.Any("request", req))

	records, err := s.nominalRepository.FindMonthMethodNominalSuccessById(req)
	if err != nil {
		s.logger.Error("failed to find monthly nominal success methods by ID", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindMonthMethodNominalSuccessById
	}

	responses := s.mapping.ToNominalsResponseMonthMethodSuccess(records)

	s.logger.Debug("Successfully fetched monthly nominal success methods by ID",
		zap.Int("nominal_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindYearMethodNominalSuccessById(req *requests.YearMethodNominalByIdRequest) ([]*response.YearMethodNominalResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly nominal success methods by ID", zap.Any("request", req))

	records, err := s.nominalRepository.FindYearMethodNominalSuccessById(req)
	if err != nil {
		s.logger.Error("failed to find yearly nominal success methods by ID", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindYearMethodNominalSuccessById
	}

	responses := s.mapping.ToNominalsResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly nominal success methods by ID",
		zap.Int("nominal_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindMonthMethodNominalFailedById(req *requests.MonthMethodNominalByIdRequest) ([]*response.MonthMethodNominalResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly nominal failed methods by ID", zap.Any("request", req))

	records, err := s.nominalRepository.FindMonthMethodNominalFailedById(req)
	if err != nil {
		s.logger.Error("failed to find monthly nominal failed methods by ID", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindMonthMethodNominalFailedById
	}

	responses := s.mapping.ToNominalsResponseMonthMethodFailed(records)

	s.logger.Debug("Successfully fetched monthly nominal failed methods by ID",
		zap.Int("nominal_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindYearMethodNominalFailedById(req *requests.YearMethodNominalByIdRequest) ([]*response.YearMethodNominalResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly nominal failed methods by ID", zap.Any("request", req))

	records, err := s.nominalRepository.FindYearMethodNominalFailedById(req)
	if err != nil {
		s.logger.Error("failed to find yearly nominal failed methods by ID", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindYearMethodNominalFailedById
	}

	responses := s.mapping.ToNominalsResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly nominal failed methods by ID",
		zap.Int("nominal_id", req.ID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindMonthAmountNominalSuccessByMerchant(req *requests.MonthAmountNominalByMerchantRequest) ([]*response.MonthAmountNominalSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly nominal success amounts by ID", zap.Any("request", req))

	records, err := s.nominalRepository.FindMonthAmountNominalSuccessByMerchant(req)
	if err != nil {
		s.logger.Error("failed to find monthly nominal success amounts by ID", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindMonthAmountNominalSuccessByMerchant
	}

	responses := s.mapping.ToNominalsResponseMonthAmountSuccess(records)

	s.logger.Debug("Successfully fetched monthly nominal success amounts by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindYearAmountNominalSuccessByMerchant(req *requests.YearAmountNominalByMerchantRequest) ([]*response.YearAmountNominalSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly nominal success amounts by ID", zap.Any("request", req))

	records, err := s.nominalRepository.FindYearAmountNominalSuccessByMerchant(req)
	if err != nil {
		s.logger.Error("failed to find yearly nominal success amounts by ID", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindYearAmountNominalSuccessByMerchant
	}

	responses := s.mapping.ToNominalsResponseYearAmountSuccess(records)

	s.logger.Debug("Successfully fetched yearly nominal success amounts by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindMonthAmountNominalFailedByMerchant(req *requests.MonthAmountNominalByMerchantRequest) ([]*response.MonthAmountNominalFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly nominal failed amounts by ID", zap.Any("request", req))

	records, err := s.nominalRepository.FindMonthAmountNominalFailedByMerchant(req)
	if err != nil {
		s.logger.Error("failed to find monthly nominal failed amounts by ID", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindMonthAmountNominalFailedByMerchant
	}

	responses := s.mapping.ToNominalsResponseMonthAmountFailed(records)

	s.logger.Debug("Successfully fetched monthly nominal failed amounts by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindYearAmountNominalFailedByMerchant(req *requests.YearAmountNominalByMerchantRequest) ([]*response.YearAmountNominalFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly nominal failed amounts by ID", zap.Any("request", req))

	records, err := s.nominalRepository.FindYearAmountNominalFailedByMerchant(req)
	if err != nil {
		s.logger.Error("failed to find yearly nominal failed amounts by ID", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindYearAmountNominalFailedByMerchant
	}

	responses := s.mapping.ToNominalsResponseYearAmountFailed(records)

	s.logger.Debug("Successfully fetched yearly nominal failed amounts by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindMonthMethodNominalSuccessByMerchant(req *requests.MonthMethodNominalByMerchantRequest) ([]*response.MonthMethodNominalResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly nominal success methods by ID", zap.Any("request", req))

	records, err := s.nominalRepository.FindMonthMethodNominalSuccessByMerchant(req)
	if err != nil {
		s.logger.Error("failed to find monthly nominal success methods by ID", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindMonthMethodNominalSuccessByMerchant
	}

	responses := s.mapping.ToNominalsResponseMonthMethodSuccess(records)

	s.logger.Debug("Successfully fetched monthly nominal success methods by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindYearMethodNominalSuccessByMerchant(req *requests.YearMethodNominalByMerchantRequest) ([]*response.YearMethodNominalResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly nominal success methods by ID", zap.Any("request", req))

	records, err := s.nominalRepository.FindYearMethodNominalSuccessByMerchant(req)
	if err != nil {
		s.logger.Error("failed to find yearly nominal success methods by ID", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindYearMethodNominalSuccessByMerchant
	}

	responses := s.mapping.ToNominalsResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly nominal success methods by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindMonthMethodNominalFailedByMerchant(req *requests.MonthMethodNominalByMerchantRequest) ([]*response.MonthMethodNominalResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly nominal failed methods by ID", zap.Any("request", req))

	records, err := s.nominalRepository.FindMonthMethodNominalFailedByMerchant(req)
	if err != nil {
		s.logger.Error("failed to find monthly nominal failed methods by ID", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindMonthMethodNominalFailedByMerchant
	}

	responses := s.mapping.ToNominalsResponseMonthMethodFailed(records)

	s.logger.Debug("Successfully fetched monthly nominal failed methods by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindYearMethodNominalFailedByMerchant(req *requests.YearMethodNominalByMerchantRequest) ([]*response.YearMethodNominalResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly nominal failed methods by ID", zap.Any("request", req))

	records, err := s.nominalRepository.FindYearMethodNominalFailedByMerchant(req)
	if err != nil {
		s.logger.Error("failed to find yearly nominal failed methods by ID", zap.Error(err))
		return nil, nominal_errors.ErrFailedFindYearMethodNominalFailedByMerchant
	}

	responses := s.mapping.ToNominalsResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly nominal failed methods by ID",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *nominalService) FindByID(id int) (*response.NominalResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching nominals by id", zap.Int("nominal_id", id))

	user, err := s.nominalRepository.FindById(id)

	if err != nil {
		s.logger.Error("failed to find nominal by ID", zap.Error(err))
		return nil, nominal_errors.ErrNominalNotFoundRes
	}

	so := s.mapping.ToNominalResponse(user)

	s.logger.Debug("Successfully fetched nominal", zap.Int("nominal_id", id))

	return so, nil
}

func (s *nominalService) FindByActive(request *requests.FindAllNominals) ([]*response.NominalResponseDeleteAt, *int, *response.ErrorResponse) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

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

	users, totalRecords, err := s.nominalRepository.FindByActiveNominal(request)

	if err != nil {
		s.logger.Error("Failed to fetch active nominals",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, nil, nominal_errors.ErrFailedFindActive
	}

	so := s.mapping.ToNominalsResponseDeleteAt(users)

	s.logger.Debug("Successfully fetched active nominals",
		zap.Int("totalRecords", *totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return so, totalRecords, nil
}

func (s *nominalService) FindByTrashed(request *requests.FindAllNominals) ([]*response.NominalResponseDeleteAt, *int, *response.ErrorResponse) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

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

	banks, totalRecords, err := s.nominalRepository.FindByTrashedNominal(request)

	if err != nil {
		s.logger.Error("Failed to find trashed nominals", zap.Error(err))

		return nil, nil, nominal_errors.ErrFailedFindTrashed
	}

	so := s.mapping.ToNominalsResponseDeleteAt(banks)

	s.logger.Debug("Successfully fetched trashed nominals",
		zap.Int("totalRecords", *totalRecords),
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

		return nil, nominal_errors.ErrFailedCreateNominal
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

			return nil, voucher_errors.ErrVoucherNotFoundRes
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

		return nil, nominal_errors.ErrFailedUpdateNominal
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

		return nil, nominal_errors.ErrFailedTrashedNominal
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

		return nil, nominal_errors.ErrFailedRestoreNominal
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

		return false, nominal_errors.ErrFailedDeletePermanent
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
		return false, nominal_errors.ErrFailedRestoreAll
	}

	s.logger.Debug("Successfully restored all nominals")
	return true, nil
}

func (s *nominalService) DeleteAllPermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all nominals")

	_, err := s.nominalRepository.DeleteAllNominalsPermanent()

	if err != nil {
		s.logger.Error("Failed to permanently delete all nominals", zap.Error(err))
		return false, nominal_errors.ErrFailedDeleteAll
	}

	s.logger.Debug("Successfully deleted all nominals permanently")
	return true, nil
}
