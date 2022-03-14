package services

import (
	"bytes"
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type minioService struct {
	minioClient *minio.Client
}

func NewMinioService(host, id, secret string) (MinioService, error) {

	minioClient, err := minio.New(
		host,
		&minio.Options{
			Creds:  credentials.NewStaticV4(id, secret, ""),
			Secure: false,
		},
	)
	if err != nil {
		return nil, err
	}

	return &minioService{
		minioClient: minioClient,
	}, nil
}

func (ms *minioService) GetBucket(ctx context.Context, bucketName string) ([]string, error) {
	objects := []string{}

	for bucket := range ms.minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{}) {
		objects = append(objects, bucket.Key)
	}

	return objects, nil
}

func (ms *minioService) GetObject(ctx context.Context, bucketName, objectName string) ([]byte, error) {
	reader, err := ms.minioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (ms *minioService) PutObject(ctx context.Context, bucketName, objectName string, data []byte) error {
	_, err := ms.minioClient.PutObject(ctx, bucketName, objectName, bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (ms *minioService) DeleteObject(ctx context.Context, bucketName, objectName string) error {
	return ms.minioClient.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
}
