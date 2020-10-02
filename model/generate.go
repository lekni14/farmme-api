package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Generate object for db
type Generate struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name            string             `json:"name" bson:"name"`
	Description     string             `json:"description" bson:"description"`	
	Cover           string             `json:"cover" bson:"cover"`
	CoverThumb      string             `json:"cover_thumb" bson:"cover_thumb"`	
	OwnerID         primitive.ObjectID `json:"owner_id" bson:"owner_id"`
	Status          string             `json:"status" bson:"status"`
	Location        string             `json:"location" bson:"location"`
	IsActive        bool               `json:"is_active" bson:"is_active"`
	FarmCode          string			`json:"farm_code" bson:"farm_code"`	
	Address     	[]Farmaddress       `json:"address" bson:"address"`
	CreatedTime     time.Time          `json:"created_time" bson:"created_time"`
	UpdatedTime     time.Time          `json:"updated_time" bson:"updated_time"`
}
