package response

type TransactionResponse struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	MerchantID    int    `json:"merchant_id"`
	VoucherID     int    `json:"voucher_id"`
	NominalID     int    `json:"nominal_id"`
	CategoryID    int    `json:"category_id"`
	BankID        int    `json:"bank_id"`
	PaymentMethod string `json:"payment_method"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type TransactionResponseDeleteAt struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	MerchantID    int    `json:"merchant_id"`
	VoucherID     int    `json:"voucher_id"`
	NominalID     int    `json:"nominal_id"`
	CategoryID    int    `json:"category_id"`
	BankID        int    `json:"bank_id"`
	PaymentMethod string `json:"payment_method"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeletedAt     string `json:"deleted_at"`
}

type ApiResponseTransactionAll struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ApiResponseTransactionDelete struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ApiResponseTransaction struct {
	Status  string               `json:"status"`
	Message string               `json:"message"`
	Data    *TransactionResponse `json:"data"`
}

type ApiResponseTransactionDeleteAt struct {
	Status  string                       `json:"status"`
	Message string                       `json:"message"`
	Data    *TransactionResponseDeleteAt `json:"data"`
}

type ApiResponsesTransaction struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    []*TransactionResponse `json:"data"`
}

type ApiResponsePaginationTransaction struct {
	Status     string                 `json:"status"`
	Message    string                 `json:"message"`
	Data       []*TransactionResponse `json:"data"`
	Pagination *PaginationMeta        `json:"pagination"`
}

type ApiResponsePaginationTransactionDeleteAt struct {
	Status     string                         `json:"status"`
	Message    string                         `json:"message"`
	Data       []*TransactionResponseDeleteAt `json:"data"`
	Pagination *PaginationMeta                `json:"pagination"`
}
