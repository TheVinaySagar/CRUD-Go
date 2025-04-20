package controllers

import (
	"log"
	"net/http"

	"myproject/config"

	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Uploadimage(c *gin.Context) {
	file, err := c.FormFile("image")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	cld, ctx := config.Credentials()

	err = c.SaveUploadedFile(file, "assets/uploads/"+file.Filename)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to save the image"})
		return
	}

	val := uuid.New().String()

	uploadResult, err := cld.Upload.Upload(
		ctx,
		"assets/uploads/"+file.Filename,
		uploader.UploadParams{PublicID: val,
			UniqueFilename: api.Bool(false),
			Overwrite:      api.Bool(true)})

	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
		return
	}
	log.Println(uploadResult.SecureURL)

	c.JSON(http.StatusAccepted, gin.H{"massage": uploadResult.SecureURL})
}
