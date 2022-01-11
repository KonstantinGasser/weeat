package response

import (
	"encoding/json"
	"net/http"
)

type RespErr interface {
	Code() int
	Msg() string
	Err() error
	Write(w http.ResponseWriter)
}

type err struct {
	code int
	err  error
	msg  string
}

func (e err) Code() int   { return e.code }
func (e err) Msg() string { return e.msg }
func (e err) Err() error  { return e.err }

func (e err) Write(w http.ResponseWriter) {
	http.Error(w, e.msg, e.code)
}

func Err(e error, code int, msg string) RespErr {
	return &err{
		code: code,
		msg:  msg,
		err:  e,
	}
}

type RespReply interface {
	Write(code int, body []byte)
	JSON(code int, data interface{})
}

type reply struct {
	w    http.ResponseWriter
	code int
}

func Reply(w http.ResponseWriter) RespReply {
	return &reply{
		w: w,
	}
}

func (r reply) Write(code int, body []byte) {
	r.w.WriteHeader(code)
	if body != nil && len(body) > 0 {
		r.w.Write(body)
	}
}

func (r reply) JSON(code int, data interface{}) {
	r.w.Header().Set("content-type", "application/json")
	json.NewEncoder(r.w).Encode(map[string]interface{}{
		"status": r.code,
		"data":   data,
	})
}
