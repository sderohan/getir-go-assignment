package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/sderohan/getir-go-assignment/app/config"
	"github.com/sderohan/getir-go-assignment/app/controller"
	"github.com/sderohan/getir-go-assignment/app/models"
	"github.com/sderohan/getir-go-assignment/pkg/db"
)

func main() {

	if len(os.Args) < 2 || len(os.Args) > 2 {
		log.Fatalf("Invalid config filepath")
	}

	// read yml file from command line
	cfgFile := os.Args[1]

	// read database and server configuration from the yml file
	appConfig, err := config.NewAppConfig(cfgFile)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// get the db configuration
	dbConfig := appConfig.GetAppConfig().DB

	// db optional parameters
	dbOptions := make(map[string]string)
	dbOptions["retryWrites"] = "true"

	// form the connection object
	connParam := db.ConnectionURL{
		UserName: dbConfig.UserName,
		Password: dbConfig.Password,
		Host:     dbConfig.Host,
		Database: dbConfig.Database,
		Options:  dbOptions,
	}
	// create the mongodb connection
	db := config.Init(connParam)

	// instantiate app
	rts := &controller.Controller{
		&models.DBWriter{DB: db},
	}

	// read server configuraion
	serverConfig := appConfig.GetAppConfig().Server

	// request handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// read the request path
		url := strings.Split(html.EscapeString(r.URL.Path), "/")
		path := ""

		if len(url) > 0 {
			path = strings.Trim(url[1], " ")
		} else {
			http.Error(w, "Requested resource not found", http.StatusNotFound)
		}

		switch r.Method {

		case http.MethodGet:
			if path == "in-memory" {
				rts.PrintParams(w, r)
			} else {
				http.Error(w, "Requested resource not found", http.StatusNotFound)
			}

		case http.MethodPost:
			if path == "in-memory" {
				rts.EchoAPI(w, r)
			} else if path == "fetch-data" {
				rts.FetchData(w, r)
			} else {
				http.Error(w, "Requested resource not found", http.StatusNotFound)
			}

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// start server and listen for requests
	log.Printf("Server listening on port %s", serverConfig.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", serverConfig.Port), nil); err != nil {
		log.Fatalf(err.Error())
	}
}
