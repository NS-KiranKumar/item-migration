package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RBO struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	RBOName string             `bson:"rboName" json:"rboName"`
	RBOCode string             `bson:"rboCode" json:"rboCode"`

	CreatedBy primitive.ObjectID `bson:"createdBy" json:"createdBy"`

	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}
