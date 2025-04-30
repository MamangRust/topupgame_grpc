package record

type TransactionRecord struct {
	ID            int     `json:"id"`
	UserID        int     `json:"user_id"`
	MerchantID    int     `json:"merchant_id"`
	VoucherID     int     `json:"voucher_id"`
	NominalID     int     `json:"nominal_id"`
	CategoryID    int     `json:"category_id"`
	BankID        int     `json:"bank_id"`
	PaymentMethod string  `json:"payment_method"`
	Status        string  `json:"status"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
	DeletedAt     *string `json:"deleted_at"`
}

type MonthAmountTransactionSuccessRecord struct {
	Year         string `json:"year"`
	Month        string `json:"month"`
	TotalSuccess int    `json:"total_success"`
	TotalAmount  int    `json:"total_amount"`
}

type YearAmountTransactionSuccessRecord struct {
	Year         string `json:"year"`
	TotalSuccess int    `json:"total_success"`
	TotalAmount  int    `json:"total_amount"`
}

type MonthAmountTransactionFailedRecord struct {
	Year        string `json:"year"`
	Month       string `json:"month"`
	TotalFailed int    `json:"total_failed"`
	TotalAmount int    `json:"total_amount"`
}

type YearAmountTransactionFailedRecord struct {
	Year        string `json:"year"`
	TotalFailed int    `json:"total_failed"`
	TotalAmount int    `json:"total_amount"`
}

type MonthMethodTransactionRecord struct {
	Month             string `json:"month"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}

type YearMethodTransactionRecord struct {
	Year              string `json:"year"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}
