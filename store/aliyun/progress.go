package aliyun

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/schollz/progressbar/v3"
)

type OssProgressListener struct {
	bar *progressbar.ProgressBar
}

func NewOssProgressListener() *OssProgressListener {
	return &OssProgressListener{}
}

func (p *OssProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	//fmt.Printf("%d  %d  %d  %d\n", event.EventType, event.TotalBytes, event.RwBytes, event.ConsumedBytes)
	switch event.EventType {
	case oss.TransferStartedEvent: // 开始上传
		p.bar = progressbar.DefaultBytes(
			event.TotalBytes,
			"文件上传中",
		)
	case oss.TransferDataEvent: // 上传中
		_ = p.bar.Add64(event.RwBytes)
	case oss.TransferCompletedEvent: // 上传完成
		fmt.Println("上传完成")
	case oss.TransferFailedEvent: // 上传失败
		fmt.Println("上传失败")
	default:
	}
}
