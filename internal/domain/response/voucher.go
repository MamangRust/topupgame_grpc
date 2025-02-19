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
