package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID                  primitive.ObjectID   `bson:"_id" json:"id"`
	Name                string               `bson:"name" json:"name" validate:"required"`
	StartAt             time.Time            `bson:"start_at" json:"start_at" validate:"required"`
	EndAt               time.Time            `bson:"end_at" json:"end_at" validate:"required,gt={start_at}"`
	RequireRSVPApproval bool                 `bson:"require_rsvp_approval" json:"require_rsvp_approval" default:"false"`
	GeoLatitude         float64              `bson:"geo_latitude" json:"geo_latitude" default:"12.9716"`
	GeoLongitude        float64              `bson:"geo_longitude" json:"geo_longitude" default:"77.5946"`
	Price               float64              `bson:"price" json:"price" default:"0"`
	Capacity            int                  `bson:"capacity" json:"capacity" default:"0"` // unlimited
	IsPublic            bool                 `bson:"is_public" json:"is_public" default:"true"`
	Summary             string               `bson:"summary" json:"summary"`
	Banner              string               `bson:"banner" json:"banner"` // add our link here
	IsVirtual           bool                 `bson:"is_virtual" json:"is_virtual" default:"false"`
	MeetingURL          string               `bson:"meeting_url" json:"meeting_url"`
	GeoAddress          string               `bson:"geo_address" json:"geo_address"`
	Participants        []primitive.ObjectID `bson:"participants" json:"participants"`
	Attendees           []primitive.ObjectID `bson:"attendees" json:"attendees"`
	Hosts               []primitive.ObjectID `bson:"hosts" json:"hosts"`
	Trainers            []primitive.ObjectID `bson:"trainers" json:"trainers"`
	CreatedAt           time.Time            `bson:"created_at" json:"created_at"`
	CreatedBy           primitive.ObjectID   `bson:"created_by" json:"created_by"`
}
