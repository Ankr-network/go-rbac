// Copyright 2019 The Ankr crop. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
/*
Package go_rbac provides RBAC client implementations.

create new client:
	c, err := New("127.0.0.1:6801")
    ...

authorize something:
   rsp, err := c.Authorize(&Request{
		Subject:  "ankr",
		Resource: "/v1/path/to/rc",
		Action:   "read",
	})
   ...
*/

package go_rbac
