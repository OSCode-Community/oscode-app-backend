package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TrainerName struct {
	FirstName string `bson:"first_name" json:"first_name" validate:"required,min=2,max=100"`
	LastName  string `bson:"last_name" json:"last_name"`
}

type Trainer struct {
	ID              primitive.ObjectID `bson:"_id"`
	Name            *TrainerName       `bson:"name" json:"name"`
	Password        string             `bson:"password" json:"-" validate:"required,min=4"`
	Email           string             `bson:"email" json:"email" validate:"email,required"`
	Phone           string             `bson:"phone" json:"phone" validate:"required"`
	IsPhoneVerified bool               `bson:"IsPhoneVerified" json:"IsPhoneVerified"`
	ProfilePic      string             `bson:"ProfilePic" json:"ProfilePic"`
	Birthday        time.Time          `bson:"birthday" json:"birthday"`
	Bio             string             `bson:"bio" json:"bio"`
	Links           *Socials           `bson:"links" json:"links"`
	Organisation    string             `bson:"organisation" json:"organisation"`
	IsMentor        bool               `bson:"is_mentor" json:"is_mentor"`
	CreatedAt       time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt       time.Time          `bson:"updatedAt" json:"updatedAt"`
}
