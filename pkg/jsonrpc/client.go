package jsonrpc

import (
	"net/http"
)

type Client struct {
	client *http.Client
}

func (c *Client) Call(request Request) (Response, error) {
	return Response{}, nil
}
