package config

import (
	"context"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func Credentials() (*cloudinary.Cloudinary, context.Context) {
	CLOUDINARY_URL := os.Getenv("CLOUDINARY_URL")
	cld, err := cloudinary.NewFromURL(CLOUDINARY_URL)

	if err != nil {
		log.Fatalf("Cloudinary is not setup, %v", err)
	}

	log.Println("Connected to Cloudinary successfully")
	cld.Config.URL.Secure = true
	ctx := context.Background()
	return cld, ctx
}
