package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Caller interface {
	Get(url string) (resp *http.Response, err error)
}

type HTTP struct {
}

func (HTTP) Get(url string) (resp *http.Response, err error) {
	return http.Get(url)
}

func MakeHTTPCall(url string) (*Response, error) {
	return makeHTTPCall(HTTP{}, url)
}

func makeHTTPCall(client Caller, url string) (*Response, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r := &Response{}
	if err := json.Unmarshal(body, r); err != nil {
		return nil, err
	}
	return r, nil
}
