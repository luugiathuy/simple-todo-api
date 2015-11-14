package apperror

import (
	"fmt"
	"net/http"
)

type JsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func (err *JsonErr) ToJSON(resp http.ResponseWriter) {
	resp.WriteHeader(err.Code)
	fmt.Fprintf(resp, `{"code":%d,"text":"%s"}`, err.Code, err.Text)
}
