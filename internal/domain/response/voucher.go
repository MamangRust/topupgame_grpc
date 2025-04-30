package response

type VoucherResponse struct {
	ID         int     `json:"id"`
	MerchantID int     `json:"merchant_id"`
	CategoryID int     `json:"category_id"`
	Name       string  `json:"name"`
	ImageName  string  `json:"image_name"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	DeletedAt  *string `json:"deleted_at"`
}

type VoucherResponseDeleteAt struct {
	ID         int    `json:"id"`
	MerchantID int    `json:"merchant_id"`
	CategoryID int    `json:"category_id"`
	Name       string `json:"name"`
	ImageName  string `json:"image_name"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
}

type MonthAmountVoucherSuccessResponse struct {
	ID           int    `json:"id"`
	VoucherName  string `json:"voucher_name"`
	Year         string `json:"year"`
	Month        string `json:"month"`
	TotalSuccess int    `json:"total_success"`
	TotalAmount  int    `json:"total_amount"`
}

type YearAmountVoucherSuccessResponse struct {
	ID           int    `json:"id"`
	VoucherName  string `json:"voucher_name"`
	Year         string `json:"year"`
	TotalSuccess int    `json:"total_success"`
	TotalAmount  int    `json:"total_amount"`
}

type MonthAmountVoucherFailedResponse struct {
	ID          int    `json:"id"`
	VoucherName string `json:"voucher_name"`
	Year        string `json:"year"`
	Month       string `json:"month"`
	TotalFailed int    `json:"total_failed"`
	TotalAmount int    `json:"total_amount"`
}

type YearAmountVoucherFailedResponse struct {
	ID          int    `json:"id"`
	VoucherName string `json:"voucher_name"`
	Year        string `json:"year"`
	TotalFailed int    `json:"total_failed"`
	TotalAmount int    `json:"total_amount"`
}

type MonthMethodVoucherResponse struct {
	ID                int    `json:"id"`
	Month             string `json:"month"`
	VoucherName       string `json:"voucher_name"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}

type YearMethodVoucherResponse struct {
	Year              string `json:"year"`
	ID                int    `json:"id"`
	VoucherName       string `json:"voucher_name"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}

type ApiResponsesVoucherMonthSuccess struct {
	Status  string                               `json:"status"`
	Message string                               `json:"message"`
	Data    []*MonthAmountVoucherSuccessResponse `json:"data"`
}

type ApiResponsesVoucherMonthFailed struct {
	Status  string                              `json:"status"`
	Message string                              `json:"message"`
	Data    []*MonthAmountVoucherFailedResponse `json:"data"`
}

type ApiResponsesVoucherYearSuccess struct {
	Status  string                              `json:"status"`
	Message string                              `json:"message"`
	Data    []*YearAmountVoucherSuccessResponse `json:"data"`
}

type ApiResponsesVoucherYearFailed struct {
	Status  string                             `json:"status"`
	Message string                             `json:"message"`
	Data    []*YearAmountVoucherFailedResponse `json:"data"`
}

type ApiResponsesVoucherMonthMethod struct {
	Status  string                        `json:"status"`
	Message string                        `json:"message"`
	Data    []*MonthMethodVoucherResponse `json:"data"`
}

type ApiResponsesVoucherYearMethod struct {
	Status  string                       `json:"status"`
	Message string                       `json:"message"`
	Data    []*YearMethodVoucherResponse `json:"data"`
}

type ApiResponseVoucherAll struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ApiResponseVoucherDelete struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ApiResponseVoucher struct {
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Data    *VoucherResponse `json:"data"`
}

type ApiResponseVoucherDeleteAt struct {
	Status  string                   `json:"status"`
	Message string                   `json:"message"`
	Data    *VoucherResponseDeleteAt `json:"data"`
}

type ApiResponsesVoucher struct {
	Status  string             `json:"status"`
	Message string             `json:"message"`
	Data    []*VoucherResponse `json:"data"`
}

type ApiResponsePaginationVoucher struct {
	Status     string             `json:"status"`
	Message    string             `json:"message"`
	Data       []*VoucherResponse `json:"data"`
	Pagination *PaginationMeta    `json:"pagination"`
}

type ApiResponsePaginationVoucherDeleteAt struct {
	Status     string                     `json:"status"`
	Message    string                     `json:"message"`
	Data       []*VoucherResponseDeleteAt `json:"data"`
	Pagination *PaginationMeta            `json:"pagination"`
}
