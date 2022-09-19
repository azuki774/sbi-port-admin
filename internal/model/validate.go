package model

import validation "github.com/go-ozzo/ozzo-validation"

func ValidateDate(date string) (err error) {
	return validation.Validate(date, validation.Date("20060102"))
}
