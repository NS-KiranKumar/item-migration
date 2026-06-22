package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string             `bson:"name" json:"name"`
	Email        string             `bson:"email" json:"email"`
	Password     string             `bson:"password" json:"-"` // never expose
	NamyaID      string             `bson:"namyaId" json:"namyaId"`
	Team         string             `bson:"team" json:"team"`
	MobileNumber string             `bson:"mobileNumber" json:"mobileNumber"`
	Designation  Designation        `bson:"designation" json:"designation"`

	IsActive  bool `bson:"isActive" json:"isActive"`
	IsDeleted bool `bson:"isDeleted" json:"isDeleted"`

	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

type Designation struct {
	GPMResource string `bson:"gpmResource,omitempty" json:"gpmResource,omitempty"`
	VDResource  string `bson:"vdResource,omitempty" json:"vdResource,omitempty"`
	TeamLead    string `bson:"teamLead,omitempty" json:"teamLead,omitempty"`
	Admin1      string `bson:"admin1,omitempty" json:"admin1,omitempty"`
	Admin2      string `bson:"admin2,omitempty" json:"admin2,omitempty"`
}

type LoginRequest struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

type ResetPassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
