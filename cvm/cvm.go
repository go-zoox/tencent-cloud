package cvm

const HOST = "cvm.tencentcloudapi.com"
const SERVICE = "cvm"

type Config struct {
	// *tencentcloud.TencentCloudClientConfig
	SecretId  string
	SecretKey string
	Region    string
}

type CvmService struct {
	config *Config
}

func New(config *Config) *CvmService {
	if config == nil {
		panic("config is nil")
	}

	if config.SecretId == "" {
		panic("config.SecretId is empty")
	}

	if config.SecretKey == "" {
		panic("config.SecretKey is empty")
	}

	if config.Region == "" {
		panic("config.Region is empty")
	}

	return &CvmService{
		config: config,
	}
}
