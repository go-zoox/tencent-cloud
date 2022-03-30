package request

import "fmt"

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (re ResponseError) Error() string {
	return fmt.Sprintf("[%s] %s", re.Code, re.Message)
}
