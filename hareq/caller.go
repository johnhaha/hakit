package hareq

import (
	"encoding/json"
	"errors"
	"log"
)

type Caller struct {
	Url    string
	Body   interface{}
	Data   []byte
	Err    error
	Param  []string
	Auth   string
	Header map[string]string
}

func NewCaller(url string) *Caller {
	return &Caller{Url: url}
}

func (caller *Caller) SetParam(param []string) *Caller {
	caller.Param = param
	return caller
}

func (caller *Caller) SetBody(body interface{}) *Caller {
	caller.Body = body
	return caller
}

func (caller *Caller) SetHeader(header map[string]string) *Caller {
	caller.Header = header
	return caller
}

func (caller *Caller) SetAuth(auth string) *Caller {
	caller.Auth = auth
	return caller
}

func (caller *Caller) GetUrl() string {
	if len(caller.Param) == 0 {
		return caller.Url
	}
	var ot string = caller.Url
	for _, p := range caller.Param {
		ot += "/" + p
	}
	return ot
}

func (caller *Caller) GetHeader() map[string]string {
	h := caller.Header
	if h == nil {
		h = make(map[string]string)
	}
	if caller.Auth != "" {
		h["Authorization"] = caller.Auth
	}
	if _, ok := h["Content-Type"]; !ok {
		h["Content-Type"] = "application/json; charset=utf-8"
	}
	return h
}

func (caller *Caller) Get() *Caller {
	data, err := Get(caller.GetUrl(), caller.GetHeader())
	if err != nil {
		log.Println("err happened on get call", err)
		caller.Err = err
		return caller
	}
	caller.Data = data
	return caller
}

func (caller *Caller) Post() *Caller {
	data, err := Post(caller.GetUrl(), caller.Body, caller.GetHeader())
	if err != nil {
		// panic(err)
		log.Println("err happened on post call", err)
		caller.Err = err
		return caller
	}
	caller.Data = data
	return caller
}

func (caller *Caller) Decode(res interface{}) *Caller {
	if caller.Data == nil {
		log.Println("nothing to decode")
		caller.Err = errors.New("nothing to decode")
		return caller
	}
	err := json.Unmarshal(caller.Data, res)
	caller.Err = err
	return caller
}

func (caller *Caller) OnErr(handler func(error)) *Caller {
	if caller.Err != nil {
		handler(caller.Err)
	}
	return caller
}

func (caller *Caller) GetData() string {
	return string(caller.Data)
}
