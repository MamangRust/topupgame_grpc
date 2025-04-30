package voucher_errors

import "errors"

var (
	ErrVoucherNotFound     = errors.New("voucher not found")
	ErrFindAllVouchers     = errors.New("failed to find all Vouchers")
	ErrFindActiveVouchers  = errors.New("failed to find active Vouchers")
	ErrFindTrashedVouchers = errors.New("failed to find trashed Vouchers")
	ErrVoucherConflict     = errors.New("failed Voucher already exists")

	ErrCreateVoucher = errors.New("failed to create Voucher")
	ErrUpdateVoucher = errors.New("failed to update Voucher")

	ErrTrashedVoucher         = errors.New("failed to move Voucher to trash")
	ErrRestoreVoucher         = errors.New("failed to restore Voucher from trash")
	ErrDeleteVoucherPermanent = errors.New("failed to permanently delete Voucher")

	ErrRestoreAllVouchers = errors.New("failed to restore all Vouchers")
	ErrDeleteAllVouchers  = errors.New("failed to permanently delete all Vouchers")

	ErrFindMonthAmountVoucherSuccess = errors.New("failed to find monthly Voucher success amounts")
	ErrFindYearAmountVoucherSuccess  = errors.New("failed to find yearly Voucher success amounts")
	ErrFindMonthAmountVoucherFailed  = errors.New("failed to find monthly Voucher failed amounts")
	ErrFindYearAmountVoucherFailed   = errors.New("failed to find yearly Voucher failed amounts")

	ErrFindMonthMethodVoucherSuccess = errors.New("failed to find monthly Voucher success methods")
	ErrFindYearMethodVoucherSuccess  = errors.New("failed to find yearly Voucher success methods")
	ErrFindMonthMethodVoucherFailed  = errors.New("failed to find monthly Voucher failed methods")
	ErrFindYearMethodVoucherFailed   = errors.New("failed to find yearly Voucher failed methods")

	ErrFindMonthAmountVoucherSuccessById = errors.New("failed to find monthly Voucher success amounts by ID")
	ErrFindYearAmountVoucherSuccessById  = errors.New("failed to find yearly Voucher success amounts by ID")
	ErrFindMonthAmountVoucherFailedById  = errors.New("failed to find monthly Voucher failed amounts by ID")
	ErrFindYearAmountVoucherFailedById   = errors.New("failed to find yearly Voucher failed amounts by ID")

	ErrFindMonthMethodVoucherSuccessById = errors.New("failed to find monthly Voucher success methods by ID")
	ErrFindYearMethodVoucherSuccessById  = errors.New("failed to find yearly Voucher success methods by ID")
	ErrFindMonthMethodVoucherFailedById  = errors.New("failed to find monthly Voucher failed methods by ID")
	ErrFindYearMethodVoucherFailedById   = errors.New("failed to find yearly Voucher failed methods by ID")

	ErrFindMonthAmountVoucherSuccessByMerchant = errors.New("failed to find monthly Voucher success amounts by merchant")
	ErrFindYearAmountVoucherSuccessByMerchant  = errors.New("failed to find yearly Voucher success amounts by merchant")
	ErrFindMonthAmountVoucherFailedByMerchant  = errors.New("failed to find monthly Voucher failed amounts by merchant")
	ErrFindYearAmountVoucherFailedByMerchant   = errors.New("failed to find yearly Voucher failed amounts by merchant")

	ErrFindMonthMethodVoucherSuccessByMerchant = errors.New("failed to find monthly Voucher success methods by merchant")
	ErrFindYearMethodVoucherSuccessByMerchant  = errors.New("failed to find yearly Voucher success methods by merchant")
	ErrFindMonthMethodVoucherFailedByMerchant  = errors.New("failed to find monthly Voucher failed methods by merchant")
	ErrFindYearMethodVoucherFailedByMerchant   = errors.New("failed to find yearly Voucher failed methods by merchant")
)
