package rest

import (
	"encoding/json"

	"github.com/sendgrid/rest"
)

type clienRest struct {
	request rest.Request
	baseUrl string
	header  map[string]string
}

type ClientRest interface {
	Get(endpoint string) (*rest.Response, error)
	Post(endpoint string, body interface{}) (*rest.Response, error)
	Put(endpoint string, body interface{}) (*rest.Response, error)
	Delete(endpoint string, body interface{}) (*rest.Response, error)
}

func NewRequest(baseUrl string, header map[string]string) ClientRest {
	return &clienRest{
		header:  header,
		baseUrl: baseUrl,
	}
}

func (r *clienRest) Get(endpoint string) (*rest.Response, error) {
	method := rest.Get
	r.createRequest(endpoint, method, nil)
	return rest.Send(r.request)
}

func (r *clienRest) Post(endpoint string, body interface{}) (*rest.Response, error) {
	method := rest.Post
	r.createRequest(endpoint, method, body)
	return rest.Send(r.request)
}
func (r *clienRest) Put(endpoint string, body interface{}) (*rest.Response, error) {
	method := rest.Put
	r.createRequest(endpoint, method, body)
	return rest.Send(r.request)
}
func (r *clienRest) Delete(endpoint string, body interface{}) (*rest.Response, error) {
	method := rest.Delete
	r.createRequest(endpoint, method, body)
	return rest.Send(r.request)
}

func (r *clienRest) decodeInfo(body interface{}) []byte {

	data, _ := json.Marshal(&body)
	return data
}
func (r *clienRest) createRequest(endpoint string, method rest.Method, body interface{}) {
	r.request = rest.Request{
		Method:      method,
		BaseURL:     r.baseUrl + endpoint,
		Headers:     r.header,
		QueryParams: nil,
		Body:        r.decodeInfo(body),
	}
}
