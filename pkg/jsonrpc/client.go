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
		return Response{error: clientError{err}, id: nullID{}}
	}
	response, err := c.client.Post(c.url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return Response{error: clientError{err}, id: nullID{}}
	}
	if response.StatusCode != http.StatusOK {
		_ = response.Body.Close()
		return Response{error: notOKError{response.Status}, id: nullID{}}
	}
	if request.IsNotification() {
		_ = response.Body.Close()
		return Response{result: "OK", id: notificationID{}}
	}
	r, e := ParseResponse(response.Body)
	if e != nil {
		_ = response.Body.Close()
		return Response{error: e, id: nullID{}}
	}
	if err := response.Body.Close(); err != nil {
		return Response{error: clientError{err}, id: nullID{}}
	}
	if !r.HasError() && r.id != request.id {
		return Response{error: idMismatch{request.id, r.id}, id: nullID{}}
	}
	return r
}
