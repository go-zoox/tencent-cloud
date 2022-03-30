package request

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/go-zoox/fetch"
	"github.com/go-zoox/tencent-cloud/sign"
)

type Config struct {
	SecretId  string
	SecretKey string
	//
	Region string
	//
	RequestURI string
}

func Get(service string, action string, config *Config) (*fetch.Response, error) {
	timestamp := time.Now().Unix()
	method := "GET"
	parsedRequestURI, error := url.Parse(config.RequestURI)
	if error != nil {
		return nil, error
	}

	requestPath := parsedRequestURI.Path
	if requestPath == "" {
		requestPath = "/"
	}

	requestQuery := map[string]string{}
	if parsedRequestURI.RawQuery != "" {
		requestQuery = make(map[string]string)
		for k, v := range parsedRequestURI.Query() {
			requestQuery[k] = v[0]
		}
	}

	headers, err := sign.GetHeaders(&sign.SignConfig{
		SecretId:       config.SecretId,
		SecretKey:      config.SecretKey,
		Timestamp:      timestamp,
		Service:        service,
		Action:         action,
		Region:         config.Region,
		RequestHost:    parsedRequestURI.Host,
		RequestMethod:  method,
		RequestURI:     requestPath,
		RequestQuery:   requestQuery,
		RequestHeaders: nil,
		RequestPayload: nil,
	})
	if err != nil {
		return nil, err
	}

	if DEBUG {
		headersString, _ := json.MarshalIndent(headers, "", "  ")
		fmt.Println("request headers:", string(headersString))
	}

	response, err := fetch.Get(config.RequestURI, &fetch.Config{
		Headers: headers,
	})
	if err != nil {
		return nil, err
	}

	if DEBUG {
		fmt.Println("response status:", response.Status)
		fmt.Println("response headers:", response.Headers)
		fmt.Println("response body:", response.String())
	}

	if response.Get("Response.Error").String() != "" {
		if DEBUG {
			fmt.Println("response error:", response.Get("Response.Error").String())
		}

		Code := response.Get("Response.Error.Code").String()
		Message := response.Get("Response.Error.Message").String()

		// if Code == "AuthFailure.SecretIdNotFound" {

		// }

		return nil, &ResponseError{Code, Message}
	}

	return response, nil
}
