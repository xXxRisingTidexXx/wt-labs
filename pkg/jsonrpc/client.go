package jsonrpc

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	client *http.Client
	url    string
}

func NewClient(client *http.Client, url string) *Client {
	return &Client{client, url}
}

func (c *Client) Call(request Request) Response {
	data, err := json.Marshal(request)
	if err != nil {
		return NewError(clientError{err})
	}
	response, err := c.client.Post(c.url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return NewError(clientError{err})
	}
	if response.StatusCode != http.StatusOK {
		_ = response.Body.Close()
		return NewError(notOKError{response.Status})
	}
	if request.IsNotification() {
		_ = response.Body.Close()
		return NewResult("OK")
	}
	r, e := ParseResponse(response.Body)
	if e != nil {
		_ = response.Body.Close()
		return NewError(e)
	}
	if err := response.Body.Close(); err != nil {
		return NewError(clientError{err})
	}
	if !r.HasError() && r.id != request.id {
		return NewError(idMismatch{request.id, r.id})
	}
	return r
}
