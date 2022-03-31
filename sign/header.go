package sign

func GetHeaders(config *SignConfig) (map[string]string, error) {
	signResult, err := Sign(config)
	if err != nil {
		return nil, err
	}

	authorization := GetAuthorization(&AuthorizationConfig{
		SecretId:        config.SecretId,
		CredentialScope: signResult.CredentialScope,
		SignedHeaders:   signResult.SignedHeaders,
		Signature:       signResult.Signature,
	})

	return map[string]string{
		"Content-Type":   signResult.RequestHeaders["content-type"],
		"Host":           signResult.RequestHeaders["host"],
		"X-TC-Action":    config.Action,
		"X-TC-Timestamp": signResult.Timestamp,
		"X-TC-Version":   config.Version,
		"X-TC-Region":    config.Region,
		"Authorization":  authorization,
	}, nil
}
