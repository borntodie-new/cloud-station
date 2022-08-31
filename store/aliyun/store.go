package aliyun

import (
	"fmt"
	"gitee.com/go-course/cloud-station-g7/store"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

var (
	// 对象是否实现了接口的约束
	_ store.Uploader = &AliOssStore{}
)

// 处理异常
func handleErr(msg string, err error) {
	fmt.Println(msg, err.Error())
	os.Exit(1)
}

type AliOssStore struct {
	client   *oss.Client // oss对象存储客户端
	listener oss.ProgressListener
}

func NewAliOssStore(option Option) (*AliOssStore, error) {
	if err := option.validate(); err != nil {
		return nil, err
	}
	c, err := oss.New(option.EndPoint, option.AccessKey, option.SecretKey)
	if err != nil {
		return nil, err
	}
	listener := NewOssProgressListener()
	return &AliOssStore{
		client:   c,
		listener: listener,
	}, nil
}

func (s *AliOssStore) Upload(bucketName, objectKey, fileName string) error {

	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		handleErr("获取bucket对象失败：", err)
	}

	if err := bucket.PutObjectFromFile(objectKey, fileName, oss.Progress(s.listener)); err != nil {
		handleErr("上传文件失败：", err)
	}

	signURL, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60*24)
	if err != nil {
		handleErr("获取文件下载地址失败：", err)
	}
	fmt.Printf("文件：【%s】的下载地址是：【%s】\n", fileName, signURL)
	return nil
}

type Option struct {
	EndPoint  string
	AccessKey string
	SecretKey string
}

func (o *Option) validate() error {
	if o.EndPoint == "" {
		return fmt.Errorf("endpoint 参数缺失")
	}
	if o.AccessKey == "" {
		return fmt.Errorf("accessKey 参数缺失")
	}
	if o.SecretKey == "" {
		return fmt.Errorf("secretKey 参数缺失")
	}
	return nil
}
