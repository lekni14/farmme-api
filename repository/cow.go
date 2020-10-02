package repository

import (
	"context"
	"fmt"
	"log"
	// "strings"
	"time"

	"farmme-api/model"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

// CowRepository is a contract 
type CowRepository interface {
	AddCow(cow model.Cow) (string, error)	
}
// CowRepositoryMongo connect db
type CowRepositoryMongo struct {
	ConnectionDB *mongo.Database
}
const (
	cowCollection = "cow"
)

//AddCow in farm 
func (cowMongo CowRepositoryMongo) AddCow(cow model.Cow) (string, error) {

	cow.CreatedTime = time.Now()
	cow.UpdatedTime = time.Now()
	res, err := cowMongo.ConnectionDB.Collection(cowCollection).InsertOne(context.TODO(), cow)
	if err != nil {
		log.Fatal(res)
	}

	fmt.Println("Inserted a single document: ", res.InsertedID)
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}
