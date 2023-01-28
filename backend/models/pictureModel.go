package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Picture struct {
	Id               primitive.ObjectID `json:"id,omitempty"`
	Location         string             `json:"location" :"Location" validate:"required"`
	Base64DefaultImg string             `json:"base64DefaultImg" :"Base64DefaultImg" validate:"required"`
	Base64RiderImg   string             `json:"base64RiderImg" :"Base64RiderImg" validate:"required"`
	ImgName          string             `json:"imgName" :"ImgName" validate:"required"`
	CreateAt         time.Time          `json:"createAt" :"CreateAt" validate:"required"`
}

type PictureRequest struct {
	Location         string `json:"location" :"Location" validate:"required"`
	Base64DefaultImg string `json:"base64DefaultImg" :"Base64DefaultImg" validate:"required"`
	Base64RiderImg   string `json:"base64RiderImg" :"Base64RiderImg" validate:"required"`
}
