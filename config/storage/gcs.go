package storage

import (
	"cloud.google.com/go/storage"
	"context"
	"io"
	"os"
)

type GCSProvider struct {
	client *storage.Client
}

func (p *GCSProvider) UploadFile(ctx context.Context, bucket string, prefix string, filepath string) error {
	writer := p.client.Bucket(bucket).Object(prefix).NewWriter(ctx)
	defer writer.Close()

	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	if _, err := writer.Write(data); err != nil {
		return err
	}

	return writer.Close()
}

func (p *GCSProvider) DownloadFile(ctx context.Context, bucket string, object string, filePath string) error {
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
