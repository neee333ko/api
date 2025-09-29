package v1

import (
	"github.com/neee333ko/component-base/pkg/validation"
	"github.com/neee333ko/component-base/pkg/validation/field"
)

func (u *User) ValidateCreate() field.ErrorList {
	v := validation.NewValidator(u)
	err := v.Validate()

	if result := validation.IsValidPassword(u.Password); len(result) != 0 {
		err = append(err, field.Invalid(field.NewPath("password"), result, ""))
	}

	return err
}

func (u *User) ValidateUpdate() field.ErrorList {
	v := validation.NewValidator(u)
	err := v.Validate()

	return err
}

func (s *Secret) Validate() field.ErrorList {
	v := validation.NewValidator(s)
	err := v.Validate()

	return err
}

func (p *Policy) Validate() field.ErrorList {
	v := validation.NewValidator(p)
	err := v.Validate()

	return err
}
