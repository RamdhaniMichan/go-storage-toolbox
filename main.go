package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	. "go-storage/config/storage"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	ctx := context.Background()
	// ======== sample GCS ==========

	//err := NewStorageProvider().InitGCS().UploadFile(ctx, os.Getenv("GCS_BUCKET"), "Hotpot.png", "./Hotpot.png")
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//err = NewStorageProvider().InitGCS().DownloadFile(ctx, os.Getenv("GCS_BUCKET"), "Hotpot.png", "./file.png")
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	// ======== sample MinIO =========

	//err = NewStorageProvider().InitMinIO().UploadFile(ctx, os.Getenv("MINIO_BUCKET"), "images/", "./Hotpot.png")
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	//fmt.Println("success upload to storage")

	err := NewStorageProvider().InitMinIO().DownloadFile(ctx, os.Getenv("MINIO_BUCKET"), "images/Hotpot.png", "./file.png")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("success download from storage")
}
