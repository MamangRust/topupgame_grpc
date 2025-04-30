package service

import (
	"math"
	"reflect"
	"topup_game/internal/domain/requests"
	"topup_game/internal/domain/response"
	response_service "topup_game/internal/mapper/response/service"
	"topup_game/internal/repository"
	"topup_game/pkg/errors/nominal_errors"
	"topup_game/pkg/errors/transaction_errors"
	"topup_game/pkg/logger"

	"go.uber.org/zap"
)

type transactionService struct {
	userRepository        repository.UserRepository
	merchantRepository    repository.MerchantRepository
	voucherRepository     repository.VoucherRepository
	nominalRepository     repository.NominalRepository
	categoryRepository    repository.CategoryRepository
	bankRepository        repository.BankRepository
	transactionRepository repository.TransactionRepository
	logger                logger.LoggerInterface
	mapping               response_service.TransactionResponseMapper
}

func NewTransactionService(
	userRepository repository.UserRepository,
	merchantRepository repository.MerchantRepository,
	voucherRepository repository.VoucherRepository,
	nominalRepository repository.NominalRepository,
	categoryRepository repository.CategoryRepository,
	bankRepository repository.BankRepository,
	transactionRepository repository.TransactionRepository,
	logger logger.LoggerInterface,
	mapping response_service.TransactionResponseMapper,
) *transactionService {
	return &transactionService{
		userRepository:        userRepository,
		merchantRepository:    merchantRepository,
		voucherRepository:     voucherRepository,
		nominalRepository:     nominalRepository,
		categoryRepository:    categoryRepository,
		bankRepository:        bankRepository,
		transactionRepository: transactionRepository,
		logger:                logger,
		mapping:               mapping,
	}
}

func (s *transactionService) FindAll(request *requests.FindAllTransactions) ([]*response.TransactionResponse, *int, *response.ErrorResponse) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	s.logger.Debug("Fetching transaction",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.transactionRepository.FindAllTransactions(request)

	if err != nil {
		s.logger.Error("Failed to fetch transaction",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, nil, transaction_errors.ErrFailedFindAll
	}

	s.logger.Debug("Successfully fetched transaction",
		zap.Int("totalRecords", *totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	so := s.mapping.ToTransactionsResponse(res)

	return so, totalRecords, nil
}

func (s *transactionService) FindMonthAmountTransactionSuccess(req *requests.MonthAmountTransactionRequest) ([]*response.MonthAmountTransactionSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly transaction success amounts", zap.Any("request", req))

	records, err := s.transactionRepository.FindMonthAmountTransactionSuccess(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly transaction success amounts", zap.Error(err))
		return nil, transaction_errors.ErrFailedFindMonthAmountTransactionSuccess
	}

	responses := s.mapping.ToTransactionsResponseMonthAmountSuccess(records)

	s.logger.Debug("Successfully fetched monthly transaction success amounts",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *transactionService) FindYearAmountTransactionSuccess(year int) ([]*response.YearAmountTransactionSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly transaction success amounts", zap.Int("year", year))

	records, err := s.transactionRepository.FindYearAmountTransactionSuccess(year)
	if err != nil {
		s.logger.Error("Failed to fetch yearly transaction success amounts", zap.Error(err))
		return nil, transaction_errors.ErrFailedFindYearAmountTransactionSuccess
	}

	responses := s.mapping.ToTransactionsResponseYearAmountSuccess(records)

	s.logger.Debug("Successfully fetched yearly transaction success amounts",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *transactionService) FindMonthAmountTransactionFailed(req *requests.MonthAmountTransactionRequest) ([]*response.MonthAmountTransactionFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly transaction failed amounts", zap.Any("request", req))

	records, err := s.transactionRepository.FindMonthAmountTransactionFailed(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly transaction failed amounts", zap.Error(err))
		return nil, transaction_errors.ErrFailedFindMonthAmountTransactionFailed
	}

	responses := s.mapping.ToTransactionsResponseMonthAmountFailed(records)

	s.logger.Debug("Successfully fetched monthly transaction failed amounts",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *transactionService) FindYearAmountTransactionFailed(year int) ([]*response.YearAmountTransactionFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly transaction failed amounts", zap.Int("year", year))

	records, err := s.transactionRepository.FindYearAmountTransactionFailed(year)
	if err != nil {
		s.logger.Error("Failed to fetch yearly transaction failed amounts", zap.Error(err))
		return nil, transaction_errors.ErrFailedFindYearAmountTransactionFailed
	}

	responses := s.mapping.ToTransactionsResponseYearAmountFailed(records)

	s.logger.Debug("Successfully fetched yearly transaction failed amounts",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *transactionService) FindMonthMethodTransactionSuccess(year int) ([]*response.MonthMethodTransactionResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly transaction success methods", zap.Int("year", year))

	records, err := s.transactionRepository.FindMonthMethodTransactionSuccess(year)
	if err != nil {
		s.logger.Error("Failed to fetch monthly transaction success methods", zap.Error(err))
		return nil, transaction_errors.ErrFailedFindMonthMethodTransactionSuccess
	}

	responses := s.mapping.ToTransactionsResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly transaction success methods",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *transactionService) FindYearMethodTransactionSuccess(year int) ([]*response.YearMethodTransactionResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly transaction success methods", zap.Int("year", year))

	records, err := s.transactionRepository.FindYearMethodTransactionSuccess(year)
	if err != nil {
		s.logger.Error("Failed to fetch yearly transaction success methods", zap.Error(err))
		return nil, transaction_errors.ErrFailedFindYearMethodTransactionSuccess
	}

	responses := s.mapping.ToTransactionsResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly transaction success methods",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *transactionService) FindMonthMethodTransactionFailed(year int) ([]*response.MonthMethodTransactionResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly transaction failed methods", zap.Int("year", year))

	records, err := s.transactionRepository.FindMonthMethodTransactionFailed(year)
	if err != nil {
		s.logger.Error("Failed to fetch monthly transaction failed methods", zap.Error(err))
		return nil, transaction_errors.ErrFailedFindMonthMethodTransactionFailed
	}

	responses := s.mapping.ToTransactionsResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly transaction failed methods",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *transactionService) FindYearMethodTransactionFailed(year int) ([]*response.YearMethodTransactionResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly transaction failed methods", zap.Int("year", year))

	records, err := s.transactionRepository.FindYearMethodTransactionFailed(year)
	if err != nil {
		s.logger.Error("Failed to fetch yearly transaction failed methods", zap.Error(err))
		return nil, transaction_errors.ErrFailedFindYearMethodTransactionFailed
	}

	responses := s.mapping.ToTransactionsResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly transaction failed methods",
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *transactionService) FindMonthAmountTransactionSuccessByMerchant(req *requests.MonthAmountTransactionByMerchantRequest) ([]*response.MonthAmountTransactionSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly transaction success amounts by merchant",
		zap.Any("request", req),
		zap.Int("merchant_id", req.MerchantID))

	records, err := s.transactionRepository.FindMonthAmountTransactionSuccessByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly transaction success amounts by merchant",
			zap.Error(err),
			zap.Int("merchant_id", req.MerchantID))
		return nil, transaction_errors.ErrFailedFindMonthAmountTransactionSuccessByMerchant
	}

	responses := s.mapping.ToTransactionsResponseMonthAmountSuccess(records)

	s.logger.Debug("Successfully fetched monthly transaction success amounts by merchant",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *transactionService) FindYearAmountTransactionSuccessByMerchant(req *requests.YearAmountTransactionByMerchantRequest) ([]*response.YearAmountTransactionSuccessResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly transaction success amounts by merchant",
		zap.Any("request", req),
		zap.Int("merchant_id", req.MerchantID))

	records, err := s.transactionRepository.FindYearAmountTransactionSuccessByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch yearly transaction success amounts by merchant",
			zap.Error(err),
			zap.Int("merchant_id", req.MerchantID))
		return nil, transaction_errors.ErrFailedFindYearAmountTransactionSuccessByMerchant
	}

	responses := s.mapping.ToTransactionsResponseYearAmountSuccess(records)

	s.logger.Debug("Successfully fetched yearly transaction success amounts by merchant",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *transactionService) FindMonthAmountTransactionFailedByMerchant(req *requests.MonthAmountTransactionByMerchantRequest) ([]*response.MonthAmountTransactionFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly transaction failed amounts by merchant",
		zap.Any("request", req),
		zap.Int("merchant_id", req.MerchantID))

	records, err := s.transactionRepository.FindMonthAmountTransactionFailedByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly transaction failed amounts by merchant",
			zap.Error(err),
			zap.Int("merchant_id", req.MerchantID))
		return nil, transaction_errors.ErrFailedFindMonthAmountTransactionFailedByMerchant
	}

	responses := s.mapping.ToTransactionsResponseMonthAmountFailed(records)

	s.logger.Debug("Successfully fetched monthly transaction failed amounts by merchant",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *transactionService) FindYearAmountTransactionFailedByMerchant(req *requests.YearAmountTransactionByMerchantRequest) ([]*response.YearAmountTransactionFailedResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly transaction failed amounts by merchant",
		zap.Any("request", req),
		zap.Int("merchant_id", req.MerchantID))

	records, err := s.transactionRepository.FindYearAmountTransactionFailedByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch yearly transaction failed amounts by merchant",
			zap.Error(err),
			zap.Int("merchant_id", req.MerchantID))
		return nil, transaction_errors.ErrFailedFindYearAmountTransactionFailedByMerchant
	}

	responses := s.mapping.ToTransactionsResponseYearAmountFailed(records)

	s.logger.Debug("Successfully fetched yearly transaction failed amounts by merchant",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *transactionService) FindMonthMethodTransactionSuccessByMerchant(req *requests.MonthMethodTransactionByMerchantRequest) ([]*response.MonthMethodTransactionResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly transaction success methods by merchant",
		zap.Any("request", req),
		zap.Int("merchant_id", req.MerchantID))

	records, err := s.transactionRepository.FindMonthMethodTransactionSuccessByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly transaction success methods by merchant",
			zap.Error(err),
			zap.Int("merchant_id", req.MerchantID))
		return nil, transaction_errors.ErrFailedFindMonthMethodTransactionSuccessByMerchant
	}

	responses := s.mapping.ToTransactionsResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly transaction success methods by merchant",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *transactionService) FindYearMethodTransactionSuccessByMerchant(req *requests.YearMethodTransactionByMerchantRequest) ([]*response.YearMethodTransactionResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly transaction success methods by merchant",
		zap.Any("request", req),
		zap.Int("merchant_id", req.MerchantID))

	records, err := s.transactionRepository.FindYearMethodTransactionSuccessByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch yearly transaction success methods by merchant",
			zap.Error(err),
			zap.Int("merchant_id", req.MerchantID))
		return nil, transaction_errors.ErrFailedFindYearMethodTransactionSuccessByMerchant
	}

	responses := s.mapping.ToTransactionsResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly transaction success methods by merchant",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *transactionService) FindMonthMethodTransactionFailedByMerchant(req *requests.MonthMethodTransactionByMerchantRequest) ([]*response.MonthMethodTransactionResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching monthly transaction failed methods by merchant",
		zap.Any("request", req),
		zap.Int("merchant_id", req.MerchantID))

	records, err := s.transactionRepository.FindMonthMethodTransactionFailedByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch monthly transaction failed methods by merchant",
			zap.Error(err),
			zap.Int("merchant_id", req.MerchantID))
		return nil, transaction_errors.ErrFailedFindMonthMethodTransactionFailedByMerchant
	}

	responses := s.mapping.ToTransactionsResponseMonthMethod(records)

	s.logger.Debug("Successfully fetched monthly transaction failed methods by merchant",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *transactionService) FindYearMethodTransactionFailedByMerchant(req *requests.YearMethodTransactionByMerchantRequest) ([]*response.YearMethodTransactionResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching yearly transaction failed methods by merchant",
		zap.Any("request", req),
		zap.Int("merchant_id", req.MerchantID))

	records, err := s.transactionRepository.FindYearMethodTransactionFailedByMerchant(req)
	if err != nil {
		s.logger.Error("Failed to fetch yearly transaction failed methods by merchant",
			zap.Error(err),
			zap.Int("merchant_id", req.MerchantID))
		return nil, transaction_errors.ErrFailedFindYearMethodTransactionFailedByMerchant
	}

	responses := s.mapping.ToTransactionsResponseYearMethod(records)

	s.logger.Debug("Successfully fetched yearly transaction failed methods by merchant",
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(responses)))

	return responses, nil
}

func (s *transactionService) FindById(id int) (*response.TransactionResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching transaction by ID", zap.Int("id", id))

	res, err := s.transactionRepository.FindById(id)

	if err != nil {
		s.logger.Error("Failed to fetch transaction record by ID", zap.Error(err))

		return nil, transaction_errors.ErrTransactionNotFoundRes
	}

	s.logger.Debug("Successfully fetched transaction", zap.Int("id", id))

	so := s.mapping.ToTransactionResponse(res)

	return so, nil
}

func (s *transactionService) FindByActive(request *requests.FindAllTransactions) ([]*response.TransactionResponseDeleteAt, *int, *response.ErrorResponse) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	s.logger.Debug("Fetching active transaction",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.transactionRepository.FindByActive(request)

	if err != nil {
		s.logger.Error("Failed to fetch active transaction",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, nil, transaction_errors.ErrFailedFindActive
	}

	s.logger.Debug("Successfully fetched active transaction",
		zap.Int("totalRecords", *totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	so := s.mapping.ToTransactionsResponseDeleteAt(res)

	return so, totalRecords, nil
}

func (s *transactionService) FindByTrashed(request *requests.FindAllTransactions) ([]*response.TransactionResponseDeleteAt, *int, *response.ErrorResponse) {
	page := request.Page
	pageSize := request.PageSize
	search := request.Search

	s.logger.Debug("Fetching trashed role",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	res, totalRecords, err := s.transactionRepository.FindByTrashed(request)

	if err != nil {
		s.logger.Error("Failed to fetch trashed transaction",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, nil, transaction_errors.ErrFailedFindTrashed
	}

	s.logger.Debug("Successfully fetched trashed transaction",
		zap.Int("totalRecords", *totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	so := s.mapping.ToTransactionsResponseDeleteAt(res)

	return so, totalRecords, nil
}

func (s *transactionService) Create(request *requests.CreateTransactionRequest) (*response.TransactionResponse, *response.ErrorResponse) {
	s.logger.Debug("Starting Create Transaction process",
		zap.Int("user_id", request.UserID),
		zap.Int("merchant_id", request.MerchantID),
		zap.Int("voucher_id", request.VoucherID),
		zap.Int("nominal_id", request.NominalID),
		zap.Int("category_id", request.CategoryID),
		zap.Int("bank_id", request.BankID),
		zap.String("payment_method", request.PaymentMethod),
	)

	if err := validFindEntities(s, request); err != nil {
		return nil, err
	}

	nominal, err := s.nominalRepository.FindById(request.NominalID)
	if err != nil {
		s.logger.Error("Failed to find nominal", zap.Error(err))
		return nil, nominal_errors.ErrNominalNotFoundRes
	}

	basePrice := nominal.Price
	ppn := basePrice * 0.11
	totalAmount := basePrice + ppn

	amountInt := int(math.Round(totalAmount))
	request.Amount = &amountInt

	res, err := s.transactionRepository.CreateTransaction(request, "pending")
	if err != nil {
		s.logger.Error("Failed to create transaction", zap.Error(err))
		return nil, transaction_errors.ErrFailedCreateTransaction
	}

	transactionSuccess := true

	if transactionSuccess {
		_, err = s.nominalRepository.UpdateQuantity(request.NominalID, 1)
		if err != nil {
			s.logger.Error("Failed to decrease voucher quantity", zap.Error(err))
			return nil, nominal_errors.ErrFailedUpdateNominal
		}

		_, err = s.transactionRepository.UpdateTransactionStatus(res.ID, "success")
		if err != nil {
			s.logger.Error("Failed to update transaction status", zap.Error(err))
			return nil, transaction_errors.ErrFailedUpdateTransaction
		}
		res.Status = "success"
	} else {
		_, err = s.transactionRepository.UpdateTransactionStatus(res.ID, "failed")
		if err != nil {
			s.logger.Error("Failed to update transaction status", zap.Error(err))
			return nil, transaction_errors.ErrFailedUpdateTransaction
		}
		res.Status = "failed"
	}

	s.logger.Debug("Create Transaction process completed",
		zap.Int("transaction_id", res.ID),
		zap.String("status", res.Status),
	)

	so := s.mapping.ToTransactionResponse(res)

	return so, nil
}

func (s *transactionService) Update(request *requests.UpdateTransactionRequest) (*response.TransactionResponse, *response.ErrorResponse) {
	s.logger.Debug("Starting Update Transaction process",
		zap.Int("transaction_id", request.ID),
		zap.Int("user_id", request.UserID),
		zap.Int("merchant_id", request.MerchantID),
		zap.Int("voucher_id", request.VoucherID),
		zap.Int("nominal_id", request.NominalID),
		zap.Int("category_id", request.CategoryID),
		zap.Int("bank_id", request.BankID),
		zap.String("payment_method", request.PaymentMethod),
	)

	if err := validFindEntities(s, request); err != nil {
		return nil, err
	}

	nominal, err := s.nominalRepository.FindById(request.NominalID)
	if err != nil {
		s.logger.Error("Failed to find nominal", zap.Error(err))
		return nil, nominal_errors.ErrNominalNotFoundRes
	}

	basePrice := nominal.Price
	ppn := basePrice * 0.11
	totalAmount := basePrice + ppn

	amountInt := int(math.Round(totalAmount))
	request.Amount = &amountInt

	updatedTransaction, err := s.transactionRepository.UpdateTransaction(request)
	if err != nil {
		s.logger.Error("Failed to update transaction", zap.Error(err))
		return nil, transaction_errors.ErrFailedUpdateTransaction
	}

	s.logger.Debug("Update Transaction process completed",
		zap.Int("transaction_id", updatedTransaction.ID),
		zap.String("status", updatedTransaction.Status),
	)

	so := s.mapping.ToTransactionResponse(updatedTransaction)

	return so, nil
}

func (s *transactionService) Trashed(id int) (*response.TransactionResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Starting Trashed Transaction process",
		zap.Int("transaction_id", id),
	)

	res, err := s.transactionRepository.TrashTransaction(id)

	if err != nil {
		s.logger.Error("Failed to move transaction to trash",
			zap.Int("transaction_id", id),
			zap.Error(err),
		)

		return nil, transaction_errors.ErrFailedTrashedTransaction
	}

	so := s.mapping.ToTransactionResponseDeleteAt(res)

	s.logger.Debug("Trashed Transaction process completed",
		zap.Int("transaction_id", id),
	)

	return so, nil
}

func (s *transactionService) Restore(id int) (*response.TransactionResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Starting RestoreTransaction process",
		zap.Int("transaction_id", id),
	)

	res, err := s.transactionRepository.RestoreTransaction(id)

	if err != nil {
		s.logger.Error("Failed to restore transaction", zap.Error(err))

		return nil, transaction_errors.ErrFailedRestoreTransaction
	}

	so := s.mapping.ToTransactionResponseDeleteAt(res)

	s.logger.Debug("RestoreTransaction process completed",
		zap.Int("transaction_id", id),
	)

	return so, nil
}

func (s *transactionService) DeletePermanent(id int) (bool, *response.ErrorResponse) {
	s.logger.Debug("Starting DeleteTransactionPermanent process",
		zap.Int("transaction_id", id),
	)

	_, err := s.transactionRepository.DeleteTransactionPermanent(id)

	if err != nil {
		s.logger.Error("Failed to delete transaction permanently",
			zap.Int("category_id", id),
			zap.Error(err),
		)

		return false, transaction_errors.ErrFailedDeletePermanent
	}

	s.logger.Debug("DeleteTransactionPermanent process completed",
		zap.Int("transaction_id", id),
	)

	return true, nil
}

func (s *transactionService) RestoreAll() (bool, *response.ErrorResponse) {
	s.logger.Debug("Restoring all Transaction")

	_, err := s.transactionRepository.RestoreAllTransactions()

	if err != nil {
		s.logger.Error("Failed to restore all Transaction", zap.Error(err))
		return false, transaction_errors.ErrFailedRestoreAll
	}

	s.logger.Debug("Successfully restored all Transaction")
	return true, nil
}

func (s *transactionService) DeleteAllPermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all Transaction")

	_, err := s.transactionRepository.DeleteAllTransactionsPermanent()

	if err != nil {
		s.logger.Error("Failed to permanently delete all Transaction", zap.Error(err))
		return false, transaction_errors.ErrFailedDeleteAll
	}

	s.logger.Debug("Successfully deleted all Transaction permanently")
	return true, nil
}

func validFindEntities[T any](s *transactionService, req T) *response.ErrorResponse {
	v := reflect.ValueOf(req)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	merchantID := v.FieldByName("MerchantID").Int()
	if _, err := s.merchantRepository.FindById(int(merchantID)); err != nil {
		return &response.ErrorResponse{Status: "error", Message: "Merchant not found"}
	}

	voucherID := v.FieldByName("VoucherID").Int()
	if _, err := s.voucherRepository.FindById(int(voucherID)); err != nil {
		return &response.ErrorResponse{Status: "error", Message: "Voucher not found"}
	}

	categoryID := v.FieldByName("CategoryID").Int()
	if _, err := s.categoryRepository.FindById(int(categoryID)); err != nil {
		return &response.ErrorResponse{Status: "error", Message: "Category not found"}
	}

	bankID := v.FieldByName("BankID").Int()
	if _, err := s.bankRepository.FindById(int(bankID)); err != nil {
		return &response.ErrorResponse{Status: "error", Message: "Bank not found"}
	}

	return nil
}
