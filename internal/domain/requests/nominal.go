package requests

import "github.com/go-playground/validator/v10"

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
