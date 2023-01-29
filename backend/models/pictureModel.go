package models

import (
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Picture struct {
	Id               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Location         string             `json:"location" :"Location" validate:"required" bson:"location"`
	Base64DefaultImg string             `json:"base64DefaultImg" :"Base64DefaultImg" validate:"required" bson:"Base64DefaultImg"`
	Base64RiderImg   string             `json:"base64RiderImg" :"Base64RiderImg" validate:"required" bson:"Base64RiderImg"`
	ImgName          string             `json:"imgName" :"ImgName" validate:"required" bson:"ImgName"`
	CreateAt         time.Time          `json:"createAt" :"CreateAt" validate:"required" bson:"CreateAt"`
}

type PictureRequest struct {
	Location         string `json:"location" :"Location" validate:"required"`
	Base64DefaultImg string `json:"base64DefaultImg" :"Base64DefaultImg" validate:"required"`
	Base64RiderImg   string `json:"base64RiderImg" :"Base64RiderImg" validate:"required"`
}
