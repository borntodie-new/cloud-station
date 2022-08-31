package etc

type TestConfig struct {
	Endpoint     string `mapstructure:"end_point"`
	AccessKey    string `mapstructure:"access_key"`
	AccessSecret string `mapstructure:"access_secret"`
	BucketName   string `mapstructure:"bucket_name"`
}
