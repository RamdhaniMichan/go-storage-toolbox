package storage

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"os"
)

type MinioProvider struct {
	client *minio.Client
}

func (p *MinioProvider) UploadFile(ctx context.Context, bucket string, prefix string, filepath string) error {
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

func (p *MinioProvider) DownloadFile(ctx context.Context, bucket string, object string, filePath string) error {
	err := p.client.FGetObject(ctx, bucket, object, filePath, minio.GetObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}
