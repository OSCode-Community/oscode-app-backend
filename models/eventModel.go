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
	Timezone            string               `bson:"timezone" json:"timezone"`
	RequireRSVPApproval bool                 `bson:"require_rsvp_approval" json:"require_rsvp_approval"`
	GeoLatitude         float64              `bson:"geo_latitude" json:"geo_latitude"`
	GeoLongitude        float64              `bson:"geo_longitude" json:"geo_longitude"`
	Price               float64              `bson:"price" json:"price" validate:"required,gte=0"`
	Capacity            int                  `bson:"capacity" json:"capacity" validate:"required,gt=0"`
	IsPublic            bool                 `bson:"is_public" json:"is_public"`
	Summary             string               `bson:"summary" json:"summary" validate:"required"`
	Banner              string               `bson:"banner" json:"banner"`
	IsVirtual           bool                 `bson:"is_virtual" json:"is_virtual"`
	MeetingURL          string               `bson:"meeting_url" json:"meeting_url" validate:"required_if=is_virtual"`
	GeoAddress          string               `bson:"geo_address" json:"geo_address"`
	Participants        []primitive.ObjectID `bson:"participants" json:"participants"`
	Attendees           []primitive.ObjectID `bson:"attendees" json:"attendees"`
	Hosts               []primitive.ObjectID `bson:"hosts" json:"hosts"`
	Trainers            []primitive.ObjectID `bson:"trainers" json:"trainers"`
	CreatedAt           time.Time            `bson:"created_at" json:"created_at"`
	CreatedBy           primitive.ObjectID   `bson:"created_by" json:"created_by"`
}
