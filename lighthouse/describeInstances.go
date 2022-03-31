package lighthouse

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

	// @TODO
	Raw []byte
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
	PlatformProjectId string   `json:"PlatformProjectId"`
	PrivateAddresses  []string `json:"PrivateAddresses"`
	PublicAddresses   []string `json:"PublicAddresses"`
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

func (cs *LighthouseService) DescribeInstances(conditions ...*DescribeInstancesConditions) (*DescribeInstancesResponse, error) {
	var requestQuery map[string]string
	if len(conditions) > 0 {
		conditionsX := conditions[0]
		requestQuery = map[string]string{
			"Offset": fmt.Sprintf("%d", conditionsX.Offset),
			"Limit":  fmt.Sprintf("%d", conditionsX.Limit),
		}
	}

	requestURI := fmt.Sprintf("https://%s", HOST)
	response, err := request.Get(SERVICE, "DescribeInstances", &request.RequestConfig{
		SecretId:     cs.config.SecretId,
		SecretKey:    cs.config.SecretKey,
		Version:      "2020-03-24",
		Region:       cs.config.Region,
		RequestURI:   requestURI,
		RequestQuery: requestQuery,
	})
	if err != nil {
		return nil, err
	}

	var result = &DescribeInstancesRawResponse{}
	err = response.UnmarshalJSON(result)
	// @TODO
	result.Response.Raw = response.Body
	if err != nil {
		return nil, err
	}

	return &result.Response, nil
}

func (cs *LighthouseService) DescribeInstancesPost(conditions map[string]interface{}) (*DescribeInstancesResponse, error) {
	requestURI := fmt.Sprintf("https://%s", HOST)
	response, err := request.Post(
		SERVICE,
		"DescribeInstances",
		conditions,
		&request.RequestConfig{
			SecretId:   cs.config.SecretId,
			SecretKey:  cs.config.SecretKey,
			Version:    "2020-03-24",
			Region:     cs.config.Region,
			RequestURI: requestURI,
		})
	if err != nil {
		return nil, err
	}

	var result = &DescribeInstancesRawResponse{}
	err = response.UnmarshalJSON(result)
	if err != nil {
		return nil, err
	}

	return &result.Response, nil
}
