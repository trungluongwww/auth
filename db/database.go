package db

import (
	"fmt"
	"github.com/trungluongwww/auth/config"
	"net/url"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/pressly/goose"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	maxIdleConnections = 30
	maxOpenConnections = 30
	lifetimeRangeMax   = 65
)

var GormConfig = &gorm.Config{
	NamingStrategy: &schema.NamingStrategy{
		SingularTable: true,
	},
}

func NewDB(env config.Env) (*gorm.DB, error) {
	return newDB(env, dsn(env))
}

func newDB(env config.Env, dsn string) (*gorm.DB, error) {
	fmt.Println(dsn)
	dialect := mysql.Open(dsn)
	gdb, err := gorm.Open(dialect, GormConfig)
	if err != nil {
		log.Infof("Failed Connecting to database : %s", err)
		return nil, err
	}
	db, err := gdb.DB()
	if err != nil {
		log.Infof("Failed get sql.DB : %s", err)
		return nil, err
	}
	db.SetMaxIdleConns(maxIdleConnections)
	db.SetMaxOpenConns(maxOpenConnections)
	db.SetConnMaxLifetime(lifetimeRangeMax * time.Minute)

	if env.IsLocal() {
		gdb.Logger = gdb.Logger.LogMode(logger.Info)
	}

	err = Migrate(gdb)
	if err != nil {
		log.Infof("Failed migration ddl : %s", err)
		panic(err)
	}
	return gdb, nil
}

func Migrate(gdb *gorm.DB) error {
	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}
	db, err := gdb.DB()
	if err != nil {
		log.Infof("Failed get sql.DB : %s", err)
		return err
	}
	return goose.Up(db, "./mysql/ddl")
}

func dsn(env config.Env) string {
	var (
		user      = env.MysqlUser
		password  = env.MysqlPassword
		protocol  = env.MysqlProtocol
		dbName    = env.MysqlDatabase
		charset   = "utf8mb4"
		parseTime = "true"
		loc       = url.PathEscape("UTC") // default timezone is UTC
	)
	// TODO: SSL
	return fmt.Sprintf("%s:%s@%s/%s?charset=%s&parseTime=%s&loc=%s&sql_safe_updates=1&tls=skip-verify", user, password, protocol, dbName, charset, parseTime, loc)
}
