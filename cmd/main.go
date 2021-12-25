package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sderohan/getir-go-assignment/app/config"
	"github.com/sderohan/getir-go-assignment/app/controller"
	"github.com/sderohan/getir-go-assignment/app/models"
	"github.com/sderohan/getir-go-assignment/app/routes.go"
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
	rts := &routes.Router{
		&controller.Controller{
			&models.DBWriter{DB: db},
		},
	}

	// read server configuraion
	serverConfig := appConfig.GetAppConfig().Server

	// start server and listen for requests
	if err := http.ListenAndServe(fmt.Sprintf(":%s", serverConfig.Port), rts.InitRoutes()); err != nil {
		log.Fatalf(err.Error())
	}

	// Test database
	// collection := clientDB.Database(connParam.Database).Collection("records")
	// cursor, err := collection.Find(context.TODO(), bson.D{})
	// if err != nil {
	// 	panic(err)
	// }

	// var results []bson.M
	// if err = cursor.All(context.TODO(), &results); err != nil {
	// 	log.Fatal(err)
	// }
	// for _, result := range results {
	// 	fmt.Println(result)
	// }
}

// mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getir-case-study?retryWrites=true
