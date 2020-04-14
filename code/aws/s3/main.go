package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/gommon/log"
)

func cp() {

	bucket := "s3://my-bucket-1"
	key := "path/to/file.txt"

	sess := session.Must(session.NewSession())
	svc := s3.New(sess)
	ctx := context.Background()
	s3object := s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   os.Stdin,
	}
	_, err := svc.PutObjectWithContext(ctx, &s3object)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info("Ok.")
}

func main() {

	fmt.Println("[TRACE] ### START ###")

	cp()

	fmt.Println("[TRACE] --- END ---")
}
