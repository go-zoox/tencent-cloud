package tencentcloud

import (
	"fmt"
	"os"

	"github.com/go-zoox/dotenv"
)

type TencentCloudClient struct {
	Config *TencentCloudClientConfig
}

type TencentCloudClientConfig struct {
	SecretId  string
	SecretKey string
	//
	Region string
}

func New(config *TencentCloudClientConfig) *TencentCloudClient {
	return &TencentCloudClient{
		Config: config,
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
