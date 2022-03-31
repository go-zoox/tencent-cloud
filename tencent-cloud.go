package tencentcloud

import (
	"github.com/go-zoox/tencent-cloud/config"
	"github.com/go-zoox/tencent-cloud/cvm"
)

type TencentCloudClient struct {
	Config *TencentCloudClientConfig
	Cvm    *cvm.CvmService
}

type TencentCloudClientConfig struct {
	SecretId  string
	SecretKey string
	//
	Region string
}

func New(config *TencentCloudClientConfig) *TencentCloudClient {
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

	return &TencentCloudClient{
		Config: config,
		// @TODO
		Cvm: cvm.New(&cvm.CvmConfig{
			SecretId:  config.SecretId,
			SecretKey: config.SecretKey,
			Region:    config.Region,
		}),
	}
}

func LoadConfig(cfg interface{}) error {
	return config.LoadConfig(cfg)
}
