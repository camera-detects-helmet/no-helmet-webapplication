package controller

import (
	"context"
	"detection-no-helmet-web-application/api/configs"
	"detection-no-helmet-web-application/api/models"
	"detection-no-helmet-web-application/api/responses"
	"detection-no-helmet-web-application/api/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type CarType struct {
	Location string `json:"location"`
	Picture  string `json:"picture"`
	Time     string `json:"time"`
}

var imageCollection = configs.GetCollection(configs.DB, "images")
var validate = validator.New()

func CreateImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var imagesRequest models.PictureRequest
		defer cancel()

		fmt.Println(imagesRequest)

		//validate the request body
		if err := c.BindJSON(&imagesRequest); err != nil {
			c.JSON(http.StatusBadRequest, responses.DefaultResponse{StatusCode: http.StatusBadRequest, Message: "Bad Request Validate body error", Data: map[string]interface{}{"error": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&imagesRequest); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.DefaultResponse{StatusCode: http.StatusBadRequest, Message: "Bad Request validate required fields", Data: map[string]interface{}{"error": validationErr.Error()}})
			return
		}
		loc, _ := time.LoadLocation("Asia/Bangkok")
		_time := time.Now().In(loc)
		imgDefaultName := imagesRequest.Location + "_" + _time.Format("02_01_2006_15_04_05") + "_default"
		imgRiderName := imagesRequest.Location + "_" + _time.Format("02_01_2006_15_04_05") + "_rider"
		pathSaveImg := "assets/images/"

		var imageDefault = &models.SavePicture{
			Base64: imagesRequest.Base64DefaultImg,
			Name:   imgDefaultName,
		}
		var imageRider = &models.SavePicture{
			Base64: imagesRequest.Base64RiderImg,
			Name:   imgRiderName,
		}
		//create new image
		services.SavePicture(imageDefault, imageRider, pathSaveImg)
		uriPathDefaultImg := configs.EnvIP_PORT() + "/images/" + imgDefaultName + ".jpg"
		uriPathRiderImg := configs.EnvIP_PORT() + "/images/" + imgRiderName + ".jpg"

		fmt.Println(uriPathDefaultImg)
		fmt.Println(uriPathRiderImg)

		newImg := models.Picture{
			Id:             primitive.NewObjectID(),
			Location:       imagesRequest.Location,
			PathDefaultImg: uriPathDefaultImg,
			PathRiderImg:   uriPathRiderImg,
			ImgName:        imagesRequest.Location + "_" + _time.Format("02-01-2006 15:04:05"),
			CreateAt:       time.Date(_time.Year(), _time.Month(), _time.Day(), _time.Hour(), _time.Minute(), _time.Second(), _time.Nanosecond(), loc),
		}
		fmt.Println(time.Date(_time.Year(), _time.Month(), _time.Day(), _time.Hour(), _time.Minute(), _time.Second(), _time.Nanosecond(), loc))
		result, err := imageCollection.InsertOne(ctx, newImg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.DefaultResponse{StatusCode: http.StatusInternalServerError, Message: "Internal Server Error", Data: map[string]interface{}{"error": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, responses.DefaultResponse{StatusCode: http.StatusCreated, Message: "Success", Data: map[string]interface{}{"data": result}})
	}
}

func GetAllImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var image models.Picture

		var images []models.Picture
		cursor, err := imageCollection.Find(ctx, bson.D{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.DefaultResponse{StatusCode: http.StatusInternalServerError, Message: "Internal Server Error", Data: map[string]interface{}{"error": err.Error()}})
			return
		}
		if err = cursor.All(ctx, &images); err != nil {
			c.JSON(http.StatusInternalServerError, responses.DefaultResponse{StatusCode: http.StatusInternalServerError, Message: "Internal Server Error", Data: map[string]interface{}{"error": err.Error()}})
			return
		}
		defer func(cursor *mongo.Cursor, ctx context.Context) {
			err := cursor.Close(ctx)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.DefaultResponse{StatusCode: http.StatusInternalServerError, Message: "Internal Server Error", Data: map[string]interface{}{"error": err.Error()}})
				return
			}
		}(cursor, ctx)

		for cursor.Next(ctx) {
			err := cursor.Decode(&image)
			if err != nil {
				return
			}
			images = append(images, image)

		}
		fmt.Println(len(images))

		var res = map[string]interface{}{"data": images, "size": len(images)}
		c.JSON(http.StatusOK, responses.DefaultResponse{StatusCode: http.StatusOK, Message: "Success", Data: res})
	}
}

func GetImageById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var image models.Picture
		id := c.Param("id")
		fmt.Println(id)
		objId, _ := primitive.ObjectIDFromHex(id)
		err := imageCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&image)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.DefaultResponse{StatusCode: http.StatusInternalServerError, Message: "Internal Server Error", Data: map[string]interface{}{"error": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.DefaultResponse{StatusCode: http.StatusOK, Message: "Success", Data: map[string]interface{}{"data": image}})
	}
}
