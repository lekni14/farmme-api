package repository

import (
	"context"
	"fmt"
	"log"
	// "strings"
	"time"

	"farmme-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)
// FarmRepository is a contract for consuming data by business logic layer
type FarmRepository interface {
	AddFarm(farm model.Farm) (string, error)
	ExistByName(name string) (bool, error)
	ExistByNameForEdit(name string, farmID string) (bool, error)
	GetFarmByUser(userID string) ([]model.Farm, error) 
}
// FarmRepositoryMongo connect db
type FarmRepositoryMongo struct {
	ConnectionDB *mongo.Database
}

const (
	farmCollection = "farm"
)
//AddFarm new farm 
func (farmMongo FarmRepositoryMongo) AddFarm(farm model.Farm) (string, error) {

	farm.CreatedTime = time.Now()
	farm.UpdatedTime = time.Now()
	res, err := farmMongo.ConnectionDB.Collection(farmCollection).InsertOne(context.TODO(), farm)
	if err != nil {
		log.Fatal(res)
	}

	fmt.Println("Inserted a single document: ", res.InsertedID)
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}
//ExistByName find nameFarm
func (farmMongo FarmRepositoryMongo) ExistByName(name string) (bool, error) {

	filter := bson.D{{"name", name}}
	count, err := farmMongo.ConnectionDB.Collection(farmCollection).CountDocuments(context.TODO(), filter)
	log.Printf("[info] count %s", count)
	if err != nil {
		log.Println(err)
	}
	if count > 0 {
		return true, nil
	}

	return false, nil
}
//ExistByNameForEdit find Name farmMongo
func (farmMongo FarmRepositoryMongo) ExistByNameForEdit(name string, farmID string) (bool, error) {

	var farm model.Farm
	id, err := primitive.ObjectIDFromHex(farmID)
	filter := bson.M{"_id": id}
	err2 := farmMongo.ConnectionDB.Collection(farmCollection).FindOne(context.TODO(), filter).Decode(&farm)
	log.Printf("[info] farm %s", err2)
	if err2 != nil {
		log.Fatal(err2)
		//return true, err2
	}
	if farm.Name == name {
		return false, nil
	}

	filter2 := bson.D{{"name", name}}
	count, err := farmMongo.ConnectionDB.Collection(farmCollection).CountDocuments(context.TODO(), filter2)
	log.Printf("[info] count %s", count)
	if err != nil {
		log.Println(err)
	}
	if count > 0 {
		return true, nil
	}

	return false, nil
}
// GetFarmByUser for farm
func (farmMongo FarmRepositoryMongo) GetFarmByUser(userID string) ([]model.Farm, error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
	var farms []model.Farm
	filter := bson.D{{"owner_id", objectID}}
	cur, err := farmMongo.ConnectionDB.Collection(farmCollection).Find(context.TODO(), filter)
	//log.Printf("[info] cur %s", cur)
	if err != nil {
		log.Println(err)
	}

	for cur.Next(context.TODO()) {
		var u model.Farm
		// decode the document
		if err := cur.Decode(&u); err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("post: %+v\n", p)
		farms = append(farms, u)
	}

	return farms, err
}