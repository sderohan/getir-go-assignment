package config

import (
	"context"
	"log"

	"github.com/sderohan/getir-go-assignment/pkg/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBAccessor struct {
	DB *mongo.Client
}

func (d *DBAccessor) GetDBConfig() *mongo.Client {
	return d.DB
}

func Init(connParam db.ConnectionURL) *mongo.Client {
	// Connect to the database
	clientDB, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connParam.String()))
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer func() {
		if err := clientDB.Disconnect(context.TODO()); err != nil {
			log.Fatalf(err.Error())
		}
	}()
	return clientDB
}
