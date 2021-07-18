package install

import (
	"errors"
	"fmt"
	"github.com/ouqiang/gocron/internal/lang"
	"strconv"

	macaron "gopkg.in/macaron.v1"

	"github.com/go-macaron/binding"
	"github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
	"github.com/ouqiang/gocron/internal/models"
	"github.com/ouqiang/gocron/internal/modules/app"
	"github.com/ouqiang/gocron/internal/modules/setting"
	"github.com/ouqiang/gocron/internal/modules/utils"
	"github.com/ouqiang/gocron/internal/service"
	log "github.com/sirupsen/logrus"
)

// System install

type InstallForm struct {
	DbType               string `binding:"In(mysql,postgres,sqlite3)"`
	DbHost               string `binding:"Required;MaxSize(50)"`
	DbPort               int    `binding:"Required;Range(1,65535)"`
	DbUsername           string `binding:"Required;MaxSize(50)"`
	DbPassword           string `binding:"Required;MaxSize(30)"`
	DbName               string `binding:"Required;MaxSize(50)"`
	DbTablePrefix        string `binding:"MaxSize(20)"`
	AdminUsername        string `binding:"Required;MinSize(3)"`
	AdminPassword        string `binding:"Required;MinSize(6)"`
	ConfirmAdminPassword string `binding:"Required;MinSize(6)"`
	AdminEmail           string `binding:"Required;Email;MaxSize(50)"`
}

func (f InstallForm) Error(ctx *macaron.Context, errs binding.Errors) {
	if len(errs) == 0 {
		return
	}
	json := utils.JsonResponse{}
	content := json.CommonFailure(lang.Tr("web_form_validate_fail"))
	ctx.Write([]byte(content))
}

// Installation
func Store(ctx *macaron.Context, form InstallForm) string {
	log.Infof("Install form: %v", form)
	json := utils.JsonResponse{}
	if app.Installed {
		return json.CommonFailure(lang.Tr("system_installed_already"))
	}
	if form.AdminPassword != form.ConfirmAdminPassword {
		return json.CommonFailure(lang.Tr("password_mismatch"))
	}
	err := testDbConnection(form)
	if err != nil {
		return json.CommonFailure(err.Error())
	}
	// 写入数据库配置
	err = writeConfig(form)
	if err != nil {
		return json.CommonFailure(lang.Tr("app_config_generate_failed"), err)
	}

	appConfig, err := setting.Read(app.AppConfig)
	if err != nil {
		return json.CommonFailure(lang.Tr("failed_read_app_ini"), err)
	}
	app.Setting = appConfig

	models.Db = models.CreateDb()
	// 创建数据库表
	migration := new(models.Migration)
	err = migration.Install(form.DbName)
	if err != nil {
		return json.CommonFailure(fmt.Sprintf(lang.Tr("failed_create_db_table"), err.Error()), err)
	}

	// 创建管理员账号
	err = createAdminUser(form)
	if err != nil {
		return json.CommonFailure(lang.Tr("failed_create_admin"), err)
	}

	// 创建安装锁
	err = app.CreateInstallLock()
	if err != nil {
		return json.CommonFailure(lang.Tr("failed_create_install_lock"), err)
	}

	// 更新版本号文件
	app.UpdateVersionFile()

	app.Installed = true
	// 初始化定时任务
	service.ServiceTask.Initialize()

	return json.Success(lang.Tr("app_installed"), nil)
}

// Write config file
func writeConfig(form InstallForm) error {
	dbConfig := []string{
		"db.engine", form.DbType,
		"db.host", form.DbHost,
		"db.port", strconv.Itoa(form.DbPort),
		"db.user", form.DbUsername,
		"db.password", form.DbPassword,
		"db.database", form.DbName,
		"db.prefix", form.DbTablePrefix,
		"db.charset", "utf8",
		"db.max.idle.conns", "5",
		"db.max.open.conns", "100",
		"allow_ips", "",
		"app.name", "Simple-GoCron", // Application name
		"api.key", "",
		"api.secret", "",
		"enable_tls", "false",
		"concurrency.queue", "500",
		"auth_secret", utils.RandAuthToken(),
		"ca_file", "",
		"cert_file", "",
		"key_file", "",
		"app.lang", "en",
	}

	return setting.Write(dbConfig, app.AppConfig)
}

// 创建管理员账号
func createAdminUser(form InstallForm) error {
	user := new(models.User)
	user.Name = form.AdminUsername
	user.Password = form.AdminPassword
	user.Email = form.AdminEmail
	user.IsAdmin = 1
	_, err := user.Create()

	return err
}

// 测试数据库连接
func testDbConnection(form InstallForm) error {
	var s setting.Setting
	s.Db.Engine = form.DbType
	s.Db.Host = form.DbHost
	s.Db.Port = form.DbPort
	s.Db.User = form.DbUsername
	s.Db.Password = form.DbPassword
	s.Db.Database = form.DbName
	s.Db.Charset = "utf8"
	db, err := models.CreateTmpDb(&s)
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Ping()
	if s.Db.Engine == "postgres" && err != nil {
		pgError, ok := err.(*pq.Error)
		if ok && pgError.Code == "3D000" {
			err = errors.New(lang.Tr("db_not_existed"))
		}
		return err
	}

	if s.Db.Engine == "mysql" && err != nil {
		mysqlError, ok := err.(*mysql.MySQLError)
		if ok && mysqlError.Number == 1049 {
			err = errors.New(lang.Tr("db_not_existed"))
		}
		return err
	}

	return err

}
