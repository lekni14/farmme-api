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

// GenRepository is a contract 
type GenRepository interface {
	AddGen(gen model.Generate)  (string, error)	
}
// GenRepositoryMongo connect db
type GenRepositoryMongo struct {
	ConnectionDB *mongo.Database
}
const (
	generateCollection = "generate"
)

//AddGen in farm 
func (genMongo CowRepositoryMongo) AddGen(gen model.Generate) (string, error) {

	gen.CreatedTime = time.Now()
	gen.UpdatedTime = time.Now()
	res, err := genMongo.ConnectionDB.Collection(generateCollection).InsertOne(context.TODO(), gen)
	if err != nil {
		log.Fatal(res)
	}

	fmt.Println("Inserted a single document: ", res.InsertedID)
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}
