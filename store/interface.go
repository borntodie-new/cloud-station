package store

// 定义如何上传文件到bucket
// 做了抽象，并不关心我们需要上传到那个OSS厂商
type Uploader interface {
	Upload(bucketName, objectKey, fileName string) error
}
