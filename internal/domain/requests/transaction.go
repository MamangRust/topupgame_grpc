package requests

import "github.com/go-playground/validator/v10"

type MonthAmountTransactionRequest struct {
	Year  int `json:"year" validate:"required"`
	Month int `json:"month" validate:"required"`
}

type MonthAmountTransactionByMerchantRequest struct {
	MerchantID int `json:"merchant_id" validate:"required"`
	Year       int `json:"year" validate:"required"`
	Month      int `json:"month" validate:"required"`
}

type YearAmountTransactionByMerchantRequest struct {
	MerchantID int `json:"merchant_id" validate:"required"`
	Year       int `json:"year" validate:"required"`
}

type MonthMethodTransactionByMerchantRequest struct {
	MerchantID int `json:"merchant_id" validate:"required"`
	Year       int `json:"year" validate:"required"`
}

type YearMethodTransactionByMerchantRequest struct {
	MerchantID int `json:"merchant_id" validate:"required"`
	Year       int `json:"year" validate:"required"`
}

type FindAllTransactions struct {
	Search   string `json:"search" validate:"required"`
	Page     int    `json:"page" validate:"min=1"`
	PageSize int    `json:"page_size" validate:"min=1,max=100"`
}

type CreateTransactionRequest struct {
	UserID        int    `json:"user_id" validate:"required,min=1"`
	MerchantID    int    `json:"merchant_id" validate:"required,min=1"`
	VoucherID     int    `json:"voucher_id" validate:"required,min=1"`
	NominalID     int    `json:"nominal_id" validate:"required,min=1"`
	CategoryID    int    `json:"category_id" validate:"required,min=1"`
	BankID        int    `json:"bank_id" validate:"required,min=1"`
	Amount        *int   `json:"amount"`
	PaymentMethod string `json:"payment_method" validate:"required"`
}

type UpdateTransactionRequest struct {
	ID            int     `json:"id" validate:"required,min=1"`
	UserID        int     `json:"user_id" validate:"omitempty,min=1"`
	MerchantID    int     `json:"merchant_id" validate:"omitempty,min=1"`
	VoucherID     int     `json:"voucher_id" validate:"omitempty,min=1"`
	NominalID     int     `json:"nominal_id" validate:"omitempty,min=1"`
	CategoryID    int     `json:"category_id" validate:"omitempty,min=1"`
	BankID        int     `json:"bank_id" validate:"omitempty,min=1"`
	Amount        *int    `json:"amount"`
	PaymentMethod string  `json:"payment_method" validate:"omitempty"`
	Status        *string `json:"status"`
}

func (r *CreateTransactionRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}

func (r *UpdateTransactionRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}
