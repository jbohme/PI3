package validation

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
	"github.com/jbohme/crud/configs/rest_err"
)

var (
	Validate   = validator.New()
	translator ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		translator, _ = unt.GetTranslator("en")
		err := enTranslation.RegisterDefaultTranslations(val, translator)
		if err != nil {
			return
		}
	}
}

func ValidateUserError(
	validatorErr error,
) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validatorErr, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid filed type")
	} else if errors.As(validatorErr, &jsonValidationError) {
		var errorsCauses []rest_err.Causes

		for _, e := range validatorErr.(validator.ValidationErrors) {

			cause := rest_err.Causes{
				Message: e.Translate(translator),
				Field:   e.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}

		fmt.Printf("%s", errorsCauses)

		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}
