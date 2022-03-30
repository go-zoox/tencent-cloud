package sign

import (
	"testing"
)

func TestSign(t *testing.T) {
	config := &SignConfig{
		SecretId:       "aaa",
		SecretKey:      "bbb",
		Timestamp:      1648653655,
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

	signResult, err := Sign(config)
	if err != nil {
		t.Fatal(err)
	}

	if signResult.Signature != "51fcc22ee98cc3531404833980a2c5c6f0470da8f59fb74e1348756254ecf409" {
		t.Error("Expected signature 51fcc22ee98cc3531404833980a2c5c6f0470da8f59fb74e1348756254ecf409, got", signResult.Signature)
	}
}
