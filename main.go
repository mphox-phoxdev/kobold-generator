package main

import (
	"flag"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/juju/errors"
	"github.com/mphox-phoxdev/kobold-generator/configs"
	"github.com/mphox-phoxdev/kobold-generator/core"
	"github.com/mphox-phoxdev/kobold-generator/handlers"
	"github.com/mphox-phoxdev/kobold-generator/kobold"
	"github.com/mphox-phoxdev/kobold-generator/mysql/kobolddb"
	log "github.com/sirupsen/logrus"
)

var err error

func main() {
	var handler handlers.Handler
	var config configs.Config
	var err error

	// read in command line flags
	portPtr := flag.String("port", "", "port number without semicolon")
	localVerboseOutputPtr := flag.Bool("lv", false, "output to local console without json format")
	configLocation := flag.String("config", "configs/conf.json", "the location of the config file loaded at runtime, defaults to configs/conf.json")
	flag.Parse()

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	if !*localVerboseOutputPtr {
		log.SetFormatter(&log.JSONFormatter{})
		logpath := filepath.Join(".", "logs")
		err := os.MkdirAll(logpath, os.ModePerm)
		if err != nil {
			core.LogErrorWithMessageAndStack(err)
			os.Exit(1)
		}

		logFile, err := os.OpenFile("logs/kobolds.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			core.LogErrorWithMessageAndStack(err)
			os.Exit(1)
		}
		defer logFile.Close()
		log.SetOutput(logFile)
	}

	err = config.LoadConfigurationFile(*configLocation)
	if err != nil {
		core.LogErrorWithMessageAndStack(err)
		os.Exit(1)
	}

	// Override port based on the incoming flat
	if *portPtr != "" {
		config.Port = *portPtr
	}

	switch logLevel := config.LogLevel; logLevel {
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "PANIC":
		log.SetLevel(log.PanicLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
		err = errors.Errorf("Log level set to unrecognized value \"%s\"", logLevel)
		core.LogErrorWithMessageAndStack(err)
	}

	// Create an sql.DB and check for errors
	log.Info("Establishing database connection")
	InitializeDatabaseConnection(config)
	// sql.DB should be long lived "defer" closes it once this function ends
	defer db.Close()

	// Regularly ping the database and try to re establish a connection on failure
	go TestDatabaseConnection(config)

	handler.KoboldDB = kobolddb.New(db)
	err = kobold.InitializeSkillMap(handler.KoboldDB)
	if err != nil {
		core.LogErrorWithMessageAndStack(err)
		os.Exit(1)
	}

	err = kobold.InitializeRoleSkillMap(handler.KoboldDB)
	if err != nil {
		core.LogErrorWithMessageAndStack(err)
		os.Exit(1)
	}

	err = kobold.InitializeRandomSkillMap(handler.KoboldDB)
	if err != nil {
		core.LogErrorWithMessageAndStack(err)
		os.Exit(1)
	}

	// handler.SkillDB = userdb.New(db)

	// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := gorillaHandlers.AllowedOrigins(config.AllowedHosts)
	credentialsOk := gorillaHandlers.AllowCredentials()
	methodsOk := gorillaHandlers.AllowedMethods([]string{"GET", "READ", "POST", "PUT", "OPTIONS"})
	allowedHeaders := gorillaHandlers.AllowedHeaders([]string{"Content-Type", "application/json"})

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/random-kobold", handler.RandomKoboldHandler).Methods("GET")

	// start server listen
	log.Infof("Running on port %s", config.Port)
	log.Fatal(http.ListenAndServe(
		":"+config.Port,
		gorillaHandlers.CORS(
			originsOk,
			credentialsOk,
			methodsOk,
			allowedHeaders,
		)(r)))
}
