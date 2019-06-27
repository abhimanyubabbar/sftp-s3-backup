package backup

import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Upload(bucket string, key string, body io.Reader) error {
	fmt.Printf("Uploading file to s3, bucket: %s, key: %s\n", bucket, key)

	conf := aws.Config{Region: aws.String("eu-west-1")}
	sess := session.New(&conf)
	svc := s3manager.NewUploader(sess)

	_, err := svc.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   body,
	})

	return err
}
