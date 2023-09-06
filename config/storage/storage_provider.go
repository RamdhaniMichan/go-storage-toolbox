package storage

import (
	gcs "cloud.google.com/go/storage"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"os"
)

type provider struct{}

type StorageProvider interface {
	UploadFile(ctx context.Context, bucket string, prefix string, filepath string) error
	DownloadFile(ctx context.Context, bucket string, object string, filePath string) error
}

func NewStorageProvider() *provider {
	return &provider{}
}

func (p *provider) InitGCS() StorageProvider {
	os.Setenv(os.Getenv("GCS_KEY_CREDENTIAL"), os.Getenv("GCS_FILE_CREDENTIAL"))
	ctx := context.Background()
	gcsClient, err := gcs.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer gcsClient.Close()

	return &GCSProvider{gcsClient}
}

func (p *provider) InitMinIO() StorageProvider {
	minioClient, err := minio.New(os.Getenv("MINIO_HOST"), &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_ID"), os.Getenv("MINIO_SECRET"), os.Getenv("MINIO_TOKEN")),
		Secure: false,
	})

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &MinioProvider{minioClient}
}
