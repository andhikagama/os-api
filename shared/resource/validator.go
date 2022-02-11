package resource

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

// Validator .
type Validator struct {
	Validator   *validator.Validate
	Translators []*ut.Translator
}

// NewValidator .
func NewValidator() *Validator {
	val := new(Validator)
	val.Validator = validator.New()
	en := en.New()
	id := id.New()
	uni := ut.New(en, en, id)

	transEN, _ := uni.GetTranslator("en")
	transID, _ := uni.GetTranslator("id")
	availTrans := make([]*ut.Translator, 2)
	availTrans[0] = &transEN
	availTrans[1] = &transID

	val.Translators = availTrans

	val.RegisterTranslations()

	return val
}

// Validate .
func (v *Validator) Validate(i interface{}) error {
	return v.Validator.Struct(i)
}

// RegisterTranslations .
func (v *Validator) RegisterTranslations() {
	en_translations.RegisterDefaultTranslations(v.Validator, *v.Translators[0])
	id_translations.RegisterDefaultTranslations(v.Validator, *v.Translators[1])
}
