package sign

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"

	oshmac "crypto/hmac"

	hash "github.com/go-zoox/crypto/hash"
)

const SIGN_ALGORITHM = "TC3-HMAC-SHA256"

type SignConfig struct {
	SecretId  string
	SecretKey string
	// Timestamp, unit: second
	Timestamp int64
	//
	Service string
	Action  string
	Region  string
	//
	RequestHost    string
	RequestMethod  string
	RequestURI     string
	RequestQuery   map[string]string
	RequestHeaders map[string]string
	RequestPayload map[string]interface{}
}

type SignResult struct {
	Signature       string
	CredentialScope string
	SignedHeaders   string
	RequestHeaders  map[string]string
	Timestamp       string
}

func Sign(data *SignConfig) (*SignResult, error) {
	canonicalRequestQuery := ""
	if data.RequestQuery != nil {
		canonicalRequestQueryUrlValue := url.Values{}
		for k, v := range data.RequestQuery {
			// canonicalRequestQuery += fmt.Sprintf("%s=%s&", k, v)
			// canonicalRequestQuerySlice = append(canonicalRequestQuerySlice, fmt.Sprintf("%s=%s", k, v))
			canonicalRequestQueryUrlValue.Add(k, v)
		}
		// canonicalRequestQuery = strings.Join(canonicalRequestQuerySlice, "&")
		canonicalRequestQuery = canonicalRequestQueryUrlValue.Encode()
	}

	// @TODO
	requestHeaders := map[string]string{}
	if data.RequestHeaders != nil {
		for k, v := range data.RequestHeaders {
			requestHeaders[strings.ToLower(k)] = v
		}
	}
	if requestHeaders["host"] == "" {
		requestHeaders["host"] = data.RequestHost
	}
	if requestHeaders["content-type"] == "" {
		if data.RequestMethod == "GET" {
			requestHeaders["content-type"] = "application/x-www-form-urlencoded"
		} else {
			requestHeaders["content-type"] = "application/json"
		}
	}

	// @TODO
	headerKeys := []string{}
	for k, _ := range requestHeaders {
		headerKeys = append(headerKeys, k)
	}
	sort.Strings(headerKeys)
	canonicalSignHeaders := strings.Join(headerKeys, ";")

	canonicalRequestHeaders := ""
	for _, key := range headerKeys {
		canonicalRequestHeaders += fmt.Sprintf("%s:%s\n", key, requestHeaders[key])
	}

	canonicalRequestPayload := ""
	if data.RequestPayload != nil {
		canonicalRequestPayloadBytes, err := json.Marshal(data.RequestPayload)
		if err != nil {
			return nil, err
		}

		canonicalRequestPayload = string(canonicalRequestPayloadBytes)
	}

	hashedCanonicalRequestPayload := hash.Sha256(canonicalRequestPayload)

	canonicalRequest := strings.Join([]string{
		data.RequestMethod,
		data.RequestURI,
		canonicalRequestQuery,
		canonicalRequestHeaders,
		canonicalSignHeaders,
		hashedCanonicalRequestPayload,
	}, "\n")

	if DEBUG {
		fmt.Println("[SIGN][START]:")
		fmt.Printf("canonicalRequest:\n%s\n\n", canonicalRequest)
	}

	date := time.Unix(data.Timestamp, 0).UTC().Format("2006-01-02")
	timestampStr := fmt.Sprintf("%d", data.Timestamp)

	hashedCanonicalRequest := hash.Sha256(canonicalRequest)
	credentialScope := fmt.Sprintf("%s/%s/tc3_request", date, data.Service)
	stringToSign := strings.Join([]string{
		SIGN_ALGORITHM,
		timestampStr,
		credentialScope,
		hashedCanonicalRequest,
	}, "\n")

	if DEBUG {
		fmt.Printf("stringToSign:\n%s\n\n", stringToSign)
		fmt.Println("[SIGN][END]")
		fmt.Println("")
	}

	// @TODO
	secretDate := hmacsha256("TC3"+data.SecretKey, date)
	secretService := hmacsha256(secretDate, data.Service)
	secretSigning := hmacsha256(secretService, "tc3_request")

	signature := fmt.Sprintf("%x", hmacsha256(secretSigning, stringToSign))

	return &SignResult{
		Signature:       signature,
		CredentialScope: credentialScope,
		SignedHeaders:   canonicalSignHeaders,
		RequestHeaders:  requestHeaders,
		Timestamp:       timestampStr,
	}, nil
}

// @TODO
func hmacsha256(key string, s string) string {
	// return hmac.Sha256(key, s)
	hashed := oshmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return string(hashed.Sum(nil))
}
