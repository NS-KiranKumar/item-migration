package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	GLID              string             `bson:"glid" json:"glid"`
	ItemType          string             `bson:"itemType" json:"itemType"`
	TransactionNumber string             `bson:"transactionNumber" json:"transactionNumber"`

	RBOCode string `bson:"rboCode" json:"rboCode"`
	RBOName string `bson:"rboName" json:"rboName"`

	Comment string `bson:"comment" json:"comment"`

	VPFileName  string `bson:"vpFileName" json:"vpFileName"`
	VLPFileName string `bson:"vlpFileName" json:"vlpFileName"`

	ItemStatus     ItemStatus         `bson:"itemStatus" json:"itemStatus"`
	ItemAssignedTo primitive.ObjectID `bson:"itemAssignedTo" json:"itemAssignedTo"`

	CreatedBy primitive.ObjectID `bson:"createdBy" json:"createdBy"`

	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

type ItemStatus struct {
	PS         string `bson:"ps" json:"ps"`
	VD         string `bson:"vd" json:"vd"`
	AWLayering string `bson:"awLayering" json:"awLayering"`
	GPMDesign  string `bson:"gpmDesign" json:"gpmDesign"`
	GPMDev     string `bson:"gpmDev" json:"gpmDev"`
	CPQQC      string `bson:"cpqQc" json:"cpqQc"`
	MPQC       string `bson:"mpQc" json:"mpQc"`
	BizzPass   string `bson:"bizzPass" json:"bizzPass"`
	BizzFail   string `bson:"bizzFail" json:"bizzFail"`
}
