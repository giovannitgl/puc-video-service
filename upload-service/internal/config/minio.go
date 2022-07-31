package config

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

type MinioInstance struct {
	Client *minio.Client
	Bucket string
}

var Minio MinioInstance

func SetupMinio() {
	ctx := context.Background()

	// Initialize minio client object.
	minioClient, errInit := minio.New(MinioEndpoint(), &minio.Options{
		Creds:  credentials.NewStaticV4(MinioAccessKey(), MinioSecretKey(), ""),
		Secure: false,
	})
	if errInit != nil {
		log.Fatalln(errInit)
	}

	location := "us-east-1"

	// Check to see if we already own this bucket (which happens if you run this twice)
	exists, err := minioClient.BucketExists(ctx, MinioBucket())
	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		err := minioClient.MakeBucket(ctx, MinioBucket(), minio.MakeBucketOptions{Region: location})
		if err != nil {
			log.Fatal(err)
		}

	}
	Minio = MinioInstance{minioClient, MinioBucket()}
}
