package models

import "go.mongodb.org/mongo-driver/mongo"

type DBWriter struct {
	DB *mongo.Client
}

func (d *DBWriter) ReadCollections() {

}
