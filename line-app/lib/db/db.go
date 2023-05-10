package db

import (
	"io/ioutil"
	"time"

	"github.com/shota-imoto/line-app/lib/models/app_user"
	"github.com/shota-imoto/line-app/lib/models/line_model"
	"github.com/shota-imoto/line-app/lib/utils/app_env"
	"gopkg.in/yaml.v2"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

// mysqlの接続
type ConnectOpenConfig struct {
	Interval int
	Retry    int
	Count    int
}

func (config ConnectOpenConfig) intervalSleep() {
	time.Sleep(time.Second * time.Duration(config.Interval))
}

func init() {
	config := ConnectOpenConfig{Interval: 5, Retry: 20, Count: 0}
	var err error

	for {
		config.intervalSleep()

		Db, err = gorm.Open(mysql.Open(dsn()), &gorm.Config{})
		if err == nil {
			break
		} else {
			if config.Count > config.Retry {
				panic(err)
			}
			config.Count++
		}
	}

	Db.AutoMigrate(
		&app_user.User{},
		&line_model.LineGroup{},
		&line_model.LineGroupUserMap{},
	)

	if app_env.LineAppEnv == "test" {
		Db = Db.Begin()
	}
}

type DbConfig struct {
	DbName string `yaml:"db_name"`
}

func (config DbConfig) dsn() string {
	return "root@tcp(db:3306)/" + config.DbName + "?charset=utf8mb4&parseTime=true"
}

func dsn() string {
	var err error
	buf, err := ioutil.ReadFile(app_env.RootPath + "/lib/config/db/" + app_env.LineAppEnv + ".yml")
	if err != nil {
		panic(err)
	}

	db_config := DbConfig{}
	err = yaml.Unmarshal(buf, &db_config)

	if err != nil {
		panic(err)
	}

	return db_config.dsn()
}
