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
