package requests

import "github.com/go-playground/validator/v10"

type MonthAmountCategoryRequest struct {
	Year  int `json:"year" validate:"required"`
	Month int `json:"month" validate:"required"`
}

type MonthAmountCategoryByIdRequest struct {
	ID    int `json:"id" validate:"required,min=1"`
	Year  int `json:"year" validate:"required"`
	Month int `json:"month" validate:"required"`
}

type MonthAmountCategoryByMerchantRequest struct {
	MerchantID int `json:"merchant_id" validate:"required"`
	Year       int `json:"year" validate:"required"`
	Month      int `json:"month" validate:"required"`
}

type YearAmountCategoryByIdRequest struct {
	ID   int `json:"id" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}

type YearAmountCategoryByMerchantRequest struct {
	MerchantID int `json:"merchant_id" validate:"required"`
	Year       int `json:"year" validate:"required"`
}

type MonthMethodCategoryByIdRequest struct {
	ID   int `json:"id" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}

type YearMethodCategoryByIdRequest struct {
	ID   int `json:"id" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}

type MonthMethodCategoryByMerchantRequest struct {
	MerchantID int `json:"merchant_id" validate:"required"`
	Year       int `json:"year" validate:"required"`
}

type YearMethodCategoryByMerchantRequest struct {
	MerchantID int `json:"merchant_id" validate:"required"`
	Year       int `json:"year" validate:"required"`
}

type FindAllCategory struct {
	Search   string `json:"search" validate:"required"`
	Page     int    `json:"page" validate:"min=1"`
	PageSize int    `json:"page_size" validate:"min=1,max=100"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateCategoryRequest struct {
	ID   int    `json:"id" validate:"required,min=1"`
	Name string `json:"name" validate:"omitempty"`
}

func (r *CreateCategoryRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}

func (r *UpdateCategoryRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}
