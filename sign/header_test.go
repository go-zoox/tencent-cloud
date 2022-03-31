package sign

import "testing"

func TestGetHeaders(t *testing.T) {
	config := &SignConfig{
		SecretId:       "aaa",
		SecretKey:      "bbb",
		Timestamp:      1648653655,
		Version:        "2017-03-12",
		Service:        "cvm",
		Action:         "DescribeInstances",
		Region:         "ap-shanghai",
		RequestHost:    "cvm.tencentcloudapi.com",
		RequestMethod:  "GET",
		RequestURI:     "/",
		RequestQuery:   nil,
		RequestHeaders: nil,
		RequestPayload: nil,
	}

	headers, err := GetHeaders(config)
	if err != nil {
		t.Fatal(err)
	}

	if headers["Content-Type"] != "application/x-www-form-urlencoded" {
		t.Errorf("Expected Content-Type %s, got %s", "application/x-www-form-urlencoded", headers["Content-Type"])
	}
	if headers["Host"] != "cvm.tencentcloudapi.com" {
		t.Errorf("Expected Host %s, got %s", "cvm.tencentcloudapi.com", headers["Host"])
	}
	if headers["X-TC-Action"] != "DescribeInstances" {
		t.Errorf("Expected X-TC-Action %s, got %s", "DescribeInstances", headers["X-TC-Action"])
	}
	if headers["X-TC-Timestamp"] != "1648653655" {
		t.Errorf("Expected X-TC-Timestamp %s, got %s", "1648653655", headers["X-TC-Timestamp"])
	}
	if headers["X-TC-Version"] != config.Version {
		t.Errorf("Expected X-TC-Version %s, got %s", config.Version, headers["X-TC-Version"])
	}
	if headers["X-TC-Region"] != "ap-shanghai" {
		t.Errorf("Expected X-TC-Region %s, got %s", "ap-shanghai", headers["X-TC-Region"])
	}
	if headers["Authorization"] != "TC3-HMAC-SHA256 Credential=aaa/2022-03-30/cvm/tc3_request, SignedHeaders=content-type;host, Signature=51fcc22ee98cc3531404833980a2c5c6f0470da8f59fb74e1348756254ecf409" {
		t.Errorf("Expected Authorization %s, got %s", "TC3-HMAC-SHA256 Credential=aaa/2022-03-30/cvm/tc3_request, SignedHeaders=content-type;host, Signature=51fcc22ee98cc3531404833980a2c5c6f0470da8f59fb74e1348756254ecf409", headers["Authorization"])
	}
}
