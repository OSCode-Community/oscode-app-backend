package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Socials struct {
	Linkedin string `bson:"linkedin" json:"linkedin"`
	Github   string `bson:"github" json:"github"`
}

type User struct {
	ID              primitive.ObjectID `bson:"_id"`
	FirstName       string             `bson:"first_name" json:"first_name" validate:"required,min=2,max=100"`
	LastName        string             `bson:"last_name" json:"last_name"`
	Password        string             `bson:"password" json:"-"`
	Email           string             `bson:"email" json:"email"`
	Phone           string             `bson:"phone" json:"phone"`
	IsPhoneVerified bool               `bson:"IsPhoneVerified" json:"IsPhoneVerified"`
	ProfilePic      string             `bson:"ProfilePic" json:"ProfilePic"`
	Birthday        time.Time          `bson:"birthday" json:"birthday"`
	Bio             string             `bson:"bio" json:"bio"`
	Links           *Socials           `bson:"links" json:"links"`
	Organisation    string             `bson:"organisation" json:"organisation"`
	Domain          string             `bson:"domain" json:"domain"`
	Course          string             `bson:"course" json:"course" validate:"eq=BTECH|eq=BSC"`
	PassoutYear     string             `bson:"passout_year" json:"passout_year"`
	UserType        string             `bson:"user_type" json:"user_type" validate:"required,eq=MEMBER|eq=STUDENT"`
	CreatedAt       time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt       time.Time          `bson:"updatedAt" json:"updatedAt"`
}
