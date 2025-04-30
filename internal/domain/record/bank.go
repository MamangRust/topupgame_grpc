package record

type BankRecord struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}

type MonthAmountBankSuccessRecord struct {
	ID           int    `json:"id"`
	BankName     string `json:"bank_name"`
	Year         string `json:"year"`
	Month        string `json:"month"`
	TotalSuccess int    `json:"total_success"`
	TotalAmount  int    `json:"total_amount"`
}

type YearAmountBankSuccessRecord struct {
	ID           int    `json:"id"`
	BankName     string `json:"bank_name"`
	Year         string `json:"year"`
	TotalSuccess int    `json:"total_success"`
	TotalAmount  int    `json:"total_amount"`
}

type MonthAmountBankFailedRecord struct {
	ID          int    `json:"id"`
	BankName    string `json:"bank_name"`
	Year        string `json:"year"`
	Month       string `json:"month"`
	TotalFailed int    `json:"total_failed"`
	TotalAmount int    `json:"total_amount"`
}

type YearAmountBankFailedRecord struct {
	ID          int    `json:"id"`
	BankName    string `json:"bank_name"`
	Year        string `json:"year"`
	TotalFailed int    `json:"total_failed"`
	TotalAmount int    `json:"total_amount"`
}

type MonthMethodBankRecord struct {
	ID                int    `json:"id"`
	Month             string `json:"month"`
	BankName          string `json:"bank_name"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}

type YearMethodBankRecord struct {
	ID                int    `json:"id"`
	Year              string `json:"year"`
	BankName          string `json:"bank_name"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}
