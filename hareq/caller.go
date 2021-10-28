package hareq

import (
	"encoding/json"
	"errors"
	"log"
)

type Caller struct {
	Url  string
	Body interface{}
	Data string
	Err  error
}

func NewCaller(url string) *Caller {
	return &Caller{Url: url}
}

func (caller *Caller) Get() *Caller {
	data, err := DataGet(caller.Url)
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
