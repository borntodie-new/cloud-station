package aliyun_test

import (
	"fmt"
	"gitee.com/go-course/cloud-station-g7/etc"
	"gitee.com/go-course/cloud-station-g7/store"
	"gitee.com/go-course/cloud-station-g7/store/aliyun"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	uploader   store.Uploader
	testConfig etc.TestConfig
)

func init() {
	// 1. 实例化全局的配置文件对象
	v := viper.New()
	v.SetConfigFile("../../etc/test_config.yaml")
	if err := v.ReadInConfig(); err != nil {
		fmt.Println("读取配置文件失败：", err.Error())
		os.Exit(1)
	}
	if err := v.Unmarshal(&testConfig); err != nil {
		fmt.Println("解析配置文件失败：", err.Error())
		os.Exit(1)
	}

	// 2. 初始化oss服务实例
	c, err := aliyun.NewAliOssStore(aliyun.Option{
		EndPoint:  testConfig.Endpoint,
		AccessKey: testConfig.AccessKey,
		SecretKey: testConfig.AccessSecret,
	})
	if err != nil {
		fmt.Println("实例化oss服务对象失败：", err.Error())
		os.Exit(1)
	}
	uploader = c

}

func TestUpload(t *testing.T) {
	should := assert.New(t)
	err := uploader.Upload(testConfig.BucketName, "content.txt", "store.go")
	should.NoError(err)
}
