package tencentcloud

import "testing"

func TestCvm(t *testing.T) {
	var config TencentCloudClientConfig
	if err := LoadConfig(&config); err != nil {
		t.Fatal(err)
	}

	client := New(&config)
	response, err := client.Cvm.DescribeInstances()
	if err != nil {
		t.Fatal(err)
	}

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
		t.Log("PrivateIpAddresses:", instance.PrivateIpAddresses)
		t.Log("PublicIpAddresses:", instance.PublicIpAddresses)
		t.Log("")
	}
}
