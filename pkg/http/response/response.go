package response

import "net/http"

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
	Write(w http.ResponseWriter)
}

type reply struct {
	code int
	body []byte
}

func Reply(code int, body []byte) RespReply {
	return &reply{
		code: code,
		body: body,
	}
}

func (r reply) Write(w http.ResponseWriter) {
	w.WriteHeader(r.code)
	if len(r.body) > 0 {
		w.Write(r.body)
	}
}
