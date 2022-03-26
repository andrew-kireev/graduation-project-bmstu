package s3_music

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3MusicBucket interface {
	GetObject(context.Context, string, string) ([]byte, error)
	PutObject(context.Context, []byte, string, string) error
}

type client struct {
	client *s3.Client
}

func New() (S3MusicBucket, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("region=ru-central1"),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:   "yandex",
				URL:           "https://storage.yandexcloud.net",
				SigningRegion: "ru-central1",
			}, nil
		})),
	)

	if err != nil {
		return nil, err
	}

	s3Client := s3.NewFromConfig(cfg)

	return &client{
		client: s3Client,
	}, nil
}

func (s3Client *client) GetObject(ctx context.Context, bucketName, fileKey string) ([]byte, error) {
	res, err := s3Client.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &fileKey,
	})

	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)

	defer res.Body.Close()

	return buf.Bytes(), nil
}

func (s3Client *client) PutObject(ctx context.Context, file []byte, bucketName, fileKey string) error {
	reader := bytes.NewReader(file)

	_, err := s3Client.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &fileKey,
		Body:   reader,
	})

	return err
}
