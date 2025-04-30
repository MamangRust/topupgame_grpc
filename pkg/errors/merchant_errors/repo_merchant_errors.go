package merchant_errors

import "errors"

var (
	ErrMerchantNotFound     = errors.New("merchant not found")
	ErrFindAllMerchants     = errors.New("failed to find all Merchants")
	ErrFindActiveMerchants  = errors.New("failed to find active Merchants")
	ErrFindTrashedMerchants = errors.New("failed to find trashed Merchants")
	ErrMerchantConflict     = errors.New("failed Merchant already exists")

	ErrCreateMerchant = errors.New("failed to create Merchant")
	ErrUpdateMerchant = errors.New("failed to update Merchant")

	ErrTrashedMerchant         = errors.New("failed to move Merchant to trash")
	ErrRestoreMerchant         = errors.New("failed to restore Merchant from trash")
	ErrDeleteMerchantPermanent = errors.New("failed to permanently delete Merchant")

	ErrRestoreAllMerchants = errors.New("failed to restore all Merchants")
	ErrDeleteAllMerchants  = errors.New("failed to permanently delete all Merchants")
)
