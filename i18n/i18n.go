package i18n

import (
	"encoding/json"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle
var localizer *i18n.Localizer

func LoadFiles() error {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	_, err := bundle.LoadMessageFile("i18n/en.json")
	if err != nil {
		return err
	}

	localizer = i18n.NewLocalizer(bundle, "en")
	return nil
}

func Localize(id string, data interface{}) string {
	toReturn, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: id, TemplateData: data})
	return toReturn
}
