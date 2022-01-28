package api

import (
	"github.com/GolfGrab/journey-to-complete-backend/util"
	"github.com/go-playground/validator/v10"
)

// validCurrency custom validation check for supported currency.
var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
