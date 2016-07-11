package p2pubapi

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"testing"
)

func testCall(t *testing.T) {
	t.Log("call test")
	api := NewAPI(os.Getenv("IIJAPI_ACCESS_KEY"), os.Getenv("IIJAPI_SECRET_KEY"))
	u, err := url.Parse(EndpointJSON)
	u.Path = path.Join(u.Path, "/r/"+APIVersion20151130+"/gises/"+os.Getenv("GISSERVICECODE")+".json")
	t.Log("u", u, u.RawPath, "err", err)
	resp, err := api.Get(*u)
	t.Log("resp", resp, "err", err)
	t.Log("resp header", resp.Header)
	b, err := ioutil.ReadAll(resp.Body)
	t.Log("read err", err)
	t.Log("resp body", string(b))
	var maps map[string]interface{}
	json.Unmarshal(b, &maps)
	t.Log("decoded", maps)
}

func testCall2(t *testing.T) {
	t.Log("call test")
	api := NewAPI(os.Getenv("IIJAPI_ACCESS_KEY"), os.Getenv("IIJAPI_SECRET_KEY"))
	u, err := url.Parse(EndpointJSON)
	u.Path = path.Join(u.Path, "/r/"+APIVersion20151130+"/gises.json")
	q := u.Query()
	q.Set("Item", "ServiceCode")
	u.RawQuery = q.Encode()
	t.Log("u", u, u.RawPath, "err", err)
	resp, err := api.Get(*u)
	t.Log("resp", resp, "err", err)
	t.Log("resp header", resp.Header)
	b, err := ioutil.ReadAll(resp.Body)
	t.Log("read err", err)
	t.Log("resp body", string(b))
	var maps map[string]interface{}
	json.Unmarshal(b, &maps)
	t.Log("decoded", maps)
}
