package requests

import "github.com/go-playground/validator/v10"

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
