package main

import (
	"testing"
	"github.com/golang/mock/gomock"
	mock "./s3_mock"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

func TestS3Request(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s3mock := mock.NewMockS3API(ctrl)
	expected := &s3.GetObjectInput{
		Bucket: aws.String("bucket-name"),
		Key:    aws.String("path/to/file"),
	}

	s3mock.EXPECT().GetObject(expected)
	// OR s3mock.EXPECT().GetObject(gomock.Any())

	model := S3HandlerModel{}
	model.Cli = s3mock
	model.main()
}
