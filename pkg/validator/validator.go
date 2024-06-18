package validator

import "github.com/go-playground/validator/v10"

type Validator struct {
    validator *validator.Validate
}

func (v *Validator) Validate(data any) error {
    return v.validator.Struct(data)
}

func New() *Validator {
    return &Validator{validator.New()}
}
