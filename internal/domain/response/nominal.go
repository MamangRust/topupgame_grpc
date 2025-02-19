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
