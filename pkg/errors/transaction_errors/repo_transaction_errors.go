package transaction_errors

import "errors"

var (
	ErrTransactionNotFound     = errors.New("transaction not found")
	ErrFindAllTransactions     = errors.New("failed to find all Transactions")
	ErrFindActiveTransactions  = errors.New("failed to find active Transactions")
	ErrFindTrashedTransactions = errors.New("failed to find trashed Transactions")
	ErrTransactionConflict     = errors.New("failed Transaction already exists")

	ErrCreateTransaction = errors.New("failed to create Transaction")
	ErrUpdateTransaction = errors.New("failed to update Transaction")

	ErrTrashedTransaction         = errors.New("failed to move Transaction to trash")
	ErrRestoreTransaction         = errors.New("failed to restore Transaction from trash")
	ErrDeleteTransactionPermanent = errors.New("failed to permanently delete Transaction")

	ErrRestoreAllTransactions = errors.New("failed to restore all Transactions")
	ErrDeleteAllTransactions  = errors.New("failed to permanently delete all Transactions")

	ErrFindMonthAmountTransactionSuccess = errors.New("failed to find monthly Transaction success amounts")
	ErrFindYearAmountTransactionSuccess  = errors.New("failed to find yearly Transaction success amounts")
	ErrFindMonthAmountTransactionFailed  = errors.New("failed to find monthly Transaction failed amounts")
	ErrFindYearAmountTransactionFailed   = errors.New("failed to find yearly Transaction failed amounts")

	ErrFindMonthMethodTransactionSuccess = errors.New("failed to find monthly Transaction success methods")
	ErrFindYearMethodTransactionSuccess  = errors.New("failed to find yearly Transaction success methods")
	ErrFindMonthMethodTransactionFailed  = errors.New("failed to find monthly Transaction failed methods")
	ErrFindYearMethodTransactionFailed   = errors.New("failed to find yearly Transaction failed methods")

	ErrFindMonthAmountTransactionSuccessById = errors.New("failed to find monthly Transaction success amounts by ID")
	ErrFindYearAmountTransactionSuccessById  = errors.New("failed to find yearly Transaction success amounts by ID")
	ErrFindMonthAmountTransactionFailedById  = errors.New("failed to find monthly Transaction failed amounts by ID")
	ErrFindYearAmountTransactionFailedById   = errors.New("failed to find yearly Transaction failed amounts by ID")

	ErrFindMonthMethodTransactionSuccessById = errors.New("failed to find monthly Transaction success methods by ID")
	ErrFindYearMethodTransactionSuccessById  = errors.New("failed to find yearly Transaction success methods by ID")
	ErrFindMonthMethodTransactionFailedById  = errors.New("failed to find monthly Transaction failed methods by ID")
	ErrFindYearMethodTransactionFailedById   = errors.New("failed to find yearly Transaction failed methods by ID")

	ErrFindMonthAmountTransactionSuccessByMerchant = errors.New("failed to find monthly Transaction success amounts by merchant")
	ErrFindYearAmountTransactionSuccessByMerchant  = errors.New("failed to find yearly Transaction success amounts by merchant")
	ErrFindMonthAmountTransactionFailedByMerchant  = errors.New("failed to find monthly Transaction failed amounts by merchant")
	ErrFindYearAmountTransactionFailedByMerchant   = errors.New("failed to find yearly Transaction failed amounts by merchant")

	ErrFindMonthMethodTransactionSuccessByMerchant = errors.New("failed to find monthly Transaction success methods by merchant")
	ErrFindYearMethodTransactionSuccessByMerchant  = errors.New("failed to find yearly Transaction success methods by merchant")
	ErrFindMonthMethodTransactionFailedByMerchant  = errors.New("failed to find monthly Transaction failed methods by merchant")
	ErrFindYearMethodTransactionFailedByMerchant   = errors.New("failed to find yearly Transaction failed methods by merchant")
)
