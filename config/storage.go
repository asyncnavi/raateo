package config

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"log/slog"
)

func SetupStorage(url string) (*cloudinary.Cloudinary, context.Context) {

	fmt.Println("Attempting to connect storage")
	cld, err := cloudinary.NewFromURL(url)

	if err != nil {
		slog.Error("Could not connect to cloudinary storage", err)
		return nil, nil
	}

	cld.Config.URL.Secure = true

	ctx := context.Background()

	slog.Debug("Connected to cloudinary storage")

	return cld, ctx
}
