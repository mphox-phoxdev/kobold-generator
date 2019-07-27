package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juju/errors"
	"github.com/mphox-phoxdev/kobold-generator/configs"
	"github.com/mphox-phoxdev/kobold-generator/core"
	log "github.com/sirupsen/logrus"
)

var db *sql.DB
var databaseIsUp bool

func InitializeDatabaseConnection(conf configs.Config) {
	db, err = sql.Open("mysql", buildDSNFromConfig(conf))
	if err != nil {
		err = errors.Trace(err)
		core.LogErrorWithMessageAndStack(err)
		os.Exit(1)
	}
}

func TestDatabaseConnection(config configs.Config) {
	for true {
		if databaseIsUp {
			log.Debug("Testing database connection: started")
		} else {
			log.Info("Testing database connection: started")
		}
		databaseError := db.Ping()
		if databaseError != nil {
			databaseIsUp = false
			databaseError = errors.Trace(databaseError)
			log.Error("Database connection fail")
			core.LogErrorWithMessageAndStack(databaseError)
			InitializeDatabaseConnection(config)
		} else {
			databaseIsUp = true
			log.Info("Database connection success")
		}
		if databaseIsUp {
			log.Debug("Testing database connection: completed")
			time.Sleep(1 * time.Minute)
		} else {
			log.Info("Testing database connection: completed")
			time.Sleep(5 * time.Second)
		}
	}
}

func TestDatabaseConnection2() (err error) {
	err = db.Ping()
	if err != nil {
		err = errors.Trace(err)
		return
	}
	return
}

func buildDSNFromConfig(conf configs.Config) string {
	databaseConfig := conf.Database
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		databaseConfig.UserName,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Name)
}
