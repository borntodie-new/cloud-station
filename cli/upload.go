package cli

import (
	"fmt"
	"gitee.com/go-course/cloud-station-g7/store/aliyun"
	"gitee.com/go-course/cloud-station-g7/store/aws"
	"gitee.com/go-course/cloud-station-g7/store/tx"
	"github.com/spf13/cobra"
)

var (
	uploadCmd = &cobra.Command{
		Use:   "upload",
		Short: "上传文件到中转站",
		Long:  `上传文件到中转站`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// 在这里做厂商判断分发
			err := dispatch()
			if err != nil {
				fmt.Println("厂商分发失败：", err.Error())
				return err
			}
			return nil
		},
	}
)
var (
	endPoint   = ""
	accessKey  = ""
	secretKey  = ""
	uploadFile = ""
	bucketName = ""
	ossType    = ""
)

func dispatch() error {
	option := aliyun.Option{
		EndPoint:  endPoint,
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
	switch ossType {
	case "aliyun":
		ali, err := aliyun.NewAliOssStore(option)
		if err != nil {
			return err
		}
		if err := ali.Upload(bucketName, uploadFile, uploadFile); err != nil {
			return err
		}
	case "tx":
		tx, err := tx.NewTxOssStore(option)
		if err != nil {
			return err
		}
		if err := tx.Upload(bucketName, uploadFile, uploadFile); err != nil {
			return err
		}
	case "aws":
		aws, err := aws.NewAwsOssStore(option)
		if err != nil {
			return err
		}
		if err := aws.Upload(bucketName, uploadFile, uploadFile); err != nil {
			return err
		}
	default:
		return fmt.Errorf("选择的厂商[%s]不在我的服务范围内", ossType)
	}
	return nil
}

/*
1. end_point
2. access_key
3. secret_key
4. upload_file
5. bucket_name
*/

func init() {
	f := uploadCmd.PersistentFlags()
	f.StringVarP(&endPoint, "end_point", "e", "oss-cn-guangzhou.aliyuncs.com", "默认端点")
	f.StringVarP(&accessKey, "access_key", "a", "", "OSS服务的accessKey")
	f.StringVarP(&secretKey, "secret_key", "s", "", "OSS服务的secretKey")
	f.StringVarP(&uploadFile, "upload_file", "u", "", "需要上传的文件")
	f.StringVarP(&bucketName, "bucket_name", "b", "cloud-state", "文件上传的桶")
	f.StringVarP(&ossType, "oss_type", "o", "aliyun", "OSS服务厂商[aliyun/tx/aws]")
	RootCmd.AddCommand(uploadCmd)
}
