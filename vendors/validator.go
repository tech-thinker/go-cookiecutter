package vendors

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ModelValidator interface {
	Struct(doc interface{}) []error
}

type modelValidator struct {
	validator *validator.Validate
}

func (v *modelValidator) Struct(doc interface{}) []error {
	var res []error
	err := v.validator.Struct(doc)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			res = append(res, fmt.Errorf(`'%v' is not a valid %v`, e.Value(), e.StructField()))
		}
	}
	return res
}

func NewModelValidator(
	validator *validator.Validate,
) ModelValidator {
	return &modelValidator{
		validator: validator,
	}
}
