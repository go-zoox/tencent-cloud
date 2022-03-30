package sign

import (
	"testing"
)

func TestGetAuthorization(t *testing.T) {
	config := &AuthorizationConfig{
		SecretId:        "aaa",
		CredentialScope: "2022-03-30/cvm/tc3_request",
		SignedHeaders:   "content-type;host",
		Signature:       "51fcc22ee98cc3531404833980a2c5c6f0470da8f59fb74e1348756254ecf409",
	}

	authorization := GetAuthorization(config)

	expected := "TC3-HMAC-SHA256 Credential=aaa/2022-03-30/cvm/tc3_request, SignedHeaders=content-type;host, Signature=51fcc22ee98cc3531404833980a2c5c6f0470da8f59fb74e1348756254ecf409"
	if authorization != expected {
		t.Fatalf("Expected authorization %s, got %s", expected, authorization)
	}
}
