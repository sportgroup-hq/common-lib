package validation

import "github.com/go-playground/validator/v10"

func Register(v *validator.Validate) {
	v.RegisterAlias("password", "required,min=8,max=32")
}
