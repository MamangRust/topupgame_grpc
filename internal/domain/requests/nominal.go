package requests

import "github.com/go-playground/validator/v10"

type MonthAmountNominalRequest struct {
	Year  int `json:"year" validate:"required"`
	Month int `json:"month" validate:"required"`
}

type MonthAmountNominalByIdRequest struct {
	ID    int `json:"id" validate:"required,min=1"`
	Year  int `json:"year" validate:"required"`
	Month int `json:"month" validate:"required"`
}

type MonthAmountNominalByMerchantRequest struct {
	MerchantID int `json:"merchant_id" validate:"required"`
	Year       int `json:"year" validate:"required"`
	Month      int `json:"month" validate:"required"`
}

type YearAmountNominalByIdRequest struct {
	ID   int `json:"id" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}

type YearAmountNominalByMerchantRequest struct {
	MerchantID int `json:"merchant_id" validate:"required"`
	Year       int `json:"year" validate:"required"`
}

type MonthMethodNominalByIdRequest struct {
	ID   int `json:"id" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}

type YearMethodNominalByIdRequest struct {
	ID   int `json:"id" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}

type MonthMethodNominalByMerchantRequest struct {
	MerchantID int `json:"merchant_id" validate:"required"`
	Year       int `json:"year" validate:"required"`
}

type YearMethodNominalByMerchantRequest struct {
	MerchantID int `json:"merchant_id" validate:"required"`
	Year       int `json:"year" validate:"required"`
}

type FindAllNominals struct {
	Search   string `json:"search" validate:"required"`
	Page     int    `json:"page" validate:"min=1"`
	PageSize int    `json:"page_size" validate:"min=1,max=100"`
}

type CreateNominalRequest struct {
	VoucherID int     `json:"voucher_id" validate:"required,min=1"`
	Name      string  `json:"name" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required,min=1"`
	Price     float64 `json:"price" validate:"required,min=0"`
}

type UpdateNominalRequest struct {
	ID        int     `json:"id" validate:"required,min=1"`
	VoucherID int     `json:"voucher_id" validate:"omitempty,min=1"`
	Name      string  `json:"name" validate:"omitempty"`
	Quantity  int     `json:"quantity" validate:"omitempty,min=1"`
	Price     float64 `json:"price" validate:"omitempty,min=0"`
}

func (r *CreateNominalRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}

func (r *UpdateNominalRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}
