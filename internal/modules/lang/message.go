package lang

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	ErrTimeout     error
	ErrCancel      error
	ErrUnavailable error
)

func init() {
	InitLangResource(language.English.String())
}

func InitLangResource(lang string) {
	bundle := i18n.NewBundle(language.English)

	// Load lang resources
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile("en.json")
	bundle.MustLoadMessageFile("zh.json")

	loc := i18n.NewLocalizer(bundle, lang)

	hostUnavailableInfo := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "node_unreachable",
	})
	fmt.Println(hostUnavailableInfo)

	timeoutInfo := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "timeout",
	})
	fmt.Println(timeoutInfo)

	cancelInfo := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "manual_cancel",
	})
	fmt.Println(cancelInfo)

	ErrUnavailable = errors.New(hostUnavailableInfo)
	ErrTimeout = errors.New(timeoutInfo)
	ErrCancel = errors.New(cancelInfo)
}
