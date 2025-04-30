package response

type CategoryResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CategoryResponseDeleteAt struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type MonthAmountCategorySuccessResponse struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
	Year         string `json:"year"`
	Month        string `json:"month"`
	TotalSuccess int    `json:"total_success"`
	TotalAmount  int    `json:"total_amount"`
}

type YearAmountCategorySuccessResponse struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
	Year         string `json:"year"`
	TotalSuccess int    `json:"total_success"`
	TotalAmount  int    `json:"total_amount"`
}

type MonthAmountCategoryFailedResponse struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
	Year         string `json:"year"`
	Month        string `json:"month"`
	TotalFailed  int    `json:"total_failed"`
	TotalAmount  int    `json:"total_amount"`
}

type YearAmountCategoryFailedResponse struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
	Year         string `json:"year"`
	TotalFailed  int    `json:"total_failed"`
	TotalAmount  int    `json:"total_amount"`
}

type MonthMethodCategoryResponse struct {
	ID                int    `json:"id"`
	Month             string `json:"month"`
	CategoryName      string `json:"category_name"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}

type YearMethodCategoryResponse struct {
	ID                int    `json:"id"`
	Year              string `json:"year"`
	CategoryName      string `json:"category_name"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}

type ApiResponseCategoryMonthAmountSuccess struct {
	Status  string                                `json:"status"`
	Message string                                `json:"message"`
	Data    []*MonthAmountCategorySuccessResponse `json:"data"`
}

type ApiResponseCategoryYearAmountSuccess struct {
	Status  string                               `json:"status"`
	Message string                               `json:"message"`
	Data    []*YearAmountCategorySuccessResponse `json:"data"`
}

type ApiResponseCategoryMonthAmountFailed struct {
	Status  string                               `json:"status"`
	Message string                               `json:"message"`
	Data    []*MonthAmountCategoryFailedResponse `json:"data"`
}

type ApiResponseCategoryYearAmountFailed struct {
	Status  string                              `json:"status"`
	Message string                              `json:"message"`
	Data    []*YearAmountCategoryFailedResponse `json:"data"`
}

type ApiResponseCategoryMonthMethod struct {
	Status  string                         `json:"status"`
	Message string                         `json:"message"`
	Data    []*MonthMethodCategoryResponse `json:"data"`
}

type ApiResponseCategoryYearMethod struct {
	Status  string                        `json:"status"`
	Message string                        `json:"message"`
	Data    []*YearMethodCategoryResponse `json:"data"`
}

type ApiResponseCategoryAll struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ApiResponseCategoryDelete struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ApiResponseCategory struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    *CategoryResponse `json:"data"`
}

type ApiResponseCategoryDeleteAt struct {
	Status  string                    `json:"status"`
	Message string                    `json:"message"`
	Data    *CategoryResponseDeleteAt `json:"data"`
}

type ApiResponsesCategory struct {
	Status  string              `json:"status"`
	Message string              `json:"message"`
	Data    []*CategoryResponse `json:"data"`
}

type ApiResponsePaginationCategory struct {
	Status     string              `json:"status"`
	Message    string              `json:"message"`
	Data       []*CategoryResponse `json:"data"`
	Pagination *PaginationMeta     `json:"pagination"`
}

type ApiResponsePaginationCategoryDeleteAt struct {
	Status     string                      `json:"status"`
	Message    string                      `json:"message"`
	Data       []*CategoryResponseDeleteAt `json:"data"`
	Pagination *PaginationMeta             `json:"pagination"`
}
