package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Cow object for db
type Cow struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Cover       string             `json:"cover" bson:"cover"`
	CoverThumb  string             `json:"cover_thumb" bson:"cover_thumb"`
	OwnerID     primitive.ObjectID `json:"owner_id" bson:"owner_id"`
	Status      string             `json:"status" bson:"status"`
	FarmID     primitive.ObjectID  `json:"farm_id" bson:"farm_id"`	
	CowCode		string             `json:"cow_code" bson:"cow_code"`
	MomID     	primitive.ObjectID `json:"mom_id" bson:"mom_id"`
	CreatedTime time.Time          `json:"created_time" bson:"created_time"`
	UpdatedTime time.Time          `json:"updated_time" bson:"updated_time"`
}
