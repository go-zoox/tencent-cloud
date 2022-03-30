package sign

import "fmt"

type AuthorizationConfig struct {
	SecretId        string
	CredentialScope string
	SignedHeaders   string
	Signature       string
}

func GetAuthorization(config *AuthorizationConfig) string {
	algorithm := SIGN_ALGORITHM
	secretId := config.SecretId

	credentialScope := config.CredentialScope
	signedHeaders := config.SignedHeaders
	signature := config.Signature

	authorization := fmt.Sprintf(
		"%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		algorithm, secretId, credentialScope, signedHeaders, signature,
	)
	return authorization
}
