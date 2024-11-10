package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate *validator.Validate

func InitValidation() {
	validate = validator.New()
}

func ValidateStruct(w http.ResponseWriter, data interface{}) error {
	err := validate.Struct(data)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			http.Error(w, fmt.Sprintf("Validation error: %v", validationErrors), http.StatusBadRequest)
			return err
		}
	}
	return nil
}
