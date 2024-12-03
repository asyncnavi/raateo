package validate

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var validate *validator.Validate

var translator ut.Translator

func init() {
	validate := validator.New()

	var validators []*validator.Validate

	validators = append(validators, validate)

	if ginValidator, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validators = append(validators, ginValidator)
	}

	for _, validator := range validators {
		translator, _ = ut.New(en.New(), en.New()).GetTranslator("en")

		en_translations.RegisterDefaultTranslations(validator, translator)
		// Use JSON tag names for errors instead of Go struct names.
		validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("trans"), ",", 2)[0]
			if name != "-" && name != "" {
				return name
			}
			name = strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name != "-" && name != "" {
				return name
			}
			name = strings.SplitN(fld.Tag.Get("uri"), ",", 2)[0]
			if name != "-" && name != "" {
				return name
			}
			name = strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]
			if name != "-" && name != "" {
				return name
			}
			return ""
		})
	}
}

func ToFieldErrors(verrors validator.ValidationErrors) FieldErrors {
	var fields FieldErrors
	for _, verror := range verrors {
		field := FieldError{
			Field: verror.Field(),
			Error: verror.Translate(translator),
		}
		fields = append(fields, field)
	}
	return fields
}

// Check validates the provided model against it's declared tags.
func Check(val any) error {
	if err := validate.Struct(val); err != nil {
		// Use a type assertion to get the real error value.
		verrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}
		return ToFieldErrors(verrors)
	}
	return nil
}
