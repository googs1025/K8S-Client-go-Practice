package api

import (
	"fmt"
	"net/http"
)

var OK = Result{Code: 0, HTTPCode: http.StatusOK, Msg: "ok"}

var (
	InternalErr = &Result{Code: 10000, HTTPCode: http.StatusInternalServerError, Msg: "internal error"}
)




type Result struct {
	Code	 int
	HTTPCode int
	Msg 	 string
	Data     interface{}
}

func (r Result) Error() string {
	return fmt.Sprintf("Error[code=%d, msg=%s]", r.Code, r.Msg)
}

func (r Result) FindMsg(msg string) *Result {
	r.Msg = msg
	return &r
}

func (r Result) FindData(data interface{}) *Result {
	r.Data = data
	return &r
}
