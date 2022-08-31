package aws

import (
	"fmt"
	"gitee.com/go-course/cloud-station-g7/store/aliyun"
)

type AwsOssStore struct {
}

func  NewAwsOssStore(option aliyun.Option) (*AwsOssStore, error) {
	return &AwsOssStore{}, nil
}

func (a *AwsOssStore)Upload(bucketName, objectKey, fileName string) error {
	fmt.Println("AwsOssStore is not implement")
	return nil
}
