package lighthouse

import (
	"testing"

	config "github.com/go-zoox/tencent-cloud/config"
)

func TestDescribeInstances(t *testing.T) {
	var cfg Config
	if err := config.LoadConfig(&cfg); err != nil {
		t.Fatal(err)
	}

	client := New(&Config{
		SecretId:  cfg.SecretId,
		SecretKey: cfg.SecretKey,
		Region:    cfg.Region,
	})

	response, err := client.DescribeInstances(&DescribeInstancesConditions{
		Offset: 0,
		Limit:  10,
	})
	if err != nil {
		t.Fatal(err)
	}

	// fmt.Println("Raw:", string(response.Raw))

	t.Log("TotalCount:", response.TotalCount)
	for _, instance := range response.InstanceSet {
		t.Log("Uuid:", instance.Uuid)
		t.Log("InstanceId:", instance.InstanceId)
		t.Log("InstanceName:", instance.InstanceName)
		t.Log("InstanceState:", instance.InstanceState)
		t.Log("InstanceType:", instance.InstanceType)
		t.Log("InstanceChargeType:", instance.InstanceChargeType)
		t.Log("CPU:", instance.CPU)
		t.Log("Memory:", instance.Memory)
		t.Log("PrivateIpAddresses:", instance.PrivateAddresses)
		t.Log("PublicIpAddresses:", instance.PublicAddresses)
		t.Log("")
	}
}

// func TestDescribeInstancesPost(t *testing.T) {
// 	var cfg Config
// 	if err := config.LoadConfig(&cfg); err != nil {
// 		t.Fatal(err)
// 	}

// 	client := New(&Config{
// 		SecretId:  cfg.SecretId,
// 		SecretKey: cfg.SecretKey,
// 		Region:    cfg.Region,
// 	})

// 	response, err := client.DescribeInstancesPost(map[string]interface{}{
// 		"Limit": 1,
// 		"Filters": []map[string]interface{}{
// 			{
// 				"Values": []string{"\u672a\u547d\u540d"},
// 				"Name":   "instance-name",
// 			},
// 		},
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Log("TotalCount:", response.TotalCount)
// 	for _, instance := range response.InstanceSet {
// 		t.Log("Uuid:", instance.Uuid)
// 		t.Log("InstanceId:", instance.InstanceId)
// 		t.Log("InstanceName:", instance.InstanceName)
// 		t.Log("InstanceState:", instance.InstanceState)
// 		t.Log("InstanceType:", instance.InstanceType)
// 		t.Log("InstanceChargeType:", instance.InstanceChargeType)
// 		t.Log("CPU:", instance.CPU)
// 		t.Log("Memory:", instance.Memory)
// 		t.Log("PrivateIpAddresses:", instance.PrivateIpAddresses)
// 		t.Log("PublicIpAddresses:", instance.PublicIpAddresses)
// 		t.Log("")
// 	}
// }
