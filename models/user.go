package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email       string             `bson:"email" json:"email" validate:"required"`
	DisplayName string             `bson:"displayname" json:"displayname"`
	Password    string             `bson:"password" json:"password" validate:"required"`
	PhotoURL    string             `bson:"photourl" json:"photourl"`
	IsAdmin     bool               `bson:"isadmin" json:"isadmin"`
	AdminKey    string             `bson:"adminKey,omitempty" json:"adminKey,omitempty"`
	CreatedAt   time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
