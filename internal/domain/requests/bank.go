package requests

import "github.com/go-playground/validator/v10"

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
