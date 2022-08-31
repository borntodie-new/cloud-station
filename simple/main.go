package main

import (
	"flag"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

var (
	endpoint   = ""
	accessKey  = ""
	secretKey  = ""
	bucketName = ""
	uploadFile = ""
	help       = false
)


func handleErr(msg string, err error) {
	fmt.Println(msg, err.Error())
	os.Exit(1)
}

// 实现文件上传功能
func upload() error {
	// 1. 创建oss对象实例
	client, err := oss.New(endpoint, accessKey, secretKey)
	if err != nil {
		handleErr("创建oss对象实例失败：", err)
	}
	// 2. 获取bucket对象
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleErr("获取bucket对象失败：", err)
	}

	// 3. 上传文件
	if err := bucket.PutObjectFromFile(uploadFile, uploadFile); err != nil {
		handleErr("上传文件失败：", err)
	}

	// 4. 输出文件下载地址
	signURL, err := bucket.SignURL(uploadFile, oss.HTTPGet, 60*60*24)
	if err != nil {
		handleErr("获取文件下载地址失败：", err)
	}
	fmt.Printf("文件：【%s】的下载地址是：【%s】\n", uploadFile, signURL)
	return nil
}

// 实现参数校验
func validate() error {
	if endpoint == "" {
		return fmt.Errorf("endpoint 参数缺失")
	}
	if accessKey == "" {
		return fmt.Errorf("accessKey 参数缺失")
	}
	if secretKey == "" {
		return fmt.Errorf("secretKey 参数缺失")
	}
	if bucketName == "" {
		return fmt.Errorf("bucketName 参数不能缺失")
	}
	return nil
}

// 参数加载
func loadParams() {
	flag.BoolVar(&help, "h", false, "输出使用文档")
	flag.StringVar(&uploadFile, "f", "", "需要上传的文件名")
	flag.Parse()

	if help {
		usage()
		os.Exit(0)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `cloud-state version: 0.0.1
Usage: cloud-station [-h] -f <uplaod_file_path>
Options:
`)
	flag.PrintDefaults() // 输出参数信息
}

func main() {

	// 参数加载
	loadParams()

	// 参数校验
	if err := validate(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := upload(); err != nil {
		handleErr("upload函数执行失败：", err)
	}
	fmt.Printf("上传文件【%s】成功\n", uploadFile)
}
