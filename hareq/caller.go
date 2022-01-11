package hareq

import (
	"encoding/json"
	"errors"
	"log"
)

type Caller struct {
	Url   string
	Body  interface{}
	Data  string
	Err   error
	Param []string
}

func NewCaller(url string) *Caller {
	return &Caller{Url: url}
}

func (caller *Caller) SetParam(param []string) *Caller {
	caller.Param = param
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

func (caller *Caller) Get() *Caller {
	data, err := DataGet(caller.GetUrl())
	if err != nil {
		log.Println("err happened on get call", err)
		caller.Err = err
		return caller
	}
	caller.Data = data
	return caller
}

func (caller *Caller) Decode(res interface{}) *Caller {
	if caller.Data == "" {
		log.Println("nothing to decode")
		caller.Err = errors.New("nothing to decode")
		return caller
	}
	err := json.Unmarshal([]byte(caller.Data), res)
	caller.Err = err
	return caller
}

func (caller *Caller) OnErr(handler func(error)) *Caller {
	if caller.Err != nil {
		handler(caller.Err)
	}
	return caller
}
