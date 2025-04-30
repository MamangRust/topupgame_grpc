package requests

import "github.com/go-playground/validator/v10"

type MonthAmountBankRequest struct {
	Year  int `json:"year" validate:"required"`
	Month int `json:"month" validate:"required"`
}

type MonthAmountBankByIdRequest struct {
	ID    int `json:"id" validate:"required,min=1"`
	Year  int `json:"year" validate:"required"`
	Month int `json:"month" validate:"required"`
}

type MonthAmountBankByMerchantRequest struct {
	MerchantID int `json:"merchant_id" validate:"required"`
	Year       int `json:"year" validate:"required"`
	Month      int `json:"month" validate:"required"`
}

type YearAmountBankByIdRequest struct {
	ID   int `json:"id" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}

type YearAmountBankByMerchantRequest struct {
	MerchantID int `json:"merchant_id" validate:"required"`
	Year       int `json:"year" validate:"required"`
}

type MonthMethodBankByIdRequest struct {
	ID   int `json:"id" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}

type YearMethodBankByIdRequest struct {
	ID   int `json:"id" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}

type MonthMethodBankByMerchantRequest struct {
	MerchantID int `json:"merchant_id" validate:"required"`
	Year       int `json:"year" validate:"required"`
}

type YearMethodBankByMerchantRequest struct {
	MerchantID int `json:"merchant_id" validate:"required"`
	Year       int `json:"year" validate:"required"`
}

type FindAllBanks struct {
	Search   string `json:"search" validate:"required"`
	Page     int    `json:"page" validate:"min=1"`
	PageSize int    `json:"page_size" validate:"min=1,max=100"`
}

type CreateBankRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateBankRequest struct {
	ID   int    `json:"id" validate:"required,min=1"`
	Name string `json:"name" validate:"omitempty"`
}

func (r *CreateBankRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}

func (r *UpdateBankRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}
