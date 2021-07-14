package lang

import (
	"encoding/json"
	"errors"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"io/ioutil"
	"net/http"
	"os"

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
	MsgFailedReadConf               string
	MsgSignalReceived               string
	MsgSignalTerminalEnd            string
	MsgExitAlready                  string
	MsgSystemToExit                 string
	MsgStopScheduler                string
	MsgUpgradeVersion               string
	MsgUpgradeVersionDone           string
	MsgSchedulerToInit              string
	MsgSchedulerInitErrorJobs       string
	MsgSchedulerInitDone            string
	MsgLogFmt                       string
	MsgNotAllowSubtaskSchedule      string
	MsgNotSupportProtocolTaskToJob  string
	MsgFailedAddTaskToScheduler     string
	MsgHttpStatusNot200             string
	MsgNodeErrorMessage             string
	MsgJobStart                     string
	MsgJobDone                      string
	MsgJobStartFailOnLog            string
	MsgJobCmd                       string
	MsgJobEndUpdateLogFailed        string
	MsgParentFailedChildWouldNotRun string
	MsgFailedGetDependencyJob       string
	MsgDependencyJobEmpty           string
	MsgDependencyJob                string
	StatusFailed                    string
	StatusSucceed                   string
	MsgJobFailed                    string
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

func InitLangResource(lang string) {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	for _, json := range listBundles() {
		content := loadStatFs("/" + json.Name())
		if _, err := bundle.ParseMessageFileBytes(content, json.Name()); err != nil {
			panic(err)
		}
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

	msgNotAllowSubtaskSchedule := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_not_allow_subtask_schedule",
	})
	MsgNotAllowSubtaskSchedule = msgNotAllowSubtaskSchedule

	msgNotSupportProtocolTaskToJob := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_not_support_protocol_task_to_job",
	})
	MsgNotSupportProtocolTaskToJob = msgNotSupportProtocolTaskToJob

	msgFailedAddTaskToScheduler := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_failed_add_task_to_scheduler",
	})
	MsgFailedAddTaskToScheduler = msgFailedAddTaskToScheduler

	msgHttpStatusNot200 := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_http_status_not_200",
	})
	MsgHttpStatusNot200 = msgHttpStatusNot200

	msgNodeErrorMessage := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_node_error_message",
	})
	MsgNodeErrorMessage = msgNodeErrorMessage

	msgJobStart := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_job_start",
	})
	MsgJobStart = msgJobStart

	msgJobDone := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_job_done",
	})
	MsgJobDone = msgJobDone

	msgJobStartFailOnLog := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_job_start_fail_on_log",
	})
	MsgJobStartFailOnLog = msgJobStartFailOnLog

	msgJobCmd := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_job_cmd",
	})
	MsgJobCmd = msgJobCmd

	msgJobEndUpdateLogFailed := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_job_end_update_log_failed",
	})
	MsgJobEndUpdateLogFailed = msgJobEndUpdateLogFailed

	msgParentFailedChildWouldNotRun := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_parent_failed_child_would_not_run",
	})
	MsgParentFailedChildWouldNotRun = msgParentFailedChildWouldNotRun

	msgFailedGetDependencyJob := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_failed_get_dependency_job",
	})
	MsgFailedGetDependencyJob = msgFailedGetDependencyJob

	msgDependencyJobEmpty := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_dependency_job_empty",
	})
	MsgDependencyJobEmpty = msgDependencyJobEmpty

	msgDependencyJob := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_dependency_job",
	})
	MsgDependencyJob = msgDependencyJob

	statusFailed := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "status_failed",
	})
	StatusFailed = statusFailed

	statusSucceed := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "status_succeed",
	})
	StatusSucceed = statusSucceed

	msgJobFailed := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "msg_job_failed",
	})
	MsgJobFailed = msgJobFailed

}
