package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go-storage/config/storage"
	"os"
	"testing"
)

func TestSuccessUploadFile(t *testing.T) {
	godotenv.Load()
	ctx := context.Background()
	err := storage.NewStorageProvider().InitMinIO().UploadFile(ctx, os.Getenv("MINIO_BUCKET"), "images/", "./Hotpot.png")
	assert.NoError(t, err)
}

func TestSuccessDownloadFile(t *testing.T) {
	godotenv.Load()
	ctx := context.Background()
	err := storage.NewStorageProvider().InitMinIO().DownloadFile(ctx, os.Getenv("MINIO_BUCKET"), "images/Hotpot.png", "./file.png")
	assert.NoError(t, err)
}
