package services

import (
	"detection-no-helmet-web-application/api/responses"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func SavePicture(base64 string, location string) interface{} {
	newBase := strings.Replace(base64, "data:image/png;base64,", "", -1)

	_, res := base64toPng(newBase, location)
	//base64toJpg(base64, location)
	resSuccess := responses.BaseSingleResponse{Status: 200, Message: "success", Value: res}

	return resSuccess
}

func base64toPng(dataBase64 string, location string) (bool, string) {
	data := dataBase64

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	bounds := m.Bounds()
	fmt.Println(bounds, formatString)
	//univertity_17_11_2022_unixtime.png
	//Encode from image format to writer
	t := time.Now()
	year, month, day := t.Date()
	hour := strconv.FormatInt(int64(t.Hour()), 10)
	minute := strconv.FormatInt(int64(t.Minute()), 10)
	second := strconv.FormatInt(int64(t.Second()), 10)
	timestamp := hour + "_" + minute + "_" + second
	pngFilename := location + "_" + strconv.FormatInt(int64(day), 10) + "_" + strconv.FormatInt(int64(month), 10) + "_" + strconv.FormatInt(int64(year), 10) + "_" + timestamp + ".png"
	f, err := os.OpenFile("assets/images/"+pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return false, "Error"
	}

	err = png.Encode(f, m)
	if err != nil {
		log.Fatal(err)
		return false, "Error"
	}
	fmt.Println("Png file", pngFilename, "created")

	return true, pngFilename

}
