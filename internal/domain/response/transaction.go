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

type MonthAmountTransactionSuccessResponse struct {
	Year         string `json:"year"`
	Month        string `json:"month"`
	TotalSuccess int    `json:"total_success"`
	TotalAmount  int    `json:"total_amount"`
}

type YearAmountTransactionSuccessResponse struct {
	Year         string `json:"year"`
	TotalSuccess int    `json:"total_success"`
	TotalAmount  int    `json:"total_amount"`
}

type MonthAmountTransactionFailedResponse struct {
	Year        string `json:"year"`
	Month       string `json:"month"`
	TotalFailed int    `json:"total_failed"`
	TotalAmount int    `json:"total_amount"`
}

type YearAmountTransactionFailedResponse struct {
	Year        string `json:"year"`
	TotalFailed int    `json:"total_failed"`
	TotalAmount int    `json:"total_amount"`
}

type MonthMethodTransactionResponse struct {
	Month             string `json:"month"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}

type YearMethodTransactionResponse struct {
	Year              string `json:"year"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}

type ApiResponsesTransactionMonthSuccess struct {
	Status  string                                   `json:"status"`
	Message string                                   `json:"message"`
	Data    []*MonthAmountTransactionSuccessResponse `json:"data"`
}

type ApiResponsesTransactionMonthFailed struct {
	Status  string                                  `json:"status"`
	Message string                                  `json:"message"`
	Data    []*MonthAmountTransactionFailedResponse `json:"data"`
}

type ApiResponsesTransactionYearSuccess struct {
	Status  string                                  `json:"status"`
	Message string                                  `json:"message"`
	Data    []*YearAmountTransactionSuccessResponse `json:"data"`
}

type ApiResponsesTransactionYearFailed struct {
	Status  string                                 `json:"status"`
	Message string                                 `json:"message"`
	Data    []*YearAmountTransactionFailedResponse `json:"data"`
}

type ApiResponsesTransactionMonthMethod struct {
	Status  string                            `json:"status"`
	Message string                            `json:"message"`
	Data    []*MonthMethodTransactionResponse `json:"data"`
}

type ApiResponsesTransactionYearMethod struct {
	Status  string                           `json:"status"`
	Message string                           `json:"message"`
	Data    []*YearMethodTransactionResponse `json:"data"`
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
