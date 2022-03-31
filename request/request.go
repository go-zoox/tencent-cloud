package request

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/go-zoox/fetch"
	"github.com/go-zoox/tencent-cloud/sign"
)

type RequestConfig struct {
	SecretId  string
	SecretKey string
	//
	Version string
	//
	Service string
	Action  string
	Region  string
	//
	RequestMethod string
	RequestURI    string
	//
	// RequestHeaders map[string]string
	RequestQuery map[string]string
	RequestBody  map[string]interface{}
}

func Request(config *RequestConfig) (*fetch.Response, error) {
	timestamp := time.Now().Unix()
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
	if config.RequestQuery != nil {
		for k, v := range config.RequestQuery {
			requestQuery[k] = v
		}
	}

	headers, err := sign.GetHeaders(&sign.SignConfig{
		SecretId:       config.SecretId,
		SecretKey:      config.SecretKey,
		Timestamp:      timestamp,
		Version:        config.Version,
		Service:        config.Service,
		Action:         config.Action,
		Region:         config.Region,
		RequestHost:    parsedRequestURI.Host,
		RequestMethod:  config.RequestMethod,
		RequestURI:     requestPath,
		RequestQuery:   requestQuery,
		RequestHeaders: nil, // config.RequestHeaders,
		RequestPayload: config.RequestBody,
	})
	if err != nil {
		return nil, err
	}

	if DEBUG {
		headersString, _ := json.MarshalIndent(headers, "", "  ")
		requestQueryString, _ := json.MarshalIndent(requestQuery, "", "  ")
		fmt.Println("[REQUEST][START]")
		fmt.Println("request method:", config.RequestMethod)
		fmt.Println("request uri:", config.RequestURI)
		fmt.Println("request query:", string(requestQueryString))
		fmt.Println("request headers:", string(headersString))
		fmt.Println("[REQUEST][END]")
		fmt.Println("")
	}

	f := fetch.New(&fetch.Config{
		Url:     config.RequestURI,
		Method:  config.RequestMethod,
		Headers: headers,
		Query:   requestQuery,
		Body:    config.RequestBody,
	})
	response, err := f.Send()
	if err != nil {
		return nil, err
	}

	if DEBUG {
		fmt.Println("[RESPONSE][START]")
		fmt.Println("response status:", response.Status)
		fmt.Println("response headers:", response.Headers)
		fmt.Println("response body:", response.String())
		fmt.Println("[RESPONSE][END]")
		fmt.Println("")
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

func Get(service string, action string, config *RequestConfig) (*fetch.Response, error) {
	return Request(&RequestConfig{
		SecretId:  config.SecretId,
		SecretKey: config.SecretKey,
		//
		Version: config.Version,
		//
		Service: service,
		Action:  action,
		Region:  config.Region,
		//
		RequestMethod: "GET",
		RequestURI:    config.RequestURI,
		//
		// RequestHeaders: config.RequestHeaders,
		RequestQuery: config.RequestQuery,
		RequestBody:  nil,
	})
}

func Post(service string, action string, data map[string]interface{}, config *RequestConfig) (*fetch.Response, error) {
	return Request(&RequestConfig{
		SecretId:  config.SecretId,
		SecretKey: config.SecretKey,
		//
		Version: config.Version,
		//
		Service: service,
		Action:  action,
		Region:  config.Region,
		//
		RequestMethod: "POST",
		RequestURI:    config.RequestURI,
		//
		// RequestHeaders: config.RequestHeaders,
		RequestQuery: config.RequestQuery,
		RequestBody:  data,
	})
}
