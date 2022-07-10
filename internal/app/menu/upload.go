package menu

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// S3
var uploader *s3manager.Uploader

func menu() {
	uploader = NewUploader()

	upload()
}

func NewUploader() *s3manager.Uploader {
	s3Config := &aws.Config{
		Region:      aws.String("ap-southeast-3"),
		Credentials: credentials.NewStaticCredentials("AKIAXIFHHPJNFMHL7TZG", "zMKhtuurnebNfYkSxZc1RjaSlaJFDJH4gaHxvgDr", ""),
	}

	s3Session := session.New(s3Config)

	uploader := s3manager.NewUploader(s3Session)
	return uploader
}

func upload() {
	log.Println("uploading")

	file, err := ioutil.ReadFile("./my_cat.jpg")
	if err != nil {
		log.Fatal(err)
	}

	upInput := &s3manager.UploadInput{
		Bucket:      aws.String("menu-service-bucket"), // bucket's name
		Key:         aws.String("menu/my_cat.jpg"),     // files destination location
		Body:        bytes.NewReader(file),             // content of the file
		ContentType: aws.String("image/jpg"),           // content type
	}
	res, err := uploader.UploadWithContext(context.Background(), upInput)
	log.Printf("res %+v\n", res)
	log.Printf("err %+v\n", err)
}
