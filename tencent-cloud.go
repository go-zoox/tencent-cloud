package tencentcloud

import (
	"github.com/go-zoox/tencent-cloud/config"
	"github.com/go-zoox/tencent-cloud/cvm"
	"github.com/go-zoox/tencent-cloud/lighthouse"
)

type TencentCloudClient struct {
	Config     *TencentCloudClientConfig
	Cvm        *cvm.CvmService
	Lighthouse *lighthouse.LighthouseService
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
		Cvm: cvm.New(&cvm.Config{
			SecretId:  config.SecretId,
			SecretKey: config.SecretKey,
			Region:    config.Region,
		}),
		Lighthouse: lighthouse.New(&lighthouse.Config{
			SecretId:  config.SecretId,
			SecretKey: config.SecretKey,
			Region:    config.Region,
		}),
	}
}

func LoadConfig(cfg interface{}) error {
	return config.LoadConfig(cfg)
}
