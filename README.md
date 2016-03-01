# aws-sdk-go s3 mock sample

## usage

```
go get github.com/mattn/gom
GOM_VENDOR_NAME=. gom install
```

```
go test .
```

## how to generate mock

```
$ mockgen -source ${GOPATH}/src/github.com/aws/aws-sdk-go/service/s3/s3iface/interface.go -destination src/example/s3_mock/s3mock.go -package s3mock
```
