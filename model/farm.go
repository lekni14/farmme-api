package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Farm object for db
type Farm struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name            string             `json:"name" bson:"name"`
	Description     string             `json:"description" bson:"description"`	
	Cover           string             `json:"cover" bson:"cover"`
	CoverThumb      string             `json:"cover_thumb" bson:"cover_thumb"`	
	Product         []ProduceFarm     `json:"product" bson:"product"`	
	OwnerID         primitive.ObjectID `json:"owner_id" bson:"owner_id"`
	Status          string             `json:"status" bson:"status"`
	Location        string             `json:"location" bson:"location"`
	IsActive        bool               `json:"is_active" bson:"is_active"`
	FarmID          string			   `json:"farm_id" bson:"farm_id"`	
	Address     	[]FarmAddress          `json:"address" bson:"address"`
	CreatedTime     time.Time          `json:"created_time" bson:"created_time"`
	UpdatedTime     time.Time          `json:"updated_time" bson:"updated_time"`
}

type ProduceFarm struct {
	ProductID primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name" binding:"required"`
	Image     []ProductImage     `json:"image" bson:"image"`
	Detail    string             `json:"detail" bson:"detail"`
	Type      []ProductType      `json:"type" bson:"type"`
	Unit      int                `json:"unit" bson:"unit"`
	Currency  string             `json:"currency" bson:"currency"`
	Status    string             `json:"status" bson:"status"`
	Reuse     bool               `json:"reuse" bson:"reuse"`	
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	
}

type ProductImage struct {
	PathURL string `json:"path_url" bson:"path_url"`
}

type ProductType struct {
	Name   string  `json:"name" bson:"name"`
	Remark string  `json:"remark" bson:"remark"`
	Price  float64 `json:"price" bson:"price"`
}

// Address farm
type FarmAddress struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Address   string             `json:"address" bson:"address" binding:"required"`
	Province  string             `json:"province" bson:"province" binding:"required"`
	District  string             `json:"district" bson:"district" binding:"required"`
	City      string             `json:"city" bson:"city" binding:"required"`
	ZipCode   string             `json:"zipcode" bson:"zipcode" binding:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}