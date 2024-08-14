package validators

import (
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	twtranslations "github.com/go-playground/validator/v10/translations/zh_tw"
	"github.com/pkg/errors"
	"strings"
)

type ErrorFields map[string]string

var (
	Validator      *validator.Validate
	ValidatorTrans ut.Translator
)

func NewTransValidator() error {
	english := zh_Hant_TW.New()
	uni := ut.New(english, english)
	validate := validator.New(validator.WithRequiredStructEnabled())
	trans, _ := uni.GetTranslator("zh_tw")
	if err := twtranslations.RegisterDefaultTranslations(validate, trans); err != nil {
		return errors.Wrap(err, "NewTransValidator")
	}
	Validator = validate
	ValidatorTrans = trans
	return nil
}

func Validate(data interface{}) (ErrorFields, error) {
	err := Validator.Struct(data)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil, err
		}
		errs := err.(validator.ValidationErrors)
		errMap := make(map[string]string)
		for _, e := range errs {
			errMap[strings.ToLower(e.Field())] = e.Translate(ValidatorTrans)
		}
		return errMap, nil
	}
	return nil, nil
}
