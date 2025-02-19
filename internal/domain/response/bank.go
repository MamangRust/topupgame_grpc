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
