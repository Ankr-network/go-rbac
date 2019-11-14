package go_rbac

import (
	"context"

	"github.com/Ankr-network/go-rbac/rbac"
)

type Request struct {
	// the executor which execute the spec action
	Subject string
	// the resource such as  the name of data
	// or the name of url(/v1/path/to/role) and so on.
	Resource string
	// the action such as Read, Write, PUT, POST and so on,
	// it which exists in reality is valid and reasonable.
	Action string
}

type Client struct {
	c rbac.InternalRoleClient
}

// New create new client for remote rbac server
func New(srvAddr string) (*Client, error) {
	c, err := rbac.NewAnkrInternalRoleClient(srvAddr)
	if err != nil {
		return nil, err
	}
	return &Client{c}, nil
}

// Close ends this communication
func (c *Client) Close() error {
	return c.c.Close()
}

// judge result response
type Response struct {
	// true means passed.
	// false means failed.
	OK bool
}

// Authorize judge the request
func (c *Client) Authorize(req *Request) (*Response, error) {
	in := &rbac.AuthorizeRequest{
		Action:   req.Action,
		Subject:  req.Subject,
		Resource: req.Resource,
	}
	rsp, err := c.c.Authorize(context.Background(), in)
	if err != nil {
		return nil, err
	}
	return &Response{OK: rsp.Ok}, nil
}
