package requests

import "github.com/go-playground/validator/v10"

type MonthAmountVoucherRequest struct {
	Year  int `json:"year" validate:"required"`
	Month int `json:"month" validate:"required"`
}

type MonthAmountVoucherByIdRequest struct {
	ID    int `json:"id" validate:"required,min=1"`
	Year  int `json:"year" validate:"required"`
	Month int `json:"month" validate:"required"`
}

type MonthAmountVoucherByMerchantRequest struct {
	MerchantID int `json:"merchant_id" validate:"required"`
	Year       int `json:"year" validate:"required"`
	Month      int `json:"month" validate:"required"`
}

type YearAmountVoucherByIdRequest struct {
	ID   int `json:"id" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}

type YearAmountVoucherByMerchantRequest struct {
	MerchantID int `json:"merchant_id" validate:"required"`
	Year       int `json:"year" validate:"required"`
}

type MonthMethodVoucherByIdRequest struct {
	ID   int `json:"id" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}

type YearMethodVoucherByIdRequest struct {
	ID   int `json:"id" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}

type MonthMethodVoucherByMerchantRequest struct {
	MerchantID int `json:"merchant_id" validate:"required"`
	Year       int `json:"year" validate:"required"`
}

type YearMethodVoucherByMerchantRequest struct {
	MerchantID int `json:"merchant_id" validate:"required"`
	Year       int `json:"year" validate:"required"`
}

type FindAllVouchers struct {
	Search   string `json:"search" validate:"required"`
	Page     int    `json:"page" validate:"min=1"`
	PageSize int    `json:"page_size" validate:"min=1,max=100"`
}

type VoucherFormData struct {
	MerchantID int    `json:"merchant_id"`
	CategoryID int    `json:"category_id"`
	Name       string `json:"name"`
	ImagePath  string `json:"image_path"`
}

type CreateVoucherRequest struct {
	MerchantID int    `json:"merchant_id" validate:"required,min=1"`
	CategoryID int    `json:"category_id" validate:"required,min=1"`
	Name       string `json:"name" validate:"required"`
	ImageName  string `json:"image_name" validate:"required"`
}

type UpdateVoucherRequest struct {
	ID         int    `json:"id" validate:"required,min=1"`
	MerchantID int    `json:"merchant_id" validate:"omitempty,min=1"`
	CategoryID int    `json:"category_id" validate:"omitempty,min=1"`
	Name       string `json:"name" validate:"omitempty"`
	ImageName  string `json:"image_name" validate:"omitempty"`
}

func (r *CreateVoucherRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}

func (r *UpdateVoucherRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}
