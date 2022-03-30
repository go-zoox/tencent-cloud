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
	//
	RequestQuery map[string]string
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
	if config.RequestQuery != nil {
		for k, v := range config.RequestQuery {
			requestQuery[k] = v
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
		requestQueryString, _ := json.MarshalIndent(requestQuery, "", "  ")
		fmt.Println("[REQUEST][START]")
		fmt.Println("request method:", method)
		fmt.Println("request uri:", config.RequestURI)
		fmt.Println("request query:", string(requestQueryString))
		fmt.Println("request headers:", string(headersString))
		fmt.Println("[REQUEST][END]")
		fmt.Println("")
	}

	response, err := fetch.Get(config.RequestURI, &fetch.Config{
		Headers: headers,
		Query:   requestQuery,
	})
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

func Post(service string, action string, data map[string]interface{}, config *Config) (*fetch.Response, error) {
	timestamp := time.Now().Unix()
	method := "POST"
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
		RequestPayload: data,
	})
	if err != nil {
		return nil, err
	}

	if DEBUG {
		headersString, _ := json.MarshalIndent(headers, "", "  ")
		requestQueryString, _ := json.MarshalIndent(requestQuery, "", "  ")
		fmt.Println("[REQUEST][START]")
		fmt.Println("request method:", method)
		fmt.Println("request uri:", config.RequestURI)
		fmt.Println("request query:", string(requestQueryString))
		fmt.Println("request headers:", string(headersString))
		fmt.Println("[REQUEST][END]")
		fmt.Println("")
	}

	response, err := fetch.Post(config.RequestURI, &fetch.Config{
		Headers: headers,
		Body:    data,
	})
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
