package http

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en2 "github.com/go-playground/validator/translations/en"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"gopkg.in/go-playground/validator.v9"
)

func fields(errors validator.ValidationErrors, trans ut.Translator) FieldValidations {
	l := len(errors)
	if l > 0 {

		fields := make(FieldValidations, l)
		for _, e := range errors {
			fields[e.Field()] = e.Translate(trans)
		}

		return fields
	}
	return nil
}

// ValidateStruct validates struct based on their tags
func ValidateStruct(s interface{}) error {
	v, trans := newValidator()
	err := v.Struct(s)
	if err != nil {
		errPtr := malformedRequestErr(err.(validator.ValidationErrors), trans)
		return &errPtr
	}
	return nil
}

func malformedRequestErr(err validator.ValidationErrors, trans ut.Translator) ValidationError {
	return ValidationError{
		Code:    400,
		Message: "Malformed request.",
		Fields:  fields(err, trans),
	}
}

// Validator returns an instance of "gopkg.in/go-playground/validator.v9"
func newValidator() (*validator.Validate, ut.Translator) {

	en := en.New()
	uni := ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	v := validator.New()

	en2.RegisterDefaultTranslations(v, trans)

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	_ = v.RegisterValidation("objectId", func(fl validator.FieldLevel) bool {
		if fl.Field().String() == "" {
			return true
		}
		_, err := primitive.ObjectIDFromHex(fl.Field().String())
		return err == nil
	})

	_ = v.RegisterTranslation("objectId", trans, func(ut ut.Translator) error {
		return ut.Add("objectId", "{0} must have a valid primitive objectId value", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("objectId", fe.Field())

		return t
	})

	return v, trans
}
