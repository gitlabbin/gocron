package lang

import (
	"encoding/json"
	"errors"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	_ "github.com/ouqiang/gocron/internal/lang/statik"
	"github.com/rakyll/statik/fs"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/language"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
)

var (
	ErrTimeout     error
	ErrCancel      error
	ErrUnavailable error
)

var msgMap = make(map[string]string)
var statikFS http.FileSystem

func init() {
	var err error
	statikFS, err = fs.New()
	if err != nil {
		log.Fatal(err)
	}
	InitLangResource(language.English.String())
}

func loadStatFs(path string) []byte {
	r, err := statikFS.Open(path)
	if err != nil {
		log.Fatal("Failed to read lang json file: %s err %v", path, err)
	}
	defer r.Close()
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	return contents
}

func listBundles() []os.FileInfo {
	r, err := statikFS.Open("/")
	if err != nil {
		log.Fatal("Failed to read lang json dir / err %v", err)
		return nil
	}
	defer r.Close()

	list, err := r.Readdir(-1)
	if err != nil {
		log.Errorf("error reading directory: %v", err)
		return nil
	} else {
		return list
	}
}

func Tr(key string) string {
	return msgMap[key]
}

func InitLangResource(lang string) {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	var keys []reflect.Value
	for _, jsonFile := range listBundles() {
		content := loadStatFs("/" + jsonFile.Name())
		if keys == nil {
			keys = loadAllKeys(content)
		}

		if _, err := bundle.ParseMessageFileBytes(content, jsonFile.Name()); err != nil {
			panic(err)
		}
	}
	loc := i18n.NewLocalizer(bundle, lang)

	for _, k := range keys {
		msg := loc.MustLocalize(&i18n.LocalizeConfig{
			MessageID: k.Interface().(string),
		})
		msgMap[k.String()] = msg
	}

	ErrUnavailable = errors.New(Tr("node_unreachable"))
	ErrTimeout = errors.New(Tr("timeout"))
	ErrCancel = errors.New(Tr("manual_cancel"))
}

func loadAllKeys(content []byte) []reflect.Value {
	var result map[string]interface{}
	if err := json.Unmarshal(content, &result); err != nil {
		panic(err)
	}
	return reflect.ValueOf(result).MapKeys()
}
