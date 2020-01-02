// Copyright 2019 The Ankr crop. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
/*
Package go_rbac provides RBAC client implementations.

create new client:
	c, err := New("127.0.0.1:6801")
    ...

authorize something:
   rsp, err := c.Authorize(ctx, &Request{
		Subject:  "ankr",
		Resource: "/v1/path/to/rc",
		Action:   "read",
	})
   ...
*/
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

type Client interface {
	Close() error
	Authorize(ctx context.Context, req *Request) (*Response, error)
}

type client struct {
	c rbac.InternalRoleClient
}

// New create new client for remote rbac server
func New(srvAddr string) (Client, error) {
	c, err := rbac.NewAnkrInternalRoleClient(srvAddr)
	if err != nil {
		return nil, err
	}
	return &client{c}, nil
}

// Close ends this communication
func (c *client) Close() error {
	return c.c.Close()
}

// judge result response
type Response struct {
	// true means passed.
	// false means failed.
	OK bool
}

// Authorize judge the request
func (c *client) Authorize(ctx context.Context, req *Request) (*Response, error) {
	in := &rbac.AuthorizeRequest{
		Action:   req.Action,
		Subject:  req.Subject,
		Resource: req.Resource,
	}
	rsp, err := c.c.Authorize(ctx, in)
	if err != nil {
		return nil, err
	}
	return &Response{OK: rsp.Ok}, nil
}
