package db

import (
	"farmme-api/config"
	"context"
	"log"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const db_host = "mongodb://localhost:27017"
const db_user = "jenni"
const db_pass = "password"

func GetDBCollection() (*mongo.Database, error) {
	
	clientOptions := options.Client().SetAuth(options.Credential{
		AuthSource: "admin", Username: db_user,
		Password: db_pass, PasswordSet: true,
	}).ApplyURI(db_host)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Print("can't connect database!!!")
		return nil, err
	}
	db := client.Database("farmme")
	return db, nil
}

func connectDB(ctx context.Context) (*mongo.Database, error) {
	uri := fmt.Sprintf(db_host, db_user, db_pass, db_host, config.DB_NAME)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to mongo: %v", err)
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("mongo client couldn't connect with background context: %v", err)
	}
	db := client.Database(config.DB_NAME)
	return db, nil
}