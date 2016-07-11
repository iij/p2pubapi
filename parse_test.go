package p2pubapi

import (
	"testing"

	"github.com/iij/p2pubapi/protocol"
)

type TCommon struct {
	Hello   string `json:"-"`
	World   string `json:"-"`
	InJson  string
	InParam string `json:"-" p2pub:",query"`
}

func (t TCommon) URI() string {
	return "/a/p/i/{{.Hello}}/{{.World}}.json"
}

func (t TCommon) APIName() string {
	return "helloworld"
}

func (t TCommon) JPName() string {
	return "はろう"
}

func (t TCommon) Method() string {
	return "GET"
}

func (t TCommon) Document() string {
	return "http://www.example.com"
}

func Test1(t *testing.T) {
	t.Log("hello")
	cmn := TCommon{Hello: "hello", World: "world", InJson: "in-json"}
	cmn.Hello = "hello"
	cmn.World = "world"
	cmn.InJson = "test value"
	cmn.InParam = "test param"
	api := NewAPI("accesskey", "secretkey")
	t.Log("body", GetBody(cmn))
	t.Log("path", GetPath(cmn))
	param := GetParam(*api, cmn)
	t.Log("param", param)
}

func TestAPIList(t *testing.T) {
	for _, a := range protocol.APIlist {
		t.Logf("%v %v %v", a.Method(), a.APIName(), a.URI())
		req, opt := ArgumentList(a)
		t.Log("req", req, "opt", opt)
	}
	arg := protocol.VMGet{}
	err := Validate(arg)
	t.Log(err)
	t.Log(ArgumentList(arg))
}
