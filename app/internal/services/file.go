package services

import (
	"log"
	"time"

	"example.com/internal/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetFilePath(user *models.Swms) models.Swms {
	if user.FileName != nil {
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("ap-southeast-2")},
		)

		// Create S3 service client
		svc := s3.New(sess)

		req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String("reslasian"),
			Key:    aws.String(*user.FileName),
		})
		urlStr, err := req.Presign(15 * time.Minute)

		if err != nil {
			log.Println("Failed to sign request", err)
		}
		//get presiged url from s3
		user.FilePath = &urlStr

	}
	return *user
}
