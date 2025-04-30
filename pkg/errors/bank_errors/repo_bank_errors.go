package bank_errors

import "errors"

var (
	ErrBankNotFound     = errors.New("bank not found")
	ErrFindAllBanks     = errors.New("failed to find all Banks")
	ErrFindActiveBanks  = errors.New("failed to find active Banks")
	ErrFindTrashedBanks = errors.New("failed to find trashed Banks")
	ErrBankConflict     = errors.New("failed Bank already exists")

	ErrCreateBank = errors.New("failed to create Bank")
	ErrUpdateBank = errors.New("failed to update Bank")

	ErrTrashedBank         = errors.New("failed to move Bank to trash")
	ErrRestoreBank         = errors.New("failed to restore Bank from trash")
	ErrDeleteBankPermanent = errors.New("failed to permanently delete Bank")

	ErrRestoreAllBanks = errors.New("failed to restore all Banks")
	ErrDeleteAllBanks  = errors.New("failed to permanently delete all Banks")

	ErrFindMonthAmountBankSuccess = errors.New("failed to find monthly bank success amounts")
	ErrFindYearAmountBankSuccess  = errors.New("failed to find yearly bank success amounts")
	ErrFindMonthAmountBankFailed  = errors.New("failed to find monthly bank failed amounts")
	ErrFindYearAmountBankFailed   = errors.New("failed to find yearly bank failed amounts")

	ErrFindMonthMethodBankSuccess = errors.New("failed to find monthly bank success methods")
	ErrFindYearMethodBankSuccess  = errors.New("failed to find yearly bank success methods")
	ErrFindMonthMethodBankFailed  = errors.New("failed to find monthly bank failed methods")
	ErrFindYearMethodBankFailed   = errors.New("failed to find yearly bank failed methods")

	ErrFindMonthAmountBankSuccessById = errors.New("failed to find monthly bank success amounts by ID")
	ErrFindYearAmountBankSuccessById  = errors.New("failed to find yearly bank success amounts by ID")
	ErrFindMonthAmountBankFailedById  = errors.New("failed to find monthly bank failed amounts by ID")
	ErrFindYearAmountBankFailedById   = errors.New("failed to find yearly bank failed amounts by ID")

	ErrFindMonthMethodBankSuccessById = errors.New("failed to find monthly bank success methods by ID")
	ErrFindYearMethodBankSuccessById  = errors.New("failed to find yearly bank success methods by ID")
	ErrFindMonthMethodBankFailedById  = errors.New("failed to find monthly bank failed methods by ID")
	ErrFindYearMethodBankFailedById   = errors.New("failed to find yearly bank failed methods by ID")

	ErrFindMonthAmountBankSuccessByMerchant = errors.New("failed to find monthly bank success amounts by merchant")
	ErrFindYearAmountBankSuccessByMerchant  = errors.New("failed to find yearly bank success amounts by merchant")
	ErrFindMonthAmountBankFailedByMerchant  = errors.New("failed to find monthly bank failed amounts by merchant")
	ErrFindYearAmountBankFailedByMerchant   = errors.New("failed to find yearly bank failed amounts by merchant")

	ErrFindMonthMethodBankSuccessByMerchant = errors.New("failed to find monthly bank success methods by merchant")
	ErrFindYearMethodBankSuccessByMerchant  = errors.New("failed to find yearly bank success methods by merchant")
	ErrFindMonthMethodBankFailedByMerchant  = errors.New("failed to find monthly bank failed methods by merchant")
	ErrFindYearMethodBankFailedByMerchant   = errors.New("failed to find yearly bank failed methods by merchant")
)
