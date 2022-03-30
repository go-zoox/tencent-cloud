package cvm

import (
	"testing"

	tencentcloud "github.com/go-zoox/tencent-cloud"
)

func TestDescribeInstances(t *testing.T) {
	var config CvmConfig
	if err := tencentcloud.LoadConfig(&config); err != nil {
		t.Fatal(err)
	}

	client := New(&CvmConfig{
		SecretId:  config.SecretId,
		SecretKey: config.SecretKey,
		Region:    config.Region,
	})

	response, err := client.DescribeInstances(&DescribeInstancesConditions{
		Offset: 0,
		Limit:  10,
	})
	if err != nil {
		t.Fatal(err)
	}

	// xxx, _ := json.MarshalIndent(response, "", "  ")
	// t.Log(string(xxx))
	for _, instance := range response.InstanceSet {
		t.Log("Uuid:", instance.Uuid)
		t.Log("InstanceId:", instance.InstanceId)
		t.Log("InstanceName:", instance.InstanceName)
		t.Log("InstanceState:", instance.InstanceState)
		t.Log("InstanceType:", instance.InstanceType)
		t.Log("InstanceChargeType:", instance.InstanceChargeType)
		t.Log("CPU:", instance.CPU)
		t.Log("Memory:", instance.Memory)
		t.Log("PrivateIpAddresses:", instance.PrivateIpAddresses)
		t.Log("PublicIpAddresses:", instance.PublicIpAddresses)
		t.Log("")
	}
}

func TestDescribeInstancesPost(t *testing.T) {
	var config CvmConfig
	if err := tencentcloud.LoadConfig(&config); err != nil {
		t.Fatal(err)
	}

	client := New(&CvmConfig{
		SecretId:  config.SecretId,
		SecretKey: config.SecretKey,
		Region:    config.Region,
	})

	response, err := client.DescribeInstancesPost(map[string]interface{}{
		"Limit": 1,
		"Filters": []map[string]interface{}{
			{
				"Values": []string{"\u672a\u547d\u540d"},
				"Name":   "instance-name",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	// xxx, _ := json.MarshalIndent(response, "", "  ")
	// t.Log(string(xxx))
	for _, instance := range response.InstanceSet {
		t.Log("Uuid:", instance.Uuid)
		t.Log("InstanceId:", instance.InstanceId)
		t.Log("InstanceName:", instance.InstanceName)
		t.Log("InstanceState:", instance.InstanceState)
		t.Log("InstanceType:", instance.InstanceType)
		t.Log("InstanceChargeType:", instance.InstanceChargeType)
		t.Log("CPU:", instance.CPU)
		t.Log("Memory:", instance.Memory)
		t.Log("PrivateIpAddresses:", instance.PrivateIpAddresses)
		t.Log("PublicIpAddresses:", instance.PublicIpAddresses)
		t.Log("")
	}
}