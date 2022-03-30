package cvm

const CVM_HOST = "cvm.tencentcloudapi.com"
const CVM_SERVICE = "cvm"

type CvmConfig struct {
	// *tencentcloud.TencentCloudClientConfig
	SecretId  string
	SecretKey string
	Region    string
}

type CvmService struct {
	config *CvmConfig
}

func New(config *CvmConfig) *CvmService {
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
