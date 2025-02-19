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
