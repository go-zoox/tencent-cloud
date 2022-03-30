package tencentcloud

import (
	"fmt"
	"os"

	"github.com/go-zoox/dotenv"
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

func LoadConfig(config interface{}) error {
	currentPath, _ := os.Getwd()
	homedirPath := os.Getenv("HOME")
	dotenvPathsTry := []string{
		fmt.Sprintf("%s/.env", currentPath),
		fmt.Sprintf("%s/../.env", currentPath),
		fmt.Sprintf("%s/.env", homedirPath),
	}
	dotenvPaths := []string{}
	for _, dotenvPathTry := range dotenvPathsTry {
		if _, err := os.Stat(dotenvPathTry); err == nil {
			dotenvPaths = append(dotenvPaths, dotenvPathTry)
		}
	}

	if err := dotenv.Load(config, dotenvPaths...); err != nil {
		return err
	}

	return nil
}
