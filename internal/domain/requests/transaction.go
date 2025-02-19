package requests

import "github.com/go-playground/validator/v10"

type CreateTransactionRequest struct {
	UserID        int    `json:"user_id" validate:"required,min=1"`
	MerchantID    int    `json:"merchant_id" validate:"required,min=1"`
	VoucherID     int    `json:"voucher_id"`
	NominalID     int    `json:"nominal_id" validate:"required,min=1"`
	CategoryID    int    `json:"category_id" validate:"required,min=1"`
	BankID        int    `json:"bank_id" validate:"required,min=1"`
	PaymentMethod string `json:"payment_method" validate:"required"`
}

type UpdateTransactionRequest struct {
	ID            int    `json:"id" validate:"required,min=1"`
	UserID        int    `json:"user_id" validate:"omitempty,min=1"`
	MerchantID    int    `json:"merchant_id" validate:"omitempty,min=1"`
	VoucherID     int    `json:"voucher_id" validate:"omitempty,min=1"`
	NominalID     int    `json:"nominal_id" validate:"omitempty,min=1"`
	CategoryID    int    `json:"category_id" validate:"omitempty,min=1"`
	BankID        int    `json:"bank_id" validate:"omitempty,min=1"`
	PaymentMethod string `json:"payment_method" validate:"omitempty"`
	Status        string `json:"status"`
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
