package commons

import (
	"crypto/tls"

	"github.com/parnurzeal/gorequest"
)

const (
	contentType     = "Content-Type"
	contentTypeJson = "application/json"
	GET             = "GET"
	POST            = "POST"
	PUT             = "PUT"
)

type ConnectionFactory struct {
	BaseRequest *gorequest.SuperAgent
}

type Response gorequest.Response

var baseRequest = gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true}).Set(contentType, contentTypeJson)

func HttpGet(url string) *ConnectionFactory {
	return HttpRequest(GET, url, nil)
}

func HttpPost(url string, body interface{}) *ConnectionFactory {
	return HttpRequest(POST, url, body)
}

func HttpPut(url string, body interface{}) *ConnectionFactory {
	return HttpRequest(PUT, url, body)
}

func HttpRequest(method string, url string, body interface{}) *ConnectionFactory {
	req := baseRequest.Clone()

	switch method {
	case GET:
		req.Get(url)
	case POST:
		req.Post(url).Send(body)
	case PUT:
		req.Put(url).Send(body)
	}
	c := &ConnectionFactory{
		BaseRequest: req,
	}
	return c
}

func (c *ConnectionFactory) Headers(headers map[string]string) *ConnectionFactory {
	for k, v := range headers {
		c.BaseRequest.Set(k, v)
	}
	return c
}

func (c *ConnectionFactory) Header(name string, value string) *ConnectionFactory {
	c.BaseRequest.Set(name, value)
	return c
}
