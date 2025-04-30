package nominal_errors

import "errors"

var (
	ErrNominalNotFound     = errors.New("nominal not found")
	ErrFindAllNominals     = errors.New("failed to find all Nominals")
	ErrFindActiveNominals  = errors.New("failed to find active Nominals")
	ErrFindTrashedNominals = errors.New("failed to find trashed Nominals")
	ErrNominalConflict     = errors.New("failed Nominal already exists")

	ErrCreateNominal = errors.New("failed to create Nominal")
	ErrUpdateNominal = errors.New("failed to update Nominal")

	ErrTrashedNominal         = errors.New("failed to move Nominal to trash")
	ErrRestoreNominal         = errors.New("failed to restore Nominal from trash")
	ErrDeleteNominalPermanent = errors.New("failed to permanently delete Nominal")

	ErrRestoreAllNominals = errors.New("failed to restore all Nominals")
	ErrDeleteAllNominals  = errors.New("failed to permanently delete all Nominals")

	ErrFindMonthAmountNominalSuccess = errors.New("failed to find monthly Nominal success amounts")
	ErrFindYearAmountNominalSuccess  = errors.New("failed to find yearly Nominal success amounts")
	ErrFindMonthAmountNominalFailed  = errors.New("failed to find monthly Nominal failed amounts")
	ErrFindYearAmountNominalFailed   = errors.New("failed to find yearly Nominal failed amounts")

	ErrFindMonthMethodNominalSuccess = errors.New("failed to find monthly Nominal success methods")
	ErrFindYearMethodNominalSuccess  = errors.New("failed to find yearly Nominal success methods")
	ErrFindMonthMethodNominalFailed  = errors.New("failed to find monthly Nominal failed methods")
	ErrFindYearMethodNominalFailed   = errors.New("failed to find yearly Nominal failed methods")

	ErrFindMonthAmountNominalSuccessById = errors.New("failed to find monthly Nominal success amounts by ID")
	ErrFindYearAmountNominalSuccessById  = errors.New("failed to find yearly Nominal success amounts by ID")
	ErrFindMonthAmountNominalFailedById  = errors.New("failed to find monthly Nominal failed amounts by ID")
	ErrFindYearAmountNominalFailedById   = errors.New("failed to find yearly Nominal failed amounts by ID")

	ErrFindMonthMethodNominalSuccessById = errors.New("failed to find monthly Nominal success methods by ID")
	ErrFindYearMethodNominalSuccessById  = errors.New("failed to find yearly Nominal success methods by ID")
	ErrFindMonthMethodNominalFailedById  = errors.New("failed to find monthly Nominal failed methods by ID")
	ErrFindYearMethodNominalFailedById   = errors.New("failed to find yearly Nominal failed methods by ID")

	ErrFindMonthAmountNominalSuccessByMerchant = errors.New("failed to find monthly Nominal success amounts by merchant")
	ErrFindYearAmountNominalSuccessByMerchant  = errors.New("failed to find yearly Nominal success amounts by merchant")
	ErrFindMonthAmountNominalFailedByMerchant  = errors.New("failed to find monthly Nominal failed amounts by merchant")
	ErrFindYearAmountNominalFailedByMerchant   = errors.New("failed to find yearly Nominal failed amounts by merchant")

	ErrFindMonthMethodNominalSuccessByMerchant = errors.New("failed to find monthly Nominal success methods by merchant")
	ErrFindYearMethodNominalSuccessByMerchant  = errors.New("failed to find yearly Nominal success methods by merchant")
	ErrFindMonthMethodNominalFailedByMerchant  = errors.New("failed to find monthly Nominal failed methods by merchant")
	ErrFindYearMethodNominalFailedByMerchant   = errors.New("failed to find yearly Nominal failed methods by merchant")
)
