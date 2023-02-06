package services

import (
	"detection-no-helmet-web-application/api/models"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"strings"
)

func SavePicture(imgDefault *models.SavePicture, imgRider *models.SavePicture, path string) {

	CheckExistDir(path)
	base64toJpg(imgDefault.Base64, path, imgDefault.Name)
	base64toJpg(imgRider.Base64, path, imgRider.Name)

}

func base64toJpg(data string, path string, name string) bool {

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	fmt.Println("base64toJpg", bounds, formatString)

	//Encode from image format to writer
	pngFilename := path + "/" + name + ".jpg"
	f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return false
	}

	err = jpeg.Encode(f, m, &jpeg.Options{Quality: 75})
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("Jpg file", pngFilename, "created")
	return true
}
