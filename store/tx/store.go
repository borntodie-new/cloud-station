package tx

import (
	"fmt"
	"gitee.com/go-course/cloud-station-g7/store/aliyun"
)

type TxOssStore struct {
}

func NewTxOssStore(option aliyun.Option) (*TxOssStore, error) {
	return &TxOssStore{}, nil
}
func (t *TxOssStore)Upload(bucketName, objectKey, fileName string) error {
	fmt.Println("TxOssStore is not implement")
	return nil
}

