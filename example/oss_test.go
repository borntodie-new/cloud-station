package example

import (
	"fmt"
	"gitee.com/go-course/cloud-station-g7/etc"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
	"os"
	"testing"
)

var (
	client     *oss.Client
	testConfig etc.TestConfig
)

func init() {
	// 1. 实例化全局的配置文件对象
	v := viper.New()
	v.SetConfigFile("../etc/test_config.yaml")
	if err := v.ReadInConfig(); err != nil {
		fmt.Println("读取配置文件失败：", err.Error())
		os.Exit(1)
	}
	if err := v.Unmarshal(&testConfig); err != nil {
		fmt.Println("解析配置文件失败：", err.Error())
		os.Exit(1)
	}

	// 2. 初始化oss服务实例
	c, err := oss.New(testConfig.Endpoint, testConfig.AccessKey, testConfig.AccessSecret)
	if err != nil {
		fmt.Println("实例化oss服务实例失败：", err.Error())
		os.Exit(1)
	}
	client = c

}

// 测试阿里云OSSSDK的基本使用
func TestBucketList(t *testing.T) {

	lsRes, err := client.ListBuckets()
	if err != nil {
		fmt.Println("获取bucket列表失败：", err.Error())
		os.Exit(1)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}
}


func TestUpload(t *testing.T) {
	bucket, err := client.Bucket(testConfig.BucketName)
	if err != nil {
		fmt.Println("测试上传文件失败：", err.Error())
		os.Exit(1)
	}
	if err := bucket.PutObjectFromFile("dir/oss_test.go", "oss_test.go"); err != nil {
		fmt.Println("上传文件失败：", err.Error())
		os.Exit(1)
	}
}