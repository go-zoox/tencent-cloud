package cvm

import (
	"fmt"

	"github.com/go-zoox/tencent-cloud/request"
)

type DescribeInstancesRawResponse struct {
	Response DescribeInstancesResponse
}

type DescribeInstancesResponse struct {
	InstanceSet []Instance

	RequestId  string `json:"RequestId"`
	TotalCount int    `json:"TotalCount"`
}

type Instance struct {
	CPU         int    `json:"CPU"`
	CamRoleName string `json:"CamRoleName"`
	CreatedTime string `json:"CreatedTime"`
	// DataRisks					[]string `json:"DataRisks"`
	DeicayedClusterId      string `json:"DeicayedClusterId"`
	DefaultLoginPort       int    `json:"DefaultLoginPort"`
	DefaultLoginUser       string `json:"DefaultLoginUser"`
	DisableApiTermination  bool   `json:"DisableApiTermination"`
	DisasterRecoverGroupId string `json:"DisasterRecoverGroupId"`
	// ExpiredTime           time.Time `json:"ExpiredTime"`
	HpcClusterId string `json:"HpcClusterId"`
	// IPv6Addresses 				[]string `json:"IPv6Addresses"`
	ImageId            string `json:"ImageId"`
	InstanceChargeType string `json:"InstanceChargeType"`
	InstanceId         string `json:"InstanceId"`
	InstanceName       string `json:"InstanceName"`
	InstanceState      string `json:"InstanceState"`
	InstanceType       string `json:"InstanceType"`
	InternetAccessible struct {
		InternetChargeType      string `json:"InternetChargeType"`
		InternetMaxBandwidthOut int    `json:"InternetMaxBandwidthOut"`
	}
	IsolatedSource string `json:"IsolatedSource"`
	// LatestOperation 			string `json:"LatestOperation"`
	// LatestOperationRequestId 	string `json:"LatestOperationRequestId"`
	// LatestOperationState 		string `json:"LatestOperationState"`
	LoginSettings struct {
		KeyIds []string `json:"KeyIds"`
	}
	Memory    int    `json:"Memory"`
	OsName    string `json:"OsName"`
	Placement struct {
		// HostId 					string `json:"HostId"`
		ProjectId int    `json:"ProjectId"`
		Zone      string `json:"Zone"`
	}
	PlatformProjectId  string   `json:"PlatformProjectId"`
	PrivateIpAddresses []string `json:"PrivateIpAddresses"`
	PublicIpAddresses  []string `json:"PublicIpAddresses"`
	// RdmaIpAddresses 				[]string `json:"RdmaIpAddresses"`
	// RenewFlag 					string `json:"RenewFlag"`
	RestrictState    string   `json:"RestrictState"`
	SecurityGroupIds []string `json:"SecurityGroupIds"`
	StopChargingMode string   `json:"StopChargingMode"`
	SystemDisk       struct {
		// CdcId 						string `json:"CdcId"`
		DiskId   string `json:"DiskId"`
		DiskSize int    `json:"DiskSize"`
		DiskType string `json:"DiskType"`
		Encrypt  bool   `json:"Encrypt"`
		// KmsKeyId 					string `json:"KmsKeyId"`
		ThroughputPerformance int `json:"ThroughputPerformance"`
	}
	// Tags 						[]string `json:"Tags"`
	Uuid                string `json:"Uuid"`
	VirtualPrivateCloud struct {
		AsVpcGateway bool   `json:"AsVpcGateway"`
		SubnetId     string `json:"SubnetId"`
		VpcId        string `json:"VpcId"`
	}
}

type DescribeInstancesConditions struct {
	Offset int
	Limit  int
}

func (cs *CvmService) DescribeInstances(conditions ...*DescribeInstancesConditions) (*DescribeInstancesResponse, error) {
	var requestQuery map[string]string
	if len(conditions) > 0 {
		conditionsX := conditions[0]
		requestQuery = map[string]string{
			"Offset": fmt.Sprintf("%d", conditionsX.Offset),
			"Limit":  fmt.Sprintf("%d", conditionsX.Limit),
		}
	}

	requestURI := fmt.Sprintf("https://%s", CVM_HOST)
	response, err := request.Get(CVM_SERVICE, "DescribeInstances", &request.Config{
		SecretId:     cs.config.SecretId,
		SecretKey:    cs.config.SecretKey,
		Region:       cs.config.Region,
		RequestURI:   requestURI,
		RequestQuery: requestQuery,
	})
	if err != nil {
		return nil, err
	}

	// json, _ := response.JSON()
	// fmt.Println("response:", json)

	var result = &DescribeInstancesRawResponse{}
	err = response.Unmarshal(result)
	if err != nil {
		return nil, err
	}

	return &result.Response, nil
}

func (cs *CvmService) DescribeInstancesPost(conditions map[string]interface{}) (*DescribeInstancesResponse, error) {
	requestURI := fmt.Sprintf("https://%s", CVM_HOST)
	response, err := request.Post(
		CVM_SERVICE,
		"DescribeInstances",
		conditions,
		&request.Config{
			SecretId:   cs.config.SecretId,
			SecretKey:  cs.config.SecretKey,
			Region:     cs.config.Region,
			RequestURI: requestURI,
		})
	if err != nil {
		return nil, err
	}

	// json, _ := response.JSON()
	// fmt.Println("response:", json)

	var result = &DescribeInstancesRawResponse{}
	err = response.Unmarshal(result)
	if err != nil {
		return nil, err
	}

	return &result.Response, nil
}
