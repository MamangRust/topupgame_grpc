package record

type VoucherRecord struct {
	ID         int     `json:"id"`
	MerchantID int     `json:"merchant_id"`
	CategoryID int     `json:"category_id"`
	Name       string  `json:"name"`
	ImageName  string  `json:"image_name"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	DeletedAt  *string `json:"deleted_at"`
}

type MonthAmountVoucherSuccessRecord struct {
	ID           int    `json:"id"`
	VoucherName  string `json:"voucher_name"`
	Year         string `json:"year"`
	Month        string `json:"month"`
	TotalSuccess int    `json:"total_success"`
	TotalAmount  int    `json:"total_amount"`
}

type YearAmountVoucherSuccessRecord struct {
	ID           int    `json:"id"`
	VoucherName  string `json:"voucher_name"`
	Year         string `json:"year"`
	TotalSuccess int    `json:"total_success"`
	TotalAmount  int    `json:"total_amount"`
}

type MonthAmountVoucherFailedRecord struct {
	ID          int    `json:"id"`
	VoucherName string `json:"voucher_name"`
	Year        string `json:"year"`
	Month       string `json:"month"`
	TotalFailed int    `json:"total_failed"`
	TotalAmount int    `json:"total_amount"`
}

type YearAmountVoucherFailedRecord struct {
	ID          int    `json:"id"`
	VoucherName string `json:"voucher_name"`
	Year        string `json:"year"`
	TotalFailed int    `json:"total_failed"`
	TotalAmount int    `json:"total_amount"`
}

type MonthMethodVoucherRecord struct {
	ID                int    `json:"id"`
	Month             string `json:"month"`
	VoucherName       string `json:"voucher_name"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}

type YearMethodVoucherRecord struct {
	Year              string `json:"year"`
	ID                int    `json:"id"`
	VoucherName       string `json:"voucher_name"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}
