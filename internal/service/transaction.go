package service

import (
	"reflect"
	"topup_game/internal/domain/requests"
	"topup_game/internal/domain/response"
	response_service "topup_game/internal/mapper/response/service"
	"topup_game/internal/repository"
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

func (s *transactionService) FindAll(page int, pageSize int, search string) ([]*response.TransactionResponse, int, *response.ErrorResponse) {
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

	res, totalRecords, err := s.transactionRepository.FindAllTransactions(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to fetch transaction",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transaction records",
		}
	}

	s.logger.Debug("Successfully fetched transaction",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	so := s.mapping.ToTransactionsResponse(res)

	return so, totalRecords, nil
}

func (s *transactionService) FindById(id int) (*response.TransactionResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching transaction by ID", zap.Int("id", id))

	res, err := s.transactionRepository.FindById(id)

	if err != nil {
		s.logger.Error("Failed to fetch transaction record by ID", zap.Error(err))

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transaction record by ID",
		}
	}

	s.logger.Debug("Successfully fetched transaction", zap.Int("id", id))

	so := s.mapping.ToTransactionResponse(res)

	return so, nil
}

func (s *transactionService) FindByActive(page int, pageSize int, search string) ([]*response.TransactionResponseDeleteAt, int, *response.ErrorResponse) {
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

	res, totalRecords, err := s.transactionRepository.FindByActive(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to fetch active transaction",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transaction records",
		}
	}

	s.logger.Debug("Successfully fetched active transaction",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	so := s.mapping.ToTransactionsResponseDeleteAt(res)

	return so, totalRecords, nil
}

func (s *transactionService) FindByTrashed(page int, pageSize int, search string) ([]*response.TransactionResponseDeleteAt, int, *response.ErrorResponse) {
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

	res, totalRecords, err := s.transactionRepository.FindByTrashed(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to fetch trashed transaction",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch transaction records",
		}
	}

	s.logger.Debug("Successfully fetched trashed transaction",
		zap.Int("totalRecords", totalRecords),
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

	res, err := s.transactionRepository.CreateTransaction(request, "pending")
	if err != nil {
		s.logger.Error("Failed to create transaction", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create transaction",
		}
	}

	transactionSuccess := true

	if transactionSuccess {
		_, err = s.nominalRepository.UpdateQuantity(request.NominalID, 1)
		if err != nil {
			s.logger.Error("Failed to decrease voucher quantity", zap.Error(err))
			return nil, &response.ErrorResponse{
				Status:  "error",
				Message: "Failed to update voucher stock",
			}
		}

		_, err = s.transactionRepository.UpdateTransactionStatus(res.ID, "success")
		if err != nil {
			s.logger.Error("Failed to update transaction status", zap.Error(err))
			return nil, &response.ErrorResponse{
				Status:  "error",
				Message: "Failed to update transaction status",
			}
		}
		res.Status = "success"
	} else {

		_, err = s.transactionRepository.UpdateTransactionStatus(res.ID, "failed")
		if err != nil {
			s.logger.Error("Failed to update transaction status", zap.Error(err))
			return nil, &response.ErrorResponse{
				Status:  "error",
				Message: "Failed to update transaction status",
			}
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

	_, err := s.transactionRepository.FindById(request.ID)
	if err != nil {
		s.logger.Error("Failed to find transaction", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Transaction not found",
		}
	}

	updatedTransaction, err := s.transactionRepository.UpdateTransaction(request)
	if err != nil {
		s.logger.Error("Failed to update transaction", zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update transaction",
		}
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

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trashed category record",
		}
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

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore transaction record",
		}
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

		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete transaction record",
		}
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
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all Transaction: " + err.Error(),
		}
	}

	s.logger.Debug("Successfully restored all Transaction")
	return true, nil
}

func (s *transactionService) DeleteAllPermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all Transaction")

	_, err := s.transactionRepository.DeleteAllTransactionsPermanent()

	if err != nil {
		s.logger.Error("Failed to permanently delete all Transaction", zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete all Transaction: " + err.Error(),
		}
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
