package validator

import (
	"log"
	"strings"
	"sync"

	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	es_translation "github.com/go-playground/validator/v10/translations/es"
)

var once sync.Once

var validate *validator.Validate
var trans ut.Translator

func NewValidator() *validator.Validate {
	once.Do(func() {
		validate = validator.New()
		es := es.New()
		uni := ut.New(es, es)

		trans, found := uni.GetTranslator("es")
		if !found {
			log.Fatal("Traductor no encontrado")
		}
		if err := es_translation.RegisterDefaultTranslations(validate, trans); err != nil {
			log.Fatal(err)
		}
	})
	return validate
}

type Error struct {
	Key     string
	Message string
}

func GenerateMessage(v *validator.Validate, errs validator.ValidationErrors) []Error {

	var errors []Error

	for _, e := range errs {
		error := Error{
			Key:     strings.Split(e.Translate(trans), " ")[1],
			Message: e.Translate((trans)),
		}
		errors = append(errors, error)
	}

	return errors

}
