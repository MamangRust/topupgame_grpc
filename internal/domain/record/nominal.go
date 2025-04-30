package record

type NominalRecord struct {
	ID        int     `json:"id"`
	VoucherID int     `json:"voucher_id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}

type MonthAmountNominalSuccessRecord struct {
	ID           int    `json:"id"`
	NominalName  string `json:"nominal_name"`
	Year         string `json:"year"`
	Month        string `json:"month"`
	TotalSuccess int    `json:"total_success"`
	TotalAmount  int    `json:"total_amount"`
}

type YearAmountNominalSuccessRecord struct {
	ID           int    `json:"id"`
	NominalName  string `json:"nominal_name"`
	Year         string `json:"year"`
	TotalSuccess int    `json:"total_success"`
	TotalAmount  int    `json:"total_amount"`
}

type MonthAmountNominalFailedRecord struct {
	ID          int    `json:"id"`
	NominalName string `json:"nominal_name"`
	Year        string `json:"year"`
	Month       string `json:"month"`
	TotalFailed int    `json:"total_failed"`
	TotalAmount int    `json:"total_amount"`
}

type YearAmountNominalFailedRecord struct {
	ID          int    `json:"id"`
	NominalName string `json:"nominal_name"`
	Year        string `json:"year"`
	TotalFailed int    `json:"total_failed"`
	TotalAmount int    `json:"total_amount"`
}

type MonthMethodNominalRecord struct {
	ID                int    `json:"id"`
	Month             string `json:"month"`
	NominalName       string `json:"nominal_name"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}

type YearMethodNominalRecord struct {
	ID                int    `json:"id"`
	Year              string `json:"year"`
	NominalName       string `json:"nominal_name"`
	PaymentMethod     string `json:"payment_method"`
	TotalTransactions int    `json:"total_transactions"`
	TotalAmount       int    `json:"total_amount"`
}
