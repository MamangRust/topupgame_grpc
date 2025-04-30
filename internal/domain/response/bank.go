package response

type BankResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type BankResponseDeleteAt struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type MonthAmountBankSuccessResponse struct {
	ID           int    `json:"id"`
	BankName     string `json:"bank_name"`
	Year         string `json:"year"`
	Month        string `json:"month"`
	TotalSuccess int    `json:"total_success"`
	TotalAmount  int    `json:"total_amount"`
}

type YearAmountBankSuccessResponse struct {
	ID           int    `json:"id"`
	BankName     string `json:"bank_name"`
	Year         string `json:"year"`
	TotalSuccess int    `json:"total_success"`
	TotalAmount  int    `json:"total_amount"`
}

type MonthAmountBankFailedResponse struct {
	ID          int    `json:"id"`
	BankName    string `json:"bank_name"`
	Year        string `json:"year"`
	Month       string `json:"month"`
	TotalFailed int    `json:"total_failed"`
	TotalAmount int    `json:"total_amount"`
}

type YearAmountBankFailedResponse struct {
	ID          int    `json:"id"`
	BankName    string `json:"bank_name"`
	Year        string `json:"year"`
	TotalFailed int    `json:"total_failed"`
	TotalAmount int    `json:"total_amount"`
}

type MonthMethodBankResponse struct {
	ID                int    `json:"id"`
	Month             string `json:"month"`
	BankName          string `json:"bank_name"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}

type YearMethodBankResponse struct {
	ID                int    `json:"id"`
	Year              string `json:"year"`
	BankName          string `json:"bank_name"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}

type ApiResponseBankMonthAmountSuccess struct {
	Status  string                            `json:"status"`
	Message string                            `json:"message"`
	Data    []*MonthAmountBankSuccessResponse `json:"data"`
}

type ApiResponseBankYearAmountSuccess struct {
	Status  string                           `json:"status"`
	Message string                           `json:"message"`
	Data    []*YearAmountBankSuccessResponse `json:"data"`
}

type ApiResponseBankMonthAmountFailed struct {
	Status  string                           `json:"status"`
	Message string                           `json:"message"`
	Data    []*MonthAmountBankFailedResponse `json:"data"`
}

type ApiResponseBankYearAmountFailed struct {
	Status  string                          `json:"status"`
	Message string                          `json:"message"`
	Data    []*YearAmountBankFailedResponse `json:"data"`
}

type ApiResponseBankMonthMethod struct {
	Status  string                     `json:"status"`
	Message string                     `json:"message"`
	Data    []*MonthMethodBankResponse `json:"data"`
}

type ApiResponseBankYearMethod struct {
	Status  string                    `json:"status"`
	Message string                    `json:"message"`
	Data    []*YearMethodBankResponse `json:"data"`
}

type ApiResponseBankAll struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ApiResponseBankDelete struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ApiResponseBank struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    *BankResponse `json:"data"`
}

type ApiResponseBankDeleteAt struct {
	Status  string                `json:"status"`
	Message string                `json:"message"`
	Data    *BankResponseDeleteAt `json:"data"`
}

type ApiResponsesBank struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    []*BankResponse `json:"data"`
}

type ApiResponsePaginationBank struct {
	Status     string          `json:"status"`
	Message    string          `json:"message"`
	Data       []*BankResponse `json:"data"`
	Pagination *PaginationMeta `json:"pagination"`
}

type ApiResponsePaginationBankDeleteAt struct {
	Status     string                  `json:"status"`
	Message    string                  `json:"message"`
	Data       []*BankResponseDeleteAt `json:"data"`
	Pagination *PaginationMeta         `json:"pagination"`
}
