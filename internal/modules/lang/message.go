package lang

import (
	"encoding/json"
	"errors"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"io/ioutil"
	"net/http"

	_ "github.com/ouqiang/gocron/internal/lang/statik"
	"github.com/rakyll/statik/fs"
	log "github.com/sirupsen/logrus"
)

var (
	ErrTimeout     error
	ErrCancel      error
	ErrUnavailable error
)

var (
	MsgFailedReadConf         string
	MsgSignalReceived         string
	MsgSignalTerminalEnd      string
	MsgExitAlready            string
	MsgSystemToExit           string
	MsgStopScheduler          string
	MsgUpgradeVersion         string
	MsgUpgradeVersionDone     string
	MsgSchedulerToInit        string
	MsgSchedulerInitErrorJobs string
	MsgSchedulerInitDone      string
	MsgLogFmt                 string
)

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

func InitLangResource(lang string) {
	bundle := i18n.NewBundle(language.English)

	// Load lang resources
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	//bundle.MustLoadMessageFile("en.json")
	//bundle.MustLoadMessageFile("zh.json")

	if _, err := bundle.ParseMessageFileBytes(loadStatFs("/en.json"), "en.json"); err != nil {
		panic(err)
	}
	if _, err := bundle.ParseMessageFileBytes(loadStatFs("/zh.json"), "zh.json"); err != nil {
		panic(err)
	}

	loc := i18n.NewLocalizer(bundle, lang)

	hostUnavailableInfo := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "node_unreachable",
	})

	timeoutInfo := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "timeout",
	})

	cancelInfo := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "manual_cancel",
	})

	ErrUnavailable = errors.New(hostUnavailableInfo)
	ErrTimeout = errors.New(timeoutInfo)
	ErrCancel = errors.New(cancelInfo)

	msgFailedReadConf := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_failed_read_conf",
	})
	MsgFailedReadConf = msgFailedReadConf

	msgSignalReceived := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_signal_received",
	})
	MsgSignalReceived = msgSignalReceived

	msgSignalTerminalEnd := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_signal_terminal_end",
	})
	MsgSignalTerminalEnd = msgSignalTerminalEnd

	msgExitAlready := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_exit_already",
	})
	MsgExitAlready = msgExitAlready

	msgSystemToExit := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_system_to_exit",
	})
	MsgSystemToExit = msgSystemToExit

	msgStopScheduler := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_stop_scheduler",
	})
	MsgStopScheduler = msgStopScheduler

	msgUpgradeVersion := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_upgrade_version",
	})
	MsgUpgradeVersion = msgUpgradeVersion

	msgUpgradeVersionDone := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_upgrade_version_done",
	})
	MsgUpgradeVersionDone = msgUpgradeVersionDone

	msgSchedulerToInit := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_scheduler_to_init",
	})
	MsgSchedulerToInit = msgSchedulerToInit

	msgSchedulerInitErrorJobs := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_scheduler_init_error_jobs",
	})
	MsgSchedulerInitErrorJobs = msgSchedulerInitErrorJobs

	msgSchedulerInitDone := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_scheduler_init_done",
	})
	MsgSchedulerInitDone = msgSchedulerInitDone

	msgLogFmt := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_log_fmt",
	})
	MsgLogFmt = msgLogFmt
}
