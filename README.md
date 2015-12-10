# aws-sdk-go s3 mock sample

## usage

```
$ go get github.com/golang/mock/gomock
$ go get github.com/golang/mock/mockgen
$ go get github.com/aws/aws-sdk-go
```

```
go test .
```

## how to generate mock

```
$ mockgen -source ${GOPATH}/src/github.com/aws/aws-sdk-go/service/s3/s3iface/interface.go -destination s3_mock/s3mock.go -package s3mock
```
