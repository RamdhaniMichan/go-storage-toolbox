package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type StorageProvider interface {
	UploadFile(bucket string, object_file string, filepath string) error
	DownloadFile(bucket string, object string, filePath string) error
}

type GCSProvider struct {
	client *storage.Client
}

type MinioProvider struct {
	client *minio.Client
}

func (p *GCSProvider) UploadFile(bucket string, prefix string, filepath string) error {
	ctx := context.Background()
	writer := p.client.Bucket(bucket).Object(prefix).NewWriter(ctx)
	defer writer.Close()

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	if _, err := writer.Write(data); err != nil {
		return err
	}

	return writer.Close()
}

func (p *GCSProvider) DownloadFile(bucket string, object string, filePath string) error {
	ctx := context.Background()
	reader, err := p.client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		return err
	}
	defer reader.Close()

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	if err != nil {
		return err
	}

	return nil
}

func (p *MinioProvider) UploadFile(bucket string, prefix string, filepath string) error {
	ctx := context.Background()
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	_, err = p.client.PutObject(ctx, bucket, fmt.Sprintf("%s%s", prefix, stat.Name()), file, stat.Size(), minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (p *MinioProvider) DownloadFile(bucket string, object string, filePath string) error {
	ctx := context.Background()
	err := p.client.FGetObject(ctx, bucket, object, filePath, minio.GetObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "your-credential.json")
	ctx := context.Background()
	gcsClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	defer gcsClient.Close()

	gcsConnection := gcsClient

	minioClient, err := minio.New("host:port", &minio.Options{
		Creds:  credentials.NewStaticV4("access_key_id", "secret_key_id", "token"),
		Secure: false,
	})

	if err != nil {
		log.Println(err)
		return
	}

	minioConnection := minioClient

	storageProvider := make(map[string]StorageProvider)
	storageProvider["gcs"] = &GCSProvider{client: gcsConnection}
	storageProvider["minio"] = &MinioProvider{client: minioConnection}

	provider := storageProvider["minio"]

	// ======== sample GCS ==========

	err = provider.UploadFile("privy-acceleration", "Hotpot.png", "./Hotpot.png")
	if err != nil {
		log.Println(err)
		return
	}

	err = provider.DownloadFile("privy-acceleration", "Hotpot.png", "./file.png")
	if err != nil {
		log.Println(err)
		return
	}

	// ======== sample MinIO =========

	err = provider.UploadFile("accel", "images/", "./Hotpot.png")
	if err != nil {
		log.Println(err)
		return
	}

	err = provider.DownloadFile("accel", "images/Hotpot.png", "./file.png")
	if err != nil {
		log.Println(err)
		return
	}
}
