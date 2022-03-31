package main

import (
	log "github.com/go-zoox/logger"

	tencentcloud "github.com/go-zoox/tencent-cloud"
)

func main() {
	var config tencentcloud.TencentCloudClientConfig
	if err := tencentcloud.LoadConfig(&config); err != nil {
		log.Fatal(err)
	}

	client := tencentcloud.New(&config)
	response, err := client.Cvm.DescribeInstances()
	if err != nil {
		log.Fatal(err)
	}

	log.Info("TotalCount:", response.TotalCount)
	for _, instance := range response.InstanceSet {
		log.Info("Uuid:", instance.Uuid)
		log.Info("InstanceId:", instance.InstanceId)
		log.Info("InstanceName:", instance.InstanceName)
		log.Info("InstanceState:", instance.InstanceState)
		log.Info("InstanceType:", instance.InstanceType)
		log.Info("InstanceChargeType:", instance.InstanceChargeType)
		log.Info("CPU:", instance.CPU)
		log.Info("Memory:", instance.Memory)
		log.Info("PrivateIpAddresses:", instance.PrivateIpAddresses)
		log.Info("PublicIpAddresses:", instance.PublicIpAddresses)
		log.Info("")
	}
}
