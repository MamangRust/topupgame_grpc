package record

type CategoryRecord struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}

type MonthAmountCategorySuccessRecord struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
	Year         string `json:"year"`
	Month        string `json:"month"`
	TotalSuccess int    `json:"total_success"`
	TotalAmount  int    `json:"total_amount"`
}

type YearAmountCategorySuccessRecord struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
	Year         string `json:"year"`
	TotalSuccess int    `json:"total_success"`
	TotalAmount  int    `json:"total_amount"`
}

type MonthAmountCategoryFailedRecord struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
	Year         string `json:"year"`
	Month        string `json:"month"`
	TotalFailed  int    `json:"total_failed"`
	TotalAmount  int    `json:"total_amount"`
}

type YearAmountCategoryFailedRecord struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
	Year         string `json:"year"`
	TotalFailed  int    `json:"total_failed"`
	TotalAmount  int    `json:"total_amount"`
}

type MonthMethodCategoryRecord struct {
	ID                int    `json:"id"`
	Month             string `json:"month"`
	CategoryName      string `json:"category_name"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}

type YearMethodCategoryRecord struct {
	ID                int    `json:"id"`
	Year              string `json:"year"`
	CategoryName      string `json:"category_name"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}
