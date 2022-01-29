package hareq

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Caller struct {
	Url    string
	Body   interface{}
	Data   []byte
	Err    error
	Param  []string
	Auth   string
	Header map[string]string
	Query  map[string]string
}

func NewCaller(url string) *Caller {
	return &Caller{Url: url}
}

func (caller *Caller) SetParam(param []string) *Caller {
	caller.Param = param
	return caller
}

func (caller *Caller) SetQuery(query map[string]string) *Caller {
	caller.Query = query
	return caller
}

func (caller *Caller) AddQuery(k string, v string) *Caller {
	if caller.Query == nil {
		caller.Query = make(map[string]string)
	}
	caller.Query[k] = v
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

func (caller *Caller) SetUrl(url string) *Caller {
	caller.Url = url
	return caller
}

func (caller *Caller) GetUrl() string {
	var ot string = caller.Url
	for _, p := range caller.Param {
		ot += "/" + p
	}
	if caller.Query != nil {
		ot += "?"
		for k, v := range caller.Query {
			ot += fmt.Sprintf("%v=%v&", k, v)
		}
	}
	ot = strings.TrimRight(ot, "&")
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
		caller.Err = err
		return caller
	}
	caller.Data = data
	return caller
}

func (caller *Caller) Post() *Caller {
	data, err := Post(caller.GetUrl(), caller.Body, caller.GetHeader())
	if err != nil {
		caller.Err = err
		return caller
	}
	caller.Data = data
	return caller
}

func (caller *Caller) Put() *Caller {
	data, err := Put(caller.GetUrl(), caller.Body, caller.GetHeader())
	if err != nil {
		caller.Err = err
		return caller
	}
	caller.Data = data
	return caller
}

func (caller *Caller) Delete() *Caller {
	data, err := Delete(caller.GetUrl(), caller.GetHeader())
	if err != nil {
		caller.Err = err
		return caller
	}
	caller.Data = data
	return caller
}

func (caller *Caller) Decode(res interface{}) error {
	if caller.Err != nil {
		return caller.Err
	}
	if caller.Data == nil {
		return errors.New("nothing to decode")
	}
	err := json.Unmarshal(caller.Data, res)
	return err
}

func (caller *Caller) GetData() string {
	return string(caller.Data)
}
