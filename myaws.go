package mys3

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type S3HandlerModel struct {
	Cli s3iface.S3API //interface of *s3.S3
}


func (self S3HandlerModel) GetObject() *s3.GetObjectOutput {
	input := &s3.GetObjectInput{
		Bucket: aws.String("bucket-name"),
		Key:    aws.String("path/to/file"),
	}
	resp, err := self.Cli.GetObject(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	return resp
}

func (self S3HandlerModel) main() {
	self.GetObject()
}
