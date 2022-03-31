package lighthouse

const HOST = "lighthouse.tencentcloudapi.com"
const SERVICE = "lighthouse"

type Config struct {
	// *tencentcloud.TencentCloudClientConfig
	SecretId  string
	SecretKey string
	Region    string
}

type LighthouseService struct {
	config *Config
}

func New(config *Config) *LighthouseService {
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

	return &LighthouseService{
		config: config,
	}
}
