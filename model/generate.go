package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Generate object for db
type Generate struct {
	ID              primitive.ObjectID 	`json:"id" bson:"_id,omitempty"`
	
	Count           string             	`json:"count" bson:"count"`
	Gene     string             `json:"gene" bson:"gene"`	
	GenerateDate	time.Time             `json:"gen_date" bson:"gen_date"`	
	Pg	bool             `json:"pg" bson:"pg"`	
	Cow     		[]CowDetail          `json:"cow" bson:"cow"`
	Farm     		[]FarmDetail          `json:"farm" bson:"farm"`
	Cover           string             `json:"cover" bson:"cover"`
	CoverThumb      string             `json:"cover_thumb" bson:"cover_thumb"`	
	OwnerID         primitive.ObjectID `json:"owner_id" bson:"owner_id"`
	
	Status          string             `json:"status" bson:"status"`
	
	FarmCode          string			`json:"farm_code" bson:"farm_code"`
	CreatedTime     time.Time          `json:"created_time" bson:"created_time"`
	UpdatedTime     time.Time          `json:"updated_time" bson:"updated_time"`
}

// CowDetail data cow
type CowDetail struct {
	ID              primitive.ObjectID 	`json:"id" bson:"_id,omitempty"`
	Name        	string             `json:"name" bson:"name"`	
	Age        		string             `json:"age" bson:"age"`
	Lactation		string				`json:"lactation" bson:"lactation"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

// FarmDetail data cow
type FarmDetail struct {
	ID              primitive.ObjectID 	`json:"id" bson:"_id,omitempty"`
	FarmCode          string			`json:"farm_code" bson:"farm_code"`
	FarmName     string             `json:"farm_name" bson:"farm_name"`	
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
