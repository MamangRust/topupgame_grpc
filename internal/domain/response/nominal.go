package response

type NominalResponse struct {
	ID        int     `json:"id"`
	VoucherID int     `json:"voucher_id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}

type NominalResponseDeleteAt struct {
	ID        int     `json:"id"`
	VoucherID int     `json:"voucher_id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt string  `json:"deleted_at"`
}

type MonthAmountNominalSuccessResponse struct {
	ID           int    `json:"id"`
	NominalName  string `json:"nominal_name"`
	Year         string `json:"mnc_year"`
	Month        string `json:"month"`
	TotalSuccess int    `json:"total_success"`
	TotalAmount  int    `json:"total_amount"`
}

type YearAmountNominalSuccessResponse struct {
	ID           int    `json:"id"`
	NominalName  string `json:"nominal_name"`
	Year         string `json:"year"`
	TotalSuccess int    `json:"total_success"`
	TotalAmount  int    `json:"total_amount"`
}

type MonthAmountNominalFailedResponse struct {
	ID          int    `json:"id"`
	NominalName string `json:"nominal_name"`
	Year        string `json:"year"`
	Month       string `json:"month"`
	TotalFailed int    `json:"total_failed"`
	TotalAmount int    `json:"total_amount"`
}

type YearAmountNominalFailedResponse struct {
	ID          int    `json:"id"`
	NominalName string `json:"nominal_name"`
	Year        string `json:"year"`
	TotalFailed int    `json:"total_failed"`
	TotalAmount int    `json:"total_amount"`
}

type MonthMethodNominalResponse struct {
	ID                int    `json:"id"`
	Month             string `json:"month"`
	NominalName       string `json:"nominal_name"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}

type YearMethodNominalResponse struct {
	ID                int    `json:"id"`
	Year              string `json:"year"`
	NominalName       string `json:"nominal_name"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}

type ApiResponseNominalMonthAmountSuccess struct {
	Status  string                               `json:"status"`
	Message string                               `json:"message"`
	Data    []*MonthAmountNominalSuccessResponse `json:"data"`
}

type ApiResponseNominalYearAmountSuccess struct {
	Status  string                              `json:"status"`
	Message string                              `json:"message"`
	Data    []*YearAmountNominalSuccessResponse `json:"data"`
}

type ApiResponseNominalMonthAmountFailed struct {
	Status  string                              `json:"status"`
	Message string                              `json:"message"`
	Data    []*MonthAmountNominalFailedResponse `json:"data"`
}

type ApiResponseNominalYearAmountFailed struct {
	Status  string                             `json:"status"`
	Message string                             `json:"message"`
	Data    []*YearAmountNominalFailedResponse `json:"data"`
}

type ApiResponseNominalMonthMethod struct {
	Status  string                        `json:"status"`
	Message string                        `json:"message"`
	Data    []*MonthMethodNominalResponse `json:"data"`
}

type ApiResponseNominalYearMethod struct {
	Status  string                       `json:"status"`
	Message string                       `json:"message"`
	Data    []*YearMethodNominalResponse `json:"data"`
}

type ApiResponseNominalAll struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ApiResponseNominalDelete struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ApiResponseNominal struct {
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Data    *NominalResponse `json:"data"`
}

type ApiResponseNominalDeleteAt struct {
	Status  string                   `json:"status"`
	Message string                   `json:"message"`
	Data    *NominalResponseDeleteAt `json:"data"`
}

type ApiResponsesNominal struct {
	Status  string             `json:"status"`
	Message string             `json:"message"`
	Data    []*NominalResponse `json:"data"`
}

type ApiResponsePaginationNominal struct {
	Status     string             `json:"status"`
	Message    string             `json:"message"`
	Data       []*NominalResponse `json:"data"`
	Pagination *PaginationMeta    `json:"pagination"`
}

type ApiResponsePaginationNominalDeleteAt struct {
	Status     string                     `json:"status"`
	Message    string                     `json:"message"`
	Data       []*NominalResponseDeleteAt `json:"data"`
	Pagination *PaginationMeta            `json:"pagination"`
}
