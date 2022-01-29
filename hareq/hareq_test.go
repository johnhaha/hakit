package hareq_test

import (
	"testing"

	"github.com/johnhaha/hakit/hareq"
)

func TestGet(t *testing.T) {
	var msg hareq.MsgRes
	err := hareq.FastGet("http://localhost:8001/api/ping", &msg)
	if err != nil {
		t.Fatal(err)
	}
	if !msg.Success {
		t.Fatal(msg)
	}
}

func TestPost(t *testing.T) {
	type Res struct {
		hareq.MsgRes
		Data string `json:"data"`
	}
	var res Res
	err := hareq.FastPost(map[string]string{
		"name": "tom",
	}, "http://localhost:8001/api/pong", &res)
	if err != nil {
		t.Fatal(err)
	}
	if !res.Success {
		t.Fatal(res)
	}
}

func TestCallerGet(t *testing.T) {
	var msg hareq.MsgRes
	caller := hareq.NewCaller("http://localhost:8001/api/ping")
	err := caller.Get().Decode(&msg)
	if err != nil {
		t.Fatal(err)
	}
	if !msg.Success {
		t.Fatal(msg)
	}
}

func TestCallerPost(t *testing.T) {
	type Res struct {
		hareq.MsgRes
		Data string `json:"data"`
	}
	var res Res
	hareq.NewCaller("http://localhost:8001/api/pong").SetBody(map[string]string{
		"name": "tom",
	}).Post().Decode(&res)
	if !res.Success {
		t.Fatal(res)
	}
}

func TestGetCallerUrl(t *testing.T) {
	caller := hareq.NewCaller("http://test")
	url := caller.SetQuery(map[string]string{"ko": "ok"}).AddQuery("fd", "sd").AddQuery("mn", "zx").GetUrl()
	t.Fatal(url)
}
