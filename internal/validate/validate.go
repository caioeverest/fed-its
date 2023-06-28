package validate

import "github.com/go-playground/validator/v10"

type Validate struct {
	*validator.Validate
}

func New() *Validate {
	return &Validate{validator.New()}
}
